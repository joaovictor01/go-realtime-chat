package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joaovictor01/go-realtime-chat/backend/pkg/websocket"
)

func serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)

	// upgrade this connection to a WebSocket connection
	ws, err := websocket.Upgrade(w, r, nil)
	if err != nil {
		log.Fprintf(w, "%+V\n", err)
	}
	go websocket.Writer(ws)
	websocket.Reader(ws)

}

func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Simple Server")
	})

	http.HandleFunc("/ws", serveWs)
}

func main() {
	fmt.Println("Chat App v1.0")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}
