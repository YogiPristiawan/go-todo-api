package services

import (
	"fmt"
	"go_todo_api/src/shared/entities"
	"go_todo_api/src/todo/dto"
	"go_todo_api/src/todo/models"
	"go_todo_api/src/todo/repositories"
	"go_todo_api/src/todo/validators"
)

// TodoService is an abstract that contains
// methods to handle todo business logic
type TodoService interface {
	Store(in dto.StoreTodoRequest) (out entities.BaseResponse[dto.StoreTodoResponse])
	Find(in dto.FindTodoRequest) (out entities.BaseResponseArray[dto.FindTodoResponse])
	Detail(in dto.DetailTodoRequest) (out entities.BaseResponse[dto.DetailTodoResponse])
}

// todoService is a struct that has methods
// to handle todo bisiness logic
type todoService struct {
	validator validators.TodoValidator
	todoRepo  repositories.TodoRepository
}

// NewTodoService creates an instance of todoService
func NewTodoService(
	validator validators.TodoValidator,
	todoRepo repositories.TodoRepository,
) TodoService {
	return &todoService{
		validator: validator,
		todoRepo:  todoRepo,
	}
}

// Store handle business logic action to store todo
func (t *todoService) Store(in dto.StoreTodoRequest) (out entities.BaseResponse[dto.StoreTodoResponse]) {
	// validate request
	if err := t.validator.ValidateStore(in); err != nil {
		out.SetResponse(400, err)
		return
	}

	// create todo data to store into database
	todo := models.Todo{
		UserId:     in.RequestMetaData.UserId,
		Todo:       in.Todo,
		Date:       in.Date,
		IsFinished: in.IsFinished,
	}
	err := t.todoRepo.Store(&todo)
	switch wrapDBErr(err) {
	case 500:
		out.SetResponse(500, err)
		return
	}

	mapStoreToResponse(&out.Data, &todo)

	out.Message = "todo created"
	out.SetResponse(201, nil)
	return
}

// Find handle business logic action to find user todo datas
func (t *todoService) Find(in dto.FindTodoRequest) (out entities.BaseResponseArray[dto.FindTodoResponse]) {
	// find todo data
	todos, err := t.todoRepo.Find(in.UserId)
	switch wrapDBErr(err) {
	case 404:
		out.SetResponse(404, fmt.Errorf("todo not found"))
		return
	case 500:
		out.SetResponse(500, err)
		return
	}

	mapFindToResponse(&out.Data, todos)

	out.Message = "list of todos"
	return
}

// Detail handle business logic action to ge detail of user todo data
func (t *todoService) Detail(in dto.DetailTodoRequest) (out entities.BaseResponse[dto.DetailTodoResponse]) {
	// validate request
	if err := t.validator.ValidateDetail(in); err != nil {
		out.SetResponse(400, err)
		return
	}

	// get todo detail
	todo, err := t.todoRepo.Detail(in.UserId, in.Id)
	switch wrapDBErr(err) {
	case 404:
		out.SetResponse(404, fmt.Errorf("todo data not found"))
		return
	case 500:
		out.SetResponse(500, err)
		return
	}

	mapDetailToResponse(&out.Data, todo)

	out.Message = "detail of todo"
	return
}