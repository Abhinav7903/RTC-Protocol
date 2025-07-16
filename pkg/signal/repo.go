package signal

import "rtc/factory"

type SignalRepository interface {
	CreateSignal(roomID, senderID, signalType string, payload []byte) (*factory.Signal, error)
	GetSignalsByRoom(roomID string) ([]factory.Signal, error)
	DeleteSignalsByRoom(roomID string) error
	GetSignalsByRoomExcludingSender(roomID, senderID string) ([]factory.Signal, error)
}
