package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/muhrizqiardi/spendtracker/internal/database/model"
	"github.com/muhrizqiardi/spendtracker/internal/dto"
	"github.com/muhrizqiardi/spendtracker/internal/repository"
	"golang.org/x/crypto/bcrypt"
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
	validate := validator.New()
	if err := validate.Struct(payload); err != nil {
		return model.User{}, err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return model.User{}, err
	}
	user, err := us.ur.Insert(payload.Email, payload.FullName, string(hashedPassword))
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (us *userService) GetOneByID(id int) (model.User, error) {
	user, err := us.ur.GetOneByID(id)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (us *userService) GetOneByEmail(email string) (model.User, error) {
	if err := validator.New().Var(email, "required,email"); err != nil {
		return model.User{}, err
	}
	user, err := us.ur.GetOneByEmail(email)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (us *userService) UpdateOneByID(id int, payload dto.UpdateUserDTO) (model.User, error) {
	validate := validator.New()
	if err := validate.Struct(payload); err != nil {
		return model.User{}, err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return model.User{}, err
	}
	user, err := us.ur.UpdateOneByID(id, payload.Email, payload.FullName, string(hashedPassword))
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (us *userService) DeleteOneByID(id int) (model.User, error) {
	return model.User{}, nil
}
