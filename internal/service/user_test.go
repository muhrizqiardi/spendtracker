package service

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/muhrizqiardi/spendtracker/internal/database/model"
	"github.com/muhrizqiardi/spendtracker/internal/dto"
	mock_repository "github.com/muhrizqiardi/spendtracker/internal/repository/mock"
	"github.com/muhrizqiardi/spendtracker/tests/testutil"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func TestUserService_Register(t *testing.T) {
	ctrl := gomock.NewController(t)
	mur := mock_repository.NewMockUserRepository(ctrl)
	us := NewUserService(mur)
	opts := []cmp.Option{
		cmpopts.IgnoreFields(model.User{}, "Model"),
	}

	t.Run("should return error if payload is invalid", func(t *testing.T) {
		if _, err := us.Register(dto.RegisterUserDTO{
			Email:    "invalid.email.example.com",
			Password: "inscure",
		}); err == nil {
			t.Error("exp error; got nil")
		}
		if _, err := us.Register(dto.RegisterUserDTO{
			Email:    "validemail@example.com",
			Password: "securepass",
		}); err == nil {
			t.Error("exp error; got nil")
		}
	})
	t.Run("should return error when repository returns error", func(t *testing.T) {
		mur.
			EXPECT().
			Insert(gomock.Eq("test@example.com"), gomock.Eq("topsecret")).
			DoAndReturn(func(email string, password string) (model.User, error) {
				return model.User{}, errors.New("")
			})

		if _, err := us.Register(dto.RegisterUserDTO{
			Email:    "test@example.com",
			Password: "topsecret",
		}); err == nil {
			t.Error("exp error; got nil")
		}
	})
	t.Run("should register user with a hashed password and return user data", func(t *testing.T) {
		mur.
			EXPECT().
			Insert(gomock.Eq("test@example.com"), gomock.Eq("topsecret")).
			DoAndReturn(func(email string, password string) (model.User, error) {
				return model.User{
					Email:    email,
					Password: password,
				}, nil
			})

		exp := model.User{
			Email:    "",
			Password: "",
		}
		got, err := us.Register(dto.RegisterUserDTO{
			Email:    "test@example.com",
			Password: "topsecret",
		})
		if err != nil {
			t.Error("exp nil; got error:", err)
		}

		testutil.CompareAndAssert(t, exp, got, opts...)
	})
}

func TestUserService_GetOneByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	mur := mock_repository.NewMockUserRepository(ctrl)
	us := NewUserService(mur)
	opts := []cmp.Option{
		cmpopts.IgnoreFields(
			model.User{},
			"Model.CreatedAt",
			"Model.UpdatedAt",
			"Model.DeletedAt",
		),
	}

	t.Run("should return error when repository returns error", func(t *testing.T) {
		mur.
			EXPECT().
			GetOneByID(gomock.Eq(1)).
			DoAndReturn(func(id int) (model.User, error) {
				return model.User{}, errors.New("")
			})

		if _, err := us.GetOneByID(1); err == nil {
			t.Error("exp error; got nil")
		}
	})
	t.Run("should return user", func(t *testing.T) {
		mur.
			EXPECT().
			GetOneByID(gomock.Eq(1)).
			DoAndReturn(func(id int) (model.User, error) {
				return model.User{
					Model: gorm.Model{
						ID: uint(id),
					},
					Email:    "test@example.com",
					Password: "topsecret",
				}, nil
			})

		exp := model.User{
			Model: gorm.Model{
				ID: uint(1),
			},
			Email:    "test@example.com",
			Password: "topsecret",
		}
		got, err := us.GetOneByID(1)
		if err != nil {
			t.Error("exp nil; got error:", err)
		}

		testutil.CompareAndAssert(t, exp, got, opts...)
	})
}

func TestUserService_GetOneByEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	mur := mock_repository.NewMockUserRepository(ctrl)
	us := NewUserService(mur)
	opts := []cmp.Option{
		cmpopts.IgnoreFields(model.User{}, "Model"),
	}

	t.Run("should return error if payload is invalid", func(t *testing.T) {
		if _, err := us.GetOneByEmail("invalid.email.example.com"); err == nil {
			t.Error("exp error; got nil")
		}
	})
	t.Run("should return error when repository returns error", func(t *testing.T) {
		mur.
			EXPECT().
			GetOneByEmail(gomock.Eq("test@example.com")).
			DoAndReturn(func(email string) (model.User, error) {
				return model.User{}, errors.New("")
			})

		if _, err := us.GetOneByEmail("test@example.com"); err != nil {
			t.Error("exp error; got nil")
		}
	})
	t.Run("should return user", func(t *testing.T) {
		mur.
			EXPECT().
			GetOneByEmail(gomock.Eq("test@example.com")).
			DoAndReturn(func(email string) (model.User, error) {
				return model.User{
					Model:    gorm.Model{},
					Email:    email,
					Password: "topsecret",
				}, nil
			})

		exp := model.User{
			Model:    gorm.Model{},
			Email:    "test@example.com",
			Password: "topsecret",
		}
		got, err := us.GetOneByEmail("test@example.com")
		if err != nil {
			t.Error("exp nil; got error:", err)
		}

		testutil.CompareAndAssert(t, exp, got, opts...)
	})
}

func TestUserService_UpdateOneByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	mur := mock_repository.NewMockUserRepository(ctrl)
	us := NewUserService(mur)

	t.Run("should return error if payload is invalid", func(t *testing.T) {
		if _, err := us.UpdateOneByID(1, dto.UpdateUserDTO{
			Email:    "invalid.email.example.com",
			Password: "inscure",
		}); err == nil {
			t.Error("exp error; got nil")
		}
		if _, err := us.UpdateOneByID(1, dto.UpdateUserDTO{
			Email:    "validemail@example.com",
			Password: "securepass",
		}); err == nil {
			t.Error("exp error; got nil")
		}
	})
	t.Run("should return error when repository returns error", func(t *testing.T) {
		mur.
			EXPECT().
			UpdateOneByID(gomock.Eq(1), gomock.Any()).
			DoAndReturn(func(id int, payload dto.UpdateUserDTO) (model.User, error) {
				return model.User{}, errors.New("")
			})

		if _, err := us.UpdateOneByID(1, dto.UpdateUserDTO{
			Email:    "test@example.com",
			Password: "topsecret",
		}); err == nil {
			t.Error("exp error; got nil")
		}
	})
	t.Run("should update user and return the updated user", func(t *testing.T) {
		mur.
			EXPECT().
			UpdateOneByID(gomock.Eq(1), gomock.Eq("test@example.com"), gomock.Eq("topsecret")).
			DoAndReturn(func(email string) (model.User, error) {
				return model.User{
					Model:    gorm.Model{},
					Email:    email,
					Password: "",
				}, nil
			})

		if _, err := us.UpdateOneByID(1, dto.UpdateUserDTO{
			Email:    "test@example.com",
			Password: "topsecret",
		}); err == nil {
			t.Error("exp error; got nil")
		}
	})
}

func TestUserService_DeleteOneByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	mur := mock_repository.NewMockUserRepository(ctrl)
	us := NewUserService(mur)

	t.Run("should return error if payload is invalid", func(t *testing.T) {})
	t.Run("should return error when repository returns error", func(t *testing.T) {})
	t.Run("should delete user and return nil", func(t *testing.T) {})
}
