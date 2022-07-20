package group

import (
	"mux/internal/models"
	"mux/pkg/utils/errHandler"
)

type UseCase interface {
	Create(*models.Group, models.User) *errHandler.ServiceError
}
