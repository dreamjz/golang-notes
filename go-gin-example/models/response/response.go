package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type EmptyData map[string]interface{}

func Result(code int, msg string, data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func OK(c *gin.Context) {
	Result(Success, GetMsg(Success), EmptyData{}, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(Success, GetMsg(Success), data, c)
}

func OkWithDetails(data interface{}, msg string, c *gin.Context) {
	Result(Success, msg, data, c)
}

func Fail(c *gin.Context) {
	Result(Error, GetMsg(Error), EmptyData{}, c)
}

func FailWithCode(code int, c *gin.Context) {
	Result(code, GetMsg(code), EmptyData{}, c)
}

func FailWithMessage(msg string, c *gin.Context) {
	Result(Error, msg, EmptyData{}, c)
}

func FailWithDetails(msg string, data interface{}, c *gin.Context) {
	Result(Error, msg, data, c)
}
