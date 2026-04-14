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

func NewAccountHandler(s services.AccountService) *AccountHandler {
	return &AccountHandler{
		s: s,
	}
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
