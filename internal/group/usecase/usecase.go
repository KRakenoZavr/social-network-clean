package usecase

import (
	"errors"
	"log"
	"mux/internal/group"
	"mux/internal/models"
	"mux/pkg/utils"
	"mux/pkg/utils/errHandler"
	"net/http"
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

func (u *groupUC) Create(group *models.Group) *errHandler.ServiceError {
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

	// create group
	err = u.groupRepo.Create(group)
	if err != nil {
		return &errHandler.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: []string{"group: db access error"},
			Err:     err,
		}
	}

	return &errHandler.ServiceError{}
}
