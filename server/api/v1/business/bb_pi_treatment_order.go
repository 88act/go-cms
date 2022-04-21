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

type BbPiTreatmentOrderApi struct {
}

 

// CreateBbPiTreatmentOrder 创建BbPiTreatmentOrder
// @Tags BbPiTreatmentOrder
// @Summary 创建BbPiTreatmentOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.BbPiTreatmentOrder true "创建BbPiTreatmentOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiTreatmentOrder/createBbPiTreatmentOrder [post]
func (bbPiTreatmentOrderApi *BbPiTreatmentOrderApi) CreateBbPiTreatmentOrder(c *gin.Context) {
	var dataObj business.BbPiTreatmentOrder
	_ = c.ShouldBindJSON(&dataObj)
	
	if err := gvalid.CheckStruct(c,dataObj, nil); err != nil {
		response.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	}

 
	if id,err := bizSev.GetBbPiTreatmentOrderSev().Create(dataObj); err != nil {
        global.LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
	    idResp := &response.IdResp{Id: id}
		response.OkWithData(idResp, c)
	}
}

// DeleteBbPiTreatmentOrder 删除BbPiTreatmentOrder
// @Tags BbPiTreatmentOrder
// @Summary 删除BbPiTreatmentOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.BbPiTreatmentOrder true "删除BbPiTreatmentOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bbPiTreatmentOrder/deleteBbPiTreatmentOrder [delete]
func (bbPiTreatmentOrderApi *BbPiTreatmentOrderApi) DeleteBbPiTreatmentOrder(c *gin.Context) {
	var bbPiTreatmentOrder business.BbPiTreatmentOrder
	_ = c.ShouldBindJSON(&bbPiTreatmentOrder)
	if err := bizSev.GetBbPiTreatmentOrderSev().Delete(bbPiTreatmentOrder); err != nil {
        global.LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBbPiTreatmentOrderByIds 批量删除BbPiTreatmentOrder
// @Tags BbPiTreatmentOrder
// @Summary 批量删除BbPiTreatmentOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BbPiTreatmentOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /bbPiTreatmentOrder/deleteBbPiTreatmentOrderByIds [delete]
func (bbPiTreatmentOrderApi *BbPiTreatmentOrderApi) DeleteBbPiTreatmentOrderByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := bizSev.GetBbPiTreatmentOrderSev().DeleteByIds(IDS); err != nil {
        global.LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBbPiTreatmentOrder 更新BbPiTreatmentOrder
// @Tags BbPiTreatmentOrder
// @Summary 更新BbPiTreatmentOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.BbPiTreatmentOrder true "更新BbPiTreatmentOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /bbPiTreatmentOrder/updateBbPiTreatmentOrder [put]
func (bbPiTreatmentOrderApi *BbPiTreatmentOrderApi) UpdateBbPiTreatmentOrder(c *gin.Context) {
	var dataObj business.BbPiTreatmentOrder
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		response.FailWithMessage("更新失败,"+err.FirstString(), c)
		return
	}

	if err := bizSev.GetBbPiTreatmentOrderSev().Update(dataObj); err != nil {
        global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBbPiTreatmentOrder 用id查询BbPiTreatmentOrder
// @Tags BbPiTreatmentOrder
// @Summary 用id查询BbPiTreatmentOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.BbPiTreatmentOrder true "用id查询BbPiTreatmentOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /bbPiTreatmentOrder/findBbPiTreatmentOrder [get]
func (bbPiTreatmentOrderApi *BbPiTreatmentOrderApi) FindBbPiTreatmentOrder(c *gin.Context) {
	var bbPiTreatmentOrder business.BbPiTreatmentOrder
	_ = c.ShouldBindQuery(&bbPiTreatmentOrder) 
	 rebbPiTreatmentOrder,err:= bizSev.GetBbPiTreatmentOrderSev().Get(bbPiTreatmentOrder.ID,""); 
	 if errors.Is(err, gorm.ErrRecordNotFound) { 
		response.OkWithData(gin.H{"bbPiTreatmentOrder": nil}, c)
	} else if err != nil { 
        global.LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else { 
		response.OkWithData(gin.H{"bbPiTreatmentOrder": rebbPiTreatmentOrder}, c)
	}
}

// GetBbPiTreatmentOrderList 分页获取BbPiTreatmentOrder列表
// @Tags BbPiTreatmentOrder
// @Summary 分页获取BbPiTreatmentOrder列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.BbPiTreatmentOrderSearch true "分页获取BbPiTreatmentOrder列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiTreatmentOrder/getBbPiTreatmentOrderList [get]
func (bbPiTreatmentOrderApi *BbPiTreatmentOrderApi) GetBbPiTreatmentOrderList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")

	var pageInfo bizReq.BbPiTreatmentOrderSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if  list, total, err := bizSev.GetBbPiTreatmentOrderSev().GetList(pageInfo,createdAtBetween,""); err != nil {
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
// @Param data body business.BbPiTreatmentOrder true "快速更新BbPiTreatmentOrder" 
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router  /bbPiTreatmentOrder/quickEdit [post] 
func (bbPiTreatmentOrderApi *BbPiTreatmentOrderApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "bb_pi_treatment_order" 
	if err := commSev.GetCommonDbSev().QuickEdit(quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}


// excelList 分页导出excel BbPiTreatmentOrder列表
// @Tags BbPiTreatmentOrder
// @Summary 分页导出excel BbPiTreatmentOrder列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.BbPiTreatmentOrderSearch true "分页导出excel BbPiTreatmentOrder列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiTreatmentOrder/excelList [get]
func (bbPiTreatmentOrderApi *BbPiTreatmentOrderApi) ExcelList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")
	var pageInfo bizReq.BbPiTreatmentOrderSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list,_,err:= bizSev.GetBbPiTreatmentOrderSev().GetListAll(pageInfo,createdAtBetween,""); err != nil {
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
					sheetFields = append(sheetFields, "处方编号")  
					sheetFields = append(sheetFields, "处方明细ID")  
					sheetFields = append(sheetFields, "卡号")  
					sheetFields = append(sheetFields, "卡类型")  
					sheetFields = append(sheetFields, "处方开立日期")  
					sheetFields = append(sheetFields, "处方有效天数")  
					sheetFields = append(sheetFields, "处方科室编码")  
					sheetFields = append(sheetFields, "处方科室名称")  
					sheetFields = append(sheetFields, "门诊号")  
					sheetFields = append(sheetFields, "姓名")  
					sheetFields = append(sheetFields, "性别代码")  
					sheetFields = append(sheetFields, "出生日期")  
					sheetFields = append(sheetFields, "年龄岁")  
					sheetFields = append(sheetFields, "年龄月")  
					sheetFields = append(sheetFields, "就诊日期时间")  
					sheetFields = append(sheetFields, "医嘱说明")  
					sheetFields = append(sheetFields, "排序号")  
					sheetFields = append(sheetFields, "医嘱项目类型代码")  
					sheetFields = append(sheetFields, "药品处方属性")  
					sheetFields = append(sheetFields, "中药类别代码")  
					sheetFields = append(sheetFields, "处方明细医保编码")  
					sheetFields = append(sheetFields, "药监药品ID")  
					sheetFields = append(sheetFields, "处方明细名称")  
					sheetFields = append(sheetFields, "药物名称")  
					sheetFields = append(sheetFields, "药品规格")  
					sheetFields = append(sheetFields, "DDD值")  
					sheetFields = append(sheetFields, "药物剂型代码")  
					sheetFields = append(sheetFields, "药物使用次剂量")  
					sheetFields = append(sheetFields, "药物使用剂量单位")  
					sheetFields = append(sheetFields, "药物使用频次代码")  
					sheetFields = append(sheetFields, "药物使用频次名称")  
					sheetFields = append(sheetFields, "用药途径代码")  
					sheetFields = append(sheetFields, "用药途径名称")  
					sheetFields = append(sheetFields, "药物使用总剂量")  
					sheetFields = append(sheetFields, "处方药品组号")  
					sheetFields = append(sheetFields, "中药饮片处方")  
					sheetFields = append(sheetFields, "中药饮片剂数")  
					sheetFields = append(sheetFields, "中药饮片煎煮法")  
					sheetFields = append(sheetFields, "中药用药方法")  
					sheetFields = append(sheetFields, "发药剂量")  
					sheetFields = append(sheetFields, "发药剂量单位")  
					sheetFields = append(sheetFields, "单价")  
					sheetFields = append(sheetFields, "总金额")  
					sheetFields = append(sheetFields, "处方开立医师工号")  
					sheetFields = append(sheetFields, "处方开立医师签名")  
					sheetFields = append(sheetFields, "处方审核工号")  
					sheetFields = append(sheetFields, "处方审核签名")  
					sheetFields = append(sheetFields, "处方调配工号")  
					sheetFields = append(sheetFields, "处方调配签名")  
					sheetFields = append(sheetFields, "处方核对工号")  
					sheetFields = append(sheetFields, "处方核对签名")  
					sheetFields = append(sheetFields, "处方发药工号")  
					sheetFields = append(sheetFields, "处方发药签名")  
					sheetFields = append(sheetFields, "执行结果")  
					sheetFields = append(sheetFields, "备注")  
					sheetFields = append(sheetFields, "取药机构代码")  
					sheetFields = append(sheetFields, "取药机构名称")  
					sheetFields = append(sheetFields, "数据生成时间")  
					sheetFields = append(sheetFields, "填报日期")  
					sheetFields = append(sheetFields, "撤销标志")  
					sheetFields = append(sheetFields, "密级") 

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
				arr = append(arr, v.Cfbh)
				arr = append(arr, v.Cfmxid)
				arr = append(arr, v.Kh)
				arr = append(arr, v.Klx)
				arr = append(arr, v.Cfklsj)
				arr = append(arr, v.Cfyxts)
				arr = append(arr, v.Cfklksbm)
				arr = append(arr, v.Cfklksmc)
				arr = append(arr, v.Mzh)
				arr = append(arr, v.Xm)
				arr = append(arr, v.Xbdm)
				arr = append(arr, v.Csrq)
				arr = append(arr, v.Nls)
				arr = append(arr, v.Nly)
				arr = append(arr, v.Jzrqsj)
				arr = append(arr, v.Yzsm)
				arr = append(arr, v.Pxh)
				arr = append(arr, v.Yzxmlxdm)
				arr = append(arr, v.Ypcfsx)
				arr = append(arr, v.Zylbdm)
				arr = append(arr, v.Cfmxxmbm)
				arr = append(arr, v.Ypid)
				arr = append(arr, v.Cfmxmc)
				arr = append(arr, v.Ywmc)
				arr = append(arr, v.Ypgg)
				arr = append(arr, v.Dddz)
				arr = append(arr, v.Ywjxdm)
				arr = append(arr, v.Ywsycjl)
				arr = append(arr, v.Ywsyjldw)
				arr = append(arr, v.Ywsypcdm)
				arr = append(arr, v.Ywsypcmc)
				arr = append(arr, v.Yytjdm)
				arr = append(arr, v.Yytjmc)
				arr = append(arr, v.Ywsyzjl)
				arr = append(arr, v.Cfypzh)
				arr = append(arr, v.Zyypcf)
				arr = append(arr, v.Zyypjs)
				arr = append(arr, v.Zyypjzf)
				arr = append(arr, v.Zyyyff)
				arr = append(arr, v.Fyjl)
				arr = append(arr, v.Fyjldw)
				arr = append(arr, v.Dj)
				arr = append(arr, v.Zje)
				arr = append(arr, v.Cfklysgh)
				arr = append(arr, v.Cfklysqm)
				arr = append(arr, v.Cfshyjsgh)
				arr = append(arr, v.Cfshyjsqm)
				arr = append(arr, v.Cftpyjsgh)
				arr = append(arr, v.Cftpyjsqm)
				arr = append(arr, v.Cfhdyjsgh)
				arr = append(arr, v.Cfhdyjsqm)
				arr = append(arr, v.Cffyyjsgh)
				arr = append(arr, v.Cffyyjsqm)
				arr = append(arr, v.Zxjg)
				arr = append(arr, v.Bz)
				arr = append(arr, v.Qyjgdm)
				arr = append(arr, v.Qyjgmc)
				arr = append(arr, v.Sjscsj)
				arr = append(arr, v.Tbrqsj)
				arr = append(arr, v.Cxbz)
				arr = append(arr, v.Mj)   
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


 