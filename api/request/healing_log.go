package request

// GenerateReportRequest AI生成报告请求
type GenerateReportRequest struct {
	ChildArchiveID uint   `json:"child_archive_id" binding:"required" example:"1"`
	StartDate      string `json:"start_date,omitempty" example:"2024-01-01"`
	EndDate        string `json:"end_date,omitempty" example:"2024-01-31"`
	ReportType     string `json:"report_type" binding:"required,oneof=summary suggestion" example:"summary"`
}

// UpdateGeneratedContentRequest 更新AI生成内容请求
type UpdateGeneratedContentRequest struct {
	Content string `json:"content" binding:"required" example:"更新后的内容"`
}