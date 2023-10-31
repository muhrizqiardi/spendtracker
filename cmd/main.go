package main

import (
	"github.com/labstack/echo/v4"
	_ "github.com/muhrizqiardi/spendtracker/docs"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title			Spendtracker API
// @version		1.0
// @description	API for Spendtracker
func main() {
	e := echo.New()

	e.GET("/docs/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
