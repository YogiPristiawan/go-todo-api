package helpers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ResponseContract struct {
	Message string      `json:"message"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
}

func ResponseJsonHttpOk(c echo.Context, data *ResponseContract) error {
	return c.JSON(http.StatusOK, data)
}
