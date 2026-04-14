package handlers

import (
	"compass-wealth/services"
	"compass-wealth/views"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AccountHandler struct {
	s services.AccountService
}

func NewAccountHandler(e *echo.Group,s services.AccountService) *AccountHandler {
	h:= &AccountHandler{
		s: s,
	}

	e.POST("/account", h.CreateAccount)
	e.POST("/account/login", h.Login)

	return h
}

func (h *AccountHandler) CreateAccount(c echo.Context) error {
	var req views.AccountRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	if err := h.s.CreateAccount(req); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, "Account created successfully")
}


func (h *AccountHandler) Login(c echo.Context) error {
	var req views.LoginAccount
	if err := c.Bind(&req); err != nil {
		return err
	}

	res, err := h.s.LoginAccount(req); 

	if err != nil {
		return  err
	}

	return c.JSON(http.StatusOK, views.Response{
		Message: "Logged in successfully",
		Data:    res,
	})
}	