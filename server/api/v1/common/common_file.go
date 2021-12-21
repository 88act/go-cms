package common

import (
	"strconv"

	"go-cms/global"
	"go-cms/model/business"
	"go-cms/model/common/response"
	"go-cms/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CommonFileApi struct {
}

var commonFileService = service.ServiceGroupApp.CommonServiceGroup.CommonFileService

// @Tags UploadFile
// @Summary 上传文件
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "上传文件 "
// @Success 200 {string} string "{"success":true,"data":{},"msg":"上传成功"}"
// @Router /commFile/upload [post]
func (commonFileApi *CommonFileApi) UploadFile(c *gin.Context) {
	module := c.DefaultPostForm("module", "0")
	media_type := c.DefaultPostForm("media_type", "0") // 此方法可以设置默认值
	size := c.DefaultPostForm("size", "0")
	ext := c.DefaultPostForm("ext", "")
	md5 := c.DefaultPostForm("md5", "")
	sha1 := c.DefaultPostForm("sha1", "")
	media_type_v, err := strconv.Atoi(media_type)
	module_v, err := strconv.Atoi(module)
	size_v, err := strconv.Atoi(size)

	basicFileMini := business.BasicFileMini{
		Size:      &size_v,
		Ext:       ext,
		MediaType: &media_type_v,
		Module:    &module_v,
	}
	basicFile := business.BasicFile{
		BasicFileMini: basicFileMini,
		Md5:           md5,
		Sha1:          sha1,
	}
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		global.LOG.Error("接收文件失败!", zap.Any("err", err))
		response.FailWithMessage("接收文件失败", c)
		return
	}
	//var basicFile business.BasicFile
	//basicFileService.CreateBasicFile(basicFile)
	//err, file = commFileService.UploadFile(header, noSave)
	var newBasicFile business.BasicFile
	err, newBasicFile = commonFileService.UploadFile(header, basicFile) // 文件上传后拿到文件路径
	//err, file = basicFileService.CreateBasicFile(header, noSave) // 文件上传后拿到文件路径
	if err != nil {
		global.LOG.Error("修改数据库链接失败!", zap.Any("err", err))
		response.FailWithMessage("修改数据库链接失败", c)
		return
	}
	newBasicFile.Path = global.CONFIG.Local.BaseUrl + newBasicFile.Path
	response.OkWithDetailed(newBasicFile, "上传成功", c)
	//response.OkWithDetailed({File: file}, "上传成功", c)

}
