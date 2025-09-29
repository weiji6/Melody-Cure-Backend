package model

import (
	"time"
	"gorm.io/gorm"
)

// ImageToken 图床token模型
type ImageToken struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	UserID    uint           `gorm:"index;not null" json:"user_id"`
	Token     string         `gorm:"unique;not null" json:"token"`
	Purpose   string         `json:"purpose"`   // 用途描述
	ExpiresAt time.Time      `json:"expires_at"` // 过期时间
	IsUsed    bool           `gorm:"default:false" json:"is_used"` // 是否已使用
	UsedAt    *time.Time     `json:"used_at"`   // 使用时间
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}