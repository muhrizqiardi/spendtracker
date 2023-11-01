package dto

type CreateAccountDTO struct {
	CurrencyID    uint   `json:"currencyId" validate:"required"`
	Name          string `json:"name" validate:"required"`
	InitialAmount int    `json:"initialAmount" validate:"required"`
}

type UpdateAccountDTO struct {
	CurrencyID    uint   `json:"currencyId" validate:"required"`
	Name          string `json:"name" validate:"required"`
	InitialAmount int    `json:"initialAmount" validate:"required"`
}
