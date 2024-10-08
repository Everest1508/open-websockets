package handlers

import (
	"log"
	"net/http"
	"websocket/entity"
	"websocket/repository"
	"websocket/service"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type WebSocketHandler struct {
	messageservice *service.MessageService
}

func NewWebSocketHandler() *WebSocketHandler {
	roomRepo := repository.NewRoomRepository()
	messageservice := service.NewMessageService(roomRepo)

	return &WebSocketHandler{
		messageservice: messageservice,
	}
}

func (h *WebSocketHandler) HandleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade:", err)
		return
	}
	room := r.URL.Query().Get("room")
	role := r.URL.Query().Get("role")
	if room == "" || role == "" {
		log.Println("Room or role not provided")
		conn.Close()
		return
	}

	connection := &entity.Connection{
		ID:   r.RemoteAddr,
		Room: room,
		Role: role,
		Conn: conn,
		Send: make(chan []byte),
	}

	h.messageservice.AddConnection(room, connection)

	go h.readMessages(connection)
	go h.writeMessages(connection)
}

func (h *WebSocketHandler) readMessages(conn *entity.Connection) {
	defer h.messageservice.RemoveConnection(conn.Room, conn)
	defer conn.Close()

	for {
		_, message, err := conn.Conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			return
		}

		h.messageservice.HandleMessage(conn, string(message))
	}
}

func (h *WebSocketHandler) writeMessages(conn *entity.Connection) {
	defer conn.Close()

	for msg := range conn.Send {
		if err := conn.Conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			log.Println("Write error:", err)
			return
		}

	}
}
