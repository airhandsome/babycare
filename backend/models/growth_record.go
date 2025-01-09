package models

import "time"

type GrowthRecord struct {
    ID        uint      `json:"id" gorm:"primarykey"`
    UserID    uint      `json:"user_id" gorm:"not null"`
    User      User      `json:"user" gorm:"foreignKey:UserID"`
    Type      string    `json:"type"`      // 记录类型：身高、体重、里程碑等
    Value     float64   `json:"value"`     // 记录值
    Unit      string    `json:"unit"`      // 单位：cm、kg等
    Note      string    `json:"note"`      // 备注
    Date      time.Time `json:"date"`      // 记录日期
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

type Milestone struct {
    ID          uint      `json:"id" gorm:"primarykey"`
    UserID      uint      `json:"user_id" gorm:"not null"`
    User        User      `json:"user" gorm:"foreignKey:UserID"`
    Category    string    `json:"category"`    // 类别：运动、语言、认知等
    Title       string    `json:"title"`       // 里程碑标题
    Completed   bool      `json:"completed"`   // 是否完成
    CompletedAt time.Time `json:"completed_at"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
} 