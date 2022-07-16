package usecase

import (
	"log"

	"mux/internal/models"
	"mux/internal/user"
)

type userUC struct {
	userRepo user.Repository
	logger   *log.Logger
}

func NewUserUseCase(userRepo user.Repository, logger *log.Logger) user.UseCase {
	return &userUC{userRepo: userRepo, logger: logger}
}

func (u *userUC) Create(user *models.User) error {
	return u.userRepo.Create(user)
}
