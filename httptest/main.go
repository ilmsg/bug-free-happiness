package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	http.HandleFunc("/upper", upperCaseHandler)
	log.Fatal(http.ListenAndServe(":7009", nil))
}

// Req: http://localhost:7009/upper?word=Cat
// Res: CAT
func upperCaseHandler(w http.ResponseWriter, r *http.Request) {
	query, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "invalid request")
		return
	}

	word := query.Get("word")
	if len(word) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "missing word")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, strings.ToUpper(word))
}
