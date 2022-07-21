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
}
