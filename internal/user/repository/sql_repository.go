package repository

import (
	"database/sql"
	"github.com/satori/uuid"
	"mux/internal/models"

	"mux/internal/user"
)

type usersRepo struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) user.Repository {
	return &usersRepo{db: db}
}

func (r *usersRepo) Create(user *models.User) error {
	id := uuid.NewV4()
	_, err := r.db.Exec(createUser, id, user.Username, user.Name, user.Password, user.Age)
	if err != nil {
		return err
	}
	return nil
}
