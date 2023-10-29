package model

import (
	"gorm.io/gorm"
)

type Currency struct {
	gorm.Model
	Code string
}

type User struct {
	gorm.Model
	Email    string `gorm:"uniqueIndex" json:"email"`
	FullName string `json:"fullName"`
	Password string `json:"password"`
}

type Account struct {
	gorm.Model
	CurrencyID    int    `json:"currencyId"`
	Name          string `json:"name"`
	InitialAmount int    `json:"initialAmount"`
}

type Category struct {
	gorm.Model
	UserID int    `gorm:"uniqueIndex:users_categories" json:"userId"`
	Name   string `gorm:"uniqueIndex:users_categories" json:"name"`
}

type Expense struct {
	gorm.Model
	UserID      int    `json:"userId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Amount      int    `json:"amount"`
}
