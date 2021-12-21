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

type K8sNodesApi struct {
}

// CreateK8sNodes 创建K8sNodes
// @Tags K8sNodes
// @Summary 创建K8sNodes
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.K8sNodes true "创建K8sNodes"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /k8sNodes/createK8sNodes [post]
func (k8sNodesApi *K8sNodesApi) CreateK8sNodes(c *gin.Context) {
	var dataObj business.K8sNodes
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		response.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	}

	if err := bizSev.GetK8sNodesService().CreateK8sNodes(dataObj); err != nil {
		global.LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteK8sNodes 删除K8sNodes
// @Tags K8sNodes
// @Summary 删除K8sNodes
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.K8sNodes true "删除K8sNodes"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /k8sNodes/deleteK8sNodes [delete]
func (k8sNodesApi *K8sNodesApi) DeleteK8sNodes(c *gin.Context) {
	var k8sNodes business.K8sNodes
	_ = c.ShouldBindJSON(&k8sNodes)
	if err := bizSev.GetK8sNodesService().DeleteK8sNodes(k8sNodes); err != nil {
		global.LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteK8sNodesByIds 批量删除K8sNodes
// @Tags K8sNodes
// @Summary 批量删除K8sNodes
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除K8sNodes"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /k8sNodes/deleteK8sNodesByIds [delete]
func (k8sNodesApi *K8sNodesApi) DeleteK8sNodesByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := bizSev.GetK8sNodesService().DeleteK8sNodesByIds(IDS); err != nil {
		global.LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateK8sNodes 更新K8sNodes
// @Tags K8sNodes
// @Summary 更新K8sNodes
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.K8sNodes true "更新K8sNodes"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /k8sNodes/updateK8sNodes [put]
func (k8sNodesApi *K8sNodesApi) UpdateK8sNodes(c *gin.Context) {
	var dataObj business.K8sNodes
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		response.FailWithMessage("更新失败,"+err.FirstString(), c)
		return
	}

	if err := bizSev.GetK8sNodesService().UpdateK8sNodes(dataObj); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindK8sNodes 用id查询K8sNodes
// @Tags K8sNodes
// @Summary 用id查询K8sNodes
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.K8sNodes true "用id查询K8sNodes"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /k8sNodes/findK8sNodes [get]
func (k8sNodesApi *K8sNodesApi) FindK8sNodes(c *gin.Context) {
	var k8sNodes business.K8sNodes
	_ = c.ShouldBindQuery(&k8sNodes)
	rek8sNodes, err := bizSev.GetK8sNodesService().GetK8sNodes(k8sNodes.ID, "")
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.OkWithData(gin.H{"k8sNodes": nil}, c)
	} else if err != nil {
		global.LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"k8sNodes": rek8sNodes}, c)
	}
}

// GetK8sNodesList 分页获取K8sNodes列表
// @Tags K8sNodes
// @Summary 分页获取K8sNodes列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.K8sNodesSearch true "分页获取K8sNodes列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /k8sNodes/getK8sNodesList [get]
func (k8sNodesApi *K8sNodesApi) GetK8sNodesList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")

	var pageInfo bizReq.K8sNodesSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := bizSev.GetK8sNodesService().GetK8sNodesInfoList(pageInfo, createdAtBetween, ""); err != nil {
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
// @Param data body business.K8sNodes true "快速更新K8sNodes"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router  /k8sNodes/quickEdit [post]
func (k8sNodesApi *K8sNodesApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "k8s_nodes"
	//var_dump.Dump(quickEdit)
	if err := commSev.GetCommonDbService().QuickEdit(quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// GetK8sNodesList 分页导出excel K8sNodes列表
// @Tags K8sNodes
// @Summary 分页导出excel K8sNodes列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.K8sNodesSearch true "分页导出excel K8sNodes列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /k8sNodes/excelList [get]
func (k8sNodesApi *K8sNodesApi) ExcelList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")

	var pageInfo bizReq.K8sNodesSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := bizSev.GetK8sNodesService().GetK8sNodesInfoList(pageInfo, createdAtBetween, ""); err != nil {
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
