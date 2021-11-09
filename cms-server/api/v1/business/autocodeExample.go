package business

import (
	"github.com/88act/go-cms/server/global"
	"github.com/88act/go-cms/server/model/business"
	businessReq "github.com/88act/go-cms/server/model/business/request"
	"github.com/88act/go-cms/server/model/common/response"
	"github.com/88act/go-cms/server/service"
	"github.com/88act/go-cms/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BusinessExampleApi struct {
}

var businessExampleService = service.ServiceGroupApp.BusinessServiceGroup.BusinessExampleService

// @Tags BusinessExample
// @Summary 创建BusinessExample
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.BusinessExample true "BusinessExample模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /businessExample/createBusinessExample [post]
func (businessExampleApi *BusinessExampleApi) CreateBusinessExample(c *gin.Context) {
	var businessExample business.BusinessExample
	_ = c.ShouldBindJSON(&businessExample)
	if err := businessExampleService.CreateBusinessExample(businessExample); err != nil {
		global.LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags BusinessExample
// @Summary 删除BusinessExample
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.BusinessExample true "BusinessExample模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /businessExample/deleteBusinessExample [delete]
func (businessExampleApi *BusinessExampleApi) DeleteBusinessExample(c *gin.Context) {
	var businessExample business.BusinessExample
	_ = c.ShouldBindJSON(&businessExample)
	if err := businessExampleService.DeleteBusinessExample(businessExample); err != nil {
		global.LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags BusinessExample
// @Summary 更新BusinessExample
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.BusinessExample true "更新BusinessExample"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /businessExample/updateBusinessExample [put]
func (businessExampleApi *BusinessExampleApi) UpdateBusinessExample(c *gin.Context) {
	var businessExample business.BusinessExample
	_ = c.ShouldBindJSON(&businessExample)
	if err := businessExampleService.UpdateBusinessExample(&businessExample); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags BusinessExample
// @Summary 用id查询BusinessExample
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.BusinessExample true "用id查询BusinessExample"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /businessExample/findBusinessExample [get]
func (businessExampleApi *BusinessExampleApi) FindBusinessExample(c *gin.Context) {
	var businessExample business.BusinessExample
	_ = c.ShouldBindQuery(&businessExample)
	if err := utils.Verify(businessExample, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, reBusinessExample := businessExampleService.GetBusinessExample(businessExample.ID); err != nil {
		global.LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithDetailed(gin.H{"reBusinessExample": reBusinessExample}, "查询成功", c)
	}
}

// @Tags BusinessExample
// @Summary 分页获取BusinessExample列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query businessReq.BusinessExampleSearch true "页码, 每页大小, 搜索条件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /businessExample/getBusinessExampleList [get]
func (businessExampleApi *BusinessExampleApi) GetBusinessExampleList(c *gin.Context) {
	var pageInfo businessReq.BusinessExampleSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := businessExampleService.GetBusinessExampleInfoList(pageInfo); err != nil {
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
