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

func TestCategoryService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	mcr := mock_repository.NewMockCategoryRepository(ctrl)
	cs := NewCategoryService(mcr)

	t.Run("should return new category", func(t *testing.T) {
		mcr.EXPECT().Insert(gomock.Eq(uint(1)), gomock.Eq("Bill")).DoAndReturn(func(userID uint, name string) (model.Category, error) {
			return model.Category{
				Model:  gorm.Model{},
				UserID: uint(userID),
				Name:   name,
			}, nil
		})

		exp := model.Category{
			UserID: 1,
			Name:   "Bill",
		}
		got, err := cs.Create(1, dto.CreateCategoryDTO{
			Name: "Bill",
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

func TestCategoryService_GetOneByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	mcr := mock_repository.NewMockCategoryRepository(ctrl)
	cs := NewCategoryService(mcr)

	t.Run("should return category", func(t *testing.T) {
		mcr.EXPECT().GetOneByID(gomock.Eq(uint(1))).DoAndReturn(func(id uint) (model.Category, error) {
			return model.Category{
				Model: gorm.Model{
					ID: id,
				},
			}, nil
		})

		exp := model.Category{
			Model: gorm.Model{
				ID: 1,
			},
		}
		got, err := cs.GetOneByID(1)
		if err != nil {
			t.Error("exp nil; got error:", err)
		}
		opts := []cmp.Option{
			cmpopts.IgnoreFields(model.Category{}, "Model.CreatedAt"),
			cmpopts.IgnoreFields(model.Category{}, "Model.UpdatedAt"),
			cmpopts.IgnoreFields(model.Category{}, "Model.DeletedAt"),
			cmpopts.IgnoreFields(model.Category{}, "Name"),
			cmpopts.IgnoreFields(model.Category{}, "UserID"),
		}
		testutil.CompareAndAssert(t, exp, got, opts...)
	})
}

func TestCategoryService_GetMany(t *testing.T) {
	ctrl := gomock.NewController(t)
	mcr := mock_repository.NewMockCategoryRepository(ctrl)
	cs := NewCategoryService(mcr)

	t.Run("should return categories", func(t *testing.T) {
		mcr.EXPECT().GetMany(gomock.Eq(uint(1)), gomock.Eq(10), gomock.Eq(10)).DoAndReturn(func(userID uint, itemPerPage, page int) ([]model.Category, error) {
			return []model.Category{
				{
					UserID: 1,
					Name:   "Bill",
				},
				{
					UserID: 1,
					Name:   "Rent",
				},
				{
					UserID: 1,
					Name:   "Entertainment",
				},
			}, nil
		})

		exp := []model.Category{
			{
				UserID: 1,
				Name:   "Bill",
			},
			{
				UserID: 1,
				Name:   "Rent",
			},
			{
				UserID: 1,
				Name:   "Entertainment",
			},
		}
		got, err := cs.GetMany(1, 10, 2)
		if err != nil {
			t.Error("exp nil; got error:", err)
		}
		opts := []cmp.Option{
			cmpopts.IgnoreFields(model.Category{}, "Model"),
		}
		testutil.CompareAndAssert(t, exp, got, opts...)
	})
}

func TestCategoryService_DeleteOneByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	mcr := mock_repository.NewMockCategoryRepository(ctrl)
	cs := NewCategoryService(mcr)

	t.Run("should return error when repository returns error", func(t *testing.T) {
		mcr.EXPECT().Delete(gomock.Eq(uint(1))).DoAndReturn(func(id uint) error {
			return errors.New("")
		})

		if err := cs.DeleteOneByID(1); err == nil {
			t.Error("exp error; got nil")
		}
	})
	t.Run("should delete category", func(t *testing.T) {
		mcr.EXPECT().Delete(gomock.Eq(uint(1))).DoAndReturn(func(id uint) error {
			return nil
		})

		if err := cs.DeleteOneByID(1); err != nil {
			t.Error("exp nil; got error:", err)
		}
	})
}
