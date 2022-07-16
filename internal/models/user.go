package models

import "github.com/satori/uuid"

type User struct {
	UserID uuid.UUID `json:"user_id" db:"user_id" validate:"omitempty"`
}
