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

type BbPiContactsApi struct {
}

 

// CreateBbPiContacts 创建BbPiContacts
// @Tags BbPiContacts
// @Summary 创建BbPiContacts
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.BbPiContacts true "创建BbPiContacts"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiContacts/createBbPiContacts [post]
func (bbPiContactsApi *BbPiContactsApi) CreateBbPiContacts(c *gin.Context) {
	var dataObj business.BbPiContacts
	_ = c.ShouldBindJSON(&dataObj)
	
	if err := gvalid.CheckStruct(c,dataObj, nil); err != nil {
		response.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	}

 
	if id,err := bizSev.GetBbPiContactsSev().Create(dataObj); err != nil {
        global.LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
	    idResp := &response.IdResp{Id: id}
		response.OkWithData(idResp, c)
	}
}

// DeleteBbPiContacts 删除BbPiContacts
// @Tags BbPiContacts
// @Summary 删除BbPiContacts
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.BbPiContacts true "删除BbPiContacts"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bbPiContacts/deleteBbPiContacts [delete]
func (bbPiContactsApi *BbPiContactsApi) DeleteBbPiContacts(c *gin.Context) {
	var bbPiContacts business.BbPiContacts
	_ = c.ShouldBindJSON(&bbPiContacts)
	if err := bizSev.GetBbPiContactsSev().Delete(bbPiContacts); err != nil {
        global.LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBbPiContactsByIds 批量删除BbPiContacts
// @Tags BbPiContacts
// @Summary 批量删除BbPiContacts
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BbPiContacts"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /bbPiContacts/deleteBbPiContactsByIds [delete]
func (bbPiContactsApi *BbPiContactsApi) DeleteBbPiContactsByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := bizSev.GetBbPiContactsSev().DeleteByIds(IDS); err != nil {
        global.LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBbPiContacts 更新BbPiContacts
// @Tags BbPiContacts
// @Summary 更新BbPiContacts
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.BbPiContacts true "更新BbPiContacts"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /bbPiContacts/updateBbPiContacts [put]
func (bbPiContactsApi *BbPiContactsApi) UpdateBbPiContacts(c *gin.Context) {
	var dataObj business.BbPiContacts
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		response.FailWithMessage("更新失败,"+err.FirstString(), c)
		return
	}

	if err := bizSev.GetBbPiContactsSev().Update(dataObj); err != nil {
        global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBbPiContacts 用id查询BbPiContacts
// @Tags BbPiContacts
// @Summary 用id查询BbPiContacts
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.BbPiContacts true "用id查询BbPiContacts"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /bbPiContacts/findBbPiContacts [get]
func (bbPiContactsApi *BbPiContactsApi) FindBbPiContacts(c *gin.Context) {
	var bbPiContacts business.BbPiContacts
	_ = c.ShouldBindQuery(&bbPiContacts) 
	 rebbPiContacts,err:= bizSev.GetBbPiContactsSev().Get(bbPiContacts.ID,""); 
	 if errors.Is(err, gorm.ErrRecordNotFound) { 
		response.OkWithData(gin.H{"bbPiContacts": nil}, c)
	} else if err != nil { 
        global.LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else { 
		response.OkWithData(gin.H{"bbPiContacts": rebbPiContacts}, c)
	}
}

// GetBbPiContactsList 分页获取BbPiContacts列表
// @Tags BbPiContacts
// @Summary 分页获取BbPiContacts列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.BbPiContactsSearch true "分页获取BbPiContacts列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiContacts/getBbPiContactsList [get]
func (bbPiContactsApi *BbPiContactsApi) GetBbPiContactsList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")

	var pageInfo bizReq.BbPiContactsSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if  list, total, err := bizSev.GetBbPiContactsSev().GetList(pageInfo,createdAtBetween,""); err != nil {
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
// @Param data body business.BbPiContacts true "快速更新BbPiContacts" 
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router  /bbPiContacts/quickEdit [post] 
func (bbPiContactsApi *BbPiContactsApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "bb_pi_contacts" 
	if err := commSev.GetCommonDbSev().QuickEdit(quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}


// excelList 分页导出excel BbPiContacts列表
// @Tags BbPiContacts
// @Summary 分页导出excel BbPiContacts列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.BbPiContactsSearch true "分页导出excel BbPiContacts列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiContacts/excelList [get]
func (bbPiContactsApi *BbPiContactsApi) ExcelList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")
	var pageInfo bizReq.BbPiContactsSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list,_,err:= bizSev.GetBbPiContactsSev().GetListAll(pageInfo,createdAtBetween,""); err != nil {
	    global.LOG.Error("获取失败!", zap.Any("err", err))
        response.FailWithMessage("获取失败", c)
    } else {
        if len(list) == 0 {
			response.FailWithMessage("没有数据", c)
		} else { 
			sheetFields := []string{}  
					sheetFields = append(sheetFields, "上传")  
					sheetFields = append(sheetFields, "机构标识")  
					sheetFields = append(sheetFields, "部门")  
					sheetFields = append(sheetFields, "负责人姓名")  
					sheetFields = append(sheetFields, "联系电话")  
					sheetFields = append(sheetFields, "数据生成日期时间")  
					sheetFields = append(sheetFields, "填报日期")  
					sheetFields = append(sheetFields, "撤销标志") 

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
				arr = append(arr, v.Jgdm)
				arr = append(arr, v.Bm)
				arr = append(arr, v.Fzrxm)
				arr = append(arr, v.Lxdh)
				arr = append(arr, v.Sjscsj)
				arr = append(arr, v.Tbrqsj)
				arr = append(arr, v.Cxbz)   
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


 