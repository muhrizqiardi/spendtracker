package repository

import (
	"github.com/muhrizqiardi/spendtracker/internal/database/model"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	Insert(userID uint, name string) (model.Category, error)
	GetOneByID(id uint) (model.Category, error)
	GetOneByName(name string) (model.Category, error)
	GetMany(userID uint, limit int, offset int) ([]model.Category, error)
	Delete(id uint)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{db}
}

func (cr *categoryRepository) Insert(userID uint, name string) (model.Category, error) {
	category := model.Category{
		UserID: userID,
		Name:   name,
	}
	if err := cr.db.Save(&category).Error; err != nil {
		return model.Category{}, err
	}

	return category, nil
}

func (cr *categoryRepository) GetOneByID(id uint) (model.Category, error) {
	var category model.Category
	if err := cr.db.First(&category, "id = ?", id).Error; err != nil {
		return model.Category{}, err
	}

	return category, nil
}

func (cr *categoryRepository) GetOneByName(name string) (model.Category, error) {
	var category model.Category
	if err := cr.db.First(&category, "name = ?", name).Error; err != nil {
		return model.Category{}, err
	}

	return category, nil
}

func (cr *categoryRepository) GetMany(userID uint, limit int, offset int) ([]model.Category, error) {
	var categories []model.Category
	if err := cr.db.Limit(limit).Offset(offset).Find(&categories, "user_id = ?", userID).Error; err != nil {
		return nil, err
	}

	return categories, nil
}

func (cr *categoryRepository) Delete(id uint) error {
	var category model.Category
	if err := cr.db.Where("id = ?", id).Delete(&category).Error; err != nil {
		return err
	}

	return nil
}
