package usecase

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"time"

	"mux/internal/models"
	"mux/internal/user"
	"mux/internal/user/dto"
	"mux/pkg/utils"
	"mux/pkg/utils/errHandler"

	"github.com/satori/uuid"
	"golang.org/x/crypto/bcrypt"
)

type userUC struct {
	userRepo user.Repository
	logger   *log.Logger
}

const cryptCost = 10

func NewUseCase(userRepo user.Repository, logger *log.Logger) user.UseCase {
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

func (u *userUC) Create(user *models.User) (models.User, *http.Cookie, *errHandler.ServiceError) {
	// validate user fields
	listOfErrors := u.validateUser(user)
	if listOfErrors != nil {
		return models.User{}, nil, &errHandler.ServiceError{
			Code:    http.StatusBadRequest,
			Message: utils.ErrArrayToStringArray(listOfErrors),
			Err:     errors.New("user: fields validation error"),
		}
	}

	// check if exist
	isExist, err := u.userRepo.CheckUserByEmail(user.Email)
	if err != nil {
		return models.User{}, nil, &errHandler.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: []string{"user: db access error"},
			Err:     err,
		}
	}

	if isExist {
		return models.User{}, nil, &errHandler.ServiceError{
			Code:    http.StatusBadRequest,
			Message: []string{"user: user with email already exists"},
			Err:     errors.New("user already exists"),
		}
	}

	cryptedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), cryptCost)
	if err != nil {
		return models.User{}, nil, &errHandler.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: []string{"cant crypt pass"},
			Err:     err,
		}
	}
	user.Password = string(cryptedPass)

	// create user
	dbUser, err := u.userRepo.Create(user)
	if err != nil {
		return models.User{}, nil, &errHandler.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: []string{"user: db access error"},
			Err:     err,
		}
	}

	cookie, err := u.createCookie(dbUser.UserID)
	if err != nil {
		return models.User{}, nil, &errHandler.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: []string{"cannot save token"},
			Err:     err,
		}
	}

	return dbUser, cookie, &errHandler.ServiceError{
		Err: nil,
	}
}

func (u *userUC) Login(user *models.UserLogin) (*http.Cookie, *errHandler.ServiceError) {
	// check if exist
	dbUser, err := u.userRepo.GetUserByEmail(user.Email)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, &errHandler.ServiceError{
				Code:    http.StatusBadRequest,
				Message: []string{"wrong credentials"},
				Err:     err,
			}
		}

		return nil, &errHandler.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: []string{"user: db access error"},
			Err:     err,
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if err != nil {
		return nil, &errHandler.ServiceError{
			Code:    http.StatusBadRequest,
			Message: []string{"wrong credentials"},
			Err:     err,
		}
	}

	cookie, err := u.createCookie(dbUser.UserID)
	if err != nil {
		return nil, &errHandler.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: []string{"cannot save token"},
			Err:     err,
		}
	}

	return cookie, &errHandler.ServiceError{Err: nil}
}

func (u *userUC) createCookie(userId uuid.UUID) (*http.Cookie, error) {
	sessionToken := uuid.NewV4()
	expires := time.Now().Add(1 * time.Hour)
	err := u.userRepo.CreateUserAuth(&models.UserAuth{UserID: userId, Expires: expires, Session: sessionToken})
	if err != nil {
		return nil, err
	}

	cookie := &http.Cookie{
		Name:     "session_token",
		Value:    sessionToken.String(),
		Expires:  expires,
		HttpOnly: true,
		Path:     "/",
	}

	return cookie, nil
}

func (u *userUC) Follow(userFollow *models.UserFollow, user models.User) *errHandler.ServiceError {
	// check if exist
	dbUser, err := u.userRepo.GetUserByID(userFollow.UserID2)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return &errHandler.ServiceError{
				Code:    http.StatusBadRequest,
				Message: []string{"user: no user with such id"},
				Err:     errors.New("user not exists"),
			}
		}

		return &errHandler.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: []string{"user: db access error"},
			Err:     err,
		}
	}

	userFollow.UserID1 = user.UserID
	userFollow.CreatedAt = time.Now()

	// create follow
	err = u.userRepo.Follow(userFollow, dbUser)
	if err != nil {
		return &errHandler.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: []string{"user: db access error"},
			Err:     err,
		}
	}

	return &errHandler.ServiceError{
		Err: nil,
	}
}

func (u *userUC) GetFollow(user models.User) ([]dto.Follow, *errHandler.ServiceError) {
	follows, err := u.userRepo.GetFollow(user)
	if err != nil {
		return nil, &errHandler.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: []string{"user: db access error"},
			Err:     err,
		}
	}

	return follows, &errHandler.ServiceError{
		Err: nil,
	}
}

func (u *userUC) Resolve(resolve *dto.ModelResolve, user models.User) *errHandler.ServiceError {
	err := u.userRepo.Resolve(resolve, user)
	if err != nil {
		return &errHandler.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: []string{"user: db access error"},
			Err:     err,
		}
	}

	return &errHandler.ServiceError{
		Err: nil,
	}
}
