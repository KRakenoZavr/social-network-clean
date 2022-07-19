package repository

import (
	"database/sql"
	"errors"
	"log"

	"mux/internal/models"

	"github.com/satori/uuid"

	"mux/internal/user"
)

type usersRepo struct {
	db     *sql.DB
	logger *log.Logger
}

var (
	UserNoAuth = errors.New("user is not authenticated")
)

func NewUserRepository(db *sql.DB, logger *log.Logger) user.Repository {
	return &usersRepo{db: db, logger: logger}
}

func (r *usersRepo) Create(user *models.User) (uuid.UUID, error) {
	id := uuid.NewV4()

	query, err := r.db.Prepare(createUserQuery)
	if err != nil {
		return uuid.Nil, err
	}

	tx, err := r.db.Begin()
	if err != nil {
		return uuid.Nil, err
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

	return id, nil
}

func (r *usersRepo) CreateUserAuth(userAuth *models.UserAuth) error {
	query, err := r.db.Prepare(createUserAuthQuery)
	if err != nil {
		return err
	}

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Stmt(query).Exec(userAuth.UserID, userAuth.Expires, userAuth.Session)
	if err != nil {
		r.logger.Println("doing rollback")
		r.logger.Println(err.Error())
		tx.Rollback()
	} else {
		tx.Commit()
	}

	return nil
}

func (r *usersRepo) GetUserAuth(session string) (models.UserAuth, error) {
	var user models.UserAuth
	row := r.db.QueryRow(getUserAuthQuery, session)
	err := row.Scan(&user.UserID, &user.Expires, &user.Session)

	switch err {
	case sql.ErrNoRows:
		return models.UserAuth{}, UserNoAuth
	case nil:
		return user, nil
	default:
		return models.UserAuth{}, err
	}
}
