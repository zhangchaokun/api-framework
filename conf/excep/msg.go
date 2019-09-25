package excep

var MsgFlags = map[int]string {
	SUCCESS:                         "处理成功",
	ERROR:                           "未知错误",
	INVALID_PARAMS:                  "请求参数错误",


	//10001


	//20001
	ERROR_AUTH_CHECK_TOKEN_FAIL:     "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT:  "Token已超时",
	ERROR_AUTH_TOKEN:                "Token生成失败",
	ERROR_AUTH:                      "Token错误",
}

func GetCodeMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
