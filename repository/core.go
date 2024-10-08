package repository

import (
	"sync"
	"websocket/entity"
)

type RoomRepository struct {
	rooms map[string][]*entity.Connection
	mu    sync.RWMutex
}

func NewRoomRepository() *RoomRepository {
	return &RoomRepository{
		rooms: make(map[string][]*entity.Connection),
	}
}

func (r *RoomRepository) AddConnection(roomID string, conn *entity.Connection) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.rooms[roomID] = append(r.rooms[roomID], conn)
}

func (r *RoomRepository) GetConnections(roomID string) []*entity.Connection {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.rooms[roomID]
}

func (r *RoomRepository) RemoveConnection(roomID string, conn *entity.Connection) {
	r.mu.Lock()
	defer r.mu.Unlock()
	connections := r.rooms[roomID]
	for i, c := range connections {
		if c == conn {
			r.rooms[roomID] = append(connections[:i], connections[i+1:]...)
			break
		}
	}
}

func (r *RoomRepository) RoomExists(roomID string) bool {
	r.mu.RLock()
	defer r.mu.RUnlock()
	_, exists := r.rooms[roomID]
	return exists
}
