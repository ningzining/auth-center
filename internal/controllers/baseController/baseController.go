package baseController

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/ningzining/auth-center/api/base/v1"
	"github.com/ningzining/auth-center/api/result"
	"github.com/ningzining/auth-center/internal/services/baseService"
)

// Login
//
//	@Summary	登录
//	@Tags		base
//	@Accept		json
//	@Produce	json
//	@Param		LoginReq	body		v1.LoginReq	true	"登录请求参数"
//	@Success	200			{object}	result.Result(data=null)
//	@Router		/base/login [post]
func Login(ctx *gin.Context) {
	var req v1.LoginReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		result.Error(ctx, err)
		return
	}
	if err := baseService.Login(ctx, req); err != nil {
		result.Error(ctx, err)
		return
	}
	result.Success(ctx)
}
