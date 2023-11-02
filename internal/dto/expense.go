package dto

type CreateExpenseDTO struct {
	CategoryID  int    `json:"categoryId" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	Amount      int    `json:"amount" validate:"required"`
}

type UpdateExpenseDTO struct {
	CategoryID  int    `json:"categoryId" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	Amount      int    `json:"amount" validate:"required"`
}
