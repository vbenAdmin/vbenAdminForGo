package Resful

import "github.com/gin-gonic/gin"

var msgMap = map[int]string{
	1001: "登录信息失效",
	1002: "表单验证失败",
	1004: "请求参数缺失",
}

func Result(code int, msg string, data interface{}) gin.H {
	var resMsg string = msg
	if resMsg == "" {
		resMsg = msgMap[code]
	}
	return gin.H{
		"code":   code,
		"msg":    resMsg,
		"result": data,
	}
}
