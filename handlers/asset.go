package handlers

import (
	"compass-wealth/middlewares"
	"compass-wealth/services"
	"compass-wealth/views"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AssetHandler struct {
	s services.AssetService
}

func NewAssetHandler(e *echo.Group, s services.AssetService) *AssetHandler {
	h := &AssetHandler{
		s: s,
	}

	e.GET("/assets", h.GetAssets)
	e.POST("/asset", h.CreateAsset)
	e.PUT("/asset", h.UpdateAsset)
	// e.DELETE("/assets", h.DeleteAsset)

	return h
}

func (h *AssetHandler) GetAssets(c echo.Context) error {
	authUser, err := middlewares.GetAuthUser(c)
	if err != nil {
		return err
	}

	assets, err := h.s.GetAssets(authUser.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, views.ListResponse{
		Message: "Assets fetched successfully",
		Data:    assets,
		Total:   len(assets),
	})
}

func (h *AssetHandler) CreateAsset(c echo.Context) error {
	var req views.AssetRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}
	user, err := middlewares.GetAuthUser(c)

	if err != nil {
		return err
	}

	if err := h.s.CreateAsset(*user, req); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, views.Response{
		Message: "created successfully",
	})
}

func (h *AssetHandler) UpdateAsset(c echo.Context) error {
	var req views.AssetRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}

	user, err := middlewares.GetAuthUser(c)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid user")
	}

	if err := h.s.UpdateAsset(*user, req); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, views.Response{
		Message: "updated successfully",
	})
}
