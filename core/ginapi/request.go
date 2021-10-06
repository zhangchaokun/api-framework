package ginapi

import (
	"api-framework/conf/excep"
	"api-framework/core/logging"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
)

func BindAndValid(c *gin.Context, form interface{}) (int, int,string) {
	err := c.Bind(form)
	if err != nil {
		fmt.Println(err)
		return http.StatusBadRequest, excep.INVALID_PARAMS, err.Error()
	}

	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		return http.StatusInternalServerError, excep.ERROR, err.Error()
	}
	if !check {
		validErrors(valid.Errors)
		return http.StatusBadRequest, excep.INVALID_PARAMS,valid.Errors[0].Message
	}

	return http.StatusOK, excep.SUCCESS, excep.GetCodeMsg(excep.SUCCESS)
}


func validErrors(errors []*validation.Error) {
	for _, err := range errors {
		logging.Info(err.Key, err.Message)
	}
}