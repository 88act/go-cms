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

type ImTxMsgFileApi struct {
}

// CreateImTxMsgFile 创建ImTxMsgFile
// @Tags ImTxMsgFile
// @Summary 创建ImTxMsgFile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.ImTxMsgFile true "创建ImTxMsgFile"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /imTxMsgFile/createImTxMsgFile [post]
func (imTxMsgFileApi *ImTxMsgFileApi) CreateImTxMsgFile(c *gin.Context) {
	var dataObj business.ImTxMsgFile
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		response.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	}

	if id, err := bizSev.GetImTxMsgFileSev().Create(dataObj); err != nil {
		global.LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		idResp := &response.IdResp{Id: id}
		response.OkWithData(idResp, c)
	}
}

// DeleteImTxMsgFile 删除ImTxMsgFile
// @Tags ImTxMsgFile
// @Summary 删除ImTxMsgFile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.ImTxMsgFile true "删除ImTxMsgFile"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /imTxMsgFile/deleteImTxMsgFile [delete]
func (imTxMsgFileApi *ImTxMsgFileApi) DeleteImTxMsgFile(c *gin.Context) {
	var imTxMsgFile business.ImTxMsgFile
	_ = c.ShouldBindJSON(&imTxMsgFile)
	if err := bizSev.GetImTxMsgFileSev().Delete(imTxMsgFile); err != nil {
		global.LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteImTxMsgFileByIds 批量删除ImTxMsgFile
// @Tags ImTxMsgFile
// @Summary 批量删除ImTxMsgFile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ImTxMsgFile"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /imTxMsgFile/deleteImTxMsgFileByIds [delete]
func (imTxMsgFileApi *ImTxMsgFileApi) DeleteImTxMsgFileByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := bizSev.GetImTxMsgFileSev().DeleteByIds(IDS); err != nil {
		global.LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateImTxMsgFile 更新ImTxMsgFile
// @Tags ImTxMsgFile
// @Summary 更新ImTxMsgFile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.ImTxMsgFile true "更新ImTxMsgFile"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /imTxMsgFile/updateImTxMsgFile [put]
func (imTxMsgFileApi *ImTxMsgFileApi) UpdateImTxMsgFile(c *gin.Context) {
	var dataObj business.ImTxMsgFile
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		response.FailWithMessage("更新失败,"+err.FirstString(), c)
		return
	}

	if err := bizSev.GetImTxMsgFileSev().Update(dataObj); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindImTxMsgFile 用id查询ImTxMsgFile
// @Tags ImTxMsgFile
// @Summary 用id查询ImTxMsgFile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.ImTxMsgFile true "用id查询ImTxMsgFile"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /imTxMsgFile/findImTxMsgFile [get]
func (imTxMsgFileApi *ImTxMsgFileApi) FindImTxMsgFile(c *gin.Context) {
	var imTxMsgFile business.ImTxMsgFile
	_ = c.ShouldBindQuery(&imTxMsgFile)
	reimTxMsgFile, err := bizSev.GetImTxMsgFileSev().Get(imTxMsgFile.ID, "")
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.OkWithData(gin.H{"imTxMsgFile": nil}, c)
	} else if err != nil {
		global.LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"imTxMsgFile": reimTxMsgFile}, c)
	}
}

// GetImTxMsgFileList 分页获取ImTxMsgFile列表
// @Tags ImTxMsgFile
// @Summary 分页获取ImTxMsgFile列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.ImTxMsgFileSearch true "分页获取ImTxMsgFile列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /imTxMsgFile/getImTxMsgFileList [get]
func (imTxMsgFileApi *ImTxMsgFileApi) GetImTxMsgFileList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")

	var pageInfo bizReq.ImTxMsgFileSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := bizSev.GetImTxMsgFileSev().GetList(pageInfo, createdAtBetween, ""); err != nil {
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
// @Param data body business.ImTxMsgFile true "快速更新ImTxMsgFile"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router  /imTxMsgFile/quickEdit [post]
func (imTxMsgFileApi *ImTxMsgFileApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "im_tx_msg_file"
	if err := commSev.GetCommonDbSev().QuickEdit(quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// excelList 分页导出excel ImTxMsgFile列表
// @Tags ImTxMsgFile
// @Summary 分页导出excel ImTxMsgFile列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.ImTxMsgFileSearch true "分页导出excel ImTxMsgFile列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /imTxMsgFile/excelList [get]
func (imTxMsgFileApi *ImTxMsgFileApi) ExcelList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")
	var pageInfo bizReq.ImTxMsgFileSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, _, err := bizSev.GetImTxMsgFileSev().GetListAll(pageInfo, createdAtBetween, ""); err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		if len(list) == 0 {
			response.FailWithMessage("没有数据", c)
		} else {
			sheetFields := []string{}
			sheetFields = append(sheetFields, "消息类型")
			sheetFields = append(sheetFields, "消息时间")
			sheetFields = append(sheetFields, "下载路径")
			sheetFields = append(sheetFields, "过期时间")
			sheetFields = append(sheetFields, "文件大小")
			sheetFields = append(sheetFields, "文件md5")
			sheetFields = append(sheetFields, "压缩大小")
			sheetFields = append(sheetFields, "压缩md5")
			sheetFields = append(sheetFields, "请求code")
			sheetFields = append(sheetFields, "请求信息")
			sheetFields = append(sheetFields, "请求状态")
			sheetFields = append(sheetFields, "本地路径")
			sheetFields = append(sheetFields, "状态")

			excel := excelize.NewFile()
			excel.SetSheetRow("Sheet1", "A1", &sheetFields)
			for i, v := range list {
				axis := fmt.Sprintf("A%d", i+2)
				var arr = []interface{}{}
				arr = append(arr, v.ChatType)
				arr = append(arr, v.MsgTime)
				arr = append(arr, v.Url)
				arr = append(arr, v.ExpireTime)

				if v.FileSize == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.FileSize)
				}

				arr = append(arr, v.FileMd5)

				if v.GzipSize == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.GzipSize)
				}

				arr = append(arr, v.GzipMd5)
				arr = append(arr, v.ErrorCode)
				arr = append(arr, v.ErrorInfo)
				arr = append(arr, v.ActionStatus)
				arr = append(arr, v.LocalFile)

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
