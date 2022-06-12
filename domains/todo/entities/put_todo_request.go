package entities

type PutTodoRequest struct {
	UserId     uint
	Todo       string `json:"todo" validate:"required,max=255"`
	Date       string `json:"date" validate:"required,datetime=2006-01-02 15:04"`
	IsFinished *bool  `json:"is_finished"`
}
