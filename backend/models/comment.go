package models

import "time"

type Comment struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	User      User      `json:"user" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	PostID    uint      `json:"post_id" gorm:"not null"`
	ForumPost ForumPost `json:"post" gorm:"foreignKey:PostID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Content   string    `json:"content"`
	ParentID  *uint     `json:"parent_id"` // 父评论ID，用于回复功能
	Likes     int       `json:"likes"`      // 点赞数
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
