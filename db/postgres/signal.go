package postgres

import (
	"rtc/factory"
)

// CreateSignal stores an offer/answer/candidate for a room
func (p *Postgres) CreateSignal(roomID, senderID, signalType string, payload []byte) (*factory.Signal, error) {
	query := `INSERT INTO signals (room_id, sender_id, signal_type, signal_payload) 
	          VALUES ($1, $2, $3, $4) 
	          RETURNING id, room_id, sender_id, signal_type, signal_payload, created_at`

	row := p.dbConn.QueryRow(query, roomID, senderID, signalType, payload)

	var signal factory.Signal
	err := row.Scan(&signal.ID, &signal.RoomID, &signal.SenderID, &signal.SignalType, &signal.SignalPayload, &signal.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &signal, nil
}

// GetSignalsByRoom fetches all signals for a room (useful for joining late or syncing)
func (p *Postgres) GetSignalsByRoom(roomID string) ([]factory.Signal, error) {
	query := `SELECT id, room_id, sender_id, signal_type, signal_payload, created_at 
	          FROM signals 
	          WHERE room_id = $1 
	          ORDER BY created_at ASC`

	rows, err := p.dbConn.Query(query, roomID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var signals []factory.Signal
	for rows.Next() {
		var signal factory.Signal
		err := rows.Scan(&signal.ID, &signal.RoomID, &signal.SenderID, &signal.SignalType, &signal.SignalPayload, &signal.CreatedAt)
		if err != nil {
			return nil, err
		}
		signals = append(signals, signal)
	}
	return signals, nil
}

// DeleteSignalsByRoom optionally clean up all signals when room is removed (though CASCADE already does this)
func (p *Postgres) DeleteSignalsByRoom(roomID string) error {
	query := `DELETE FROM signals WHERE room_id = $1`
	_, err := p.dbConn.Exec(query, roomID)
	return err
}

func (p *Postgres) GetSignalsByRoomExcludingSender(roomID, senderID string) ([]factory.Signal, error) {
	query := `
	SELECT id, room_id, sender_id, signal_type, signal_payload, created_at
	FROM signals
	WHERE room_id = $1 AND sender_id != $2
	ORDER BY created_at ASC
	`
	rows, err := p.dbConn.Query(query, roomID, senderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var signals []factory.Signal
	for rows.Next() {
		var s factory.Signal
		if err := rows.Scan(&s.ID, &s.RoomID, &s.SenderID, &s.SignalType, &s.SignalPayload, &s.CreatedAt); err != nil {
			return nil, err
		}
		signals = append(signals, s)
	}
	return signals, nil
}
