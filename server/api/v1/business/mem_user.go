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

type MemUserApi struct {
}

 

// CreateMemUser 创建MemUser
// @Tags MemUser
// @Summary 创建MemUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.MemUser true "创建MemUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /memUser/createMemUser [post]
func (memUserApi *MemUserApi) CreateMemUser(c *gin.Context) {
	var dataObj business.MemUser
	_ = c.ShouldBindJSON(&dataObj)
	
	if err := gvalid.CheckStruct(c,dataObj, nil); err != nil {
		response.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	}

 
	if err := bizSev.GetMemUserService().CreateMemUser(dataObj); err != nil {
        global.LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMemUser 删除MemUser
// @Tags MemUser
// @Summary 删除MemUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.MemUser true "删除MemUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /memUser/deleteMemUser [delete]
func (memUserApi *MemUserApi) DeleteMemUser(c *gin.Context) {
	var memUser business.MemUser
	_ = c.ShouldBindJSON(&memUser)
	if err := bizSev.GetMemUserService().DeleteMemUser(memUser); err != nil {
        global.LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMemUserByIds 批量删除MemUser
// @Tags MemUser
// @Summary 批量删除MemUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除MemUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /memUser/deleteMemUserByIds [delete]
func (memUserApi *MemUserApi) DeleteMemUserByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := bizSev.GetMemUserService().DeleteMemUserByIds(IDS); err != nil {
        global.LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMemUser 更新MemUser
// @Tags MemUser
// @Summary 更新MemUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.MemUser true "更新MemUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /memUser/updateMemUser [put]
func (memUserApi *MemUserApi) UpdateMemUser(c *gin.Context) {
	var dataObj business.MemUser
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		response.FailWithMessage("更新失败,"+err.FirstString(), c)
		return
	}

	if err := bizSev.GetMemUserService().UpdateMemUser(dataObj); err != nil {
        global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMemUser 用id查询MemUser
// @Tags MemUser
// @Summary 用id查询MemUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.MemUser true "用id查询MemUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /memUser/findMemUser [get]
func (memUserApi *MemUserApi) FindMemUser(c *gin.Context) {
	var memUser business.MemUser
	_ = c.ShouldBindQuery(&memUser) 
	 rememUser,err:= bizSev.GetMemUserService().GetMemUser(memUser.ID,""); 
	 if errors.Is(err, gorm.ErrRecordNotFound) { 
		response.OkWithData(gin.H{"memUser": nil}, c)
	} else if err != nil { 
        global.LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else { 
		response.OkWithData(gin.H{"memUser": rememUser}, c)
	}
}

// GetMemUserList 分页获取MemUser列表
// @Tags MemUser
// @Summary 分页获取MemUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.MemUserSearch true "分页获取MemUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /memUser/getMemUserList [get]
func (memUserApi *MemUserApi) GetMemUserList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")

	var pageInfo bizReq.MemUserSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if  list, total, err := bizSev.GetMemUserService().GetMemUserInfoList(pageInfo,createdAtBetween,""); err != nil {
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
// @Param data body business.MemUser true "快速更新MemUser" 
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router  /memUser/quickEdit [post] 
func (memUserApi *MemUserApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "mem_user" 
	if err := commSev.GetCommonDbService().QuickEdit(quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}


// GetMemUserList 分页导出excel MemUser列表
// @Tags MemUser
// @Summary 分页导出excel MemUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.MemUserSearch true "分页导出excel MemUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /memUser/excelList [get]
func (memUserApi *MemUserApi) ExcelList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")

	var pageInfo bizReq.MemUserSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total,err:= bizSev.GetMemUserService().GetMemUserInfoList(pageInfo,createdAtBetween,""); err != nil {
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


