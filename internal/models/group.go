package models

import (
	"time"

	"github.com/satori/uuid"
)

type InviteType int

const (
	// InviteAccepted invitation accepted
	InviteAccepted InviteType = iota
	// InviteUser invite from group to user
	InviteUser
	// InviteGroup request to join group from user
	InviteGroup
)

type Group struct {
	GroupID   uuid.UUID `json:"groupID"`
	UserID    uuid.UUID `json:"userID"`
	CreatedAt time.Time `json:"createdAt"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
}

type GroupUser struct {
	ID        int        `json:"id"`
	GroupID   uuid.UUID  `json:"groupID"`
	UserID    uuid.UUID  `json:"userID"`
	CreatedAt time.Time  `json:"createdAt"`
	Invite    InviteType `json:"invite"`
}
