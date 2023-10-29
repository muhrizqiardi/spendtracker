package repository

import (
	"github.com/muhrizqiardi/spendtracker/internal/database/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Insert(email string, password string) (model.User, error)
	GetOneByEmail(email string) (model.User, error)
	GetOneByID(id int) (model.User, error)
	UpdateOneByID(id int, email string, password string) (model.User, error)
	DeleteOneByID(id int) (model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (ur *userRepository) Insert(email string, password string) (model.User, error) {
	newUser := model.User{
		Email:    email,
		Password: password,
	}
	if err := ur.db.Save(&newUser).Error; err != nil {
		return model.User{}, err
	}

	return newUser, nil
}

func (ur *userRepository) GetOneByEmail(email string) (model.User, error) {
	var user model.User
	if err := ur.db.First(&user, "email = ?", email).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (ur *userRepository) GetOneByID(id int) (model.User, error) {
	var user model.User
	if err := ur.db.First(&user, "id = ?", id).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (ur *userRepository) UpdateOneByID(id int, email string, password string) (model.User, error) {
	user := model.User{
		Model: gorm.Model{
			ID: uint(id),
		},
		Email:    email,
		Password: password,
	}
	if err := ur.db.Where("id = ?", id).First(&model.User{}).Error; err != nil {
		return model.User{}, err
	}
	if err := ur.db.Save(&user).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (ur *userRepository) DeleteOneByID(id int) error {
	user := model.User{
		Model: gorm.Model{
			ID: uint(id),
		},
	}
	if err := ur.db.Where("id = ?", id).First(&model.User{}).Error; err != nil {
		return err
	}
	if err := ur.db.Unscoped().Delete(&user).Error; err != nil {
		return err
	}

	return nil
}
