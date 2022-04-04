package dbrepo

import (
	"database/sql"

	"github.com/s0a1qq/booking-mess/internal/config"
	"github.com/s0a1qq/booking-mess/internal/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresDBRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}
