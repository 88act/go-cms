package common

import (
	"mime/multipart"

	"github.com/88act/go-cms/server/global"
	"github.com/88act/go-cms/server/model/business"
	"github.com/88act/go-cms/server/utils"
	"github.com/88act/go-cms/server/utils/upload"
)

type CommonFileService struct {
}

//@author: [ljd]
//@function: UploadFile 上传文件
//@description: 根据配置文件判断是文件上传到本地或者七牛云
//@param: header *multipart.FileHeader,basicFile business.BasicFile
//@return: err error, file model.ExaFileUploadAndDownload

func (e *CommonFileService) UploadFile(header *multipart.FileHeader, basicFile business.BasicFile) (err error, file business.BasicFile) {
	oss := upload.NewOss()

	//var dictSer = new(systemService.DictionaryDetailService)
	//moduleName, _ = dictSer.GetNameByValue(13, module)
	var module int
	module = *basicFile.Module
	filePath, guid, uploadErr := oss.UploadFile(header, module, 1)
	if uploadErr != nil {
		panic(err)
	}
	//是否保存到数据库表 basic_file
	//s := strings.Split(header.Filename, ".")
	basicFile.Guid = guid
	basicFile.Path = filePath
	basicFile.Name = header.Filename

	driver := 0
	switch global.CONFIG.System.OssType {
	case "local":
		driver = 0
	case "qiniu":
		driver = 1
	case "tencent-cos":
		driver = 2
	case "aliyun-oss":
		driver = 3
	default:
		driver = 4
	}
	basicFile.Driver = &driver
	err = global.DB.Create(&basicFile).Error
	return err, basicFile
}

func (e *CommonFileService) GetPathByGuid(guid string) (err error, path string) {
	//basicFS := new(businessSev.BasicFileService)
	// 先从 redis 获取 ，没有则读取数据库 ，缓存2小时
	if utils.IsEmpty(guid) {
		return nil, ""
	}

	basicFile := business.BasicFile{}
	err = global.DB.Where("guid = ?", guid).First(&basicFile).Error
	if err != nil {
		return err, ""
	}
	path = basicFile.Path
	if !utils.IsEmpty(path) {
		path = global.CONFIG.Local.BaseUrl + path
	}
	return err, path
}
