package usecase

import (
	"errors"
	"log"
	"mux/internal/models"
	"mux/internal/user"
	"mux/pkg/utils"
	"mux/pkg/utils/errHandler"
	"net/http"
)

type userUC struct {
	userRepo user.Repository
	logger   *log.Logger
}

func NewUserUseCase(userRepo user.Repository, logger *log.Logger) user.UseCase {
	return &userUC{userRepo: userRepo, logger: logger}
}

func (u *userUC) validateUser(user *models.User) []error {
	validator := utils.NewValidator()
	validator.CheckNull(user.Email, "email")
	validator.CheckNull(user.FName, "first name")
	validator.CheckNull(user.LName, "last name")
	validator.CheckNull(user.Password, "password")

	validator.CheckBDay(user.DateOfBirth)
	validator.CheckEmail(user.Email)

	validator.CheckLen(user.Password, 6)
	validator.OnlyAlphabet(user.FName)
	validator.OnlyAlphabet(user.LName)

	return validator.Errors()
}

func (u *userUC) Create(user *models.User) *errHandler.ServiceError {
	listOfErrors := u.validateUser(user)
	if listOfErrors != nil {
		return &errHandler.ServiceError{
			Code:    http.StatusBadRequest,
			Message: utils.ErrArrayToStringArray(listOfErrors),
			Err:     errors.New("user: fields validation error"),
		}
	}

	err := u.userRepo.Create(user)
	if err != nil {
		return &errHandler.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: []string{},
			Err:     errors.New("user: db access error"),
		}
	}

	return &errHandler.ServiceError{}
}
