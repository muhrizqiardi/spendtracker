package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/muhrizqiardi/spendtracker/internal/database/model"
	"github.com/muhrizqiardi/spendtracker/internal/response"
	"github.com/muhrizqiardi/spendtracker/internal/service"
	"github.com/muhrizqiardi/spendtracker/internal/util"
)

type AdviceHandler interface {
	GetAdvice(c echo.Context) error
}

type adviceHandler struct {
	ads service.AdviceService
}

func NewAdviceHandler(ads service.AdviceService) *adviceHandler {
	return &adviceHandler{ads}
}

func (adh *adviceHandler) GetAdvice(c echo.Context) error {
	user := c.Get("user").(model.User)

	res, err := adh.ads.GetAdvice(int(user.ID))
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(
			http.StatusInternalServerError,
			util.CreateBaseResponse[any](false, "Internal Server Error", nil),
		)
	}

	return c.JSON(
		http.StatusOK,
		util.CreateBaseResponse[response.GetAdviceResponse](
			true, "Getting Advice Success",
			response.GetAdviceResponse{
				Advice: res,
			},
		),
	)
}
