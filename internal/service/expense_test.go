package service

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/muhrizqiardi/spendtracker/internal/database/model"
	"github.com/muhrizqiardi/spendtracker/internal/dto"
	mock_repository "github.com/muhrizqiardi/spendtracker/internal/repository/mock"
	mock_service "github.com/muhrizqiardi/spendtracker/internal/service/mock"
	"github.com/muhrizqiardi/spendtracker/tests/testutil"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func TestExpenseService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	mer := mock_repository.NewMockExpenseRepository(ctrl)
	mas := mock_service.NewMockAccountService(ctrl)
	es := NewExpenseService(mer, mas)

	t.Run("should return error when account service call returns error", func(t *testing.T) {
		mas.EXPECT().GetOneByID(gomock.Eq(2)).DoAndReturn(func(id int) (model.Account, error) {
			return model.Account{}, errors.New("")
		})
		if _, err := es.Create(1, 2, dto.CreateExpenseDTO{
			Name:        "Dinner",
			Description: "Eating out with friends",
			Amount:      120000,
		}); err == nil {
			t.Error("exp error; got nil")
		}
	})
	t.Run("should return error when repository call returns error", func(t *testing.T) {
		mas.EXPECT().GetOneByID(gomock.Eq(2)).DoAndReturn(func(id int) (model.Account, error) {
			return model.Account{
				UserID: uint(1),
			}, nil
		})
		mer.EXPECT().Insert(gomock.Eq(uint(1)), gomock.Eq(uint(2)), gomock.Eq("Dinner"), gomock.Eq("Eating out with friends"), gomock.Eq(120000)).
			DoAndReturn(func(userID uint, accountID uint, name string, description string, amount int) (model.Expense, error) {
				return model.Expense{}, errors.New("")
			})
		if _, err := es.Create(1, 2, dto.CreateExpenseDTO{
			Name:        "Dinner",
			Description: "Eating out with friends",
			Amount:      120000,
		}); err == nil {
			t.Error("exp error; got nil")
		}
	})
	t.Run("should return new expense", func(t *testing.T) {
		mas.EXPECT().GetOneByID(gomock.Eq(2)).DoAndReturn(func(id int) (model.Account, error) {
			return model.Account{
				UserID: uint(1),
			}, nil
		})
		mer.EXPECT().Insert(gomock.Eq(uint(1)), gomock.Eq(uint(2)), gomock.Eq("Dinner"), gomock.Eq("Eating out with friends"), gomock.Eq(120000)).
			DoAndReturn(func(userID uint, accountID uint, name string, description string, amount int) (model.Expense, error) {
				return model.Expense{
					UserID:      userID,
					AccountID:   accountID,
					Name:        name,
					Description: description,
					Amount:      amount,
				}, nil
			})

		exp := model.Expense{
			UserID:      uint(1),
			AccountID:   uint(2),
			Name:        "Dinner",
			Description: "Eating out with friends",
			Amount:      120000,
		}
		got, err := es.Create(1, 2, dto.CreateExpenseDTO{
			Name:        "Dinner",
			Description: "Eating out with friends",
			Amount:      120000,
		})
		if err != nil {
			t.Error("exp nil; got error:", err)
		}

		opts := []cmp.Option{
			cmpopts.IgnoreFields(model.Category{}, "Model"),
		}
		testutil.CompareAndAssert(t, exp, got, opts...)
	})
}

func TestExpenseService_GetOneByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	mer := mock_repository.NewMockExpenseRepository(ctrl)
	mas := mock_service.NewMockAccountService(ctrl)
	es := NewExpenseService(mer, mas)

	t.Run("should return expense", func(t *testing.T) {
		mer.EXPECT().GetOneByID(gomock.Eq(uint(1))).DoAndReturn(func(id uint) (model.Expense, error) {
			return model.Expense{
				Model: gorm.Model{
					ID: id,
				},
			}, nil
		})

		got, err := es.GetOneByID(1)
		if err != nil {
			t.Error("exp nil; got error:", err)
		}
		if got.Model.ID != uint(1) {
			t.Error("exp 1; got", got.Model.ID)
		}
	})
}

func TestExpenseService_GetMany(t *testing.T) {
	ctrl := gomock.NewController(t)
	mer := mock_repository.NewMockExpenseRepository(ctrl)
	mas := mock_service.NewMockAccountService(ctrl)
	es := NewExpenseService(mer, mas)

	t.Run("should return many expenses", func(t *testing.T) {
		mer.EXPECT().GetMany(gomock.Eq(10), gomock.Eq(20)).DoAndReturn(func(limit, offset int) ([]model.Expense, error) {
			return []model.Expense{
				{
					Name:        "Expense 1",
					Description: "Desc",
					Amount:      1000,
				},
				{
					Name:        "Expense 1",
					Description: "Desc",
					Amount:      1000,
				},
				{
					Name:        "Expense 1",
					Description: "Desc",
					Amount:      1000,
				},
			}, nil
		})

		got, err := es.GetMany(10, 3)
		if err != nil {
			t.Error("exp nil; got error:", err)
		}
		if len(got) < 1 {
			t.Error("exp => 1; got < 1")
		}
	})
}

func TestExpenseService_GetManyBelongedToUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	mer := mock_repository.NewMockExpenseRepository(ctrl)
	mas := mock_service.NewMockAccountService(ctrl)
	es := NewExpenseService(mer, mas)

	t.Run("should return many expenses", func(t *testing.T) {
		mer.EXPECT().GetManyBelongedToUser(gomock.Eq(uint(1)), gomock.Eq(10), gomock.Eq(20)).DoAndReturn(func(userID uint, limit, offset int) ([]model.Expense, error) {
			return []model.Expense{
				{
					Name:        "Expense 1",
					Description: "Desc",
					Amount:      1000,
				},
				{
					Name:        "Expense 1",
					Description: "Desc",
					Amount:      1000,
				},
				{
					Name:        "Expense 1",
					Description: "Desc",
					Amount:      1000,
				},
			}, nil
		})

		got, err := es.GetManyBelongedToUser(1, 10, 3)
		if err != nil {
			t.Error("exp nil; got error:", err)
		}
		if len(got) < 1 {
			t.Error("exp => 1; got < 1")
		}
	})
}

func TestExpenseService_GetManyBelongedToCategory(t *testing.T) {
	ctrl := gomock.NewController(t)
	mer := mock_repository.NewMockExpenseRepository(ctrl)
	mas := mock_service.NewMockAccountService(ctrl)
	es := NewExpenseService(mer, mas)

	t.Run("should return many expenses", func(t *testing.T) {
		mer.EXPECT().GetManyBelongedToCategory(gomock.Eq(uint(1)), gomock.Eq(uint(2)), gomock.Eq(10), gomock.Eq(20)).DoAndReturn(func(userID, categoryID uint, limit, offset int) ([]model.Expense, error) {
			return []model.Expense{
				{
					Name:        "Expense 1",
					Description: "Desc",
					Amount:      1000,
				},
				{
					Name:        "Expense 1",
					Description: "Desc",
					Amount:      1000,
				},
				{
					Name:        "Expense 1",
					Description: "Desc",
					Amount:      1000,
				},
			}, nil
		})

		got, err := es.GetManyBelongedToCategory(1, 2, 10, 3)
		if err != nil {
			t.Error("exp nil; got error:", err)
		}
		if len(got) < 1 {
			t.Error("exp => 1; got < 1")
		}
	})
}

func TestExpenseService_GetManyBelongedToAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	mer := mock_repository.NewMockExpenseRepository(ctrl)
	mas := mock_service.NewMockAccountService(ctrl)
	es := NewExpenseService(mer, mas)

	t.Run("should return many expenses", func(t *testing.T) {
		mer.EXPECT().GetManyBelongedToAccount(gomock.Eq(uint(1)), gomock.Eq(uint(2)), gomock.Eq(10), gomock.Eq(20)).DoAndReturn(func(userID uint, accountID uint, limit, offset int) ([]model.Expense, error) {
			return []model.Expense{
				{
					Name:        "Expense 1",
					Description: "Desc",
					Amount:      1000,
				},
				{
					Name:        "Expense 1",
					Description: "Desc",
					Amount:      1000,
				},
				{
					Name:        "Expense 1",
					Description: "Desc",
					Amount:      1000,
				},
			}, nil
		})

		got, err := es.GetManyBelongedToAccount(1, 2, 10, 3)
		if err != nil {
			t.Error("exp nil; got error:", err)
		}
		if len(got) < 1 {
			t.Error("exp => 1; got < 1")
		}
	})
}

func TestExpenseService_GetManyBelongedToCategoryAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	mer := mock_repository.NewMockExpenseRepository(ctrl)
	mas := mock_service.NewMockAccountService(ctrl)
	es := NewExpenseService(mer, mas)

	t.Run("should return many expenses", func(t *testing.T) {
		mer.EXPECT().GetManyBelongedToCategoryAccount(gomock.Eq(uint(1)), gomock.Eq(uint(2)), gomock.Eq(uint(3)), gomock.Eq(10), gomock.Eq(20)).DoAndReturn(func(userID, categoryID, accountID uint, limit, offset int) ([]model.Expense, error) {
			return []model.Expense{
				{
					Name:        "Expense 1",
					Description: "Desc",
					Amount:      1000,
				},
				{
					Name:        "Expense 1",
					Description: "Desc",
					Amount:      1000,
				},
				{
					Name:        "Expense 1",
					Description: "Desc",
					Amount:      1000,
				},
			}, nil
		})

		got, err := es.GetManyBelongedToCategoryAccount(1, 2, 3, 10, 3)
		if err != nil {
			t.Error("exp nil; got error:", err)
		}
		if len(got) < 1 {
			t.Error("exp => 1; got < 1")
		}
	})
}

func TestExpenseService_UpdateOneByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	mer := mock_repository.NewMockExpenseRepository(ctrl)
	mas := mock_service.NewMockAccountService(ctrl)
	es := NewExpenseService(mer, mas)

	t.Run("should return updated expense", func(t *testing.T) {
		mer.EXPECT().UpdateOneByID(gomock.Eq(uint(1)), gomock.Eq("Dinner"), gomock.Eq("Eating out with friends"), gomock.Eq(120000)).
			DoAndReturn(func(id uint, name string, description string, amount int) (model.Expense, error) {
				return model.Expense{
					Model: gorm.Model{
						ID: id,
					},
					Name:        name,
					Description: description,
					Amount:      amount,
				}, nil
			})

		exp := model.Expense{
			Model: gorm.Model{
				ID: uint(1),
			},
			Name:        "Dinner",
			Description: "Eating out with friends",
			Amount:      120000,
		}
		got, err := es.UpdateOneByID(1, dto.UpdateExpenseDTO{
			Name:        "Dinner",
			Description: "Eating out with friends",
			Amount:      120000,
		})
		if err != nil {
			t.Error("exp nil; got error:", err)
		}

		opts := []cmp.Option{
			cmpopts.IgnoreFields(model.Category{}, "Model.CreatedAt"),
			cmpopts.IgnoreFields(model.Category{}, "Model.UpdatedAt"),
			cmpopts.IgnoreFields(model.Category{}, "Model.DeletedAt"),
		}
		testutil.CompareAndAssert(t, exp, got, opts...)
	})
}

func TestExpenseService_DeleteOneByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	mer := mock_repository.NewMockExpenseRepository(ctrl)
	mas := mock_service.NewMockAccountService(ctrl)
	es := NewExpenseService(mer, mas)

	t.Run("should return error when repository returns error", func(t *testing.T) {
		mer.EXPECT().DeleteOneByID(gomock.Eq(uint(1))).DoAndReturn(func(id uint) error {
			return errors.New("")
		})

		if err := es.DeleteOneByID(1); err == nil {
			t.Error("exp error; got nil")
		}
	})
	t.Run("should delete and return nil", func(t *testing.T) {
		mer.EXPECT().DeleteOneByID(gomock.Eq(uint(1))).DoAndReturn(func(id uint) error {
			return nil
		})

		if err := es.DeleteOneByID(1); err != nil {
			t.Error("exp nil; got error:", err)
		}
	})
}
