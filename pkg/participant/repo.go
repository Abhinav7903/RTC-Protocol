package participant

import "rtc/factory"

type ParticipantRepository interface {
	CreateParticipant(roomID int, displayName string) (*factory.Participant, error)
	GetParticipantsByRoom(roomID int) ([]factory.Participant, error)
	DeleteParticipant(id int) error
}
