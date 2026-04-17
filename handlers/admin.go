package handlers

import (
	"compass-wealth/middlewares"
	"compass-wealth/services"
	"compass-wealth/views"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AdminHandler struct {
	s services.AdminService
}

func NewAdminHandler(e *echo.Group, s services.AdminService) *AdminHandler {
	h := &AdminHandler{
		s: s,
	}

	e.GET("/users", h.GetUsers)
	e.PUT("/user/role", h.UpdateUserRole, middlewares.AdminOnly)

	return h
}

func (h *AdminHandler) GetUsers(c echo.Context) error {
	users, err := h.s.ListAllUsers()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, views.ListResponse{
		Message: "Users fetched successfully",
		Data:    users,
		Total:   len(users),
	})
}

func (h *AdminHandler) UpdateUserRole(c echo.Context) error {

	var req views.EditUserRole

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}

	if err := h.s.UpdateUserRole(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, "Role updated successfully")
}
