package result

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	success         = "success"
	ParamParseError = "参数解析错误"
)

type Result struct {
	Code int         `json:"code"` // 返回码
	Msg  string      `json:"msg"`  // 返回信息
	Data interface{} `json:"data"` // 返回数据
}

func Success(ctx *gin.Context) {
	SuccessWithData(ctx, success, nil)
}

func SuccessWithMsg(ctx *gin.Context, msg string) {
	SuccessWithData(ctx, msg, nil)
}

func SuccessWithData(ctx *gin.Context, msg string, data interface{}) {
	ctx.JSON(http.StatusOK, Result{
		Code: 1,
		Msg:  msg,
		Data: data,
	})
}

func Error(ctx *gin.Context, err error) {
	ErrorWithData(ctx, err, nil)
}

func ErrorWithData(ctx *gin.Context, err error, data interface{}) {
	ctx.JSON(http.StatusOK, Result{
		Code: 500,
		Msg:  err.Error(),
		Data: data,
	})
}
