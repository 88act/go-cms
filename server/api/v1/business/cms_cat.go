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

type CmsCatApi struct {
}

// CreateCmsCat 创建CmsCat
// @Tags CmsCat
// @Summary 创建CmsCat
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.CmsCat true "创建CmsCat"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmsCat/createCmsCat [post]
func (cmsCatApi *CmsCatApi) CreateCmsCat(c *gin.Context) {
	var dataObj business.CmsCat
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		response.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	}

	if err := bizSev.GetCmsCatService().CreateCmsCat(dataObj); err != nil {
		global.LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCmsCat 删除CmsCat
// @Tags CmsCat
// @Summary 删除CmsCat
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.CmsCat true "删除CmsCat"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmsCat/deleteCmsCat [delete]
func (cmsCatApi *CmsCatApi) DeleteCmsCat(c *gin.Context) {
	var cmsCat business.CmsCat
	_ = c.ShouldBindJSON(&cmsCat)
	if err := bizSev.GetCmsCatService().DeleteCmsCat(cmsCat); err != nil {
		global.LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCmsCatByIds 批量删除CmsCat
// @Tags CmsCat
// @Summary 批量删除CmsCat
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CmsCat"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /cmsCat/deleteCmsCatByIds [delete]
func (cmsCatApi *CmsCatApi) DeleteCmsCatByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := bizSev.GetCmsCatService().DeleteCmsCatByIds(IDS); err != nil {
		global.LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCmsCat 更新CmsCat
// @Tags CmsCat
// @Summary 更新CmsCat
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.CmsCat true "更新CmsCat"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cmsCat/updateCmsCat [put]
func (cmsCatApi *CmsCatApi) UpdateCmsCat(c *gin.Context) {
	var dataObj business.CmsCat
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		response.FailWithMessage("更新失败,"+err.FirstString(), c)
		return
	}

	if err := bizSev.GetCmsCatService().UpdateCmsCat(dataObj); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCmsCat 用id查询CmsCat
// @Tags CmsCat
// @Summary 用id查询CmsCat
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.CmsCat true "用id查询CmsCat"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cmsCat/findCmsCat [get]
func (cmsCatApi *CmsCatApi) FindCmsCat(c *gin.Context) {
	var cmsCat business.CmsCat
	_ = c.ShouldBindQuery(&cmsCat)
	recmsCat, err := bizSev.GetCmsCatService().GetCmsCat(cmsCat.ID, "")
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.OkWithData(gin.H{"cmsCat": nil}, c)
	} else if err != nil {
		global.LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"cmsCat": recmsCat}, c)
	}
}

// GetCmsCatList 分页获取CmsCat列表
// @Tags CmsCat
// @Summary 分页获取CmsCat列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.CmsCatSearch true "分页获取CmsCat列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmsCat/getCmsCatList [get]
func (cmsCatApi *CmsCatApi) GetCmsCatList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")

	var pageInfo bizReq.CmsCatSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := bizSev.GetCmsCatService().GetCmsCatInfoList(pageInfo, createdAtBetween, ""); err != nil {
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
// @Param data body business.CmsCat true "快速更新CmsCat"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router  /cmsCat/quickEdit [post]
func (cmsCatApi *CmsCatApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "cms_cat"
	//var_dump.Dump(quickEdit)
	if err := commSev.GetCommonDbService().QuickEdit(quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// GetCmsCatList 分页导出excel CmsCat列表
// @Tags CmsCat
// @Summary 分页导出excel CmsCat列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.CmsCatSearch true "分页导出excel CmsCat列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmsCat/excelList [get]
func (cmsCatApi *CmsCatApi) ExcelList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")

	var pageInfo bizReq.CmsCatSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := bizSev.GetCmsCatService().GetCmsCatInfoList(pageInfo, createdAtBetween, ""); err != nil {
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
