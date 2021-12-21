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

type ImTximApi struct {
}

// CreateImTxim 创建ImTxim
// @Tags ImTxim
// @Summary 创建ImTxim
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.ImTxim true "创建ImTxim"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /imTxim/createImTxim [post]
func (imTximApi *ImTximApi) CreateImTxim(c *gin.Context) {
	var dataObj business.ImTxim
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		response.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	}

	if err := bizSev.GetImTximService().CreateImTxim(dataObj); err != nil {
		global.LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteImTxim 删除ImTxim
// @Tags ImTxim
// @Summary 删除ImTxim
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.ImTxim true "删除ImTxim"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /imTxim/deleteImTxim [delete]
func (imTximApi *ImTximApi) DeleteImTxim(c *gin.Context) {
	var imTxim business.ImTxim
	_ = c.ShouldBindJSON(&imTxim)
	if err := bizSev.GetImTximService().DeleteImTxim(imTxim); err != nil {
		global.LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteImTximByIds 批量删除ImTxim
// @Tags ImTxim
// @Summary 批量删除ImTxim
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ImTxim"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /imTxim/deleteImTximByIds [delete]
func (imTximApi *ImTximApi) DeleteImTximByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := bizSev.GetImTximService().DeleteImTximByIds(IDS); err != nil {
		global.LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateImTxim 更新ImTxim
// @Tags ImTxim
// @Summary 更新ImTxim
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.ImTxim true "更新ImTxim"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /imTxim/updateImTxim [put]
func (imTximApi *ImTximApi) UpdateImTxim(c *gin.Context) {
	var dataObj business.ImTxim
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		response.FailWithMessage("更新失败,"+err.FirstString(), c)
		return
	}

	if err := bizSev.GetImTximService().UpdateImTxim(dataObj); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindImTxim 用id查询ImTxim
// @Tags ImTxim
// @Summary 用id查询ImTxim
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.ImTxim true "用id查询ImTxim"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /imTxim/findImTxim [get]
func (imTximApi *ImTximApi) FindImTxim(c *gin.Context) {
	var imTxim business.ImTxim
	_ = c.ShouldBindQuery(&imTxim)
	reimTxim, err := bizSev.GetImTximService().GetImTxim(imTxim.ID, "")
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.OkWithData(gin.H{"imTxim": nil}, c)
	} else if err != nil {
		global.LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"imTxim": reimTxim}, c)
	}
}

// GetImTximList 分页获取ImTxim列表
// @Tags ImTxim
// @Summary 分页获取ImTxim列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.ImTximSearch true "分页获取ImTxim列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /imTxim/getImTximList [get]
func (imTximApi *ImTximApi) GetImTximList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")

	var pageInfo bizReq.ImTximSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := bizSev.GetImTximService().GetImTximInfoList(pageInfo, createdAtBetween, ""); err != nil {
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
// @Param data body business.ImTxim true "快速更新ImTxim"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router  /imTxim/quickEdit [post]
func (imTximApi *ImTximApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "im_txim"
	//var_dump.Dump(quickEdit)
	if err := commSev.GetCommonDbService().QuickEdit(quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// GetImTximList 分页导出excel ImTxim列表
// @Tags ImTxim
// @Summary 分页导出excel ImTxim列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.ImTximSearch true "分页导出excel ImTxim列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /imTxim/excelList [get]
func (imTximApi *ImTximApi) ExcelList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")

	var pageInfo bizReq.ImTximSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := bizSev.GetImTximService().GetImTximInfoList(pageInfo, createdAtBetween, ""); err != nil {
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
