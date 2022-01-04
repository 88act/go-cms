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

type MemUserLogApi struct {
}

 

// CreateMemUserLog 创建MemUserLog
// @Tags MemUserLog
// @Summary 创建MemUserLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.MemUserLog true "创建MemUserLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /memUserLog/createMemUserLog [post]
func (memUserLogApi *MemUserLogApi) CreateMemUserLog(c *gin.Context) {
	var dataObj business.MemUserLog
	_ = c.ShouldBindJSON(&dataObj)
	
	if err := gvalid.CheckStruct(c,dataObj, nil); err != nil {
		response.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	}

 
	if id,err := bizSev.GetMemUserLogSev().Create(dataObj); err != nil {
        global.LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
	    idResp := &response.IdResp{Id: id}
		response.OkWithData(idResp, c)
	}
}

// DeleteMemUserLog 删除MemUserLog
// @Tags MemUserLog
// @Summary 删除MemUserLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.MemUserLog true "删除MemUserLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /memUserLog/deleteMemUserLog [delete]
func (memUserLogApi *MemUserLogApi) DeleteMemUserLog(c *gin.Context) {
	var memUserLog business.MemUserLog
	_ = c.ShouldBindJSON(&memUserLog)
	if err := bizSev.GetMemUserLogSev().Delete(memUserLog); err != nil {
        global.LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMemUserLogByIds 批量删除MemUserLog
// @Tags MemUserLog
// @Summary 批量删除MemUserLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除MemUserLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /memUserLog/deleteMemUserLogByIds [delete]
func (memUserLogApi *MemUserLogApi) DeleteMemUserLogByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := bizSev.GetMemUserLogSev().DeleteByIds(IDS); err != nil {
        global.LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMemUserLog 更新MemUserLog
// @Tags MemUserLog
// @Summary 更新MemUserLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.MemUserLog true "更新MemUserLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /memUserLog/updateMemUserLog [put]
func (memUserLogApi *MemUserLogApi) UpdateMemUserLog(c *gin.Context) {
	var dataObj business.MemUserLog
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		response.FailWithMessage("更新失败,"+err.FirstString(), c)
		return
	}

	if err := bizSev.GetMemUserLogSev().Update(dataObj); err != nil {
        global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMemUserLog 用id查询MemUserLog
// @Tags MemUserLog
// @Summary 用id查询MemUserLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.MemUserLog true "用id查询MemUserLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /memUserLog/findMemUserLog [get]
func (memUserLogApi *MemUserLogApi) FindMemUserLog(c *gin.Context) {
	var memUserLog business.MemUserLog
	_ = c.ShouldBindQuery(&memUserLog) 
	 rememUserLog,err:= bizSev.GetMemUserLogSev().Get(memUserLog.ID,""); 
	 if errors.Is(err, gorm.ErrRecordNotFound) { 
		response.OkWithData(gin.H{"memUserLog": nil}, c)
	} else if err != nil { 
        global.LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else { 
		response.OkWithData(gin.H{"memUserLog": rememUserLog}, c)
	}
}

// GetMemUserLogList 分页获取MemUserLog列表
// @Tags MemUserLog
// @Summary 分页获取MemUserLog列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.MemUserLogSearch true "分页获取MemUserLog列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /memUserLog/getMemUserLogList [get]
func (memUserLogApi *MemUserLogApi) GetMemUserLogList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")

	var pageInfo bizReq.MemUserLogSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if  list, total, err := bizSev.GetMemUserLogSev().GetList(pageInfo,createdAtBetween,""); err != nil {
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
// @Param data body business.MemUserLog true "快速更新MemUserLog" 
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router  /memUserLog/quickEdit [post] 
func (memUserLogApi *MemUserLogApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "mem_user_log" 
	if err := commSev.GetCommonDbSev().QuickEdit(quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}


// excelList 分页导出excel MemUserLog列表
// @Tags MemUserLog
// @Summary 分页导出excel MemUserLog列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.MemUserLogSearch true "分页导出excel MemUserLog列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /memUserLog/excelList [get]
func (memUserLogApi *MemUserLogApi) ExcelList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")
	var pageInfo bizReq.MemUserLogSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list,_,err:= bizSev.GetMemUserLogSev().GetListAll(pageInfo,createdAtBetween,""); err != nil {
	    global.LOG.Error("获取失败!", zap.Any("err", err))
        response.FailWithMessage("获取失败", c)
    } else {
        if len(list) == 0 {
			response.FailWithMessage("没有数据", c)
		} else { 
			sheetFields := []string{}  
					sheetFields = append(sheetFields, "用户id")  
					sheetFields = append(sheetFields, "类型")  
					sheetFields = append(sheetFields, "状态")  
					sheetFields = append(sheetFields, "ip") 

			excel := excelize.NewFile()
			excel.SetSheetRow("Sheet1", "A1", &sheetFields)
			for i, v := range list {
				axis := fmt.Sprintf("A%d", i+2)
				var arr = []interface{}{}
				if v.UserId == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.UserId)
				}
				if v.Type == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.Type)
				}
				if v.Status == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.Status)
				}
				arr = append(arr, v.Ip)   
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


 