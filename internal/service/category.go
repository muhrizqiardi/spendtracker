package service

import (
	"github.com/muhrizqiardi/spendtracker/internal/database/model"
	"github.com/muhrizqiardi/spendtracker/internal/dto"
	"github.com/muhrizqiardi/spendtracker/internal/repository"
)

type CategoryService interface {
	Create(userID int, payload dto.CreateCategoryDTO) (model.Category, error)
	GetOneByID(id int) (model.Category, error)
	GetMany(userID, itemPerPage, page int) ([]model.Category, error)
	DeleteOneByID(id int) error
}

type categoryService struct {
	cr repository.CategoryRepository
}

func NewCategoryService(cr repository.CategoryRepository) *categoryService {
	return &categoryService{cr}
}

func (cs *categoryService) Create(userID int, payload dto.CreateCategoryDTO) (model.Category, error) {
	category, err := cs.cr.Insert(uint(userID), payload.Name)
	if err != nil {
		return model.Category{}, err
	}

	return category, nil
}

func (cs *categoryService) GetOneByID(id int) (model.Category, error) {
	category, err := cs.cr.GetOneByID(uint(id))
	if err != nil {
		return model.Category{}, err
	}

	return category, nil
}

func (cs *categoryService) GetMany(userID, itemPerPage, page int) ([]model.Category, error) {
	category, err := cs.cr.GetMany(uint(userID), itemPerPage, (page-1)*itemPerPage)
	if err != nil {
		return nil, nil
	}

	return category, nil
}

func (cs *categoryService) DeleteOneByID(id int) error {
	if err := cs.cr.Delete(uint(id)); err != nil {
		return err
	}

	return nil
}
