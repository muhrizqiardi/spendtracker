package service

import (
	"errors"
	"fmt"

	"github.com/muhrizqiardi/spendtracker/internal/database/model"
	"github.com/muhrizqiardi/spendtracker/internal/dto"
	"github.com/muhrizqiardi/spendtracker/internal/repository"
)

var ErrAccountNotBelongedToUser = errors.New("Account doesn't belong to current user")

type ExpenseService interface {
	Create(userID int, accountID int, payload dto.CreateExpenseDTO) (model.Expense, error)
	GetOneByID(id int) (model.Expense, error)
	GetMany(itemPerPage, page int) ([]model.Expense, error)
	GetManyBelongedToUser(userID, itemPerPage, page int) ([]model.Expense, error)
	GetManyBelongedToAccount(userID, accountID, itemPerPage, page int) ([]model.Expense, error)
	GetManyBelongedToCategory(userID, categoryID, itemPerPage, page int) ([]model.Expense, error)
	GetManyBelongedToCategoryAccount(userID, categoryID, accountID, itemPerPage, page int) ([]model.Expense, error)
	UpdateOneByID(id int, payload dto.UpdateExpenseDTO) (model.Expense, error)
	DeleteOneByID(id int) error
}

type expenseService struct {
	er repository.ExpenseRepository
	as AccountService
}

func NewExpenseService(er repository.ExpenseRepository, as AccountService) *expenseService {
	return &expenseService{er, as}
}

func (es *expenseService) Create(userID int, accountID int, payload dto.CreateExpenseDTO) (model.Expense, error) {
	fmt.Println("mamamia")
	if account, err := es.as.GetOneByID(accountID); err != nil || account.UserID != uint(userID) {
		return model.Expense{}, ErrAccountNotBelongedToUser
	}

	expense, err := es.er.Insert(uint(userID), uint(accountID), payload.Name, payload.Description, payload.Amount)
	if err != nil {
		return model.Expense{}, err
	}

	return expense, nil
}

func (es *expenseService) GetOneByID(id int) (model.Expense, error) {
	expense, err := es.er.GetOneByID(uint(id))
	if err != nil {
		return model.Expense{}, err
	}

	return expense, nil
}

func (es *expenseService) GetMany(itemPerPage, page int) ([]model.Expense, error) {
	expenses, err := es.er.GetMany(itemPerPage, (page-1)*itemPerPage)
	if err != nil {
		return nil, err
	}

	return expenses, nil
}

func (es *expenseService) GetManyBelongedToUser(userID, itemPerPage, page int) ([]model.Expense, error) {
	expenses, err := es.er.GetManyBelongedToUser(uint(userID), itemPerPage, (page-1)*itemPerPage)
	if err != nil {
		return nil, err
	}

	return expenses, nil
}

func (es *expenseService) GetManyBelongedToAccount(userID, accountID, itemPerPage, page int) ([]model.Expense, error) {
	expenses, err := es.er.GetManyBelongedToAccount(uint(userID), uint(accountID), itemPerPage, (page-1)*itemPerPage)
	if err != nil {
		return nil, err
	}

	return expenses, nil
}

func (es *expenseService) GetManyBelongedToCategory(userID, categoryID, itemPerPage, page int) ([]model.Expense, error) {
	expenses, err := es.er.GetManyBelongedToCategory(uint(userID), uint(categoryID), itemPerPage, (page-1)*itemPerPage)
	if err != nil {
		return nil, err
	}

	return expenses, nil
}

func (es *expenseService) GetManyBelongedToCategoryAccount(userID, categoryID, accountID, itemPerPage, page int) ([]model.Expense, error) {
	expenses, err := es.er.GetManyBelongedToCategoryAccount(uint(userID), uint(categoryID), uint(accountID), itemPerPage, (page-1)*itemPerPage)
	if err != nil {
		return nil, err
	}

	return expenses, nil
}

func (es *expenseService) UpdateOneByID(id int, payload dto.UpdateExpenseDTO) (model.Expense, error) {
	expense, err := es.er.UpdateOneByID(uint(id), payload.Name, payload.Description, payload.Amount)
	if err != nil {
		return model.Expense{}, err
	}

	return expense, nil
}

func (es *expenseService) DeleteOneByID(id int) error {
	if err := es.er.DeleteOneByID(uint(id)); err != nil {
		return err
	}

	return nil
}
