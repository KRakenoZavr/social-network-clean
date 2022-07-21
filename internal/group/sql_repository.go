package group

import (
	"mux/internal/group/dto"
	"mux/internal/models"

	"github.com/satori/uuid"
)

type Repository interface {
	Create(*models.Group) error
	JoinRequest(*models.GroupUser) error
	Invite(*models.GroupUser, models.User) error
	GetInvites(models.User) ([]models.Group, error)

	GetAllGroups() ([]models.Group, error)
	GetRequests(models.User) ([]dto.ModelJoinRequest, error)

	CheckGroupByTitle(string) (bool, error)
	CheckGroupByID(uuid.UUID) (bool, error)
	CheckAdmin(uuid.UUID, uuid.UUID) (bool, error)
}
