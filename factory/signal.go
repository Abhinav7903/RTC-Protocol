package factory

import (
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
