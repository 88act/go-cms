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

type MemAddressApi struct {
}

 

// CreateMemAddress 创建MemAddress
// @Tags MemAddress
// @Summary 创建MemAddress
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.MemAddress true "创建MemAddress"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /memAddress/createMemAddress [post]
func (memAddressApi *MemAddressApi) CreateMemAddress(c *gin.Context) {
	var dataObj business.MemAddress
	_ = c.ShouldBindJSON(&dataObj)
	
	if err := gvalid.CheckStruct(c,dataObj, nil); err != nil {
		response.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	}

 
	if id,err := bizSev.GetMemAddressSev().Create(dataObj); err != nil {
        global.LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
	    idResp := &response.IdResp{Id: id}
		response.OkWithData(idResp, c)
	}
}

// DeleteMemAddress 删除MemAddress
// @Tags MemAddress
// @Summary 删除MemAddress
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.MemAddress true "删除MemAddress"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /memAddress/deleteMemAddress [delete]
func (memAddressApi *MemAddressApi) DeleteMemAddress(c *gin.Context) {
	var memAddress business.MemAddress
	_ = c.ShouldBindJSON(&memAddress)
	if err := bizSev.GetMemAddressSev().Delete(memAddress); err != nil {
        global.LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMemAddressByIds 批量删除MemAddress
// @Tags MemAddress
// @Summary 批量删除MemAddress
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除MemAddress"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /memAddress/deleteMemAddressByIds [delete]
func (memAddressApi *MemAddressApi) DeleteMemAddressByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := bizSev.GetMemAddressSev().DeleteByIds(IDS); err != nil {
        global.LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMemAddress 更新MemAddress
// @Tags MemAddress
// @Summary 更新MemAddress
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.MemAddress true "更新MemAddress"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /memAddress/updateMemAddress [put]
func (memAddressApi *MemAddressApi) UpdateMemAddress(c *gin.Context) {
	var dataObj business.MemAddress
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		response.FailWithMessage("更新失败,"+err.FirstString(), c)
		return
	}

	if err := bizSev.GetMemAddressSev().Update(dataObj); err != nil {
        global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMemAddress 用id查询MemAddress
// @Tags MemAddress
// @Summary 用id查询MemAddress
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.MemAddress true "用id查询MemAddress"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /memAddress/findMemAddress [get]
func (memAddressApi *MemAddressApi) FindMemAddress(c *gin.Context) {
	var memAddress business.MemAddress
	_ = c.ShouldBindQuery(&memAddress) 
	 rememAddress,err:= bizSev.GetMemAddressSev().Get(memAddress.ID,""); 
	 if errors.Is(err, gorm.ErrRecordNotFound) { 
		response.OkWithData(gin.H{"memAddress": nil}, c)
	} else if err != nil { 
        global.LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else { 
		response.OkWithData(gin.H{"memAddress": rememAddress}, c)
	}
}

// GetMemAddressList 分页获取MemAddress列表
// @Tags MemAddress
// @Summary 分页获取MemAddress列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.MemAddressSearch true "分页获取MemAddress列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /memAddress/getMemAddressList [get]
func (memAddressApi *MemAddressApi) GetMemAddressList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")

	var pageInfo bizReq.MemAddressSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if  list, total, err := bizSev.GetMemAddressSev().GetList(pageInfo,createdAtBetween,""); err != nil {
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
// @Param data body business.MemAddress true "快速更新MemAddress" 
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router  /memAddress/quickEdit [post] 
func (memAddressApi *MemAddressApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "mem_address" 
	if err := commSev.GetCommonDbSev().QuickEdit(quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}


// excelList 分页导出excel MemAddress列表
// @Tags MemAddress
// @Summary 分页导出excel MemAddress列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.MemAddressSearch true "分页导出excel MemAddress列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /memAddress/excelList [get]
func (memAddressApi *MemAddressApi) ExcelList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")
	var pageInfo bizReq.MemAddressSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list,_,err:= bizSev.GetMemAddressSev().GetListAll(pageInfo,createdAtBetween,""); err != nil {
	    global.LOG.Error("获取失败!", zap.Any("err", err))
        response.FailWithMessage("获取失败", c)
    } else {
        if len(list) == 0 {
			response.FailWithMessage("没有数据", c)
		} else { 
			sheetFields := []string{}  
					sheetFields = append(sheetFields, "用户id")  
					sheetFields = append(sheetFields, "收货人")  
					sheetFields = append(sheetFields, "手机")  
					sheetFields = append(sheetFields, "邮箱地址")  
					sheetFields = append(sheetFields, "地区编码")  
					sheetFields = append(sheetFields, "详细地址")  
					sheetFields = append(sheetFields, "邮政编码")  
					sheetFields = append(sheetFields, "默认地址")  
					sheetFields = append(sheetFields, "状态") 

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
				arr = append(arr, v.Realname)
				arr = append(arr, v.Phone)
				arr = append(arr, v.Email)
				arr = append(arr, v.AreaCode)
				arr = append(arr, v.Address)
				arr = append(arr, v.ZipCode)
				if v.IsDefault == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.IsDefault)
				}
				if v.Status == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.Status)
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


 