package middlewares

import (
	"net/http"
	"strconv"

	"github.com/k2sebeom/go-echo-server-template/config"
	"github.com/k2sebeom/go-echo-server-template/utils"

	"github.com/labstack/echo/v4"
)

func Jwt(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenStr := c.Request().Header.Get("Authorization")

		// Token should be present
		if tokenStr == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Token Missing")
		}

		userId, err := utils.ValidateToken(tokenStr, config.GlobalConfig.Auth.AccessSecret)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Token Invalid")
		}
		c.Request().Header.Add("X-User-Id", strconv.Itoa(int(userId)))
		return next(c)
	}
}
