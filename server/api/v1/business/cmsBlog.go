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

type CmsBlogApi struct {
	commSev.BaseApi
} 

// CreateCmsBlog 创建CmsBlog
func (m *CmsBlogApi) CreateCmsBlog(c *gin.Context) {
	var dataObj business.CmsBlog
	_ = c.ShouldBindJSON(&dataObj)
	
	if err := gvalid.CheckStruct(c,dataObj, nil); err != nil {
		m.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	} 
    
	dataObj.UserId = utils.GetUserID(c)
	if id,err := bizSev.GetCmsBlogSev().Create(c, dataObj); err != nil {
        global.LOG.Error("创建失败!", zap.Any("err", err))
		m.FailWithMessage("创建失败,"+err.Error(), c) 
	} else {
	    idResp := &response.IdResp{Id: id}
		m.OkWithData(idResp, c)
	}
}

// DeleteCmsBlog 删除CmsBlog

func (m *CmsBlogApi) DeleteCmsBlog(c *gin.Context) {
	var cmsBlog business.CmsBlog
	_ = c.ShouldBindJSON(&cmsBlog)
	if err := bizSev.GetCmsBlogSev().Delete(c, cmsBlog); err != nil {
        global.LOG.Error("删除失败!", zap.Any("err", err))
		m.FailWithMessage("删除失败,"+err.Error(), c) 
	} else {
		m.OkWithMessage("删除成功", c)
	}
}

// DeleteCmsBlogByIds 批量删除CmsBlog
func (m *CmsBlogApi) DeleteCmsBlogByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := bizSev.GetCmsBlogSev().DeleteByIds(c, IDS); err != nil {
        global.LOG.Error("删除失败!", zap.Any("err", err))
		m.FailWithMessage("删除失败,"+err.Error(), c) 
	} else {
		m.OkWithMessage("删除成功", c)
	}
}

// UpdateCmsBlog 更新CmsBlog
func (m *CmsBlogApi) UpdateCmsBlog(c *gin.Context) {
	var dataObj business.CmsBlog
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		m.FailWithMessage("更新失败,"+err.FirstString(), c)
		return
	}

	if err := bizSev.GetCmsBlogSev().Update(c, dataObj); err != nil {
        global.LOG.Error("更新失败!", zap.Any("err", err))
		m.FailWithMessage("更新失败,"+err.Error(), c) 
	} else {
		m.OkWithMessage("更新成功", c)
	}
}

// FindCmsBlog 用id查询CmsBlog
func (m *CmsBlogApi) FindCmsBlog(c *gin.Context) {
	var cmsBlog business.CmsBlog
	_ = c.ShouldBindQuery(&cmsBlog) 	 
    data,err:= bizSev.GetCmsBlogSev().Get(c, cmsBlog.Id,""); 
	if err != nil { 
        global.LOG.Error("查询失败!", zap.Any("err", err))
		m.FailWithMessage("查询失败,"+err.Error(), c) 
	} else {
		m.OkWithData(data, c)  
	}
}

// GetCmsBlogList 分页获取CmsBlog列表
func (m *CmsBlogApi) GetCmsBlogList(c *gin.Context) {
	var pageInfo business.CmsBlogSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if  list, total, err := bizSev.GetCmsBlogSev().GetList(c, pageInfo,business.Field_CmsBlog_mini); err != nil {
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
func (m *CmsBlogApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "cms_blog" 
	if err := commSev.GetCommonDbSev().QuickEdit(c, quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		m.FailWithMessage("更新失败,"+err.Error(), c) 
	} else {
		m.OkWithMessage("更新成功", c)
	}
}


// excelList 分页导出excel CmsBlog列表
func (m *CmsBlogApi) ExcelList(c *gin.Context) {	 
	var pageInfo business.CmsBlogSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list,_,err:= bizSev.GetCmsBlogSev().GetList(c, pageInfo,""); err != nil {
	    global.LOG.Error("获取失败!", zap.Any("err", err))
        m.FailWithMessage("获取失败,"+err.Error(), c) 
    } else {
        if len(list) == 0 {
			m.FailWithMessage("没有数据", c)
		} else { 
			sheetFields := []string{}  
					sheetFields = append(sheetFields, "用户id")  
					sheetFields = append(sheetFields, "标题")  
					sheetFields = append(sheetFields, "简介")  
					sheetFields = append(sheetFields, "缩略图")  
					sheetFields = append(sheetFields, "文件列表")  
					sheetFields = append(sheetFields, "邮件")  
					sheetFields = append(sheetFields, "地区id")  
					sheetFields = append(sheetFields, "地址")  
					sheetFields = append(sheetFields, "综合指数")  
					sheetFields = append(sheetFields, "总点击")  
					sheetFields = append(sheetFields, "总粉丝")  
					sheetFields = append(sheetFields, "总赞")  
					sheetFields = append(sheetFields, "总踩")  
					sheetFields = append(sheetFields, "状态") 

			excel := excelize.NewFile()
			excel.SetSheetRow("Sheet1", "A1", &sheetFields)
			for i, v := range list {
				axis := fmt.Sprintf("A%d", i+2)
				var arr = []interface{}{} 
				arr = append(arr, v.UserId) 
				arr = append(arr, v.Title) 
				arr = append(arr, v.Desc) 
				arr = append(arr, v.Image) 
				arr = append(arr, v.FileList) 
				arr = append(arr, v.Email) 
				arr = append(arr, v.AreaId) 
				arr = append(arr, v.Address) 
				arr = append(arr, v.TotalWhole) 
				arr = append(arr, v.TotalClick) 
				arr = append(arr, v.TotalFan) 
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


 