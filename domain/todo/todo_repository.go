package todo

type TodoRepository interface {
	Store(*TodoModel) (*TodoModel, error)
	GetByUserId(userId uint) ([]*TodoModel, error)
	FindById(todoId uint) (*TodoModel, error)
	UpdateById(todoId uint, todo *TodoModel) (*TodoModel, error)
	DeleteById(todoId uint) error

	VerifyTodoAccess(userId uint, todoId uint) error
}
