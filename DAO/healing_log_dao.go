package DAO

import (
	"melody_cure/model"

	"gorm.io/gorm"
)

type HealingLogDAO struct {
	db *gorm.DB
}

func NewHealingLogDAO(db *gorm.DB) *HealingLogDAO {
	return &HealingLogDAO{db: db}
}

// CreateHealingLog 创建疗愈日志
func (dao *HealingLogDAO) CreateHealingLog(log *model.HealingLog) error {
	return dao.db.Create(log).Error
}

// GetHealingLogsByChildID 获取指定儿童的所有疗愈日志
func (dao *HealingLogDAO) GetHealingLogsByChildID(childID uint) ([]model.HealingLog, error) {
	var logs []model.HealingLog
	err := dao.db.Preload("Media").Where("child_archive_id = ?", childID).Order("created_at desc").Find(&logs).Error
	return logs, err
}

// GetHealingLogByID 获取单个疗愈日志详情
func (dao *HealingLogDAO) GetHealingLogByID(logID uint) (*model.HealingLog, error) {
	var log model.HealingLog
	err := dao.db.Preload("Media").First(&log, logID).Error
	return &log, err
}

// DeleteHealingLog 删除疗愈日志
func (dao *HealingLogDAO) DeleteHealingLog(logID uint) error {
	// 同时删除关联的媒体文件
	tx := dao.db.Begin()
	if err := tx.Where("healing_log_id = ?", logID).Delete(&model.LogMedia{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Delete(&model.HealingLog{}, logID).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}