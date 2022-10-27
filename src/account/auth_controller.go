package account

import (
	"go_todo_api/src/account/dto"
	"go_todo_api/src/account/services"
	"go_todo_api/src/shared/entities"
	"go_todo_api/src/shared/presentation"

	"github.com/gin-gonic/gin"
)

// AuthController is an abstract that contains
// methods to handle auth related request
type AuthController interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
}

// authController is a struct that has methods
// to handle auth related request
type authController struct {
	service services.AuthService
}

// NewAuthController creates an instance of authContoller
func NewAuthController(
	service services.AuthService,
) AuthController {
	return &authController{
		service: service,
	}
}

// Login collect the login request data
// and call the service
func (a *authController) Login(c *gin.Context) {
	in := dto.LoginRequest{}
	if err := presentation.ReadRestIn(c, &in); err != nil {
		o := entities.CommonResult{}
		o.SetResponse(400, err)
		presentation.WriteRestOut(c, o, o)
		return
	}

	out := a.service.Login(in)

	presentation.WriteRestOut(c, out, out.CommonResult)
}

// Register collect the register request data
// and call the service
func (a *authController) Register(c *gin.Context) {
	in := dto.RegisterRequest{}
	if err := presentation.ReadRestIn(c, &in); err != nil {
		o := entities.CommonResult{}
		o.SetResponse(400, err)
		presentation.WriteRestOut(c, o, o)
		return
	}

	out := a.service.Register(in)
	presentation.WriteRestOut(c, out, out.CommonResult)
}
