package usecase

import (
	"database/sql"
	"errors"
	"log"
	"mux/internal/models"
	"mux/internal/user"
	"mux/pkg/utils"
	"mux/pkg/utils/errHandler"
	"net/http"
	"time"

	"github.com/satori/uuid"
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

func (u *userUC) Create(user *models.User) (*http.Cookie, *errHandler.ServiceError) {
	// validate user fields
	listOfErrors := u.validateUser(user)
	if listOfErrors != nil {
		return nil, &errHandler.ServiceError{
			Code:    http.StatusBadRequest,
			Message: utils.ErrArrayToStringArray(listOfErrors),
			Err:     errors.New("user: fields validation error"),
		}
	}

	// check if exist
	isExist, err := u.userRepo.CheckUserByEmail(user.Email)
	if err != nil {
		return nil, &errHandler.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: []string{"user: db access error"},
			Err:     err,
		}
	}

	if isExist {
		return nil, &errHandler.ServiceError{
			Code:    http.StatusBadRequest,
			Message: []string{"user: user with email already exists"},
			Err:     errors.New("user already exists"),
		}
	}

	// create user
	id, err := u.userRepo.Create(user)
	if err != nil {
		return nil, &errHandler.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: []string{"user: db access error"},
			Err:     err,
		}
	}

	return u.createCookie(id), &errHandler.ServiceError{}
}

func (u *userUC) Login(user *models.UserLogin) (*http.Cookie, *errHandler.ServiceError) {
	// check if exist
	dbUser, err := u.userRepo.GetUserByEmail(user.Email)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, &errHandler.ServiceError{
				Code:    http.StatusBadRequest,
				Message: []string{"check your credentials"},
				Err:     err,
			}
		}

		return nil, &errHandler.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: []string{"user: db access error"},
			Err:     err,
		}
	}

	if dbUser.Password != user.Password {
		return nil, &errHandler.ServiceError{
			Code:    http.StatusBadRequest,
			Message: []string{"check your credentials"},
			Err:     errors.New("wrong credentials"),
		}
	}

	return u.createCookie(dbUser.UserID), &errHandler.ServiceError{}
}

func (u *userUC) createCookie(userId uuid.UUID) *http.Cookie {
	sessionToken := uuid.NewV4()
	expires := time.Now().Add(1 * time.Hour)
	u.userRepo.CreateUserAuth(&models.UserAuth{UserID: userId, Expires: expires, Session: sessionToken})

	cookie := &http.Cookie{
		Name:     "session_token",
		Value:    sessionToken.String(),
		Expires:  expires,
		HttpOnly: true,
		Path:     "/",
	}

	return cookie
}
