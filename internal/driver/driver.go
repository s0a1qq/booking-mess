package driver

import (
	"database/sql"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

//DB holds the db connection pool
type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

const maxOpenDbConn = 10
const maxIdleDbConn = 5
const maxDbLifeTime = 5 * time.Minute

//ConnectSQL creates db pool for postgres
func ConnectSQL(dsn string) (*DB, error) {
	d, err := NewDB(dsn)
	if err != nil {
		panic(err)
	}

	d.SetMaxOpenConns(maxOpenDbConn)
	d.SetConnMaxIdleTime(maxIdleDbConn)
	d.SetConnMaxLifetime(maxDbLifeTime)

	dbConn.SQL = d

	err = TestDB(dbConn.SQL)
	if err != nil {
		return nil, err
	}

	return dbConn, nil
}

//TestDB tries to ping to DB
func TestDB(d *sql.DB) error {
	err := d.Ping()
	if err != nil {
		return err
	}

	return nil
}

//NewDB creates new DB for app
func NewDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = TestDB(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}
