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

type CmsAdApi struct {
}

 

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
	
	if err := gvalid.CheckStruct(c,dataObj, nil); err != nil {
		response.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	}

 
	if id,err := bizSev.GetCmsAdSev().Create(dataObj); err != nil {
        global.LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
	    idResp := &response.IdResp{Id: id}
		response.OkWithData(idResp, c)
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
	if err := bizSev.GetCmsAdSev().Delete(cmsAd); err != nil {
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
	if err := bizSev.GetCmsAdSev().DeleteByIds(IDS); err != nil {
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

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		response.FailWithMessage("更新失败,"+err.FirstString(), c)
		return
	}

	if err := bizSev.GetCmsAdSev().Update(dataObj); err != nil {
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
	 recmsAd,err:= bizSev.GetCmsAdSev().Get(cmsAd.ID,""); 
	 if errors.Is(err, gorm.ErrRecordNotFound) { 
		response.OkWithData(gin.H{"cmsAd": nil}, c)
	} else if err != nil { 
        global.LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else { 
		response.OkWithData(gin.H{"cmsAd": recmsAd}, c)
	}
}

// GetCmsAdList 分页获取CmsAd列表
// @Tags CmsAd
// @Summary 分页获取CmsAd列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.CmsAdSearch true "分页获取CmsAd列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmsAd/getCmsAdList [get]
func (cmsAdApi *CmsAdApi) GetCmsAdList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")

	var pageInfo bizReq.CmsAdSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if  list, total, err := bizSev.GetCmsAdSev().GetList(pageInfo,createdAtBetween,""); err != nil {
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
	if err := commSev.GetCommonDbSev().QuickEdit(quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}


// excelList 分页导出excel CmsAd列表
// @Tags CmsAd
// @Summary 分页导出excel CmsAd列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.CmsAdSearch true "分页导出excel CmsAd列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmsAd/excelList [get]
func (cmsAdApi *CmsAdApi) ExcelList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")
	var pageInfo bizReq.CmsAdSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list,_,err:= bizSev.GetCmsAdSev().GetListAll(pageInfo,createdAtBetween,""); err != nil {
	    global.LOG.Error("获取失败!", zap.Any("err", err))
        response.FailWithMessage("获取失败", c)
    } else {
        if len(list) == 0 {
			response.FailWithMessage("没有数据", c)
		} else { 
			sheetFields := []string{}  
					sheetFields = append(sheetFields, "广告位置")  
					sheetFields = append(sheetFields, "广告名称")  
					sheetFields = append(sheetFields, "广告类型")  
					sheetFields = append(sheetFields, "是否新窗")  
					sheetFields = append(sheetFields, "广告图片")  
					sheetFields = append(sheetFields, "广告链接")  
					sheetFields = append(sheetFields, "广告内容")  
					sheetFields = append(sheetFields, "过期内容")  
					sheetFields = append(sheetFields, "投放时间")  
					sheetFields = append(sheetFields, "结束时间")  
					sheetFields = append(sheetFields, "点击量")  
					sheetFields = append(sheetFields, "排序")  
					sheetFields = append(sheetFields, "状态") 

			excel := excelize.NewFile()
			excel.SetSheetRow("Sheet1", "A1", &sheetFields)
			for i, v := range list {
				axis := fmt.Sprintf("A%d", i+2)
				var arr = []interface{}{}
				arr = append(arr, *v.SeatId)
				arr = append(arr, v.Name)
				arr = append(arr, *v.MediaType)
				arr = append(arr, *v.BeTarget)
				arr = append(arr, v.Image)
				arr = append(arr, v.Link)
				arr = append(arr, v.Detail)
				arr = append(arr, v.ExpDetail)
				arr = append(arr, v.StartTime)
				arr = append(arr, v.EndTime)
				arr = append(arr, *v.TotalClick)
				arr = append(arr, *v.Sort)
				arr = append(arr, *v.Status)   
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


 