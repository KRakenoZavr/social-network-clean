package models

import (
	"time"

	"github.com/satori/uuid"
)

type InviteType int

const (
	// invitation accepted
	InviteAccepted InviteType = iota
	// invite from group to user
	InviteUser
	// request to join group from user
	InviteGroup
)

type Group struct {
	GroupID   uuid.UUID `json:"group_id"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
}

type GroupUser struct {
	ID        int        `json:"id"`
	GroupID   uuid.UUID  `json:"group_id"`
	UserID    uuid.UUID  `json:"user_id"`
	CreatedAt time.Time  `json:"created_at"`
	Invite    InviteType `json:"invite"`
}
