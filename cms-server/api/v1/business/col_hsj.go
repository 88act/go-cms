package business

import (
	"errors"
	"github.com/88act/go-cms/server/global"
    "github.com/88act/go-cms/server/model/business"
    "github.com/88act/go-cms/server/model/common/request"
    businessReq "github.com/88act/go-cms/server/model/business/request"
    "github.com/88act/go-cms/server/model/common/response"
    "github.com/88act/go-cms/server/service"
    "github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gvalid"
    "go.uber.org/zap" 
	"gorm.io/gorm"
) 

type ColHsjApi struct {
}

var colHsjService = service.ServiceGroupApp.BusinessServiceGroup.ColHsjService
 

// CreateColHsj 创建ColHsj
// @Tags ColHsj
// @Summary 创建ColHsj
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.ColHsj true "创建ColHsj"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /colHsj/createColHsj [post]
func (colHsjApi *ColHsjApi) CreateColHsj(c *gin.Context) {
	var dataObj business.ColHsj
	_ = c.ShouldBindJSON(&dataObj)
	
	if err := gvalid.CheckStruct(c,dataObj, nil); err != nil {
		response.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	}

	if err := colHsjService.CreateColHsj(dataObj); err != nil {
        global.LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteColHsj 删除ColHsj
// @Tags ColHsj
// @Summary 删除ColHsj
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.ColHsj true "删除ColHsj"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /colHsj/deleteColHsj [delete]
func (colHsjApi *ColHsjApi) DeleteColHsj(c *gin.Context) {
	var colHsj business.ColHsj
	_ = c.ShouldBindJSON(&colHsj)
	if err := colHsjService.DeleteColHsj(colHsj); err != nil {
        global.LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteColHsjByIds 批量删除ColHsj
// @Tags ColHsj
// @Summary 批量删除ColHsj
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ColHsj"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /colHsj/deleteColHsjByIds [delete]
func (colHsjApi *ColHsjApi) DeleteColHsjByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := colHsjService.DeleteColHsjByIds(IDS); err != nil {
        global.LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateColHsj 更新ColHsj
// @Tags ColHsj
// @Summary 更新ColHsj
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.ColHsj true "更新ColHsj"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /colHsj/updateColHsj [put]
func (colHsjApi *ColHsjApi) UpdateColHsj(c *gin.Context) {
	var dataObj business.ColHsj
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		response.FailWithMessage("更新失败,"+err.FirstString(), c)
		return
	}

	if err := colHsjService.UpdateColHsj(dataObj); err != nil {
        global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindColHsj 用id查询ColHsj
// @Tags ColHsj
// @Summary 用id查询ColHsj
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.ColHsj true "用id查询ColHsj"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /colHsj/findColHsj [get]
func (colHsjApi *ColHsjApi) FindColHsj(c *gin.Context) {
	var colHsj business.ColHsj
	_ = c.ShouldBindQuery(&colHsj) 
	 err, recolHsj := colHsjService.GetColHsj(colHsj.ID,""); 
	 if errors.Is(err, gorm.ErrRecordNotFound) {
		//fmt.Println("查询不到数据")
		response.OkWithData(gin.H{"colHsj": nil}, c)
	} else if err != nil { 
        global.LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		//response.OkWithData(gin.H{"recolHsj": recolHsj}, c)
		response.OkWithData(gin.H{"colHsj": recolHsj}, c)
	}
}

// GetColHsjList 分页获取ColHsj列表
// @Tags ColHsj
// @Summary 分页获取ColHsj列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query businessReq.ColHsjSearch true "分页获取ColHsj列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /colHsj/getColHsjList [get]
func (colHsjApi *ColHsjApi) GetColHsjList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")

	var pageInfo businessReq.ColHsjSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := colHsjService.GetColHsjInfoList(pageInfo,createdAtBetween,""); err != nil {
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
// @Param data body business.ColHsj true "快速更新ColHsj" 
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router  /colHsj/quickEdit [post] 
func (colHsjApi *ColHsjApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "col_hsj"
	//var_dump.Dump(quickEdit)
	if err := commDbService.QuickEdit(quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}