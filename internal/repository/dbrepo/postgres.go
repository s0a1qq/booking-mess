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
func (m *postgresDBRepo) InsertReservation(res models.Reservation) error {

	cts, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `insert into reservations 
	(first_name, last_name, email, phone, start_date, end_date, room_id, created_at, updated_at)
	values ($1, $2, $3, $4, $5, $6, $7, $8, $9) returning id`

	_, err := m.DB.ExecContext(cts, stmt,
		res.FirstName,
		res.LastName,
		res.Email,
		res.Phone,
		res.StardDate,
		res.EndDate,
		res.RoomID,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return err
	}

	return nil
}
