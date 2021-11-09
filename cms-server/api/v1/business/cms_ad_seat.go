package business

import (
	"errors"

	"github.com/88act/go-cms/server/global"
	"github.com/88act/go-cms/server/model/business"
	businessReq "github.com/88act/go-cms/server/model/business/request"
	"github.com/88act/go-cms/server/model/common/request"
	"github.com/88act/go-cms/server/model/common/response"
	"github.com/88act/go-cms/server/service"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gvalid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CmsAdSeatApi struct {
}

var cmsAdSeatService = service.ServiceGroupApp.BusinessServiceGroup.CmsAdSeatService

// CreateCmsAdSeat 创建CmsAdSeat
// @Tags CmsAdSeat
// @Summary 创建CmsAdSeat
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.CmsAdSeat true "创建CmsAdSeat"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmsAdSeat/createCmsAdSeat [post]
func (cmsAdSeatApi *CmsAdSeatApi) CreateCmsAdSeat(c *gin.Context) {
	var dataObj business.CmsAdSeat
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		response.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	}

	if err := cmsAdSeatService.CreateCmsAdSeat(dataObj); err != nil {
		global.LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCmsAdSeat 删除CmsAdSeat
// @Tags CmsAdSeat
// @Summary 删除CmsAdSeat
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.CmsAdSeat true "删除CmsAdSeat"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmsAdSeat/deleteCmsAdSeat [delete]
func (cmsAdSeatApi *CmsAdSeatApi) DeleteCmsAdSeat(c *gin.Context) {
	var cmsAdSeat business.CmsAdSeat
	_ = c.ShouldBindJSON(&cmsAdSeat)
	if err := cmsAdSeatService.DeleteCmsAdSeat(cmsAdSeat); err != nil {
		global.LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功2222", c)
	}
}

// DeleteCmsAdSeatByIds 批量删除CmsAdSeat
// @Tags CmsAdSeat
// @Summary 批量删除CmsAdSeat
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CmsAdSeat"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /cmsAdSeat/deleteCmsAdSeatByIds [delete]
func (cmsAdSeatApi *CmsAdSeatApi) DeleteCmsAdSeatByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := cmsAdSeatService.DeleteCmsAdSeatByIds(IDS); err != nil {
		global.LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// UpdateCmsAdSeat 更新CmsAdSeat
// @Tags CmsAdSeat
// @Summary 更新CmsAdSeat
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.CmsAdSeat true "更新CmsAdSeat"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cmsAdSeat/updateCmsAdSeat [put]
func (cmsAdSeatApi *CmsAdSeatApi) UpdateCmsAdSeat(c *gin.Context) {
	var dataObj business.CmsAdSeat
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		response.FailWithMessage("更新失败,"+err.FirstString(), c)
		return
	}

	if err := cmsAdSeatService.UpdateCmsAdSeat(dataObj); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCmsAdSeat 用id查询CmsAdSeat
// @Tags CmsAdSeat
// @Summary 用id查询CmsAdSeat
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.CmsAdSeat true "用id查询CmsAdSeat"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cmsAdSeat/findCmsAdSeat [get]
func (cmsAdSeatApi *CmsAdSeatApi) FindCmsAdSeat(c *gin.Context) {
	var cmsAdSeat business.CmsAdSeat
	_ = c.ShouldBindQuery(&cmsAdSeat)
	err, recmsAdSeat := cmsAdSeatService.GetCmsAdSeat(cmsAdSeat.ID, "")
	if errors.Is(err, gorm.ErrRecordNotFound) {
		//fmt.Println("查询不到数据")
		response.OkWithData(gin.H{"cmsAdSeat": nil}, c)
	} else if err != nil {
		global.LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		//response.OkWithData(gin.H{"recmsAdSeat": recmsAdSeat}, c)
		response.OkWithData(gin.H{"cmsAdSeat": recmsAdSeat}, c)
	}
}

// GetCmsAdSeatList 分页获取CmsAdSeat列表
// @Tags CmsAdSeat
// @Summary 分页获取CmsAdSeat列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query businessReq.CmsAdSeatSearch true "分页获取CmsAdSeat列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmsAdSeat/getCmsAdSeatList [get]
func (cmsAdSeatApi *CmsAdSeatApi) GetCmsAdSeatList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")

	var pageInfo businessReq.CmsAdSeatSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := cmsAdSeatService.GetCmsAdSeatInfoList(pageInfo, createdAtBetween, ""); err != nil {
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
// @Param data body business.CmsAdSeat true "快速更新CmsAdSeat"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router  /cmsAdSeat/quickEdit [post]
func (cmsAdSeatApi *CmsAdSeatApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "cms_ad_seat"
	//var_dump.Dump(quickEdit)
	if err := commDbService.QuickEdit(quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}
