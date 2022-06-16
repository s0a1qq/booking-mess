package dbrepo

import (
	"context"
	"errors"
	"time"

	"github.com/s0a1qq/booking-mess/internal/models"
	"golang.org/x/crypto/bcrypt"
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
		res.StartDate,
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
		r.StartDate,
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

// SearchAvailabilityForAllRooms return a slive of available rooms if any, for given date range
func (m *postgresDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {

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

//GetUserByID returns a user by id
func (m *postgresDBRepo) GetUserByID(id int) (models.User, error) {
	cts, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, first_name, last_name, email, password, access_level, created_at, updated_at
			from users where id = $1`

	row := m.DB.QueryRowContext(cts, query, id)

	var u models.User

	err := row.Scan(
		&u.ID,
		&u.FirstName,
		&u.LastName,
		&u.Email,
		&u.Password,
		&u.AccessLevel,
		&u.CreatedAt,
		&u.UpdatetAt,
	)

	if err != nil {
		return u, err
	}

	return u, nil
}

//UpdateUser returns a user by id
func (m *postgresDBRepo) UpdateUser(u models.User) error {
	cts, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `update users set first_name=$1, last_name=$2, email=$3, access_level=$4, updated_at=$5
	where id = $6`

	_, err := m.DB.ExecContext(cts, query,
		u.FirstName,
		u.LastName,
		u.Email,
		u.AccessLevel,
		time.Now(),
		u.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

//Authenticate authenticates a user
func (m *postgresDBRepo) Authenticate(email, testPassword string) (int, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var id int
	var hashedPassword string

	row := m.DB.QueryRowContext(ctx, "select id, password from users where email = $1", email)
	err := row.Scan(&id, &hashedPassword)
	if err != nil {
		return id, hashedPassword, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(testPassword))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return 0, "", errors.New("incorrect password")
	} else if err != nil {
		return 0, "", err
	}

	return id, hashedPassword, nil
}
