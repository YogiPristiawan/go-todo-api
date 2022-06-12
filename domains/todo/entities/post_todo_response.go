package entities

type PostTodoResponse struct {
	Todo       string `json:"todo"`
	Date       string `json:"date"`
	IsFinished *bool  `json:"is_finished"`
	CreatedAt  int64  `json:"created_at"`
	UpdatedAt  int64  `json:"updated_at"`
}

func MapPostTodoResponse(todo *TodoModel) *PostTodoResponse {
	return &PostTodoResponse{
		Todo:       todo.Todo,
		Date:       todo.Date,
		IsFinished: todo.IsFinished,
		CreatedAt:  todo.CreatedAt,
		UpdatedAt:  todo.UpdatedAt,
	}
}
