package dto

import (
	"time"

	"github.com/satori/uuid"
)

type Follow struct {
	UserID      uuid.UUID
	Email       string
	FName       string
	LName       string
	DateOfBirth time.Time
	IsPrivate   bool
	Avatar      string
}

type ModelResolve struct {
	UserID uuid.UUID `json:"userId"`
	Accept bool      `json:"accept"`
}
