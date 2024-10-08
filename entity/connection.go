package entity

import (
	"github.com/gorilla/websocket"
)

type Connection struct {
	ID   string
	Room string 
	Role string // "publisher" or "subscriber"
	Conn *websocket.Conn
	Send chan []byte
}

func (c *Connection) Close() error {
	return c.Conn.Close()
}
