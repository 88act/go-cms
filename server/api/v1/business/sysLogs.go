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

type SysLogsApi struct {
	commSev.BaseApi
}

// CreateSysLogs 创建SysLogs
func (m *SysLogsApi) CreateSysLogs(c *gin.Context) {
	var dataObj business.SysLogs
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		m.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	}

	dataObj.UserId = utils.GetUserID(c)
	if id, err := bizSev.GetSysLogsSev().Create(c, dataObj); err != nil {
		global.LOG.Error("创建失败!", zap.Any("err", err))
		m.FailWithMessage("创建失败,"+err.Error(), c)
	} else {
		idResp := &response.IdResp{Id: id}
		m.OkWithData(idResp, c)
	}
}

// DeleteSysLogs 删除SysLogs

func (m *SysLogsApi) DeleteSysLogs(c *gin.Context) {
	var sysLogs business.SysLogs
	_ = c.ShouldBindJSON(&sysLogs)
	if err := bizSev.GetSysLogsSev().Delete(c, sysLogs); err != nil {
		global.LOG.Error("删除失败!", zap.Any("err", err))
		m.FailWithMessage("删除失败,"+err.Error(), c)
	} else {
		m.OkWithMessage("删除成功", c)
	}
}

// DeleteSysLogsByIds 批量删除SysLogs
func (m *SysLogsApi) DeleteSysLogsByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := bizSev.GetSysLogsSev().DeleteByIds(c, IDS); err != nil {
		global.LOG.Error("删除失败!", zap.Any("err", err))
		m.FailWithMessage("删除失败,"+err.Error(), c)
	} else {
		m.OkWithMessage("删除成功", c)
	}
}

// UpdateSysLogs 更新SysLogs
func (m *SysLogsApi) UpdateSysLogs(c *gin.Context) {
	var dataObj business.SysLogs
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		m.FailWithMessage("更新失败,"+err.FirstString(), c)
		return
	}

	if err := bizSev.GetSysLogsSev().Update(c, dataObj); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		m.FailWithMessage("更新失败,"+err.Error(), c)
	} else {
		m.OkWithMessage("更新成功", c)
	}
}

// FindSysLogs 用id查询SysLogs
func (m *SysLogsApi) FindSysLogs(c *gin.Context) {
	var sysLogs business.SysLogs
	_ = c.ShouldBindQuery(&sysLogs)
	data, err := bizSev.GetSysLogsSev().Get(c, sysLogs.Id, "")
	if err != nil {
		global.LOG.Error("查询失败!", zap.Any("err", err))
		m.FailWithMessage("查询失败,"+err.Error(), c)
	} else {
		m.OkWithData(data, c)
	}
}

// GetSysLogsList 分页获取SysLogs列表
func (m *SysLogsApi) GetSysLogsList(c *gin.Context) {
	var pageInfo business.SysLogsSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := bizSev.GetSysLogsSev().GetList(c, pageInfo, business.Field_SysLogs_mini); err != nil {
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
func (m *SysLogsApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "sys_logs"
	if err := commSev.GetCommonDbSev().QuickEdit(c, quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		m.FailWithMessage("更新失败,"+err.Error(), c)
	} else {
		m.OkWithMessage("更新成功", c)
	}
}

// excelList 分页导出excel SysLogs列表
func (m *SysLogsApi) ExcelList(c *gin.Context) {
	var pageInfo business.SysLogsSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, _, err := bizSev.GetSysLogsSev().GetList(c, pageInfo, ""); err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		m.FailWithMessage("获取失败,"+err.Error(), c)
	} else {
		if len(list) == 0 {
			m.FailWithMessage("没有数据", c)
		} else {
			sheetFields := []string{}

			sheetFields = append(sheetFields, "用户id")
			sheetFields = append(sheetFields, "log类型")
			sheetFields = append(sheetFields, "请求方法")
			sheetFields = append(sheetFields, "请求路径")
			sheetFields = append(sheetFields, "延迟")
			sheetFields = append(sheetFields, "代理")
			sheetFields = append(sheetFields, "错误信息")
			sheetFields = append(sheetFields, "请求Body")
			sheetFields = append(sheetFields, "响应Body")
			sheetFields = append(sheetFields, "请求ip")
			sheetFields = append(sheetFields, "地址")
			sheetFields = append(sheetFields, "请求状态")

			excel := excelize.NewFile()
			excel.SetSheetRow("Sheet1", "A1", &sheetFields)
			for i, v := range list {
				axis := fmt.Sprintf("A%d", i+2)
				var arr = []interface{}{}

				arr = append(arr, v.UserId)
				arr = append(arr, v.LogType)
				arr = append(arr, v.Method)
				arr = append(arr, v.Path)
				arr = append(arr, v.Latency)
				arr = append(arr, v.Agent)
				arr = append(arr, v.ErrorMessage)
				arr = append(arr, v.Body)
				arr = append(arr, v.Resp)
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
