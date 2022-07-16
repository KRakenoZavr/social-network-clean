package user

import "mux/internal/models"

type UseCase interface {
	Create(user *models.User) error
}
