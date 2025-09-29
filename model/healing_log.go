package model

import (
	"gorm.io/gorm"
)

// HealingLog 疗愈日志模型
type HealingLog struct {
	gorm.Model
	UserID         uint       `gorm:"not null;comment:用户ID"`
	ChildArchiveID uint       `gorm:"not null;comment:儿童档案ID"`
	Content        string     `gorm:"type:text;comment:日志内容"`
	Media          []LogMedia `gorm:"foreignKey:HealingLogID;comment:日志媒体"`
}

// LogMedia 日志媒体模型
type LogMedia struct {
	gorm.Model
	HealingLogID uint   `gorm:"not null;comment:疗愈日志ID"`
	MediaType    string `gorm:"type:varchar(20);not null;comment:媒体类型(image, video)"`
	URL          string `gorm:"type:varchar(255);not null;comment:媒体URL"`
}

func (HealingLog) TableName() string {
	return "healing_logs"
}

func (LogMedia) TableName() string {
	return "log_media"
}