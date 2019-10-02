package api

import (
	"api-framework/conf/excep"
	"api-framework/core/ginapi"
	"api-framework/helper/utils"
	"api-framework/service/auth_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type auth struct {
	Username string `form:"username" valid:"Required; MaxSize(30)"`
	Password string `form:"password" valid:"Required; MaxSize(50)"`
}

// @Summary 获取用户Token
// @Produce  json
// @Param username query string true "userName"
// @Param password query string true "password"
// @Success 200 {object} ginapi.Response
// @Failure 500 {object} ginapi.Response
// @Router /auth [get]
func GetAuth(c *gin.Context) {
	ginApi := ginapi.Gin{C: c}


	username := c.Query("username")
	password := c.Query("password")


	form := auth{Username: username, Password: password}
	httpCode, errCode, msg := ginapi.BindAndValid(c, &form)
	if errCode != excep.SUCCESS {
		ginApi.Response(httpCode, errCode, nil,msg)
		return
	}


	authService := auth_service.Auth{Username: username, Password: password}
	isExist, err := authService.Check()
	if err != nil {
		ginApi.Response(http.StatusOK, excep.ERROR_AUTH_CHECK_TOKEN_FAIL, nil, err.Error())
		return
	}

	if !isExist {
		ginApi.Response(http.StatusOK, excep.ERROR_AUTH, nil, "")
		return
	}

	token, err := utils.GenerateToken(username, password)
	if err != nil {
		ginApi.Response(http.StatusOK, excep.ERROR_AUTH_TOKEN, nil, "")
		return
	}

	ginApi.Response(http.StatusOK, excep.SUCCESS, map[string]string{"token": token,},"")
}
