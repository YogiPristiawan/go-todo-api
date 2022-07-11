package todo

type TodoUseCase interface {
	Store(payload *StoreTodoRequest) (*StoreTodoResponse, error)
	GetByUserId(userId uint) ([]*GetTodosResponse, error)
	UpdateById(authUserid uint, todoId uint, payload *UpdateTodoRequest) (*UpdateTodoResponse, error)
	DeleteById(authUserId uint, todoId uint) error
}
