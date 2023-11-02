package route

import (
	"github.com/labstack/echo/v4"
	"github.com/muhrizqiardi/spendtracker/internal/handler"
	"github.com/muhrizqiardi/spendtracker/internal/middleware"
)

type Router interface {
	Define(e *echo.Echo) *echo.Echo
}

type router struct {
	e         *echo.Echo
	authh     handler.AuthHandler
	authm     middleware.AuthMiddleware
	userh     handler.UserHandler
	accounth  handler.AccountHandler
	categoryh handler.CategoryHandler
	expenseh  handler.ExpenseHandler
}

func NewRouter(
	e *echo.Echo,
	authh handler.AuthHandler,
	authm middleware.AuthMiddleware,
	userh handler.UserHandler,
	accounth handler.AccountHandler,
	categoryh handler.CategoryHandler,
	expenseh handler.ExpenseHandler,
) *router {
	return &router{e, authh, authm, userh, accounth, categoryh, expenseh}
}

func (r *router) Define() *echo.Echo {
	r.e.POST("/auth", r.authh.LogIn)
	r.e.POST("/users", r.userh.Register)

	protected := r.e.Group("/")

	protected.Use(
		r.authm.Authenticate,
	)

	protected.GET("/users/:userID", r.userh.GetOneByID)
	protected.PUT("/users/:userID", r.userh.UpdateOneByID)

	protected.POST("/accounts", r.accounth.Create)
	protected.GET("/accounts", r.accounth.GetMany)
	protected.GET("/accounts/:accountID", r.accounth.GetOneByID)
	protected.PUT("/accounts/:accountID", r.accounth.UpdateOneByID)
	protected.DELETE("/accounts/:accountID", r.accounth.DeleteOneByID)

	protected.POST("/categories", r.categoryh.Create)
	protected.GET("/categories/:categoryID", r.categoryh.GetOneByID)
	protected.GET("/categories", r.categoryh.GetMany)
	protected.DELETE("/categories/:categoryID", r.categoryh.DeleteOneByID)

	protected.POST("/expenses", r.expenseh.Create)
	protected.GET("/expenses/:expenseID", r.expenseh.GetOneByID)
	protected.GET("/expenses", r.expenseh.GetMany)
	protected.PUT("/expenses/:expenseID", r.expenseh.UpdateOneByID)
	protected.DELETE("/expenses/:expenseID", r.expenseh.DeleteOneByID)

	return r.e
}
