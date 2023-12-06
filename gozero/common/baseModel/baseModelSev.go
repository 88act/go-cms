// 自动生成模板MemUser
package baseModel

import (
	"context"
	"go-cms/common/mycache"
	"go-cms/common/myconfig"
	"go-cms/common/utils"
	"strings"

	"github.com/gogf/gf/util/gconv"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

type BaseModelSev struct {
	Db              *gorm.DB
	Cache           cache.ClusterConf
	CacheKeyListPre string
	CacheKeyPre     string
}

// func (m *BaseModelSev) GetPathByGuid(ctx context.Context, guid string) (path string, err error) {

// 	//basicFS := new(businessSev.BasicFileService)
// 	// 先从 redis 获取 ，没有则读取数据库 ，缓存2小时
// 	if utils.IsEmptyStr(guid) {
// 		return "", nil
// 	}
// 	if len(guid) < 32 {
// 		logx.Errorf(" GetPathByGuid无效guid=" + guid)
// 		return "", nil
// 	}
// 	///  byteData := []byte("ssss112112qweqweqwe2")
// 	//	GetCache().Set("ssss", byteData)
// 	cacheKey := m.GetCacheKey("file", guid)
// 	cacheData, err := cache.GetCache().Get(cacheKey)
// 	if err == nil {
// 		//global.LOG.Info("from cacheKey=" + cacheKey)
// 		path = cacheData.(string)
// 		//global.LOG.Info("from path=" + path)
// 		return path, err
// 	}
// 	basicFile := BasicFile{}
// 	err =  m.Db.WithContext(ctx).Model(basicFile).Where("guid = ?", guid).First(&basicFile).Error
// 	if err != nil {
// 		//cacheData = []byte("")
// 		cache.GetCache().Set(cacheKey, "")
// 		return "", err
// 	}
// 	path = basicFile.Path
// 	if !utils.IsEmpty(path) {
// 		path = ctx.Config.LocalRes.BaseUrl + path
// 	}
// 	//cacheData = []byte(path)
// 	cache.GetCache().Set(cacheKey, path)
// 	return path, err
// }

func (m *BaseModelSev) GetPathByGuid(ctx context.Context, guid string) (path string, err error) {
	// 如果是 http开头 ，直接返回
	if strings.HasPrefix(guid, "http") {
		return guid, nil
	}

	//basicFS := new(businessSev.BasicFileService)
	// 先从 redis 获取 ，没有则读取数据库 ，缓存2小时
	if utils.IsEmptyStr(guid) {
		return "", nil
	}
	// if len(guid) < 32 {
	// 	logx.WithContext(ctx).Infow("GetPathByGuid无效guid=" + guid)
	// 	return "", nil
	// }
	///  byteData := []byte("ssss112112qweqweqwe2")
	//	GetCache().Set("ssss", byteData)
	cacheKey := m.GetCacheKey("file", guid)
	cacheData, err := mycache.GetCache().Get(cacheKey)
	if err == nil {
		//logx.Errorf("from cacheKey=" + cacheKey)
		path = cacheData.(string)
		//global.LOG.Info("from path=" + path)
		return path, err
	}
	basicFile := BasicFile{}
	// mapData := make(map[string]interface{})
	// mapData["guid"] = guid
	// mapData["status"] = 1
	err = m.Db.WithContext(ctx).Where("guid = ?  ", guid).Select("path,guid").First(&basicFile).Error
	if err != nil {
		//cacheData = []byte("")
		mycache.GetCache().Set(cacheKey, "", 0)
		return "", err
	}
	path = basicFile.Path
	if !utils.IsEmpty(path) {
		path = myconfig.HttpRoot + path
	}
	//cacheData = []byte(path)
	mycache.GetCache().Set(cacheKey, path, 0)
	return path, err
}

func (m *BaseModelSev) DelCache(table string, ids []int64) {
	for _, v := range ids {
		cacheKey := m.GetCacheKey(table, v)
		mycache.GetCache().Delete(cacheKey)
	}
}
func (m *BaseModelSev) DelCache32(table string, ids []int) {
	for _, v := range ids {
		cacheKey := m.GetCacheKey(table, v)
		mycache.GetCache().Delete(cacheKey)
	}
}
func (m *BaseModelSev) GetCacheKey(table string, id interface{}) (key string) {
	key = table + "_" + gconv.String(id)
	return key
}

//@author: [ljd]
//@function: UploadFile 上传文件
//@description: 根据配置文件判断是文件上传到本地或者七牛云
//@param: header *multipart.FileHeader,basicFile business.BasicFile
//@return: err error, file model.ExaFileUploadAndDownload

// func (e *BasicFileService) UploadFile(header *multipart.FileHeader, basicFile business.BasicFile) (err error, file business.BasicFile) {
// 	oss := upload.NewOss()

// 	//var dictSer = new(systemService.DictionaryDetailService)
// 	//moduleName, _ = dictSer.GetNameByValue(13, module)
// 	var module int
// 	module = *basicFile.Module
// 	filePath, guid, uploadErr := oss.UploadFile(header, module, 1)
// 	if uploadErr != nil {
// 		panic(err)
// 	}
// 	//是否保存到数据库表 basic_file
// 	//s := strings.Split(header.Filename, ".")
// 	basicFile.Guid = guid
// 	basicFile.Path = filePath
// 	basicFile.Name = header.Filename

// 	driver := 0
// 	switch global.CONFIG.System.OssType {
// 	case "local":
// 		driver = 0
// 	case "qiniu":
// 		driver = 1
// 	case "tencent-cos":
// 		driver = 2
// 	case "aliyun-oss":
// 		driver = 3
// 	default:
// 		driver = 4
// 	}
// 	basicFile.Driver = &driver
// 	err = global.DB.Create(&basicFile).Error
// 	return err, basicFile
// }
