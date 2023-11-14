package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	listLock    sync.RWMutex
	connections []connectionState
)

type websocketMessage struct {
	MessageType string `json:"message_type"`
	Data        string `json:"data"`
}

type connectionState struct {
	websocket *threadSafeWriter
}

type threadSafeWriter struct {
	*websocket.Conn
	sync.Mutex
}

func websocketHandler(w http.ResponseWriter, r *http.Request) {
	unsafeConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	conn := &threadSafeWriter{
		unsafeConn,
		sync.Mutex{},
	}
	defer conn.Close()

	listLock.Lock()
	connections = append(connections, connectionState{websocket: conn})
	listLock.Unlock()

	message := &websocketMessage{}
	for {
		_, raw, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		} else if err := json.Unmarshal(raw, &message); err != nil {
			log.Println(err)
			return
		}

		for _, c := range connections {
			c.websocket.WriteJSON(message)
		}
	}
}

func (t *threadSafeWriter) WriteJSON(v interface{}) error {
	t.Lock()
	defer t.Unlock()
	return t.Conn.WriteJSON(v)
}

type pageHandler struct {
	once     sync.Once
	filename string
	tmpl     *template.Template
}

func (p *pageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.once.Do(func() {
		p.tmpl = template.Must(template.ParseFiles(filepath.Join("templates", p.filename)))
	})
	p.tmpl.Execute(w, "Hello World")
}

func main() {
	http.HandleFunc("/websocket", websocketHandler)
	http.Handle("/", &pageHandler{filename: "index.html"})
	log.Fatal(http.ListenAndServe(":7001", nil))
}
