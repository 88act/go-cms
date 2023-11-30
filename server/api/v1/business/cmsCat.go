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

type CmsCatApi struct {
	commSev.BaseApi
} 

// CreateCmsCat 创建CmsCat
func (m *CmsCatApi) CreateCmsCat(c *gin.Context) {
	var dataObj business.CmsCat
	_ = c.ShouldBindJSON(&dataObj)
	
	if err := gvalid.CheckStruct(c,dataObj, nil); err != nil {
		m.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	} 
    
	dataObj.UserId = utils.GetUserID(c)
	if id,err := bizSev.GetCmsCatSev().Create(c, dataObj); err != nil {
        global.LOG.Error("创建失败!", zap.Any("err", err))
		m.FailWithMessage("创建失败,"+err.Error(), c) 
	} else {
	    idResp := &response.IdResp{Id: id}
		m.OkWithData(idResp, c)
	}
}

// DeleteCmsCat 删除CmsCat

func (m *CmsCatApi) DeleteCmsCat(c *gin.Context) {
	var cmsCat business.CmsCat
	_ = c.ShouldBindJSON(&cmsCat)
	if err := bizSev.GetCmsCatSev().Delete(c, cmsCat); err != nil {
        global.LOG.Error("删除失败!", zap.Any("err", err))
		m.FailWithMessage("删除失败,"+err.Error(), c) 
	} else {
		m.OkWithMessage("删除成功", c)
	}
}

// DeleteCmsCatByIds 批量删除CmsCat
func (m *CmsCatApi) DeleteCmsCatByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := bizSev.GetCmsCatSev().DeleteByIds(c, IDS); err != nil {
        global.LOG.Error("删除失败!", zap.Any("err", err))
		m.FailWithMessage("删除失败,"+err.Error(), c) 
	} else {
		m.OkWithMessage("删除成功", c)
	}
}

// UpdateCmsCat 更新CmsCat
func (m *CmsCatApi) UpdateCmsCat(c *gin.Context) {
	var dataObj business.CmsCat
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		m.FailWithMessage("更新失败,"+err.FirstString(), c)
		return
	}

	if err := bizSev.GetCmsCatSev().Update(c, dataObj); err != nil {
        global.LOG.Error("更新失败!", zap.Any("err", err))
		m.FailWithMessage("更新失败,"+err.Error(), c) 
	} else {
		m.OkWithMessage("更新成功", c)
	}
}

// FindCmsCat 用id查询CmsCat
func (m *CmsCatApi) FindCmsCat(c *gin.Context) {
	var cmsCat business.CmsCat
	_ = c.ShouldBindQuery(&cmsCat) 	 
    data,err:= bizSev.GetCmsCatSev().Get(c, cmsCat.Id,""); 
	if err != nil { 
        global.LOG.Error("查询失败!", zap.Any("err", err))
		m.FailWithMessage("查询失败,"+err.Error(), c) 
	} else {
		m.OkWithData(data, c)  
	}
}

// GetCmsCatList 分页获取CmsCat列表
func (m *CmsCatApi) GetCmsCatList(c *gin.Context) {
	var pageInfo business.CmsCatSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if  list, total, err := bizSev.GetCmsCatSev().GetList(c, pageInfo,business.Field_CmsCat_mini); err != nil {
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
func (m *CmsCatApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "cms_cat" 
	if err := commSev.GetCommonDbSev().QuickEdit(c, quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		m.FailWithMessage("更新失败,"+err.Error(), c) 
	} else {
		m.OkWithMessage("更新成功", c)
	}
}


// excelList 分页导出excel CmsCat列表
func (m *CmsCatApi) ExcelList(c *gin.Context) {	 
	var pageInfo business.CmsCatSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list,_,err:= bizSev.GetCmsCatSev().GetList(c, pageInfo,""); err != nil {
	    global.LOG.Error("获取失败!", zap.Any("err", err))
        m.FailWithMessage("获取失败,"+err.Error(), c) 
    } else {
        if len(list) == 0 {
			m.FailWithMessage("没有数据", c)
		} else { 
			sheetFields := []string{}  
					sheetFields = append(sheetFields, "父ID")  
					sheetFields = append(sheetFields, "用户id")  
					sheetFields = append(sheetFields, "系统分类")  
					sheetFields = append(sheetFields, "专题id")  
					sheetFields = append(sheetFields, "栏目类型")  
					sheetFields = append(sheetFields, "标题")  
					sheetFields = append(sheetFields, "摘要")  
					sheetFields = append(sheetFields, "标签列表")  
					sheetFields = append(sheetFields, "插图")  
					sheetFields = append(sheetFields, "媒体列表")  
					sheetFields = append(sheetFields, "置顶")  
					sheetFields = append(sheetFields, "是否导航")  
					sheetFields = append(sheetFields, "排序")  
					sheetFields = append(sheetFields, "状态") 

			excel := excelize.NewFile()
			excel.SetSheetRow("Sheet1", "A1", &sheetFields)
			for i, v := range list {
				axis := fmt.Sprintf("A%d", i+2)
				var arr = []interface{}{} 
				arr = append(arr, v.Pid) 
				arr = append(arr, v.UserId) 
				arr = append(arr, v.BeSys) 
				arr = append(arr, v.ObjId) 
				arr = append(arr, v.Type) 
				arr = append(arr, v.Title) 
				arr = append(arr, v.Desc) 
				arr = append(arr, v.TagList) 
				arr = append(arr, v.Image) 
				arr = append(arr, v.FileList) 
				arr = append(arr, v.BeTop) 
				arr = append(arr, v.BeNav) 
				arr = append(arr, v.Sort) 
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


 