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

type ImTxMsgFileApi struct {
}

// CreateImTxMsgFile 创建ImTxMsgFile
// @Tags ImTxMsgFile
// @Summary 创建ImTxMsgFile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.ImTxMsgFile true "创建ImTxMsgFile"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /imTxMsgFile/createImTxMsgFile [post]
func (imTxMsgFileApi *ImTxMsgFileApi) CreateImTxMsgFile(c *gin.Context) {
	var dataObj business.ImTxMsgFile
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		response.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	}

	if err := bizSev.GetImTxMsgFileService().CreateImTxMsgFile(dataObj); err != nil {
		global.LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteImTxMsgFile 删除ImTxMsgFile
// @Tags ImTxMsgFile
// @Summary 删除ImTxMsgFile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.ImTxMsgFile true "删除ImTxMsgFile"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /imTxMsgFile/deleteImTxMsgFile [delete]
func (imTxMsgFileApi *ImTxMsgFileApi) DeleteImTxMsgFile(c *gin.Context) {
	var imTxMsgFile business.ImTxMsgFile
	_ = c.ShouldBindJSON(&imTxMsgFile)
	if err := bizSev.GetImTxMsgFileService().DeleteImTxMsgFile(imTxMsgFile); err != nil {
		global.LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteImTxMsgFileByIds 批量删除ImTxMsgFile
// @Tags ImTxMsgFile
// @Summary 批量删除ImTxMsgFile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ImTxMsgFile"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /imTxMsgFile/deleteImTxMsgFileByIds [delete]
func (imTxMsgFileApi *ImTxMsgFileApi) DeleteImTxMsgFileByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := bizSev.GetImTxMsgFileService().DeleteImTxMsgFileByIds(IDS); err != nil {
		global.LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateImTxMsgFile 更新ImTxMsgFile
// @Tags ImTxMsgFile
// @Summary 更新ImTxMsgFile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.ImTxMsgFile true "更新ImTxMsgFile"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /imTxMsgFile/updateImTxMsgFile [put]
func (imTxMsgFileApi *ImTxMsgFileApi) UpdateImTxMsgFile(c *gin.Context) {
	var dataObj business.ImTxMsgFile
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		response.FailWithMessage("更新失败,"+err.FirstString(), c)
		return
	}

	if err := bizSev.GetImTxMsgFileService().UpdateImTxMsgFile(dataObj); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindImTxMsgFile 用id查询ImTxMsgFile
// @Tags ImTxMsgFile
// @Summary 用id查询ImTxMsgFile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.ImTxMsgFile true "用id查询ImTxMsgFile"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /imTxMsgFile/findImTxMsgFile [get]
func (imTxMsgFileApi *ImTxMsgFileApi) FindImTxMsgFile(c *gin.Context) {
	var imTxMsgFile business.ImTxMsgFile
	_ = c.ShouldBindQuery(&imTxMsgFile)
	reimTxMsgFile, err := bizSev.GetImTxMsgFileService().GetImTxMsgFile(imTxMsgFile.ID, "")
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.OkWithData(gin.H{"imTxMsgFile": nil}, c)
	} else if err != nil {
		global.LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"imTxMsgFile": reimTxMsgFile}, c)
	}
}

// GetImTxMsgFileList 分页获取ImTxMsgFile列表
// @Tags ImTxMsgFile
// @Summary 分页获取ImTxMsgFile列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.ImTxMsgFileSearch true "分页获取ImTxMsgFile列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /imTxMsgFile/getImTxMsgFileList [get]
func (imTxMsgFileApi *ImTxMsgFileApi) GetImTxMsgFileList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")

	var pageInfo bizReq.ImTxMsgFileSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := bizSev.GetImTxMsgFileService().GetImTxMsgFileInfoList(pageInfo, createdAtBetween, ""); err != nil {
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
// @Param data body business.ImTxMsgFile true "快速更新ImTxMsgFile"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router  /imTxMsgFile/quickEdit [post]
func (imTxMsgFileApi *ImTxMsgFileApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "im_tx_msg_file"
	//var_dump.Dump(quickEdit)
	if err := commSev.GetCommonDbService().QuickEdit(quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// GetImTxMsgFileList 分页导出excel ImTxMsgFile列表
// @Tags ImTxMsgFile
// @Summary 分页导出excel ImTxMsgFile列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.ImTxMsgFileSearch true "分页导出excel ImTxMsgFile列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /imTxMsgFile/excelList [get]
func (imTxMsgFileApi *ImTxMsgFileApi) ExcelList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")

	var pageInfo bizReq.ImTxMsgFileSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := bizSev.GetImTxMsgFileService().GetImTxMsgFileInfoList(pageInfo, createdAtBetween, ""); err != nil {
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
