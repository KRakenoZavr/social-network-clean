package user

import (
	"mux/internal/models"
	"mux/pkg/utils/errHandler"
	"net/http"
)

type UseCase interface {
	Create(*models.User) (*http.Cookie, *errHandler.ServiceError)
}
