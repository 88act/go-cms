package common

import (
	"mime/multipart"
	"sync"

	"go-cms/global"
	"go-cms/model/business"
	"go-cms/mycache"
	"go-cms/utils"
	"go-cms/utils/upload"
)

type CommonFileService struct {
	BaseService
}

var once_CommonFileService sync.Once = sync.Once{}
var obj_CommonFileService *CommonFileService

// 获取单例
func GetCommonFileSev() *CommonFileService {
	once_CommonFileService.Do(func() {
		obj_CommonFileService = new(CommonFileService)
		//instanse.init()
	})
	return obj_CommonFileService
}

//@author: [ljd]
//@function: UploadFile 上传文件
//@description: 根据配置文件判断是文件上传到本地或者七牛云
//@param: header *multipart.FileHeader,basicFile business.BasicFile
//@return: err error, file model.ExaFileUploadAndDownload

func (m *CommonFileService) UploadFile(header *multipart.FileHeader, basicFile business.BasicFile) (err error, file business.BasicFile) {
	oss := upload.NewOss()
	//var dictSer = new(systemService.DictionaryDetailService)
	//moduleName, _ = dictSer.GetNameByValue(13, module)
	var module int
	module = basicFile.Module
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
	basicFile.Driver = driver
	err = global.DB.Create(&basicFile).Error
	return err, basicFile
}

func (m *CommonFileService) GetPathByGuid(guid string) (path string, err error) {
	//basicFS := new(businessSev.BasicFileService)
	// 先从 redis 获取 ，没有则读取数据库 ，缓存2小时
	if utils.IsEmptyStr(guid) {
		return "", nil
	}
	if len(guid) < 32 {
		global.LOG.Warn(" GetPathByGuid无效guid=" + guid)
		return "", nil
	}
	///  byteData := []byte("ssss112112qweqweqwe2")
	//	GetCache().Set("ssss", byteData)
	cacheKey := m.GetCacheKey("file", guid)
	cacheData, err := mycache.GetCache().Get(cacheKey)
	if err == nil {
		//global.LOG.Info("from cacheKey=" + cacheKey)
		path = cacheData.(string)
		//global.LOG.Info("from path=" + path)
		return path, err
	}
	basicFile := business.BasicFile{}
	err = global.DB.Where("guid = ?", guid).First(&basicFile).Error
	if err != nil {
		//cacheData = []byte("")
		mycache.GetCache().Set(cacheKey, "", "basicFile")
		return "", err
	}
	path = basicFile.Path
	if !utils.IsEmpty(path) {
		path = global.CONFIG.Local.BaseUrl + path
	}
	//cacheData = []byte(path)
	mycache.GetCache().Set(cacheKey, path, "basicFile")
	return path, err
}
