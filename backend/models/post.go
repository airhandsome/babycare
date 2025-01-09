package models

import (
	"time"
)

type Post struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UserID    uint      `json:"user_id" gorm:"index"`
	User      User      `json:"user"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Views     int       `json:"views" gorm:"default:0"` // 阅读量
	WordCount int       `json:"word_count"`             // 字数
	ReadTime  int       `json:"read_time"`              // 预计阅读时间（分钟）
	Tags      string    `json:"tags"`                   // 文章标签，用逗号分隔
	Category  string    `json:"category" gorm:"index"`
	Image     string    `json:"image"`
	Icon      string    `json:"icon"`
	Duration  string    `json:"duration"`
	Materials string    `json:"materials"`
	Summary   string    `json:"summary" gorm:"type:text"` // 添加摘要字段
	Avatar    string    `json:"avatar"`
}
