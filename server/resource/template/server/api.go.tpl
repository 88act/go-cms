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

type {{.StructName}}Api struct {
	commSev.BaseApi
} 

// Create{{.StructName}} 创建{{.StructName}}
func (m *{{.StructName}}Api) Create{{.StructName}}(c *gin.Context) {
	var dataObj business.{{.StructName}}
	_ = c.ShouldBindJSON(&dataObj)
	
	if err := gvalid.CheckStruct(c,dataObj, nil); err != nil {
		m.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	} 
    
	dataObj.UserId = utils.GetUserID(c)
	if id,err := bizSev.Get{{.StructName}}Sev().Create(c, dataObj); err != nil {
        global.LOG.Error("创建失败!", zap.Any("err", err))
		m.FailWithMessage("创建失败,"+err.Error(), c) 
	} else {
	    idResp := &response.IdResp{Id: id}
		m.OkWithData(idResp, c)
	}
}

// Delete{{.StructName}} 删除{{.StructName}}

func (m *{{.StructName}}Api) Delete{{.StructName}}(c *gin.Context) {
	var {{.Abbreviation}} business.{{.StructName}}
	_ = c.ShouldBindJSON(&{{.Abbreviation}})
	if err := bizSev.Get{{.StructName}}Sev().Delete(c, {{.Abbreviation}}); err != nil {
        global.LOG.Error("删除失败!", zap.Any("err", err))
		m.FailWithMessage("删除失败,"+err.Error(), c) 
	} else {
		m.OkWithMessage("删除成功", c)
	}
}

// Delete{{.StructName}}ByIds 批量删除{{.StructName}}
func (m *{{.StructName}}Api) Delete{{.StructName}}ByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := bizSev.Get{{.StructName}}Sev().DeleteByIds(c, IDS); err != nil {
        global.LOG.Error("删除失败!", zap.Any("err", err))
		m.FailWithMessage("删除失败,"+err.Error(), c) 
	} else {
		m.OkWithMessage("删除成功", c)
	}
}

// Update{{.StructName}} 更新{{.StructName}}
func (m *{{.StructName}}Api) Update{{.StructName}}(c *gin.Context) {
	var dataObj business.{{.StructName}}
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		m.FailWithMessage("更新失败,"+err.FirstString(), c)
		return
	}

	if err := bizSev.Get{{.StructName}}Sev().Update(c, dataObj); err != nil {
        global.LOG.Error("更新失败!", zap.Any("err", err))
		m.FailWithMessage("更新失败,"+err.Error(), c) 
	} else {
		m.OkWithMessage("更新成功", c)
	}
}

// Find{{.StructName}} 用id查询{{.StructName}}
func (m *{{.StructName}}Api) Find{{.StructName}}(c *gin.Context) {
	var {{.Abbreviation}} business.{{.StructName}}
	_ = c.ShouldBindQuery(&{{.Abbreviation}}) 	 
    data,err:= bizSev.Get{{.StructName}}Sev().Get(c, {{.Abbreviation}}.Id,""); 
	if err != nil { 
        global.LOG.Error("查询失败!", zap.Any("err", err))
		m.FailWithMessage("查询失败,"+err.Error(), c) 
	} else {
		m.OkWithData(data, c)  
	}
}

// Get{{.StructName}}List 分页获取{{.StructName}}列表
func (m *{{.StructName}}Api) Get{{.StructName}}List(c *gin.Context) {
	var pageInfo business.{{.StructName}}Search
	_ = c.ShouldBindQuery(&pageInfo)
	if  list, total, err := bizSev.Get{{.StructName}}Sev().GetList(c, pageInfo,business.Field_{{.StructName}}_mini); err != nil {
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
func (m *{{.StructName}}Api) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "{{.TableName}}" 
	if err := commSev.GetCommonDbSev().QuickEdit(c, quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		m.FailWithMessage("更新失败,"+err.Error(), c) 
	} else {
		m.OkWithMessage("更新成功", c)
	}
}


// excelList 分页导出excel {{.StructName}}列表
func (m *{{.StructName}}Api) ExcelList(c *gin.Context) {	 
	var pageInfo business.{{.StructName}}Search
	_ = c.ShouldBindQuery(&pageInfo)
	if list,_,err:= bizSev.Get{{.StructName}}Sev().GetList(c, pageInfo,""); err != nil {
	    global.LOG.Error("获取失败!", zap.Any("err", err))
        m.FailWithMessage("获取失败,"+err.Error(), c) 
    } else {
        if len(list) == 0 {
			m.FailWithMessage("没有数据", c)
		} else { 
			sheetFields := []string{}  
			{{- range .Fields}}  
					sheetFields = append(sheetFields, "{{.FieldDesc}}") 
			{{- end }} 

			excel := excelize.NewFile()
			excel.SetSheetRow("Sheet1", "A1", &sheetFields)
			for i, v := range list {
				axis := fmt.Sprintf("A%d", i+2)
				var arr = []interface{}{}
			{{- range .Fields}} 
				arr = append(arr, v.{{.FieldName}})  
			{{- end }}   
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


 