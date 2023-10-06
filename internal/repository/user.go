package repository

import (
	"github.com/muhrizqiardi/spendtracker/internal/database/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Insert(email string, password string) (model.UserModel, error)
	GetOneByEmail(email string) (model.UserModel, error)
	GetOneByID(id int) (model.UserModel, error)
	UpdateOneByID(id int, email string, password string) (model.UserModel, error)
	DeleteOneByID(id int) (model.UserModel, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (ur *userRepository) Insert(email string, password string) (model.UserModel, error) {
	newUser := model.UserModel{
		Email:    email,
		Password: password,
	}
	if err := ur.db.Save(&newUser).Error; err != nil {
		return model.UserModel{}, err
	}

	return newUser, nil
}

func (ur *userRepository) GetOneByEmail(email string) (model.UserModel, error) {
	var user model.UserModel
	if err := ur.db.First(&user, "email = ?", email).Error; err != nil {
		return model.UserModel{}, err
	}

	return user, nil
}

func (ur *userRepository) GetOneByID(id int) (model.UserModel, error) {
	var user model.UserModel
	if err := ur.db.First(&user, "id = ?", id).Error; err != nil {
		return model.UserModel{}, err
	}

	return user, nil
}

func (ur *userRepository) UpdateOneByID(id int, email string, password string) (model.UserModel, error) {
	user := model.UserModel{
		Model: gorm.Model{
			ID: uint(id),
		},
		Email:    email,
		Password: password,
	}
	if err := ur.db.Where("id = ?", id).First(&model.UserModel{}).Error; err != nil {
		return model.UserModel{}, err
	}
	if err := ur.db.Save(&user).Error; err != nil {
		return model.UserModel{}, err
	}

	return user, nil
}

func (ur *userRepository) DeleteOneByID(id int) error {
	user := model.UserModel{
		Model: gorm.Model{
			ID: uint(id),
		},
	}
	if err := ur.db.Where("id = ?", id).First(&model.UserModel{}).Error; err != nil {
		return err
	}
	if err := ur.db.Unscoped().Delete(&user).Error; err != nil {
		return err
	}

	return nil
}
