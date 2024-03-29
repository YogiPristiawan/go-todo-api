package services

import (
	"go_todo_api/src/shared/databases"
	"go_todo_api/src/todo/dto"
	"go_todo_api/src/todo/models"
)

// mockError is a struct to mock an error, for testing purpose
type mockError struct {
	code int
}

func (m *mockError) Error() string {
	return ""
}

// this variables store helper functions
// to easier testing
var wrapDBErr = databases.WrapError

// mapStoreToResponse handle struct mapping from
// Todo model into TodoResponse
func mapStoreToResponse(res *dto.StoreTodoResponse, todo *models.Todo) {
	res.Id = todo.Id
	res.UserId = todo.UserId
	res.Todo = todo.Todo
	res.Date = todo.Date
	res.IsFinished = todo.IsFinished
	res.CreatedAt = todo.CreatedAt
	res.UpdatedAt = todo.UpdatedAt
}

// mapFindToResponse handle struct mapping from
// Todo model into FindTodoResponse
func mapFindToResponse(res *[]dto.FindTodoResponse, todos []models.Todo) {
	for _, todo := range todos {
		var r dto.FindTodoResponse
		r.Id = todo.Id
		r.UserId = todo.UserId
		r.Todo = todo.Todo
		r.Date = todo.Date
		r.IsFinished = todo.IsFinished
		r.CreatedAt = todo.CreatedAt
		r.UpdatedAt = todo.UpdatedAt

		*res = append(*res, r)
	}
}

// mapDetailTodResponse handle struct mapping from
// Todo model into DetailTodoResponse
func mapDetailToResponse(res *dto.DetailTodoResponse, todo models.Todo) {
	res.Id = todo.Id
	res.UserId = todo.UserId
	res.Todo = todo.Todo
	res.Date = todo.Date
	res.IsFinished = todo.IsFinished
	res.UpdatedAt = todo.UpdatedAt
	res.CreatedAt = todo.CreatedAt
}

// mapUpdateTodoResponse handle struct mapping from
// TodoModel into UpdateTodoResponse
func mapUpdateTodoResponse(res *dto.UpdateTodoResponse, todo models.Todo) {
	res.Id = todo.Id
	res.Todo = todo.Todo
	res.Date = todo.Date
	res.IsFinished = todo.IsFinished
	res.CreatedAt = todo.CreatedAt
	res.UpdatedAt = todo.UpdatedAt
}
