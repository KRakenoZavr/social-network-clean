package group

import (
	"github.com/satori/uuid"
	"mux/internal/models"
)

type Repository interface {
	Create(*models.Group) error
	JoinRequest(*models.GroupUser) error

	GetAllGroups() ([]models.Group, error)

	CheckGroupByTitle(string) (bool, error)
	CheckGroupByID(uuid.UUID) (bool, error)
}
