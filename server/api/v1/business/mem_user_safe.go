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

type MemUserSafeApi struct {
}

 

// CreateMemUserSafe 创建MemUserSafe
// @Tags MemUserSafe
// @Summary 创建MemUserSafe
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.MemUserSafe true "创建MemUserSafe"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /memUserSafe/createMemUserSafe [post]
func (memUserSafeApi *MemUserSafeApi) CreateMemUserSafe(c *gin.Context) {
	var dataObj business.MemUserSafe
	_ = c.ShouldBindJSON(&dataObj)
	
	if err := gvalid.CheckStruct(c,dataObj, nil); err != nil {
		response.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	}

 
	if id,err := bizSev.GetMemUserSafeSev().Create(dataObj); err != nil {
        global.LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
	    idResp := &response.IdResp{Id: id}
		response.OkWithData(idResp, c)
	}
}

// DeleteMemUserSafe 删除MemUserSafe
// @Tags MemUserSafe
// @Summary 删除MemUserSafe
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.MemUserSafe true "删除MemUserSafe"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /memUserSafe/deleteMemUserSafe [delete]
func (memUserSafeApi *MemUserSafeApi) DeleteMemUserSafe(c *gin.Context) {
	var memUserSafe business.MemUserSafe
	_ = c.ShouldBindJSON(&memUserSafe)
	if err := bizSev.GetMemUserSafeSev().Delete(memUserSafe); err != nil {
        global.LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMemUserSafeByIds 批量删除MemUserSafe
// @Tags MemUserSafe
// @Summary 批量删除MemUserSafe
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除MemUserSafe"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /memUserSafe/deleteMemUserSafeByIds [delete]
func (memUserSafeApi *MemUserSafeApi) DeleteMemUserSafeByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := bizSev.GetMemUserSafeSev().DeleteByIds(IDS); err != nil {
        global.LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMemUserSafe 更新MemUserSafe
// @Tags MemUserSafe
// @Summary 更新MemUserSafe
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.MemUserSafe true "更新MemUserSafe"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /memUserSafe/updateMemUserSafe [put]
func (memUserSafeApi *MemUserSafeApi) UpdateMemUserSafe(c *gin.Context) {
	var dataObj business.MemUserSafe
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		response.FailWithMessage("更新失败,"+err.FirstString(), c)
		return
	}

	if err := bizSev.GetMemUserSafeSev().Update(dataObj); err != nil {
        global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMemUserSafe 用id查询MemUserSafe
// @Tags MemUserSafe
// @Summary 用id查询MemUserSafe
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.MemUserSafe true "用id查询MemUserSafe"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /memUserSafe/findMemUserSafe [get]
func (memUserSafeApi *MemUserSafeApi) FindMemUserSafe(c *gin.Context) {
	var memUserSafe business.MemUserSafe
	_ = c.ShouldBindQuery(&memUserSafe) 
	 rememUserSafe,err:= bizSev.GetMemUserSafeSev().Get(memUserSafe.ID,""); 
	 if errors.Is(err, gorm.ErrRecordNotFound) { 
		response.OkWithData(gin.H{"memUserSafe": nil}, c)
	} else if err != nil { 
        global.LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else { 
		response.OkWithData(gin.H{"memUserSafe": rememUserSafe}, c)
	}
}

// GetMemUserSafeList 分页获取MemUserSafe列表
// @Tags MemUserSafe
// @Summary 分页获取MemUserSafe列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.MemUserSafeSearch true "分页获取MemUserSafe列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /memUserSafe/getMemUserSafeList [get]
func (memUserSafeApi *MemUserSafeApi) GetMemUserSafeList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")

	var pageInfo bizReq.MemUserSafeSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if  list, total, err := bizSev.GetMemUserSafeSev().GetList(pageInfo,createdAtBetween,""); err != nil {
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
// @Param data body business.MemUserSafe true "快速更新MemUserSafe" 
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router  /memUserSafe/quickEdit [post] 
func (memUserSafeApi *MemUserSafeApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "mem_user_safe" 
	if err := commSev.GetCommonDbSev().QuickEdit(quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}


// excelList 分页导出excel MemUserSafe列表
// @Tags MemUserSafe
// @Summary 分页导出excel MemUserSafe列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.MemUserSafeSearch true "分页导出excel MemUserSafe列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /memUserSafe/excelList [get]
func (memUserSafeApi *MemUserSafeApi) ExcelList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")
	var pageInfo bizReq.MemUserSafeSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list,_,err:= bizSev.GetMemUserSafeSev().GetListAll(pageInfo,createdAtBetween,""); err != nil {
	    global.LOG.Error("获取失败!", zap.Any("err", err))
        response.FailWithMessage("获取失败", c)
    } else {
        if len(list) == 0 {
			response.FailWithMessage("没有数据", c)
		} else { 
			sheetFields := []string{}  
					sheetFields = append(sheetFields, "用户id")  
					sheetFields = append(sheetFields, "类型")  
					sheetFields = append(sheetFields, "旧值")  
					sheetFields = append(sheetFields, "新值")  
					sheetFields = append(sheetFields, "状态") 

			excel := excelize.NewFile()
			excel.SetSheetRow("Sheet1", "A1", &sheetFields)
			for i, v := range list {
				axis := fmt.Sprintf("A%d", i+2)
				var arr = []interface{}{}
				arr = append(arr, *v.UserId)
				arr = append(arr, *v.Type)
				arr = append(arr, *v.ValOld)
				arr = append(arr, *v.ValNew)
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


 