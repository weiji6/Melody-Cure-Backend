package DAO

import (
	"melody_cure/model"
	"time"

	"gorm.io/gorm"
)

type GeneratedReportDAO struct {
	db *gorm.DB
}

func NewGeneratedReportDAO(db *gorm.DB) *GeneratedReportDAO {
	return &GeneratedReportDAO{db: db}
}

// CreateGeneratedReport 创建AI生成报告
func (dao *GeneratedReportDAO) CreateGeneratedReport(report *model.GeneratedReport) error {
	return dao.db.Create(report).Error
}

// GetGeneratedReportByChildIDAndType 根据儿童档案ID和报告类型获取报告
func (dao *GeneratedReportDAO) GetGeneratedReportByChildIDAndType(childArchiveID, reportType string) (*model.GeneratedReport, error) {
	var report model.GeneratedReport
	err := dao.db.Where("child_archive_id = ? AND report_type = ?", childArchiveID, reportType).
		Order("created_at desc").First(&report).Error
	return &report, err
}

// UpdateGeneratedReport 更新AI生成报告
func (dao *GeneratedReportDAO) UpdateGeneratedReport(report *model.GeneratedReport) error {
	return dao.db.Save(report).Error
}

// DeleteGeneratedReport 删除AI生成报告
func (dao *GeneratedReportDAO) DeleteGeneratedReport(reportID uint) error {
	return dao.db.Delete(&model.GeneratedReport{}, reportID).Error
}

// GetGeneratedReportsByChildID 获取指定儿童的所有报告
func (dao *GeneratedReportDAO) GetGeneratedReportsByChildID(childArchiveID string) ([]model.GeneratedReport, error) {
	var reports []model.GeneratedReport
	err := dao.db.Where("child_archive_id = ?", childArchiveID).
		Order("created_at desc").Find(&reports).Error
	return reports, err
}

// GetGeneratedReportsByChildIDWithDateFilter 获取指定儿童的报告，支持日期筛选
func (dao *GeneratedReportDAO) GetGeneratedReportsByChildIDWithDateFilter(childArchiveID string, startDate, endDate *time.Time) ([]model.GeneratedReport, error) {
	var reports []model.GeneratedReport
	query := dao.db.Where("child_archive_id = ?", childArchiveID)
	
	// 添加日期筛选条件
	if startDate != nil {
		query = query.Where("created_at >= ?", *startDate)
	}
	if endDate != nil {
		query = query.Where("created_at <= ?", *endDate)
	}
	
	err := query.Order("created_at desc").Find(&reports).Error
	return reports, err
}