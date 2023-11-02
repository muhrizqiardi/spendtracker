package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/muhrizqiardi/spendtracker/internal/database/model"
	"github.com/muhrizqiardi/spendtracker/internal/dto"
	"github.com/muhrizqiardi/spendtracker/internal/response"
	"github.com/muhrizqiardi/spendtracker/internal/service"
	"github.com/muhrizqiardi/spendtracker/internal/util"
)

type AccountHandler interface {
	Create(c echo.Context) error
	GetOneByID(c echo.Context) error
	GetMany(c echo.Context) error
	UpdateOneByID(c echo.Context) error
	DeleteOneByID(c echo.Context) error
}

type accountHandler struct {
	as service.AccountService
}

func NewAccountHandler(as service.AccountService) *accountHandler {
	return &accountHandler{as}
}

// @Router		/accounts [post]
// @Summary	Create to account
// @Tags		account
// @Param		payload	body		dto.CreateAccountDTO	true	"Create account DTO"
// @Success	201		{object}	util.BaseResponse[response.CommonAccountResponse]
func (ah *accountHandler) Create(c echo.Context) error {
	var payload dto.CreateAccountDTO
	if err := c.Bind(&payload); err != nil {
		c.Logger().Error(err)
		return c.JSON(
			http.StatusBadRequest,
			util.CreateBaseResponse[any](false, "Bad Request", nil),
		)
	}

	user := c.Get("user").(model.User)

	account, err := ah.as.Create(int(user.ID), payload)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(
			http.StatusInternalServerError,
			util.CreateBaseResponse[any](false, "Internal Server Error", nil),
		)
	}

	return c.JSON(
		http.StatusCreated,
		util.CreateBaseResponse[response.CommonAccountResponse](
			true,
			"Account created",
			response.CommonAccountResponse{
				ID:            account.ID,
				UserID:        account.UserID,
				CurrencyID:    account.CurrencyID,
				Name:          account.Name,
				InitialAmount: account.InitialAmount,
			},
		),
	)
}

// @Router		/accounts/{accountID} [get]
// @Summary	Get one by ID
// @Tags		account
// @Param		accountID	path		string	true	"Account ID"
// @Success	200			{object}	util.BaseResponse[response.CommonAccountResponse]
func (ah *accountHandler) GetOneByID(c echo.Context) error {
	accountID, err := strconv.Atoi(c.Param("accountID"))
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(
			http.StatusBadRequest,
			util.CreateBaseResponse[any](false, "Bad Request", nil),
		)
	}

	account, err := ah.as.GetOneByID(accountID)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(
			http.StatusInternalServerError,
			util.CreateBaseResponse[any](false, "Internal Server Error", nil),
		)
	}

	user := c.Get("user").(model.User)
	if user.ID != account.UserID {
		c.Logger().Error("Not Allowed")
		return c.JSON(
			http.StatusForbidden,
			util.CreateBaseResponse[any](false, "Forbidden", nil),
		)
	}

	return c.JSON(
		http.StatusOK,
		util.CreateBaseResponse[model.Account](true, "Account found", account),
	)
}

// @Router		/accounts/ [get]
// @Summary	Get many
// @Tags		account
// @Param		itemPerPage	query		string	false	"Amount of items per page"
// @Param		page		query		string	false	"Page number"
// @Success	200			{object}	util.BaseResponse[[]response.CommonAccountResponse]
func (ah *accountHandler) GetMany(c echo.Context) error {
	itemPerPage := 10
	page := 1
	var err error
	itemPerPage, err = strconv.Atoi(c.QueryParam("itemPerPage"))
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(
			http.StatusBadRequest,
			util.CreateBaseResponse[any](false, "Bad Request", nil),
		)
	}
	page, err = strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(
			http.StatusBadRequest,
			util.CreateBaseResponse[any](false, "Bad Request", nil),
		)
	}

	user := c.Get("user").(model.User)
	account, err := ah.as.GetMany(int(user.ID), itemPerPage, page)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(
			http.StatusInternalServerError,
			util.CreateBaseResponse[any](false, "Internal Server Error", nil),
		)
	}

	return c.JSON(
		http.StatusOK,
		util.CreateBaseResponse[[]model.Account](true, "Account(s) found", account),
	)
}

// @Router		/accounts/{accountID} [put]
// @Summary	Update account
// @Tags		account
// @Param		accountID	path		string					true	"Account ID"
// @Param		payload		body		dto.UpdateAccountDTO true	"Update account DTO"
// @Success	200			{object}	util.BaseResponse[response.CommonAccountResponse]
func (ah *accountHandler) UpdateOneByID(c echo.Context) error {
	var payload dto.UpdateAccountDTO
	if err := c.Bind(&payload); err != nil {
		c.Logger().Error(err)
		return c.JSON(
			http.StatusBadRequest,
			util.CreateBaseResponse[any](false, "Bad Request", nil),
		)
	}

	user := c.Get("user").(model.User)

	account, err := ah.as.UpdateOneByID(int(user.ID), payload)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(
			http.StatusInternalServerError,
			util.CreateBaseResponse[any](false, "Internal Server Error", nil),
		)
	}

	return c.JSON(
		http.StatusCreated,
		util.CreateBaseResponse[model.Account](true, "Account updated", account),
	)
}

// @Router		/accounts/{accountID} [delete]
// @Summary	Delete account
// @Tags		account
// @Param		accountID	path		string	true	"Account ID"
// @Success	200			{object}	util.BaseResponse[any]
func (ah *accountHandler) DeleteOneByID(c echo.Context) error {
	accountID, err := strconv.Atoi(c.Param("accountID"))
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(
			http.StatusBadRequest,
			util.CreateBaseResponse[any](false, "Bad Request", nil),
		)
	}

	if err := ah.as.DeleteOneByID(accountID); err != nil {
		c.Logger().Error(err)
		return c.JSON(
			http.StatusInternalServerError,
			util.CreateBaseResponse[any](false, "Internal Server Error", nil),
		)
	}

	return nil
}
