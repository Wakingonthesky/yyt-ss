package res

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int
	Data any
	Msg  string
}

func Result(code Code, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: int(code),
		Data: data,
		Msg:  msg,
	})
}

func OK(data any, msg string, c *gin.Context) {
	Result(OKCode, data, msg, c)
}

func OKWithData(data any, c *gin.Context) {
	Result(OKCode, data, "成功", c)
}

func OKWithMessage(msg string, c *gin.Context) {
	Result(OKCode, map[string]any{}, msg, c)
}

func Fail(data any, msg string, c *gin.Context) {
	Result(BadRequest, data, msg, c)
}

func FailWithData(data any, c *gin.Context) {
	Result(BadRequest, data, "成功", c)
}

func FailWithMessage(msg string, c *gin.Context) {
	Result(BadRequest, map[string]any{}, msg, c)
}
func FailWithCode(code Code, c *gin.Context) {
	msg, ok := ErrorMap[Code(code)]
	if ok {
		Result(code, map[string]any{}, msg, c)
		return
	}
	Result(code, map[string]any{}, "未知错误", c)
}
