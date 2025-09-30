package controller

import (
	"melody_cure/api/request"
	"melody_cure/api/response"
	"melody_cure/service"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type AIReportController struct {
	aiReportService *service.AIReportService
}

func NewAIReportController(aiReportService *service.AIReportService) *AIReportController {
	return &AIReportController{
		aiReportService: aiReportService,
	}
}

// GenerateReport 生成AI报告
func (c *AIReportController) GenerateReport(ctx *gin.Context) {
	var req request.GenerateReportRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误: " + err.Error()})
		return
	}

	// 解析日期参数
	var startDate, endDate *time.Time
	if req.StartDate != "" {
		if parsed, err := time.Parse("2006-01-02", req.StartDate); err == nil {
			startDate = &parsed
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "开始日期格式错误"})
			return
		}
	}
	if req.EndDate != "" {
		if parsed, err := time.Parse("2006-01-02", req.EndDate); err == nil {
			endDate = &parsed
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "结束日期格式错误"})
			return
		}
	}

	// 生成报告
	report, err := c.aiReportService.GenerateReport(strconv.FormatUint(uint64(req.ChildArchiveID), 10), req.ReportType, startDate, endDate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "生成报告失败: " + err.Error()})
		return
	}

	// 构建响应
	resp := response.GeneratedReportResponse{
		ID:              report.ID,
		ChildArchiveID:  req.ChildArchiveID,
		ReportType:      report.ReportType,
		Content:         report.Content,
		IsEdited:        report.IsEdited,
		GeneratedAt:     report.GeneratedAt,
		UpdatedAt:       report.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "报告生成成功",
		"data":    resp,
	})
}

// UpdateReportContent 更新报告内容
func (c *AIReportController) UpdateReportContent(ctx *gin.Context) {
	reportIDStr := ctx.Param("id")
	reportID, err := strconv.ParseUint(reportIDStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "报告ID格式错误"})
		return
	}

	var req request.UpdateGeneratedContentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误: " + err.Error()})
		return
	}

	// 更新报告内容
	err = c.aiReportService.UpdateReportContent(uint(reportID), req.Content)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "更新报告失败: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "报告更新成功",
	})
}

// GetReport 获取报告
func (c *AIReportController) GetReport(ctx *gin.Context) {
	childArchiveID := ctx.Query("child_archive_id")
	reportType := ctx.Query("report_type")

	if childArchiveID == "" || reportType == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "缺少必要参数"})
		return
	}

	// 获取报告
	report, err := c.aiReportService.GetReportByChildIDAndType(childArchiveID, reportType)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "报告不存在"})
		return
	}

	// 将字符串ID转换为uint
	childArchiveIDUint, err := strconv.ParseUint(childArchiveID, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "儿童档案ID格式错误"})
		return
	}

	// 构建响应
	resp := response.GeneratedReportResponse{
		ID:              report.ID,
		ChildArchiveID:  uint(childArchiveIDUint),
		ReportType:      report.ReportType,
		Content:         report.Content,
		IsEdited:        report.IsEdited,
		GeneratedAt:     report.GeneratedAt,
		UpdatedAt:       report.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "获取报告成功",
		"data":    resp,
	})
}