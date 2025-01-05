package mock

import (
	"time"

	"github.com/babycare/models"
	"gorm.io/gorm"
)

func GenerateMockData(db *gorm.DB) error {
	// 创建测试用户
	users := []models.User{
		{
			Username:  "test_user1",
			Email:    "user1@example.com",
			Password: "password123",
		},
		{
			Username:  "test_user2",
			Email:    "user2@example.com",
			Password: "password123",
		},
	}

	for i := range users {
		users[i].CreatedAt = time.Now()
		users[i].UpdatedAt = time.Now()
		if err := db.Create(&users[i]).Error; err != nil {
			return err
		}
	}

	// 创建测试文章
	posts := []models.Post{
		{
			Title:   "婴儿护理小技巧",
			Content: "1. 保持室温适中\n2. 定时换尿布\n3. 注意卫生",
			UserID:  users[0].ID,
		},
		{
			Title:   "宝宝营养食谱",
			Content: "6-12个月宝宝辅食添加指南...",
			UserID:  users[0].ID,
		},
		{
			Title:   "新生儿睡眠指南",
			Content: "如何帮助宝宝建立良好的作息习惯...",
			UserID:  users[1].ID,
		},
	}

	for i := range posts {
		posts[i].CreatedAt = time.Now()
		posts[i].UpdatedAt = time.Now()
		if err := db.Create(&posts[i]).Error; err != nil {
			return err
		}
	}

	return nil
} 