// File for processing the web requests.
// All the things that decode incoming web request or even know about the web
// should be here.  Things like parsing errors or figuring out what HTTP code
// to send should reside here. Beyond this barrier the logical layer of the app
// should not be concerned with how the requests came in.
package main

import (
	"encoding/json"
	"html/template"
	"io"
	"net/http"
	"strconv"

	"github.com/google/generative-ai-go/genai"
	"github.com/rs/cors" // Import the cors package
)

var chatSession *genai.ChatSession

// start a webserver.  Address is passed direclty to http library
func listenAndServe(address string) error {
	r := http.NewServeMux()

	logger.Println("Starting the server")

	r.HandleFunc("/guests/{id}", oneGuestHandler)
	r.HandleFunc("/guests", allGuestHandler)

	r.HandleFunc("/chat", chatHandler)
	r.HandleFunc("/llmodel", llmModelHandler)

	fs := http.FileServer(http.Dir("html/"))
	r.Handle("/", http.StripPrefix("/", fs))

	// Create a new CORS handler
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173",
								   "http://localhost:8080"}, // Allow your frontend origin
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Allow these methods
		AllowedHeaders:   []string{"*"}, // Allow all headers
		AllowCredentials: true, // If you need to send cookies, set this to true
	})
	handler := c.Handler(r)


	logger.Printf("Starting server on %s", address)
	return http.ListenAndServe(address, handler)
}

func chatHandler(w http.ResponseWriter, r *http.Request) {
	logger.Println("chat handler")
	ctx := r.Context()

	// TODO eventually expand to be multi session
	if chatSession == nil {
		chatSession = model.StartChat()
	}

	switch r.Method {
	case http.MethodGet:

		err := writeHistory(w)
		if err != nil {
			logger.Printf("Error writing history: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
	case http.MethodPost:

		content := r.FormValue("chat_message")

		userMsg := NewMessage(User, content)
		ChatHistory = append(ChatHistory, userMsg)

		resp, err := aiChat(ctx, chatSession, content)
		if err != nil {
			logger.Printf("Error getting a response from gemini: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		aiMsg := NewMessage(Agent, resp)
		ChatHistory = append(ChatHistory, aiMsg)

		err = writeHistory(w)
		if err != nil {
			logger.Printf("Error writing history: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func writeHistory(w http.ResponseWriter) error {
	for _, msg := range ChatHistory {
		template := "html/components/"
		if msg.Role == User {
			template += "user_msg.html"
		} else {
			template += "ai_msg.html"
		}

		err := writeChatMessage(w, template, msg)
		if err != nil {
			logger.Printf("Error writing template: %v", err)
			return err
		}
	}
	return nil
}

func llmModelHandler(w http.ResponseWriter, r *http.Request) {
	logger.Println("large language model handler")

	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(GetLangModel()))
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// oneGuestHandler handles requests to the /guests/{id} endpoint.
// It supports GET (retrieve a guest), PUT (update a guest), and DELETE (delete a guest) methods.
// Other methods will return a 405 Method Not Allowed error.
func oneGuestHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	logger.Printf("%s guest. ID: %s\n", r.Method, id)

	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		guest, err := oneGuest(uint(uid))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(guest)
		w.WriteHeader(http.StatusOK)
	case http.MethodDelete:
		err := deleteGuest(uint(uid))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusNoContent) // 204 for successful delete
	case http.MethodPut:
		var guest Guest
		err := json.NewDecoder(r.Body).Decode(&guest)
		logger.Printf("guest %v\n", guest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		guest.ID = uint(uid)
		err = updateGuest(&guest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// allGuestHandler handles requests to the /guests endpoint.
// It supports GET (list all guests) and POST (create a new guest) methods.
// Other methods will return a 405 Method Not Allowed error.
func allGuestHandler(w http.ResponseWriter, r *http.Request) {
	logger.Printf("%s guests\n", r.Method)

	switch r.Method {

	// get all the guests
	case http.MethodGet:
		guests, err := allGuests()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(guests)
		w.WriteHeader(http.StatusOK)

	// create a new guest
	case http.MethodPost:
		var guest Guest
		err := json.NewDecoder(r.Body).Decode(&guest)
		logger.Printf("guest %v\n", guest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = createGuest(&guest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}


func writeChatMessage(wr io.Writer, name string, data* Message) error {
	tmpl, err := template.ParseFiles(name)
	if err != nil {
		return err
	}
	err = tmpl.Execute(wr, data)
	return err
}