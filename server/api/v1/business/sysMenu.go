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

type SysMenuApi struct {
	commSev.BaseApi
}

// CreateSysMenu 创建SysMenu
func (m *SysMenuApi) CreateSysMenu(c *gin.Context) {
	var dataObj business.SysMenu
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		m.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	}

	//dataObj.UserId = utils.GetUserID(c)
	if id, err := bizSev.GetSysMenuSev().Create(c, dataObj); err != nil {
		global.LOG.Error("创建失败!", zap.Any("err", err))
		m.FailWithMessage("创建失败,"+err.Error(), c)
	} else {
		idResp := &response.IdResp{Id: id}
		m.OkWithData(idResp, c)
	}
}

// DeleteSysMenu 删除SysMenu

func (m *SysMenuApi) DeleteSysMenu(c *gin.Context) {
	var sysMenu business.SysMenu
	_ = c.ShouldBindJSON(&sysMenu)
	if err := bizSev.GetSysMenuSev().Delete(c, sysMenu); err != nil {
		global.LOG.Error("删除失败!", zap.Any("err", err))
		m.FailWithMessage("删除失败,"+err.Error(), c)
	} else {
		m.OkWithMessage("删除成功", c)
	}
}

// DeleteSysMenuByIds 批量删除SysMenu
func (m *SysMenuApi) DeleteSysMenuByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := bizSev.GetSysMenuSev().DeleteByIds(c, IDS); err != nil {
		global.LOG.Error("删除失败!", zap.Any("err", err))
		m.FailWithMessage("删除失败,"+err.Error(), c)
	} else {
		m.OkWithMessage("删除成功", c)
	}
}

// UpdateSysMenu 更新SysMenu
func (m *SysMenuApi) UpdateSysMenu(c *gin.Context) {
	var dataObj business.SysMenu
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		m.FailWithMessage("更新失败,"+err.FirstString(), c)
		return
	}

	if dataObj.Pid > 0 {
		if obj2, err := bizSev.GetSysMenuSev().Get(c, dataObj.Pid, ""); err == nil {
			if obj2.Pid > 0 {
				m.FailWithMessage("错误，仅支持二级菜单", c)
				c.Abort()
				return
			}
		}
	}

	if err := bizSev.GetSysMenuSev().Update(c, dataObj); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		m.FailWithMessage("更新失败,"+err.Error(), c)
	} else {
		m.OkWithMessage("更新成功", c)
	}
}

// FindSysMenu 用id查询SysMenu
func (m *SysMenuApi) FindSysMenu(c *gin.Context) {
	var sysMenu business.SysMenu
	_ = c.ShouldBindQuery(&sysMenu)
	data, err := bizSev.GetSysMenuSev().Get(c, sysMenu.Id, "")
	if err != nil {
		global.LOG.Error("查询失败!", zap.Any("err", err))
		m.FailWithMessage("查询失败,"+err.Error(), c)
	} else {
		m.OkWithData(data, c)
	}
}

// GetSysMenuList 分页获取SysMenu列表
func (m *SysMenuApi) GetSysMenuList(c *gin.Context) {
	var pageInfo business.SysMenuSearch
	_ = c.ShouldBindQuery(&pageInfo)
	list, total, err := bizSev.GetSysMenuSev().GetList(c, pageInfo, business.Field_SysMenu_mini)

	list2 := m.getTreeList(list, 0)
	if err != nil {
		m.FailWithMessage("获取失败", c)
	} else {
		m.OkWithDetailed(response.PageResult{
			List:     list2,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

func (m *SysMenuApi) getTreeList(list []*business.SysMenu, pid int64) []*business.SysMenu {
	res := make([]*business.SysMenu, 0)
	for _, v := range list {
		if v.Pid == pid {
			v.Children = m.getTreeList(list, v.Id)
			res = append(res, v)
		}
	}
	return res
}

// QuickEdit 快速更新
func (m *SysMenuApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "sys_menu"
	if err := commSev.GetCommonDbSev().QuickEdit(c, quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		m.FailWithMessage("更新失败,"+err.Error(), c)
	} else {
		m.OkWithMessage("更新成功", c)
	}
}

// excelList 分页导出excel SysMenu列表
func (m *SysMenuApi) ExcelList(c *gin.Context) {
	var pageInfo business.SysMenuSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, _, err := bizSev.GetSysMenuSev().GetList(c, pageInfo, ""); err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		m.FailWithMessage("获取失败,"+err.Error(), c)
	} else {
		if len(list) == 0 {
			m.FailWithMessage("没有数据", c)
		} else {
			sheetFields := []string{}

			sheetFields = append(sheetFields, "父菜单ID")
			sheetFields = append(sheetFields, "路由name")
			sheetFields = append(sheetFields, "路由path")
			sheetFields = append(sheetFields, "标题")
			sheetFields = append(sheetFields, "等级")
			sheetFields = append(sheetFields, "是否隐藏")
			sheetFields = append(sheetFields, "前端文件路径")
			sheetFields = append(sheetFields, "参数")
			sheetFields = append(sheetFields, "排序")
			sheetFields = append(sheetFields, "保持连接")
			sheetFields = append(sheetFields, "图标")
			sheetFields = append(sheetFields, "关闭tab")
			sheetFields = append(sheetFields, "状态")

			excel := excelize.NewFile()
			excel.SetSheetRow("Sheet1", "A1", &sheetFields)
			for i, v := range list {
				axis := fmt.Sprintf("A%d", i+2)
				var arr = []interface{}{}

				arr = append(arr, v.Pid)
				arr = append(arr, v.Name)
				arr = append(arr, v.Path)
				arr = append(arr, v.Title)
				arr = append(arr, v.Level)
				arr = append(arr, v.Component)
				arr = append(arr, v.Param)
				arr = append(arr, v.Sort)
				arr = append(arr, v.KeepAlive)
				arr = append(arr, v.Icon)

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
