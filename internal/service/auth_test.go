package service

import (
	"errors"
	"testing"

	"github.com/muhrizqiardi/spendtracker/internal/database/model"
	"github.com/muhrizqiardi/spendtracker/internal/dto"
	mock_service "github.com/muhrizqiardi/spendtracker/internal/service/mock"
	"go.uber.org/mock/gomock"
)

func TestAuthService_LogIn(t *testing.T) {
	ctrl := gomock.NewController(t)
	mus := mock_service.NewMockUserService(ctrl)
	as := NewAuthService(mus, "mocksecret")

	t.Run("should return error if UserService returns error", func(t *testing.T) {
		mus.EXPECT().GetOneByEmail(gomock.Eq("email@example.com")).DoAndReturn(
			func(email string) (model.User, error) {
				return model.User{}, errors.New("")
			},
		)

		if _, err := as.LogIn(dto.LogInDTO{
			Email:    "email@example.com",
			Password: "topsecret",
		}); err == nil {
			t.Error("exp error; got nil")
		}
	})
	t.Run("should return token", func(t *testing.T) {
		mus.EXPECT().GetOneByEmail(gomock.Eq("email@example.com")).DoAndReturn(
			func(email string) (model.User, error) {
				return model.User{
					Email:    email,
					Password: "$2a$12$htC6KUeMQ10/mBdUoVeRp.UW47NYED2gMG.mF/7oJ39p02XPJvuI2",
				}, nil
			},
		)

		got, err := as.LogIn(dto.LogInDTO{
			Email:    "email@example.com",
			Password: "topsecret",
		})
		if err != nil {
			t.Error("exp nil; got error:", err)
		}
		if got == "" {
			t.Error(`exp ""; got`, got)
		}
	})
}
