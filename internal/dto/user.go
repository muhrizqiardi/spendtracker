package dto

type RegisterUserDTO struct {
	Email    string `json:"email" validate:"email"`
	FullName string `json:"fullName" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
}

type UpdateUserDTO struct {
	Email    string `json:"email" validate:"required,email"`
	FullName string `json:"fullName" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
}
