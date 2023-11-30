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

type MemDepartApi struct {
	commSev.BaseApi
}

// CreateMemDepart 创建MemDepart
func (m *MemDepartApi) CreateMemDepart(c *gin.Context) {
	var dataObj business.MemDepart
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		m.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	}

	dataObj.UserId = utils.GetUserID(c)

	if id, err := bizSev.GetMemDepartSev().Create(c, dataObj); err != nil {
		global.LOG.Error("创建失败!", zap.Any("err", err))
		m.FailWithMessage("创建失败,"+err.Error(), c)
	} else {
		idResp := &response.IdResp{Id: id}
		m.OkWithData(idResp, c)
	}
}

// DeleteMemDepart 删除MemDepart

func (m *MemDepartApi) DeleteMemDepart(c *gin.Context) {
	var memDepart business.MemDepart
	_ = c.ShouldBindJSON(&memDepart)
	if err := bizSev.GetMemDepartSev().Delete(c, memDepart); err != nil {
		global.LOG.Error("删除失败!", zap.Any("err", err))
		m.FailWithMessage("删除失败,"+err.Error(), c)
	} else {
		m.OkWithMessage("删除成功", c)
	}
}

// DeleteMemDepartByIds 批量删除MemDepart
func (m *MemDepartApi) DeleteMemDepartByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := bizSev.GetMemDepartSev().DeleteByIds(c, IDS); err != nil {
		global.LOG.Error("删除失败!", zap.Any("err", err))
		m.FailWithMessage("删除失败,"+err.Error(), c)
	} else {
		m.OkWithMessage("删除成功", c)
	}
}

// UpdateMemDepart 更新MemDepart
func (m *MemDepartApi) UpdateMemDepart(c *gin.Context) {
	var dataObj business.MemDepart
	_ = c.ShouldBindJSON(&dataObj)

	if dataObj.Pid > 0 {
		if obj2, err := bizSev.GetMemDepartSev().Get(c, dataObj.Pid, ""); err == nil {
			if obj2.Pid > 0 {
				if obj3, err := bizSev.GetMemDepartSev().Get(c, obj2.Pid, ""); err == nil {
					if obj3.Pid > 0 {
						m.FailWithMessage("错误，仅支持三级部门", c)
						return
					}
				}
			}
		}
	}

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		m.FailWithMessage("更新失败,"+err.FirstString(), c)
		return
	}

	if err := bizSev.GetMemDepartSev().Update(c, dataObj); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		m.FailWithMessage("更新失败,"+err.Error(), c)
	} else {
		m.OkWithMessage("更新成功", c)
	}
}

// FindMemDepart 用id查询MemDepart
func (m *MemDepartApi) FindMemDepart(c *gin.Context) {
	var memDepart business.MemDepart
	_ = c.ShouldBindQuery(&memDepart)
	data, err := bizSev.GetMemDepartSev().Get(c, memDepart.Id, "")
	if err != nil {
		global.LOG.Error("查询失败!", zap.Any("err", err))
		m.FailWithMessage("查询失败,"+err.Error(), c)
	} else {
		m.OkWithData(data, c)
	}
}

// GetMemDepartList 分页获取MemDepart列表
func (m *MemDepartApi) GetMemDepartList(c *gin.Context) {
	var pageInfo business.MemDepartSearch
	_ = c.ShouldBindQuery(&pageInfo)

	if list, total, err := bizSev.GetMemDepartSev().GetList(c, pageInfo, business.Field_MemDepart_mini); err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		m.FailWithMessage("获取失败", c)
	} else {
		list2 := m.getTreeList(list, 0)
		m.OkWithDetailed(response.PageResult{
			List:     list2,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

func (m *MemDepartApi) getTreeList(list []*business.MemDepart, pid int64) []*business.MemDepart {
	res := make([]*business.MemDepart, 0)
	for _, v := range list {
		if v.Pid == pid {
			v.Children = m.getTreeList(list, v.Id)
			res = append(res, v)
		}
	}
	return res
}

// QuickEdit 快速更新
func (m *MemDepartApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "mem_depart"
	if err := commSev.GetCommonDbSev().QuickEdit(c, quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		m.FailWithMessage("更新失败,"+err.Error(), c)
	} else {
		m.OkWithMessage("更新成功", c)
	}
}

// excelList 分页导出excel MemDepart列表
func (m *MemDepartApi) ExcelList(c *gin.Context) {
	var pageInfo business.MemDepartSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, _, err := bizSev.GetMemDepartSev().GetList(c, pageInfo, ""); err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		m.FailWithMessage("获取失败,"+err.Error(), c)
	} else {
		if len(list) == 0 {
			m.FailWithMessage("没有数据", c)
		} else {
			sheetFields := []string{}

			sheetFields = append(sheetFields, "父id")
			sheetFields = append(sheetFields, "管理员id")
			sheetFields = append(sheetFields, "名称")
			sheetFields = append(sheetFields, "编号")
			sheetFields = append(sheetFields, "描述")
			sheetFields = append(sheetFields, "地址")
			sheetFields = append(sheetFields, "联系电话")
			sheetFields = append(sheetFields, "邮箱")
			sheetFields = append(sheetFields, "排序")
			sheetFields = append(sheetFields, "配图")
			sheetFields = append(sheetFields, "状态")

			excel := excelize.NewFile()
			excel.SetSheetRow("Sheet1", "A1", &sheetFields)
			for i, v := range list {
				axis := fmt.Sprintf("A%d", i+2)
				var arr = []interface{}{}

				arr = append(arr, v.Pid)
				arr = append(arr, v.UserId)
				arr = append(arr, v.Name)
				arr = append(arr, v.Encode)
				arr = append(arr, v.Desc)
				arr = append(arr, v.Address)
				arr = append(arr, v.Phone)
				arr = append(arr, v.Email)
				arr = append(arr, v.Sort)
				arr = append(arr, v.Image)
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
