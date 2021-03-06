package models

import (
	"time"

	"github.com/satori/uuid"
)

type User struct {
	UserID      uuid.UUID `json:"userID" db:"user_id"`
	Email       string    `json:"email" db:"email"`
	Password    string    `json:"password" db:"password"`
	FName       string    `json:"firstName" db:"first_name"`
	LName       string    `json:"lastName" db:"last_name"`
	DateOfBirth time.Time `json:"dateOfBirth" db:"date_of_birth"`
	IsPrivate   bool      `json:"isPrivate" db:"is_private"`
	Avatar      string    `json:"avatar" db:"avatar"`
	NickName    string    `json:"nickname" db:"nickname"`
	About       string    `json:"about" db:"about"`
}

type UserAuth struct {
	ID      int
	UserID  uuid.UUID
	Expires time.Time
	Session uuid.UUID
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserFollow struct {
	ID        int        `json:"id"`
	UserID1   uuid.UUID  `json:"userID1"`
	UserID2   uuid.UUID  `json:"userID2"`
	CreatedAt time.Time  `json:"createdAt"`
	Invite    InviteType `json:"invite"`
}
