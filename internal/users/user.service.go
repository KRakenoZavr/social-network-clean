package users

import (
	"mux/pkg/users/dto"
)

// type UserService struct {
// 	userEntity *UserEntity
// }

// func NewUserService() *UserService {
// 	return &UserService{
// 		userEntity: newUserEntity(),
// 	}
// }

func CreateUser(user dto.CreateUserDTO) {
	// createUser(user)
}

func GetUser(user dto.LoginUserDTO) {
	// getUser(user)
}

type IUserService interface {
	CreateUser(dto.CreateUserDTO)
	GetUser(dto.LoginUserDTO)
	// UpdateUser(dto.UpdateUserDTO)
	// DeleteUser(dto.LoginUserDTO)
}
