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

type ActActApi struct {
}

 

// CreateActAct 创建ActAct
// @Tags ActAct
// @Summary 创建ActAct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.ActAct true "创建ActAct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /actAct/createActAct [post]
func (m *ActActApi) CreateActAct(c *gin.Context) {
	var dataObj business.ActAct
	_ = c.ShouldBindJSON(&dataObj)
	
	if err := gvalid.CheckStruct(c,dataObj, nil); err != nil {
		response.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	}

 
	if id,err := bizSev.GetActActSev().Create(dataObj); err != nil {
        global.LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
	    idResp := &response.IdResp{Id: id}
		response.OkWithData(idResp, c)
	}
}

// DeleteActAct 删除ActAct
// @Tags ActAct
// @Summary 删除ActAct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.ActAct true "删除ActAct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /actAct/deleteActAct [delete]
func (m *ActActApi) DeleteActAct(c *gin.Context) {
	var actAct business.ActAct
	_ = c.ShouldBindJSON(&actAct)
	if err := bizSev.GetActActSev().Delete(actAct); err != nil {
        global.LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteActActByIds 批量删除ActAct
// @Tags ActAct
// @Summary 批量删除ActAct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ActAct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /actAct/deleteActActByIds [delete]
func (m *ActActApi) DeleteActActByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := bizSev.GetActActSev().DeleteByIds(IDS); err != nil {
        global.LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateActAct 更新ActAct
// @Tags ActAct
// @Summary 更新ActAct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.ActAct true "更新ActAct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /actAct/updateActAct [put]
func (m *ActActApi) UpdateActAct(c *gin.Context) {
	var dataObj business.ActAct
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		response.FailWithMessage("更新失败,"+err.FirstString(), c)
		return
	}

	if err := bizSev.GetActActSev().Update(dataObj); err != nil {
        global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindActAct 用id查询ActAct
// @Tags ActAct
// @Summary 用id查询ActAct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.ActAct true "用id查询ActAct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /actAct/findActAct [get]
func (m *ActActApi) FindActAct(c *gin.Context) {
	var actAct business.ActAct
	_ = c.ShouldBindQuery(&actAct) 
	 reactAct,err:= bizSev.GetActActSev().Get(actAct.ID,""); 
	 if errors.Is(err, gorm.ErrRecordNotFound) { 
		response.OkWithData(gin.H{"actAct": nil}, c)
	} else if err != nil { 
        global.LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else { 
		response.OkWithData(gin.H{"actAct": reactAct}, c)
	}
}

// GetActActList 分页获取ActAct列表
// @Tags ActAct
// @Summary 分页获取ActAct列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.ActActSearch true "分页获取ActAct列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /actAct/getActActList [get]
func (m *ActActApi) GetActActList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")

	var pageInfo bizReq.ActActSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if  list, total, err := bizSev.GetActActSev().GetList(pageInfo,createdAtBetween,""); err != nil {
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
// @Param data body business.ActAct true "快速更新ActAct" 
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router  /actAct/quickEdit [post] 
func (m *ActActApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "act_act" 
	if err := commSev.GetCommonDbSev().QuickEdit(quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}


// excelList 分页导出excel ActAct列表
// @Tags ActAct
// @Summary 分页导出excel ActAct列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.ActActSearch true "分页导出excel ActAct列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /actAct/excelList [get]
func (m *ActActApi) ExcelList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")
	var pageInfo bizReq.ActActSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list,_,err:= bizSev.GetActActSev().GetListAll(pageInfo,createdAtBetween,""); err != nil {
	    global.LOG.Error("获取失败!", zap.Any("err", err))
        response.FailWithMessage("获取失败", c)
    } else {
        if len(list) == 0 {
			response.FailWithMessage("没有数据", c)
		} else { 
			sheetFields := []string{}  
					sheetFields = append(sheetFields, "用户id")  
					sheetFields = append(sheetFields, "系统栏目")  
					sheetFields = append(sheetFields, "用户栏目")  
					sheetFields = append(sheetFields, "类型")  
					sheetFields = append(sheetFields, "活动类型")  
					sheetFields = append(sheetFields, "分期活动")  
					sheetFields = append(sheetFields, "分期")  
					sheetFields = append(sheetFields, "标题")  
					sheetFields = append(sheetFields, "简介")  
					sheetFields = append(sheetFields, "缩略图")  
					sheetFields = append(sheetFields, "媒体列表")  
					sheetFields = append(sheetFields, "地址")  
					sheetFields = append(sheetFields, "地区id")  
					sheetFields = append(sheetFields, "经度")  
					sheetFields = append(sheetFields, "纬度")  
					sheetFields = append(sheetFields, "开始时间")  
					sheetFields = append(sheetFields, "结束时间")  
					sheetFields = append(sheetFields, "直播开始")  
					sheetFields = append(sheetFields, "直播结束")  
					sheetFields = append(sheetFields, "主持/讲师")  
					sheetFields = append(sheetFields, "客服电话")  
					sheetFields = append(sheetFields, "合作电话")  
					sheetFields = append(sheetFields, "微信")  
					sheetFields = append(sheetFields, "QQ")  
					sheetFields = append(sheetFields, "群组id")  
					sheetFields = append(sheetFields, "结束报名")  
					sheetFields = append(sheetFields, "票价")  
					sheetFields = append(sheetFields, "vip票价")  
					sheetFields = append(sheetFields, "票价描述")  
					sheetFields = append(sheetFields, "退费类型")  
					sheetFields = append(sheetFields, "退费天")  
					sheetFields = append(sheetFields, "是否显示人数")  
					sheetFields = append(sheetFields, "活动状态")  
					sheetFields = append(sheetFields, "最大报名人数")  
					sheetFields = append(sheetFields, "当前报名人数")  
					sheetFields = append(sheetFields, "是否精选")  
					sheetFields = append(sheetFields, "是否周末")  
					sheetFields = append(sheetFields, "是否会员")  
					sheetFields = append(sheetFields, "综合指数")  
					sheetFields = append(sheetFields, "总分享")  
					sheetFields = append(sheetFields, "总收藏")  
					sheetFields = append(sheetFields, "总报名数")  
					sheetFields = append(sheetFields, "总评论")  
					sheetFields = append(sheetFields, "总点击")  
					sheetFields = append(sheetFields, "总评")  
					sheetFields = append(sheetFields, "总星评1")  
					sheetFields = append(sheetFields, "总星评2")  
					sheetFields = append(sheetFields, "总星评3")  
					sheetFields = append(sheetFields, "状态")  
					sheetFields = append(sheetFields, "审核原因") 

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
				if v.CatId == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.CatId)
				}
				if v.CatIdUser == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.CatIdUser)
				}
				if v.BeOnline == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.BeOnline)
				}
				arr = append(arr, v.ActType)
				if v.BeMulti == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.BeMulti)
				}
				if v.Phase == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.Phase)
				}
				arr = append(arr, v.Title)
				arr = append(arr, v.Desc)
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
				arr = append(arr, v.BeginTime)
				arr = append(arr, v.EndTime)
				arr = append(arr, v.LiveTime)
				arr = append(arr, v.LiveEnd)
				arr = append(arr, v.Presenter)
				arr = append(arr, v.PhoneKefu)
				arr = append(arr, v.PhoneHezu)
				arr = append(arr, v.Wx)
				arr = append(arr, v.Qq)
				arr = append(arr, v.GroupId)
				arr = append(arr, v.EndJoinTime)
				if v.Price == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.Price)
				}
				if v.PriceVip == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.PriceVip)
				}
				arr = append(arr, v.PriceDesc)
				if v.RefundType == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.RefundType)
				}
				if v.RefundDays == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.RefundDays)
				}
				if v.BeShowjoin == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.BeShowjoin)
				}
				if v.StatusAct == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.StatusAct)
				}
				if v.MaxJoin == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.MaxJoin)
				}
				if v.NowJoin == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.NowJoin)
				}
				if v.BeChosen == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.BeChosen)
				}
				if v.BeWeek == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.BeWeek)
				}
				if v.BeVip == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.BeVip)
				}
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
				if v.TotalStar1 == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.TotalStar1)
				}
				if v.TotalStar2 == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.TotalStar2)
				}
				if v.TotalStar3 == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.TotalStar3)
				}
				if v.Status == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.Status)
				}
				arr = append(arr, v.StatusMsg)   
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


 