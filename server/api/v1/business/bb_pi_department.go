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

type BbPiDepartmentApi struct {
}

 

// CreateBbPiDepartment 创建BbPiDepartment
// @Tags BbPiDepartment
// @Summary 创建BbPiDepartment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.BbPiDepartment true "创建BbPiDepartment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiDepartment/createBbPiDepartment [post]
func (bbPiDepartmentApi *BbPiDepartmentApi) CreateBbPiDepartment(c *gin.Context) {
	var dataObj business.BbPiDepartment
	_ = c.ShouldBindJSON(&dataObj)
	
	if err := gvalid.CheckStruct(c,dataObj, nil); err != nil {
		response.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	}

 
	if id,err := bizSev.GetBbPiDepartmentSev().Create(dataObj); err != nil {
        global.LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
	    idResp := &response.IdResp{Id: id}
		response.OkWithData(idResp, c)
	}
}

// DeleteBbPiDepartment 删除BbPiDepartment
// @Tags BbPiDepartment
// @Summary 删除BbPiDepartment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.BbPiDepartment true "删除BbPiDepartment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bbPiDepartment/deleteBbPiDepartment [delete]
func (bbPiDepartmentApi *BbPiDepartmentApi) DeleteBbPiDepartment(c *gin.Context) {
	var bbPiDepartment business.BbPiDepartment
	_ = c.ShouldBindJSON(&bbPiDepartment)
	if err := bizSev.GetBbPiDepartmentSev().Delete(bbPiDepartment); err != nil {
        global.LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBbPiDepartmentByIds 批量删除BbPiDepartment
// @Tags BbPiDepartment
// @Summary 批量删除BbPiDepartment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BbPiDepartment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /bbPiDepartment/deleteBbPiDepartmentByIds [delete]
func (bbPiDepartmentApi *BbPiDepartmentApi) DeleteBbPiDepartmentByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := bizSev.GetBbPiDepartmentSev().DeleteByIds(IDS); err != nil {
        global.LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBbPiDepartment 更新BbPiDepartment
// @Tags BbPiDepartment
// @Summary 更新BbPiDepartment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.BbPiDepartment true "更新BbPiDepartment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /bbPiDepartment/updateBbPiDepartment [put]
func (bbPiDepartmentApi *BbPiDepartmentApi) UpdateBbPiDepartment(c *gin.Context) {
	var dataObj business.BbPiDepartment
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		response.FailWithMessage("更新失败,"+err.FirstString(), c)
		return
	}

	if err := bizSev.GetBbPiDepartmentSev().Update(dataObj); err != nil {
        global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBbPiDepartment 用id查询BbPiDepartment
// @Tags BbPiDepartment
// @Summary 用id查询BbPiDepartment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.BbPiDepartment true "用id查询BbPiDepartment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /bbPiDepartment/findBbPiDepartment [get]
func (bbPiDepartmentApi *BbPiDepartmentApi) FindBbPiDepartment(c *gin.Context) {
	var bbPiDepartment business.BbPiDepartment
	_ = c.ShouldBindQuery(&bbPiDepartment) 
	 rebbPiDepartment,err:= bizSev.GetBbPiDepartmentSev().Get(bbPiDepartment.ID,""); 
	 if errors.Is(err, gorm.ErrRecordNotFound) { 
		response.OkWithData(gin.H{"bbPiDepartment": nil}, c)
	} else if err != nil { 
        global.LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else { 
		response.OkWithData(gin.H{"bbPiDepartment": rebbPiDepartment}, c)
	}
}

// GetBbPiDepartmentList 分页获取BbPiDepartment列表
// @Tags BbPiDepartment
// @Summary 分页获取BbPiDepartment列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.BbPiDepartmentSearch true "分页获取BbPiDepartment列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiDepartment/getBbPiDepartmentList [get]
func (bbPiDepartmentApi *BbPiDepartmentApi) GetBbPiDepartmentList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")

	var pageInfo bizReq.BbPiDepartmentSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if  list, total, err := bizSev.GetBbPiDepartmentSev().GetList(pageInfo,createdAtBetween,""); err != nil {
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
// @Param data body business.BbPiDepartment true "快速更新BbPiDepartment" 
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router  /bbPiDepartment/quickEdit [post] 
func (bbPiDepartmentApi *BbPiDepartmentApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "bb_pi_department" 
	if err := commSev.GetCommonDbSev().QuickEdit(quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}


// excelList 分页导出excel BbPiDepartment列表
// @Tags BbPiDepartment
// @Summary 分页导出excel BbPiDepartment列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.BbPiDepartmentSearch true "分页导出excel BbPiDepartment列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiDepartment/excelList [get]
func (bbPiDepartmentApi *BbPiDepartmentApi) ExcelList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")
	var pageInfo bizReq.BbPiDepartmentSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list,_,err:= bizSev.GetBbPiDepartmentSev().GetListAll(pageInfo,createdAtBetween,""); err != nil {
	    global.LOG.Error("获取失败!", zap.Any("err", err))
        response.FailWithMessage("获取失败", c)
    } else {
        if len(list) == 0 {
			response.FailWithMessage("没有数据", c)
		} else { 
			sheetFields := []string{}  
					sheetFields = append(sheetFields, "上传")  
					sheetFields = append(sheetFields, "机构标识")  
					sheetFields = append(sheetFields, "科室编码")  
					sheetFields = append(sheetFields, "科室名称")  
					sheetFields = append(sheetFields, "标准科室代码")  
					sheetFields = append(sheetFields, "医保局代码")  
					sheetFields = append(sheetFields, "科室简介")  
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
				arr = append(arr, v.Ksbm)
				arr = append(arr, v.Ksmc)
				arr = append(arr, v.Bzksdm)
				arr = append(arr, v.Ybjdm)
				arr = append(arr, v.Ksjs)
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


 