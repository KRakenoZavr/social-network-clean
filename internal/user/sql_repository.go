package user

import (
	"mux/internal/models"

	"github.com/satori/uuid"
)

type Repository interface {
	Create(*models.User) (uuid.UUID, error)
	GetUserByEmail(string) (models.User, error)
	CheckUserByEmail(string) (bool, error)

	CreateUserAuth(*models.UserAuth) error
	GetUserAuth(string) (models.UserAuth, error)
}
