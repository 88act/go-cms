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

type CmsCatArtApi struct {
	commSev.BaseApi
} 

// CreateCmsCatArt 创建CmsCatArt
func (m *CmsCatArtApi) CreateCmsCatArt(c *gin.Context) {
	var dataObj business.CmsCatArt
	_ = c.ShouldBindJSON(&dataObj)
	
	if err := gvalid.CheckStruct(c,dataObj, nil); err != nil {
		m.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	} 
    
	dataObj.UserId = utils.GetUserID(c)
	if id,err := bizSev.GetCmsCatArtSev().Create(c, dataObj); err != nil {
        global.LOG.Error("创建失败!", zap.Any("err", err))
		m.FailWithMessage("创建失败,"+err.Error(), c) 
	} else {
	    idResp := &response.IdResp{Id: id}
		m.OkWithData(idResp, c)
	}
}

// DeleteCmsCatArt 删除CmsCatArt

func (m *CmsCatArtApi) DeleteCmsCatArt(c *gin.Context) {
	var cmsCatArt business.CmsCatArt
	_ = c.ShouldBindJSON(&cmsCatArt)
	if err := bizSev.GetCmsCatArtSev().Delete(c, cmsCatArt); err != nil {
        global.LOG.Error("删除失败!", zap.Any("err", err))
		m.FailWithMessage("删除失败,"+err.Error(), c) 
	} else {
		m.OkWithMessage("删除成功", c)
	}
}

// DeleteCmsCatArtByIds 批量删除CmsCatArt
func (m *CmsCatArtApi) DeleteCmsCatArtByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := bizSev.GetCmsCatArtSev().DeleteByIds(c, IDS); err != nil {
        global.LOG.Error("删除失败!", zap.Any("err", err))
		m.FailWithMessage("删除失败,"+err.Error(), c) 
	} else {
		m.OkWithMessage("删除成功", c)
	}
}

// UpdateCmsCatArt 更新CmsCatArt
func (m *CmsCatArtApi) UpdateCmsCatArt(c *gin.Context) {
	var dataObj business.CmsCatArt
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		m.FailWithMessage("更新失败,"+err.FirstString(), c)
		return
	}

	if err := bizSev.GetCmsCatArtSev().Update(c, dataObj); err != nil {
        global.LOG.Error("更新失败!", zap.Any("err", err))
		m.FailWithMessage("更新失败,"+err.Error(), c) 
	} else {
		m.OkWithMessage("更新成功", c)
	}
}

// FindCmsCatArt 用id查询CmsCatArt
func (m *CmsCatArtApi) FindCmsCatArt(c *gin.Context) {
	var cmsCatArt business.CmsCatArt
	_ = c.ShouldBindQuery(&cmsCatArt) 	 
    data,err:= bizSev.GetCmsCatArtSev().Get(c, cmsCatArt.Id,""); 
	if err != nil { 
        global.LOG.Error("查询失败!", zap.Any("err", err))
		m.FailWithMessage("查询失败,"+err.Error(), c) 
	} else {
		m.OkWithData(data, c)  
	}
}

// GetCmsCatArtList 分页获取CmsCatArt列表
func (m *CmsCatArtApi) GetCmsCatArtList(c *gin.Context) {
	var pageInfo business.CmsCatArtSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if  list, total, err := bizSev.GetCmsCatArtSev().GetList(c, pageInfo,business.Field_CmsCatArt_mini); err != nil {
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
func (m *CmsCatArtApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "cms_cat_art" 
	if err := commSev.GetCommonDbSev().QuickEdit(c, quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		m.FailWithMessage("更新失败,"+err.Error(), c) 
	} else {
		m.OkWithMessage("更新成功", c)
	}
}


// excelList 分页导出excel CmsCatArt列表
func (m *CmsCatArtApi) ExcelList(c *gin.Context) {	 
	var pageInfo business.CmsCatArtSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list,_,err:= bizSev.GetCmsCatArtSev().GetList(c, pageInfo,""); err != nil {
	    global.LOG.Error("获取失败!", zap.Any("err", err))
        m.FailWithMessage("获取失败,"+err.Error(), c) 
    } else {
        if len(list) == 0 {
			m.FailWithMessage("没有数据", c)
		} else { 
			sheetFields := []string{}  
					sheetFields = append(sheetFields, "用户id")  
					sheetFields = append(sheetFields, "栏目id")  
					sheetFields = append(sheetFields, "文章id")  
					sheetFields = append(sheetFields, "状态") 

			excel := excelize.NewFile()
			excel.SetSheetRow("Sheet1", "A1", &sheetFields)
			for i, v := range list {
				axis := fmt.Sprintf("A%d", i+2)
				var arr = []interface{}{} 
				arr = append(arr, v.UserId) 
				arr = append(arr, v.CatId) 
				arr = append(arr, v.ArtId) 
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


 