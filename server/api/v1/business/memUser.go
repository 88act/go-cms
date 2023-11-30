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

type MemUserApi struct {
	commSev.BaseApi
}

// CreateJqCustomer
func (m *MemUserApi) MemChangePwd(c *gin.Context) {
	var userCPW business.ChangePasswordStruct
	_ = c.ShouldBindJSON(&userCPW)
	if len(userCPW.NewPassword) < 6 {
		response.FailWithMessage("新密码长度不对", c)
		c.Abort()
		return
	}

	//如果是修改其他人， 必须是 admin id = 1
	userId := utils.GetUserID(c)
	if err := bizSev.GetMemUserSev().ChangePasswordCu(c, userCPW, userId); err == nil {
		response.OkWithMessage("修改成功", c)
	} else {
		response.FailWithMessage("修改失败", c)
	}

}

// CreateMemUser 创建MemUser
func (m *MemUserApi) CreateMemUser(c *gin.Context) {
	var dataObj business.MemUser
	_ = c.ShouldBindJSON(&dataObj)

	dataObj.UserId = utils.GetUserID(c)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		m.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	}

	dataObj.UserId = utils.GetUserID(c)
	if id, err := bizSev.GetMemUserSev().Create(c, dataObj); err != nil {
		global.LOG.Error("创建失败!", zap.Any("err", err))
		m.FailWithMessage("创建失败,"+err.Error(), c)
	} else {
		idResp := &response.IdResp{Id: id}
		m.OkWithData(idResp, c)
	}
}

// DeleteMemUser 删除MemUser

func (m *MemUserApi) DeleteMemUser(c *gin.Context) {
	var memUser business.MemUser
	_ = c.ShouldBindJSON(&memUser)
	if err := bizSev.GetMemUserSev().Delete(c, memUser); err != nil {
		global.LOG.Error("删除失败!", zap.Any("err", err))
		m.FailWithMessage("删除失败,"+err.Error(), c)
	} else {
		m.OkWithMessage("删除成功", c)
	}
}

// DeleteMemUserByIds 批量删除MemUser
func (m *MemUserApi) DeleteMemUserByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := bizSev.GetMemUserSev().DeleteByIds(c, IDS); err != nil {
		global.LOG.Error("删除失败!", zap.Any("err", err))
		m.FailWithMessage("删除失败,"+err.Error(), c)
	} else {
		m.OkWithMessage("删除成功", c)
	}
}

// UpdateMemUser 更新MemUser
func (m *MemUserApi) UpdateMemUser(c *gin.Context) {
	var dataObj business.MemUser
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		m.FailWithMessage("更新失败,"+err.FirstString(), c)
		return
	}

	if err := bizSev.GetMemUserSev().Update(c, dataObj); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		m.FailWithMessage("更新失败,"+err.Error(), c)
	} else {
		m.OkWithMessage("更新成功", c)
	}
}

// FindMemUser 用id查询MemUser
func (m *MemUserApi) FindMemUser(c *gin.Context) {
	var memUser business.MemUser
	_ = c.ShouldBindQuery(&memUser)
	data, err := bizSev.GetMemUserSev().Get(c, memUser.Id, "")
	if err != nil {
		global.LOG.Error("查询失败!", zap.Any("err", err))
		m.FailWithMessage("查询失败,"+err.Error(), c)
	} else {
		m.OkWithData(data, c)
	}
}

// GetMemUserList 分页获取MemUser列表
func (m *MemUserApi) GetMemUserList(c *gin.Context) {
	var pageInfo business.MemUserSearch
	_ = c.ShouldBindQuery(&pageInfo)

	if list, total, err := bizSev.GetMemUserSev().GetList(c, pageInfo, business.Field_MemUser_mini); err != nil {
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
func (m *MemUserApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "mem_user"
	if err := commSev.GetCommonDbSev().QuickEdit(c, quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		m.FailWithMessage("更新失败,"+err.Error(), c)
	} else {
		m.OkWithMessage("更新成功", c)
	}
}

// excelList 分页导出excel MemUser列表
func (m *MemUserApi) ExcelList(c *gin.Context) {
	var pageInfo business.MemUserSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, _, err := bizSev.GetMemUserSev().GetList(c, pageInfo, ""); err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		m.FailWithMessage("获取失败,"+err.Error(), c)
	} else {
		if len(list) == 0 {
			m.FailWithMessage("没有数据", c)
		} else {
			sheetFields := []string{}

			sheetFields = append(sheetFields, "录入人")
			sheetFields = append(sheetFields, "唯一id")
			sheetFields = append(sheetFields, "用户名")
			sheetFields = append(sheetFields, "密码")
			sheetFields = append(sheetFields, "密码盐")
			sheetFields = append(sheetFields, "昵称")
			sheetFields = append(sheetFields, "真实名")
			sheetFields = append(sheetFields, "角色list")
			sheetFields = append(sheetFields, "角色")
			sheetFields = append(sheetFields, "类型")
			sheetFields = append(sheetFields, "邮件")
			sheetFields = append(sheetFields, "手机")
			sheetFields = append(sheetFields, "身份证")
			sheetFields = append(sheetFields, "性别")
			sheetFields = append(sheetFields, "生日")
			sheetFields = append(sheetFields, "头像")
			sheetFields = append(sheetFields, "岗位")
			sheetFields = append(sheetFields, "部门")
			sheetFields = append(sheetFields, "验证手机")
			sheetFields = append(sheetFields, "验证邮箱")
			sheetFields = append(sheetFields, "验证实名")
			sheetFields = append(sheetFields, "备注")
			sheetFields = append(sheetFields, "安全状态")
			sheetFields = append(sheetFields, "注册ip")
			sheetFields = append(sheetFields, "登录ip")
			sheetFields = append(sheetFields, "登录次数")
			sheetFields = append(sheetFields, "最后登录时间")
			sheetFields = append(sheetFields, "扫码")
			sheetFields = append(sheetFields, "扫码")
			sheetFields = append(sheetFields, "设置值")
			sheetFields = append(sheetFields, "远程协助模式")
			sheetFields = append(sheetFields, "状态")

			excel := excelize.NewFile()
			excel.SetSheetRow("Sheet1", "A1", &sheetFields)
			for i, v := range list {
				axis := fmt.Sprintf("A%d", i+2)
				var arr = []interface{}{}

				arr = append(arr, v.UserId)
				arr = append(arr, v.Guid)
				arr = append(arr, v.Username)
				arr = append(arr, v.Password)
				arr = append(arr, v.PasswordSlat)
				arr = append(arr, v.Nickname)
				arr = append(arr, v.Realname)
				arr = append(arr, v.RoleList)
				arr = append(arr, v.Role)
				arr = append(arr, v.UserType)
				arr = append(arr, v.Email)
				arr = append(arr, v.Mobile)
				arr = append(arr, v.CardId)
				arr = append(arr, v.Sex)
				arr = append(arr, v.Birthday)
				arr = append(arr, v.Avatar)
				arr = append(arr, v.JobId)
				arr = append(arr, v.DepartId)
				arr = append(arr, v.MobileValidated)
				arr = append(arr, v.EmailValidated)
				arr = append(arr, v.CardidValidated)
				arr = append(arr, v.Remark)
				arr = append(arr, v.StatusSafe)
				arr = append(arr, v.RegIp)
				arr = append(arr, v.LoginIp)
				arr = append(arr, v.LoginTotal)
				arr = append(arr, v.LoginTime)
				arr = append(arr, v.Scan)
				arr = append(arr, v.ScanTime)
				arr = append(arr, v.Setting)
				arr = append(arr, v.RtcModel)
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
