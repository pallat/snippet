package payment

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func New() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "OK")
	}
}
