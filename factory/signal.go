package factory

import (
	"encoding/json"
	"time"
)

type Signal struct {
	ID            int       `json:"id"`
	RoomID        string    `json:"room_id"`
	SenderID      string    `json:"sender_id"`
	SignalType    string    `json:"signal_type"`    // "offer", "answer", or "candidate"
	SignalPayload []byte    `json:"signal_payload"` // raw JSON
	CreatedAt     time.Time `json:"created_at"`
}

type SignalRequest struct {
	RoomID     string          `json:"room_id"`
	SenderID   string          `json:"sender_id"`
	SignalType string          `json:"signal_type"`
	Payload    json.RawMessage `json:"payload"`
}
