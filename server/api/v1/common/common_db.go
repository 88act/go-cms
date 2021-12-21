package common

import (
	"fmt"
	"go-cms/global"
	bizReq "go-cms/model/business/request"
	"go-cms/model/common/request"
	"go-cms/model/common/response"
	bizSev "go-cms/service/business"
	commSev "go-cms/service/common"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gconv"
	"go.uber.org/zap"
)

type CommonDbApi struct {
}

// QuickEdit 更新QuickEdit
// @Tags common
// @Summary 更新QuickEdit
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.CmsAd true "更新 QuickEdit"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /comm/updateCmsAd [put]
func (o *CommonDbApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	//var_dump.Dump(quickEdit)

	if err := commSev.GetCommonDbService().QuickEdit(quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// Test_db_get 测试Test_db_get
// @Tags common
// @Summary 测试Test_db_get
// @accept application/json
// @Produce application/json
// @Param name query string true "用户姓名"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"成功"}"
// @Router /commDb/test_db_get [get]
func (o *CommonDbApi) Test_db_get(c *gin.Context) {
	param := c.DefaultQuery("name", "没收到name参数")
	fmt.Println("测试数据库 Test_db_get param =" + param)
	global.LOG.Error("测试数据库 Test_db_get param =" + param)
	pageInfo := bizReq.CmsCatSearch{}
	pageInfo.Page = 1
	pageInfo.PageSize = 20
	if list, total, err := bizSev.GetCmsCatService().GetCmsCatInfoList(pageInfo, nil, ""); err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		global.LOG.Debug("测试数据库成功  = " + string(total))
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// Test_db 测试Test_db
// @Tags common
// @Summary 测试Test_db
// @accept application/json
// @Produce application/json
// @Param id query integer true "id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"成功"}"
// @Router /commDb/test_db [get]
func (o *CommonDbApi) Test_db(c *gin.Context) {
	param := c.DefaultQuery("id", "1")
	fmt.Println("测试数据库 Test_db   =" + param)
	global.LOG.Error("测试数据库 Test_db   =" + param)
	if obj, err := bizSev.GetCmsCatService().GetCmsCat(gconv.Uint(param), ""); err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		global.LOG.Debug("测试数据库成功  = ")
		response.OkWithData(gin.H{"cmsCat": obj}, c)
	}
}
