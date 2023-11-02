package main

import (
	"log"

	"github.com/labstack/echo/v4"
	_ "github.com/muhrizqiardi/spendtracker/docs"
	"github.com/muhrizqiardi/spendtracker/internal/database/setup"
	"github.com/muhrizqiardi/spendtracker/internal/handler"
	"github.com/muhrizqiardi/spendtracker/internal/middleware"
	"github.com/muhrizqiardi/spendtracker/internal/repository"
	"github.com/muhrizqiardi/spendtracker/internal/route"
	"github.com/muhrizqiardi/spendtracker/internal/service"
	"github.com/muhrizqiardi/spendtracker/internal/util"
	"github.com/sashabaranov/go-openai"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.uber.org/zap"
)

// @title			Spendtracker API
// @version		1.0
// @description	API for Spendtracker
func main() {
	zl, err := zap.NewProduction()
	if err != nil {
		log.Fatal()
	}
	defer zl.Sync()
	lg := util.NewLogger(zl)

	cfg := util.LoadConfig()
	db, err := setup.SetupMigrateAndSeedMySQL(cfg, lg)

	oac := openai.NewClient(cfg.OpenAIAPIKey)

	userRepo := repository.NewUserRepository(db)
	accountRepo := repository.NewAccountRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)
	expenseRepo := repository.NewExpenseRepository(db)
	openaiRepo := repository.NewOpenAIRepository(oac)

	userService := service.NewUserService(userRepo)
	authService := service.NewAuthService(userService, cfg.Secret)
	accountService := service.NewAccountService(accountRepo)
	categoryService := service.NewCategoryService(categoryRepo)
	expenseService := service.NewExpenseService(expenseRepo, accountService)
	adviceService := service.NewAdviceService(expenseService, openaiRepo)

	authHandler := handler.NewAuthHandler(authService)
	userHandler := handler.NewUserHandler(userService)
	accountHandler := handler.NewAccountHandler(accountService)
	categoryHandler := handler.NewCategoryHandler(categoryService)
	expenseHandler := handler.NewExpenseHandler(expenseService)
	adviceHandler := handler.NewAdviceHandler(adviceService)

	authMiddleware := middleware.NewAuthMiddleware(userService, cfg.Secret)

	e := echo.New()

	r := route.NewRouter(
		e,
		authHandler,
		authMiddleware,
		userHandler,
		accountHandler,
		categoryHandler,
		expenseHandler,
		adviceHandler,
	).Define()

	r.GET("/docs/*", echoSwagger.WrapHandler)

	lg.FatalError("Server failed", r.Start(":"+cfg.Port))
}
