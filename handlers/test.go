package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type TestHandler struct{}

func NewTestHandler(e *echo.Group) *TestHandler {
	h := &TestHandler{}

	e.GET("/test", h.Test)

	return h
}

func (h *TestHandler) Test(c echo.Context) error {

	indianCities := []string{"Mumbai", "Delhi", "Bangalore", "Hyderabad", "Chennai", "Kolkata", "Pune", "Ahmedabad", "Surat", "Visakhapatnam"}
	return c.JSON(http.StatusOK, indianCities)
}
