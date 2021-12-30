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

type CmsCatApi struct {
}

// CreateCmsCat 创建CmsCat
// @Tags CmsCat
// @Summary 创建CmsCat
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.CmsCat true "创建CmsCat"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmsCat/createCmsCat [post]
func (cmsCatApi *CmsCatApi) CreateCmsCat(c *gin.Context) {
	var dataObj business.CmsCat
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		response.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	}

	if id, err := bizSev.GetCmsCatSev().Create(dataObj); err != nil {
		global.LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		idResp := &response.IdResp{Id: id}
		response.OkWithData(idResp, c)
	}
}

// DeleteCmsCat 删除CmsCat
// @Tags CmsCat
// @Summary 删除CmsCat
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.CmsCat true "删除CmsCat"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmsCat/deleteCmsCat [delete]
func (cmsCatApi *CmsCatApi) DeleteCmsCat(c *gin.Context) {
	var cmsCat business.CmsCat
	_ = c.ShouldBindJSON(&cmsCat)
	if err := bizSev.GetCmsCatSev().Delete(cmsCat); err != nil {
		global.LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCmsCatByIds 批量删除CmsCat
// @Tags CmsCat
// @Summary 批量删除CmsCat
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CmsCat"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /cmsCat/deleteCmsCatByIds [delete]
func (cmsCatApi *CmsCatApi) DeleteCmsCatByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := bizSev.GetCmsCatSev().DeleteByIds(IDS); err != nil {
		global.LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCmsCat 更新CmsCat
// @Tags CmsCat
// @Summary 更新CmsCat
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.CmsCat true "更新CmsCat"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cmsCat/updateCmsCat [put]
func (cmsCatApi *CmsCatApi) UpdateCmsCat(c *gin.Context) {
	var dataObj business.CmsCat
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		response.FailWithMessage("更新失败,"+err.FirstString(), c)
		return
	}

	if err := bizSev.GetCmsCatSev().Update(dataObj); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCmsCat 用id查询CmsCat
// @Tags CmsCat
// @Summary 用id查询CmsCat
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.CmsCat true "用id查询CmsCat"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cmsCat/findCmsCat [get]
func (cmsCatApi *CmsCatApi) FindCmsCat(c *gin.Context) {
	var cmsCat business.CmsCat
	_ = c.ShouldBindQuery(&cmsCat)
	recmsCat, err := bizSev.GetCmsCatSev().Get(cmsCat.ID, "")
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.OkWithData(gin.H{"cmsCat": nil}, c)
	} else if err != nil {
		global.LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"cmsCat": recmsCat}, c)
	}
}

// GetCmsCatList 分页获取CmsCat列表
// @Tags CmsCat
// @Summary 分页获取CmsCat列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.CmsCatSearch true "分页获取CmsCat列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmsCat/getCmsCatList [get]
func (cmsCatApi *CmsCatApi) GetCmsCatList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")

	var pageInfo bizReq.CmsCatSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := bizSev.GetCmsCatSev().GetList(pageInfo, createdAtBetween, ""); err != nil {
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
// @Param data body business.CmsCat true "快速更新CmsCat"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router  /cmsCat/quickEdit [post]
func (cmsCatApi *CmsCatApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "cms_cat"
	if err := commSev.GetCommonDbSev().QuickEdit(quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// excelList 分页导出excel CmsCat列表
// @Tags CmsCat
// @Summary 分页导出excel CmsCat列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.CmsCatSearch true "分页导出excel CmsCat列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmsCat/excelList [get]
func (cmsCatApi *CmsCatApi) ExcelList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")
	var pageInfo bizReq.CmsCatSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, _, err := bizSev.GetCmsCatSev().GetListAll(pageInfo, createdAtBetween, ""); err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		if len(list) == 0 {
			response.FailWithMessage("没有数据", c)
		} else {
			sheetFields := []string{}
			sheetFields = append(sheetFields, "父ID")
			sheetFields = append(sheetFields, "系统分类")
			sheetFields = append(sheetFields, "群组id")
			sheetFields = append(sheetFields, "文章类型")
			sheetFields = append(sheetFields, "名称")
			sheetFields = append(sheetFields, "配图")
			sheetFields = append(sheetFields, "排序")
			sheetFields = append(sheetFields, "是否导航")
			sheetFields = append(sheetFields, "描述")
			sheetFields = append(sheetFields, "关键词")
			sheetFields = append(sheetFields, "别名")
			sheetFields = append(sheetFields, "状态")

			excel := excelize.NewFile()
			excel.SetSheetRow("Sheet1", "A1", &sheetFields)
			for i, v := range list {
				axis := fmt.Sprintf("A%d", i+2)
				var arr = []interface{}{}
				arr = append(arr, *v.Pid)
				arr = append(arr, *v.BeSys)
				arr = append(arr, *v.GroupId)
				arr = append(arr, *v.MediaType)
				arr = append(arr, v.Name)
				arr = append(arr, v.Thumb)
				arr = append(arr, *v.Sort)
				arr = append(arr, *v.BeNav)
				arr = append(arr, v.Desc)
				arr = append(arr, v.Keywords)
				arr = append(arr, v.Alias)
				arr = append(arr, *v.Status)
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
