package business

import (
	"errors"

	"go-cms/global"
	"go-cms/model/business"
	businessReq "go-cms/model/business/request"
	"go-cms/model/common/request"
	"go-cms/model/common/response"
	"go-cms/service"
	commSev "go-cms/service/common"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gvalid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ColCollectApi struct {
}

var colCollectService = service.ServiceGroupApp.BusinessServiceGroup.ColCollectService

// CreateColCollect 创建ColCollect
// @Tags ColCollect
// @Summary 创建ColCollect
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.ColCollect true "创建ColCollect"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /colCollect/createColCollect [post]
func (colCollectApi *ColCollectApi) CreateColCollect(c *gin.Context) {
	var dataObj business.ColCollect
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		response.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	}

	if err := colCollectService.CreateColCollect(dataObj); err != nil {
		global.LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteColCollect 删除ColCollect
// @Tags ColCollect
// @Summary 删除ColCollect
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.ColCollect true "删除ColCollect"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /colCollect/deleteColCollect [delete]
func (colCollectApi *ColCollectApi) DeleteColCollect(c *gin.Context) {
	var colCollect business.ColCollect
	_ = c.ShouldBindJSON(&colCollect)
	if err := colCollectService.DeleteColCollect(colCollect); err != nil {
		global.LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteColCollectByIds 批量删除ColCollect
// @Tags ColCollect
// @Summary 批量删除ColCollect
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ColCollect"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /colCollect/deleteColCollectByIds [delete]
func (colCollectApi *ColCollectApi) DeleteColCollectByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := colCollectService.DeleteColCollectByIds(IDS); err != nil {
		global.LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateColCollect 更新ColCollect
// @Tags ColCollect
// @Summary 更新ColCollect
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.ColCollect true "更新ColCollect"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /colCollect/updateColCollect [put]
func (colCollectApi *ColCollectApi) UpdateColCollect(c *gin.Context) {
	var dataObj business.ColCollect
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		response.FailWithMessage("更新失败,"+err.FirstString(), c)
		return
	}

	if err := colCollectService.UpdateColCollect(dataObj); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindColCollect 用id查询ColCollect
// @Tags ColCollect
// @Summary 用id查询ColCollect
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.ColCollect true "用id查询ColCollect"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /colCollect/findColCollect [get]
func (colCollectApi *ColCollectApi) FindColCollect(c *gin.Context) {
	var colCollect business.ColCollect
	_ = c.ShouldBindQuery(&colCollect)
	err, recolCollect := colCollectService.GetColCollect(colCollect.ID, "")
	if errors.Is(err, gorm.ErrRecordNotFound) {
		//fmt.Println("查询不到数据")
		response.OkWithData(gin.H{"colCollect": nil}, c)
	} else if err != nil {
		global.LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		//response.OkWithData(gin.H{"recolCollect": recolCollect}, c)
		response.OkWithData(gin.H{"colCollect": recolCollect}, c)
	}
}

// GetColCollectList 分页获取ColCollect列表
// @Tags ColCollect
// @Summary 分页获取ColCollect列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query businessReq.ColCollectSearch true "分页获取ColCollect列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /colCollect/getColCollectList [get]
func (colCollectApi *ColCollectApi) GetColCollectList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")

	var pageInfo businessReq.ColCollectSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := colCollectService.GetColCollectInfoList(pageInfo, createdAtBetween, ""); err != nil {
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
// @Param data body business.ColCollect true "快速更新ColCollect"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router  /colCollect/quickEdit [post]
func (colCollectApi *ColCollectApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "col_collect"
	//var_dump.Dump(quickEdit)
	if err := commSev.GetCommonDbService().QuickEdit(quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// GetColCollectList 分页导出excel ColCollect列表
// @Tags ColCollect
// @Summary 分页导出excel ColCollect列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query businessReq.ColCollectSearch true "分页导出excel ColCollect列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /colCollect/excelList [get]
func (colCollectApi *ColCollectApi) ExcelList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")

	var pageInfo businessReq.ColCollectSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := colCollectService.GetColCollectInfoList(pageInfo, createdAtBetween, ""); err != nil {
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
