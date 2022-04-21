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

type BbPiInstitutionBusinessApi struct {
}

 

// CreateBbPiInstitutionBusiness 创建BbPiInstitutionBusiness
// @Tags BbPiInstitutionBusiness
// @Summary 创建BbPiInstitutionBusiness
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.BbPiInstitutionBusiness true "创建BbPiInstitutionBusiness"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiInstitutionBusiness/createBbPiInstitutionBusiness [post]
func (bbPiInstitutionBusinessApi *BbPiInstitutionBusinessApi) CreateBbPiInstitutionBusiness(c *gin.Context) {
	var dataObj business.BbPiInstitutionBusiness
	_ = c.ShouldBindJSON(&dataObj)
	
	if err := gvalid.CheckStruct(c,dataObj, nil); err != nil {
		response.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	}

 
	if id,err := bizSev.GetBbPiInstitutionBusinessSev().Create(dataObj); err != nil {
        global.LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
	    idResp := &response.IdResp{Id: id}
		response.OkWithData(idResp, c)
	}
}

// DeleteBbPiInstitutionBusiness 删除BbPiInstitutionBusiness
// @Tags BbPiInstitutionBusiness
// @Summary 删除BbPiInstitutionBusiness
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.BbPiInstitutionBusiness true "删除BbPiInstitutionBusiness"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bbPiInstitutionBusiness/deleteBbPiInstitutionBusiness [delete]
func (bbPiInstitutionBusinessApi *BbPiInstitutionBusinessApi) DeleteBbPiInstitutionBusiness(c *gin.Context) {
	var bbPiInstitutionBusiness business.BbPiInstitutionBusiness
	_ = c.ShouldBindJSON(&bbPiInstitutionBusiness)
	if err := bizSev.GetBbPiInstitutionBusinessSev().Delete(bbPiInstitutionBusiness); err != nil {
        global.LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBbPiInstitutionBusinessByIds 批量删除BbPiInstitutionBusiness
// @Tags BbPiInstitutionBusiness
// @Summary 批量删除BbPiInstitutionBusiness
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BbPiInstitutionBusiness"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /bbPiInstitutionBusiness/deleteBbPiInstitutionBusinessByIds [delete]
func (bbPiInstitutionBusinessApi *BbPiInstitutionBusinessApi) DeleteBbPiInstitutionBusinessByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := bizSev.GetBbPiInstitutionBusinessSev().DeleteByIds(IDS); err != nil {
        global.LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBbPiInstitutionBusiness 更新BbPiInstitutionBusiness
// @Tags BbPiInstitutionBusiness
// @Summary 更新BbPiInstitutionBusiness
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.BbPiInstitutionBusiness true "更新BbPiInstitutionBusiness"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /bbPiInstitutionBusiness/updateBbPiInstitutionBusiness [put]
func (bbPiInstitutionBusinessApi *BbPiInstitutionBusinessApi) UpdateBbPiInstitutionBusiness(c *gin.Context) {
	var dataObj business.BbPiInstitutionBusiness
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		response.FailWithMessage("更新失败,"+err.FirstString(), c)
		return
	}

	if err := bizSev.GetBbPiInstitutionBusinessSev().Update(dataObj); err != nil {
        global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBbPiInstitutionBusiness 用id查询BbPiInstitutionBusiness
// @Tags BbPiInstitutionBusiness
// @Summary 用id查询BbPiInstitutionBusiness
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.BbPiInstitutionBusiness true "用id查询BbPiInstitutionBusiness"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /bbPiInstitutionBusiness/findBbPiInstitutionBusiness [get]
func (bbPiInstitutionBusinessApi *BbPiInstitutionBusinessApi) FindBbPiInstitutionBusiness(c *gin.Context) {
	var bbPiInstitutionBusiness business.BbPiInstitutionBusiness
	_ = c.ShouldBindQuery(&bbPiInstitutionBusiness) 
	 rebbPiInstitutionBusiness,err:= bizSev.GetBbPiInstitutionBusinessSev().Get(bbPiInstitutionBusiness.ID,""); 
	 if errors.Is(err, gorm.ErrRecordNotFound) { 
		response.OkWithData(gin.H{"bbPiInstitutionBusiness": nil}, c)
	} else if err != nil { 
        global.LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else { 
		response.OkWithData(gin.H{"bbPiInstitutionBusiness": rebbPiInstitutionBusiness}, c)
	}
}

// GetBbPiInstitutionBusinessList 分页获取BbPiInstitutionBusiness列表
// @Tags BbPiInstitutionBusiness
// @Summary 分页获取BbPiInstitutionBusiness列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.BbPiInstitutionBusinessSearch true "分页获取BbPiInstitutionBusiness列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiInstitutionBusiness/getBbPiInstitutionBusinessList [get]
func (bbPiInstitutionBusinessApi *BbPiInstitutionBusinessApi) GetBbPiInstitutionBusinessList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")

	var pageInfo bizReq.BbPiInstitutionBusinessSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if  list, total, err := bizSev.GetBbPiInstitutionBusinessSev().GetList(pageInfo,createdAtBetween,""); err != nil {
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
// @Param data body business.BbPiInstitutionBusiness true "快速更新BbPiInstitutionBusiness" 
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router  /bbPiInstitutionBusiness/quickEdit [post] 
func (bbPiInstitutionBusinessApi *BbPiInstitutionBusinessApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "bb_pi_institution_business" 
	if err := commSev.GetCommonDbSev().QuickEdit(quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}


// excelList 分页导出excel BbPiInstitutionBusiness列表
// @Tags BbPiInstitutionBusiness
// @Summary 分页导出excel BbPiInstitutionBusiness列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.BbPiInstitutionBusinessSearch true "分页导出excel BbPiInstitutionBusiness列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiInstitutionBusiness/excelList [get]
func (bbPiInstitutionBusinessApi *BbPiInstitutionBusinessApi) ExcelList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")
	var pageInfo bizReq.BbPiInstitutionBusinessSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list,_,err:= bizSev.GetBbPiInstitutionBusinessSev().GetListAll(pageInfo,createdAtBetween,""); err != nil {
	    global.LOG.Error("获取失败!", zap.Any("err", err))
        response.FailWithMessage("获取失败", c)
    } else {
        if len(list) == 0 {
			response.FailWithMessage("没有数据", c)
		} else { 
			sheetFields := []string{}  
					sheetFields = append(sheetFields, "上传")  
					sheetFields = append(sheetFields, "机构标识")  
					sheetFields = append(sheetFields, "月份")  
					sheetFields = append(sheetFields, "派出机构数量")  
					sheetFields = append(sheetFields, "职工总数")  
					sheetFields = append(sheetFields, "客户服务人员总数")  
					sheetFields = append(sheetFields, "日均坐诊医生数")  
					sheetFields = append(sheetFields, "业务用房面积")  
					sheetFields = append(sheetFields, "总收入")  
					sheetFields = append(sheetFields, "总支出")  
					sheetFields = append(sheetFields, "总资产")  
					sheetFields = append(sheetFields, "流动资产")  
					sheetFields = append(sheetFields, "对外投资")  
					sheetFields = append(sheetFields, "固定资产")  
					sheetFields = append(sheetFields, "无形资产及开办费")  
					sheetFields = append(sheetFields, "负债")  
					sheetFields = append(sheetFields, "净资产")  
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
				arr = append(arr, v.Nf)
				arr = append(arr, v.Pcjgsl)
				arr = append(arr, v.Zgzs)
				arr = append(arr, v.Khffryzs)
				arr = append(arr, v.Rjzzyspbs)
				arr = append(arr, v.Ywyfmj)
				arr = append(arr, v.Zsr)
				arr = append(arr, v.Zzc)
				arr = append(arr, v.Zzch)
				arr = append(arr, v.Ldzc)
				arr = append(arr, v.Dwtz)
				arr = append(arr, v.Gdzc)
				arr = append(arr, v.Wxzcjkbf)
				arr = append(arr, v.Fz)
				arr = append(arr, v.Jzc)
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


 