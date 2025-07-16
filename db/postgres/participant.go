package postgres

import (
	"rtc/factory"
)

// CreateParticipant inserts a participant into a room
func (p *Postgres) CreateParticipant(roomID int, displayName string) (*factory.Participant, error) {
	query := `INSERT INTO room_participants (room_id, display_name) 
	          VALUES ($1, $2) 
	          RETURNING id, room_id, display_name, joined_at`
	row := p.dbConn.QueryRow(query, roomID, displayName)

	var participant factory.Participant
	err := row.Scan(&participant.ID, &participant.RoomID, &participant.DisplayName, &participant.JoinedAt)
	if err != nil {
		return nil, err
	}
	return &participant, nil
}

// GetParticipantsByRoom lists all participants for a given room
func (p *Postgres) GetParticipantsByRoom(roomID int) ([]factory.Participant, error) {
	query := `SELECT id, room_id, display_name, joined_at 
	          FROM room_participants 
	          WHERE room_id = $1 
	          ORDER BY joined_at ASC`
	rows, err := p.dbConn.Query(query, roomID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var participants []factory.Participant
	for rows.Next() {
		var pData factory.Participant
		err := rows.Scan(&pData.ID, &pData.RoomID, &pData.DisplayName, &pData.JoinedAt)
		if err != nil {
			return nil, err
		}
		participants = append(participants, pData)
	}
	return participants, nil
}

// DeleteParticipant removes a participant by ID
func (p *Postgres) DeleteParticipant(id int) error {
	query := `DELETE FROM room_participants WHERE id = $1`
	_, err := p.dbConn.Exec(query, id)
	return err
}
