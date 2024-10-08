package service

import (
	"errors"
	"time"
	"websocket/entity"
	"websocket/repository"
)

type MessageService struct {
	RoomRepo *repository.RoomRepository
}

func NewMessageService(roomRepo *repository.RoomRepository) *MessageService {
	return &MessageService{
		RoomRepo: roomRepo,
	}
}

func (u *MessageService) HandleMessage(conn *entity.Connection, content string) (*entity.Message, error) {
	if conn.Role != "publisher" {
		return nil, errors.New("only publishers can send messages")
	}

	message := &entity.Message{
		SenderID:  conn.ID,
		Content:   content,
		Timestamp: time.Now().Unix(),
	}

	u.broadcastToRoom(conn.Room, message.Content)

	return message, nil
}

func (u *MessageService) broadcastToRoom(roomID string, message string) {
	connections := u.RoomRepo.GetConnections(roomID)
	for _, conn := range connections {
		if conn.Role == "subscriber" {
			conn.Send <- []byte(message)
		}
	}
}

func (u *MessageService) AddConnection(roomID string, conn *entity.Connection) {
	u.RoomRepo.AddConnection(roomID, conn)
}

func (u *MessageService) RemoveConnection(roomID string, conn *entity.Connection) {
	u.RoomRepo.RemoveConnection(roomID, conn)
}
