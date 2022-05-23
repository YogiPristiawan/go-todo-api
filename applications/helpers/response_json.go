package helpers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type responseContract struct {
	Message string      `json:"message"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
}

func ResponseJsonHttpOk(c echo.Context, message string, data interface{}) error {
	return c.JSON(http.StatusOK, &responseContract{
		Message: message,
		Status:  "success",
		Data:    data,
	})
}

func ResponseJsonCreated(c echo.Context, message string, data interface{}) error {
	return c.JSON(http.StatusCreated, &responseContract{
		Message: message,
		Status:  "success",
		Data:    data,
	})
}

func ResponseJsonBadRequest(c echo.Context, message string) error {
	return c.JSON(http.StatusBadRequest, &responseContract{
		Message: message,
		Status:  "error",
	})
}

func ResponseJsonUnAuthenticated(c echo.Context, message string) error {
	return c.JSON(http.StatusUnauthorized, &responseContract{
		Message: message,
		Status:  "error",
	})
}

func ResponseJsonForbidden(c echo.Context, message string) error {
	return c.JSON(http.StatusForbidden, &responseContract{
		Message: message,
		Status:  "error",
	})
}

func ResponseJsonNotFoundError(c echo.Context, message string) error {
	return c.JSON(http.StatusNotFound, &responseContract{
		Message: message,
		Status:  "error",
	})
}

func ResponseJsonServerError(c echo.Context) error {
	return c.JSON(http.StatusInternalServerError, &responseContract{
		Message: "Maaf, terjadi kesalahan di server kami",
		Status:  "error",
	})
}

func ResponseJsonMake(c echo.Context, message string, statusCode int, data interface{}) error {
	status := "success"
	if statusCode >= 400 {
		status = "error"
	}
	response := &responseContract{}
	response.Message = message
	response.Status = status
	if data != nil {
		response.Data = data
	}

	return c.JSON(statusCode, response)
}
