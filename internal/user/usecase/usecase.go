package usecase

import (
	"mux/internal/models"
	"mux/internal/user"
)

type userUC struct {
	userRepo user.Repository
}

func NewUserUseCase(u user.Repository) user.UseCase {
	return &userUC{userRepo: u}
}

func (u *userUC) Create(user *models.User) error {
	return u.userRepo.Create(user)
}
