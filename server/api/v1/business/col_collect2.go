package business

import (
	"errors"
	"fmt"

	"go-cms/global"
	"go-cms/model/common/response"
	"go-cms/myError"
	"go-cms/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ColCollectApi2 struct {
	ColCollectApi
}

// StartOrStopCollect 启动/停止收集
// @Tags ColCollect
// @Summary 用id查询ColCollect
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.ColCollect true "用id查询ColCollect"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /colCollect/StartOrStopCollect [get]
func (m *ColCollectApi2) StartOrStopCollect(c *gin.Context) {
	//var colCollect business.ColCollect
	//_ = c.ShouldBindQuery(&colCollect)
	id := utils.StrToUInt(c.DefaultQuery("ID", "0"))
	opt := utils.StrToInt(c.DefaultQuery("opt", "0"))
	fmt.Println(id, opt)
	err, collent := colCollectService.GetColCollect(id, "")
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.FailWithMessage("不存在的采集器", c)
		//response.OkWithData(gin.H{"colCollect": nil}, c)
	} else if err != nil {
		global.LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		//启动收集器
		myErr := collectManager.Start(collent, opt)
		value, ok := myErr.(myError.MyError)
		if ok {
			if value.Type == myError.ErrOK {
				response.OkWithMessage(value.Msg, c)
			} else {
				response.FailWithMessage(value.Msg, c)
			}
		} else {
			response.FailWithMessage(myErr.Error(), c)
		}

	}
}
