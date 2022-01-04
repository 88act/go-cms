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

type SysSuperBuilderHistoriesApi struct {
}

// CreateSysSuperBuilderHistories 创建SysSuperBuilderHistories
// @Tags SysSuperBuilderHistories
// @Summary 创建SysSuperBuilderHistories
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.SysSuperBuilderHistories true "创建SysSuperBuilderHistories"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysSuperBuilderHistories/createSysSuperBuilderHistories [post]
func (sysSuperBuilderHistoriesApi *SysSuperBuilderHistoriesApi) CreateSysSuperBuilderHistories(c *gin.Context) {
	var dataObj business.SysSuperBuilderHistories
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		response.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	}

	if id, err := bizSev.GetSysSuperBuilderHistoriesSev().Create(dataObj); err != nil {
		global.LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		idResp := &response.IdResp{Id: id}
		response.OkWithData(idResp, c)
	}
}

// DeleteSysSuperBuilderHistories 删除SysSuperBuilderHistories
// @Tags SysSuperBuilderHistories
// @Summary 删除SysSuperBuilderHistories
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.SysSuperBuilderHistories true "删除SysSuperBuilderHistories"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysSuperBuilderHistories/deleteSysSuperBuilderHistories [delete]
func (sysSuperBuilderHistoriesApi *SysSuperBuilderHistoriesApi) DeleteSysSuperBuilderHistories(c *gin.Context) {
	var sysSuperBuilderHistories business.SysSuperBuilderHistories
	_ = c.ShouldBindJSON(&sysSuperBuilderHistories)
	if err := bizSev.GetSysSuperBuilderHistoriesSev().Delete(sysSuperBuilderHistories); err != nil {
		global.LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSysSuperBuilderHistoriesByIds 批量删除SysSuperBuilderHistories
// @Tags SysSuperBuilderHistories
// @Summary 批量删除SysSuperBuilderHistories
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SysSuperBuilderHistories"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /sysSuperBuilderHistories/deleteSysSuperBuilderHistoriesByIds [delete]
func (sysSuperBuilderHistoriesApi *SysSuperBuilderHistoriesApi) DeleteSysSuperBuilderHistoriesByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := bizSev.GetSysSuperBuilderHistoriesSev().DeleteByIds(IDS); err != nil {
		global.LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSysSuperBuilderHistories 更新SysSuperBuilderHistories
// @Tags SysSuperBuilderHistories
// @Summary 更新SysSuperBuilderHistories
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.SysSuperBuilderHistories true "更新SysSuperBuilderHistories"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysSuperBuilderHistories/updateSysSuperBuilderHistories [put]
func (sysSuperBuilderHistoriesApi *SysSuperBuilderHistoriesApi) UpdateSysSuperBuilderHistories(c *gin.Context) {
	var dataObj business.SysSuperBuilderHistories
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		response.FailWithMessage("更新失败,"+err.FirstString(), c)
		return
	}

	if err := bizSev.GetSysSuperBuilderHistoriesSev().Update(dataObj); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSysSuperBuilderHistories 用id查询SysSuperBuilderHistories
// @Tags SysSuperBuilderHistories
// @Summary 用id查询SysSuperBuilderHistories
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.SysSuperBuilderHistories true "用id查询SysSuperBuilderHistories"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysSuperBuilderHistories/findSysSuperBuilderHistories [get]
func (sysSuperBuilderHistoriesApi *SysSuperBuilderHistoriesApi) FindSysSuperBuilderHistories(c *gin.Context) {
	var sysSuperBuilderHistories business.SysSuperBuilderHistories
	_ = c.ShouldBindQuery(&sysSuperBuilderHistories)
	resysSuperBuilderHistories, err := bizSev.GetSysSuperBuilderHistoriesSev().Get(sysSuperBuilderHistories.ID, "")
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.OkWithData(gin.H{"sysSuperBuilderHistories": nil}, c)
	} else if err != nil {
		global.LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"sysSuperBuilderHistories": resysSuperBuilderHistories}, c)
	}
}

// GetSysSuperBuilderHistoriesList 分页获取SysSuperBuilderHistories列表
// @Tags SysSuperBuilderHistories
// @Summary 分页获取SysSuperBuilderHistories列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.SysSuperBuilderHistoriesSearch true "分页获取SysSuperBuilderHistories列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysSuperBuilderHistories/getSysSuperBuilderHistoriesList [get]
func (sysSuperBuilderHistoriesApi *SysSuperBuilderHistoriesApi) GetSysSuperBuilderHistoriesList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")

	var pageInfo bizReq.SysSuperBuilderHistoriesSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := bizSev.GetSysSuperBuilderHistoriesSev().GetList(pageInfo, createdAtBetween, ""); err != nil {
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
// @Param data body business.SysSuperBuilderHistories true "快速更新SysSuperBuilderHistories"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router  /sysSuperBuilderHistories/quickEdit [post]
func (sysSuperBuilderHistoriesApi *SysSuperBuilderHistoriesApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "sys_super_builder_histories"
	if err := commSev.GetCommonDbSev().QuickEdit(quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// excelList 分页导出excel SysSuperBuilderHistories列表
// @Tags SysSuperBuilderHistories
// @Summary 分页导出excel SysSuperBuilderHistories列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.SysSuperBuilderHistoriesSearch true "分页导出excel SysSuperBuilderHistories列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysSuperBuilderHistories/excelList [get]
func (sysSuperBuilderHistoriesApi *SysSuperBuilderHistoriesApi) ExcelList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")
	var pageInfo bizReq.SysSuperBuilderHistoriesSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, _, err := bizSev.GetSysSuperBuilderHistoriesSev().GetListAll(pageInfo, createdAtBetween, ""); err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		if len(list) == 0 {
			response.FailWithMessage("没有数据", c)
		} else {
			sheetFields := []string{}
			sheetFields = append(sheetFields, "表名")
			sheetFields = append(sheetFields, "表单")
			sheetFields = append(sheetFields, "superBuilderPath字段")
			sheetFields = append(sheetFields, "injectionMeta字段")
			sheetFields = append(sheetFields, "struct名称")
			sheetFields = append(sheetFields, "struct名称")
			sheetFields = append(sheetFields, "apiIds字段")
			sheetFields = append(sheetFields, "flag字段")

			excel := excelize.NewFile()
			excel.SetSheetRow("Sheet1", "A1", &sheetFields)
			for i, v := range list {
				axis := fmt.Sprintf("A%d", i+2)
				var arr = []interface{}{}
				arr = append(arr, v.TableName)
				arr = append(arr, v.RequestMeta)
				arr = append(arr, v.SuperBuilderPath)
				arr = append(arr, v.InjectionMeta)
				arr = append(arr, v.StructName)
				arr = append(arr, v.StructCnName)
				arr = append(arr, v.ApiIds)
				if v.Flag == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.Flag)
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
