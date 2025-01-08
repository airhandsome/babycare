package main

import (
	"os"

	"github.com/babycare/configs"
	"github.com/babycare/middleware"
	"github.com/babycare/models"
	"github.com/babycare/routes"
	"github.com/babycare/utils/logger"
	"github.com/babycare/utils/mock"
	"github.com/gin-gonic/gin"
)

func main() {
	// 获取环境变量，默认为development
	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "default"
	}

	// 加载配置
	config, err := configs.LoadConfig(env)
	if err != nil {
		panic(err)
	}

	// 初始化日志
	if err := logger.InitLogger(&config.Log); err != nil {
		panic(err)
	}

	// 设置gin模式
	gin.SetMode(config.Server.Mode)

	// 初始化数据库连接
	db, err := configs.InitDatabase(config)
	if err != nil {
		logger.Log.Fatalf("Failed to connect to database: %v", err)
	}

	// 初始化Redis连接
	rdb, err := configs.InitRedis(config)
	if err != nil {
		logger.Log.Fatalf("Failed to connect to Redis: %v", err)
	}

	// 自动迁移数据库表
	if err := db.AutoMigrate(&models.User{}); err != nil {
		logger.Log.Fatalf("Failed to migrate user table: %v", err)
	}
	if err := db.AutoMigrate(&models.Post{}, &models.ForumPost{}); err != nil {
		logger.Log.Fatalf("Failed to migrate post tables: %v", err)
	}
	if err := db.AutoMigrate(&models.Comment{}); err != nil {
		logger.Log.Fatalf("Failed to migrate comment table: %v", err)
	}

	// 生成mock数据
	if env == "default" {
		if err := mock.GenerateMockData(db); err != nil {
			logger.Log.Warnf("Failed to generate mock data: %v", err)
		}
	}

	r := gin.New()

	// 配置gin不自动重定向
	r.RedirectTrailingSlash = false
	r.RedirectFixedPath = false

	r.Use(middleware.CORS())
	r.Use(middleware.Logger())
	r.Use(middleware.Recovery())
	//r.Use(middleware.ClientCheck())

	// 初始化路由
	routes.SetupRoutes(r, db, rdb)

	// 启动服务器
	logger.Log.Infof("Server starting on port %s", config.Server.Port)
	if err := r.Run(":" + config.Server.Port); err != nil {
		logger.Log.Fatalf("Server failed to start: %v", err)
	}
}
