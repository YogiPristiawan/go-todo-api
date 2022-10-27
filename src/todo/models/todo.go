package models

// Todo provides data struct of todo
// to interact with database
type Todo struct {
	Id         int64
	UserId     int64
	Todo       string
	Date       string
	IsFinished bool
	CreatedAt  int64
	UpdatedAt  int64
}
