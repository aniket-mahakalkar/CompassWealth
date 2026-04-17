package handlers

import (
	"compass-wealth/middlewares"
	"compass-wealth/services"
	"compass-wealth/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type DashboardHandler struct {
	s services.DashboardService
}

func NewDashboardHandler(e *echo.Group, s services.DashboardService) *DashboardHandler {
	h := &DashboardHandler{
		s: s,
	}

	e.GET("/all/stats", h.GetAllStats)

	return h
}

func (h *DashboardHandler) GetAllStats(c echo.Context) error {
	user, err := middlewares.GetAuthUser(c)
	if err != nil {
		return err
	}

	res, err := h.s.GetDashboardData(*user)
	if err != nil {
		utils.ThrowError(err, "error fetching dashboard")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}
