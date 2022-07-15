package dto

type UpdateUserDTO struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
}
