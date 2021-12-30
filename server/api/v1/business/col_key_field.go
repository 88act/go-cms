package business

import (
	"errors"
	"go-cms/global"
	"go-cms/model/business"
	bizReq "go-cms/model/business/request"
	"go-cms/model/common/request"
	"go-cms/model/common/response"
	bizSev "go-cms/service/business"
	commSev "go-cms/service/common"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gvalid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ColKeyFieldApi struct {
}

// CreateColKeyField 创建ColKeyField
// @Tags ColKeyField
// @Summary 创建ColKeyField
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.ColKeyField true "创建ColKeyField"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /colKeyField/createColKeyField [post]
func (colKeyFieldApi *ColKeyFieldApi) CreateColKeyField(c *gin.Context) {
	var dataObj business.ColKeyField
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		response.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	}

	if err := bizSev.GetColKeyFieldService().CreateColKeyField(dataObj); err != nil {
		global.LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteColKeyField 删除ColKeyField
// @Tags ColKeyField
// @Summary 删除ColKeyField
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.ColKeyField true "删除ColKeyField"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /colKeyField/deleteColKeyField [delete]
func (colKeyFieldApi *ColKeyFieldApi) DeleteColKeyField(c *gin.Context) {
	var colKeyField business.ColKeyField
	_ = c.ShouldBindJSON(&colKeyField)
	if err := bizSev.GetColKeyFieldService().DeleteColKeyField(colKeyField); err != nil {
		global.LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteColKeyFieldByIds 批量删除ColKeyField
// @Tags ColKeyField
// @Summary 批量删除ColKeyField
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ColKeyField"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /colKeyField/deleteColKeyFieldByIds [delete]
func (colKeyFieldApi *ColKeyFieldApi) DeleteColKeyFieldByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := bizSev.GetColKeyFieldService().DeleteColKeyFieldByIds(IDS); err != nil {
		global.LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateColKeyField 更新ColKeyField
// @Tags ColKeyField
// @Summary 更新ColKeyField
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.ColKeyField true "更新ColKeyField"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /colKeyField/updateColKeyField [put]
func (colKeyFieldApi *ColKeyFieldApi) UpdateColKeyField(c *gin.Context) {
	var dataObj business.ColKeyField
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		response.FailWithMessage("更新失败,"+err.FirstString(), c)
		return
	}

	if err := bizSev.GetColKeyFieldService().UpdateColKeyField(dataObj); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindColKeyField 用id查询ColKeyField
// @Tags ColKeyField
// @Summary 用id查询ColKeyField
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.ColKeyField true "用id查询ColKeyField"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /colKeyField/findColKeyField [get]
func (colKeyFieldApi *ColKeyFieldApi) FindColKeyField(c *gin.Context) {
	var colKeyField business.ColKeyField
	_ = c.ShouldBindQuery(&colKeyField)
	recolKeyField, err := bizSev.GetColKeyFieldService().GetColKeyField(colKeyField.ID, "")
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.OkWithData(gin.H{"colKeyField": nil}, c)
	} else if err != nil {
		global.LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"colKeyField": recolKeyField}, c)
	}
}

// GetColKeyFieldList 分页获取ColKeyField列表
// @Tags ColKeyField
// @Summary 分页获取ColKeyField列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.ColKeyFieldSearch true "分页获取ColKeyField列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /colKeyField/getColKeyFieldList [get]
func (colKeyFieldApi *ColKeyFieldApi) GetColKeyFieldList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")

	var pageInfo bizReq.ColKeyFieldSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := bizSev.GetColKeyFieldService().GetColKeyFieldInfoList(pageInfo, createdAtBetween, ""); err != nil {
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

// QuickEdit 快速更新
// @Tags QuickEdit
// @Summary 快速更新
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.ColKeyField true "快速更新ColKeyField"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router  /colKeyField/quickEdit [post]
func (colKeyFieldApi *ColKeyFieldApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "col_key_field"
	if err := commSev.GetCommonDbSev().QuickEdit(quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// GetColKeyFieldList 分页导出excel ColKeyField列表
// @Tags ColKeyField
// @Summary 分页导出excel ColKeyField列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.ColKeyFieldSearch true "分页导出excel ColKeyField列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /colKeyField/excelList [get]
func (colKeyFieldApi *ColKeyFieldApi) ExcelList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")

	var pageInfo bizReq.ColKeyFieldSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := bizSev.GetColKeyFieldService().GetColKeyFieldInfoList(pageInfo, createdAtBetween, ""); err != nil {
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
