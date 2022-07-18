package todo

import (
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
	// collect payload
	r := new(todo.StoreTodoRequest)
	if err := (&echo.DefaultBinder{}).BindBody(c, r); err != nil {
		return helper.HandleError(c, exceptions.NewInvariantError(err.Error()))
	}

	// validate payload
	if err := t.Validator.Struct(r); err != nil {
		err = helper.ValidatorErrorTranslate(err, t.ValidatorTranslation)
		return helper.HandleError(c, err)
	}

	// get authenticated user
	auth, err := helper.DecodeAuthJwtPayload(c)
	if err != nil {
		return helper.HandleError(c, err)
	}

	// call use case
	result, err := t.UseCase.Store(&todo.StoreTodoRequest{
		UserId:     auth.UserId,
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
	// get authenticated user
	auth, err := helper.DecodeAuthJwtPayload(c)
	if err != nil {
		return helper.HandleError(c, err)
	}

	// call use case
	result, err := t.UseCase.GetByUserId(auth.UserId)
	if err != nil {
		return helper.HandleError(c, err)
	}

	return helper.ResponseJsonHttpOk(c, "success", result)
}

func (t *TodoHandler) DetailById(c echo.Context) error {
	// get authenticated user
	auth, err := helper.DecodeAuthJwtPayload(c)
	if err != nil {
		return helper.HandleError(c, err)
	}

	// collect param
	todoId, err := helper.CollectParamUint(c, "id")
	if err != nil {
		return helper.HandleError(c, err)
	}

	// call use case
	result, err := t.UseCase.DetailById(auth.UserId, uint(todoId))
	if err != nil {
		return helper.HandleError(c, err)
	}

	return helper.ResponseJsonHttpOk(c, "succcess", result)
}

func (t *TodoHandler) UpdateById(c echo.Context) error {
	// collect parameter
	todoId, err := helper.CollectParamUint(c, "id")
	if err != nil {
		return helper.HandleError(c, err)
	}

	// collect paylaod
	body := new(todo.UpdateTodoRequest)
	err = (&echo.DefaultBinder{}).BindBody(c, body)
	if err != nil {
		return helper.HandleError(c, err)
	}

	// validate paylaod
	if err := t.Validator.Struct(body); err != nil {
		err = helper.ValidatorErrorTranslate(err, t.ValidatorTranslation)
		return helper.HandleError(c, err)
	}

	// get authenticated user
	auth, err := helper.DecodeAuthJwtPayload(c)
	if err != nil {
		return helper.HandleError(c, err)
	}

	// call use case
	result, err := t.UseCase.UpdateById(auth.UserId, uint(todoId), &todo.UpdateTodoRequest{
		Todo:       body.Todo,
		Date:       body.Date,
		IsFinished: body.IsFinished,
	})
	if err != nil {
		return helper.HandleError(c, err)
	}

	return helper.ResponseJsonHttpOk(c, "oke", result)
}

func (t *TodoHandler) DeleteById(c echo.Context) error {
	// collect parameter
	todoId, err := helper.CollectParamUint(c, "id")
	if err != nil {
		return helper.HandleError(c, err)
	}

	// get authenticated user
	auth, err := helper.DecodeAuthJwtPayload(c)
	if err != nil {
		return helper.HandleError(c, err)
	}

	// call use case
	err = t.UseCase.DeleteById(auth.UserId, uint(todoId))
	if err != nil {
		return helper.HandleError(c, err)
	}

	return helper.ResponseJsonHttpOk(c, "success deleted todo", nil)
}
