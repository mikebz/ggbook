// File for processing the web requests.
// All the things that decode incoming web request or even know about the web
// should be here.
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// start a webserver.  Address is passed direclty to http library
func listenAndServe(address string) error {
	r := http.NewServeMux()

	fmt.Println("Starting the server")
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/guests/{id}", oneGuestHandler)
	r.HandleFunc("/guests", allGuestHandler)

	return http.ListenAndServe(address, r)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Go Guest Book!\n")
}

// oneGuestHandler handles requests to the /guests/{id} endpoint.
// It supports GET (retrieve a guest), PUT (update a guest), and DELETE (delete a guest) methods.
// Other methods will return a 405 Method Not Allowed error.
func oneGuestHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Fprintf(w, "%s guest. ID: %s\n", r.Method, id)

	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	switch r.Method {
	case "GET":
		guest, err := oneGuest(uint(uid))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(guest)
	case "DELETE":
		err := deleteGuest(uint(uid))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
	case "PUT":
		var guest Guest
		err := json.NewDecoder(r.Body).Decode(&guest)
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

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// allGuestHandler handles requests to the /guests endpoint.
// It supports GET (list all guests) and POST (create a new guest) methods.
// Other methods will return a 405 Method Not Allowed error.
func allGuestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s guests\n", r.Method)

	switch r.Method {

	// get all the guests
	case "GET":
		guests, err := allGuests()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(guests)

	// create a new guest
	case "POST":
		var guest Guest
		err := json.NewDecoder(r.Body).Decode(&guest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = createGuest(&Guest{})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
