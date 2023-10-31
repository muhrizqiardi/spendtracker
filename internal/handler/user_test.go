package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/labstack/echo/v4"
	"github.com/muhrizqiardi/spendtracker/internal/database/model"
	"github.com/muhrizqiardi/spendtracker/internal/dto"
	mock_service "github.com/muhrizqiardi/spendtracker/internal/service/mock"
	"github.com/muhrizqiardi/spendtracker/internal/util"
	"github.com/muhrizqiardi/spendtracker/tests/testutil"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func TestUserHandler_Register(t *testing.T) {
	ctrl := gomock.NewController(t)
	mus := mock_service.NewMockUserService(ctrl)
	uh := NewUserHandler(mus)
	t.Run("should return error 400 when body is invalid", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader("invalidbody"))
		r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		w := httptest.NewRecorder()
		c := e.NewContext(r, w)

		uh.Register(c)

		exp := http.StatusBadRequest
		got := w.Code
		if exp != got {
			t.Errorf("exp %d; got %d", exp, got)
		}
	})
	t.Run("should return error 500 when service layer returns error", func(t *testing.T) {
		mus.EXPECT().Register(gomock.Eq(dto.RegisterUserDTO{
			Email:    "test@example.com",
			FullName: "Fulan",
			Password: "topsecret",
		})).DoAndReturn(func(payload dto.RegisterUserDTO) (model.User, error) {
			return model.User{}, errors.New("")
		})

		e := echo.New()
		validBodyJSON := `{"email":"test@example.com","fullName":"Fulan","password":"topsecret"}`
		r := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(validBodyJSON))
		r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		w := httptest.NewRecorder()
		c := e.NewContext(r, w)

		uh.Register(c)

		exp := http.StatusInternalServerError
		got := w.Code
		if exp != got {
			t.Errorf("exp %d; got %d", exp, got)
		}
	})
	t.Run("should return new user", func(t *testing.T) {
		mus.EXPECT().Register(gomock.Eq(dto.RegisterUserDTO{
			Email:    "test@example.com",
			FullName: "Fulan",
			Password: "topsecret",
		})).DoAndReturn(func(payload dto.RegisterUserDTO) (model.User, error) {
			return model.User{
				Email:    "test@example.com",
				FullName: "Fulan",
				Password: "hashed-topsecret",
			}, nil
		})

		e := echo.New()
		validBodyJSON := `{"email":"test@example.com","fullName":"Fulan","password":"topsecret"}`
		r := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(validBodyJSON))
		r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		w := httptest.NewRecorder()
		c := e.NewContext(r, w)

		uh.Register(c)

		{
			exp := http.StatusCreated
			got := w.Code
			if exp != got {
				t.Errorf("exp %d; got %d", exp, got)
			}
		}
		{
			var resBody util.BaseResponse[model.User]
			if err := json.Unmarshal([]byte(w.Body.String()), &resBody); err != nil {
				t.Error("exp nil; got error:", err)
			}
			exp := model.User{
				Email:    "test@example.com",
				FullName: "Fulan",
			}
			got := resBody.Data
			opts := []cmp.Option{
				cmpopts.IgnoreFields(model.User{}, "Model"),
				cmpopts.IgnoreFields(model.User{}, "Password"),
			}
			testutil.CompareAndAssert(t, exp, got, opts...)
		}
	})
}

func TestUserHandler_GetOneByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	mus := mock_service.NewMockUserService(ctrl)
	uh := NewUserHandler(mus)

	t.Run("should return error 404 when URL param is invalid", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest(http.MethodGet, "/users", strings.NewReader("invalidbody"))
		r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		w := httptest.NewRecorder()
		c := e.NewContext(r, w)
		c.SetParamNames("userID")
		c.SetParamValues("invaliduser")

		uh.GetOneByID(c)

		exp := http.StatusNotFound
		got := w.Code
		if exp != got {
			t.Errorf("exp %d; got %d", exp, got)
		}
	})
	t.Run("should return error 404 when service layer returns error", func(t *testing.T) {
		mus.EXPECT().GetOneByID(gomock.Eq(1)).DoAndReturn(func(id int) (model.User, error) {
			return model.User{}, errors.New("")
		})

		e := echo.New()
		r := httptest.NewRequest(http.MethodGet, "/users/1", nil)
		r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		w := httptest.NewRecorder()
		c := e.NewContext(r, w)
		c.SetParamNames("userID")
		c.SetParamValues("1")

		uh.GetOneByID(c)

		exp := http.StatusNotFound
		got := w.Code
		if exp != got {
			t.Errorf("exp %d; got %d", exp, got)
		}
	})
	t.Run("should return user", func(t *testing.T) {
		mus.EXPECT().GetOneByID(gomock.Eq(1)).DoAndReturn(func(id int) (model.User, error) {
			return model.User{
				Model: gorm.Model{
					ID: uint(id),
				},
			}, nil
		})

		e := echo.New()
		r := httptest.NewRequest(http.MethodGet, "/users/1", nil)
		r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		w := httptest.NewRecorder()
		c := e.NewContext(r, w)
		c.SetParamNames("userID")
		c.SetParamValues("1")

		uh.GetOneByID(c)

		{
			exp := http.StatusOK
			got := w.Code
			if exp != got {
				t.Errorf("exp %d; got %d", exp, got)
			}
		}
		{
			var resBody util.BaseResponse[model.User]
			if err := json.Unmarshal([]byte(w.Body.String()), &resBody); err != nil {
				t.Error("exp nil; got error:", err)
			}
			exp := model.User{
				Model: gorm.Model{
					ID: uint(1),
				},
			}
			got := resBody.Data
			opts := []cmp.Option{
				cmpopts.IgnoreFields(model.User{}, "Model.CreatedAt"),
				cmpopts.IgnoreFields(model.User{}, "Model.UpdatedAt"),
				cmpopts.IgnoreFields(model.User{}, "Model.DeletedAt"),
				cmpopts.IgnoreFields(model.User{}, "Email"),
				cmpopts.IgnoreFields(model.User{}, "FullName"),
				cmpopts.IgnoreFields(model.User{}, "Password"),
			}
			testutil.CompareAndAssert(t, exp, got, opts...)
		}
	})
}

func Test_UpdateOneByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	mus := mock_service.NewMockUserService(ctrl)
	uh := NewUserHandler(mus)

	t.Run("should return error 404 when URL param is invalid", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest(http.MethodGet, "/users/invaliduser", nil)
		r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		w := httptest.NewRecorder()
		c := e.NewContext(r, w)
		c.SetParamNames("userID")
		c.SetParamValues("invaliduser")

		uh.UpdateOneByID(c)

		exp := http.StatusNotFound
		got := w.Code
		if exp != got {
			t.Errorf("exp %d; got %d", exp, got)
		}
	})
	t.Run("should return error 400 when body is invalid", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest(http.MethodPut, "/users/1", strings.NewReader("invalidbody"))
		r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		w := httptest.NewRecorder()
		c := e.NewContext(r, w)
		c.SetParamNames("userID")
		c.SetParamValues("1")

		uh.UpdateOneByID(c)

		exp := http.StatusBadRequest
		got := w.Code
		if exp != got {
			t.Errorf("exp %d; got %d", exp, got)
		}
	})
	t.Run("should return error 500 when service layer return error", func(t *testing.T) {
		mus.EXPECT().UpdateOneByID(gomock.Eq(1), gomock.Eq(dto.UpdateUserDTO{
			Email:    "test@example.com",
			FullName: "Fulan",
			Password: "topsecret",
		})).DoAndReturn(func(id int, payload dto.UpdateUserDTO) (model.User, error) {
			return model.User{}, errors.New("")
		})

		e := echo.New()
		validBodyJSON := `{"email":"test@example.com","fullName":"Fulan","password":"topsecret"}`
		r := httptest.NewRequest(http.MethodPut, "/users/1", strings.NewReader(validBodyJSON))
		r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		w := httptest.NewRecorder()
		c := e.NewContext(r, w)
		c.SetParamNames("userID")
		c.SetParamValues("1")

		uh.UpdateOneByID(c)

		exp := http.StatusInternalServerError
		got := w.Code
		if exp != got {
			t.Errorf("exp %d; got %d", exp, got)
		}
	})
	t.Run("should return response with user", func(t *testing.T) {
		mus.EXPECT().UpdateOneByID(gomock.Eq(1), gomock.Eq(dto.UpdateUserDTO{
			Email:    "test@example.com",
			FullName: "Fulan",
			Password: "topsecret",
		})).DoAndReturn(func(id int, payload dto.UpdateUserDTO) (model.User, error) {
			return model.User{
				Model: gorm.Model{
					ID: uint(id),
				},
				Email:    payload.Email,
				FullName: payload.FullName,
				Password: payload.Password,
			}, nil
		})

		e := echo.New()
		validBodyJSON := `{"email":"test@example.com","fullName":"Fulan","password":"topsecret"}`
		r := httptest.NewRequest(http.MethodPut, "/users/1", strings.NewReader(validBodyJSON))
		r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		w := httptest.NewRecorder()
		c := e.NewContext(r, w)
		c.SetParamNames("userID")
		c.SetParamValues("1")

		uh.UpdateOneByID(c)

		{
			exp := http.StatusOK
			got := w.Code
			if exp != got {
				t.Errorf("exp %d; got %d", exp, got)
			}
		}
		{
			var resBody util.BaseResponse[model.User]
			if err := json.Unmarshal([]byte(w.Body.String()), &resBody); err != nil {
				t.Error("exp nil; got error:", err)
			}
			exp := model.User{
				Model: gorm.Model{
					ID: uint(1),
				},
				Email:    "test@example.com",
				FullName: "Fulan",
			}
			got := resBody.Data
			opts := []cmp.Option{
				cmpopts.IgnoreFields(model.User{}, "Model.CreatedAt"),
				cmpopts.IgnoreFields(model.User{}, "Model.UpdatedAt"),
				cmpopts.IgnoreFields(model.User{}, "Model.DeletedAt"),
				cmpopts.IgnoreFields(model.User{}, "Password"),
			}
			testutil.CompareAndAssert(t, exp, got, opts...)
		}
	})
}
