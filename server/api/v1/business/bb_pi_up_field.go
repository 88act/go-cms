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

type BbPiUpFieldApi struct {
}

 

// CreateBbPiUpField 创建BbPiUpField
// @Tags BbPiUpField
// @Summary 创建BbPiUpField
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.BbPiUpField true "创建BbPiUpField"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiUpField/createBbPiUpField [post]
func (bbPiUpFieldApi *BbPiUpFieldApi) CreateBbPiUpField(c *gin.Context) {
	var dataObj business.BbPiUpField
	_ = c.ShouldBindJSON(&dataObj)
	
	if err := gvalid.CheckStruct(c,dataObj, nil); err != nil {
		response.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	}

 
	if id,err := bizSev.GetBbPiUpFieldSev().Create(dataObj); err != nil {
        global.LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
	    idResp := &response.IdResp{Id: id}
		response.OkWithData(idResp, c)
	}
}

// DeleteBbPiUpField 删除BbPiUpField
// @Tags BbPiUpField
// @Summary 删除BbPiUpField
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.BbPiUpField true "删除BbPiUpField"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bbPiUpField/deleteBbPiUpField [delete]
func (bbPiUpFieldApi *BbPiUpFieldApi) DeleteBbPiUpField(c *gin.Context) {
	var bbPiUpField business.BbPiUpField
	_ = c.ShouldBindJSON(&bbPiUpField)
	if err := bizSev.GetBbPiUpFieldSev().Delete(bbPiUpField); err != nil {
        global.LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBbPiUpFieldByIds 批量删除BbPiUpField
// @Tags BbPiUpField
// @Summary 批量删除BbPiUpField
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BbPiUpField"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /bbPiUpField/deleteBbPiUpFieldByIds [delete]
func (bbPiUpFieldApi *BbPiUpFieldApi) DeleteBbPiUpFieldByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := bizSev.GetBbPiUpFieldSev().DeleteByIds(IDS); err != nil {
        global.LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBbPiUpField 更新BbPiUpField
// @Tags BbPiUpField
// @Summary 更新BbPiUpField
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.BbPiUpField true "更新BbPiUpField"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /bbPiUpField/updateBbPiUpField [put]
func (bbPiUpFieldApi *BbPiUpFieldApi) UpdateBbPiUpField(c *gin.Context) {
	var dataObj business.BbPiUpField
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		response.FailWithMessage("更新失败,"+err.FirstString(), c)
		return
	}

	if err := bizSev.GetBbPiUpFieldSev().Update(dataObj); err != nil {
        global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBbPiUpField 用id查询BbPiUpField
// @Tags BbPiUpField
// @Summary 用id查询BbPiUpField
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.BbPiUpField true "用id查询BbPiUpField"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /bbPiUpField/findBbPiUpField [get]
func (bbPiUpFieldApi *BbPiUpFieldApi) FindBbPiUpField(c *gin.Context) {
	var bbPiUpField business.BbPiUpField
	_ = c.ShouldBindQuery(&bbPiUpField) 
	 rebbPiUpField,err:= bizSev.GetBbPiUpFieldSev().Get(bbPiUpField.ID,""); 
	 if errors.Is(err, gorm.ErrRecordNotFound) { 
		response.OkWithData(gin.H{"bbPiUpField": nil}, c)
	} else if err != nil { 
        global.LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else { 
		response.OkWithData(gin.H{"bbPiUpField": rebbPiUpField}, c)
	}
}

// GetBbPiUpFieldList 分页获取BbPiUpField列表
// @Tags BbPiUpField
// @Summary 分页获取BbPiUpField列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.BbPiUpFieldSearch true "分页获取BbPiUpField列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiUpField/getBbPiUpFieldList [get]
func (bbPiUpFieldApi *BbPiUpFieldApi) GetBbPiUpFieldList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")

	var pageInfo bizReq.BbPiUpFieldSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if  list, total, err := bizSev.GetBbPiUpFieldSev().GetList(pageInfo,createdAtBetween,""); err != nil {
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
// @Param data body business.BbPiUpField true "快速更新BbPiUpField" 
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router  /bbPiUpField/quickEdit [post] 
func (bbPiUpFieldApi *BbPiUpFieldApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "bb_pi_up_field" 
	if err := commSev.GetCommonDbSev().QuickEdit(quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}


// excelList 分页导出excel BbPiUpField列表
// @Tags BbPiUpField
// @Summary 分页导出excel BbPiUpField列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.BbPiUpFieldSearch true "分页导出excel BbPiUpField列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiUpField/excelList [get]
func (bbPiUpFieldApi *BbPiUpFieldApi) ExcelList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")
	var pageInfo bizReq.BbPiUpFieldSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list,_,err:= bizSev.GetBbPiUpFieldSev().GetListAll(pageInfo,createdAtBetween,""); err != nil {
	    global.LOG.Error("获取失败!", zap.Any("err", err))
        response.FailWithMessage("获取失败", c)
    } else {
        if len(list) == 0 {
			response.FailWithMessage("没有数据", c)
		} else { 
			sheetFields := []string{}  
					sheetFields = append(sheetFields, "是否上传")  
					sheetFields = append(sheetFields, "表名")  
					sheetFields = append(sheetFields, "表名cn")  
					sheetFields = append(sheetFields, "字段名")  
					sheetFields = append(sheetFields, "字段名cn")  
					sheetFields = append(sheetFields, "排序") 

			excel := excelize.NewFile()
			excel.SetSheetRow("Sheet1", "A1", &sheetFields)
			for i, v := range list {
				axis := fmt.Sprintf("A%d", i+2)
				var arr = []interface{}{}
				if v.Status == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.Status)
				}
				arr = append(arr, v.TabName)
				arr = append(arr, v.TabNameCn)
				arr = append(arr, v.TabField)
				arr = append(arr, v.TabFieldCn)
				if v.Sort == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.Sort)
				}   
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


 