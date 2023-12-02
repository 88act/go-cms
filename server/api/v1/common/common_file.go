package common

import (
	"fmt"
	"path"

	"go-cms/global"
	"go-cms/model/business"
	"go-cms/service"
	"go-cms/utils"

	bizSev "go-cms/service/business"
	commSev "go-cms/service/common"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gconv"
	"go.uber.org/zap"
)

const DEFAULT_MAX_WIDTH float64 = 200
const DEFAULT_MAX_HEIGHT float64 = 200

var allow_ext map[string]string = map[string]string{
	".jpg":  "jpg",
	".png":  "png",
	".jpeg": "jpeg",
	".mp3":  "mp3",
	".wav":  "wav",
	".mp4":  "mp4",
	".pdf":  "pdf",
	".zip":  "zip",
	".docx": "docx",
	".doc":  "doc",
}

var img_ext map[string]string = map[string]string{
	".jpg":  "jpg",
	".png":  "png",
	".jpeg": "jpeg",
}
var video_ext map[string]string = map[string]string{
	".mp4": "mp4",
}
var audio_ext map[string]string = map[string]string{
	".mp3": "mp3",
	".wav": "wav",
}
var doc_ext map[string]string = map[string]string{
	".pdf":  "pdf",
	".zip":  "zip",
	".docx": "docx",
	".doc":  "doc",
}

type CommonFileApi struct {
	commSev.BaseApi
}

var commonFileService = service.ServiceGroupApp.CommonServiceGroup.CommonFileService

func (m *CommonFileApi) GetFileByKey(c *gin.Context) {
	md5 := c.Query("md5")
	if utils.IsEmptyStr(md5) {
		m.FailWithMessage("参数错误", c)
		c.Abort()
		return
	}

	mapData := make(map[string]interface{})
	mapData["md5"] = md5
	//mapData["user_id"] = utils.GetUserID(c)
	newBasicFile, err := bizSev.GetBasicFileSev().GetByMap(c, mapData, "id,name,path,guid,sha1,md5,media_type")
	// if errors.Is(err, gorm.ErrRecordNotFound) {
	// 	m.OkWithData(gin.H{"basicFile": nil}, c)
	// } else
	if err != nil {
		global.LOG.Error("查询失败!", zap.Any("err", err))
		m.FailWithMessage("查询失败", c)
	} else {
		fileMini := business.BasicFileMini{}
		fileMini.Path = global.CONFIG.Local.BaseUrl + newBasicFile.Path
		fileMini.Guid = newBasicFile.Guid
		fileMini.Name = newBasicFile.Name
		m.OkWithData(fileMini, c)
	}
}

// @Tags UploadFile
// @Summary 上传文件
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "上传文件 "
// @Success 200 {string} string "{"success":true,"data":{},"msg":"上传成功"}"
// @Router /commFile/upload [post]
func (m *CommonFileApi) UploadFile(c *gin.Context) {

	basicFile := business.BasicFile{}
	basicFile.UserId = utils.GetUserID(c)
	//basicFile.CuId = utils.GetCuId(c)
	//basicFile.CatId = gconv.Int64(c.DefaultPostForm("catId", "0"))
	basicFile.Size = gconv.Int(c.DefaultPostForm("size", "0"))
	basicFile.Ext = c.DefaultPostForm("ext", "")
	basicFile.Md5 = c.DefaultPostForm("md5", "")
	//basicFile.Sha1 = c.DefaultPostForm("sha1", "")
	basicFile.Module = gconv.Int(c.DefaultPostForm("module", "0"))
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		global.LOG.Error("接收文件失败!", zap.Any("err", err))
		m.FailWithMessage("接收文件失败", c)
		return
	}
	basicFile.Ext = path.Ext(header.Filename)
	_, ok := allow_ext[basicFile.Ext]
	if !ok {
		m.FailWithMessage("非法文件类型 ,Ext =  "+basicFile.Ext, c)
		//return nil, errors.Wrapf(xerr.NewErrCode(xerr.Fail), "非法文件类型 ,Ext = %v", basicFile.Ext)
		//return nil, errors.New("文件类型错误")
	}
	if _, ok = img_ext[basicFile.Ext]; ok {
		mtype := 1
		basicFile.MediaType = mtype
	}
	if _, ok = video_ext[basicFile.Ext]; ok {
		mtype := 2
		basicFile.MediaType = mtype
	}
	if _, ok = audio_ext[basicFile.Ext]; ok {
		mtype := 3
		basicFile.MediaType = mtype
	}
	if _, ok = doc_ext[basicFile.Ext]; ok {
		mtype := 4
		basicFile.MediaType = mtype
	}

	//var basicFile business.BasicFile
	//basicFileService.CreateBasicFile(basicFile)
	//err, file = commFileService.UploadFile(header, noSave)
	fmt.Println("basicFile==============", basicFile)
	var newBasicFile business.BasicFile
	err, newBasicFile = commonFileService.UploadFile(header, basicFile) // 文件上传后拿到文件路径
	//err, file = basicFileService.CreateBasicFile(header, noSave) // 文件上传后拿到文件路径
	if err != nil {
		m.FailWithMessage("上传失败"+err.Error(), c)
		return
	}

	fileMini := business.BasicFileMini{}
	fileMini.Path = global.CONFIG.Local.BaseUrl + newBasicFile.Path
	fileMini.Guid = newBasicFile.Guid
	fileMini.Name = newBasicFile.Name
	m.OkWithDetailed(fileMini, "上传成功", c)
	//m.OkWithDetailed({File: file}, "上传成功", c)

}

func (m *CommonFileApi) UploadFileToken(c *gin.Context) {
	fmt.Println("  UploadFileToken = ")
	userId := utils.GetUserID(c)
	fmt.Println("  userId = ", userId)

	basicFile := business.BasicFile{}
	basicFile.UserId = utils.GetUserID(c)
	//basicFile.CuId = basicFile.UserId
	basicFile.CatId = 1 //gconv.Int64(c.DefaultPostForm("catId", "0"))
	basicFile.Size = gconv.Int(c.DefaultPostForm("size", "0"))
	basicFile.Ext = c.DefaultPostForm("ext", "")
	basicFile.Md5 = c.DefaultPostForm("md5", "")
	basicFile.Sha1 = c.DefaultPostForm("sha1", "")
	basicFile.Module = 0 // gconv.Int(c.DefaultPostForm("module", "0"))

	fmt.Print("basicFil===========")
	fmt.Print(basicFile)
	// 注意 ： 这里不一样 file[]
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		global.LOG.Error("接收文件失败!", zap.Any("err", err))
		m.FailWithMessage("接收文件失败", c)
		return
	}

	basicFile.Ext = path.Ext(header.Filename)

	_, ok := allow_ext[basicFile.Ext]
	if !ok {
		m.FailWithMessage("非法文件类型 ,Ext =  "+basicFile.Ext, c)
		//return nil, errors.Wrapf(xerr.NewErrCode(xerr.Fail), "非法文件类型 ,Ext = %v", basicFile.Ext)
		//return nil, errors.New("文件类型错误")
	}
	if _, ok = img_ext[basicFile.Ext]; ok {
		mtype := 1
		basicFile.MediaType = mtype
	}
	if _, ok = video_ext[basicFile.Ext]; ok {
		mtype := 2
		basicFile.MediaType = mtype
	}
	if _, ok = audio_ext[basicFile.Ext]; ok {
		mtype := 3
		basicFile.MediaType = mtype
	}
	if _, ok = doc_ext[basicFile.Ext]; ok {
		mtype := 4
		basicFile.MediaType = mtype
	}

	//var basicFile business.BasicFile
	//basicFileService.CreateBasicFile(basicFile)
	//err, file = commFileService.UploadFile(header, noSave)
	var newBasicFile business.BasicFile
	err, newBasicFile = commonFileService.UploadFile(header, basicFile) // 文件上传后拿到文件路径
	//err, file = basicFileService.CreateBasicFile(header, noSave) // 文件上传后拿到文件路径
	if err != nil {
		global.LOG.Error("修改数据库链接失败!", zap.Any("err", err))
		m.FailWithMessage("修改数据库链接失败", c)
		return
	}

	fileMini := business.BasicFileMini{}
	fileMini.Path = global.CONFIG.Local.BaseUrl + newBasicFile.Path
	fileMini.Guid = newBasicFile.Guid
	fileMini.Name = newBasicFile.Name
	m.OkWithDetailed(fileMini, "上传成功", c)
	//m.OkWithDetailed({File: file}, "上传成功", c)

}
