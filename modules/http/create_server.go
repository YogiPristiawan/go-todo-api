package http

import (
	"github.com/labstack/echo/v4"
)

func CreateServer() *echo.Echo {
	e := echo.New()

	return e
}
