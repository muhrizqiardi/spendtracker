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

func TestAccountService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	mar := mock_repository.NewMockAccountRepository(ctrl)
	as := NewAccountService(mar)

	t.Run("should create account", func(t *testing.T) {
		mar.EXPECT().Insert(gomock.Eq(uint(1)), gomock.Eq(uint(2)), gomock.Eq("Acme Bank"), gomock.Eq(1000)).
			DoAndReturn(func(userID uint, currencyID uint, name string, initialAmount int) (model.Account, error) {
				return model.Account{
					UserID:        uint(userID),
					CurrencyID:    uint(currencyID),
					Name:          name,
					InitialAmount: initialAmount,
				}, nil
			})

		exp := model.Account{
			UserID:        1,
			CurrencyID:    2,
			Name:          "Acme Bank",
			InitialAmount: 1000,
		}
		got, err := as.Create(1, dto.CreateAccountDTO{
			CurrencyID:    2,
			Name:          "Acme Bank",
			InitialAmount: 1000,
		})
		if err != nil {
			t.Error("exp nil; got error:", err)
		}

		opts := []cmp.Option{
			cmpopts.IgnoreFields(model.Account{}, "Model"),
		}
		testutil.CompareAndAssert(t, exp, got, opts...)
	})
}

func TestAccountService_GetOneByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	mar := mock_repository.NewMockAccountRepository(ctrl)
	as := NewAccountService(mar)

	t.Run("should get one by ID", func(t *testing.T) {
		mar.EXPECT().GetOneByID(gomock.Eq(uint(1))).DoAndReturn(func(id uint) (model.Account, error) {
			return model.Account{
				Model: gorm.Model{
					ID: uint(id),
				},
			}, nil
		})

		exp := model.Account{
			Model: gorm.Model{
				ID: uint(1),
			},
		}
		got, err := as.GetOneByID(1)
		if err != nil {
			t.Error("exp nil; got error:", err)
		}

		opts := []cmp.Option{
			cmpopts.IgnoreFields(model.Account{}, "Model.CreatedAt"),
			cmpopts.IgnoreFields(model.Account{}, "Model.UpdatedAt"),
			cmpopts.IgnoreFields(model.Account{}, "Model.DeletedAt"),
			cmpopts.IgnoreFields(model.Account{}, "CurrencyID"),
			cmpopts.IgnoreFields(model.Account{}, "Name"),
			cmpopts.IgnoreFields(model.Account{}, "InitialAmount"),
		}
		testutil.CompareAndAssert(t, exp, got, opts...)
	})
}

func TestAccountService_GetMany(t *testing.T) {
	ctrl := gomock.NewController(t)
	mar := mock_repository.NewMockAccountRepository(ctrl)
	as := NewAccountService(mar)

	t.Run("should get many accoutns", func(t *testing.T) {
		mar.EXPECT().GetMany(gomock.Eq(uint(1)), gomock.Eq(10), gomock.Eq(20)).
			DoAndReturn(func(userID uint, limit int, offset int) ([]model.Account, error) {
				return []model.Account{
					{
						UserID: userID,
						Name:   "Lorem Bank",
					},
					{
						UserID: userID,
						Name:   "Acme Bank",
					},
					{
						UserID: userID,
						Name:   "Ipsum Bank",
					},
				}, nil
			})

		exp := []model.Account{
			{
				UserID: uint(1),
				Name:   "Lorem Bank",
			},
			{
				UserID: uint(1),
				Name:   "Acme Bank",
			},
			{
				UserID: uint(1),
				Name:   "Ipsum Bank",
			},
		}
		got, err := as.GetMany(1, 10, 3)
		if err != nil {
			t.Error("exp nil; got error:", err)
		}

		opts := []cmp.Option{
			cmpopts.IgnoreFields(model.Account{}, "Model"),
			cmpopts.IgnoreFields(model.Account{}, "CurrencyID"),
			cmpopts.IgnoreFields(model.Account{}, "InitialAmount"),
		}
		testutil.CompareAndAssert(t, exp, got, opts...)
	})
}

func TestAccountService_UpdateOneByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	mar := mock_repository.NewMockAccountRepository(ctrl)
	as := NewAccountService(mar)

	t.Run("should update account", func(t *testing.T) {
		mar.EXPECT().UpdateOneByID(gomock.Eq(uint(1)), gomock.Eq(uint(2)), gomock.Eq("Acme Bank"), gomock.Eq(1000)).
			DoAndReturn(func(id uint, currencyID uint, name string, initialAmount int) (model.Account, error) {
				return model.Account{
					Model: gorm.Model{
						ID: uint(id),
					},
					CurrencyID:    uint(currencyID),
					Name:          name,
					InitialAmount: initialAmount,
				}, nil
			})

		exp := model.Account{
			Model: gorm.Model{
				ID: uint(1),
			},
			CurrencyID:    2,
			Name:          "Acme Bank",
			InitialAmount: 1000,
		}
		got, err := as.UpdateOneByID(1, dto.UpdateAccountDTO{
			CurrencyID:    2,
			Name:          "Acme Bank",
			InitialAmount: 1000,
		})
		if err != nil {
			t.Error("exp nil; got error:", err)
		}

		opts := []cmp.Option{
			cmpopts.IgnoreFields(model.Account{}, "Model.CreatedAt"),
			cmpopts.IgnoreFields(model.Account{}, "Model.UpdatedAt"),
			cmpopts.IgnoreFields(model.Account{}, "Model.DeletedAt"),
		}
		testutil.CompareAndAssert(t, exp, got, opts...)
	})
}

func TestAccountService_DeleteOneByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	mar := mock_repository.NewMockAccountRepository(ctrl)
	as := NewAccountService(mar)

	t.Run("should return error when repository layer returns error", func(t *testing.T) {
		mar.EXPECT().DeleteOneByID(gomock.Eq(uint(1))).DoAndReturn(func(id uint) error {
			return errors.New("")
		})

		if err := as.DeleteOneByID(1); err == nil {
			t.Error("exp error; got nil")
		}
	})
	t.Run("should delete and return nil", func(t *testing.T) {
		mar.EXPECT().DeleteOneByID(gomock.Eq(uint(1))).DoAndReturn(func(id uint) error {
			return nil
		})

		if err := as.DeleteOneByID(1); err != nil {
			t.Error("exp nil; got error:", err)
		}
	})
}
