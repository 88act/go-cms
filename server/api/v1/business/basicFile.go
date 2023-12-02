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

type BasicFileApi struct {
	commSev.BaseApi
}

// CreateBasicFile 创建BasicFile
func (m *BasicFileApi) CreateBasicFile(c *gin.Context) {
	var dataObj business.BasicFile
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		m.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	}

	dataObj.UserId = utils.GetUserID(c)
	if id, err := bizSev.GetBasicFileSev().Create(c, dataObj); err != nil {
		global.LOG.Error("创建失败!", zap.Any("err", err))
		m.FailWithMessage("创建失败,"+err.Error(), c)
	} else {
		idResp := &response.IdResp{Id: id}
		m.OkWithData(idResp, c)
	}
}

// DeleteBasicFile 删除BasicFile

func (m *BasicFileApi) DeleteBasicFile(c *gin.Context) {
	var basicFile business.BasicFile
	_ = c.ShouldBindJSON(&basicFile)
	if err := bizSev.GetBasicFileSev().Delete(c, basicFile); err != nil {
		global.LOG.Error("删除失败!", zap.Any("err", err))
		m.FailWithMessage("删除失败,"+err.Error(), c)
	} else {
		m.OkWithMessage("删除成功", c)
	}
}

// DeleteBasicFileByIds 批量删除BasicFile
func (m *BasicFileApi) DeleteBasicFileByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := bizSev.GetBasicFileSev().DeleteByIds(c, IDS); err != nil {
		global.LOG.Error("删除失败!", zap.Any("err", err))
		m.FailWithMessage("删除失败,"+err.Error(), c)
	} else {
		m.OkWithMessage("删除成功", c)
	}
}

// UpdateBasicFile 更新BasicFile
func (m *BasicFileApi) UpdateBasicFile(c *gin.Context) {
	var dataObj business.BasicFile
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		m.FailWithMessage("更新失败,"+err.FirstString(), c)
		return
	}

	if err := bizSev.GetBasicFileSev().Update(c, dataObj); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		m.FailWithMessage("更新失败,"+err.Error(), c)
	} else {
		m.OkWithMessage("更新成功", c)
	}
}

// FindBasicFile 用id查询BasicFile
func (m *BasicFileApi) FindBasicFile(c *gin.Context) {
	var basicFile business.BasicFile
	_ = c.ShouldBindQuery(&basicFile)
	data, err := bizSev.GetBasicFileSev().Get(c, basicFile.Id, "")
	if err != nil {
		global.LOG.Error("查询失败!", zap.Any("err", err))
		m.FailWithMessage("查询失败,"+err.Error(), c)
	} else {
		m.OkWithData(data, c)
	}
}

// GetBasicFileList 分页获取BasicFile列表
func (m *BasicFileApi) GetBasicFileList(c *gin.Context) {
	var pageInfo business.BasicFileSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := bizSev.GetBasicFileSev().GetList(c, pageInfo, business.Field_BasicFile_mini); err != nil {
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
func (m *BasicFileApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "basic_file"
	if err := commSev.GetCommonDbSev().QuickEdit(c, quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		m.FailWithMessage("更新失败,"+err.Error(), c)
	} else {
		m.OkWithMessage("更新成功", c)
	}
}

// excelList 分页导出excel BasicFile列表
func (m *BasicFileApi) ExcelList(c *gin.Context) {
	var pageInfo business.BasicFileSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, _, err := bizSev.GetBasicFileSev().GetList(c, pageInfo, ""); err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		m.FailWithMessage("获取失败,"+err.Error(), c)
	} else {
		if len(list) == 0 {
			m.FailWithMessage("没有数据", c)
		} else {
			sheetFields := []string{}
			sheetFields = append(sheetFields, "唯一id")
			sheetFields = append(sheetFields, "用户id")

			sheetFields = append(sheetFields, "文件名")
			sheetFields = append(sheetFields, "栏目id")
			sheetFields = append(sheetFields, "模块id")
			sheetFields = append(sheetFields, "媒体类型")
			sheetFields = append(sheetFields, "存储位置")
			sheetFields = append(sheetFields, "文件路径")
			sheetFields = append(sheetFields, "文件类型")
			sheetFields = append(sheetFields, "文件类型")
			sheetFields = append(sheetFields, "文件大小[k]")
			sheetFields = append(sheetFields, "md5值")
			sheetFields = append(sheetFields, "sha散列值")
			sheetFields = append(sheetFields, "排序")
			sheetFields = append(sheetFields, "下载次数")
			sheetFields = append(sheetFields, "使用次数")
			sheetFields = append(sheetFields, "status字段")

			excel := excelize.NewFile()
			excel.SetSheetRow("Sheet1", "A1", &sheetFields)
			for i, v := range list {
				axis := fmt.Sprintf("A%d", i+2)
				var arr = []interface{}{}
				arr = append(arr, v.Guid)
				arr = append(arr, v.UserId)

				arr = append(arr, v.Name)
				arr = append(arr, v.CatId)
				arr = append(arr, v.Module)
				arr = append(arr, v.MediaType)
				arr = append(arr, v.Driver)
				arr = append(arr, v.Path)
				arr = append(arr, v.FileType)
				arr = append(arr, v.Ext)
				arr = append(arr, v.Size)
				arr = append(arr, v.Md5)
				arr = append(arr, v.Sha1)
				arr = append(arr, v.Sort)
				arr = append(arr, v.Download)
				arr = append(arr, v.UsedTime)
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
