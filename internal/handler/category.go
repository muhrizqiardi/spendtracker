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

type CategoryHandler interface {
	Create(c echo.Context) error
	GetOneByID(c echo.Context) error
	GetMany(c echo.Context) error
	DeleteOneByID(c echo.Context) error
}

type categoryHandler struct {
	cs service.CategoryService
}

func NewCategoryHandler(cs service.CategoryService) *categoryHandler {
	return &categoryHandler{cs}
}

//	@Router		/categories [post]
//	@Summary	Create category
//	@Tags		category
//	@Param		payload	body	dto.CreateCategoryDTO	true	"Create category DTO"
//	@Security	Bearer
//	@Success	201	{object}	util.BaseResponse[response.CommonCategoryResponse]
func (ch *categoryHandler) Create(c echo.Context) error {
	var payload dto.CreateCategoryDTO
	if err := c.Bind(&payload); err != nil {
		c.Logger().Error(err)
		return c.JSON(
			http.StatusBadRequest,
			util.CreateBaseResponse[any](false, "Bad Request", nil),
		)
	}

	user := c.Get("user").(model.User)

	category, err := ch.cs.Create(int(user.ID), payload)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(
			http.StatusInternalServerError,
			util.CreateBaseResponse[any](false, "Internal Server Error", nil),
		)
	}

	return c.JSON(
		http.StatusCreated,
		util.CreateBaseResponse[response.CommonCategoryResponse](
			true,
			"Category created",
			response.CommonCategoryResponse{
				ID:        category.ID,
				UserID:    category.UserID,
				Name:      category.Name,
				CreatedAt: category.CreatedAt,
				UpdatedAt: category.UpdatedAt,
			},
		),
	)
}

//	@Router		/categories/{categoryID} [get]
//	@Summary	Get one category by ID
//	@Tags		category
//
//	@Security	Bearer
//
//	@Success	200	{object}	util.BaseResponse[response.CommonCategoryResponse]
func (ch *categoryHandler) GetOneByID(c echo.Context) error {
	categoryID, err := strconv.Atoi(c.Param("categoryID"))
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(
			http.StatusBadRequest,
			util.CreateBaseResponse[any](false, "Bad Request", nil),
		)
	}

	category, err := ch.cs.GetOneByID(categoryID)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(
			http.StatusInternalServerError,
			util.CreateBaseResponse[any](false, "Internal Server Error", nil),
		)
	}

	user := c.Get("user").(model.User)
	if user.ID != category.ID {
		c.Logger().Error("Not Allowed")
		return c.JSON(
			http.StatusForbidden,
			util.CreateBaseResponse[any](false, "Forbidden", nil),
		)
	}

	return c.JSON(
		http.StatusOK,
		util.CreateBaseResponse[response.CommonCategoryResponse](
			true,
			"Category found",
			response.CommonCategoryResponse{
				ID:        category.ID,
				UserID:    category.UserID,
				Name:      category.Name,
				CreatedAt: category.CreatedAt,
				UpdatedAt: category.UpdatedAt,
			},
		),
	)
}

//	@Router		/categories [get]
//	@Summary	Get many categories
//	@Tags		category
//	@Param		itemPerPage	query	string	true	"Amount of items per page"
//	@Param		page		query	string	true	"Page number"
//	@Security	Bearer
//	@Success	200	{object}	util.BaseResponse[[]response.CommonCategoryResponse]
func (ch *categoryHandler) GetMany(c echo.Context) error {
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
	categories, err := ch.cs.GetMany(int(user.ID), itemPerPage, page)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(
			http.StatusInternalServerError,
			util.CreateBaseResponse[any](false, "Internal Server Error", nil),
		)
	}

	responses := make([]response.CommonCategoryResponse, 0, len(categories))
	for _, e := range categories {
		responses = append(responses, response.CommonCategoryResponse{
			ID:        e.ID,
			UserID:    e.UserID,
			Name:      e.Name,
			CreatedAt: e.CreatedAt,
			UpdatedAt: e.UpdatedAt,
		})
	}
	return c.JSON(
		http.StatusOK,
		util.CreateBaseResponse[[]response.CommonCategoryResponse](
			true, "Categories found",
			responses,
		),
	)
}

//	@Router		/categories/{categoryID} [delete]
//	@Summary	Delete one category by ID
//	@Tags		category
//	@Security	Bearer
//	@Success	200	{object}	util.BaseResponse[any]
func (ch *categoryHandler) DeleteOneByID(c echo.Context) error {
	categoryID, err := strconv.Atoi(c.Param("categoryID"))
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(
			http.StatusBadRequest,
			util.CreateBaseResponse[any](false, "Bad Request", nil),
		)
	}

	category, err := ch.cs.GetOneByID(categoryID)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(
			http.StatusInternalServerError,
			util.CreateBaseResponse[any](false, "Internal Server Error", nil),
		)
	}

	user := c.Get("user").(model.User)
	if user.ID != category.ID {
		c.Logger().Error("Not Allowed")
		return c.JSON(
			http.StatusForbidden,
			util.CreateBaseResponse[any](false, "Forbidden", nil),
		)
	}

	if err := ch.cs.DeleteOneByID(categoryID); err != nil {
		c.Logger().Error(err)
		return c.JSON(
			http.StatusInternalServerError,
			util.CreateBaseResponse[any](false, "Internal Server Error", nil),
		)
	}

	return c.JSON(
		http.StatusOK,
		util.CreateBaseResponse[any](
			true, "Category deleted", nil,
		),
	)
}
