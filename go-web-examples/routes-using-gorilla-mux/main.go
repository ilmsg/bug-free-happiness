package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You're requested the book: %s on page %s\n", title, page)
	})
	http.ListenAndServe(":8080", r)

}

// http://localhost:8080/books/go-programming-blueprint/page/10

//
// features of the gorilla/mux Router
//

// r.HandleFunc("/books/{title}", CreateBook).Methods("POST")
// r.HandleFunc("/books/{title}", ReadBook).Methods("GET")
// r.HandleFunc("/books/{title}", UpdateBook).Methods("PUT")
// r.HandleFunc("/books/{title}", DeleteBook).Methods("DELETE")

// Hostname & Subdomains
// r.HandleFunc("/books/{title}", BookHandler).Host("www.mybookstore.com")

// Schemes
// r.HandleFunc("/secure", SecureHandler).Schemes("https")
// r.HandleFunc("/insecure", IncsecureHandler).Schemes("http")

// Path Prefixes & Subrouters
// bookrouter := r.PathPrefix("/books").Subrouter()
// bookrouter.HandleFunc("/", AllBook)
// bookrouter.HandleFunc("/{title}", GetBook)
