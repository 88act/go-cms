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

type OrderOrderApi struct {
}

 

// CreateOrderOrder 创建OrderOrder
// @Tags OrderOrder
// @Summary 创建OrderOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.OrderOrder true "创建OrderOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /orderOrder/createOrderOrder [post]
func (m *OrderOrderApi) CreateOrderOrder(c *gin.Context) {
	var dataObj business.OrderOrder
	_ = c.ShouldBindJSON(&dataObj)
	
	if err := gvalid.CheckStruct(c,dataObj, nil); err != nil {
		response.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	}

 
	if id,err := bizSev.GetOrderOrderSev().Create(dataObj); err != nil {
        global.LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
	    idResp := &response.IdResp{Id: id}
		response.OkWithData(idResp, c)
	}
}

// DeleteOrderOrder 删除OrderOrder
// @Tags OrderOrder
// @Summary 删除OrderOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.OrderOrder true "删除OrderOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /orderOrder/deleteOrderOrder [delete]
func (m *OrderOrderApi) DeleteOrderOrder(c *gin.Context) {
	var orderOrder business.OrderOrder
	_ = c.ShouldBindJSON(&orderOrder)
	if err := bizSev.GetOrderOrderSev().Delete(orderOrder); err != nil {
        global.LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteOrderOrderByIds 批量删除OrderOrder
// @Tags OrderOrder
// @Summary 批量删除OrderOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除OrderOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /orderOrder/deleteOrderOrderByIds [delete]
func (m *OrderOrderApi) DeleteOrderOrderByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := bizSev.GetOrderOrderSev().DeleteByIds(IDS); err != nil {
        global.LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateOrderOrder 更新OrderOrder
// @Tags OrderOrder
// @Summary 更新OrderOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.OrderOrder true "更新OrderOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /orderOrder/updateOrderOrder [put]
func (m *OrderOrderApi) UpdateOrderOrder(c *gin.Context) {
	var dataObj business.OrderOrder
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		response.FailWithMessage("更新失败,"+err.FirstString(), c)
		return
	}

	if err := bizSev.GetOrderOrderSev().Update(dataObj); err != nil {
        global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindOrderOrder 用id查询OrderOrder
// @Tags OrderOrder
// @Summary 用id查询OrderOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.OrderOrder true "用id查询OrderOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /orderOrder/findOrderOrder [get]
func (m *OrderOrderApi) FindOrderOrder(c *gin.Context) {
	var orderOrder business.OrderOrder
	_ = c.ShouldBindQuery(&orderOrder) 
	 reorderOrder,err:= bizSev.GetOrderOrderSev().Get(orderOrder.ID,""); 
	 if errors.Is(err, gorm.ErrRecordNotFound) { 
		response.OkWithData(gin.H{"orderOrder": nil}, c)
	} else if err != nil { 
        global.LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else { 
		response.OkWithData(gin.H{"orderOrder": reorderOrder}, c)
	}
}

// GetOrderOrderList 分页获取OrderOrder列表
// @Tags OrderOrder
// @Summary 分页获取OrderOrder列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.OrderOrderSearch true "分页获取OrderOrder列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /orderOrder/getOrderOrderList [get]
func (m *OrderOrderApi) GetOrderOrderList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")

	var pageInfo bizReq.OrderOrderSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if  list, total, err := bizSev.GetOrderOrderSev().GetList(pageInfo,createdAtBetween,""); err != nil {
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
// @Param data body business.OrderOrder true "快速更新OrderOrder" 
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router  /orderOrder/quickEdit [post] 
func (m *OrderOrderApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "order_order" 
	if err := commSev.GetCommonDbSev().QuickEdit(quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}


// excelList 分页导出excel OrderOrder列表
// @Tags OrderOrder
// @Summary 分页导出excel OrderOrder列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.OrderOrderSearch true "分页导出excel OrderOrder列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /orderOrder/excelList [get]
func (m *OrderOrderApi) ExcelList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")
	var pageInfo bizReq.OrderOrderSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list,_,err:= bizSev.GetOrderOrderSev().GetListAll(pageInfo,createdAtBetween,""); err != nil {
	    global.LOG.Error("获取失败!", zap.Any("err", err))
        response.FailWithMessage("获取失败", c)
    } else {
        if len(list) == 0 {
			response.FailWithMessage("没有数据", c)
		} else { 
			sheetFields := []string{}  
					sheetFields = append(sheetFields, "订单号")  
					sheetFields = append(sheetFields, "用户id")  
					sheetFields = append(sheetFields, "订单类型")  
					sheetFields = append(sheetFields, "对象id")  
					sheetFields = append(sheetFields, "实付价")  
					sheetFields = append(sheetFields, "原价")  
					sheetFields = append(sheetFields, "优惠券ID")  
					sheetFields = append(sheetFields, "备注")  
					sheetFields = append(sheetFields, "核销码")  
					sheetFields = append(sheetFields, "支付状态")  
					sheetFields = append(sheetFields, "订单状态")  
					sheetFields = append(sheetFields, "记录状态") 

			excel := excelize.NewFile()
			excel.SetSheetRow("Sheet1", "A1", &sheetFields)
			for i, v := range list {
				axis := fmt.Sprintf("A%d", i+2)
				var arr = []interface{}{}
				arr = append(arr, v.Sn)
				if v.UserId == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.UserId)
				}
				if v.OrderType == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.OrderType)
				}
				if v.ObjId == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.ObjId)
				}
				if v.Price == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.Price)
				}
				if v.PriceSrc == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.PriceSrc)
				}
				if v.CouponId == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.CouponId)
				}
				arr = append(arr, v.Remark)
				arr = append(arr, v.OrderCode)
				if v.StatusPay == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.StatusPay)
				}
				if v.StatusOrder == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.StatusOrder)
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


 