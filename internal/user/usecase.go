package user

import (
	"mux/internal/models"
	"mux/pkg/utils/errHandler"
)

type UseCase interface {
	Create(*models.User) *errHandler.ServiceError
}
