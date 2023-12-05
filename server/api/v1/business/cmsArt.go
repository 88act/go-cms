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

type CmsArtApi struct {
	commSev.BaseApi
}

// CreateCmsArt 创建CmsArt
func (m *CmsArtApi) CreateCmsArt(c *gin.Context) {
	var dataObj business.CmsArt
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		m.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	}

	dataObj.UserId = utils.GetUserID(c)
	if id, err := bizSev.GetCmsArtSev().Create(c, dataObj); err != nil {
		global.LOG.Error("创建失败!", zap.Any("err", err))
		m.FailWithMessage("创建失败,"+err.Error(), c)
	} else {
		idResp := &response.IdResp{Id: id}
		m.OkWithData(idResp, c)
	}
}

// DeleteCmsArt 删除CmsArt

func (m *CmsArtApi) DeleteCmsArt(c *gin.Context) {
	var cmsArt business.CmsArt
	_ = c.ShouldBindJSON(&cmsArt)
	if err := bizSev.GetCmsArtSev().Delete(c, cmsArt); err != nil {
		global.LOG.Error("删除失败!", zap.Any("err", err))
		m.FailWithMessage("删除失败,"+err.Error(), c)
	} else {
		m.OkWithMessage("删除成功", c)
	}
}

// DeleteCmsArtByIds 批量删除CmsArt
func (m *CmsArtApi) DeleteCmsArtByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := bizSev.GetCmsArtSev().DeleteByIds(c, IDS); err != nil {
		global.LOG.Error("删除失败!", zap.Any("err", err))
		m.FailWithMessage("删除失败,"+err.Error(), c)
	} else {
		m.OkWithMessage("删除成功", c)
	}
}

// UpdateCmsArt 更新CmsArt
func (m *CmsArtApi) UpdateCmsArt(c *gin.Context) {
	var dataObj business.CmsArt
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		m.FailWithMessage("更新失败,"+err.FirstString(), c)
		return
	}

	if err := bizSev.GetCmsArtSev().Update(c, dataObj); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		m.FailWithMessage("更新失败,"+err.Error(), c)
	} else {
		m.OkWithMessage("更新成功", c)
	}
}

// FindCmsArt 用id查询CmsArt
func (m *CmsArtApi) FindCmsArt(c *gin.Context) {
	var cmsArt business.CmsArt
	_ = c.ShouldBindQuery(&cmsArt)
	data, err := bizSev.GetCmsArtSev().Get(c, cmsArt.Id, "")
	if err != nil {
		global.LOG.Error("查询失败!", zap.Any("err", err))
		m.FailWithMessage("查询失败,"+err.Error(), c)
	} else {
		m.OkWithData(data, c)
	}
}

// GetCmsArtList 分页获取CmsArt列表
func (m *CmsArtApi) GetCmsArtList(c *gin.Context) {
	var pageInfo business.CmsArtSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := bizSev.GetCmsArtSev().GetList(c, pageInfo, business.Field_CmsArt_mini); err != nil {
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
func (m *CmsArtApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "cms_art"
	if err := commSev.GetCommonDbSev().QuickEdit(c, quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		m.FailWithMessage("更新失败,"+err.Error(), c)
	} else {
		m.OkWithMessage("更新成功", c)
	}
}

// excelList 分页导出excel CmsArt列表
func (m *CmsArtApi) ExcelList(c *gin.Context) {
	var pageInfo business.CmsArtSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, _, err := bizSev.GetCmsArtSev().GetList(c, pageInfo, ""); err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		m.FailWithMessage("获取失败,"+err.Error(), c)
	} else {
		if len(list) == 0 {
			m.FailWithMessage("没有数据", c)
		} else {
			sheetFields := []string{}
			sheetFields = append(sheetFields, "父id")
			sheetFields = append(sheetFields, "用户id")
			sheetFields = append(sheetFields, "文章类型")
			sheetFields = append(sheetFields, "文章标题")
			sheetFields = append(sheetFields, "文章摘要")
			sheetFields = append(sheetFields, "标签列表")
			sheetFields = append(sheetFields, "来源")
			sheetFields = append(sheetFields, "插图")
			sheetFields = append(sheetFields, "媒体列表")
			sheetFields = append(sheetFields, "链接地址")

			sheetFields = append(sheetFields, "综合指数")
			sheetFields = append(sheetFields, "总分享")
			sheetFields = append(sheetFields, "总收藏")
			sheetFields = append(sheetFields, "总评论")
			sheetFields = append(sheetFields, "总点击")
			sheetFields = append(sheetFields, "总星")
			sheetFields = append(sheetFields, "总赞")
			sheetFields = append(sheetFields, "总踩")
			sheetFields = append(sheetFields, "状态")
			sheetFields = append(sheetFields, "审核信息")

			excel := excelize.NewFile()
			excel.SetSheetRow("Sheet1", "A1", &sheetFields)
			for i, v := range list {
				axis := fmt.Sprintf("A%d", i+2)
				var arr = []interface{}{}
				arr = append(arr, v.Pid)
				arr = append(arr, v.UserId)

				arr = append(arr, v.Type)
				arr = append(arr, v.Title)
				arr = append(arr, v.Desc)
				arr = append(arr, v.TagList)
				arr = append(arr, v.Source)
				arr = append(arr, v.Image)
				arr = append(arr, v.FileList)
				arr = append(arr, v.Link)

				arr = append(arr, v.TotalWhole)
				arr = append(arr, v.TotalShare)
				arr = append(arr, v.TotalFav)
				arr = append(arr, v.TotalDiscuss)
				arr = append(arr, v.TotalClick)
				arr = append(arr, v.TotalStar)
				arr = append(arr, v.TotalGood)
				arr = append(arr, v.TotalPoor)
				arr = append(arr, v.Status)
				arr = append(arr, v.VerifyMsg)
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
