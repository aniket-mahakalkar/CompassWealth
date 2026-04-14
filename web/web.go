package web

import (
	"compass-wealth/handlers"
	"compass-wealth/services"
	"net/http"
	"os"

	"compass-wealth/utils"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() {
	e := echo.New()

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	authGroup := e.Group("/auth")

	// Protected group
	apiGroup := e.Group("/api/v1")
	apiGroup.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: utils.JWTKey,
	}))

	accountService := *services.NewAccountService()
	handlers.NewAccountHandler(authGroup, accountService)
	handlers.NewTestHandler(apiGroup)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to Compass Wealth")
	})

	port := "8080"
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}
	e.Logger.Fatal(e.Start(":" + port))
}
