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
	NotFound = errors.New("no row found")
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

func (r *usersRepo) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	row := r.db.QueryRow(getUserByEmailQuery, email)
	err := row.Scan(&user.UserID, &user.Email, &user.Password,
		&user.FName, &user.LName, &user.DateOfBirth,
		&user.IsPrivate, &user.Avatar, &user.NickName, &user.About)

	switch err {
	case sql.ErrNoRows:
		return models.User{}, NotFound
	case nil:
		return user, nil
	default:
		return models.User{}, err
	}
}

func (r *usersRepo) CheckUserByEmail(email string) (bool, error) {
	var userId string
	row := r.db.QueryRow(getUserIDByEmailQuery, email)
	err := row.Scan(&userId)

	switch err {
	case sql.ErrNoRows:
		return false, nil
	case nil:
		return true, nil
	default:
		return false, err
	}
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
	err := row.Scan(&user.ID, &user.UserID, &user.Expires, &user.Session)

	switch err {
	case sql.ErrNoRows:
		return models.UserAuth{}, NotFound
	case nil:
		return user, nil
	default:
		return models.UserAuth{}, err
	}
}
