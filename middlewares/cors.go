package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewCorsMiddleware() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.DefaultCORSConfig)
}