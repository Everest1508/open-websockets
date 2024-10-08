package main

import (
	"log"
	"net/http"
	"websocket/handlers"
)

func main() {
	wsHandler := handlers.NewWebSocketHandler()

	http.HandleFunc("/ws", wsHandler.HandleConnection)

	log.Println("WebSocket server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
