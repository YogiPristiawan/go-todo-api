package dto

import "go_todo_api/src/shared/entities"

// StoreTodoRequest provides data struct
// of todo store action request
type StoreTodoRequest struct {
	entities.RequestMetaData
	Todo       string `json:"todo" validate:"required,max=255"`
	Date       string `json:"date" validate:"required,datetime=2006-01-02"`
	IsFinished bool   `json:"is_finished"`
}

// StoreTodoResponse provides data struct of
// todo store action response
type StoreTodoResponse struct {
	Id         int64  `json:"id"`
	UserId     int64  `json:"user_id"`
	Todo       string `json:"todo"`
	Date       string `json:"date"`
	IsFinished bool   `json:"is_finished"`
	CreatedAt  int64  `json:"created_at"`
	UpdatedAt  int64  `json:"updated_at"`
}

// FindTodoRequest provides required data struct
// of todo find action request
type FindTodoRequest struct {
	entities.RequestMetaData
}

// FindTodoResponse provides data struct of
// todo find action response
type FindTodoResponse struct {
	Id         int64  `json:"id"`
	UserId     int64  `json:"user_id"`
	Todo       string `json:"todo"`
	Date       string `json:"date"`
	IsFinished bool   `json:"is_finished"`
	CreatedAt  int64  `json:"created_at"`
	UpdatedAt  int64  `json:"updated_at"`
}

// DetailTodoRequest provides required data struct
// of todo get detail action request
type DetailTodoRequest struct {
	entities.RequestMetaData
	Id int64 `uri:"id" validate:"required"`
}

// DetailTodoResponse provides data struct of
// todo get detail action response
type DetailTodoResponse struct {
	Id         int64  `json:"id"`
	UserId     int64  `json:"user_id"`
	Todo       string `json:"todo"`
	Date       string `json:"date"`
	IsFinished bool   `json:"is_finished"`
	CreatedAt  int64  `json:"created_at"`
	UpdatedAt  int64  `json:"updated_at"`
}

// UpdateTodoRequest holds the data struct
// needed to update the todo
type UpdateTodoRequest struct {
	entities.RequestMetaData
	Id         int64  `uri:"id" validate:"required"`
	Todo       string `json:"todo" validate:"required,max=255"`
	Date       string `json:"date" validate:"required,datetime=2006-01-02"`
	IsFinished bool   `json:"is_finished"`
}

// UpdateTodoResponse provides data struct
// of update todo response
type UpdateTodoResponse struct {
	Id         int64  `json:"id"`
	Todo       string `json:"user_id"`
	Date       string `json:"date"`
	IsFinished bool   `json:"is_finished"`
	CreatedAt  int64  `json:"created_at"`
	UpdatedAt  int64  `json:"updated_at"`
}
