package entity

type EventType string

const (
	EventMessageReceived  EventType = "MESSAGE_RECEIVED"
	EventConnectionClosed EventType = "CONNECTION_CLOSED"
	EventErrorOccurred    EventType = "ERROR_OCCURRED"
)

type Event struct {
	Type EventType   `json:"type"`
	Data interface{} `json:"data"`
}
