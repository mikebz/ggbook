package main

import (
	"fmt"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Create a new GORM model
type Guest struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	db, err := gorm.Open(sqlite.Open("main.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Migration")
	err = db.AutoMigrate(&Guest{})
	if err != nil {
		panic("failed to migrate")
	}

	r := http.NewServeMux()

	fmt.Println("Starting the server")
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/guests/{id}", getGuestHandler)
	r.HandleFunc("/guests", guestHandler)

	http.ListenAndServe(":8080", r)
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
