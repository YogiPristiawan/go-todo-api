package presentation

import (
	"go_todo_api/src/shared/entities"

	"github.com/gin-gonic/gin"
)

func ReadRestIn[T interface{}](c *gin.Context, in *T) (err error) {
	err = c.ShouldBind(in)
	return
}

func ReadUriIn[T interface{}](c *gin.Context, in *T) (err error) {
	err = c.ShouldBindUri(c)
	return
}

func WriteRestOut[T interface{}](c *gin.Context, out T, cr entities.CommonResult) {
	if cr.GetCode() == 0 {
		c.JSON(200, out)
		return
	}

	if cr.GetCode() >= 200 && cr.GetCode() < 300 {
		c.JSON(cr.GetCode(), out)
		return
	}

	if cr.GetCode() >= 400 {
		c.AbortWithStatusJSON(cr.GetCode(), entities.BaseResponse[interface{}]{
			Message: cr.GetMessage(),
			Data:    nil,
		})
		return
	}
}
