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

type CmsDiscussApi struct {
	commSev.BaseApi
} 

// CreateCmsDiscuss 创建CmsDiscuss
func (m *CmsDiscussApi) CreateCmsDiscuss(c *gin.Context) {
	var dataObj business.CmsDiscuss
	_ = c.ShouldBindJSON(&dataObj)
	
	if err := gvalid.CheckStruct(c,dataObj, nil); err != nil {
		m.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	} 
    
	dataObj.UserId = utils.GetUserID(c)
	if id,err := bizSev.GetCmsDiscussSev().Create(c, dataObj); err != nil {
        global.LOG.Error("创建失败!", zap.Any("err", err))
		m.FailWithMessage("创建失败,"+err.Error(), c) 
	} else {
	    idResp := &response.IdResp{Id: id}
		m.OkWithData(idResp, c)
	}
}

// DeleteCmsDiscuss 删除CmsDiscuss

func (m *CmsDiscussApi) DeleteCmsDiscuss(c *gin.Context) {
	var cmsDiscuss business.CmsDiscuss
	_ = c.ShouldBindJSON(&cmsDiscuss)
	if err := bizSev.GetCmsDiscussSev().Delete(c, cmsDiscuss); err != nil {
        global.LOG.Error("删除失败!", zap.Any("err", err))
		m.FailWithMessage("删除失败,"+err.Error(), c) 
	} else {
		m.OkWithMessage("删除成功", c)
	}
}

// DeleteCmsDiscussByIds 批量删除CmsDiscuss
func (m *CmsDiscussApi) DeleteCmsDiscussByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := bizSev.GetCmsDiscussSev().DeleteByIds(c, IDS); err != nil {
        global.LOG.Error("删除失败!", zap.Any("err", err))
		m.FailWithMessage("删除失败,"+err.Error(), c) 
	} else {
		m.OkWithMessage("删除成功", c)
	}
}

// UpdateCmsDiscuss 更新CmsDiscuss
func (m *CmsDiscussApi) UpdateCmsDiscuss(c *gin.Context) {
	var dataObj business.CmsDiscuss
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		m.FailWithMessage("更新失败,"+err.FirstString(), c)
		return
	}

	if err := bizSev.GetCmsDiscussSev().Update(c, dataObj); err != nil {
        global.LOG.Error("更新失败!", zap.Any("err", err))
		m.FailWithMessage("更新失败,"+err.Error(), c) 
	} else {
		m.OkWithMessage("更新成功", c)
	}
}

// FindCmsDiscuss 用id查询CmsDiscuss
func (m *CmsDiscussApi) FindCmsDiscuss(c *gin.Context) {
	var cmsDiscuss business.CmsDiscuss
	_ = c.ShouldBindQuery(&cmsDiscuss) 	 
    data,err:= bizSev.GetCmsDiscussSev().Get(c, cmsDiscuss.Id,""); 
	if err != nil { 
        global.LOG.Error("查询失败!", zap.Any("err", err))
		m.FailWithMessage("查询失败,"+err.Error(), c) 
	} else {
		m.OkWithData(data, c)  
	}
}

// GetCmsDiscussList 分页获取CmsDiscuss列表
func (m *CmsDiscussApi) GetCmsDiscussList(c *gin.Context) {
	var pageInfo business.CmsDiscussSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if  list, total, err := bizSev.GetCmsDiscussSev().GetList(c, pageInfo,business.Field_CmsDiscuss_mini); err != nil {
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
func (m *CmsDiscussApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "cms_discuss" 
	if err := commSev.GetCommonDbSev().QuickEdit(c, quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		m.FailWithMessage("更新失败,"+err.Error(), c) 
	} else {
		m.OkWithMessage("更新成功", c)
	}
}


// excelList 分页导出excel CmsDiscuss列表
func (m *CmsDiscussApi) ExcelList(c *gin.Context) {	 
	var pageInfo business.CmsDiscussSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list,_,err:= bizSev.GetCmsDiscussSev().GetList(c, pageInfo,""); err != nil {
	    global.LOG.Error("获取失败!", zap.Any("err", err))
        m.FailWithMessage("获取失败,"+err.Error(), c) 
    } else {
        if len(list) == 0 {
			m.FailWithMessage("没有数据", c)
		} else { 
			sheetFields := []string{}  
					sheetFields = append(sheetFields, "父ID")  
					sheetFields = append(sheetFields, "id")  
					sheetFields = append(sheetFields, "用户id")  
					sheetFields = append(sheetFields, "用户At")  
					sheetFields = append(sheetFields, "标题")  
					sheetFields = append(sheetFields, "内容")  
					sheetFields = append(sheetFields, "图片")  
					sheetFields = append(sheetFields, "总赞")  
					sheetFields = append(sheetFields, "总踩")  
					sheetFields = append(sheetFields, "状态") 

			excel := excelize.NewFile()
			excel.SetSheetRow("Sheet1", "A1", &sheetFields)
			for i, v := range list {
				axis := fmt.Sprintf("A%d", i+2)
				var arr = []interface{}{} 
				arr = append(arr, v.Pid) 
				arr = append(arr, v.ArtId) 
				arr = append(arr, v.UserId) 
				arr = append(arr, v.UserIdAt) 
				arr = append(arr, v.Title) 
				arr = append(arr, v.Detail) 
				arr = append(arr, v.FileList) 
				arr = append(arr, v.TotalGood) 
				arr = append(arr, v.TotalPoor) 
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


 