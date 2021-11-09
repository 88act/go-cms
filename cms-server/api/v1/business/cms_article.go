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

type CmsArticleApi struct {
}

var cmsArticleService = service.ServiceGroupApp.BusinessServiceGroup.CmsArticleService
 

// CreateCmsArticle 创建CmsArticle
// @Tags CmsArticle
// @Summary 创建CmsArticle
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.CmsArticle true "创建CmsArticle"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmsArticle/createCmsArticle [post]
func (cmsArticleApi *CmsArticleApi) CreateCmsArticle(c *gin.Context) {
	var dataObj business.CmsArticle
	_ = c.ShouldBindJSON(&dataObj)
	
	if err := gvalid.CheckStruct(c,dataObj, nil); err != nil {
		response.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	}

	if err := cmsArticleService.CreateCmsArticle(dataObj); err != nil {
        global.LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCmsArticle 删除CmsArticle
// @Tags CmsArticle
// @Summary 删除CmsArticle
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.CmsArticle true "删除CmsArticle"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmsArticle/deleteCmsArticle [delete]
func (cmsArticleApi *CmsArticleApi) DeleteCmsArticle(c *gin.Context) {
	var cmsArticle business.CmsArticle
	_ = c.ShouldBindJSON(&cmsArticle)
	if err := cmsArticleService.DeleteCmsArticle(cmsArticle); err != nil {
        global.LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCmsArticleByIds 批量删除CmsArticle
// @Tags CmsArticle
// @Summary 批量删除CmsArticle
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CmsArticle"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /cmsArticle/deleteCmsArticleByIds [delete]
func (cmsArticleApi *CmsArticleApi) DeleteCmsArticleByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := cmsArticleService.DeleteCmsArticleByIds(IDS); err != nil {
        global.LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCmsArticle 更新CmsArticle
// @Tags CmsArticle
// @Summary 更新CmsArticle
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.CmsArticle true "更新CmsArticle"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cmsArticle/updateCmsArticle [put]
func (cmsArticleApi *CmsArticleApi) UpdateCmsArticle(c *gin.Context) {
	var dataObj business.CmsArticle
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		response.FailWithMessage("更新失败,"+err.FirstString(), c)
		return
	}

	if err := cmsArticleService.UpdateCmsArticle(dataObj); err != nil {
        global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCmsArticle 用id查询CmsArticle
// @Tags CmsArticle
// @Summary 用id查询CmsArticle
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.CmsArticle true "用id查询CmsArticle"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cmsArticle/findCmsArticle [get]
func (cmsArticleApi *CmsArticleApi) FindCmsArticle(c *gin.Context) {
	var cmsArticle business.CmsArticle
	_ = c.ShouldBindQuery(&cmsArticle) 
	 err, recmsArticle := cmsArticleService.GetCmsArticle(cmsArticle.ID,""); 
	 if errors.Is(err, gorm.ErrRecordNotFound) {
		//fmt.Println("查询不到数据")
		response.OkWithData(gin.H{"cmsArticle": nil}, c)
	} else if err != nil { 
        global.LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		//response.OkWithData(gin.H{"recmsArticle": recmsArticle}, c)
		response.OkWithData(gin.H{"cmsArticle": recmsArticle}, c)
	}
}

// GetCmsArticleList 分页获取CmsArticle列表
// @Tags CmsArticle
// @Summary 分页获取CmsArticle列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query businessReq.CmsArticleSearch true "分页获取CmsArticle列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmsArticle/getCmsArticleList [get]
func (cmsArticleApi *CmsArticleApi) GetCmsArticleList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")

	var pageInfo businessReq.CmsArticleSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := cmsArticleService.GetCmsArticleInfoList(pageInfo,createdAtBetween,""); err != nil {
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
// @Param data body business.CmsArticle true "快速更新CmsArticle" 
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router  /cmsArticle/quickEdit [post] 
func (cmsArticleApi *CmsArticleApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "cms_article"
	//var_dump.Dump(quickEdit)
	if err := commDbService.QuickEdit(quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}