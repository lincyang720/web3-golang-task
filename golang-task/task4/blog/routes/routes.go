package routes

import (
	"blog/controllers"
	"blog/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// 用户相关路由
	userGroup := r.Group("/users")
	{
		userGroup.POST("/register", controllers.Register)
		userGroup.POST("/login", controllers.Login)
	}

	// 文章相关路由
	postGroup := r.Group("/posts")
	{
		postGroup.GET("/", controllers.GetPosts)
		postGroup.GET("/:id", controllers.GetPost)
		
		// 需要认证的路由
		postGroup.Use(middleware.AuthMiddleware())
		{
			postGroup.POST("/", controllers.CreatePost)
			postGroup.PUT("/:id", controllers.UpdatePost)
			postGroup.DELETE("/:id", controllers.DeletePost)
		}
	}

	// 评论相关路由
	commentGroup := r.Group("/comments")
	{
		commentGroup.GET("/post/:post_id", controllers.GetCommentsByPost)
		
		// 需要认证的路由
		commentGroup.Use(middleware.AuthMiddleware())
		{
			commentGroup.POST("/", controllers.CreateComment)
		}
	}
}