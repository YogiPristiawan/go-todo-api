package todo

import (
	"go_todo_api/src/shared/entities"
	"go_todo_api/src/shared/presentation"
	"go_todo_api/src/todo/dto"
	"go_todo_api/src/todo/services"

	"github.com/gin-gonic/gin"
)

// TodoController is an abstract that contains
// mehods to handle todo related request
type TodoController interface {
	Store(c *gin.Context)
	Find(c *gin.Context)
	Detail(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

// todoController is a struct that has methods
// to handle related todo request
type todoController struct {
	service services.TodoService
}

// NewTodoController creates an instance of todoController
func NewTodoController(service services.TodoService) TodoController {
	return &todoController{
		service: service,
	}
}

// Store handle a request to store todo data
func (t *todoController) Store(c *gin.Context) {
	// get authenticated user
	authUserId, _ := c.Get("auth_user_id")

	// make payload
	in := dto.StoreTodoRequest{
		RequestMetaData: entities.RequestMetaData{
			AuthUserId: authUserId.(int64),
		},
	}

	// bind request data
	if err := presentation.ReadRestIn(c, &in); err != nil {
		o := entities.CommonResult{}
		presentation.WriteRestOut(c, o, o)
		return
	}

	out := t.service.Store(in)
	presentation.WriteRestOut(c, out, out.CommonResult)
}

// Find handle a request to find authenticated user todo data
func (t *todoController) Find(c *gin.Context) {
	// get authenticated user
	authUserId, _ := c.Get("auth_user_id")

	// make payload
	in := dto.FindTodoRequest{
		RequestMetaData: entities.RequestMetaData{
			AuthUserId: authUserId.(int64),
		},
	}

	out := t.service.Find(in)
	presentation.WriteRestOut(c, out, out.CommonResult)
}

// Detail handle a request to get detail of user todo
func (t *todoController) Detail(c *gin.Context) {
	// get authenticated user
	authUserId, _ := c.Get("auth_user_id")

	// make payload
	in := dto.DetailTodoRequest{
		RequestMetaData: entities.RequestMetaData{
			AuthUserId: authUserId.(int64),
		},
	}

	// bind request data
	if err := presentation.ReadUriIn(c, &in); err != nil {
		o := entities.CommonResult{}
		presentation.WriteRestOut(c, o, o)
		return
	}

}

// Update handle a request to update user todo data
func (t *todoController) Update(c *gin.Context) {

}

// Delete handle a request to delete user todo data
func (t *todoController) Delete(c *gin.Context) {

}
