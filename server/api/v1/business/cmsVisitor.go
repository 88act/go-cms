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

type CmsVisitorApi struct {
	commSev.BaseApi
} 

// CreateCmsVisitor 创建CmsVisitor
func (m *CmsVisitorApi) CreateCmsVisitor(c *gin.Context) {
	var dataObj business.CmsVisitor
	_ = c.ShouldBindJSON(&dataObj)
	
	if err := gvalid.CheckStruct(c,dataObj, nil); err != nil {
		m.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	} 
    
	dataObj.UserId = utils.GetUserID(c)
	if id,err := bizSev.GetCmsVisitorSev().Create(c, dataObj); err != nil {
        global.LOG.Error("创建失败!", zap.Any("err", err))
		m.FailWithMessage("创建失败,"+err.Error(), c) 
	} else {
	    idResp := &response.IdResp{Id: id}
		m.OkWithData(idResp, c)
	}
}

// DeleteCmsVisitor 删除CmsVisitor

func (m *CmsVisitorApi) DeleteCmsVisitor(c *gin.Context) {
	var cmsVisitor business.CmsVisitor
	_ = c.ShouldBindJSON(&cmsVisitor)
	if err := bizSev.GetCmsVisitorSev().Delete(c, cmsVisitor); err != nil {
        global.LOG.Error("删除失败!", zap.Any("err", err))
		m.FailWithMessage("删除失败,"+err.Error(), c) 
	} else {
		m.OkWithMessage("删除成功", c)
	}
}

// DeleteCmsVisitorByIds 批量删除CmsVisitor
func (m *CmsVisitorApi) DeleteCmsVisitorByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := bizSev.GetCmsVisitorSev().DeleteByIds(c, IDS); err != nil {
        global.LOG.Error("删除失败!", zap.Any("err", err))
		m.FailWithMessage("删除失败,"+err.Error(), c) 
	} else {
		m.OkWithMessage("删除成功", c)
	}
}

// UpdateCmsVisitor 更新CmsVisitor
func (m *CmsVisitorApi) UpdateCmsVisitor(c *gin.Context) {
	var dataObj business.CmsVisitor
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		m.FailWithMessage("更新失败,"+err.FirstString(), c)
		return
	}

	if err := bizSev.GetCmsVisitorSev().Update(c, dataObj); err != nil {
        global.LOG.Error("更新失败!", zap.Any("err", err))
		m.FailWithMessage("更新失败,"+err.Error(), c) 
	} else {
		m.OkWithMessage("更新成功", c)
	}
}

// FindCmsVisitor 用id查询CmsVisitor
func (m *CmsVisitorApi) FindCmsVisitor(c *gin.Context) {
	var cmsVisitor business.CmsVisitor
	_ = c.ShouldBindQuery(&cmsVisitor) 	 
    data,err:= bizSev.GetCmsVisitorSev().Get(c, cmsVisitor.Id,""); 
	if err != nil { 
        global.LOG.Error("查询失败!", zap.Any("err", err))
		m.FailWithMessage("查询失败,"+err.Error(), c) 
	} else {
		m.OkWithData(data, c)  
	}
}

// GetCmsVisitorList 分页获取CmsVisitor列表
func (m *CmsVisitorApi) GetCmsVisitorList(c *gin.Context) {
	var pageInfo business.CmsVisitorSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if  list, total, err := bizSev.GetCmsVisitorSev().GetList(c, pageInfo,business.Field_CmsVisitor_mini); err != nil {
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
func (m *CmsVisitorApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "cms_visitor" 
	if err := commSev.GetCommonDbSev().QuickEdit(c, quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		m.FailWithMessage("更新失败,"+err.Error(), c) 
	} else {
		m.OkWithMessage("更新成功", c)
	}
}


// excelList 分页导出excel CmsVisitor列表
func (m *CmsVisitorApi) ExcelList(c *gin.Context) {	 
	var pageInfo business.CmsVisitorSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list,_,err:= bizSev.GetCmsVisitorSev().GetList(c, pageInfo,""); err != nil {
	    global.LOG.Error("获取失败!", zap.Any("err", err))
        m.FailWithMessage("获取失败,"+err.Error(), c) 
    } else {
        if len(list) == 0 {
			m.FailWithMessage("没有数据", c)
		} else { 
			sheetFields := []string{}  
					sheetFields = append(sheetFields, "用户id")  
					sheetFields = append(sheetFields, "类型")  
					sheetFields = append(sheetFields, "对象Id")  
					sheetFields = append(sheetFields, "平台")  
					sheetFields = append(sheetFields, "客户端ip")  
					sheetFields = append(sheetFields, "状态") 

			excel := excelize.NewFile()
			excel.SetSheetRow("Sheet1", "A1", &sheetFields)
			for i, v := range list {
				axis := fmt.Sprintf("A%d", i+2)
				var arr = []interface{}{} 
				arr = append(arr, v.UserId) 
				arr = append(arr, v.Type) 
				arr = append(arr, v.ObjId) 
				arr = append(arr, v.PlatType) 
				arr = append(arr, v.ClientIp) 
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


 