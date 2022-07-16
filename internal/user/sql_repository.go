package user

import "mux/internal/models"

type Repository interface {
	Create(*models.User) error
}
