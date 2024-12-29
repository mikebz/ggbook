package main

import (
	"fmt"
	"net/http"
	"os"
)

// start a webserver. Use environment variables
// SERVER and PORT in order to specify the address to listen on
func listenAndServe() error {
	r := http.NewServeMux()

	server := os.Getenv("SERVER")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Starting the server")
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/guests/{id}", getGuestHandler)
	r.HandleFunc("/guests", guestHandler)

	return http.ListenAndServe(server+":"+port, r)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Go Guest Book!\n")
}

// getGuestHandler handles requests to the /guests/{id} endpoint.
// It supports GET (retrieve a guest), PUT (update a guest), and DELETE (delete a guest) methods.
// Other methods will return a 405 Method Not Allowed error.
func getGuestHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Fprintf(w, "%s guest. ID: %s\n", r.Method, id)

	switch r.Method {
	case "GET":
	case "PUT":
	case "DELETE":
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// guestHandler handles requests to the /guests endpoint.
// It supports GET (list all guests) and POST (create a new guest) methods.
// Other methods will return a 405 Method Not Allowed error.
func guestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s guests\n", r.Method)

	switch r.Method {
	case "GET":
	case "POST":
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
