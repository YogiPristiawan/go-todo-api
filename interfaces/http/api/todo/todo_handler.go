package todo

import (
	"strconv"
	"strings"

	"github.com/YogiPristiawan/go-todo-api/applications/exceptions"
	"github.com/YogiPristiawan/go-todo-api/applications/helpers"
	"github.com/YogiPristiawan/go-todo-api/domains/todo"
	"github.com/YogiPristiawan/go-todo-api/domains/todo/entities"
	"github.com/YogiPristiawan/go-todo-api/infrastructures/security/tokenize"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type todoHandler struct {
	useCase              todo.TodoUseCase
	validator            *validator.Validate
	validatorTranslation ut.Translator
	jwt                  *tokenize.JwtToken
}

func (t *todoHandler) storeTodo(c echo.Context) error {
	r := new(entities.PostTodoRequest)
	if err := (&echo.DefaultBinder{}).BindBody(c, r); err != nil {
		return helpers.HandleError(c, exceptions.NewInvariantError(err.Error()))
	}

	// validate payload
	if err := t.validator.Struct(r); err != nil {
		if he, ok := err.(validator.ValidationErrors); ok {
			errors := he.Translate(t.validatorTranslation)

			for _, val := range errors {
				return helpers.HandleError(c, exceptions.NewInvariantError(val))
			}
		}
	}

	// get authenticated user
	header := c.Request().Header.Get("Authorization")
	token := strings.Split(header, " ")[1]
	claims, err := t.jwt.DecodeAccessToken(token)
	if err != nil {
		return helpers.HandleError(c, exceptions.NewAuthenticationError(err.Error()))
	}

	result, err := t.useCase.StoreTodo(&entities.PostTodoRequest{
		UserId:     claims.UserId,
		Todo:       r.Todo,
		Date:       r.Date,
		IsFinished: r.IsFinished,
	})
	if err != nil {
		return helpers.HandleError(c, err)
	}

	return helpers.ResponseJsonCreated(c, "success add todo", result)
}

func (t *todoHandler) getTodos(c echo.Context) error {
	// get header
	authorizationHeader := c.Request().Header.Get("Authorization")
	token := strings.Split(authorizationHeader, " ")[1]
	claims, err := t.jwt.DecodeAccessToken(token)
	if err != nil {
		return helpers.HandleError(c, exceptions.NewAuthenticationError(err.Error()))
	}
	result, err := t.useCase.GetTodos(claims.UserId)
	if err != nil {
		return helpers.HandleError(c, err)
	}

	return helpers.ResponseJsonHttpOk(c, "success", result)
}

func (t *todoHandler) updateTodo(c echo.Context) error {
	// collect parameter
	todoId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return helpers.HandleError(c, exceptions.NewInvariantError("Paramter must be a valid number"))
	}

	// collect payload
	r := new(entities.PutTodoRequest)
	if err := (&echo.DefaultBinder{}).BindBody(c, r); err != nil {
		return helpers.HandleError(c, exceptions.NewInvariantError(err.Error()))
	}

	// validdate payload
	if err := t.validator.Struct(r); err != nil {
		if he, ok := err.(validator.ValidationErrors); ok {
			errors := he.Translate(t.validatorTranslation)

			for _, val := range errors {
				return helpers.HandleError(c, exceptions.NewInvariantError(val))
			}
		}
	}

	// get authenticated user
	headerAuthorization := c.Request().Header.Get("Authorization")
	token := strings.Split(headerAuthorization, " ")[1]
	claims, err := t.jwt.DecodeAccessToken(token)
	if err != nil {
		return helpers.HandleError(c, exceptions.NewAuthenticationError(err.Error()))
	}

	// call use case
	result, err := t.useCase.UpdateTodo(claims.UserId, uint(todoId), &entities.PutTodoRequest{
		Todo:       r.Todo,
		Date:       r.Date,
		IsFinished: r.IsFinished,
	})
	if err != nil {
		return helpers.HandleError(c, err)
	}

	return helpers.ResponseJsonHttpOk(c, "oke", result)
}

func (t *todoHandler) deleteTodo(c echo.Context) error {
	// collect parameter
	todoId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return helpers.HandleError(c, exceptions.NewInvariantError("Paramter must be a valid number"))
	}

	// get authenticated user
	headerAuthorization := c.Request().Header.Get("Authorization")
	token := strings.Split(headerAuthorization, " ")[1]
	claims, err := t.jwt.DecodeAccessToken(token)
	if err != nil {
		return helpers.HandleError(c, exceptions.NewAuthenticationError(err.Error()))
	}

	// call use case
	err = t.useCase.DeleteTodo(claims.UserId, uint(todoId))
	if err != nil {
		return helpers.HandleError(c, err)
	}

	return helpers.ResponseJsonHttpOk(c, "success deleted todo", nil)
}
