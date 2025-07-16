package room

import "rtc/factory"

type RoomRepository interface {
	CreateRoom(name, roomType string) (*factory.Room, error)
	GetRoomByID(id string) (*factory.Room, error)
	DeleteRoom(id string) error
	ListRooms() ([]factory.Room, error)
}
