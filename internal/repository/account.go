package repository

import (
	"github.com/muhrizqiardi/spendtracker/internal/database/model"
	"gorm.io/gorm"
)

type AccountRepository interface {
	Insert(userID uint, currencyID uint, name string, initialAmount int) (model.Account, error)
	GetOneByID(id uint) (model.Account, error)
	GetMany(userID uint, limit int, offset int) ([]model.Account, error)
	UpdateOneByID(id uint, currencyID uint, name string, initialAmount int) (model.Account, error)
	DeleteOneByID(id uint) error
}

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *accountRepository {
	return &accountRepository{db}
}

func (ar *accountRepository) Insert(userID uint, currencyID uint, name string, initialAmount int) (model.Account, error) {
	newAccount := model.Account{
		UserID:        userID,
		CurrencyID:    currencyID,
		Name:          name,
		InitialAmount: initialAmount,
	}
	if err := ar.db.Save(&newAccount).Error; err != nil {
		return model.Account{}, err
	}

	return newAccount, nil
}

func (ar *accountRepository) GetOneByID(id uint) (model.Account, error) {
	var account model.Account
	if err := ar.db.First(&account, "id = ?", id).Error; err != nil {
		return model.Account{}, err
	}

	return account, nil
}

func (ar *accountRepository) GetMany(userID uint, limit int, offset int) ([]model.Account, error) {
	var accounts []model.Account
	if err := ar.db.Limit(limit).Offset(offset).Find(&accounts, "user_id = ?", userID).Error; err != nil {
		return []model.Account{}, err
	}

	return accounts, nil
}

func (ar *accountRepository) UpdateOneByID(id uint, currencyID uint, name string, initialAmount int) (model.Account, error) {
	updatedAccount := model.Account{
		Model: gorm.Model{
			ID: uint(id),
		},
		CurrencyID:    currencyID,
		Name:          name,
		InitialAmount: initialAmount,
	}
	if err := ar.db.Save(&updatedAccount).Error; err != nil {
		return model.Account{}, err
	}

	return updatedAccount, nil
}

func (ar *accountRepository) DeleteOneByID(id uint) error {
	var account model.Account
	if err := ar.db.Where("id = ?", id).Delete(&account).Error; err != nil {
		return err
	}

	return nil
}
