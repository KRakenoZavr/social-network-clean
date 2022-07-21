package group

import (
	"mux/internal/models"

	"github.com/satori/uuid"
)

type Repository interface {
	Create(*models.Group) error
	JoinRequest(*models.GroupUser) error

	GetAllGroups() ([]models.Group, error)
	GetRequests(models.User) ([]models.User, error)

	CheckGroupByTitle(string) (bool, error)
	CheckGroupByID(uuid.UUID) (bool, error)
}
