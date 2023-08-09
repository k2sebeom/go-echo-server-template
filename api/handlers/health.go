package handlers

import (
	"net/http"

	"github.com/k2sebeom/go-echo-server-template/services"

	"github.com/labstack/echo/v4"
)

type (
	HealthResponse struct {
		Status bool `json:"status"`
	}
)

func GetHealthHandler(healthService *services.HealthService) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, HealthResponse{
			Status: healthService.GetHealthStatus(),
		})
	}
}