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

type BasicDictApi struct {
	commSev.BaseApi
}

// CreateBasicDict 创建BasicDict
func (m *BasicDictApi) CreateBasicDict(c *gin.Context) {
	var dataObj business.BasicDict
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		m.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	}

	//dataObj.UserId = utils.GetUserID(c)
	if id, err := bizSev.GetBasicDictSev().Create(c, dataObj); err != nil {
		global.LOG.Error("创建失败!", zap.Any("err", err))
		m.FailWithMessage("创建失败,"+err.Error(), c)
	} else {
		idResp := &response.IdResp{Id: id}
		m.OkWithData(idResp, c)
	}
}

// DeleteBasicDict 删除BasicDict

func (m *BasicDictApi) DeleteBasicDict(c *gin.Context) {
	var basicDict business.BasicDict
	_ = c.ShouldBindJSON(&basicDict)
	if err := bizSev.GetBasicDictSev().Delete(c, basicDict); err != nil {
		global.LOG.Error("删除失败!", zap.Any("err", err))
		m.FailWithMessage("删除失败,"+err.Error(), c)
	} else {
		m.OkWithMessage("删除成功", c)
	}
}

// DeleteBasicDictByIds 批量删除BasicDict
func (m *BasicDictApi) DeleteBasicDictByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := bizSev.GetBasicDictSev().DeleteByIds(c, IDS); err != nil {
		global.LOG.Error("删除失败!", zap.Any("err", err))
		m.FailWithMessage("删除失败,"+err.Error(), c)
	} else {
		m.OkWithMessage("删除成功", c)
	}
}

// UpdateBasicDict 更新BasicDict
func (m *BasicDictApi) UpdateBasicDict(c *gin.Context) {
	var dataObj business.BasicDict
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		m.FailWithMessage("更新失败,"+err.FirstString(), c)
		return
	}

	if err := bizSev.GetBasicDictSev().Update(c, dataObj); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		m.FailWithMessage("更新失败,"+err.Error(), c)
	} else {
		m.OkWithMessage("更新成功", c)
	}
}

// FindBasicDict 用id查询BasicDict
func (m *BasicDictApi) FindBasicDict(c *gin.Context) {
	var basicDict business.BasicDict
	_ = c.ShouldBindQuery(&basicDict)
	data, err := bizSev.GetBasicDictSev().Get(c, basicDict.Id, "")
	if err != nil {
		global.LOG.Error("查询失败!", zap.Any("err", err))
		m.FailWithMessage("查询失败,"+err.Error(), c)
	} else {
		m.OkWithData(data, c)
	}
}

// GetBasicDictList 分页获取BasicDict列表
func (m *BasicDictApi) GetBasicDictList(c *gin.Context) {
	var pageInfo business.BasicDictSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := bizSev.GetBasicDictSev().GetList(c, pageInfo, business.Field_BasicDict_mini); err != nil {
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
func (m *BasicDictApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "basic_dict"
	if err := commSev.GetCommonDbSev().QuickEdit(c, quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		m.FailWithMessage("更新失败,"+err.Error(), c)
	} else {
		m.OkWithMessage("更新成功", c)
	}
}

// excelList 分页导出excel BasicDict列表
func (m *BasicDictApi) ExcelList(c *gin.Context) {
	var pageInfo business.BasicDictSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, _, err := bizSev.GetBasicDictSev().GetList(c, pageInfo, ""); err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		m.FailWithMessage("获取失败,"+err.Error(), c)
	} else {
		if len(list) == 0 {
			m.FailWithMessage("没有数据", c)
		} else {
			sheetFields := []string{}
			sheetFields = append(sheetFields, "类别")
			sheetFields = append(sheetFields, "父id")
			sheetFields = append(sheetFields, "模块名")
			sheetFields = append(sheetFields, "名称")
			sheetFields = append(sheetFields, "key")
			sheetFields = append(sheetFields, "图标")
			sheetFields = append(sheetFields, "备注")
			sheetFields = append(sheetFields, "排序")
			sheetFields = append(sheetFields, "状态")
			sheetFields = append(sheetFields, "key字段")

			excel := excelize.NewFile()
			excel.SetSheetRow("Sheet1", "A1", &sheetFields)
			for i, v := range list {
				axis := fmt.Sprintf("A%d", i+2)
				var arr = []interface{}{}
				arr = append(arr, v.Type)
				arr = append(arr, v.Pid)
				arr = append(arr, v.Module)
				arr = append(arr, v.Label)
				arr = append(arr, v.Value)
				arr = append(arr, v.Icon)
				arr = append(arr, v.Remark)
				arr = append(arr, v.Sort)
				arr = append(arr, v.Status)
				arr = append(arr, v.Key)
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
