package database

import (
	"database/sql"
)

var DB *sql.DB

func InitializeDatabase() error {

	var err error

	DB, err = sql.Open("postgres", "postgresql://localhost/dcp")

	if err != nil {
		return err
	}

	return DB.Ping()
}
