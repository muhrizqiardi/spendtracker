package dto

type RegisterUserDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
