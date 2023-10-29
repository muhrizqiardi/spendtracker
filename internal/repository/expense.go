package repository

import (
	"github.com/muhrizqiardi/spendtracker/internal/database/model"
	"gorm.io/gorm"
)

type ExpenseRepository interface {
	Insert(userID uint, name string, description string, amount int) (model.Expense, error)
	GetOneByID(id uint) (model.Expense, error)
	GetMany(limit, offset int) ([]model.Expense, error)
	GetManyBelongedToUser(userID uint, limit, offset int) ([]model.Expense, error)
	GetManyBelongedToAccount(userID, accountID uint, limit, offset int) ([]model.Expense, error)
	GetManyBelongedToCategory(userID, categoryID uint, limit, offset int) ([]model.Expense, error)
	GetManyBelongedToCategoryAccount(userID, categoryID, accountID uint, limit, offset int) ([]model.Expense, error)
	UpdateOneByID(id uint, name string, description string, amount int) (model.Expense, error)
	DeleteOneByID(id uint) error
}

type expenseRepository struct {
	db *gorm.DB
}

func NewExpenseRepository(db *gorm.DB) *expenseRepository {
	return &expenseRepository{db}
}

func (er *expenseRepository) Insert(userID uint, name string, description string, amount int) (model.Expense, error) {
	expense := model.Expense{
		UserID:      userID,
		Name:        name,
		Description: description,
		Amount:      amount,
	}
	if err := er.db.Save(&expense).Error; err != nil {
		return model.Expense{}, err
	}

	return expense, nil
}

func (er *expenseRepository) GetOneByID(id uint) (model.Expense, error) {
	var expense model.Expense
	if err := er.db.First(&expense, "id = ?", id).Error; err != nil {
		return model.Expense{}, err
	}

	return expense, nil
}

func (er *expenseRepository) GetMany(limit, offset int) ([]model.Expense, error) {
	var expenses []model.Expense
	if err := er.db.Limit(limit).Offset(offset).Find(&expenses).Error; err != nil {
		return []model.Expense{}, err
	}

	return expenses, nil
}

func (er *expenseRepository) GetManyBelongedToUser(userID uint, limit, offset int) ([]model.Expense, error) {
	var expenses []model.Expense
	if err := er.db.Limit(limit).Offset(offset).Find(&expenses, "user_id = ?", userID).Error; err != nil {
		return []model.Expense{}, err
	}

	return expenses, nil
}

func (er *expenseRepository) GetManyBelongedToAccount(userID, accountID uint, limit, offset int) ([]model.Expense, error) {
	var expenses []model.Expense
	if err := er.db.
		Limit(limit).
		Offset(offset).
		Find(&expenses, "user_id = ? and account_id = ?", userID, accountID).
		Error; err != nil {
		return []model.Expense{}, err
	}

	return expenses, nil
}

func (er *expenseRepository) GetManyBelongedToCategory(userID, categoryID uint, limit, offset int) ([]model.Expense, error) {
	var expenses []model.Expense
	if err := er.db.
		Limit(limit).
		Offset(offset).
		Find(&expenses, "user_id = ? and category_id = ?", userID, categoryID).
		Error; err != nil {
		return []model.Expense{}, err
	}

	return expenses, nil
}

func (er *expenseRepository) GetManyBelongedToCategoryAccount(userID, categoryID, accountID uint, limit, offset int) ([]model.Expense, error) {
	var expenses []model.Expense
	if err := er.db.
		Limit(limit).
		Offset(offset).
		Find(&expenses, "user_id = ? and category_id = ? and account_id = ?", userID, categoryID, accountID).
		Error; err != nil {
		return []model.Expense{}, err
	}

	return expenses, nil
}

func (er *expenseRepository) UpdateOneByID(id uint, name string, description string, amount int) (model.Expense, error) {
	expense := model.Expense{
		Model: gorm.Model{
			ID: id,
		},
		Name:        name,
		Description: description,
		Amount:      amount,
	}
	if err := er.db.Save(&expense).Error; err != nil {
		return model.Expense{}, err
	}

	return expense, nil
}

func (er *expenseRepository) DeleteOneByID(id uint) error {
	var expense model.Expense
	if err := er.db.Where("id = ?", id).Delete(&expense).Error; err != nil {
		return err
	}

	return nil
}
