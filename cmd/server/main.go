package main

import (
	config "Demo1/internal/db_config"
	"Demo1/internal/handlers"
	"Demo1/internal/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化数据库连接
	config.InitDB()

	r := gin.Default()

	// 用户注册和登录路由
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	// JWT 认证保护的路由组
	auth := r.Group("/questions")
	auth.Use(middleware.JWTAuthMiddleware())
	auth.POST("/", handlers.CreateQuestion)
	auth.PUT("/:id", handlers.UpdateQuestion)
	auth.DELETE("/:id", handlers.DeleteQuestion)
	auth.POST("/:id/answers", handlers.AddAnswer)

	r.Run(":8080") // 启动服务器，监听8080端口
}
