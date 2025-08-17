package database

import (
	"api/src/config"
	"database/sql"
)

func Connect() (*sql.DB, error) {
	// Open a connection to the database
	db, err := sql.Open("mysql", config.StringConnection)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil { // Check if the connection is working
		return nil, err
	}

	return db, nil
}
