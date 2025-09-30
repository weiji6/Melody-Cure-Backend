package model

import (
	"time"
	"gorm.io/gorm"
)

// GeneratedReport AI生成的报告模型
type GeneratedReport struct {
	ID             uint           `gorm:"primaryKey" json:"id"`
	ChildArchiveID string         `gorm:"not null;index" json:"child_archive_id"`
	ReportType     string         `gorm:"type:varchar(50);not null" json:"report_type"` // summary, suggestion
	Content        string         `gorm:"type:text;not null" json:"content"`
	IsEdited       bool           `gorm:"default:false" json:"is_edited"`
	GeneratedAt    time.Time      `gorm:"not null" json:"generated_at"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName 指定表名
func (GeneratedReport) TableName() string {
	return "generated_reports"
}