package service

import (
	"github.com/muhrizqiardi/spendtracker/internal/database/model"
	"github.com/muhrizqiardi/spendtracker/internal/dto"
	"github.com/muhrizqiardi/spendtracker/internal/repository"
)

type AccountService interface {
	Create(userID int, payload dto.CreateAccountDTO) (model.Account, error)
	GetOneByID(id int) (model.Account, error)
	GetMany(userID, itemPerPage, page int) ([]model.Account, error)
	UpdateOneByID(id int, payload dto.UpdateAccountDTO) (model.Account, error)
	DeleteOneByID(id int) error
}

type accountService struct {
	ar repository.AccountRepository
}

func NewAccountService(ar repository.AccountRepository) *accountService {
	return &accountService{ar}
}

func (as *accountService) Create(userID int, payload dto.CreateAccountDTO) (model.Account, error) {
	account, err := as.ar.Insert(uint(userID), payload.CurrencyID, payload.Name, payload.InitialAmount)
	if err != nil {
		return model.Account{}, err
	}

	return account, nil
}

func (as *accountService) GetOneByID(id int) (model.Account, error) {
	account, err := as.ar.GetOneByID(uint(id))
	if err != nil {
		return model.Account{}, err
	}

	return account, nil
}

func (as *accountService) GetMany(userID, itemPerPage, page int) ([]model.Account, error) {
	account, err := as.ar.GetMany(uint(userID), itemPerPage, (page-1)*itemPerPage)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (as *accountService) UpdateOneByID(id int, payload dto.UpdateAccountDTO) (model.Account, error) {
	account, err := as.ar.UpdateOneByID(uint(id), payload.CurrencyID, payload.Name, payload.InitialAmount)
	if err != nil {
		return model.Account{}, err
	}

	return account, nil
}

func (as *accountService) DeleteOneByID(id int) error {
	if err := as.ar.DeleteOneByID(uint(id)); err != nil {
		return err
	}

	return nil
}
