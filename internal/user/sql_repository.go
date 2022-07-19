package user

import (
	"mux/internal/models"

	"github.com/satori/uuid"
)

type Repository interface {
	Create(*models.User) (uuid.UUID, error)
	CreateUserAuth(*models.UserAuth) error
	GetUserAuth(string) (models.UserAuth, error)
}
