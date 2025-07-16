package factory

import "time"

type Participant struct {
	ID          int       `json:"id"`
	RoomID      int       `json:"room_id"`
	DisplayName string    `json:"display_name"`
	JoinedAt    time.Time `json:"joined_at"`
}
