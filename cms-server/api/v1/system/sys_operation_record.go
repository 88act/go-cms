package system

import (
	"github.com/88act/go-cms/server/global"
	"github.com/88act/go-cms/server/model/common/request"
	"github.com/88act/go-cms/server/model/common/response"
	"github.com/88act/go-cms/server/model/system"
	systemReq "github.com/88act/go-cms/server/model/system/request"
	"github.com/88act/go-cms/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type OperationRecordApi struct {
}

// @Tags SysOperationRecord
// @Summary 创建SysOperationRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysOperationRecord true "创建SysOperationRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysOperationRecord/createSysOperationRecord [post]
func (s *OperationRecordApi) CreateSysOperationRecord(c *gin.Context) {
	var sysOperationRecord system.SysOperationRecord
	_ = c.ShouldBindJSON(&sysOperationRecord)
	if err := operationRecordService.CreateSysOperationRecord(sysOperationRecord); err != nil {
		global.LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags SysOperationRecord
// @Summary 删除SysOperationRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysOperationRecord true "SysOperationRecord模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysOperationRecord/deleteSysOperationRecord [delete]
func (s *OperationRecordApi) DeleteSysOperationRecord(c *gin.Context) {
	var sysOperationRecord system.SysOperationRecord
	_ = c.ShouldBindJSON(&sysOperationRecord)
	if err := operationRecordService.DeleteSysOperationRecord(sysOperationRecord); err != nil {
		global.LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags SysOperationRecord
// @Summary 批量删除SysOperationRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SysOperationRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /sysOperationRecord/deleteSysOperationRecordByIds [delete]
func (s *OperationRecordApi) DeleteSysOperationRecordByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := operationRecordService.DeleteSysOperationRecordByIds(IDS); err != nil {
		global.LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags SysOperationRecord
// @Summary 用id查询SysOperationRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query system.SysOperationRecord true "Id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysOperationRecord/findSysOperationRecord [get]
func (s *OperationRecordApi) FindSysOperationRecord(c *gin.Context) {
	var sysOperationRecord system.SysOperationRecord
	_ = c.ShouldBindQuery(&sysOperationRecord)
	if err := utils.Verify(sysOperationRecord, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, resysOperationRecord := operationRecordService.GetSysOperationRecord(sysOperationRecord.ID); err != nil {
		global.LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithDetailed(gin.H{"resysOperationRecord": resysOperationRecord}, "查询成功", c)
	}
}

// @Tags SysOperationRecord
// @Summary 分页获取SysOperationRecord列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.SysOperationRecordSearch true "页码, 每页大小, 搜索条件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysOperationRecord/getSysOperationRecordList [get]
func (s *OperationRecordApi) GetSysOperationRecordList(c *gin.Context) {
	var pageInfo systemReq.SysOperationRecordSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := operationRecordService.GetSysOperationRecordInfoList(pageInfo); err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
