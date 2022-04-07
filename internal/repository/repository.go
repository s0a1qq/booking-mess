package repository

import (
	"time"

	"github.com/s0a1qq/booking-mess/internal/models"
)

type DatabaseRepo interface {
	AllUsers() bool

	InsertReservation(res models.Reservation) (int, error)
	InsertRoomRestriction(r models.RoomRestriction) (int, error)
	SearchAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error)
	SearchAvailabilityByDatesForAllRooms(start, end time.Time) ([]models.Room, error)
	GetRoomByID(ID int) (models.Room, error)
}
