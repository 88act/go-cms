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

type PayPaymentApi struct {
}

// CreatePayPayment 创建PayPayment
// @Tags PayPayment
// @Summary 创建PayPayment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.PayPayment true "创建PayPayment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /payPayment/createPayPayment [post]
func (m *PayPaymentApi) CreatePayPayment(c *gin.Context) {
	var dataObj business.PayPayment
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		response.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	}

	if id, err := bizSev.GetPayPaymentSev().Create(dataObj); err != nil {
		global.LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		idResp := &response.IdResp{Id: id}
		response.OkWithData(idResp, c)
	}
}

// DeletePayPayment 删除PayPayment
// @Tags PayPayment
// @Summary 删除PayPayment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.PayPayment true "删除PayPayment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /payPayment/deletePayPayment [delete]
func (m *PayPaymentApi) DeletePayPayment(c *gin.Context) {
	var payPayment business.PayPayment
	_ = c.ShouldBindJSON(&payPayment)
	if err := bizSev.GetPayPaymentSev().Delete(payPayment); err != nil {
		global.LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeletePayPaymentByIds 批量删除PayPayment
// @Tags PayPayment
// @Summary 批量删除PayPayment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PayPayment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /payPayment/deletePayPaymentByIds [delete]
func (m *PayPaymentApi) DeletePayPaymentByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := bizSev.GetPayPaymentSev().DeleteByIds(IDS); err != nil {
		global.LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdatePayPayment 更新PayPayment
// @Tags PayPayment
// @Summary 更新PayPayment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.PayPayment true "更新PayPayment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /payPayment/updatePayPayment [put]
func (m *PayPaymentApi) UpdatePayPayment(c *gin.Context) {
	var dataObj business.PayPayment
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		response.FailWithMessage("更新失败,"+err.FirstString(), c)
		return
	}

	if err := bizSev.GetPayPaymentSev().Update(dataObj); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindPayPayment 用id查询PayPayment
// @Tags PayPayment
// @Summary 用id查询PayPayment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.PayPayment true "用id查询PayPayment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /payPayment/findPayPayment [get]
func (m *PayPaymentApi) FindPayPayment(c *gin.Context) {
	var payPayment business.PayPayment
	_ = c.ShouldBindQuery(&payPayment)
	repayPayment, err := bizSev.GetPayPaymentSev().Get(payPayment.ID, "")
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.OkWithData(gin.H{"payPayment": nil}, c)
	} else if err != nil {
		global.LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"payPayment": repayPayment}, c)
	}
}

// GetPayPaymentList 分页获取PayPayment列表
// @Tags PayPayment
// @Summary 分页获取PayPayment列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.PayPaymentSearch true "分页获取PayPayment列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /payPayment/getPayPaymentList [get]
func (m *PayPaymentApi) GetPayPaymentList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")

	var pageInfo bizReq.PayPaymentSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := bizSev.GetPayPaymentSev().GetList(pageInfo, createdAtBetween, ""); err != nil {
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
// @Param data body business.PayPayment true "快速更新PayPayment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router  /payPayment/quickEdit [post]
func (m *PayPaymentApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "pay_payment"
	if err := commSev.GetCommonDbSev().QuickEdit(quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// excelList 分页导出excel PayPayment列表
// @Tags PayPayment
// @Summary 分页导出excel PayPayment列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.PayPaymentSearch true "分页导出excel PayPayment列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /payPayment/excelList [get]
func (m *PayPaymentApi) ExcelList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")
	var pageInfo bizReq.PayPaymentSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, _, err := bizSev.GetPayPaymentSev().GetListAll(pageInfo, createdAtBetween, ""); err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		if len(list) == 0 {
			response.FailWithMessage("没有数据", c)
		} else {
			sheetFields := []string{}
			sheetFields = append(sheetFields, "订单号")
			sheetFields = append(sheetFields, "用户id")
			sheetFields = append(sheetFields, "支付方式")
			sheetFields = append(sheetFields, "三方支付类型")
			sheetFields = append(sheetFields, "三方交易状态")
			sheetFields = append(sheetFields, "总金额")
			sheetFields = append(sheetFields, "三方支付单号")
			sheetFields = append(sheetFields, "支付状态")
			sheetFields = append(sheetFields, "业务单号")
			sheetFields = append(sheetFields, "业务类型")
			sheetFields = append(sheetFields, "支付状态")
			sheetFields = append(sheetFields, "支付时间")
			sheetFields = append(sheetFields, "状态")

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
				arr = append(arr, v.PayMode)
				arr = append(arr, v.TradeType)
				arr = append(arr, v.TradeState)
				if v.Price == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.Price)
				}
				arr = append(arr, v.TransactionId)
				arr = append(arr, v.TradeStateDesc)
				arr = append(arr, v.OrderSn)
				arr = append(arr, v.ServiceType)
				if v.StatusPay == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.StatusPay)
				}
				arr = append(arr, v.PayTime)
				if v.Status == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.Status)
				}
				excel.SetSheetRow("Sheet1", axis, &arr)
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
