package account

import (
	"go_todo_api/src/account/dto"
	"go_todo_api/src/account/services"
	"go_todo_api/src/shared/entities"
	"go_todo_api/src/shared/presentation"

	"github.com/gin-gonic/gin"
)

// AccountController is an abstract that contains
// methods to handle profile related request
type AccountController interface {
	GetProfile(c *gin.Context)
}

// profileController is a struct that has methods
// to handle profile related request
type accountController struct {
	service services.AccountService
}

// NewAccountController creates an instance of accountController
func NewAccountController(service services.AccountService) AccountController {
	return &accountController{
		service: service,
	}
}

// GetProfile handle request to get authenticated
// user profile using JWT token
func (a *accountController) GetProfile(c *gin.Context) {
	authUserId, _ := c.Get("auth_user_id")
	in := dto.ProfileRequest{
		RequestMetaData: entities.RequestMetaData{
			AuthUserId: authUserId.(int64),
		},
	}

	out := a.service.GetProfile(in)
	presentation.WriteRestOut(c, out, out.CommonResult)
}
