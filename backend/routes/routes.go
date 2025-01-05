package routes

import (
	"github.com/babycare/controllers"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB, rdb *redis.Client) {
	// 创建控制器实例
	postController := controllers.NewPostController(db, rdb)

	// API 版本分组
	v1 := r.Group("/api/v1")
	{
		// 健康检查
		v1.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status": "ok",
			})
		})

		// 文章相关路由
		posts := v1.Group("/posts")
		{
			posts.POST("", postController.CreatePost)
			posts.GET("", postController.GetPosts)
			posts.GET("/:id", postController.GetPost)
			posts.PUT("/:id", postController.UpdatePost)
			posts.DELETE("/:id", postController.DeletePost)
			posts.POST("/:id/views", postController.IncrementViews)
		}
	}
}
