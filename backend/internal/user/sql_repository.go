package user

import (
	"mux/internal/models"
	"mux/internal/user/dto"

	"github.com/satori/uuid"
)

type Repository interface {
	Create(*models.User) (models.User, error)
	GetUserByEmail(string) (models.User, error)
	CheckUserByEmail(string) (bool, error)
	GetUserByID(uuid.UUID) (models.User, error)

	Follow(*models.UserFollow, models.User) error
	GetFollow(models.User) ([]dto.Follow, error)
	Resolve(*dto.ModelResolve, models.User) error
	GetFriends(models.User) ([]dto.Follow, error)

	CreateUserAuth(*models.UserAuth) error
	GetUserAuth(string) (models.UserAuth, error)
}
