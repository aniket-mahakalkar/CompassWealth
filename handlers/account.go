package handlers

import (
	"compass-wealth/services"
	"compass-wealth/utils"
	"compass-wealth/views"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AccountHandler struct {
	s services.AccountService
}

func NewAccountHandler(e *echo.Group, s services.AccountService) *AccountHandler {
	h := &AccountHandler{
		s: s,
	}

	e.POST("/create", h.CreateAccount)
	e.POST("/login", h.Login)

	return h
}

func (h *AccountHandler) CreateAccount(c echo.Context) error {
	var req views.AccountRequest
	if err := c.Bind(&req); err != nil {
		return utils.ThrowError(err, "invalid request")
	}

	if err := h.s.CreateAccount(req); err != nil {
		utils.ThrowError(err, "error in creating account")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, "Account created successfully")
}

func (h *AccountHandler) Login(c echo.Context) error {
	var req views.LoginAccount
	if err := c.Bind(&req); err != nil {
		return err
	}

	res, err := h.s.LoginAccount(req)

	if err != nil {
		utils.ThrowError(err, "error in login")
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	return c.JSON(http.StatusOK, views.Response{
		Message: "Logged in successfully",
		Data:    res,
	})
}
