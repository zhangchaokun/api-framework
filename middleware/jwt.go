package middleware

import (
	"api-framework/conf/excep"
	"api-framework/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(context *gin.Context) {
		var (
			code int
			data interface{}
		)

		code = excep.SUCCESS
		token := context.Query("token")

		if token == "" {
			code = excep.INVALID_PARAMS
		} else {
			claims, err := utils.ParseToken(token)
			if err != nil {
				code = excep.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = excep.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != excep.SUCCESS {
			context.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  excep.GetCodeMsg(code),
				"data": data,
			})
			context.Abort()
			return
		}

		context.Next()
	}
}
