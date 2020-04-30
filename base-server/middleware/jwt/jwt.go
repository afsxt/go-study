package jwt

import (
	"base-server/pkg/e"
	"base-server/pkg/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

//-----------------------------------------------------------------------------

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS
		//token := c.Query("token")
		token := c.GetHeader("authorization")
		if token == "" {
			code = e.InvalidParams
		} else {
			_, err := util.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = e.ErrorAuthCheckTokenTimeout
				default:
					code = e.ErrorAuthCheckTokenFail
				}
			}
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
