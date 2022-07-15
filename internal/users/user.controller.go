package users

import (
	"encoding/json"
	"log"
	"net/http"

	"mux/pkg/helper"
	"mux/pkg/users/dto"
)

type UserController struct {
	userService IUserService
}

func NewUserController() *UserController {
	s := struct {
		CreateUser func(dto.CreateUserDTO)
	}{
		CreateUser,
	}

	return &UserController{
		userService: s,
	}
}

func (u *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody := dto.CreateUserDTO{}

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		helper.ResponseWithErr(w, http.StatusBadRequest, err.Error())
		return
	}

	log.Println(requestBody)

	u.userService.CreateUser(requestBody)

	w.WriteHeader(http.StatusCreated)
}

type IUserController interface {
	CreateUser(http.ResponseWriter, *http.Request)
	// GetUser(http.ResponseWriter, *http.Request)
	// UpdateUser(http.ResponseWriter, *http.Request)
	// DeleteUser(http.ResponseWriter, *http.Request)
}
