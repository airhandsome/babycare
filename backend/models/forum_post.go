package models

import "time"

type ForumPost struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	User      User      `json:"user" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Category  string    `json:"category"` // 帖子分类
	Content   string    `json:"content"`  // 帖子内容
	Images    string    `json:"images"`   // 图片URLs，用逗号分隔
	Likes     int       `json:"likes"`    // 点赞数
	Comments  int       `json:"comments"` // 评论数
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
