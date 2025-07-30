package main

import (
	"log"
	"blog/routes"
	"blog/database"
	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化数据库
	database.InitDatabase()
	
	// 设置Gin路由器
	r := gin.Default()
	
	// 设置路由
	routes.SetupRoutes(r)
	
	// 启动服务器
	log.Println("Starting server on :8080")
	err := r.Run(":8080")
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}