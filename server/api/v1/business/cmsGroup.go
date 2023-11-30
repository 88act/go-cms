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

type CmsGroupApi struct {
	commSev.BaseApi
} 

// CreateCmsGroup 创建CmsGroup
func (m *CmsGroupApi) CreateCmsGroup(c *gin.Context) {
	var dataObj business.CmsGroup
	_ = c.ShouldBindJSON(&dataObj)
	
	if err := gvalid.CheckStruct(c,dataObj, nil); err != nil {
		m.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	} 
    
	dataObj.UserId = utils.GetUserID(c)
	if id,err := bizSev.GetCmsGroupSev().Create(c, dataObj); err != nil {
        global.LOG.Error("创建失败!", zap.Any("err", err))
		m.FailWithMessage("创建失败,"+err.Error(), c) 
	} else {
	    idResp := &response.IdResp{Id: id}
		m.OkWithData(idResp, c)
	}
}

// DeleteCmsGroup 删除CmsGroup

func (m *CmsGroupApi) DeleteCmsGroup(c *gin.Context) {
	var cmsGroup business.CmsGroup
	_ = c.ShouldBindJSON(&cmsGroup)
	if err := bizSev.GetCmsGroupSev().Delete(c, cmsGroup); err != nil {
        global.LOG.Error("删除失败!", zap.Any("err", err))
		m.FailWithMessage("删除失败,"+err.Error(), c) 
	} else {
		m.OkWithMessage("删除成功", c)
	}
}

// DeleteCmsGroupByIds 批量删除CmsGroup
func (m *CmsGroupApi) DeleteCmsGroupByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := bizSev.GetCmsGroupSev().DeleteByIds(c, IDS); err != nil {
        global.LOG.Error("删除失败!", zap.Any("err", err))
		m.FailWithMessage("删除失败,"+err.Error(), c) 
	} else {
		m.OkWithMessage("删除成功", c)
	}
}

// UpdateCmsGroup 更新CmsGroup
func (m *CmsGroupApi) UpdateCmsGroup(c *gin.Context) {
	var dataObj business.CmsGroup
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		m.FailWithMessage("更新失败,"+err.FirstString(), c)
		return
	}

	if err := bizSev.GetCmsGroupSev().Update(c, dataObj); err != nil {
        global.LOG.Error("更新失败!", zap.Any("err", err))
		m.FailWithMessage("更新失败,"+err.Error(), c) 
	} else {
		m.OkWithMessage("更新成功", c)
	}
}

// FindCmsGroup 用id查询CmsGroup
func (m *CmsGroupApi) FindCmsGroup(c *gin.Context) {
	var cmsGroup business.CmsGroup
	_ = c.ShouldBindQuery(&cmsGroup) 	 
    data,err:= bizSev.GetCmsGroupSev().Get(c, cmsGroup.Id,""); 
	if err != nil { 
        global.LOG.Error("查询失败!", zap.Any("err", err))
		m.FailWithMessage("查询失败,"+err.Error(), c) 
	} else {
		m.OkWithData(data, c)  
	}
}

// GetCmsGroupList 分页获取CmsGroup列表
func (m *CmsGroupApi) GetCmsGroupList(c *gin.Context) {
	var pageInfo business.CmsGroupSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if  list, total, err := bizSev.GetCmsGroupSev().GetList(c, pageInfo,business.Field_CmsGroup_mini); err != nil {
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
func (m *CmsGroupApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "cms_group" 
	if err := commSev.GetCommonDbSev().QuickEdit(c, quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		m.FailWithMessage("更新失败,"+err.Error(), c) 
	} else {
		m.OkWithMessage("更新成功", c)
	}
}


// excelList 分页导出excel CmsGroup列表
func (m *CmsGroupApi) ExcelList(c *gin.Context) {	 
	var pageInfo business.CmsGroupSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list,_,err:= bizSev.GetCmsGroupSev().GetList(c, pageInfo,""); err != nil {
	    global.LOG.Error("获取失败!", zap.Any("err", err))
        m.FailWithMessage("获取失败,"+err.Error(), c) 
    } else {
        if len(list) == 0 {
			m.FailWithMessage("没有数据", c)
		} else { 
			sheetFields := []string{}  
					sheetFields = append(sheetFields, "用户id")  
					sheetFields = append(sheetFields, "群组类别")  
					sheetFields = append(sheetFields, "名称")  
					sheetFields = append(sheetFields, "内容")  
					sheetFields = append(sheetFields, "标签")  
					sheetFields = append(sheetFields, "插图")  
					sheetFields = append(sheetFields, "文件")  
					sheetFields = append(sheetFields, "状态") 

			excel := excelize.NewFile()
			excel.SetSheetRow("Sheet1", "A1", &sheetFields)
			for i, v := range list {
				axis := fmt.Sprintf("A%d", i+2)
				var arr = []interface{}{} 
				arr = append(arr, v.UserId) 
				arr = append(arr, v.Type) 
				arr = append(arr, v.Title) 
				arr = append(arr, v.Desc) 
				arr = append(arr, v.TagList) 
				arr = append(arr, v.Image) 
				arr = append(arr, v.FileList) 
				arr = append(arr, v.Status)   
			    excel.SetSheetRow("Sheet1", axis,&arr)  
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


 