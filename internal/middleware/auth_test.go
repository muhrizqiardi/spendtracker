package middleware

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/labstack/echo/v4"
	"github.com/muhrizqiardi/spendtracker/internal/database/model"
	mock_service "github.com/muhrizqiardi/spendtracker/internal/service/mock"
	"github.com/muhrizqiardi/spendtracker/tests/testutil"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

var mockSecret string = "DO_NOT_USE_THIS"

func TestAuthMiddleware_Authenticate(t *testing.T) {
	ctrl := gomock.NewController(t)
	mus := mock_service.NewMockUserService(ctrl)
	am := NewAuthMiddleware(mus, mockSecret)

	t.Run("should return error if `Authorization` header is invalid", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodGet, "/users/42", nil)
		r.Header.Set("Authorization", "Bearerincorrect")
		w := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(r, w)
		c.Handler()

		got := am.Authenticate(func(c echo.Context) error {
			return c.String(200, "OK")
		})(c)

		if got == nil {
			t.Error("exp error; got nil")
		}
	})

	t.Run("should return error if jwt.Parse call returns error", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodGet, "/users/42", nil)
		r.Header.Set("Authorization", "Bearer invalidToken")
		w := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(r, w)
		c.Handler()

		got := am.Authenticate(func(c echo.Context) error {
			return c.String(200, "OK")
		})(c)

		if got == nil {
			t.Error("exp error; got nil")
		}
	})

	t.Run("should return error if token is invalid", func(t *testing.T) {
		expiredClaims := jwt.RegisteredClaims{
			Subject:   "1",
			ExpiresAt: jwt.NewNumericDate(time.Date(1984, time.January, 1, 1, 0, 0, 0, time.UTC)),
		}
		expiredToken := jwt.NewWithClaims(jwt.SigningMethodHS256, expiredClaims)
		expiredSS, _ := expiredToken.SignedString([]byte(mockSecret))

		r := httptest.NewRequest(http.MethodGet, "/users/42", nil)
		r.Header.Set("Authorization", "Bearer "+expiredSS)
		w := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(r, w)
		c.Handler()

		got := am.Authenticate(func(c echo.Context) error {
			return nil
		})(c)

		if got == nil {
			t.Error("exp error; got nil")
		}
	})

	t.Run("should proceed request/call the `next` handler set context's `user` value", func(t *testing.T) {
		mus.
			EXPECT().
			GetOneByID(gomock.Eq(42)).
			DoAndReturn(
				func(id int) (model.User, error) {
					return model.User{
						Model: gorm.Model{
							ID: uint(id),
						},
					}, nil
				},
			)

		validClaims := jwt.RegisteredClaims{
			Subject:   "42",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
		}
		validToken := jwt.NewWithClaims(jwt.SigningMethodHS256, validClaims)
		validSS, _ := validToken.SignedString([]byte(mockSecret))

		r := httptest.NewRequest(http.MethodGet, "/users/42", nil)
		r.Header.Set("Authorization", "Bearer "+validSS)
		w := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(r, w)
		c.Handler()

		got := am.Authenticate(func(c echo.Context) error {
			// assert context's value
			exp := model.User{
				Model: gorm.Model{
					ID: 42,
				},
			}
			got := c.Get("user")
			fmt.Printf("got: %v\n", got)

			opts := []cmp.Option{
				cmpopts.IgnoreFields(model.User{}, "Model.CreatedAt"),
				cmpopts.IgnoreFields(model.User{}, "Model.UpdatedAt"),
				cmpopts.IgnoreFields(model.User{}, "Model.DeletedAt"),
				cmpopts.IgnoreFields(model.User{}, "Email"),
				cmpopts.IgnoreFields(model.User{}, "FullName"),
				cmpopts.IgnoreFields(model.User{}, "Password"),
			}
			testutil.CompareAndAssert(t, exp, got, opts...)

			return nil
		})(c)

		if got != nil {
			t.Error("exp nil; got error:", got)
		}

	})
}
