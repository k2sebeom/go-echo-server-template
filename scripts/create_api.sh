API=$1
API_FILE=$(echo "$API" | tr '[:upper:]' '[:lower:]')


echo "package routes

import (
	\"github.com/k2sebeom/go-echo-server-template/api/handlers\"

	\"github.com/labstack/echo/v4\"
)

func ${API}Routes(router *echo.Group) *echo.Group {
	${API_FILE} := router.Group(\"/${API_FILE}\")
	{
		${API_FILE}.GET(\"/\", handlers.Default${API}Handler())
	}
	return ${API_FILE}
}" > api/routes/$API_FILE.go

echo "package handlers

import (
	\"net/http\"

	\"github.com/labstack/echo/v4\"
)

type (
	${API}Response struct {
		Status bool \`json:\"status\"\`
	}
)

func Default${API}Handler() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, ${API}Response{
			Status: true,
		})
	}
}" > api/handlers/$API_FILE.go