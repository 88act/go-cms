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

type ActShopApi struct {
}

 

// CreateActShop 创建ActShop
// @Tags ActShop
// @Summary 创建ActShop
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.ActShop true "创建ActShop"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /actShop/createActShop [post]
func (m *ActShopApi) CreateActShop(c *gin.Context) {
	var dataObj business.ActShop
	_ = c.ShouldBindJSON(&dataObj)
	
	if err := gvalid.CheckStruct(c,dataObj, nil); err != nil {
		response.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	}

 
	if id,err := bizSev.GetActShopSev().Create(dataObj); err != nil {
        global.LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
	    idResp := &response.IdResp{Id: id}
		response.OkWithData(idResp, c)
	}
}

// DeleteActShop 删除ActShop
// @Tags ActShop
// @Summary 删除ActShop
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.ActShop true "删除ActShop"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /actShop/deleteActShop [delete]
func (m *ActShopApi) DeleteActShop(c *gin.Context) {
	var actShop business.ActShop
	_ = c.ShouldBindJSON(&actShop)
	if err := bizSev.GetActShopSev().Delete(actShop); err != nil {
        global.LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteActShopByIds 批量删除ActShop
// @Tags ActShop
// @Summary 批量删除ActShop
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ActShop"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /actShop/deleteActShopByIds [delete]
func (m *ActShopApi) DeleteActShopByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := bizSev.GetActShopSev().DeleteByIds(IDS); err != nil {
        global.LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateActShop 更新ActShop
// @Tags ActShop
// @Summary 更新ActShop
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.ActShop true "更新ActShop"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /actShop/updateActShop [put]
func (m *ActShopApi) UpdateActShop(c *gin.Context) {
	var dataObj business.ActShop
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		response.FailWithMessage("更新失败,"+err.FirstString(), c)
		return
	}

	if err := bizSev.GetActShopSev().Update(dataObj); err != nil {
        global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindActShop 用id查询ActShop
// @Tags ActShop
// @Summary 用id查询ActShop
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.ActShop true "用id查询ActShop"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /actShop/findActShop [get]
func (m *ActShopApi) FindActShop(c *gin.Context) {
	var actShop business.ActShop
	_ = c.ShouldBindQuery(&actShop) 
	 reactShop,err:= bizSev.GetActShopSev().Get(actShop.ID,""); 
	 if errors.Is(err, gorm.ErrRecordNotFound) { 
		response.OkWithData(gin.H{"actShop": nil}, c)
	} else if err != nil { 
        global.LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else { 
		response.OkWithData(gin.H{"actShop": reactShop}, c)
	}
}

// GetActShopList 分页获取ActShop列表
// @Tags ActShop
// @Summary 分页获取ActShop列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.ActShopSearch true "分页获取ActShop列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /actShop/getActShopList [get]
func (m *ActShopApi) GetActShopList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")

	var pageInfo bizReq.ActShopSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if  list, total, err := bizSev.GetActShopSev().GetList(pageInfo,createdAtBetween,""); err != nil {
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
// @Param data body business.ActShop true "快速更新ActShop" 
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router  /actShop/quickEdit [post] 
func (m *ActShopApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "act_shop" 
	if err := commSev.GetCommonDbSev().QuickEdit(quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}


// excelList 分页导出excel ActShop列表
// @Tags ActShop
// @Summary 分页导出excel ActShop列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.ActShopSearch true "分页导出excel ActShop列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /actShop/excelList [get]
func (m *ActShopApi) ExcelList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")
	var pageInfo bizReq.ActShopSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list,_,err:= bizSev.GetActShopSev().GetListAll(pageInfo,createdAtBetween,""); err != nil {
	    global.LOG.Error("获取失败!", zap.Any("err", err))
        response.FailWithMessage("获取失败", c)
    } else {
        if len(list) == 0 {
			response.FailWithMessage("没有数据", c)
		} else { 
			sheetFields := []string{}  
					sheetFields = append(sheetFields, "用户id")  
					sheetFields = append(sheetFields, "标题")  
					sheetFields = append(sheetFields, "简介")  
					sheetFields = append(sheetFields, "详细内容")  
					sheetFields = append(sheetFields, "缩略图")  
					sheetFields = append(sheetFields, "媒体列表")  
					sheetFields = append(sheetFields, "地址")  
					sheetFields = append(sheetFields, "地区id")  
					sheetFields = append(sheetFields, "经度")  
					sheetFields = append(sheetFields, "纬度")  
					sheetFields = append(sheetFields, "邮件")  
					sheetFields = append(sheetFields, "手机")  
					sheetFields = append(sheetFields, "vip等级")  
					sheetFields = append(sheetFields, "vip结束时间")  
					sheetFields = append(sheetFields, "综合指数")  
					sheetFields = append(sheetFields, "总分享")  
					sheetFields = append(sheetFields, "总收藏")  
					sheetFields = append(sheetFields, "总报名人数")  
					sheetFields = append(sheetFields, "总评论")  
					sheetFields = append(sheetFields, "总点击")  
					sheetFields = append(sheetFields, "总评")  
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
				arr = append(arr, v.Name)
				arr = append(arr, v.Desc)
				arr = append(arr, v.Detail)
				arr = append(arr, v.Avater)
				arr = append(arr, v.MediaList)
				arr = append(arr, v.Address)
				if v.AreaId == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.AreaId)
				}
				if v.Lng == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.Lng)
				}
				if v.Lat == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.Lat)
				}
				arr = append(arr, v.Email)
				arr = append(arr, v.Mobile)
				if v.VipLev == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.VipLev)
				}
				arr = append(arr, v.VipTime)
				if v.TotalWhole == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.TotalWhole)
				}
				if v.TotalShare == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.TotalShare)
				}
				if v.TotalFav == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.TotalFav)
				}
				if v.TotalJoin == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.TotalJoin)
				}
				if v.TotalDiscuss == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.TotalDiscuss)
				}
				if v.TotalClick == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.TotalClick)
				}
				if v.TotalStar == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.TotalStar)
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


 