package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	key   = []byte("keyboard cat")
	store = sessions.NewCookieStore(key)
)

func secret(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	fmt.Fprintln(w, "The cake is a lie!")
}

func main() {
	http.HandleFunc("/secret", secret)
	http.ListenAndServe(":8080", nil)
}
