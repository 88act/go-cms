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

type BbPiDeviceApi struct {
}

 

// CreateBbPiDevice 创建BbPiDevice
// @Tags BbPiDevice
// @Summary 创建BbPiDevice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.BbPiDevice true "创建BbPiDevice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiDevice/createBbPiDevice [post]
func (bbPiDeviceApi *BbPiDeviceApi) CreateBbPiDevice(c *gin.Context) {
	var dataObj business.BbPiDevice
	_ = c.ShouldBindJSON(&dataObj)
	
	if err := gvalid.CheckStruct(c,dataObj, nil); err != nil {
		response.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	}

 
	if id,err := bizSev.GetBbPiDeviceSev().Create(dataObj); err != nil {
        global.LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
	    idResp := &response.IdResp{Id: id}
		response.OkWithData(idResp, c)
	}
}

// DeleteBbPiDevice 删除BbPiDevice
// @Tags BbPiDevice
// @Summary 删除BbPiDevice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.BbPiDevice true "删除BbPiDevice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bbPiDevice/deleteBbPiDevice [delete]
func (bbPiDeviceApi *BbPiDeviceApi) DeleteBbPiDevice(c *gin.Context) {
	var bbPiDevice business.BbPiDevice
	_ = c.ShouldBindJSON(&bbPiDevice)
	if err := bizSev.GetBbPiDeviceSev().Delete(bbPiDevice); err != nil {
        global.LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBbPiDeviceByIds 批量删除BbPiDevice
// @Tags BbPiDevice
// @Summary 批量删除BbPiDevice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BbPiDevice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /bbPiDevice/deleteBbPiDeviceByIds [delete]
func (bbPiDeviceApi *BbPiDeviceApi) DeleteBbPiDeviceByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := bizSev.GetBbPiDeviceSev().DeleteByIds(IDS); err != nil {
        global.LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBbPiDevice 更新BbPiDevice
// @Tags BbPiDevice
// @Summary 更新BbPiDevice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.BbPiDevice true "更新BbPiDevice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /bbPiDevice/updateBbPiDevice [put]
func (bbPiDeviceApi *BbPiDeviceApi) UpdateBbPiDevice(c *gin.Context) {
	var dataObj business.BbPiDevice
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		response.FailWithMessage("更新失败,"+err.FirstString(), c)
		return
	}

	if err := bizSev.GetBbPiDeviceSev().Update(dataObj); err != nil {
        global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBbPiDevice 用id查询BbPiDevice
// @Tags BbPiDevice
// @Summary 用id查询BbPiDevice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.BbPiDevice true "用id查询BbPiDevice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /bbPiDevice/findBbPiDevice [get]
func (bbPiDeviceApi *BbPiDeviceApi) FindBbPiDevice(c *gin.Context) {
	var bbPiDevice business.BbPiDevice
	_ = c.ShouldBindQuery(&bbPiDevice) 
	 rebbPiDevice,err:= bizSev.GetBbPiDeviceSev().Get(bbPiDevice.ID,""); 
	 if errors.Is(err, gorm.ErrRecordNotFound) { 
		response.OkWithData(gin.H{"bbPiDevice": nil}, c)
	} else if err != nil { 
        global.LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else { 
		response.OkWithData(gin.H{"bbPiDevice": rebbPiDevice}, c)
	}
}

// GetBbPiDeviceList 分页获取BbPiDevice列表
// @Tags BbPiDevice
// @Summary 分页获取BbPiDevice列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.BbPiDeviceSearch true "分页获取BbPiDevice列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiDevice/getBbPiDeviceList [get]
func (bbPiDeviceApi *BbPiDeviceApi) GetBbPiDeviceList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")

	var pageInfo bizReq.BbPiDeviceSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if  list, total, err := bizSev.GetBbPiDeviceSev().GetList(pageInfo,createdAtBetween,""); err != nil {
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
// @Param data body business.BbPiDevice true "快速更新BbPiDevice" 
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router  /bbPiDevice/quickEdit [post] 
func (bbPiDeviceApi *BbPiDeviceApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "bb_pi_device" 
	if err := commSev.GetCommonDbSev().QuickEdit(quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}


// excelList 分页导出excel BbPiDevice列表
// @Tags BbPiDevice
// @Summary 分页导出excel BbPiDevice列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.BbPiDeviceSearch true "分页导出excel BbPiDevice列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiDevice/excelList [get]
func (bbPiDeviceApi *BbPiDeviceApi) ExcelList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")
	var pageInfo bizReq.BbPiDeviceSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list,_,err:= bizSev.GetBbPiDeviceSev().GetListAll(pageInfo,createdAtBetween,""); err != nil {
	    global.LOG.Error("获取失败!", zap.Any("err", err))
        response.FailWithMessage("获取失败", c)
    } else {
        if len(list) == 0 {
			response.FailWithMessage("没有数据", c)
		} else { 
			sheetFields := []string{}  
					sheetFields = append(sheetFields, "上传")  
					sheetFields = append(sheetFields, "机构标识")  
					sheetFields = append(sheetFields, "设备代号")  
					sheetFields = append(sheetFields, "设备名称")  
					sheetFields = append(sheetFields, "同批设备台数")  
					sheetFields = append(sheetFields, "产地")  
					sheetFields = append(sheetFields, "生产厂家")  
					sheetFields = append(sheetFields, "设备型号")  
					sheetFields = append(sheetFields, "设备参数")  
					sheetFields = append(sheetFields, "设备类型")  
					sheetFields = append(sheetFields, "购买日期")  
					sheetFields = append(sheetFields, "用途")  
					sheetFields = append(sheetFields, "设备价值金额")  
					sheetFields = append(sheetFields, "购进时新旧情况")  
					sheetFields = append(sheetFields, "理论设计寿命")  
					sheetFields = append(sheetFields, "使用情况")  
					sheetFields = append(sheetFields, "备注")  
					sheetFields = append(sheetFields, "数据生成时间")  
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
				arr = append(arr, v.Sbdh)
				arr = append(arr, v.Sbmc)
				arr = append(arr, v.Tpsbts)
				arr = append(arr, v.Cd)
				arr = append(arr, v.Sccj)
				arr = append(arr, v.Sbxh)
				arr = append(arr, v.Sbcs)
				arr = append(arr, v.Sblx)
				arr = append(arr, v.Gmrq)
				arr = append(arr, v.Yt)
				arr = append(arr, v.Sbjzje)
				arr = append(arr, v.Gjsxqk)
				arr = append(arr, v.Llsjsm)
				arr = append(arr, v.Syqk)
				arr = append(arr, v.Bz)
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


 