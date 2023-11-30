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

type BasicFileCatApi struct {
	commSev.BaseApi
} 

// CreateBasicFileCat 创建BasicFileCat
func (m *BasicFileCatApi) CreateBasicFileCat(c *gin.Context) {
	var dataObj business.BasicFileCat
	_ = c.ShouldBindJSON(&dataObj)
	
	if err := gvalid.CheckStruct(c,dataObj, nil); err != nil {
		m.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	} 
    
	dataObj.UserId = utils.GetUserID(c)
	if id,err := bizSev.GetBasicFileCatSev().Create(c, dataObj); err != nil {
        global.LOG.Error("创建失败!", zap.Any("err", err))
		m.FailWithMessage("创建失败,"+err.Error(), c) 
	} else {
	    idResp := &response.IdResp{Id: id}
		m.OkWithData(idResp, c)
	}
}

// DeleteBasicFileCat 删除BasicFileCat

func (m *BasicFileCatApi) DeleteBasicFileCat(c *gin.Context) {
	var basicFileCat business.BasicFileCat
	_ = c.ShouldBindJSON(&basicFileCat)
	if err := bizSev.GetBasicFileCatSev().Delete(c, basicFileCat); err != nil {
        global.LOG.Error("删除失败!", zap.Any("err", err))
		m.FailWithMessage("删除失败,"+err.Error(), c) 
	} else {
		m.OkWithMessage("删除成功", c)
	}
}

// DeleteBasicFileCatByIds 批量删除BasicFileCat
func (m *BasicFileCatApi) DeleteBasicFileCatByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := bizSev.GetBasicFileCatSev().DeleteByIds(c, IDS); err != nil {
        global.LOG.Error("删除失败!", zap.Any("err", err))
		m.FailWithMessage("删除失败,"+err.Error(), c) 
	} else {
		m.OkWithMessage("删除成功", c)
	}
}

// UpdateBasicFileCat 更新BasicFileCat
func (m *BasicFileCatApi) UpdateBasicFileCat(c *gin.Context) {
	var dataObj business.BasicFileCat
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		m.FailWithMessage("更新失败,"+err.FirstString(), c)
		return
	}

	if err := bizSev.GetBasicFileCatSev().Update(c, dataObj); err != nil {
        global.LOG.Error("更新失败!", zap.Any("err", err))
		m.FailWithMessage("更新失败,"+err.Error(), c) 
	} else {
		m.OkWithMessage("更新成功", c)
	}
}

// FindBasicFileCat 用id查询BasicFileCat
func (m *BasicFileCatApi) FindBasicFileCat(c *gin.Context) {
	var basicFileCat business.BasicFileCat
	_ = c.ShouldBindQuery(&basicFileCat) 	 
    data,err:= bizSev.GetBasicFileCatSev().Get(c, basicFileCat.Id,""); 
	if err != nil { 
        global.LOG.Error("查询失败!", zap.Any("err", err))
		m.FailWithMessage("查询失败,"+err.Error(), c) 
	} else {
		m.OkWithData(data, c)  
	}
}

// GetBasicFileCatList 分页获取BasicFileCat列表
func (m *BasicFileCatApi) GetBasicFileCatList(c *gin.Context) {
	var pageInfo business.BasicFileCatSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if  list, total, err := bizSev.GetBasicFileCatSev().GetList(c, pageInfo,business.Field_BasicFileCat_mini); err != nil {
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
func (m *BasicFileCatApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "basic_file_cat" 
	if err := commSev.GetCommonDbSev().QuickEdit(c, quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		m.FailWithMessage("更新失败,"+err.Error(), c) 
	} else {
		m.OkWithMessage("更新成功", c)
	}
}


// excelList 分页导出excel BasicFileCat列表
func (m *BasicFileCatApi) ExcelList(c *gin.Context) {	 
	var pageInfo business.BasicFileCatSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list,_,err:= bizSev.GetBasicFileCatSev().GetList(c, pageInfo,""); err != nil {
	    global.LOG.Error("获取失败!", zap.Any("err", err))
        m.FailWithMessage("获取失败,"+err.Error(), c) 
    } else {
        if len(list) == 0 {
			m.FailWithMessage("没有数据", c)
		} else { 
			sheetFields := []string{}  
					sheetFields = append(sheetFields, "父ID")  
					sheetFields = append(sheetFields, "用户id")  
					sheetFields = append(sheetFields, "名称")  
					sheetFields = append(sheetFields, "排序")  
					sheetFields = append(sheetFields, "状态") 

			excel := excelize.NewFile()
			excel.SetSheetRow("Sheet1", "A1", &sheetFields)
			for i, v := range list {
				axis := fmt.Sprintf("A%d", i+2)
				var arr = []interface{}{} 
				arr = append(arr, v.Pid) 
				arr = append(arr, v.UserId) 
				arr = append(arr, v.Name) 
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


 