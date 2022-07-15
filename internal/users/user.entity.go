package users

import (
	"database/sql"
	"fmt"
	"log"

	"mux/pkg"
	"mux/pkg/db/sqlite"
	"mux/pkg/users/dto"
)

type UserEntity struct {
	db *sqlite.SQLiteRepository
}

type UserModel struct {
	UserId   int
	Username string
	Name     string
	Password string
	Age      int
}

func newUserEntity() *UserEntity {
	return &UserEntity{
		db: pkg.Sqlite,
	}
}

func (u *UserEntity) createUser(user dto.CreateUserDTO) {
	stmt, err := u.db.DB.Prepare("INSERT INTO user(username, name, password, age) values(?,?,?,?)")
	if err != nil {
		log.Println(err.Error())
	}

	tx, err := u.db.DB.Begin()
	if err != nil {
		log.Println(err.Error())
	}

	_, err = tx.Stmt(stmt).Exec(user.Username, user.Name, user.Password, user.Age)
	if err != nil {
		log.Println(err.Error())
		tx.Rollback()
	} else {
		tx.Commit()
	}
}

func (u *UserEntity) getUser(login dto.LoginUserDTO) (UserModel, error) {
	sqlStatement := `SELECT * FROM users WHERE username=$1;`
	var user UserModel
	row := u.db.DB.QueryRow(sqlStatement, login.Username)
	err := row.Scan(&user.UserId, &user.Username, &user.Name,
		&user.Password, &user.Age)
	switch err {
	case sql.ErrNoRows:
		return UserModel{}, fmt.Errorf("user with username: %s not found", login.Username)
	case nil:
		return user, nil
	default:
		return UserModel{}, err
	}
}

type IUserEntity interface {
	createUser(dto.CreateUserDTO)
	getUser(dto.LoginUserDTO)
}
