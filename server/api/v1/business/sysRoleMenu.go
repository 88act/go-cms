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

type SysRoleMenuApi struct {
	commSev.BaseApi
}

// CreateSysRoleMenu 创建SysRoleMenu
func (m *SysRoleMenuApi) CreateSysRoleMenu(c *gin.Context) {
	var dataObj business.SysRoleMenu
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		m.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	}

	//dataObj.UserId = utils.GetUserID(c)
	if id, err := bizSev.GetSysRoleMenuSev().Create(c, dataObj); err != nil {
		global.LOG.Error("创建失败!", zap.Any("err", err))
		m.FailWithMessage("创建失败,"+err.Error(), c)
	} else {
		idResp := &response.IdResp{Id: id}
		m.OkWithData(idResp, c)
	}
}

// DeleteSysRoleMenu 删除SysRoleMenu

func (m *SysRoleMenuApi) DeleteSysRoleMenu(c *gin.Context) {
	var sysRoleMenu business.SysRoleMenu
	_ = c.ShouldBindJSON(&sysRoleMenu)
	if err := bizSev.GetSysRoleMenuSev().Delete(c, sysRoleMenu); err != nil {
		global.LOG.Error("删除失败!", zap.Any("err", err))
		m.FailWithMessage("删除失败,"+err.Error(), c)
	} else {
		m.OkWithMessage("删除成功", c)
	}
}

// DeleteSysRoleMenuByIds 批量删除SysRoleMenu
func (m *SysRoleMenuApi) DeleteSysRoleMenuByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := bizSev.GetSysRoleMenuSev().DeleteByIds(c, IDS); err != nil {
		global.LOG.Error("删除失败!", zap.Any("err", err))
		m.FailWithMessage("删除失败,"+err.Error(), c)
	} else {
		m.OkWithMessage("删除成功", c)
	}
}

// UpdateSysRoleMenu 更新SysRoleMenu
func (m *SysRoleMenuApi) UpdateSysRoleMenu(c *gin.Context) {
	var dataObj business.SysRoleMenu
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		m.FailWithMessage("更新失败,"+err.FirstString(), c)
		return
	}

	if err := bizSev.GetSysRoleMenuSev().Update(c, dataObj); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		m.FailWithMessage("更新失败,"+err.Error(), c)
	} else {
		m.OkWithMessage("更新成功", c)
	}
}

// FindSysRoleMenu 用id查询SysRoleMenu
func (m *SysRoleMenuApi) FindSysRoleMenu(c *gin.Context) {
	var sysRoleMenu business.SysRoleMenu
	_ = c.ShouldBindQuery(&sysRoleMenu)
	data, err := bizSev.GetSysRoleMenuSev().Get(c, sysRoleMenu.Id, "")
	if err != nil {
		global.LOG.Error("查询失败!", zap.Any("err", err))
		m.FailWithMessage("查询失败,"+err.Error(), c)
	} else {
		m.OkWithData(data, c)
	}
}

// GetSysRoleMenuList 分页获取SysRoleMenu列表
func (m *SysRoleMenuApi) GetSysRoleMenuList(c *gin.Context) {
	var pageInfo business.SysRoleMenuSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := bizSev.GetSysRoleMenuSev().GetList(c, pageInfo, business.Field_SysRoleMenu_mini); err != nil {
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
func (m *SysRoleMenuApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "sys_role_menu"
	if err := commSev.GetCommonDbSev().QuickEdit(c, quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		m.FailWithMessage("更新失败,"+err.Error(), c)
	} else {
		m.OkWithMessage("更新成功", c)
	}
}

// excelList 分页导出excel SysRoleMenu列表
func (m *SysRoleMenuApi) ExcelList(c *gin.Context) {
	var pageInfo business.SysRoleMenuSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, _, err := bizSev.GetSysRoleMenuSev().GetList(c, pageInfo, ""); err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		m.FailWithMessage("获取失败,"+err.Error(), c)
	} else {
		if len(list) == 0 {
			m.FailWithMessage("没有数据", c)
		} else {
			sheetFields := []string{}
			sheetFields = append(sheetFields, "角色")
			sheetFields = append(sheetFields, "菜单")
			sheetFields = append(sheetFields, "状态")

			excel := excelize.NewFile()
			excel.SetSheetRow("Sheet1", "A1", &sheetFields)
			for i, v := range list {
				axis := fmt.Sprintf("A%d", i+2)
				var arr = []interface{}{}
				arr = append(arr, v.RoleId)
				arr = append(arr, v.MenuId)
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
