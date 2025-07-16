package factory

import (
	"encoding/json"
	"time"
)

type Signal struct {
	ID            int       `json:"id"`
	RoomID        int       `json:"room_id"`
	SenderID      int       `json:"sender_id"`
	SignalType    string    `json:"signal_type"`    // "offer", "answer", "candidate"
	SignalPayload []byte    `json:"signal_payload"` // JSONB
	CreatedAt     time.Time `json:"created_at"`
}

type SignalRequest struct {
	RoomID     int             `json:"room_id"`   // changed from string
	SenderID   int             `json:"sender_id"` // changed from string
	SignalType string          `json:"signal_type"`
	Payload    json.RawMessage `json:"payload"`
}
