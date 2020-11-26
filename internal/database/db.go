package database

import (
	"backendo-go/internal/config"
	"errors"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3" // adding sqlite driver support
)

// NewDatabase ...
func NewDatabase(conf *config.Config) (*sqlx.DB, error) {
	switch conf.DB.Type {
	case "sqlite3":
		db, err := sqlx.Open(conf.DB.Driver, conf.DB.Conn)
		if err != nil {
			return nil, err
		}
		db.Ping()
		if err != nil {
			return nil, err
		}

		return db, nil

	default:
		return nil, errors.New("Invalid db type")
	}
}
