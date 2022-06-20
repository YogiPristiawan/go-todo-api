package entities

type GetTodoResponse struct {
	ID         uint   `json:"id"`
	Todo       string `json:"todo"`
	IsFinished *bool  `json:"is_finished"`
	Date       string `json:"date"`
	CreatedAt  int64  `json:"created_at"`
}

func MapGetTodoResponse(todos []*TodoModel) []*GetTodoResponse {
	var todoResponse []*GetTodoResponse

	if len(todos) == 0 {
		return make([]*GetTodoResponse, 0)
	}

	for _, val := range todos {
		todoResponse = append(todoResponse, &GetTodoResponse{
			ID:         val.ID,
			Todo:       val.Todo,
			IsFinished: val.IsFinished,
			Date:       val.Date,
			CreatedAt:  val.CreatedAt,
		})
	}

	return todoResponse
}
