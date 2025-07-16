package participant

import "rtc/factory"

type ParticipantRepository interface {
	CreateParticipant(roomID, displayName string) (*factory.Participant, error)
	GetParticipantsByRoom(roomID string) ([]factory.Participant, error)
	DeleteParticipant(id string) error
}
