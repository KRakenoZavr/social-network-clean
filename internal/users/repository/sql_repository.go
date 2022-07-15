package repository

import (
	"database/sql"

	"mux/internal/users"
)

type usersRepo struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) users.Repository {
	return &usersRepo{db: db}
}

func (r *usersRepo) Create() {
}
