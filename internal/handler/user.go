package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/muhrizqiardi/spendtracker/internal/database/model"
	"github.com/muhrizqiardi/spendtracker/internal/dto"
	"github.com/muhrizqiardi/spendtracker/internal/response"
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

// @Router		/users [post]
// @Summary	Register user
// @Tags		user
// @Param		payload	body		dto.RegisterUserDTO	true	"Create user DTO"
// @Success	201		{object}	util.BaseResponse[response.CommonUserResponse]
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
		util.CreateBaseResponse[response.CommonUserResponse](true, "User registered",
			response.CommonUserResponse{
				ID:        int(user.ID),
				Email:     user.Email,
				FullName:  user.FullName,
				CreatedAt: user.CreatedAt,
				UpdatedAt: user.UpdatedAt,
			},
		),
	)
}

func (uh *userHandler) GetOneByID(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(
			http.StatusNotFound,
			util.CreateBaseResponse[any](false, "Not Found", nil),
		)
	}

	user, err := uh.us.GetOneByID(userID)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(
			http.StatusNotFound,
			util.CreateBaseResponse[any](false, "Not Found", nil),
		)
	}

	return c.JSON(
		http.StatusOK,
		util.CreateBaseResponse[model.User](true, "User found", user),
	)
}

// @Router		/users/{userID} [put]
// @Summary	Update user
// @Tags		user
// @Param		userID	path		string				true	"User ID"
// @Param		payload	body		dto.UpdateUserDTO	true	"Update user DTO"
// @Success	200		{object}	util.BaseResponse[response.CommonUserResponse]
func (uh *userHandler) UpdateOneByID(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(
			http.StatusNotFound,
			util.CreateBaseResponse[any](false, "Not Found", nil),
		)
	}

	var payload dto.UpdateUserDTO
	if err := c.Bind(&payload); err != nil {
		fmt.Println("bebek")
		c.Logger().Error(err)
		return c.JSON(
			http.StatusBadRequest,
			util.CreateBaseResponse[any](false, "Bad Request", nil),
		)
	}

	user, err := uh.us.UpdateOneByID(userID, payload)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(
			http.StatusInternalServerError,
			util.CreateBaseResponse[any](false, "Internal Server Error", nil),
		)
	}

	return c.JSON(
		http.StatusCreated,
		util.CreateBaseResponse[response.CommonUserResponse](true, "User updated",
			response.CommonUserResponse{
				ID:        int(user.ID),
				Email:     user.Email,
				FullName:  user.FullName,
				CreatedAt: user.CreatedAt,
				UpdatedAt: user.UpdatedAt,
			},
		),
	)
}
