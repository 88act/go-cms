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

type BbPiServicePointApi struct {
}

 

// CreateBbPiServicePoint 创建BbPiServicePoint
// @Tags BbPiServicePoint
// @Summary 创建BbPiServicePoint
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.BbPiServicePoint true "创建BbPiServicePoint"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiServicePoint/createBbPiServicePoint [post]
func (bbPiServicePointApi *BbPiServicePointApi) CreateBbPiServicePoint(c *gin.Context) {
	var dataObj business.BbPiServicePoint
	_ = c.ShouldBindJSON(&dataObj)
	
	if err := gvalid.CheckStruct(c,dataObj, nil); err != nil {
		response.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	}

 
	if id,err := bizSev.GetBbPiServicePointSev().Create(dataObj); err != nil {
        global.LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
	    idResp := &response.IdResp{Id: id}
		response.OkWithData(idResp, c)
	}
}

// DeleteBbPiServicePoint 删除BbPiServicePoint
// @Tags BbPiServicePoint
// @Summary 删除BbPiServicePoint
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.BbPiServicePoint true "删除BbPiServicePoint"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bbPiServicePoint/deleteBbPiServicePoint [delete]
func (bbPiServicePointApi *BbPiServicePointApi) DeleteBbPiServicePoint(c *gin.Context) {
	var bbPiServicePoint business.BbPiServicePoint
	_ = c.ShouldBindJSON(&bbPiServicePoint)
	if err := bizSev.GetBbPiServicePointSev().Delete(bbPiServicePoint); err != nil {
        global.LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBbPiServicePointByIds 批量删除BbPiServicePoint
// @Tags BbPiServicePoint
// @Summary 批量删除BbPiServicePoint
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BbPiServicePoint"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /bbPiServicePoint/deleteBbPiServicePointByIds [delete]
func (bbPiServicePointApi *BbPiServicePointApi) DeleteBbPiServicePointByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := bizSev.GetBbPiServicePointSev().DeleteByIds(IDS); err != nil {
        global.LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBbPiServicePoint 更新BbPiServicePoint
// @Tags BbPiServicePoint
// @Summary 更新BbPiServicePoint
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.BbPiServicePoint true "更新BbPiServicePoint"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /bbPiServicePoint/updateBbPiServicePoint [put]
func (bbPiServicePointApi *BbPiServicePointApi) UpdateBbPiServicePoint(c *gin.Context) {
	var dataObj business.BbPiServicePoint
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		response.FailWithMessage("更新失败,"+err.FirstString(), c)
		return
	}

	if err := bizSev.GetBbPiServicePointSev().Update(dataObj); err != nil {
        global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBbPiServicePoint 用id查询BbPiServicePoint
// @Tags BbPiServicePoint
// @Summary 用id查询BbPiServicePoint
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.BbPiServicePoint true "用id查询BbPiServicePoint"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /bbPiServicePoint/findBbPiServicePoint [get]
func (bbPiServicePointApi *BbPiServicePointApi) FindBbPiServicePoint(c *gin.Context) {
	var bbPiServicePoint business.BbPiServicePoint
	_ = c.ShouldBindQuery(&bbPiServicePoint) 
	 rebbPiServicePoint,err:= bizSev.GetBbPiServicePointSev().Get(bbPiServicePoint.ID,""); 
	 if errors.Is(err, gorm.ErrRecordNotFound) { 
		response.OkWithData(gin.H{"bbPiServicePoint": nil}, c)
	} else if err != nil { 
        global.LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else { 
		response.OkWithData(gin.H{"bbPiServicePoint": rebbPiServicePoint}, c)
	}
}

// GetBbPiServicePointList 分页获取BbPiServicePoint列表
// @Tags BbPiServicePoint
// @Summary 分页获取BbPiServicePoint列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.BbPiServicePointSearch true "分页获取BbPiServicePoint列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiServicePoint/getBbPiServicePointList [get]
func (bbPiServicePointApi *BbPiServicePointApi) GetBbPiServicePointList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")

	var pageInfo bizReq.BbPiServicePointSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if  list, total, err := bizSev.GetBbPiServicePointSev().GetList(pageInfo,createdAtBetween,""); err != nil {
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
// @Param data body business.BbPiServicePoint true "快速更新BbPiServicePoint" 
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router  /bbPiServicePoint/quickEdit [post] 
func (bbPiServicePointApi *BbPiServicePointApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "bb_pi_service_point" 
	if err := commSev.GetCommonDbSev().QuickEdit(quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}


// excelList 分页导出excel BbPiServicePoint列表
// @Tags BbPiServicePoint
// @Summary 分页导出excel BbPiServicePoint列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.BbPiServicePointSearch true "分页导出excel BbPiServicePoint列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiServicePoint/excelList [get]
func (bbPiServicePointApi *BbPiServicePointApi) ExcelList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")
	var pageInfo bizReq.BbPiServicePointSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list,_,err:= bizSev.GetBbPiServicePointSev().GetListAll(pageInfo,createdAtBetween,""); err != nil {
	    global.LOG.Error("获取失败!", zap.Any("err", err))
        response.FailWithMessage("获取失败", c)
    } else {
        if len(list) == 0 {
			response.FailWithMessage("没有数据", c)
		} else { 
			sheetFields := []string{}  
					sheetFields = append(sheetFields, "上传")  
					sheetFields = append(sheetFields, "机构标识")  
					sheetFields = append(sheetFields, "统一社会信用代码")  
					sheetFields = append(sheetFields, "服务网点代码")  
					sheetFields = append(sheetFields, "服务点名称")  
					sheetFields = append(sheetFields, "行政区划代码")  
					sheetFields = append(sheetFields, "服务点类型")  
					sheetFields = append(sheetFields, "服务点成立日期")  
					sheetFields = append(sheetFields, "单位隶属关系代码")  
					sheetFields = append(sheetFields, "服务点分类管理类别代码")  
					sheetFields = append(sheetFields, "服务点分类代码")  
					sheetFields = append(sheetFields, "经济类型代码")  
					sheetFields = append(sheetFields, "地址")  
					sheetFields = append(sheetFields, "服务点医院级别")  
					sheetFields = append(sheetFields, "服务点医院等级")  
					sheetFields = append(sheetFields, "许可证号码")  
					sheetFields = append(sheetFields, "许可项目名称")  
					sheetFields = append(sheetFields, "许可证有效期")  
					sheetFields = append(sheetFields, "开办资金金额数")  
					sheetFields = append(sheetFields, "法人代表")  
					sheetFields = append(sheetFields, "服务点地方标志")  
					sheetFields = append(sheetFields, "是否分支机构") 

			excel := excelize.NewFile()
			excel.SetSheetRow("Sheet1", "A1", &sheetFields)
			for i, v := range list {
				axis := fmt.Sprintf("A%d", i+2)
				var arr = []interface{}{}
				if v.Status == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.Status)
				}
				arr = append(arr, v.Jgdm)
				arr = append(arr, v.Zzjgdm)
				arr = append(arr, v.Fwwddm)
				arr = append(arr, v.Fwdmc)
				arr = append(arr, v.Xzqhdm)
				arr = append(arr, v.Fwdlx)
				arr = append(arr, v.Fwdclrq)
				arr = append(arr, v.Dwlsgxdm)
				arr = append(arr, v.Fwdflgllbdm)
				arr = append(arr, v.Fwdfldm)
				arr = append(arr, v.Jjlxdm)
				arr = append(arr, v.Dz)
				arr = append(arr, v.Fwdyyjb)
				arr = append(arr, v.Fwdyydj)
				arr = append(arr, v.Xkzhm)
				arr = append(arr, v.Xkxmmc)
				arr = append(arr, v.Xkzyxq)
				arr = append(arr, v.Kbzjjes)
				arr = append(arr, v.Frdb)
				arr = append(arr, v.Fwdszdmzzzdfbz)
				arr = append(arr, v.Sffzjg)   
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


 