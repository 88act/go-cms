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

type CmsAdSeatApi struct {
}

 

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
	
	if err := gvalid.CheckStruct(c,dataObj, nil); err != nil {
		response.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	}

 
	if id,err := bizSev.GetCmsAdSeatSev().Create(dataObj); err != nil {
        global.LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
	    idResp := &response.IdResp{Id: id}
		response.OkWithData(idResp, c)
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
	if err := bizSev.GetCmsAdSeatSev().Delete(cmsAdSeat); err != nil {
        global.LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
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
	if err := bizSev.GetCmsAdSeatSev().DeleteByIds(IDS); err != nil {
        global.LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
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

	if err := bizSev.GetCmsAdSeatSev().Update(dataObj); err != nil {
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
	 recmsAdSeat,err:= bizSev.GetCmsAdSeatSev().Get(cmsAdSeat.ID,""); 
	 if errors.Is(err, gorm.ErrRecordNotFound) { 
		response.OkWithData(gin.H{"cmsAdSeat": nil}, c)
	} else if err != nil { 
        global.LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else { 
		response.OkWithData(gin.H{"cmsAdSeat": recmsAdSeat}, c)
	}
}

// GetCmsAdSeatList 分页获取CmsAdSeat列表
// @Tags CmsAdSeat
// @Summary 分页获取CmsAdSeat列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.CmsAdSeatSearch true "分页获取CmsAdSeat列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmsAdSeat/getCmsAdSeatList [get]
func (cmsAdSeatApi *CmsAdSeatApi) GetCmsAdSeatList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")

	var pageInfo bizReq.CmsAdSeatSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if  list, total, err := bizSev.GetCmsAdSeatSev().GetList(pageInfo,createdAtBetween,""); err != nil {
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
	if err := commSev.GetCommonDbSev().QuickEdit(quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}


// excelList 分页导出excel CmsAdSeat列表
// @Tags CmsAdSeat
// @Summary 分页导出excel CmsAdSeat列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.CmsAdSeatSearch true "分页导出excel CmsAdSeat列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmsAdSeat/excelList [get]
func (cmsAdSeatApi *CmsAdSeatApi) ExcelList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")
	var pageInfo bizReq.CmsAdSeatSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list,_,err:= bizSev.GetCmsAdSeatSev().GetListAll(pageInfo,createdAtBetween,""); err != nil {
	    global.LOG.Error("获取失败!", zap.Any("err", err))
        response.FailWithMessage("获取失败", c)
    } else {
        if len(list) == 0 {
			response.FailWithMessage("没有数据", c)
		} else { 
			sheetFields := []string{}  
					sheetFields = append(sheetFields, "位置名称")  
					sheetFields = append(sheetFields, "广告位宽度")  
					sheetFields = append(sheetFields, "广告位高度")  
					sheetFields = append(sheetFields, "广告描述")  
					sheetFields = append(sheetFields, "模板")  
					sheetFields = append(sheetFields, "状态") 

			excel := excelize.NewFile()
			excel.SetSheetRow("Sheet1", "A1", &sheetFields)
			for i, v := range list {
				axis := fmt.Sprintf("A%d", i+2)
				var arr = []interface{}{}
				arr = append(arr, v.Name)
				arr = append(arr, *v.Width)
				arr = append(arr, *v.Height)
				arr = append(arr, v.Desc)
				arr = append(arr, v.Style)
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


 