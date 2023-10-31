package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/muhrizqiardi/spendtracker/internal/dto"
	"github.com/muhrizqiardi/spendtracker/internal/service"
	"github.com/muhrizqiardi/spendtracker/internal/util"
)

type UserHandler interface {
	Register(c echo.Context) error
	GetOneByID(c echo.Context) error
	UpdateOneByID(c echo.Context) error
}

type userHandler struct {
	us service.UserService
}

func NewUserHandler(us service.UserService) *userHandler {
	return &userHandler{us}
}

func (uh *userHandler) Register(c echo.Context) error {
	var payload dto.RegisterUserDTO
	if err := c.Bind(&payload); err != nil {
		c.Logger().Error(err)
		return c.JSON(
			http.StatusBadRequest,
			util.CreateBaseResponse[any](false, "Bad Request", nil),
		)
	}

	user, err := uh.us.Register(payload)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(
			http.StatusInternalServerError,
			util.CreateBaseResponse[any](false, "Internal Server Error", nil),
		)
	}

	return c.JSON(
		http.StatusCreated,
		util.CreateBaseResponse[any](true, "User Created", user),
	)
}

func (uh *userHandler) GetOneByID(c echo.Context) error {
	return nil
}

func (uh *userHandler) UpdateOneByID(c echo.Context) error {
	return nil
}
