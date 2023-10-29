package integration

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/muhrizqiardi/spendtracker/internal/database/model"
	"github.com/muhrizqiardi/spendtracker/internal/repository"
	"github.com/muhrizqiardi/spendtracker/tests/testutil"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupDBForUserTest() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"))
	if err != nil {
		return &gorm.DB{}, err
	}

	if err := db.AutoMigrate(
		&model.User{},
		&model.Expense{},
		&model.ReoccuringPaymentModel{},
		&model.GoalModel{},
	); err != nil {
		return &gorm.DB{}, err
	}

	return db, nil
}

func TestUserRepository_Insert(t *testing.T) {
	db, err := setupDBForUserTest()
	if err != nil {
		t.Error("exp nil; got error:", err)
	}
	ur := repository.NewUserRepository(db)
	opts := []cmp.Option{
		cmpopts.IgnoreFields(model.User{}, "Model"),
	}

	t.Run("should insert user", func(t *testing.T) {
		exp := model.User{
			Email:    "test@example.com",
			Password: "hashedpass",
		}
		got, err := ur.Insert("test@example.com", "hashedpass")
		if err != nil {
			t.Error("exp nil; got error:", err)
		}

		if diff, isEqual := testutil.Compare(exp, got, opts...); !isEqual {
			t.Errorf("mismatch (-exp, +got): %s", diff)
		}
	})
	t.Run("should return error when email already exists", func(t *testing.T) {
		if _, err := ur.Insert("test@example.com", "hashedpass"); err == nil {
			t.Error("exp error; got nil")
		}
	})
}

func TestUserRepository_GetOneByEmail(t *testing.T) {
	db, err := setupDBForUserTest()
	if err != nil {
		t.Error("exp nil; got error:", err)
	}
	ur := repository.NewUserRepository(db)
	opts := []cmp.Option{
		cmpopts.IgnoreFields(model.User{}, "Model"),
	}

	if _, err := ur.Insert("get.one.by.email@example.com", "hashedpass"); err != nil {
		t.Error("exp nil; got error:", err)
	}

	t.Run("should return error when user does not exist", func(t *testing.T) {
		if _, err := ur.GetOneByEmail("does.not.exist@example.com"); err == nil {
			t.Error("exp error; got nil")
		}
	})
	t.Run("should return user", func(t *testing.T) {
		exp := model.User{
			Email:    "get.one.by.email@example.com",
			Password: "hashedpass",
		}
		got, err := ur.GetOneByEmail("get.one.by.email@example.com")
		if err != nil {
			t.Error("exp nil; got error:", err)
		}

		if diff, isEqual := testutil.Compare(exp, got, opts...); !isEqual {
			t.Errorf("mismatch (-exp, +got): %s", diff)
		}
	})
}

func TestUserRepository_GetOneByID(t *testing.T) {
	db, err := setupDBForUserTest()
	if err != nil {
		t.Error("exp nil; got error:", err)
	}
	ur := repository.NewUserRepository(db)
	opts := []cmp.Option{
		cmpopts.IgnoreFields(model.User{}, "Model"),
	}

	mockUser, err := ur.Insert("get.one.by.email@example.com", "hashedpass")
	if err != nil {
		t.Error("exp nil; got error:", err)
	}

	t.Run("should return error when user does not exist", func(t *testing.T) {
		if _, err := ur.GetOneByID(1001); err == nil {
			t.Error("exp error; got nil")
		}
	})
	t.Run("should return user", func(t *testing.T) {
		exp := mockUser
		got, err := ur.GetOneByID(int(mockUser.ID))
		if err != nil {
			t.Error("exp nil; got error:", err)
		}

		if diff, isEqual := testutil.Compare(exp, got, opts...); !isEqual {
			t.Errorf("mismatch (-exp, +got): %s", diff)
		}
	})
}

func TestUserRepository_UpdateOneByID(t *testing.T) {
	db, err := setupDBForUserTest()
	if err != nil {
		t.Error("exp nil; got error:", err)
	}
	ur := repository.NewUserRepository(db)
	opts := []cmp.Option{
		cmpopts.IgnoreFields(model.User{}, "Model"),
	}

	mockUser, err := ur.Insert("before.update@example.com", "hashedpass")
	if err != nil {
		t.Error("exp nil; got error:", err)
	}

	t.Run("should return error if user does not exist", func(t *testing.T) {
		if _, err := ur.UpdateOneByID(1001, "does.not.exist@example.com", "hashedpass"); err == nil {
			t.Error("exp error; got nil")
		}
	})
	t.Run("should update user's email and returns its result", func(t *testing.T) {
		exp := model.User{
			Email:    "after.update@example.com",
			Password: "updatedhash",
		}
		got, err := ur.UpdateOneByID(int(mockUser.ID), "after.update@example.com", "updatedhash")
		if err != nil {
			t.Error("exp nil; got error:", err)
		}

		if diff, isEqual := testutil.Compare(exp, got, opts...); !isEqual {
			t.Errorf("mismatch (-exp, +got): %s", diff)
		}
	})
}

func TestUserRepository_DeleteOneByID(t *testing.T) {
	db, err := setupDBForUserTest()
	if err != nil {
		t.Error("exp nil; got error:", err)
	}
	ur := repository.NewUserRepository(db)

	mockUser, err := ur.Insert("to.be.deleted@example.com", "hashedpass")
	if err != nil {
		t.Error("exp nil; got error:", err)
	}

	t.Run("should return error if user does not exist", func(t *testing.T) {
		if err := ur.DeleteOneByID(1001); err == nil {
			t.Error("exp error; got nil")
		}
	})
	t.Run("should delete user ", func(t *testing.T) {
		if err := ur.DeleteOneByID(int(mockUser.ID)); err != nil {
			t.Error("exp nil; got error:", err)
		}

		if _, err := ur.GetOneByID(int(mockUser.ID)); err == nil {
			t.Error("exp error; got nil")
		}
	})
}
