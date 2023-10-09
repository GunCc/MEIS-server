package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

// 状态码
var (
	SUCCESS = 200
	FAIL    = 400
)

func Result(code int, data interface{}, message string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  message,
		Data: data,
	})
}

// --------------------------- 200 成功操作 ----------------------------

func Success(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func SuccessWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

func SuccessWithData(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, "操作成功", c)
}

func SuccessWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

// --------------------------- 400 失败操作 ----------------------------

func Fail(c *gin.Context) {
	Result(FAIL, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(FAIL, map[string]interface{}{}, message, c)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(FAIL, data, message, c)
}
