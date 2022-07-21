package group

import (
	"mux/internal/group/dto"
	"mux/internal/models"
	"mux/pkg/utils/errHandler"
)

type UseCase interface {
	Create(*models.Group, models.User) *errHandler.ServiceError
	JoinRequest(*models.GroupUser, models.User) *errHandler.ServiceError
	GetAllGroups() ([]models.Group, *errHandler.ServiceError)
	GetRequests(models.User) ([]dto.ModelJoinRequest, *errHandler.ServiceError)
	Invite(*models.GroupUser, models.User) *errHandler.ServiceError
	GetInvites(models.User) ([]models.Group, *errHandler.ServiceError)
	Resolve(*dto.ModelResolve, models.User) *errHandler.ServiceError
}
