package model

import (
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
}

type ExpenseModel struct {
	gorm.Model
	UserModel
	Name        string `json:"name"`
	Description string `json:"description"`
	Amount      int    `json:"amount"`
}

type ReoccuringPaymentModel struct {
	gorm.Model
	UserModel
	Name          string `json:"name"`
	Description   string `json:"description"`
	Amount        int    `json:"amount"`
	IntervalHours int    `json:"intervalHours"`
}

type Goal struct {
	gorm.Model
	Name          string `json:"email"`
	MaxSpending   int    `json:"maxSpending"`
	IntervalHours int    `json:"intervalHours"`
}
