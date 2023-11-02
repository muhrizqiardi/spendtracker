package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/muhrizqiardi/spendtracker/internal/dto"
	"github.com/muhrizqiardi/spendtracker/internal/response"
	"github.com/muhrizqiardi/spendtracker/internal/service"
	"github.com/muhrizqiardi/spendtracker/internal/util"
)

type AuthHandler interface {
	LogIn(c echo.Context) error
}

type authHandler struct {
	as service.AuthService
}

func NewAuthHandler(as service.AuthService) *authHandler {
	return &authHandler{as}
}

// LogIn
//
//	@Router		/auth [post]
//	@Summary	Log in to account
//	@Tags		auth
//	@Param		payload	body		dto.LogInDTO	true	"log in DTO"
//	@Success	200		{object}	response.LogInResponse
func (ah *authHandler) LogIn(c echo.Context) error {
	var payload dto.LogInDTO
	if err := c.Bind(&payload); err != nil {
		c.Logger().Error(err)
		return c.JSON(
			http.StatusBadRequest,
			util.CreateBaseResponse[any](false, "Bad Request", nil),
		)
	}

	token, err := ah.as.LogIn(payload)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(
			http.StatusInternalServerError,
			util.CreateBaseResponse[any](false, "Internal Server Error", nil),
		)
	}

	// return c.JSON(200, util.CreateBaseResponse[response.log]()(response.LogInResponse{Token: token}))
	return c.JSON(
		http.StatusCreated,
		util.CreateBaseResponse[response.LogInResponse](
			true,
			"Log in success",
			response.LogInResponse{
				Token: token,
			},
		),
	)
}
