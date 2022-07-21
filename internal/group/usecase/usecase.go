package usecase

import (
	"errors"
	"log"
	"net/http"
	"time"

	"mux/internal/group"
	"mux/internal/models"
	"mux/pkg/utils"
	"mux/pkg/utils/errHandler"
)

type groupUC struct {
	groupRepo group.Repository
	logger    *log.Logger
}

func NewUseCase(groupRepo group.Repository, logger *log.Logger) group.UseCase {
	return &groupUC{groupRepo: groupRepo, logger: logger}
}

func (u *groupUC) validateGroup(group *models.Group) []error {
	validator := utils.NewValidator()
	validator.CheckNull(group.Title, "title")
	validator.CheckNull(group.Body, "description")

	return validator.Errors()
}

func (u *groupUC) Create(group *models.Group, user models.User) *errHandler.ServiceError {
	// validate group fields
	listOfErrors := u.validateGroup(group)
	if listOfErrors != nil {
		return &errHandler.ServiceError{
			Code:    http.StatusBadRequest,
			Message: utils.ErrArrayToStringArray(listOfErrors),
			Err:     errors.New("group: fields validation error"),
		}
	}

	// check if exist
	isExist, err := u.groupRepo.CheckGroupByTitle(group.Title)
	if err != nil {
		return &errHandler.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: []string{"group: db access error"},
			Err:     err,
		}
	}

	if isExist {
		return &errHandler.ServiceError{
			Code:    http.StatusBadRequest,
			Message: []string{"group: group with title already exists"},
			Err:     errors.New("group already exists"),
		}
	}

	group.UserID = user.UserID
	group.CreatedAt = time.Now()

	// create group
	err = u.groupRepo.Create(group)
	if err != nil {
		return &errHandler.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: []string{"group: db access error"},
			Err:     err,
		}
	}

	return &errHandler.ServiceError{
		Err: nil,
	}
}

func (u *groupUC) JoinRequest(gUser *models.GroupUser, user models.User) *errHandler.ServiceError {
	// check if exist
	isExist, err := u.groupRepo.CheckGroupByID(gUser.GroupID)
	if err != nil {
		return &errHandler.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: []string{"group user: db access error"},
			Err:     err,
		}
	}

	if !isExist {
		return &errHandler.ServiceError{
			Code:    http.StatusBadRequest,
			Message: []string{"group user: no group with such id"},
			Err:     errors.New("group not exists"),
		}
	}

	gUser.UserID = user.UserID
	gUser.CreatedAt = time.Now()
	gUser.Invite = models.InviteGroup

	// create group
	err = u.groupRepo.JoinRequest(gUser)
	if err != nil {
		return &errHandler.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: []string{"group user: db access error"},
			Err:     err,
		}
	}

	return &errHandler.ServiceError{
		Err: nil,
	}
}

func (u *groupUC) GetAllGroups() ([]models.Group, *errHandler.ServiceError) {
	groups, err := u.groupRepo.GetAllGroups()
	if err != nil {
		return nil, &errHandler.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: []string{"group: db access error"},
			Err:     err,
		}
	}

	return groups, &errHandler.ServiceError{
		Err: err,
	}
}

func (u *groupUC) GetRequests(user models.User) ([]models.User, *errHandler.ServiceError) {
	gUsers, err := u.groupRepo.GetRequests(user)
	if err != nil {
		return nil, &errHandler.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: []string{"group: db access error"},
			Err:     err,
		}
	}

	return gUsers, &errHandler.ServiceError{
		Err: err,
	}
}
