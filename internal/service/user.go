package service

import (
	"github.com/muhrizqiardi/spendtracker/internal/database/model"
	"github.com/muhrizqiardi/spendtracker/internal/dto"
	"github.com/muhrizqiardi/spendtracker/internal/repository"
)

type UserService interface {
	Register(payload dto.RegisterUserDTO) (model.User, error)
	GetOneByID(id int) (model.User, error)
	GetOneByEmail(email string) (model.User, error)
	UpdateOneByID(id int, payload dto.UpdateUserDTO) (model.User, error)
	DeleteOneByID(id int) (model.User, error)
}

type userService struct {
	ur repository.UserRepository
}

func NewUserService(ur repository.UserRepository) *userService {
	return &userService{ur}
}

func (us *userService) Register(payload dto.RegisterUserDTO) (model.User, error) {
	// TODO: add validation
	user, err := us.ur.Insert(payload.Email, payload.Password)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (us *userService) GetOneByID(id int) (model.User, error) {
	return model.User{}, nil
}

func (us *userService) GetOneByEmail(email string) (model.User, error) {
	return model.User{}, nil
}

func (us *userService) UpdateOneByID(id int, payload dto.UpdateUserDTO) (model.User, error) {
	return model.User{}, nil
}

func (us *userService) DeleteOneByID(id int) (model.User, error) {
	return model.User{}, nil
}
