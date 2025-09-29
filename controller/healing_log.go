package controller

import (
	"melody_cure/api/response"
	"melody_cure/model"
	"melody_cure/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HealingLogController struct {
	healingLogService *service.HealingLogService
}

func NewHealingLogController(healingLogService *service.HealingLogService) *HealingLogController {
	return &HealingLogController{healingLogService: healingLogService}
}

// CreateHealingLog 创建疗愈日志
// @Summary 创建疗愈日志
// @Description 创建一条新的疗愈日志，记录儿童成长进步和疗愈前后对比
// @Tags 疗愈日志
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param healing_log body model.HealingLog true "疗愈日志信息"
// @Success 200 {object} response.SuccessResponse "创建成功"
// @Failure 400 {object} response.ErrorResponse "参数错误"
// @Failure 401 {object} response.ErrorResponse "未认证"
// @Failure 500 {object} response.ErrorResponse "创建失败"
// @Router /api/healing-log [post]
func (c *HealingLogController) CreateHealingLog(ctx *gin.Context) {
	var log model.HealingLog
	if err := ctx.ShouldBindJSON(&log); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{Code: http.StatusBadRequest, Message: "参数错误"})
		return
	}

	// 从JWT获取用户ID
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, response.ErrorResponse{Code: http.StatusUnauthorized, Message: "未认证"})
		return
	}
	log.UserID = userID.(uint)

	if err := c.healingLogService.CreateHealingLog(&log); err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{Code: http.StatusInternalServerError, Message: "创建失败: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse{Code: http.StatusOK, Message: "创建成功"})
}

// GetHealingLogsByChildID 获取指定儿童的所有疗愈日志
// @Summary 获取儿童疗愈日志列表
// @Description 获取指定儿童的所有疗愈日志，按时间线排序显示成长进步
// @Tags 疗愈日志
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param child_id path int true "儿童档案ID"
// @Success 200 {object} object{code=int,data=[]model.HealingLog} "获取成功"
// @Failure 400 {object} response.ErrorResponse "无效的儿童ID"
// @Failure 401 {object} response.ErrorResponse "未认证"
// @Failure 500 {object} response.ErrorResponse "获取失败"
// @Router /api/healing-log/child/{child_id} [get]
func (c *HealingLogController) GetHealingLogsByChildID(ctx *gin.Context) {
	childIDStr := ctx.Param("child_id")
	childID, err := strconv.ParseUint(childIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{Code: http.StatusBadRequest, Message: "无效的儿童ID"})
		return
	}

	logs, err := c.healingLogService.GetHealingLogsByChildID(uint(childID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{Code: http.StatusInternalServerError, Message: "获取失败: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": logs})
}

// GetHealingLogByID 获取单个疗愈日志详情
// @Summary 获取疗愈日志详情
// @Description 获取单个疗愈日志的详细信息，包括文字内容和媒体文件
// @Tags 疗愈日志
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param log_id path int true "日志ID"
// @Success 200 {object} object{code=int,data=model.HealingLog} "获取成功"
// @Failure 400 {object} response.ErrorResponse "无效的日志ID"
// @Failure 401 {object} response.ErrorResponse "未认证"
// @Failure 500 {object} response.ErrorResponse "获取失败"
// @Router /api/healing-log/{log_id} [get]
func (c *HealingLogController) GetHealingLogByID(ctx *gin.Context) {
	logIDStr := ctx.Param("log_id")
	logID, err := strconv.ParseUint(logIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{Code: http.StatusBadRequest, Message: "无效的日志ID"})
		return
	}

	log, err := c.healingLogService.GetHealingLogByID(uint(logID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{Code: http.StatusInternalServerError, Message: "获取失败: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": log})
}

// DeleteHealingLog 删除疗愈日志
// @Summary 删除疗愈日志
// @Description 删除指定的疗愈日志及其关联的媒体文件
// @Tags 疗愈日志
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param log_id path int true "日志ID"
// @Success 200 {object} response.SuccessResponse "删除成功"
// @Failure 400 {object} response.ErrorResponse "无效的日志ID"
// @Failure 401 {object} response.ErrorResponse "未认证"
// @Failure 500 {object} response.ErrorResponse "删除失败"
// @Router /api/healing-log/{log_id} [delete]
func (c *HealingLogController) DeleteHealingLog(ctx *gin.Context) {
	logIDStr := ctx.Param("log_id")
	logID, err := strconv.ParseUint(logIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{Code: http.StatusBadRequest, Message: "无效的日志ID"})
		return
	}

	if err := c.healingLogService.DeleteHealingLog(uint(logID)); err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{Code: http.StatusInternalServerError, Message: "删除失败: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse{Code: http.StatusOK, Message: "删除成功"})
}