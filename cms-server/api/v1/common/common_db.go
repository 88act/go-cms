package common

import (
	"github.com/88act/go-cms/server/global"
	"github.com/88act/go-cms/server/model/common/request"
	"github.com/88act/go-cms/server/model/common/response"
	"github.com/88act/go-cms/server/service"
	var_dump "github.com/favframework/debug"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CommonDbApi struct {
}

var commDbService = service.ServiceGroupApp.CommonServiceGroup.CommonDbService

// QuickEdit 更新QuickEdit
// @Tags QuickEdit
// @Summary 更新QuickEdit
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.CmsAd true "更新 QuickEdit"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cmsAd/updateCmsAd [put]
func (CommonDbApi *CommonDbApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	var_dump.Dump(quickEdit)

	if err := commDbService.QuickEdit(quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}
