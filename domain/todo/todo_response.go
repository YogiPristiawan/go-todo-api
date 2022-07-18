package todo

import (
	"time"
)

type GetTodosResponse struct {
	ID         uint   `json:"id"`
	Todo       string `json:"todo"`
	IsFinished bool   `json:"is_finished"`
	Date       string `json:"date"`
	CreatedAt  int64  `json:"created_at"`
}

func MapGetTodosResponse(todos []*TodoModel) []*GetTodosResponse {
	var todo []*GetTodosResponse

	if len(todos) == 0 {
		return make([]*GetTodosResponse, 0)
	}

	for _, val := range todos {
		todo = append(todo, &GetTodosResponse{
			ID:         val.ID,
			Todo:       val.Todo,
			IsFinished: val.IsFinished,
			Date:       val.Date,
			CreatedAt:  val.CreatedAt,
		})
	}

	return todo
}

type DetailTodoResponse struct {
	ID         uint   `json:"id"`
	Todo       string `json:"todo"`
	IsFinished bool   `json:"is_finished"`
	Date       string `json:"date"`
	CreatedAt  int64  `json:"created_at"`
}

func MapDetailTodoReponse(todo *TodoModel) *DetailTodoResponse {
	return &DetailTodoResponse{
		ID:         todo.ID,
		Todo:       todo.Todo,
		IsFinished: todo.IsFinished,
		Date:       todo.Date,
		CreatedAt:  todo.CreatedAt,
	}
}

type StoreTodoResponse struct {
	Todo       string `json:"todo"`
	Date       string `json:"date"`
	IsFinished bool   `json:"is_finished"`
	CreatedAt  int64  `json:"created_at"`
	UpdatedAt  int64  `json:"updated_at"`
}

func MapStoreTodoResponse(todo *TodoModel) *StoreTodoResponse {
	return &StoreTodoResponse{
		Todo:       todo.Todo,
		Date:       todo.Date,
		IsFinished: todo.IsFinished,
		CreatedAt:  todo.CreatedAt,
		UpdatedAt:  todo.UpdatedAt,
	}
}

type UpdateTodoResponse struct {
	ID         uint   `json:"id"`
	Todo       string `json:"todo"`
	Date       string `json:"date"`
	IsFinished bool   `json:"is_finished"`
	UpdatedAt  string `json:"updated_at"`
}

func MapUpdateTodoResponse(todo *TodoModel) *UpdateTodoResponse {
	return &UpdateTodoResponse{
		ID:         todo.ID,
		Todo:       todo.Todo,
		Date:       todo.Date,
		IsFinished: todo.IsFinished,
		UpdatedAt:  time.Unix(int64(todo.UpdatedAt), 0).UTC().String(),
	}
}
