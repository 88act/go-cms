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

type ImTxFileApi struct {
}

// CreateImTxFile 创建ImTxFile
// @Tags ImTxFile
// @Summary 创建ImTxFile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.ImTxFile true "创建ImTxFile"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /imTxFile/createImTxFile [post]
func (imTxFileApi *ImTxFileApi) CreateImTxFile(c *gin.Context) {
	var dataObj business.ImTxFile
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		response.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	}

	if id, err := bizSev.GetImTxFileSev().Create(dataObj); err != nil {
		global.LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		idResp := &response.IdResp{Id: id}
		response.OkWithData(idResp, c)
	}
}

// DeleteImTxFile 删除ImTxFile
// @Tags ImTxFile
// @Summary 删除ImTxFile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.ImTxFile true "删除ImTxFile"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /imTxFile/deleteImTxFile [delete]
func (imTxFileApi *ImTxFileApi) DeleteImTxFile(c *gin.Context) {
	var imTxFile business.ImTxFile
	_ = c.ShouldBindJSON(&imTxFile)
	if err := bizSev.GetImTxFileSev().Delete(imTxFile); err != nil {
		global.LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteImTxFileByIds 批量删除ImTxFile
// @Tags ImTxFile
// @Summary 批量删除ImTxFile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ImTxFile"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /imTxFile/deleteImTxFileByIds [delete]
func (imTxFileApi *ImTxFileApi) DeleteImTxFileByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := bizSev.GetImTxFileSev().DeleteByIds(IDS); err != nil {
		global.LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateImTxFile 更新ImTxFile
// @Tags ImTxFile
// @Summary 更新ImTxFile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.ImTxFile true "更新ImTxFile"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /imTxFile/updateImTxFile [put]
func (imTxFileApi *ImTxFileApi) UpdateImTxFile(c *gin.Context) {
	var dataObj business.ImTxFile
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		response.FailWithMessage("更新失败,"+err.FirstString(), c)
		return
	}

	if err := bizSev.GetImTxFileSev().Update(dataObj); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindImTxFile 用id查询ImTxFile
// @Tags ImTxFile
// @Summary 用id查询ImTxFile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.ImTxFile true "用id查询ImTxFile"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /imTxFile/findImTxFile [get]
func (imTxFileApi *ImTxFileApi) FindImTxFile(c *gin.Context) {
	var imTxFile business.ImTxFile
	_ = c.ShouldBindQuery(&imTxFile)
	reimTxFile, err := bizSev.GetImTxFileSev().Get(imTxFile.ID, "")
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.OkWithData(gin.H{"imTxFile": nil}, c)
	} else if err != nil {
		global.LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"imTxFile": reimTxFile}, c)
	}
}

// GetImTxFileList 分页获取ImTxFile列表
// @Tags ImTxFile
// @Summary 分页获取ImTxFile列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.ImTxFileSearch true "分页获取ImTxFile列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /imTxFile/getImTxFileList [get]
func (imTxFileApi *ImTxFileApi) GetImTxFileList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")

	var pageInfo bizReq.ImTxFileSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := bizSev.GetImTxFileSev().GetList(pageInfo, createdAtBetween, ""); err != nil {
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
// @Param data body business.ImTxFile true "快速更新ImTxFile"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router  /imTxFile/quickEdit [post]
func (imTxFileApi *ImTxFileApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "im_tx_file"
	if err := commSev.GetCommonDbSev().QuickEdit(quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// excelList 分页导出excel ImTxFile列表
// @Tags ImTxFile
// @Summary 分页导出excel ImTxFile列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.ImTxFileSearch true "分页导出excel ImTxFile列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /imTxFile/excelList [get]
func (imTxFileApi *ImTxFileApi) ExcelList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")
	var pageInfo bizReq.ImTxFileSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, _, err := bizSev.GetImTxFileSev().GetListAll(pageInfo, createdAtBetween, ""); err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		if len(list) == 0 {
			response.FailWithMessage("没有数据", c)
		} else {
			sheetFields := []string{}
			sheetFields = append(sheetFields, "消息id")
			sheetFields = append(sheetFields, "文件类型")
			sheetFields = append(sheetFields, "远程地址")
			sheetFields = append(sheetFields, "本地路径")
			sheetFields = append(sheetFields, "下载状态")

			excel := excelize.NewFile()
			excel.SetSheetRow("Sheet1", "A1", &sheetFields)
			for i, v := range list {
				axis := fmt.Sprintf("A%d", i+2)
				var arr = []interface{}{}
				if v.MsgId == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.MsgId)
				}
				if v.MediaType == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.MediaType)
				}
				arr = append(arr, v.Url)
				arr = append(arr, v.Local)
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
