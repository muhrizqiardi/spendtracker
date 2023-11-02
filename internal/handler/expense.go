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

type ExpenseHandler interface {
	Create(c echo.Context) error
	GetOneByID(c echo.Context) error
	GetMany(c echo.Context) error
	UpdateOneByID(c echo.Context) error
	DeleteOneByID(c echo.Context) error
}

type expenseHandler struct {
	es service.ExpenseService
}

func NewExpenseHandler(es service.ExpenseService) *expenseHandler {
	return &expenseHandler{es}
}

// @Router		/accounts/{accountID}/expenses [post]
// @Summary	Create expense
// @Tags		expense
// @Param		accountId	path		string					true	"Account ID"
// @Param		payload		body		dto.CreateExpenseDTO	true	"Create expense DTO"
// @Success	201			{object}	util.BaseResponse[response.CommonExpenseResponse]
func (eh *expenseHandler) Create(c echo.Context) error {
	accountID, err := strconv.Atoi(c.Param("accountID"))
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(
			http.StatusBadRequest,
			util.CreateBaseResponse[any](false, "Bad Request", nil),
		)
	}

	var payload dto.CreateExpenseDTO
	if err := c.Bind(&payload); err != nil {
		c.Logger().Error(err)
		return c.JSON(
			http.StatusBadRequest,
			util.CreateBaseResponse[any](false, "Bad Request", nil),
		)
	}

	user := c.Get("user").(model.User)
	expense, err := eh.es.Create(int(user.ID), accountID, payload)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(
			http.StatusInternalServerError,
			util.CreateBaseResponse[any](false, "Internal Server Error", nil),
		)
	}

	return c.JSON(
		http.StatusCreated,
		util.CreateBaseResponse[response.CommonExpenseResponse](
			true, "User created",
			response.CommonExpenseResponse{
				ID:          int(expense.ID),
				UserID:      expense.UserID,
				AccountID:   expense.AccountID,
				Name:        expense.Name,
				Description: expense.Description,
				Amount:      expense.Amount,
				CreatedAt:   expense.CreatedAt,
				UpdatedAt:   expense.UpdatedAt,
			},
		),
	)
}

// @Router		/expenses/{expenseID} [get]
// @Summary	Get one expense by ID
// @Tags		expense
// @Param		expenseID	path		string	true	"Expense ID"
// @Success	200			{object}	util.BaseResponse[response.CommonExpenseResponse]
func (eh *expenseHandler) GetOneByID(c echo.Context) error {
	expenseID, err := strconv.Atoi(c.Param("expenseID"))
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(
			http.StatusBadRequest,
			util.CreateBaseResponse[any](false, "Bad Request", nil),
		)
	}

	expense, err := eh.es.GetOneByID(expenseID)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(
			http.StatusInternalServerError,
			util.CreateBaseResponse[any](false, "Internal Server Error", nil),
		)
	}

	return c.JSON(
		http.StatusOK,
		util.CreateBaseResponse[response.CommonExpenseResponse](
			true, "Expense found",
			response.CommonExpenseResponse{
				ID:          int(expense.ID),
				UserID:      expense.UserID,
				AccountID:   expense.AccountID,
				Name:        expense.Name,
				Description: expense.Description,
				Amount:      expense.Amount,
				CreatedAt:   expense.CreatedAt,
				UpdatedAt:   expense.UpdatedAt,
			},
		),
	)
}

// @Router		/expenses [get]
// @Summary	Get many expenses
// @Tags		expense
// @Param		accountId	query		string	false	"Account ID"
// @Param		categoryId	query		string	false	"Category ID"
// @Param		itemPerPage	query		string	true	"Amount of items per page"
// @Param		page		query		string	true	"Page number"
// @Success	200			{object}	util.BaseResponse[[]response.CommonExpenseResponse]
func (eh *expenseHandler) GetMany(c echo.Context) error {
	user := c.Get("user").(model.User)
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

	if c.QueryParam("categoryId") != "" && c.QueryParam("accountId") != "" {
		categoryID, err := strconv.Atoi(c.QueryParam("categoryId"))
		if err != nil {
			c.Logger().Error(err)
			return c.JSON(
				http.StatusBadRequest,
				util.CreateBaseResponse[any](false, "Bad Request", nil),
			)
		}
		accountID, err := strconv.Atoi(c.QueryParam("accountId"))
		if err != nil {
			c.Logger().Error(err)
			return c.JSON(
				http.StatusBadRequest,
				util.CreateBaseResponse[any](false, "Bad Request", nil),
			)
		}

		expenses, err := eh.es.GetManyBelongedToCategoryAccount(int(user.ID), categoryID, accountID, itemPerPage, page)
		if err != nil {
			c.Logger().Error(err)
			return c.JSON(
				http.StatusInternalServerError,
				util.CreateBaseResponse[any](false, "Internal Server Error", nil),
			)
		}

		responses := make([]response.CommonExpenseResponse, 0, len(expenses))
		for _, e := range expenses {
			responses = append(responses, response.CommonExpenseResponse{
				ID:          int(e.ID),
				UserID:      e.UserID,
				AccountID:   e.AccountID,
				Name:        e.Name,
				Description: e.Description,
				Amount:      e.Amount,
				CreatedAt:   e.CreatedAt,
				UpdatedAt:   e.UpdatedAt,
			})
		}
		return c.JSON(
			http.StatusOK,
			util.CreateBaseResponse[[]response.CommonExpenseResponse](
				true, "Expenses found", responses,
			),
		)
	} else if c.QueryParam("accountId") == "" {
		categoryID, err := strconv.Atoi(c.QueryParam("categoryId"))
		if err != nil {
			c.Logger().Error(err)
			return c.JSON(
				http.StatusBadRequest,
				util.CreateBaseResponse[any](false, "Bad Request", nil),
			)
		}
		expenses, err := eh.es.GetManyBelongedToCategory(int(user.ID), categoryID, itemPerPage, page)
		if err != nil {
			c.Logger().Error(err)
			return c.JSON(
				http.StatusInternalServerError,
				util.CreateBaseResponse[any](false, "Internal Server Error", nil),
			)
		}

		responses := make([]response.CommonExpenseResponse, 0, len(expenses))
		for _, e := range expenses {
			responses = append(responses, response.CommonExpenseResponse{
				ID:          int(e.ID),
				UserID:      e.UserID,
				AccountID:   e.AccountID,
				Name:        e.Name,
				Description: e.Description,
				Amount:      e.Amount,
				CreatedAt:   e.CreatedAt,
				UpdatedAt:   e.UpdatedAt,
			})
		}
		return c.JSON(
			http.StatusOK,
			util.CreateBaseResponse[[]response.CommonExpenseResponse](
				true, "Expenses found", responses,
			),
		)
	} else if c.QueryParam("categoryId") == "" {
		accountID, err := strconv.Atoi(c.QueryParam("accountId"))
		if err != nil {
			c.Logger().Error(err)
			return c.JSON(
				http.StatusBadRequest,
				util.CreateBaseResponse[any](false, "Bad Request", nil),
			)
		}
		expenses, err := eh.es.GetManyBelongedToAccount(int(user.ID), accountID, itemPerPage, page)
		if err != nil {
			c.Logger().Error(err)
			return c.JSON(
				http.StatusInternalServerError,
				util.CreateBaseResponse[any](false, "Internal Server Error", nil),
			)
		}

		responses := make([]response.CommonExpenseResponse, 0, len(expenses))
		for _, e := range expenses {
			responses = append(responses, response.CommonExpenseResponse{
				ID:          int(e.ID),
				UserID:      e.UserID,
				AccountID:   e.AccountID,
				Name:        e.Name,
				Description: e.Description,
				Amount:      e.Amount,
				CreatedAt:   e.CreatedAt,
				UpdatedAt:   e.UpdatedAt,
			})
		}
		return c.JSON(
			http.StatusOK,
			util.CreateBaseResponse[[]response.CommonExpenseResponse](
				true, "Expenses found", responses,
			),
		)
	} else {
		expenses, err := eh.es.GetManyBelongedToUser(int(user.ID), itemPerPage, page)
		if err != nil {
			c.Logger().Error(err)
			return c.JSON(
				http.StatusInternalServerError,
				util.CreateBaseResponse[any](false, "Internal Server Error", nil),
			)
		}

		responses := make([]response.CommonExpenseResponse, 0, len(expenses))
		for _, e := range expenses {
			responses = append(responses, response.CommonExpenseResponse{
				ID:          int(e.ID),
				UserID:      e.UserID,
				AccountID:   e.AccountID,
				Name:        e.Name,
				Description: e.Description,
				Amount:      e.Amount,
				CreatedAt:   e.CreatedAt,
				UpdatedAt:   e.UpdatedAt,
			})
		}
		return c.JSON(
			http.StatusOK,
			util.CreateBaseResponse[[]response.CommonExpenseResponse](
				true, "Expenses found", responses,
			),
		)
	}
}

// @Router		/expenses/{expenseID} [put]
// @Summary	Update expense
// @Tags		expense
// @Param		expenseID	path		string					true	"Expense ID"
// @Param		payload		body		dto.UpdateExpenseDTO	true	"Update expense DTO"
// @Success	200			{object}	util.BaseResponse[response.CommonExpenseResponse]
func (eh *expenseHandler) UpdateOneByID(c echo.Context) error {
	expenseID, err := strconv.Atoi(c.Param("expenseID"))
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(
			http.StatusBadRequest,
			util.CreateBaseResponse[any](false, "Bad Request", nil),
		)
	}

	var payload dto.UpdateExpenseDTO
	if err := c.Bind(&payload); err != nil {
		c.Logger().Error(err)
		return c.JSON(
			http.StatusBadRequest,
			util.CreateBaseResponse[any](false, "Bad Request", nil),
		)
	}

	user := c.Get("user").(model.User)
	if expense, err := eh.es.GetOneByID(expenseID); expense.UserID != user.ID || err != nil {
		c.Logger().Error(err)
		return c.JSON(
			http.StatusForbidden,
			util.CreateBaseResponse[any](false, "Forbidden", nil),
		)
	}

	expense, err := eh.es.UpdateOneByID(expenseID, payload)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(
			http.StatusInternalServerError,
			util.CreateBaseResponse[any](false, "Internal Server Error", nil),
		)
	}

	return c.JSON(
		http.StatusOK,
		util.CreateBaseResponse[response.CommonExpenseResponse](
			true, "User updated",
			response.CommonExpenseResponse{
				ID:          int(expense.ID),
				UserID:      expense.UserID,
				AccountID:   expense.AccountID,
				Name:        expense.Name,
				Description: expense.Description,
				Amount:      expense.Amount,
				CreatedAt:   expense.CreatedAt,
				UpdatedAt:   expense.UpdatedAt,
			},
		),
	)
}

// @Router		/expenses/{expenseID} [delete]
// @Summary	Delete one expense by ID
// @Tags		expense
// @Param		expenseID	path		string	true	"Expense ID"
// @Success	200			{object}	util.BaseResponse[any]
func (eh *expenseHandler) DeleteOneByID(c echo.Context) error {
	expenseID, err := strconv.Atoi(c.Param("expenseID"))
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(
			http.StatusBadRequest,
			util.CreateBaseResponse[any](false, "Bad Request", nil),
		)
	}

	if err := eh.es.DeleteOneByID(expenseID); err != nil {
		c.Logger().Error(err)
		return c.JSON(
			http.StatusInternalServerError,
			util.CreateBaseResponse[any](false, "Internal Server Error", nil),
		)
	}

	return c.JSON(
		http.StatusOK,
		util.CreateBaseResponse[any](
			true, "Expense deleted", nil,
		),
	)
}
