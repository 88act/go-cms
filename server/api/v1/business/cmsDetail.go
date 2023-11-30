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

type CmsDetailApi struct {
	commSev.BaseApi
}

// CreateCmsDetail 创建CmsDetail
func (m *CmsDetailApi) CreateCmsDetail(c *gin.Context) {
	var dataObj business.CmsDetail
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		m.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	}

	//dataObj.UserId = utils.GetUserID(c)
	if id, err := bizSev.GetCmsDetailSev().Create(c, dataObj); err != nil {
		global.LOG.Error("创建失败!", zap.Any("err", err))
		m.FailWithMessage("创建失败,"+err.Error(), c)
	} else {
		idResp := &response.IdResp{Id: id}
		m.OkWithData(idResp, c)
	}
}

// DeleteCmsDetail 删除CmsDetail

func (m *CmsDetailApi) DeleteCmsDetail(c *gin.Context) {
	var cmsDetail business.CmsDetail
	_ = c.ShouldBindJSON(&cmsDetail)
	if err := bizSev.GetCmsDetailSev().Delete(c, cmsDetail); err != nil {
		global.LOG.Error("删除失败!", zap.Any("err", err))
		m.FailWithMessage("删除失败,"+err.Error(), c)
	} else {
		m.OkWithMessage("删除成功", c)
	}
}

// DeleteCmsDetailByIds 批量删除CmsDetail
func (m *CmsDetailApi) DeleteCmsDetailByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := bizSev.GetCmsDetailSev().DeleteByIds(c, IDS); err != nil {
		global.LOG.Error("删除失败!", zap.Any("err", err))
		m.FailWithMessage("删除失败,"+err.Error(), c)
	} else {
		m.OkWithMessage("删除成功", c)
	}
}

// UpdateCmsDetail 更新CmsDetail
func (m *CmsDetailApi) UpdateCmsDetail(c *gin.Context) {
	var dataObj business.CmsDetail
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		m.FailWithMessage("更新失败,"+err.FirstString(), c)
		return
	}

	if err := bizSev.GetCmsDetailSev().Update(c, dataObj); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		m.FailWithMessage("更新失败,"+err.Error(), c)
	} else {
		m.OkWithMessage("更新成功", c)
	}
}

// FindCmsDetail 用id查询CmsDetail
func (m *CmsDetailApi) FindCmsDetail(c *gin.Context) {
	var cmsDetail business.CmsDetail
	_ = c.ShouldBindQuery(&cmsDetail)
	data, err := bizSev.GetCmsDetailSev().Get(c, cmsDetail.Id, "")
	if err != nil {
		global.LOG.Error("查询失败!", zap.Any("err", err))
		m.FailWithMessage("查询失败,"+err.Error(), c)
	} else {
		m.OkWithData(data, c)
	}
}

// GetCmsDetailList 分页获取CmsDetail列表
func (m *CmsDetailApi) GetCmsDetailList(c *gin.Context) {
	var pageInfo business.CmsDetailSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := bizSev.GetCmsDetailSev().GetList(c, pageInfo, business.Field_CmsDetail_mini); err != nil {
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
func (m *CmsDetailApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "cms_detail"
	if err := commSev.GetCommonDbSev().QuickEdit(c, quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		m.FailWithMessage("更新失败,"+err.Error(), c)
	} else {
		m.OkWithMessage("更新成功", c)
	}
}

// excelList 分页导出excel CmsDetail列表
func (m *CmsDetailApi) ExcelList(c *gin.Context) {
	var pageInfo business.CmsDetailSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, _, err := bizSev.GetCmsDetailSev().GetList(c, pageInfo, ""); err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		m.FailWithMessage("获取失败,"+err.Error(), c)
	} else {
		if len(list) == 0 {
			m.FailWithMessage("没有数据", c)
		} else {
			sheetFields := []string{}
			sheetFields = append(sheetFields, "文章id")
			sheetFields = append(sheetFields, "详细")

			excel := excelize.NewFile()
			excel.SetSheetRow("Sheet1", "A1", &sheetFields)
			for i, v := range list {
				axis := fmt.Sprintf("A%d", i+2)
				var arr = []interface{}{}
				arr = append(arr, v.ArtId)
				arr = append(arr, v.Detail)
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
