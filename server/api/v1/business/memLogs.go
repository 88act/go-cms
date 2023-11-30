package business

import (
	"fmt"
	"go-cms/global"
	"go-cms/model/business"
	"go-cms/model/common/request"
	"go-cms/model/common/response"
	bizSev "go-cms/service/business"
	commSev "go-cms/service/common"
	"go-cms/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gvalid"
	"github.com/xuri/excelize/v2"
	"go.uber.org/zap"
)

type MemLogsApi struct {
	commSev.BaseApi
}

// CreateMemLogs 创建MemLogs
func (m *MemLogsApi) CreateMemLogs(c *gin.Context) {
	var dataObj business.MemLogs
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		m.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	}

	dataObj.UserId = utils.GetUserID(c)

	if id, err := bizSev.GetMemLogsSev().Create(c, dataObj); err != nil {
		global.LOG.Error("创建失败!", zap.Any("err", err))
		m.FailWithMessage("创建失败,"+err.Error(), c)
	} else {
		idResp := &response.IdResp{Id: id}
		m.OkWithData(idResp, c)
	}
}

// DeleteMemLogs 删除MemLogs

func (m *MemLogsApi) DeleteMemLogs(c *gin.Context) {
	var memLogs business.MemLogs
	_ = c.ShouldBindJSON(&memLogs)
	if err := bizSev.GetMemLogsSev().Delete(c, memLogs); err != nil {
		global.LOG.Error("删除失败!", zap.Any("err", err))
		m.FailWithMessage("删除失败,"+err.Error(), c)
	} else {
		m.OkWithMessage("删除成功", c)
	}
}

// DeleteMemLogsByIds 批量删除MemLogs
func (m *MemLogsApi) DeleteMemLogsByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := bizSev.GetMemLogsSev().DeleteByIds(c, IDS); err != nil {
		global.LOG.Error("删除失败!", zap.Any("err", err))
		m.FailWithMessage("删除失败,"+err.Error(), c)
	} else {
		m.OkWithMessage("删除成功", c)
	}
}

// UpdateMemLogs 更新MemLogs
func (m *MemLogsApi) UpdateMemLogs(c *gin.Context) {
	var dataObj business.MemLogs
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		m.FailWithMessage("更新失败,"+err.FirstString(), c)
		return
	}

	if err := bizSev.GetMemLogsSev().Update(c, dataObj); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		m.FailWithMessage("更新失败,"+err.Error(), c)
	} else {
		m.OkWithMessage("更新成功", c)
	}
}

// FindMemLogs 用id查询MemLogs
func (m *MemLogsApi) FindMemLogs(c *gin.Context) {
	var memLogs business.MemLogs
	_ = c.ShouldBindQuery(&memLogs)
	data, err := bizSev.GetMemLogsSev().Get(c, memLogs.Id, "")
	if err != nil {
		global.LOG.Error("查询失败!", zap.Any("err", err))
		m.FailWithMessage("查询失败,"+err.Error(), c)
	} else {
		m.OkWithData(data, c)
	}
}

// GetMemLogsList 分页获取MemLogs列表
func (m *MemLogsApi) GetMemLogsList(c *gin.Context) {
	var pageInfo business.MemLogsSearch
	_ = c.ShouldBindQuery(&pageInfo)

	if list, total, err := bizSev.GetMemLogsSev().GetList(c, pageInfo, business.Field_MemLogs_mini); err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		m.FailWithMessage("获取失败", c)
	} else {
		m.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// QuickEdit 快速更新
func (m *MemLogsApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "mem_logs"
	if err := commSev.GetCommonDbSev().QuickEdit(c, quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		m.FailWithMessage("更新失败,"+err.Error(), c)
	} else {
		m.OkWithMessage("更新成功", c)
	}
}

// excelList 分页导出excel MemLogs列表
func (m *MemLogsApi) ExcelList(c *gin.Context) {
	var pageInfo business.MemLogsSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, _, err := bizSev.GetMemLogsSev().GetList(c, pageInfo, ""); err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		m.FailWithMessage("获取失败,"+err.Error(), c)
	} else {
		if len(list) == 0 {
			m.FailWithMessage("没有数据", c)
		} else {
			sheetFields := []string{}

			sheetFields = append(sheetFields, "商户id")
			sheetFields = append(sheetFields, "类型")
			sheetFields = append(sheetFields, "备注")
			sheetFields = append(sheetFields, "ip")
			sheetFields = append(sheetFields, "ip城市")
			sheetFields = append(sheetFields, "状态")

			excel := excelize.NewFile()
			excel.SetSheetRow("Sheet1", "A1", &sheetFields)
			for i, v := range list {
				axis := fmt.Sprintf("A%d", i+2)
				var arr = []interface{}{}

				arr = append(arr, v.UserId)
				arr = append(arr, v.Type)
				arr = append(arr, v.Remark)
				arr = append(arr, v.Ip)
				arr = append(arr, v.IpAddr)
				arr = append(arr, v.Status)
				excel.SetSheetRow("Sheet1", axis, &arr)
			}
			filename := fmt.Sprintf("ecl%d.xlsx", time.Now().Unix())
			filePath := global.CONFIG.Local.BasePath + global.CONFIG.Local.Path + "/excel/" + filename
			url := global.CONFIG.Local.BaseUrl + global.CONFIG.Local.Path + "/excel/" + filename
			err := excel.SaveAs(filePath)
			if err != nil {
				global.LOG.Error(err.Error())
				m.FailWithMessage("获取失败", c)
			} else {
				resData := map[string]string{"url": url, "filename": filename}
				m.OkWithData(resData, c)
			}
		}
	}
}
