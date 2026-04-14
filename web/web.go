package web

import (
	"compass-wealth/handlers"
	"compass-wealth/services"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() {
	e := echo.New()

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	authGroup := e.Group("/auth")

	accountService := *services.NewAccountService()
	handlers.NewAccountHandler(authGroup, accountService)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	port := "8080"
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}
	e.Logger.Fatal(e.Start(":" + port))
}
