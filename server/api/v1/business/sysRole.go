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

type SysRoleApi struct {
	commSev.BaseApi
}

// CreateSysRole 创建SysRole
func (m *SysRoleApi) CreateSysRole(c *gin.Context) {
	var dataObj business.SysRole
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		m.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	}

	dataObj.UserId = utils.GetUserID(c)

	if id, err := bizSev.GetSysRoleSev().Create(c, dataObj); err != nil {
		global.LOG.Error("创建失败!", zap.Any("err", err))
		m.FailWithMessage("创建失败,"+err.Error(), c)
	} else {
		idResp := &response.IdResp{Id: id}
		m.OkWithData(idResp, c)
	}
}

// DeleteSysRole 删除SysRole

func (m *SysRoleApi) DeleteSysRole(c *gin.Context) {
	var sysRole business.SysRole
	_ = c.ShouldBindJSON(&sysRole)
	if err := bizSev.GetSysRoleSev().Delete(c, sysRole); err != nil {
		global.LOG.Error("删除失败!", zap.Any("err", err))
		m.FailWithMessage("删除失败,"+err.Error(), c)
	} else {
		m.OkWithMessage("删除成功", c)
	}
}

// DeleteSysRoleByIds 批量删除SysRole
func (m *SysRoleApi) DeleteSysRoleByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := bizSev.GetSysRoleSev().DeleteByIds(c, IDS); err != nil {
		global.LOG.Error("删除失败!", zap.Any("err", err))
		m.FailWithMessage("删除失败,"+err.Error(), c)
	} else {
		m.OkWithMessage("删除成功", c)
	}
}

// UpdateSysRole 更新SysRole
func (m *SysRoleApi) UpdateSysRole(c *gin.Context) {
	var dataObj business.SysRole
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		m.FailWithMessage("更新失败,"+err.FirstString(), c)
		return
	}

	if err := bizSev.GetSysRoleSev().Update(c, dataObj); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		m.FailWithMessage("更新失败,"+err.Error(), c)
	} else {
		m.OkWithMessage("更新成功", c)
	}
}

// FindSysRole 用id查询SysRole
func (m *SysRoleApi) FindSysRole(c *gin.Context) {
	var sysRole business.SysRole
	_ = c.ShouldBindQuery(&sysRole)
	data, err := bizSev.GetSysRoleSev().Get(c, sysRole.Id, "")
	if err != nil {
		global.LOG.Error("查询失败!", zap.Any("err", err))
		m.FailWithMessage("查询失败,"+err.Error(), c)
	} else {
		m.OkWithData(data, c)
	}
}

// GetSysRoleList 分页获取SysRole列表
func (m *SysRoleApi) GetSysRoleList(c *gin.Context) {
	var pageInfo business.SysRoleSearch
	_ = c.ShouldBindQuery(&pageInfo)

	if list, total, err := bizSev.GetSysRoleSev().GetList(c, pageInfo, business.Field_SysRole_mini); err != nil {
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
func (m *SysRoleApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "sys_role"
	if err := commSev.GetCommonDbSev().QuickEdit(c, quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		m.FailWithMessage("更新失败,"+err.Error(), c)
	} else {
		m.OkWithMessage("更新成功", c)
	}
}

// excelList 分页导出excel SysRole列表
func (m *SysRoleApi) ExcelList(c *gin.Context) {
	var pageInfo business.SysRoleSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, _, err := bizSev.GetSysRoleSev().GetList(c, pageInfo, ""); err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		m.FailWithMessage("获取失败,"+err.Error(), c)
	} else {
		if len(list) == 0 {
			m.FailWithMessage("没有数据", c)
		} else {
			sheetFields := []string{}

			sheetFields = append(sheetFields, "父id")
			sheetFields = append(sheetFields, "角色名")
			sheetFields = append(sheetFields, "角色编码")
			sheetFields = append(sheetFields, "默认菜单")
			sheetFields = append(sheetFields, "状态")

			excel := excelize.NewFile()
			excel.SetSheetRow("Sheet1", "A1", &sheetFields)
			for i, v := range list {
				axis := fmt.Sprintf("A%d", i+2)
				var arr = []interface{}{}

				arr = append(arr, v.Pid)
				arr = append(arr, v.RoleName)
				arr = append(arr, v.RoleCode)
				arr = append(arr, v.DefaultMenu)
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
