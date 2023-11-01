package dto

type CreateExpenseDTO struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	Amount      int    `json:"amount" validate:"required"`
}

type UpdateExpenseDTO struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	Amount      int    `json:"amount" validate:"required"`
}
