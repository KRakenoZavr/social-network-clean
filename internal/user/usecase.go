package user

import "mux/internal/models"

type UseCase interface {
	Create(*models.User) error
}
