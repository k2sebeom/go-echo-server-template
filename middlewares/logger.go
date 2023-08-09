package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Logger() echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${remote_ip} | ${time_custom} | ${method} | ${uri} | ${status} | ${latency_human}\n",
		CustomTimeFormat: "2006-01-02 15:04:05",
	})
}
