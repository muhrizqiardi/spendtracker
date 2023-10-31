package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/muhrizqiardi/spendtracker/internal/dto"
	"github.com/muhrizqiardi/spendtracker/internal/response"
	mock_service "github.com/muhrizqiardi/spendtracker/internal/service/mock"
	"github.com/muhrizqiardi/spendtracker/internal/util"
	"github.com/muhrizqiardi/spendtracker/tests/testutil"
	"go.uber.org/mock/gomock"
)

func TestAuthHandler_LogIn(t *testing.T) {
	ctrl := gomock.NewController(t)
	mas := mock_service.NewMockAuthService(ctrl)
	ah := NewAuthHandler(mas)

	t.Run("should return error when body is invalid", func(t *testing.T) {
		e := echo.New()
		invalidLogInDTOJSON := `{"email": "email@example.com",password:"88888888"}`
		r := httptest.NewRequest(http.MethodPost, "/auth", strings.NewReader(invalidLogInDTOJSON))
		r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		w := httptest.NewRecorder()
		c := e.NewContext(r, w)

		ah.LogIn(c)

		exp := http.StatusBadRequest
		got := w.Code
		if exp != got {
			t.Errorf("exp %v; got %v", exp, got)
		}
	})
	t.Run("should return error when service layer returns error", func(t *testing.T) {
		mas.EXPECT().LogIn(gomock.Eq(dto.LogInDTO{
			Email:    "test@example.com",
			Password: "topsecret",
		})).DoAndReturn(func(payload dto.LogInDTO) (string, error) {
			return "", errors.New("")
		})

		e := echo.New()
		validLogInDTOJSON := `{"email": "test@example.com","password":"topsecret"}`
		req := httptest.NewRequest(http.MethodPost, "/auth", strings.NewReader(validLogInDTOJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		ah.LogIn(c)

		fmt.Printf("w: %v\n", rec)

		exp := http.StatusInternalServerError
		got := rec.Result().StatusCode
		if exp != got {
			t.Errorf("exp %v; got %v", exp, got)
		}
	})
	t.Run("should return token", func(t *testing.T) {
		mas.EXPECT().LogIn(gomock.Eq(dto.LogInDTO{
			Email:    "test@example.com",
			Password: "topsecret",
		})).DoAndReturn(func(payload dto.LogInDTO) (string, error) {
			return "mocktoken", nil
		})

		e := echo.New()
		validLogInDTOJSON := `{"email": "test@example.com","password":"topsecret"}`
		r := httptest.NewRequest(http.MethodPost, "/auth", strings.NewReader(validLogInDTOJSON))
		r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		w := httptest.NewRecorder()
		c := e.NewContext(r, w)

		ah.LogIn(c)

		var resBody util.BaseResponse[response.LogInResponse]
		if err := json.Unmarshal([]byte(w.Body.String()), &resBody); err != nil {
			t.Error("exp nil; got error:", err)
		}
		exp := response.LogInResponse{
			Token: "mocktoken",
		}
		got := resBody.Data
		testutil.CompareAndAssert(t, exp, got)
	})
}
