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

type ImTxMsgApi struct {
}

// CreateImTxMsg 创建ImTxMsg
// @Tags ImTxMsg
// @Summary 创建ImTxMsg
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.ImTxMsg true "创建ImTxMsg"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /imTxMsg/createImTxMsg [post]
func (imTxMsgApi *ImTxMsgApi) CreateImTxMsg(c *gin.Context) {
	var dataObj business.ImTxMsg
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		response.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	}

	if id, err := bizSev.GetImTxMsgSev().Create(dataObj); err != nil {
		global.LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		idResp := &response.IdResp{Id: id}
		response.OkWithData(idResp, c)
	}
}

// DeleteImTxMsg 删除ImTxMsg
// @Tags ImTxMsg
// @Summary 删除ImTxMsg
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.ImTxMsg true "删除ImTxMsg"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /imTxMsg/deleteImTxMsg [delete]
func (imTxMsgApi *ImTxMsgApi) DeleteImTxMsg(c *gin.Context) {
	var imTxMsg business.ImTxMsg
	_ = c.ShouldBindJSON(&imTxMsg)
	if err := bizSev.GetImTxMsgSev().Delete(imTxMsg); err != nil {
		global.LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteImTxMsgByIds 批量删除ImTxMsg
// @Tags ImTxMsg
// @Summary 批量删除ImTxMsg
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ImTxMsg"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /imTxMsg/deleteImTxMsgByIds [delete]
func (imTxMsgApi *ImTxMsgApi) DeleteImTxMsgByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := bizSev.GetImTxMsgSev().DeleteByIds(IDS); err != nil {
		global.LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateImTxMsg 更新ImTxMsg
// @Tags ImTxMsg
// @Summary 更新ImTxMsg
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.ImTxMsg true "更新ImTxMsg"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /imTxMsg/updateImTxMsg [put]
func (imTxMsgApi *ImTxMsgApi) UpdateImTxMsg(c *gin.Context) {
	var dataObj business.ImTxMsg
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		response.FailWithMessage("更新失败,"+err.FirstString(), c)
		return
	}

	if err := bizSev.GetImTxMsgSev().Update(dataObj); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindImTxMsg 用id查询ImTxMsg
// @Tags ImTxMsg
// @Summary 用id查询ImTxMsg
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.ImTxMsg true "用id查询ImTxMsg"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /imTxMsg/findImTxMsg [get]
func (imTxMsgApi *ImTxMsgApi) FindImTxMsg(c *gin.Context) {
	var imTxMsg business.ImTxMsg
	_ = c.ShouldBindQuery(&imTxMsg)
	reimTxMsg, err := bizSev.GetImTxMsgSev().Get(imTxMsg.ID, "")
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.OkWithData(gin.H{"imTxMsg": nil}, c)
	} else if err != nil {
		global.LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"imTxMsg": reimTxMsg}, c)
	}
}

// GetImTxMsgList 分页获取ImTxMsg列表
// @Tags ImTxMsg
// @Summary 分页获取ImTxMsg列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.ImTxMsgSearch true "分页获取ImTxMsg列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /imTxMsg/getImTxMsgList [get]
func (imTxMsgApi *ImTxMsgApi) GetImTxMsgList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")

	var pageInfo bizReq.ImTxMsgSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := bizSev.GetImTxMsgSev().GetList(pageInfo, createdAtBetween, ""); err != nil {
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
// @Param data body business.ImTxMsg true "快速更新ImTxMsg"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router  /imTxMsg/quickEdit [post]
func (imTxMsgApi *ImTxMsgApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "im_tx_msg"
	if err := commSev.GetCommonDbSev().QuickEdit(quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// excelList 分页导出excel ImTxMsg列表
// @Tags ImTxMsg
// @Summary 分页导出excel ImTxMsg列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.ImTxMsgSearch true "分页导出excel ImTxMsg列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /imTxMsg/excelList [get]
func (imTxMsgApi *ImTxMsgApi) ExcelList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")
	var pageInfo bizReq.ImTxMsgSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, _, err := bizSev.GetImTxMsgSev().GetListAll(pageInfo, createdAtBetween, ""); err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		if len(list) == 0 {
			response.FailWithMessage("没有数据", c)
		} else {
			sheetFields := []string{}
			sheetFields = append(sheetFields, "消息类型")
			sheetFields = append(sheetFields, "消息时间")
			sheetFields = append(sheetFields, "发送人")
			sheetFields = append(sheetFields, "接收人")
			sheetFields = append(sheetFields, "时间撮")
			sheetFields = append(sheetFields, "seq")
			sheetFields = append(sheetFields, "random")
			sheetFields = append(sheetFields, "消息类型")
			sheetFields = append(sheetFields, "内容")
			sheetFields = append(sheetFields, "文本")
			sheetFields = append(sheetFields, "媒体文件")
			sheetFields = append(sheetFields, "媒体远程文件")
			sheetFields = append(sheetFields, "媒体下载")
			sheetFields = append(sheetFields, "IP地址")
			sheetFields = append(sheetFields, "平台")
			sheetFields = append(sheetFields, "状态")

			excel := excelize.NewFile()
			excel.SetSheetRow("Sheet1", "A1", &sheetFields)
			for i, v := range list {
				axis := fmt.Sprintf("A%d", i+2)
				var arr = []interface{}{}
				arr = append(arr, v.ChatType)
				arr = append(arr, v.MsgTime)
				arr = append(arr, v.FromAccount)
				arr = append(arr, v.ToAccount)
				arr = append(arr, v.MsgTimestamp)
				arr = append(arr, *v.MsgSeq)
				arr = append(arr, *v.MsgRandom)
				arr = append(arr, v.MsgType)
				arr = append(arr, v.MsgContent)
				arr = append(arr, v.MsgText)
				arr = append(arr, v.MediaList)
				arr = append(arr, v.MediaListTx)
				arr = append(arr, *v.StatusMedia)
				arr = append(arr, v.ClientIp)
				arr = append(arr, v.MsgFromPlatform)
				arr = append(arr, *v.Status)
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
