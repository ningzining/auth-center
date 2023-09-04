package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ningzining/auth-center/internal/global"
	"github.com/ningzining/auth-center/internal/middleware"
	"github.com/ningzining/auth-center/internal/router"
	"github.com/ningzining/auth-center/pkg/log"
)

func main() {
	global.Logger = log.Init()

	engine := gin.Default()
	engine.Use(middleware.Cors())
	// 初始化路由
	router.Init(engine)

	global.Logger.Infof("http server start at 8080")
	err := engine.Run(":8080")
	if err != nil {
		return
	}
}
