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

type BasicFileApi struct {
}

// CreateBasicFile 创建BasicFile
// @Tags BasicFile
// @Summary 创建BasicFile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.BasicFile true "创建BasicFile"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /basicFile/createBasicFile [post]
func (basicFileApi *BasicFileApi) CreateBasicFile(c *gin.Context) {
	var dataObj business.BasicFile
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		response.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	}

	if err := bizSev.GetBasicFileService().CreateBasicFile(dataObj); err != nil {
		global.LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBasicFile 删除BasicFile
// @Tags BasicFile
// @Summary 删除BasicFile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.BasicFile true "删除BasicFile"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /basicFile/deleteBasicFile [delete]
func (basicFileApi *BasicFileApi) DeleteBasicFile(c *gin.Context) {
	var basicFile business.BasicFile
	_ = c.ShouldBindJSON(&basicFile)
	if err := bizSev.GetBasicFileService().DeleteBasicFile(basicFile); err != nil {
		global.LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBasicFileByIds 批量删除BasicFile
// @Tags BasicFile
// @Summary 批量删除BasicFile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BasicFile"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /basicFile/deleteBasicFileByIds [delete]
func (basicFileApi *BasicFileApi) DeleteBasicFileByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := bizSev.GetBasicFileService().DeleteBasicFileByIds(IDS); err != nil {
		global.LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBasicFile 更新BasicFile
// @Tags BasicFile
// @Summary 更新BasicFile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.BasicFile true "更新BasicFile"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /basicFile/updateBasicFile [put]
func (basicFileApi *BasicFileApi) UpdateBasicFile(c *gin.Context) {
	var dataObj business.BasicFile
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		response.FailWithMessage("更新失败,"+err.FirstString(), c)
		return
	}

	if err := bizSev.GetBasicFileService().UpdateBasicFile(dataObj); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBasicFile 用id查询BasicFile
// @Tags BasicFile
// @Summary 用id查询BasicFile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.BasicFile true "用id查询BasicFile"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /basicFile/findBasicFile [get]
func (basicFileApi *BasicFileApi) FindBasicFile(c *gin.Context) {
	var basicFile business.BasicFile
	_ = c.ShouldBindQuery(&basicFile)
	rebasicFile, err := bizSev.GetBasicFileService().GetBasicFile(basicFile.ID, "")
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.OkWithData(gin.H{"basicFile": nil}, c)
	} else if err != nil {
		global.LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"basicFile": rebasicFile}, c)
	}
}

// GetBasicFileList 分页获取BasicFile列表
// @Tags BasicFile
// @Summary 分页获取BasicFile列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.BasicFileSearch true "分页获取BasicFile列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /basicFile/getBasicFileList [get]
func (basicFileApi *BasicFileApi) GetBasicFileList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")

	var pageInfo bizReq.BasicFileSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := bizSev.GetBasicFileService().GetBasicFileInfoList(pageInfo, createdAtBetween, ""); err != nil {
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
// @Param data body business.BasicFile true "快速更新BasicFile"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router  /basicFile/quickEdit [post]
func (basicFileApi *BasicFileApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "basic_file"
	//var_dump.Dump(quickEdit)
	if err := commSev.GetCommonDbService().QuickEdit(quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// GetBasicFileList 分页导出excel BasicFile列表
// @Tags BasicFile
// @Summary 分页导出excel BasicFile列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.BasicFileSearch true "分页导出excel BasicFile列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /basicFile/excelList [get]
func (basicFileApi *BasicFileApi) ExcelList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")

	var pageInfo bizReq.BasicFileSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := bizSev.GetBasicFileService().GetBasicFileInfoList(pageInfo, createdAtBetween, ""); err != nil {
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
