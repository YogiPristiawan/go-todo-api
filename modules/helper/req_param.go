package helper

import (
	"fmt"
	"strconv"

	"go_todo_api/modules/exceptions"

	"github.com/labstack/echo/v4"
)

func CollectParamUint(c echo.Context, key string) (uint64, error) {
	param, err := strconv.ParseUint(c.Param(key), 10, 64)
	if err != nil {
		return 0, exceptions.NewInvariantError(fmt.Sprintf("parameter %s must be an number", key))
	}

	return param, nil
}
