package repositories

import (
	"api/src/models"
	"database/sql"
)

type usuarios struct {
	db *sql.DB
}

func NewRepositoryUsers(db *sql.DB) *usuarios {
	return &usuarios{db} // Return a new instance of the repository with the database connection
}

func (u usuarios) Create(usuario models.User) (uint64, error) {
	return 0, nil
}
