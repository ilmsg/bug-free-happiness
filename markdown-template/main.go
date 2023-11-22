package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/russross/blackfriday"
)

type Page struct {
	Title string
	Body  string
}

func main() {
	http.HandleFunc("/", pageHandle)
	http.ListenAndServe(":7000", nil)
}

func markDowner(args ...interface{}) template.HTML {
	s := blackfriday.MarkdownCommon([]byte(fmt.Sprintf("%s", args...)))
	return template.HTML(s)
}

func pageHandle(w http.ResponseWriter, r *http.Request) {
	body, err := os.ReadFile("page.md")
	if err != nil {
		log.Fatal(err)
	}

	p := &Page{Title: "Page Title", Body: string(body)}

	tmpl := template.
		Must(template.
			New("page.html").
			Funcs(template.FuncMap{"markDown": markDowner}).
			ParseFiles("page.html"))

	if err := tmpl.ExecuteTemplate(w, "page.html", p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
