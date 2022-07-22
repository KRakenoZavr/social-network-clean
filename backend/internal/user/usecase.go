package user

import (
	"net/http"

	"mux/internal/models"
	"mux/pkg/utils/errHandler"
)

type UseCase interface {
	Create(*models.User) (models.User, *http.Cookie, *errHandler.ServiceError)
	Login(*models.UserLogin) (*http.Cookie, *errHandler.ServiceError)
	Follow(*models.UserFollow, models.User) *errHandler.ServiceError
}
