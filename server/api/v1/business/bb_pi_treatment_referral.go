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

type BbPiTreatmentReferralApi struct {
}

 

// CreateBbPiTreatmentReferral 创建BbPiTreatmentReferral
// @Tags BbPiTreatmentReferral
// @Summary 创建BbPiTreatmentReferral
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.BbPiTreatmentReferral true "创建BbPiTreatmentReferral"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiTreatmentReferral/createBbPiTreatmentReferral [post]
func (bbPiTreatmentReferralApi *BbPiTreatmentReferralApi) CreateBbPiTreatmentReferral(c *gin.Context) {
	var dataObj business.BbPiTreatmentReferral
	_ = c.ShouldBindJSON(&dataObj)
	
	if err := gvalid.CheckStruct(c,dataObj, nil); err != nil {
		response.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	}

 
	if id,err := bizSev.GetBbPiTreatmentReferralSev().Create(dataObj); err != nil {
        global.LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
	    idResp := &response.IdResp{Id: id}
		response.OkWithData(idResp, c)
	}
}

// DeleteBbPiTreatmentReferral 删除BbPiTreatmentReferral
// @Tags BbPiTreatmentReferral
// @Summary 删除BbPiTreatmentReferral
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.BbPiTreatmentReferral true "删除BbPiTreatmentReferral"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bbPiTreatmentReferral/deleteBbPiTreatmentReferral [delete]
func (bbPiTreatmentReferralApi *BbPiTreatmentReferralApi) DeleteBbPiTreatmentReferral(c *gin.Context) {
	var bbPiTreatmentReferral business.BbPiTreatmentReferral
	_ = c.ShouldBindJSON(&bbPiTreatmentReferral)
	if err := bizSev.GetBbPiTreatmentReferralSev().Delete(bbPiTreatmentReferral); err != nil {
        global.LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBbPiTreatmentReferralByIds 批量删除BbPiTreatmentReferral
// @Tags BbPiTreatmentReferral
// @Summary 批量删除BbPiTreatmentReferral
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BbPiTreatmentReferral"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /bbPiTreatmentReferral/deleteBbPiTreatmentReferralByIds [delete]
func (bbPiTreatmentReferralApi *BbPiTreatmentReferralApi) DeleteBbPiTreatmentReferralByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := bizSev.GetBbPiTreatmentReferralSev().DeleteByIds(IDS); err != nil {
        global.LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBbPiTreatmentReferral 更新BbPiTreatmentReferral
// @Tags BbPiTreatmentReferral
// @Summary 更新BbPiTreatmentReferral
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.BbPiTreatmentReferral true "更新BbPiTreatmentReferral"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /bbPiTreatmentReferral/updateBbPiTreatmentReferral [put]
func (bbPiTreatmentReferralApi *BbPiTreatmentReferralApi) UpdateBbPiTreatmentReferral(c *gin.Context) {
	var dataObj business.BbPiTreatmentReferral
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		response.FailWithMessage("更新失败,"+err.FirstString(), c)
		return
	}

	if err := bizSev.GetBbPiTreatmentReferralSev().Update(dataObj); err != nil {
        global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBbPiTreatmentReferral 用id查询BbPiTreatmentReferral
// @Tags BbPiTreatmentReferral
// @Summary 用id查询BbPiTreatmentReferral
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.BbPiTreatmentReferral true "用id查询BbPiTreatmentReferral"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /bbPiTreatmentReferral/findBbPiTreatmentReferral [get]
func (bbPiTreatmentReferralApi *BbPiTreatmentReferralApi) FindBbPiTreatmentReferral(c *gin.Context) {
	var bbPiTreatmentReferral business.BbPiTreatmentReferral
	_ = c.ShouldBindQuery(&bbPiTreatmentReferral) 
	 rebbPiTreatmentReferral,err:= bizSev.GetBbPiTreatmentReferralSev().Get(bbPiTreatmentReferral.ID,""); 
	 if errors.Is(err, gorm.ErrRecordNotFound) { 
		response.OkWithData(gin.H{"bbPiTreatmentReferral": nil}, c)
	} else if err != nil { 
        global.LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else { 
		response.OkWithData(gin.H{"bbPiTreatmentReferral": rebbPiTreatmentReferral}, c)
	}
}

// GetBbPiTreatmentReferralList 分页获取BbPiTreatmentReferral列表
// @Tags BbPiTreatmentReferral
// @Summary 分页获取BbPiTreatmentReferral列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.BbPiTreatmentReferralSearch true "分页获取BbPiTreatmentReferral列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiTreatmentReferral/getBbPiTreatmentReferralList [get]
func (bbPiTreatmentReferralApi *BbPiTreatmentReferralApi) GetBbPiTreatmentReferralList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")

	var pageInfo bizReq.BbPiTreatmentReferralSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if  list, total, err := bizSev.GetBbPiTreatmentReferralSev().GetList(pageInfo,createdAtBetween,""); err != nil {
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
// @Param data body business.BbPiTreatmentReferral true "快速更新BbPiTreatmentReferral" 
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router  /bbPiTreatmentReferral/quickEdit [post] 
func (bbPiTreatmentReferralApi *BbPiTreatmentReferralApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "bb_pi_treatment_referral" 
	if err := commSev.GetCommonDbSev().QuickEdit(quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}


// excelList 分页导出excel BbPiTreatmentReferral列表
// @Tags BbPiTreatmentReferral
// @Summary 分页导出excel BbPiTreatmentReferral列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.BbPiTreatmentReferralSearch true "分页导出excel BbPiTreatmentReferral列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiTreatmentReferral/excelList [get]
func (bbPiTreatmentReferralApi *BbPiTreatmentReferralApi) ExcelList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")
	var pageInfo bizReq.BbPiTreatmentReferralSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list,_,err:= bizSev.GetBbPiTreatmentReferralSev().GetListAll(pageInfo,createdAtBetween,""); err != nil {
	    global.LOG.Error("获取失败!", zap.Any("err", err))
        response.FailWithMessage("获取失败", c)
    } else {
        if len(list) == 0 {
			response.FailWithMessage("没有数据", c)
		} else { 
			sheetFields := []string{}  
					sheetFields = append(sheetFields, "上传")  
					sheetFields = append(sheetFields, "机构标识")  
					sheetFields = append(sheetFields, "服务网点代码")  
					sheetFields = append(sheetFields, "转诊记录")  
					sheetFields = append(sheetFields, "卡号")  
					sheetFields = append(sheetFields, "卡类型")  
					sheetFields = append(sheetFields, "门诊号")  
					sheetFields = append(sheetFields, "姓名")  
					sheetFields = append(sheetFields, "性别")  
					sheetFields = append(sheetFields, "出生日期")  
					sheetFields = append(sheetFields, "年龄岁")  
					sheetFields = append(sheetFields, "年龄月")  
					sheetFields = append(sheetFields, "就诊时间")  
					sheetFields = append(sheetFields, "负责医生工号")  
					sheetFields = append(sheetFields, "负责医生姓名")  
					sheetFields = append(sheetFields, "负责科室编码")  
					sheetFields = append(sheetFields, "负责科室名称")  
					sheetFields = append(sheetFields, "转诊原因")  
					sheetFields = append(sheetFields, "转诊时间")  
					sheetFields = append(sheetFields, "转向机构代码")  
					sheetFields = append(sheetFields, "转向机构名称")  
					sheetFields = append(sheetFields, "转向科室代码")  
					sheetFields = append(sheetFields, "转向科室名称")  
					sheetFields = append(sheetFields, "转向医生工号")  
					sheetFields = append(sheetFields, "转向医生姓名")  
					sheetFields = append(sheetFields, "主要既往史")  
					sheetFields = append(sheetFields, "治疗经过")  
					sheetFields = append(sheetFields, "下一步治疗方案")  
					sheetFields = append(sheetFields, "转诊标志")  
					sheetFields = append(sheetFields, "数据生成时间")  
					sheetFields = append(sheetFields, "填报日期")  
					sheetFields = append(sheetFields, "密级")  
					sheetFields = append(sheetFields, "撤销标志") 

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
				arr = append(arr, v.Fwwddm)
				arr = append(arr, v.Zzjlid)
				arr = append(arr, v.Kh)
				arr = append(arr, v.Klx)
				arr = append(arr, v.Mzh)
				arr = append(arr, v.Xm)
				arr = append(arr, v.Xb)
				arr = append(arr, v.Csrq)
				arr = append(arr, v.Nls)
				arr = append(arr, v.Nly)
				arr = append(arr, v.Jzrqsj)
				arr = append(arr, v.Fzysgh)
				arr = append(arr, v.Fzysxm)
				arr = append(arr, v.Fzksbm)
				arr = append(arr, v.Fzksmc)
				arr = append(arr, v.Zzyy)
				arr = append(arr, v.Zzsj)
				arr = append(arr, v.Zxjgdm)
				arr = append(arr, v.Zxjgmc)
				arr = append(arr, v.Zxksdm)
				arr = append(arr, v.Zxksmc)
				arr = append(arr, v.Zxysgh)
				arr = append(arr, v.Zxysxm)
				arr = append(arr, v.Zyjws)
				arr = append(arr, v.Zljg)
				arr = append(arr, v.Xybzlfa)
				arr = append(arr, v.Zzbz)
				arr = append(arr, v.Sjscsj)
				arr = append(arr, v.Tbrqsj)
				arr = append(arr, v.Mj)
				arr = append(arr, v.Cxbz)   
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


 