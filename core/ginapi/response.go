package ginapi

import (
	"api-framework/conf/excep"
	"github.com/gin-gonic/gin"
	"strings"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (g *Gin) Response(httpCode, errCode int, data interface{}, msg string) {
	msg = strings.TrimSpace(msg)
	if msg == "" {
		msg = excep.GetCodeMsg(errCode)
	}
	resp := Response{
		Code: errCode,
		Msg:  msg,
		Data: data,
	}
	g.C.JSON(httpCode, resp)
}