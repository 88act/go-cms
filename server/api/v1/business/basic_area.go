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

type BasicAreaApi struct {
}

// CreateBasicArea 创建BasicArea
// @Tags BasicArea
// @Summary 创建BasicArea
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.BasicArea true "创建BasicArea"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /basicArea/createBasicArea [post]
func (basicAreaApi *BasicAreaApi) CreateBasicArea(c *gin.Context) {
	var dataObj business.BasicArea
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		response.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	}

	if id, err := bizSev.GetBasicAreaSev().Create(dataObj); err != nil {
		global.LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		idResp := &response.IdResp{Id: id}
		response.OkWithData(idResp, c)
	}
}

// DeleteBasicArea 删除BasicArea
// @Tags BasicArea
// @Summary 删除BasicArea
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.BasicArea true "删除BasicArea"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /basicArea/deleteBasicArea [delete]
func (basicAreaApi *BasicAreaApi) DeleteBasicArea(c *gin.Context) {
	var basicArea business.BasicArea
	_ = c.ShouldBindJSON(&basicArea)
	if err := bizSev.GetBasicAreaSev().Delete(basicArea); err != nil {
		global.LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBasicAreaByIds 批量删除BasicArea
// @Tags BasicArea
// @Summary 批量删除BasicArea
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BasicArea"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /basicArea/deleteBasicAreaByIds [delete]
func (basicAreaApi *BasicAreaApi) DeleteBasicAreaByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := bizSev.GetBasicAreaSev().DeleteByIds(IDS); err != nil {
		global.LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBasicArea 更新BasicArea
// @Tags BasicArea
// @Summary 更新BasicArea
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.BasicArea true "更新BasicArea"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /basicArea/updateBasicArea [put]
func (basicAreaApi *BasicAreaApi) UpdateBasicArea(c *gin.Context) {
	var dataObj business.BasicArea
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		response.FailWithMessage("更新失败,"+err.FirstString(), c)
		return
	}

	if err := bizSev.GetBasicAreaSev().Update(dataObj); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBasicArea 用id查询BasicArea
// @Tags BasicArea
// @Summary 用id查询BasicArea
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.BasicArea true "用id查询BasicArea"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /basicArea/findBasicArea [get]
func (basicAreaApi *BasicAreaApi) FindBasicArea(c *gin.Context) {
	var basicArea business.BasicArea
	_ = c.ShouldBindQuery(&basicArea)
	rebasicArea, err := bizSev.GetBasicAreaSev().Get(basicArea.ID, "")
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.OkWithData(gin.H{"basicArea": nil}, c)
	} else if err != nil {
		global.LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"basicArea": rebasicArea}, c)
	}
}

// GetBasicAreaList 分页获取BasicArea列表
// @Tags BasicArea
// @Summary 分页获取BasicArea列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.BasicAreaSearch true "分页获取BasicArea列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /basicArea/getBasicAreaList [get]
func (basicAreaApi *BasicAreaApi) GetBasicAreaList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")

	var pageInfo bizReq.BasicAreaSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := bizSev.GetBasicAreaSev().GetList(pageInfo, createdAtBetween, ""); err != nil {
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
// @Param data body business.BasicArea true "快速更新BasicArea"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router  /basicArea/quickEdit [post]
func (basicAreaApi *BasicAreaApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "basic_area"
	if err := commSev.GetCommonDbSev().QuickEdit(quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// excelList 分页导出excel BasicArea列表
// @Tags BasicArea
// @Summary 分页导出excel BasicArea列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.BasicAreaSearch true "分页导出excel BasicArea列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /basicArea/excelList [get]
func (basicAreaApi *BasicAreaApi) ExcelList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")
	var pageInfo bizReq.BasicAreaSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, _, err := bizSev.GetBasicAreaSev().GetListAll(pageInfo, createdAtBetween, ""); err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		if len(list) == 0 {
			response.FailWithMessage("没有数据", c)
		} else {
			sheetFields := []string{}
			sheetFields = append(sheetFields, "名称")
			sheetFields = append(sheetFields, "父栏目")
			sheetFields = append(sheetFields, "简称")
			sheetFields = append(sheetFields, "电话区号")
			sheetFields = append(sheetFields, "邮编")
			sheetFields = append(sheetFields, "拼音")
			sheetFields = append(sheetFields, "拼音简写")
			sheetFields = append(sheetFields, "经度")
			sheetFields = append(sheetFields, "纬度")
			sheetFields = append(sheetFields, "级别")
			sheetFields = append(sheetFields, "position字段")
			sheetFields = append(sheetFields, "组合名称")
			sheetFields = append(sheetFields, "排序")

			excel := excelize.NewFile()
			excel.SetSheetRow("Sheet1", "A1", &sheetFields)
			for i, v := range list {
				axis := fmt.Sprintf("A%d", i+2)
				var arr = []interface{}{}
				arr = append(arr, v.Areaname)
				if v.Parentid == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.Parentid)
				}
				arr = append(arr, v.Shortname)
				arr = append(arr, v.Areacode)
				arr = append(arr, v.Zipcode)
				arr = append(arr, v.Pinyin)
				arr = append(arr, v.Py)
				arr = append(arr, v.Lng)
				arr = append(arr, v.Lat)
				if v.Level == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.Level)
				}
				arr = append(arr, v.Position)
				arr = append(arr, v.Mergername)
				if v.Sort == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.Sort)
				}
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
