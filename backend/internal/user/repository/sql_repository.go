package repository

import (
	"database/sql"
	"errors"
	"log"

	"mux/internal/models"
	"mux/internal/user/dto"

	"github.com/satori/uuid"

	"mux/internal/user"
)

type usersRepo struct {
	db     *sql.DB
	logger *log.Logger
}

var NotFound = errors.New("no row found")

func NewRepository(db *sql.DB, logger *log.Logger) user.Repository {
	return &usersRepo{db: db, logger: logger}
}

func (r *usersRepo) Create(user *models.User) (models.User, error) {
	id := uuid.NewV4()

	query, err := r.db.Prepare(createUserQuery)
	if err != nil {
		return models.User{}, err
	}

	tx, err := r.db.Begin()
	if err != nil {
		return models.User{}, err
	}

	_, err = tx.Stmt(query).Exec(id, user.Email, user.Password,
		user.FName, user.LName, user.DateOfBirth, user.IsPrivate,
		user.Avatar, user.NickName, user.About)
	if err != nil {
		r.logger.Println("doing rollback")
		r.logger.Println(err.Error())
		tx.Rollback()
		return models.User{}, err
	} else {
		tx.Commit()
	}

	dbUser, err := r.GetUserByID(id)
	if err != nil {
		return dbUser, err
	}

	return dbUser, nil
}

func (r *usersRepo) GetUserByEmail(email string) (models.User, error) {
	var dbUser models.User
	row := r.db.QueryRow(getUserByEmailQuery, email)
	err := row.Scan(&dbUser.UserID, &dbUser.Email, &dbUser.Password,
		&dbUser.FName, &dbUser.LName, &dbUser.DateOfBirth,
		&dbUser.IsPrivate, &dbUser.Avatar, &dbUser.NickName, &dbUser.About)
	if err != nil {
		return models.User{}, err
	}
	return dbUser, nil
}

func (r *usersRepo) GetUserByID(id uuid.UUID) (models.User, error) {
	var dbUser models.User
	row := r.db.QueryRow(getUserByIDQuery, id)
	err := row.Scan(&dbUser.UserID, &dbUser.Email, &dbUser.Password,
		&dbUser.FName, &dbUser.LName, &dbUser.DateOfBirth,
		&dbUser.IsPrivate, &dbUser.Avatar, &dbUser.NickName, &dbUser.About)
	if err != nil {
		return models.User{}, err
	}
	return dbUser, nil
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
		return err
	} else {
		tx.Commit()
	}

	return nil
}

func (r *usersRepo) GetUserAuth(session string) (models.UserAuth, error) {
	var dbUser models.UserAuth
	row := r.db.QueryRow(getUserAuthQuery, session)
	err := row.Scan(&dbUser.ID, &dbUser.UserID, &dbUser.Expires, &dbUser.Session)

	switch err {
	case sql.ErrNoRows:
		return models.UserAuth{}, NotFound
	case nil:
		return dbUser, nil
	default:
		return models.UserAuth{}, err
	}
}

func (r *usersRepo) Follow(userFollow *models.UserFollow, dbUser models.User) error {
	query, err := r.db.Prepare(createFollowQuery)
	if err != nil {
		return err
	}

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Stmt(query).Exec(userFollow.UserID1, userFollow.UserID2, userFollow.CreatedAt, dbUser.IsPrivate)
	if err != nil {
		r.logger.Println("doing rollback")
		r.logger.Println(err.Error())
		tx.Rollback()
		return err
	} else {
		tx.Commit()
	}

	return nil
}

func (r *usersRepo) GetFollow(user models.User) ([]dto.Follow, error) {
	var follows []dto.Follow
	rows, err := r.db.Query(getInvites, user.UserID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var dbUser dto.Follow
		err = rows.Scan(&dbUser.UserID, &dbUser.Email, &dbUser.FName,
			&dbUser.LName, &dbUser.DateOfBirth,
			&dbUser.IsPrivate, &dbUser.Avatar)
		if err != nil {
			return nil, err
		}
		follows = append(follows, dbUser)
	}

	err = rows.Err()
	if err != nil {
		r.logger.Println(err)
	}

	return follows, nil
}

func (r *usersRepo) Resolve(resolve *dto.ModelResolve, user models.User) error {
	var preparedQuery string
	if resolve.Accept {
		preparedQuery = acceptInvite
	} else {
		preparedQuery = declineInvite
	}

	query, err := r.db.Prepare(preparedQuery)
	if err != nil {
		return err
	}

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Stmt(query).Exec(resolve.UserID, user.UserID)
	if err != nil {
		r.logger.Println("doing rollback")
		r.logger.Println(err.Error())
		tx.Rollback()
		return err
	} else {
		tx.Commit()
	}

	return nil
}

func (r *usersRepo) GetFriends(user models.User) ([]dto.Follow, error) {
	var friends []dto.Follow
	rows, err := r.db.Query(getFriends, user.UserID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var dbUser dto.Follow
		err = rows.Scan(&dbUser.UserID, &dbUser.Email, &dbUser.FName,
			&dbUser.LName, &dbUser.DateOfBirth,
			&dbUser.IsPrivate, &dbUser.Avatar)
		if err != nil {
			return nil, err
		}
		friends = append(friends, dbUser)
	}

	err = rows.Err()
	if err != nil {
		r.logger.Println(err)
	}

	return friends, nil
}
