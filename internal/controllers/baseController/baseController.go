package baseController

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/ningzining/auth-center/api/base/v1"
	"github.com/ningzining/auth-center/api/result"
	"github.com/ningzining/auth-center/internal/services/baseService"
)

func Login(ctx *gin.Context) {
	var req v1.LoginReq
	if err := ctx.ShouldBindQuery(&req); err != nil {
		result.Error(ctx, err)
		return
	}
	if err := baseService.Login(ctx, req); err != nil {
		result.Error(ctx, err)
		return
	}
	result.Success(ctx)
}
