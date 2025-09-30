package service

import (
	"melody_cure/DAO"
	"melody_cure/model"
	"time"
)

type HealingLogService struct {
	healingLogDAO *DAO.HealingLogDAO
}

func NewHealingLogService(healingLogDAO *DAO.HealingLogDAO) *HealingLogService {
	return &HealingLogService{healingLogDAO: healingLogDAO}
}

// CreateHealingLog 创建疗愈日志
func (s *HealingLogService) CreateHealingLog(log *model.HealingLog) error {
	return s.healingLogDAO.CreateHealingLog(log)
}

// GetHealingLogsByChildID 获取指定儿童的所有疗愈日志
func (s *HealingLogService) GetHealingLogsByChildID(childID uint) ([]model.HealingLog, error) {
	return s.healingLogDAO.GetHealingLogsByChildID(childID)
}

// GetHealingLogByID 获取单个疗愈日志详情
func (s *HealingLogService) GetHealingLogByID(logID uint) (*model.HealingLog, error) {
	return s.healingLogDAO.GetHealingLogByID(logID)
}

// DeleteHealingLog 删除疗愈日志
func (s *HealingLogService) DeleteHealingLog(logID uint) error {
	return s.healingLogDAO.DeleteHealingLog(logID)
}

// GetHealingLogsByChildIDWithDateFilter 获取指定儿童的疗愈日志，支持日期筛选
func (s *HealingLogService) GetHealingLogsByChildIDWithDateFilter(childID uint, startDate, endDate *time.Time) ([]model.HealingLog, error) {
	return s.healingLogDAO.GetHealingLogsByChildIDWithDateFilter(childID, startDate, endDate)
}