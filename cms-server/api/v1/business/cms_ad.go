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

type CmsAdApi struct {
}

var cmsAdService = service.ServiceGroupApp.BusinessServiceGroup.CmsAdService

// CreateCmsAd 创建CmsAd
// @Tags CmsAd
// @Summary 创建CmsAd
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.CmsAd true "创建CmsAd"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmsAd/createCmsAd [post]
func (cmsAdApi *CmsAdApi) CreateCmsAd(c *gin.Context) {
	var dataObj business.CmsAd
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		response.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	}

	if err := cmsAdService.CreateCmsAd(dataObj); err != nil {
		global.LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCmsAd 删除CmsAd
// @Tags CmsAd
// @Summary 删除CmsAd
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.CmsAd true "删除CmsAd"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmsAd/deleteCmsAd [delete]
func (cmsAdApi *CmsAdApi) DeleteCmsAd(c *gin.Context) {
	var cmsAd business.CmsAd
	_ = c.ShouldBindJSON(&cmsAd)
	if err := cmsAdService.DeleteCmsAd(cmsAd); err != nil {
		global.LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCmsAdByIds 批量删除CmsAd
// @Tags CmsAd
// @Summary 批量删除CmsAd
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CmsAd"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /cmsAd/deleteCmsAdByIds [delete]
func (cmsAdApi *CmsAdApi) DeleteCmsAdByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := cmsAdService.DeleteCmsAdByIds(IDS); err != nil {
		global.LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCmsAd 更新CmsAd
// @Tags CmsAd
// @Summary 更新CmsAd
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.CmsAd true "更新CmsAd"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cmsAd/updateCmsAd [put]
func (cmsAdApi *CmsAdApi) UpdateCmsAd(c *gin.Context) {
	var dataObj business.CmsAd
	_ = c.ShouldBindJSON(&dataObj)

	if err := cmsAdService.UpdateCmsAd(dataObj); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCmsAd 用id查询CmsAd
// @Tags CmsAd
// @Summary 用id查询CmsAd
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.CmsAd true "用id查询CmsAd"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cmsAd/findCmsAd [get]
func (cmsAdApi *CmsAdApi) FindCmsAd(c *gin.Context) {
	var cmsAd business.CmsAd
	_ = c.ShouldBindQuery(&cmsAd)
	err, recmsAd := cmsAdService.GetCmsAd(cmsAd.ID, "")
	if errors.Is(err, gorm.ErrRecordNotFound) {
		//fmt.Println("查询不到数据")
		response.OkWithData(gin.H{"cmsAd": nil}, c)
	} else if err != nil {
		global.LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		//response.OkWithData(gin.H{"recmsAd": recmsAd}, c)
		response.OkWithData(gin.H{"cmsAd": recmsAd}, c)
	}
}

// GetCmsAdList 分页获取CmsAd列表
// @Tags CmsAd
// @Summary 分页获取CmsAd列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query businessReq.CmsAdSearch true "分页获取CmsAd列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmsAd/getCmsAdList [get]
func (cmsAdApi *CmsAdApi) GetCmsAdList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")

	var pageInfo businessReq.CmsAdSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := cmsAdService.GetCmsAdInfoList(pageInfo, createdAtBetween, ""); err != nil {
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
// @Param data body business.CmsAd true "快速更新CmsAd"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router  /cmsAd/quickEdit [post]
func (cmsAdApi *CmsAdApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "cms_ad"
	//var_dump.Dump(quickEdit)
	if err := commDbService.QuickEdit(quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}
