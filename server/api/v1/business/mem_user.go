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

type MemUserApi struct {
}

// CreateMemUser 创建MemUser
// @Tags MemUser
// @Summary 创建MemUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.MemUser true "创建MemUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /memUser/createMemUser [post]
func (memUserApi *MemUserApi) CreateMemUser(c *gin.Context) {
	var dataObj business.MemUser
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		response.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	}

	if id, err := bizSev.GetMemUserSev().Create(dataObj); err != nil {
		global.LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		idResp := &response.IdResp{Id: id}
		response.OkWithData(idResp, c)
	}
}

// DeleteMemUser 删除MemUser
// @Tags MemUser
// @Summary 删除MemUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.MemUser true "删除MemUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /memUser/deleteMemUser [delete]
func (memUserApi *MemUserApi) DeleteMemUser(c *gin.Context) {
	var memUser business.MemUser
	_ = c.ShouldBindJSON(&memUser)
	if err := bizSev.GetMemUserSev().Delete(memUser); err != nil {
		global.LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMemUserByIds 批量删除MemUser
// @Tags MemUser
// @Summary 批量删除MemUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除MemUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /memUser/deleteMemUserByIds [delete]
func (memUserApi *MemUserApi) DeleteMemUserByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := bizSev.GetMemUserSev().DeleteByIds(IDS); err != nil {
		global.LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMemUser 更新MemUser
// @Tags MemUser
// @Summary 更新MemUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.MemUser true "更新MemUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /memUser/updateMemUser [put]
func (memUserApi *MemUserApi) UpdateMemUser(c *gin.Context) {
	var dataObj business.MemUser
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		response.FailWithMessage("更新失败,"+err.FirstString(), c)
		return
	}

	if err := bizSev.GetMemUserSev().Update(dataObj); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMemUser 用id查询MemUser
// @Tags MemUser
// @Summary 用id查询MemUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.MemUser true "用id查询MemUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /memUser/findMemUser [get]
func (memUserApi *MemUserApi) FindMemUser(c *gin.Context) {
	var memUser business.MemUser
	_ = c.ShouldBindQuery(&memUser)
	rememUser, err := bizSev.GetMemUserSev().Get(memUser.ID, "")
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.OkWithData(gin.H{"memUser": nil}, c)
	} else if err != nil {
		global.LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"memUser": rememUser}, c)
	}
}

// GetMemUserList 分页获取MemUser列表
// @Tags MemUser
// @Summary 分页获取MemUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.MemUserSearch true "分页获取MemUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /memUser/getMemUserList [get]
func (memUserApi *MemUserApi) GetMemUserList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")

	var pageInfo bizReq.MemUserSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := bizSev.GetMemUserSev().GetList(pageInfo, createdAtBetween, ""); err != nil {
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
// @Param data body business.MemUser true "快速更新MemUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router  /memUser/quickEdit [post]
func (memUserApi *MemUserApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "mem_user"
	if err := commSev.GetCommonDbSev().QuickEdit(quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// excelList 分页导出excel MemUser列表
// @Tags MemUser
// @Summary 分页导出excel MemUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.MemUserSearch true "分页导出excel MemUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /memUser/excelList [get]
func (memUserApi *MemUserApi) ExcelList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")
	var pageInfo bizReq.MemUserSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, _, err := bizSev.GetMemUserSev().GetListAll(pageInfo, createdAtBetween, ""); err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		if len(list) == 0 {
			response.FailWithMessage("没有数据", c)
		} else {
			sheetFields := []string{}
			sheetFields = append(sheetFields, "用户名")
			sheetFields = append(sheetFields, "密码")
			sheetFields = append(sheetFields, "密码盐")
			sheetFields = append(sheetFields, "邮件")
			sheetFields = append(sheetFields, "手机")
			sheetFields = append(sheetFields, "昵称")
			sheetFields = append(sheetFields, "真实名")
			sheetFields = append(sheetFields, "身份证")
			sheetFields = append(sheetFields, "性别")
			sheetFields = append(sheetFields, "生日")
			sheetFields = append(sheetFields, "头像")
			sheetFields = append(sheetFields, "验证手机")
			sheetFields = append(sheetFields, "验证邮箱")
			sheetFields = append(sheetFields, "验证实名")
			sheetFields = append(sheetFields, "登录次数")
			sheetFields = append(sheetFields, "推荐人ID")
			sheetFields = append(sheetFields, "推荐人ID2")
			sheetFields = append(sheetFields, "推荐码16位（自己的）")
			sheetFields = append(sheetFields, "状态")
			sheetFields = append(sheetFields, "安全状态")
			sheetFields = append(sheetFields, "注册ip")
			sheetFields = append(sheetFields, "登录ip")
			sheetFields = append(sheetFields, "最后登录时间")

			excel := excelize.NewFile()
			excel.SetSheetRow("Sheet1", "A1", &sheetFields)
			for i, v := range list {
				axis := fmt.Sprintf("A%d", i+2)
				var arr = []interface{}{}
				arr = append(arr, v.Username)
				arr = append(arr, v.Pws)
				arr = append(arr, v.PwsSlat)
				arr = append(arr, v.Email)
				arr = append(arr, v.Mobile)
				arr = append(arr, v.Nickname)
				arr = append(arr, v.Realname)
				arr = append(arr, v.CardId)
				if v.Sex == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.Sex)
				}
				arr = append(arr, v.Birthday)
				arr = append(arr, v.Avatar)
				if v.MobileValidated == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.MobileValidated)
				}
				if v.EmailValidated == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.EmailValidated)
				}
				if v.RealnameValidated == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.RealnameValidated)
				}
				if v.LoginTimes == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.LoginTimes)
				}
				if v.RecommendId == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.RecommendId)
				}
				if v.RecommendId2 == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.RecommendId2)
				}
				arr = append(arr, v.RecommendCode)
				if v.Status == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.Status)
				}
				if v.StatusSafe == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.StatusSafe)
				}
				if v.RegIp == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.RegIp)
				}
				if v.LastLoginIp == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.LastLoginIp)
				}
				arr = append(arr, v.LastLoginTime)
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
