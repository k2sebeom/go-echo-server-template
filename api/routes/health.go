package routes

import (
	"github.com/k2sebeom/go-echo-server-template/api/handlers"
	"github.com/k2sebeom/go-echo-server-template/middlewares"
	"github.com/k2sebeom/go-echo-server-template/services"

	"github.com/labstack/echo/v4"
)

func HealthRoutes(router *echo.Group, healthService *services.HealthService) *echo.Group {
	health := router.Group("/health", middlewares.Jwt)
	{
		health.GET("/", handlers.GetHealthHandler(healthService))
	}
	return health
}