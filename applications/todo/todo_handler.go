package todo

import (
	"strconv"
	"strings"

	"github.com/YogiPristiawan/go-todo-api/domain/todo"
	"github.com/YogiPristiawan/go-todo-api/modules/exceptions"
	"github.com/YogiPristiawan/go-todo-api/modules/helper"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type TodoHandler struct {
	UseCase              todo.TodoUseCase
	Validator            *validator.Validate
	ValidatorTranslation ut.Translator
}

func (t *TodoHandler) Store(c echo.Context) error {
	r := new(todo.StoreTodoRequest)
	if err := (&echo.DefaultBinder{}).BindBody(c, r); err != nil {
		return helper.HandleError(c, exceptions.NewInvariantError(err.Error()))
	}

	// validate payload
	if err := t.Validator.Struct(r); err != nil {
		if he, ok := err.(validator.ValidationErrors); ok {
			errors := he.Translate(t.ValidatorTranslation)

			for _, val := range errors {
				return helper.HandleError(c, exceptions.NewInvariantError(val))
			}
		}
	}

	// get authenticated user
	header := c.Request().Header.Get("Authorization")
	token := strings.Split(header, " ")[1]
	claims, err := helper.DecodeAccessToken(token)
	if err != nil {
		return helper.HandleError(c, exceptions.NewAuthenticationError(err.Error()))
	}

	result, err := t.UseCase.Store(&todo.StoreTodoRequest{
		UserId:     claims.UserId,
		Todo:       r.Todo,
		Date:       r.Date,
		IsFinished: r.IsFinished,
	})
	if err != nil {
		return helper.HandleError(c, err)
	}

	return helper.ResponseJsonCreated(c, "success add todo", result)
}

func (t *TodoHandler) GetByUserId(c echo.Context) error {
	// get header
	authorizationHeader := c.Request().Header.Get("Authorization")
	token := strings.Split(authorizationHeader, " ")[1]
	claims, err := helper.DecodeAccessToken(token)
	if err != nil {
		return helper.HandleError(c, exceptions.NewAuthenticationError(err.Error()))
	}
	result, err := t.UseCase.GetByUserId(claims.UserId)
	if err != nil {
		return helper.HandleError(c, err)
	}

	return helper.ResponseJsonHttpOk(c, "success", result)
}

func (t *TodoHandler) UpdateById(c echo.Context) error {
	// collect parameter
	todoId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return helper.HandleError(c, exceptions.NewInvariantError("Paramter must be a valid number"))
	}

	// collect payload
	r := new(todo.UpdateTodoRequest)
	if err := (&echo.DefaultBinder{}).BindBody(c, r); err != nil {
		return helper.HandleError(c, exceptions.NewInvariantError(err.Error()))
	}

	// validdate payload
	if err := t.Validator.Struct(r); err != nil {
		if he, ok := err.(validator.ValidationErrors); ok {
			errors := he.Translate(t.ValidatorTranslation)

			for _, val := range errors {
				return helper.HandleError(c, exceptions.NewInvariantError(val))
			}
		}
	}

	// get authenticated user
	headerAuthorization := c.Request().Header.Get("Authorization")
	token := strings.Split(headerAuthorization, " ")[1]
	claims, err := helper.DecodeAccessToken(token)
	if err != nil {
		return helper.HandleError(c, exceptions.NewAuthenticationError(err.Error()))
	}

	// call use case
	result, err := t.UseCase.UpdateById(claims.UserId, uint(todoId), &todo.UpdateTodoRequest{
		Todo:       r.Todo,
		Date:       r.Date,
		IsFinished: r.IsFinished,
	})
	if err != nil {
		return helper.HandleError(c, err)
	}

	return helper.ResponseJsonHttpOk(c, "oke", result)
}

func (t *TodoHandler) DeleteById(c echo.Context) error {
	// collect parameter
	todoId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return helper.HandleError(c, exceptions.NewInvariantError("Paramter must be a valid number"))
	}

	// get authenticated user
	headerAuthorization := c.Request().Header.Get("Authorization")
	token := strings.Split(headerAuthorization, " ")[1]
	claims, err := helper.DecodeAccessToken(token)
	if err != nil {
		return helper.HandleError(c, exceptions.NewAuthenticationError(err.Error()))
	}

	// call use case
	err = t.UseCase.DeleteById(claims.UserId, uint(todoId))
	if err != nil {
		return helper.HandleError(c, err)
	}

	return helper.ResponseJsonHttpOk(c, "success deleted todo", nil)
}
