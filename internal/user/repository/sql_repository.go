package repository

import (
	"database/sql"

	"mux/internal/user"
)

type usersRepo struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) user.Repository {
	return &usersRepo{db: db}
}

func (r *usersRepo) Create() {
}
