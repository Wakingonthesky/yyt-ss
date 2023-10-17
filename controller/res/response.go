package res

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int
	Data any
	Msg  string
}

func Result(code Code, data any, msg string) *gin.H {
	return &gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	}
}

func OK(data any, msg string) *gin.H {
	return Result(OKCode, data, msg)
}

func OKWithData(data any) *gin.H {
	return Result(OKCode, data, "成功")
}

func OKWithCode(code Code) *gin.H {
	return Result(code, map[string]any{}, ErrorMap[code])
}
func OKWithMessage(msg string) *gin.H {
	return Result(OKCode, map[string]any{}, msg)
}

func Fail(data any, msg string) *gin.H {
	return Result(BadRequest, data, msg)
}

func FailWithData(data any) *gin.H {
	return Result(BadRequest, data, "成功")
}

func FailWithMessage(msg string) *gin.H {
	return Result(BadRequest, map[string]any{}, msg)
}
func FailWithCode(code Code) *gin.H {
	msg, ok := ErrorMap[Code(code)]
	if ok {
		return Result(code, map[string]any{}, msg)

	}
	return Result(code, map[string]any{}, "未知错误")
}
