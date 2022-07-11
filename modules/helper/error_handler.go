package helper

import (
	"fmt"

	"github.com/YogiPristiawan/go-todo-api/modules/exceptions"
	"github.com/labstack/echo/v4"
)

func HandleError(c echo.Context, err error) error {
	switch err.(type) {
	case *exceptions.InvariantError:
		he, _ := err.(*exceptions.InvariantError)
		return ResponseJsonBadRequest(c, he.Message)
	case *exceptions.AuthenticationError:
		he, _ := err.(*exceptions.AuthenticationError)
		return ResponseJsonUnAuthenticated(c, he.Message)
	case *exceptions.AuthorizationError:
		he, _ := err.(*exceptions.AuthorizationError)
		return ResponseJsonForbidden(c, he.Message)
	case *exceptions.NotFoundError:
		he, _ := err.(*exceptions.NotFoundError)
		return ResponseJsonNotFoundError(c, he.Message)
	default:
		fmt.Println(err)
		return ResponseJsonServerError(c)
	}
}
