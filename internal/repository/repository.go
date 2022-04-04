package repository

import "github.com/s0a1qq/booking-mess/internal/models"

type DatabaseRepo interface {
	AllUsers() bool

	InsertReservation(res models.Reservation) error
}
