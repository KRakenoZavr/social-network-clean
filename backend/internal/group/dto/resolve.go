package dto

import "github.com/satori/uuid"

type ModelResolve struct {
	GroupID uuid.UUID `json:"groupId"`
	Accept  bool      `json:"accept"`
}
