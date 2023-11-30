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

type MemUserSafeApi struct {
	commSev.BaseApi
}

// CreateMemUserSafe 创建MemUserSafe
func (m *MemUserSafeApi) CreateMemUserSafe(c *gin.Context) {
	var dataObj business.MemUserSafe
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		m.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	}

	//dataObj.UserId = utils.GetUserID(c)
	if id, err := bizSev.GetMemUserSafeSev().Create(c, dataObj); err != nil {
		global.LOG.Error("创建失败!", zap.Any("err", err))
		m.FailWithMessage("创建失败,"+err.Error(), c)
	} else {
		idResp := &response.IdResp{Id: id}
		m.OkWithData(idResp, c)
	}
}

// DeleteMemUserSafe 删除MemUserSafe

func (m *MemUserSafeApi) DeleteMemUserSafe(c *gin.Context) {
	var memUserSafe business.MemUserSafe
	_ = c.ShouldBindJSON(&memUserSafe)
	if err := bizSev.GetMemUserSafeSev().Delete(c, memUserSafe); err != nil {
		global.LOG.Error("删除失败!", zap.Any("err", err))
		m.FailWithMessage("删除失败,"+err.Error(), c)
	} else {
		m.OkWithMessage("删除成功", c)
	}
}

// DeleteMemUserSafeByIds 批量删除MemUserSafe
func (m *MemUserSafeApi) DeleteMemUserSafeByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := bizSev.GetMemUserSafeSev().DeleteByIds(c, IDS); err != nil {
		global.LOG.Error("删除失败!", zap.Any("err", err))
		m.FailWithMessage("删除失败,"+err.Error(), c)
	} else {
		m.OkWithMessage("删除成功", c)
	}
}

// UpdateMemUserSafe 更新MemUserSafe
func (m *MemUserSafeApi) UpdateMemUserSafe(c *gin.Context) {
	var dataObj business.MemUserSafe
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		m.FailWithMessage("更新失败,"+err.FirstString(), c)
		return
	}

	if err := bizSev.GetMemUserSafeSev().Update(c, dataObj); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		m.FailWithMessage("更新失败,"+err.Error(), c)
	} else {
		m.OkWithMessage("更新成功", c)
	}
}

// FindMemUserSafe 用id查询MemUserSafe
func (m *MemUserSafeApi) FindMemUserSafe(c *gin.Context) {
	var memUserSafe business.MemUserSafe
	_ = c.ShouldBindQuery(&memUserSafe)
	data, err := bizSev.GetMemUserSafeSev().Get(c, memUserSafe.Id, "")
	if err != nil {
		global.LOG.Error("查询失败!", zap.Any("err", err))
		m.FailWithMessage("查询失败,"+err.Error(), c)
	} else {
		m.OkWithData(data, c)
	}
}

// GetMemUserSafeList 分页获取MemUserSafe列表
func (m *MemUserSafeApi) GetMemUserSafeList(c *gin.Context) {
	var pageInfo business.MemUserSafeSearch
	_ = c.ShouldBindQuery(&pageInfo)

	if list, total, err := bizSev.GetMemUserSafeSev().GetList(c, pageInfo, business.Field_MemUserSafe_mini); err != nil {
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
func (m *MemUserSafeApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "mem_user_safe"
	if err := commSev.GetCommonDbSev().QuickEdit(c, quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		m.FailWithMessage("更新失败,"+err.Error(), c)
	} else {
		m.OkWithMessage("更新成功", c)
	}
}

// excelList 分页导出excel MemUserSafe列表
func (m *MemUserSafeApi) ExcelList(c *gin.Context) {
	var pageInfo business.MemUserSafeSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, _, err := bizSev.GetMemUserSafeSev().GetList(c, pageInfo, ""); err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		m.FailWithMessage("获取失败,"+err.Error(), c)
	} else {
		if len(list) == 0 {
			m.FailWithMessage("没有数据", c)
		} else {
			sheetFields := []string{}

			sheetFields = append(sheetFields, "类型")
			sheetFields = append(sheetFields, "旧值")
			sheetFields = append(sheetFields, "新值")
			sheetFields = append(sheetFields, "ip")
			sheetFields = append(sheetFields, "ip城市")
			sheetFields = append(sheetFields, "状态")

			excel := excelize.NewFile()
			excel.SetSheetRow("Sheet1", "A1", &sheetFields)
			for i, v := range list {
				axis := fmt.Sprintf("A%d", i+2)
				var arr = []interface{}{}

				arr = append(arr, v.SafeType)
				arr = append(arr, v.ValOld)
				arr = append(arr, v.ValNew)
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
