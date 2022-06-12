package entities

import "time"

type PutTodoResponse struct {
	ID         uint   `json:"id"`
	Todo       string `json:"todo"`
	Date       string `json:"date"`
	IsFinished *bool  `json:"is_finished"`
	UpdatedAt  string `json:"updated_at"`
}

func MapPutTodoResponse(todo *TodoModel) *PutTodoResponse {
	return &PutTodoResponse{
		ID:         todo.ID,
		Todo:       todo.Todo,
		Date:       todo.Date,
		IsFinished: todo.IsFinished,
		UpdatedAt:  time.Unix(todo.UpdatedAt, 0).UTC().String(),
	}
}
