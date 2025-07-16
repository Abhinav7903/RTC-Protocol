package signal

import (
	"rtc/factory"
)

type SignalRepository interface {
	CreateSignal(roomID, senderID int, signalType string, payload []byte) (*factory.Signal, error)
	GetSignalsByRoom(roomID int) ([]factory.Signal, error)
	DeleteSignalsByRoom(roomID int) error
	GetSignalsByRoomExcludingSender(roomID int, senderID int) ([]factory.Signal, error)
}
