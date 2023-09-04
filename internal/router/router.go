package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ningzining/auth-center/internal/controllers/baseController"
)

func Init(group *gin.Engine) {
	v1Group := group.Group("/v1")
	baseGroup(v1Group)
}
func baseGroup(v1Group *gin.RouterGroup) {
	base := v1Group.Group("/base")
	base.POST("/login", baseController.Login)
}
