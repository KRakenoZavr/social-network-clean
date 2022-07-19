package repository

import (
	"database/sql"
	"log"

	"mux/internal/models"

	"github.com/satori/uuid"

	"mux/internal/user"
)

type usersRepo struct {
	db     *sql.DB
	logger *log.Logger
}

func NewUserRepository(db *sql.DB, logger *log.Logger) user.Repository {
	return &usersRepo{db: db, logger: logger}
}

func (r *usersRepo) Create(user *models.User) error {
	id := uuid.NewV4()

	query, err := r.db.Prepare(createUser)
	if err != nil {
		return err
	}

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Stmt(query).Exec(id, user.Email, user.Password,
		user.FName, user.LName, user.DateOfBirth, user.IsPrivate,
		user.Avatar, user.NickName, user.About)
	if err != nil {
		r.logger.Println("doing rollback")
		r.logger.Println(err.Error())
		tx.Rollback()
	} else {
		tx.Commit()
	}

	return nil
}
