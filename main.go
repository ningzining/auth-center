package main

import (
	"auth-center/internal/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.Use(middleware.Cors())
	err := engine.Run(":8080")
	if err != nil {
		return
	}
}
