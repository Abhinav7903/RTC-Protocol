package factory

import "time"

type Participant struct {
	ID          string    `json:"id"`
	RoomID      string    `json:"room_id"`
	DisplayName string    `json:"display_name"`
	JoinedAt    time.Time `json:"joined_at"`
}
