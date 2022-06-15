package models

import (
	"time"
)

//User is user model
type User struct {
	ID          int
	FirstName   string
	LastName    string
	Email       string
	Password    string
	AccessLevel int
	CreatedAt   time.Time
	UpdatetAt   time.Time
}

//Rooms is Room model
type Room struct {
	ID        int
	RoomName  string
	CreatedAt time.Time
	UpdatetAt time.Time
}

//Restriction is restriction model
type Restriction struct {
	ID              int
	RestrictionName string
	CreatedAt       time.Time
	UpdatetAt       time.Time
}

//Reservation is reservation model
type Reservation struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Phone     string
	StartDate time.Time
	EndDate   time.Time
	RoomID    int
	CreatedAt time.Time
	UpdatetAt time.Time
	Room      Room
}

//RoomRestriction is room restriction model
type RoomRestriction struct {
	ID            int
	StartDate     time.Time
	EndDate       time.Time
	RoomID        int
	ReservationID int
	RestrictionID int
	CreatedAt     time.Time
	UpdatetAt     time.Time
	Room          Room
	Reservation   Reservation
	Restriction   Restriction
}

//Mail data holds an email message
type MailData struct {
	To       string
	From     string
	Subject  string
	Content  string
	Template string
}
