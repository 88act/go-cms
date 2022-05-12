package business

import (
 "errors"
	"fmt"
	"go-cms/global"
	"go-cms/model/business"
	bizReq "go-cms/model/business/request"
	"go-cms/model/common/request"
	"go-cms/model/common/response"
	bizSev "go-cms/service/business"
	commSev "go-cms/service/common"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gvalid"
	"github.com/xuri/excelize/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
) 

type ActFavApi struct {
}

 

// CreateActFav 创建ActFav
// @Tags ActFav
// @Summary 创建ActFav
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.ActFav true "创建ActFav"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /actFav/createActFav [post]
func (m *ActFavApi) CreateActFav(c *gin.Context) {
	var dataObj business.ActFav
	_ = c.ShouldBindJSON(&dataObj)
	
	if err := gvalid.CheckStruct(c,dataObj, nil); err != nil {
		response.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	}

 
	if id,err := bizSev.GetActFavSev().Create(dataObj); err != nil {
        global.LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
	    idResp := &response.IdResp{Id: id}
		response.OkWithData(idResp, c)
	}
}

// DeleteActFav 删除ActFav
// @Tags ActFav
// @Summary 删除ActFav
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.ActFav true "删除ActFav"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /actFav/deleteActFav [delete]
func (m *ActFavApi) DeleteActFav(c *gin.Context) {
	var actFav business.ActFav
	_ = c.ShouldBindJSON(&actFav)
	if err := bizSev.GetActFavSev().Delete(actFav); err != nil {
        global.LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteActFavByIds 批量删除ActFav
// @Tags ActFav
// @Summary 批量删除ActFav
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ActFav"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /actFav/deleteActFavByIds [delete]
func (m *ActFavApi) DeleteActFavByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := bizSev.GetActFavSev().DeleteByIds(IDS); err != nil {
        global.LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateActFav 更新ActFav
// @Tags ActFav
// @Summary 更新ActFav
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.ActFav true "更新ActFav"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /actFav/updateActFav [put]
func (m *ActFavApi) UpdateActFav(c *gin.Context) {
	var dataObj business.ActFav
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		response.FailWithMessage("更新失败,"+err.FirstString(), c)
		return
	}

	if err := bizSev.GetActFavSev().Update(dataObj); err != nil {
        global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindActFav 用id查询ActFav
// @Tags ActFav
// @Summary 用id查询ActFav
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.ActFav true "用id查询ActFav"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /actFav/findActFav [get]
func (m *ActFavApi) FindActFav(c *gin.Context) {
	var actFav business.ActFav
	_ = c.ShouldBindQuery(&actFav) 
	 reactFav,err:= bizSev.GetActFavSev().Get(actFav.ID,""); 
	 if errors.Is(err, gorm.ErrRecordNotFound) { 
		response.OkWithData(gin.H{"actFav": nil}, c)
	} else if err != nil { 
        global.LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else { 
		response.OkWithData(gin.H{"actFav": reactFav}, c)
	}
}

// GetActFavList 分页获取ActFav列表
// @Tags ActFav
// @Summary 分页获取ActFav列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.ActFavSearch true "分页获取ActFav列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /actFav/getActFavList [get]
func (m *ActFavApi) GetActFavList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")

	var pageInfo bizReq.ActFavSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if  list, total, err := bizSev.GetActFavSev().GetList(pageInfo,createdAtBetween,""); err != nil {
	    global.LOG.Error("获取失败!", zap.Any("err", err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}



// QuickEdit 快速更新
// @Tags QuickEdit
// @Summary 快速更新
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.ActFav true "快速更新ActFav" 
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router  /actFav/quickEdit [post] 
func (m *ActFavApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "act_fav" 
	if err := commSev.GetCommonDbSev().QuickEdit(quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}


// excelList 分页导出excel ActFav列表
// @Tags ActFav
// @Summary 分页导出excel ActFav列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.ActFavSearch true "分页导出excel ActFav列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /actFav/excelList [get]
func (m *ActFavApi) ExcelList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")
	var pageInfo bizReq.ActFavSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list,_,err:= bizSev.GetActFavSev().GetListAll(pageInfo,createdAtBetween,""); err != nil {
	    global.LOG.Error("获取失败!", zap.Any("err", err))
        response.FailWithMessage("获取失败", c)
    } else {
        if len(list) == 0 {
			response.FailWithMessage("没有数据", c)
		} else { 
			sheetFields := []string{}  
					sheetFields = append(sheetFields, "用户id")  
					sheetFields = append(sheetFields, "收藏类型")  
					sheetFields = append(sheetFields, "对象id")  
					sheetFields = append(sheetFields, "状态") 

			excel := excelize.NewFile()
			excel.SetSheetRow("Sheet1", "A1", &sheetFields)
			for i, v := range list {
				axis := fmt.Sprintf("A%d", i+2)
				var arr = []interface{}{}
				if v.UserId == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.UserId)
				}
				if v.Type == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.Type)
				}
				if v.ObjId == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.ObjId)
				}
				if v.Status == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.Status)
				}   
			    excel.SetSheetRow("Sheet1", axis,&arr)  
			}
			filename := fmt.Sprintf("ecl%d.xlsx", time.Now().Unix())
			filePath := global.CONFIG.Local.BasePath + global.CONFIG.Local.Path + "/excel/" + filename
			url := global.CONFIG.Local.BaseUrl + global.CONFIG.Local.Path + "/excel/" + filename
			err := excel.SaveAs(filePath)
			if err != nil {
				global.LOG.Error(err.Error())
				response.FailWithMessage("获取失败", c)
			} else {
				resData := map[string]string{"url": url, "filename": filename} 
				response.OkWithData(resData, c)
			} 
		}
    }
}


 