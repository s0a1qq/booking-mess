package dbrepo

import (
	"context"
	"time"

	"github.com/s0a1qq/booking-mess/internal/models"
)

func (m *postgresDBRepo) AllUsers() bool {
	return true
}

// InsertReservation Inserts a reservation into the database
func (m *postgresDBRepo) InsertReservation(res models.Reservation) (int, error) {

	cts, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var newID int

	stmt := `insert into reservations 
	(first_name, last_name, email, phone, start_date, end_date, room_id, created_at, updated_at)
	values ($1, $2, $3, $4, $5, $6, $7, $8, $9) returning id`

	err := m.DB.QueryRowContext(cts, stmt,
		res.FirstName,
		res.LastName,
		res.Email,
		res.Phone,
		res.StardDate,
		res.EndDate,
		res.RoomID,
		time.Now(),
		time.Now(),
	).Scan(&newID)

	if err != nil {
		return 0, err
	}

	return newID, nil
}

// InsertRoomRestriction Inserts a room_restrictions into the database
func (m *postgresDBRepo) InsertRoomRestriction(r models.RoomRestriction) (int, error) {

	cts, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var newID int

	stmt := `insert into room_restrictions
	(start_date, end_date, room_id, reservation_id, created_at, updated_at, restriction_id)
	values ($1, $2, $3, $4, $5, $6, $7) returning id`

	err := m.DB.QueryRowContext(cts, stmt,
		r.StardDate,
		r.EndDate,
		r.RoomID,
		r.ReservationID,
		time.Now(),
		time.Now(),
		r.RestrictionID,
	).Scan(&newID)

	if err != nil {
		return 0, err
	}

	return newID, nil
}

// SearchAvailabilityByDatesByRoomID return true if availability exist for room_id and false if no availability exist
func (m *postgresDBRepo) SearchAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error) {

	cts, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var numRows int

	stmt := `
	select 
		count(id)
	from 
		room_restrictions
	where 
		room_id = $1
		and $2 < end_date and $3 > start_date`

	row := m.DB.QueryRowContext(cts, stmt,
		roomID,
		start,
		end,
	)

	err := row.Scan(&numRows)
	if err != nil {
		return false, err
	}

	if numRows == 0 {
		return true, nil
	}
	return false, nil
}

// SearchAvailabilityByDatesForAllRooms return a slive of available rooms if any, for given date range
func (m *postgresDBRepo) SearchAvailabilityByDatesForAllRooms(start, end time.Time) ([]models.Room, error) {

	cts, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var rooms []models.Room

	query := `
	select
		r.id,
		r.room_name
	from
		rooms r
	where
		r.id not in (
		select
			rr.room_id
		from
			room_restrictions rr
		where
			$1 < rr.end_date
			and $2 > rr.start_date)
	`

	rows, err := m.DB.QueryContext(cts, query,
		start,
		end,
	)
	if err != nil {
		return rooms, err
	}

	for rows.Next() {
		var room models.Room
		err := rows.Scan(
			&room.ID,
			&room.RoomName,
		)
		if err != nil {
			return rooms, err
		}

		rooms = append(rooms, room)
	}

	if err = rows.Err(); err != nil {
		return rooms, err
	}

	return rooms, nil
}

//GetRoomByID get room by ID
func (m *postgresDBRepo) GetRoomByID(ID int) (models.Room, error) {
	cts, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var room models.Room
	query := `
	select
		r.*
	from
		rooms r
	where
		r.id = $1	
`

	row := m.DB.QueryRowContext(cts, query, ID)
	err := row.Scan(
		&room.ID,
		&room.RoomName,
		&room.CreatedAt,
		&room.UpdatetAt,
	)

	if err != nil {
		return room, err
	}

	return room, nil
}
