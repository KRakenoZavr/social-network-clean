package models

import "github.com/satori/uuid"

type User struct {
	UserID 		uuid.UUID 	`json:"user_id" db:"user_id" validate:"omitempty"`
	Username 	string 		`json:"username" db:"username"`
	Name 		string 		`json:"name" db:"name"`
	Password 	string 		`json:"password" db:"password"`
	Age 		int 		`json:"age" db:"age"`
}
