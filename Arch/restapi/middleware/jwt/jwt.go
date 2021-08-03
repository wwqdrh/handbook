package jwt

import (
	"archbook/restapi/pkg/cons"
	"archbook/restapi/pkg/utils"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			code int
			data interface{}
		)

		code = cons.SUCCESS
		token := c.Query("token")
		if token == "" {
			code = cons.INVALID_PARAMS
		} else {
			_, err := utils.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = cons.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				default:
					code = cons.ERROR_AUTH_CHECK_TOKEN_FAIL

				}
			}
		}

		if code != cons.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  cons.GetMsg(code),
				"data": data,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
