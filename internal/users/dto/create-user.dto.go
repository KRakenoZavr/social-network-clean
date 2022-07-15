package dto

type CreateUserDTO struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Age      int    `json:"age"`
}
