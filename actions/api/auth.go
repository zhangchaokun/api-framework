package api

import (
	"api-framework/conf/excep"
	"api-framework/core/ginapi"
	"api-framework/service/auth_service"
	"api-framework/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary 获取用户Token
// @Produce  json
// @Param username query string true "userName"
// @Param password query string true "password"
// @Success 200 {object} ginapi.Response
// @Failure 500 {object} ginapi.Response
// @Router /auth [get]
func GetAuth(c *gin.Context) {
	ginApi := ginapi.Gin{C: c}
	//valid := validation.Validation{}

	username := c.Query("username")
	password := c.Query("password")

	/*
		a := auth{Username: username, Password: password}
		ok, _ := valid.Valid(&a)

		if !ok {
			app.MarkErrors(valid.Errors) //logging.Info()
			appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
			return
		}
	*/

	authService := auth_service.Auth{Username: username, Password: password}
	isExist, err := authService.Check()
	if err != nil {
		ginApi.Response(http.StatusOK, excep.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}

	if !isExist {
		ginApi.Response(http.StatusOK, excep.ERROR_AUTH, nil)
		return
	}

	token, err := utils.GenerateToken(username, password)
	if err != nil {
		ginApi.Response(http.StatusOK, excep.ERROR_AUTH_TOKEN, nil)
		return
	}

	ginApi.Response(http.StatusOK, excep.SUCCESS, map[string]string{"token": token,})
}
