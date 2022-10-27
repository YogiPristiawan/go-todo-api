package middlewares

import (
	"go_todo_api/src/shared/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

var decodeAccessToken = jwt.DecodeAccessToken

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")
		if header == "" {
			c.AbortWithStatusJSON(401, gin.H{
				"message": "unautorize",
				"data":    nil,
			})
			return
		}

		var token string
		if strings.HasPrefix(header, "Bearer ") {
			token = strings.Split(header, "Bearer ")[1]
		} else {
			token = header
		}

		claims, err := decodeAccessToken(token)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{
				"message": err.Error(),
				"data":    nil,
			})
			return
		}

		// pass claims into context
		c.Set("auth_user_id", claims.UserId)
		c.Next()
	}
}
