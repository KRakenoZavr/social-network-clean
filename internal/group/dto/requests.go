package dto

import "github.com/satori/uuid"

type ModelJoinRequest struct {
	UserID  uuid.UUID
	Email   string
	FName   string
	LName   string
	Avatar  string
	GroupID uuid.UUID
	Title   string
}
