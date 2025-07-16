package factory

import "time"

type Room struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	RoomType  string    `json:"room_type"`
	CreatedAt time.Time `json:"created_at"`
}
