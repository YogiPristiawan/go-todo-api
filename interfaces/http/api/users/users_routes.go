package users

import "github.com/labstack/echo/v4"

func Users(e *echo.Echo) {
	usersHandler := &UsersHandler{}
	users := e.Group("/users")

	users.GET("", usersHandler.GetAllUsers)
}
