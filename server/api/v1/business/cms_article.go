package business

import (
 "errors"
	"fmt"
	"go-cms/global"
	"go-cms/model/business"
	bizReq "go-cms/model/business/request"
	"go-cms/model/common/request"
	"go-cms/model/common/response"
	bizSev "go-cms/service/business"
	commSev "go-cms/service/common"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gvalid"
	"github.com/xuri/excelize/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
) 

type CmsArticleApi struct {
}

 

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

 
	if id,err := bizSev.GetCmsArticleSev().Create(dataObj); err != nil {
        global.LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
	    idResp := &response.IdResp{Id: id}
		response.OkWithData(idResp, c)
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
	if err := bizSev.GetCmsArticleSev().Delete(cmsArticle); err != nil {
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
	if err := bizSev.GetCmsArticleSev().DeleteByIds(IDS); err != nil {
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

	if err := bizSev.GetCmsArticleSev().Update(dataObj); err != nil {
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
	 recmsArticle,err:= bizSev.GetCmsArticleSev().Get(cmsArticle.ID,""); 
	 if errors.Is(err, gorm.ErrRecordNotFound) { 
		response.OkWithData(gin.H{"cmsArticle": nil}, c)
	} else if err != nil { 
        global.LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else { 
		response.OkWithData(gin.H{"cmsArticle": recmsArticle}, c)
	}
}

// GetCmsArticleList 分页获取CmsArticle列表
// @Tags CmsArticle
// @Summary 分页获取CmsArticle列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.CmsArticleSearch true "分页获取CmsArticle列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmsArticle/getCmsArticleList [get]
func (cmsArticleApi *CmsArticleApi) GetCmsArticleList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")

	var pageInfo bizReq.CmsArticleSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if  list, total, err := bizSev.GetCmsArticleSev().GetList(pageInfo,createdAtBetween,""); err != nil {
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
	if err := commSev.GetCommonDbSev().QuickEdit(quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}


// excelList 分页导出excel CmsArticle列表
// @Tags CmsArticle
// @Summary 分页导出excel CmsArticle列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.CmsArticleSearch true "分页导出excel CmsArticle列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmsArticle/excelList [get]
func (cmsArticleApi *CmsArticleApi) ExcelList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")
	var pageInfo bizReq.CmsArticleSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list,_,err:= bizSev.GetCmsArticleSev().GetListAll(pageInfo,createdAtBetween,""); err != nil {
	    global.LOG.Error("获取失败!", zap.Any("err", err))
        response.FailWithMessage("获取失败", c)
    } else {
        if len(list) == 0 {
			response.FailWithMessage("没有数据", c)
		} else { 
			sheetFields := []string{}  
					sheetFields = append(sheetFields, "用户id")  
					sheetFields = append(sheetFields, "类别ID")  
					sheetFields = append(sheetFields, "系统类别ID")  
					sheetFields = append(sheetFields, "文章类型")  
					sheetFields = append(sheetFields, "文章标题")  
					sheetFields = append(sheetFields, "文章摘要")  
					sheetFields = append(sheetFields, "文章内容")  
					sheetFields = append(sheetFields, "文章作者")  
					sheetFields = append(sheetFields, "标签列表")  
					sheetFields = append(sheetFields, "作者邮箱")  
					sheetFields = append(sheetFields, "来源")  
					sheetFields = append(sheetFields, "插图")  
					sheetFields = append(sheetFields, "二维码图片")  
					sheetFields = append(sheetFields, "图片列表")  
					sheetFields = append(sheetFields, "媒体列表")  
					sheetFields = append(sheetFields, "链接地址")  
					sheetFields = append(sheetFields, "置顶")  
					sheetFields = append(sheetFields, "热门")  
					sheetFields = append(sheetFields, "总评论")  
					sheetFields = append(sheetFields, "总点击")  
					sheetFields = append(sheetFields, "总评")  
					sheetFields = append(sheetFields, "总星评1")  
					sheetFields = append(sheetFields, "总星评2")  
					sheetFields = append(sheetFields, "总星评3")  
					sheetFields = append(sheetFields, "总星评4")  
					sheetFields = append(sheetFields, "总星评5")  
					sheetFields = append(sheetFields, "排序")  
					sheetFields = append(sheetFields, "父id")  
					sheetFields = append(sheetFields, "章节集合")  
					sheetFields = append(sheetFields, "状态")  
					sheetFields = append(sheetFields, "审核信息") 

			excel := excelize.NewFile()
			excel.SetSheetRow("Sheet1", "A1", &sheetFields)
			for i, v := range list {
				axis := fmt.Sprintf("A%d", i+2)
				var arr = []interface{}{}
				arr = append(arr, *v.UserId)
				arr = append(arr, *v.CatId)
				arr = append(arr, *v.CatIdSys)
				arr = append(arr, *v.MediaType)
				arr = append(arr, v.Name)
				arr = append(arr, v.Sketch)
				arr = append(arr, v.Detail)
				arr = append(arr, v.Author)
				arr = append(arr, v.TagList)
				arr = append(arr, v.AuthorEmail)
				arr = append(arr, v.Referer)
				arr = append(arr, v.Thumb)
				arr = append(arr, v.Qrcode)
				arr = append(arr, v.ImageList)
				arr = append(arr, v.MediaList)
				arr = append(arr, v.Link)
				arr = append(arr, *v.IsTop)
				arr = append(arr, *v.IsHot)
				arr = append(arr, *v.TotalDiscuss)
				arr = append(arr, *v.TotalClick)
				arr = append(arr, *v.TotalStar)
				arr = append(arr, *v.TotalStar1)
				arr = append(arr, *v.TotalStar2)
				arr = append(arr, *v.TotalStar3)
				arr = append(arr, *v.TotalStar4)
				arr = append(arr, *v.TotalStar5)
				arr = append(arr, *v.Sort)
				arr = append(arr, *v.Pid)
				arr = append(arr, v.Chapter)
				arr = append(arr, *v.Status)
				arr = append(arr, v.VerifyMsg)   
			    excel.SetSheetRow("Sheet1", axis,&arr)  
			}
			filename := fmt.Sprintf("ecl%d.xlsx", time.Now().Unix())
			filePath := global.CONFIG.Local.BasePath + global.CONFIG.Local.Path + "/excel/" + filename
			url := global.CONFIG.Local.BaseUrl + global.CONFIG.Local.Path + "/excel/" + filename
			err := excel.SaveAs(filePath)
			if err != nil {
				global.LOG.Error(err.Error())
				response.FailWithMessage("获取失败", c)
			} else {
				resData := map[string]string{"url": url, "filename": filename} 
				response.OkWithData(resData, c)
			} 
		}
    }
}


 