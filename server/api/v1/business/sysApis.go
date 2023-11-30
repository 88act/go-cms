package business

import (
	"fmt"
	"go-cms/global"
	"go-cms/model/business"
	"go-cms/model/common/request"
	"go-cms/model/common/response"
	bizSev "go-cms/service/business"
	commSev "go-cms/service/common"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gvalid"
	"github.com/xuri/excelize/v2"
	"go.uber.org/zap"
)

type SysApisApi struct {
	commSev.BaseApi
}

// CreateSysApis 创建SysApis
func (m *SysApisApi) CreateSysApis(c *gin.Context) {
	var dataObj business.SysApis
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		m.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	}

	//dataObj.UserId = utils.GetUserID(c)
	if id, err := bizSev.GetSysApisSev().Create(c, dataObj); err != nil {
		global.LOG.Error("创建失败!", zap.Any("err", err))
		m.FailWithMessage("创建失败,"+err.Error(), c)
	} else {
		idResp := &response.IdResp{Id: id}
		m.OkWithData(idResp, c)
	}
}

// DeleteSysApis 删除SysApis

func (m *SysApisApi) DeleteSysApis(c *gin.Context) {
	var sysApis business.SysApis
	_ = c.ShouldBindJSON(&sysApis)
	if err := bizSev.GetSysApisSev().Delete(c, sysApis); err != nil {
		global.LOG.Error("删除失败!", zap.Any("err", err))
		m.FailWithMessage("删除失败,"+err.Error(), c)
	} else {
		m.OkWithMessage("删除成功", c)
	}
}

// DeleteSysApisByIds 批量删除SysApis
func (m *SysApisApi) DeleteSysApisByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := bizSev.GetSysApisSev().DeleteByIds(c, IDS); err != nil {
		global.LOG.Error("删除失败!", zap.Any("err", err))
		m.FailWithMessage("删除失败,"+err.Error(), c)
	} else {
		m.OkWithMessage("删除成功", c)
	}
}

// UpdateSysApis 更新SysApis
func (m *SysApisApi) UpdateSysApis(c *gin.Context) {
	var dataObj business.SysApis
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		m.FailWithMessage("更新失败,"+err.FirstString(), c)
		return
	}

	if err := bizSev.GetSysApisSev().Update(c, dataObj); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		m.FailWithMessage("更新失败,"+err.Error(), c)
	} else {
		m.OkWithMessage("更新成功", c)
	}
}

// FindSysApis 用id查询SysApis
func (m *SysApisApi) FindSysApis(c *gin.Context) {
	var sysApis business.SysApis
	_ = c.ShouldBindQuery(&sysApis)
	data, err := bizSev.GetSysApisSev().Get(c, sysApis.Id, "")
	if err != nil {
		global.LOG.Error("查询失败!", zap.Any("err", err))
		m.FailWithMessage("查询失败,"+err.Error(), c)
	} else {
		m.OkWithData(data, c)
	}
}

// GetSysApisList 分页获取SysApis列表
func (m *SysApisApi) GetSysApisList(c *gin.Context) {
	var pageInfo business.SysApisSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := bizSev.GetSysApisSev().GetList(c, pageInfo, business.Field_SysApis_mini); err != nil {
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
func (m *SysApisApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "sys_apis"
	if err := commSev.GetCommonDbSev().QuickEdit(c, quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		m.FailWithMessage("更新失败,"+err.Error(), c)
	} else {
		m.OkWithMessage("更新成功", c)
	}
}

// excelList 分页导出excel SysApis列表
func (m *SysApisApi) ExcelList(c *gin.Context) {
	var pageInfo business.SysApisSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, _, err := bizSev.GetSysApisSev().GetList(c, pageInfo, ""); err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		m.FailWithMessage("获取失败,"+err.Error(), c)
	} else {
		if len(list) == 0 {
			m.FailWithMessage("没有数据", c)
		} else {
			sheetFields := []string{}
			sheetFields = append(sheetFields, "api路径")
			sheetFields = append(sheetFields, "api中文描述")
			sheetFields = append(sheetFields, "所属模块")
			sheetFields = append(sheetFields, "api组")
			sheetFields = append(sheetFields, "方法")
			sheetFields = append(sheetFields, "状态")

			excel := excelize.NewFile()
			excel.SetSheetRow("Sheet1", "A1", &sheetFields)
			for i, v := range list {
				axis := fmt.Sprintf("A%d", i+2)
				var arr = []interface{}{}
				arr = append(arr, v.Path)
				arr = append(arr, v.Desc)
				arr = append(arr, v.Model)
				arr = append(arr, v.ApiGroup)
				arr = append(arr, v.Method)
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
