package postgres

import (
	"rtc/factory"
)

func (p *Postgres) CreateRoom(name, roomType string) (*factory.Room, error) {
	query := `INSERT INTO rooms (name, room_type) VALUES ($1, $2) RETURNING id, name, room_type, created_at`
	row := p.dbConn.QueryRow(query, name, roomType)

	var room factory.Room
	err := row.Scan(&room.ID, &room.Name, &room.RoomType, &room.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &room, nil
}

func (p *Postgres) GetRoomByID(id string) (*factory.Room, error) {
	query := `SELECT id, name, room_type, created_at FROM rooms WHERE id = $1`
	row := p.dbConn.QueryRow(query, id)

	var room factory.Room
	err := row.Scan(&room.ID, &room.Name, &room.RoomType, &room.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &room, nil
}

func (p *Postgres) DeleteRoom(id string) error {
	query := `DELETE FROM rooms WHERE id = $1`
	_, err := p.dbConn.Exec(query, id)
	return err
}

func (p *Postgres) ListRooms() ([]factory.Room, error) {
	query := `SELECT id, name, room_type, created_at FROM rooms ORDER BY created_at DESC`
	rows, err := p.dbConn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rooms []factory.Room
	for rows.Next() {
		var room factory.Room
		err := rows.Scan(&room.ID, &room.Name, &room.RoomType, &room.CreatedAt)
		if err != nil {
			return nil, err
		}
		rooms = append(rooms, room)
	}
	return rooms, nil
}
