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

type BbPiPersonApi struct {
}

 

// CreateBbPiPerson 创建BbPiPerson
// @Tags BbPiPerson
// @Summary 创建BbPiPerson
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.BbPiPerson true "创建BbPiPerson"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiPerson/createBbPiPerson [post]
func (bbPiPersonApi *BbPiPersonApi) CreateBbPiPerson(c *gin.Context) {
	var dataObj business.BbPiPerson
	_ = c.ShouldBindJSON(&dataObj)
	
	if err := gvalid.CheckStruct(c,dataObj, nil); err != nil {
		response.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	}

 
	if id,err := bizSev.GetBbPiPersonSev().Create(dataObj); err != nil {
        global.LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
	    idResp := &response.IdResp{Id: id}
		response.OkWithData(idResp, c)
	}
}

// DeleteBbPiPerson 删除BbPiPerson
// @Tags BbPiPerson
// @Summary 删除BbPiPerson
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.BbPiPerson true "删除BbPiPerson"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bbPiPerson/deleteBbPiPerson [delete]
func (bbPiPersonApi *BbPiPersonApi) DeleteBbPiPerson(c *gin.Context) {
	var bbPiPerson business.BbPiPerson
	_ = c.ShouldBindJSON(&bbPiPerson)
	if err := bizSev.GetBbPiPersonSev().Delete(bbPiPerson); err != nil {
        global.LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBbPiPersonByIds 批量删除BbPiPerson
// @Tags BbPiPerson
// @Summary 批量删除BbPiPerson
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BbPiPerson"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /bbPiPerson/deleteBbPiPersonByIds [delete]
func (bbPiPersonApi *BbPiPersonApi) DeleteBbPiPersonByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := bizSev.GetBbPiPersonSev().DeleteByIds(IDS); err != nil {
        global.LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBbPiPerson 更新BbPiPerson
// @Tags BbPiPerson
// @Summary 更新BbPiPerson
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.BbPiPerson true "更新BbPiPerson"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /bbPiPerson/updateBbPiPerson [put]
func (bbPiPersonApi *BbPiPersonApi) UpdateBbPiPerson(c *gin.Context) {
	var dataObj business.BbPiPerson
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		response.FailWithMessage("更新失败,"+err.FirstString(), c)
		return
	}

	if err := bizSev.GetBbPiPersonSev().Update(dataObj); err != nil {
        global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBbPiPerson 用id查询BbPiPerson
// @Tags BbPiPerson
// @Summary 用id查询BbPiPerson
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.BbPiPerson true "用id查询BbPiPerson"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /bbPiPerson/findBbPiPerson [get]
func (bbPiPersonApi *BbPiPersonApi) FindBbPiPerson(c *gin.Context) {
	var bbPiPerson business.BbPiPerson
	_ = c.ShouldBindQuery(&bbPiPerson) 
	 rebbPiPerson,err:= bizSev.GetBbPiPersonSev().Get(bbPiPerson.ID,""); 
	 if errors.Is(err, gorm.ErrRecordNotFound) { 
		response.OkWithData(gin.H{"bbPiPerson": nil}, c)
	} else if err != nil { 
        global.LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else { 
		response.OkWithData(gin.H{"bbPiPerson": rebbPiPerson}, c)
	}
}

// GetBbPiPersonList 分页获取BbPiPerson列表
// @Tags BbPiPerson
// @Summary 分页获取BbPiPerson列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.BbPiPersonSearch true "分页获取BbPiPerson列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiPerson/getBbPiPersonList [get]
func (bbPiPersonApi *BbPiPersonApi) GetBbPiPersonList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")

	var pageInfo bizReq.BbPiPersonSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if  list, total, err := bizSev.GetBbPiPersonSev().GetList(pageInfo,createdAtBetween,""); err != nil {
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
// @Param data body business.BbPiPerson true "快速更新BbPiPerson" 
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router  /bbPiPerson/quickEdit [post] 
func (bbPiPersonApi *BbPiPersonApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "bb_pi_person" 
	if err := commSev.GetCommonDbSev().QuickEdit(quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}


// excelList 分页导出excel BbPiPerson列表
// @Tags BbPiPerson
// @Summary 分页导出excel BbPiPerson列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.BbPiPersonSearch true "分页导出excel BbPiPerson列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiPerson/excelList [get]
func (bbPiPersonApi *BbPiPersonApi) ExcelList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")
	var pageInfo bizReq.BbPiPersonSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list,_,err:= bizSev.GetBbPiPersonSev().GetListAll(pageInfo,createdAtBetween,""); err != nil {
	    global.LOG.Error("获取失败!", zap.Any("err", err))
        response.FailWithMessage("获取失败", c)
    } else {
        if len(list) == 0 {
			response.FailWithMessage("没有数据", c)
		} else { 
			sheetFields := []string{}  
					sheetFields = append(sheetFields, "userid")  
					sheetFields = append(sheetFields, "上传")  
					sheetFields = append(sheetFields, "机构标识")  
					sheetFields = append(sheetFields, "卡号")  
					sheetFields = append(sheetFields, "卡类型")  
					sheetFields = append(sheetFields, "身份证件号码")  
					sheetFields = append(sheetFields, "身份证件类别")  
					sheetFields = append(sheetFields, "姓名")  
					sheetFields = append(sheetFields, "性别代码")  
					sheetFields = append(sheetFields, "性别名称")  
					sheetFields = append(sheetFields, "患者属性")  
					sheetFields = append(sheetFields, "婚姻状况代码")  
					sheetFields = append(sheetFields, "婚姻状态名称")  
					sheetFields = append(sheetFields, "出生日期")  
					sheetFields = append(sheetFields, "民族代码")  
					sheetFields = append(sheetFields, "民族名称")  
					sheetFields = append(sheetFields, "国籍代码")  
					sheetFields = append(sheetFields, "国籍名称")  
					sheetFields = append(sheetFields, "机构内部档案号")  
					sheetFields = append(sheetFields, "固定电话")  
					sheetFields = append(sheetFields, "手机号码")  
					sheetFields = append(sheetFields, "电子邮件")  
					sheetFields = append(sheetFields, "文化程度代码")  
					sheetFields = append(sheetFields, "文化程度名称")  
					sheetFields = append(sheetFields, "职业类别代码")  
					sheetFields = append(sheetFields, "职业类别名称")  
					sheetFields = append(sheetFields, "出生地行政区划码")  
					sheetFields = append(sheetFields, "出生地")  
					sheetFields = append(sheetFields, "居住地行政区划码")  
					sheetFields = append(sheetFields, "居住地址")  
					sheetFields = append(sheetFields, "户口地行政区划码")  
					sheetFields = append(sheetFields, "户口地址")  
					sheetFields = append(sheetFields, "联系人姓名")  
					sheetFields = append(sheetFields, "联系人电话")  
					sheetFields = append(sheetFields, "ABO 血型")  
					sheetFields = append(sheetFields, "RH 血型")  
					sheetFields = append(sheetFields, "参保类别代码")  
					sheetFields = append(sheetFields, "个人档案ID")  
					sheetFields = append(sheetFields, "预留一")  
					sheetFields = append(sheetFields, "预留二")  
					sheetFields = append(sheetFields, "数据生成时间")  
					sheetFields = append(sheetFields, "填报日期")  
					sheetFields = append(sheetFields, "密级")  
					sheetFields = append(sheetFields, "撤销标志") 

			excel := excelize.NewFile()
			excel.SetSheetRow("Sheet1", "A1", &sheetFields)
			for i, v := range list {
				axis := fmt.Sprintf("A%d", i+2)
				var arr = []interface{}{}
				if v.Userid == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.Userid)
				}
				if v.Status == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.Status)
				}
				arr = append(arr, v.Jgdm)
				arr = append(arr, v.Kh)
				arr = append(arr, v.Klx)
				arr = append(arr, v.Zjhm)
				arr = append(arr, v.Zjlbdm)
				arr = append(arr, v.Xm)
				arr = append(arr, v.Xbdm)
				arr = append(arr, v.Xbmc)
				arr = append(arr, v.Hzsx)
				arr = append(arr, v.Hyztdm)
				arr = append(arr, v.Hyztmc)
				arr = append(arr, v.Csrq)
				arr = append(arr, v.Mzdm)
				arr = append(arr, v.Mzmc)
				arr = append(arr, v.Gjdm)
				arr = append(arr, v.Gjmc)
				arr = append(arr, v.Jgnbdah)
				arr = append(arr, v.Gddhhm)
				arr = append(arr, v.Sjhm)
				arr = append(arr, v.Dzyj)
				arr = append(arr, v.Whcddm)
				arr = append(arr, v.Whcdmc)
				arr = append(arr, v.Zylbdm)
				arr = append(arr, v.Zylbmc)
				arr = append(arr, v.Csdxzqhm)
				arr = append(arr, v.Csd)
				arr = append(arr, v.Jzdxzqhm)
				arr = append(arr, v.Jzdz)
				arr = append(arr, v.Hkdxzqhm)
				arr = append(arr, v.Hkdz)
				arr = append(arr, v.Lxrxm)
				arr = append(arr, v.Lxrdh)
				arr = append(arr, v.Abo)
				arr = append(arr, v.Rh)
				arr = append(arr, v.Cblbdm)
				arr = append(arr, v.Grdaid)
				arr = append(arr, v.Yl1)
				arr = append(arr, v.Yl2)
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


 