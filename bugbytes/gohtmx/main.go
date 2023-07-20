package main

import (
	"fmt"
	"net/http"
	"text/template"
	"time"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	fmt.Println("Hello")

	http.HandleFunc("/", h1)
	http.HandleFunc("/add-film/", h2)
	http.ListenAndServe(":3000", nil)
}

func h1(w http.ResponseWriter, r *http.Request) {
	// io.WriteString(w, "Hello, world!")
	// io.WriteString(w, r.Method)

	// fmt.Fprintf(w, "Hi, world!\n")
	// fmt.Fprintf(w, r.Method+"\n")

	films := map[string][]Film{
		"films": {
			{Title: "The Godfather", Director: "Francis Ford Coppola"},
			{Title: "Blade Runner", Director: "Ridley Scott"},
			{Title: "The Thing", Director: "John Carpenter"},
		},
	}

	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, films)
}

func h2(w http.ResponseWriter, r *http.Request) {
	// log.Print("HTMX request received")
	// log.Print(r.Header.Get("HX-Request"))

	time.Sleep(time.Second)
	title := r.PostFormValue("title")
	director := r.PostFormValue("director")
	// log.Print(title)
	// log.Print(director)

	// htmlStr := fmt.Sprintf("<li class='list-group-item bg-primary text-white'>%s - %s</li>", title, director)
	// tmpl, _ := template.New("t").Parse(htmlStr)
	// tmpl.Execute(w, nil)

	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.ExecuteTemplate(w, "film-list-element", Film{Title: title, Director: director})
}
