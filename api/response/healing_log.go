package response

import "time"

// GeneratedReportResponse AI生成报告响应
type GeneratedReportResponse struct {
	ID             uint      `json:"id" example:"1"`
	ChildArchiveID uint      `json:"child_archive_id" example:"1"`
	ReportType     string    `json:"report_type" example:"summary"`
	Content        string    `json:"content" example:"本月儿童在情绪管理方面有显著进步..."`
	IsEdited       bool      `json:"is_edited" example:"false"`
	GeneratedAt    time.Time `json:"generated_at" example:"2024-01-15T10:30:00Z"`
	UpdatedAt      time.Time `json:"updated_at" example:"2024-01-15T10:30:00Z"`
}

// HealingLogWithReportResponse 包含AI生成报告的疗愈记录响应
type HealingLogWithReportResponse struct {
	Logs           []HealingLogResponse      `json:"logs"`
	GeneratedReport *GeneratedReportResponse `json:"generated_report,omitempty"`
}

// HealingLogResponse 疗愈记录响应
type HealingLogResponse struct {
	ID             uint                `json:"id" example:"1"`
	UserID         uint                `json:"user_id" example:"1"`
	ChildArchiveID uint                `json:"child_archive_id" example:"1"`
	Content        string              `json:"content" example:"今天孩子情绪很稳定"`
	Media          []LogMediaResponse  `json:"media,omitempty"`
	CreatedAt      time.Time           `json:"created_at" example:"2024-01-15T10:30:00Z"`
}

// LogMediaResponse 日志媒体响应
type LogMediaResponse struct {
	ID           uint   `json:"id" example:"1"`
	MediaType    string `json:"media_type" example:"image"`
	MediaURL     string `json:"media_url" example:"https://example.com/image.jpg"`
	Description  string `json:"description,omitempty" example:"孩子的笑脸照片"`
}