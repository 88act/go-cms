package common

import (
	"context"
	"go-cms/global"
	"go-cms/model/business"
	"go-cms/mycache"
	"go-cms/utils"

	"strings"

	"github.com/gogf/gf/util/gconv"
	"gorm.io/gorm"
)

var UserType_plat = 9 //平台
var UserType_cuAdmin = 3
var UserType_cuChild = 2
var UserType_cuWorker = 1
var Status_ok = 1

// 客户管理员 默认权限
var DefaultRoleId int64 = 3

type BaseService struct {
	Db *gorm.DB
}

// 获取文件路径
func (m *BaseService) GetPathByGuid(ctx context.Context, guid string) (path string, err error) {
	// 如果是 http开头 ，直接返回
	//global.LOG.Info("GetPathByGuid  guid=" + guid)
	if strings.HasPrefix(guid, "http") {
		return guid, nil
	}

	//basicFS := new(businessSev.BasicFileService)
	// 先从 redis 获取 ，没有则读取数据库 ，缓存2小时
	if utils.IsEmptyStr(guid) {
		return "", nil
	}
	// if len(guid) < 32 {
	// 	global.LOG.Error("guid长度无效")
	// 	return "", nil
	// }
	///  byteData := []byte("ssss112112qweqweqwe2")
	//	GetCache().Set("ssss", byteData)
	cacheKey := m.GetCacheKey("file", guid)
	cacheData, err := mycache.GetCache().Get(cacheKey)
	if err == nil {
		//logx.Errorf("from cacheKey=" + cacheKey)
		path = cacheData.(string)
		//global.LOG.Info("from cacheKey  path=" + path)
		return path, err
	}
	basicFile := business.BasicFile{}
	// mapData := make(map[string]interface{})
	// mapData["guid"] = guid
	// mapData["status"] = 1
	err = global.DB.WithContext(ctx).Where("guid = ? ", guid).Select("path,guid").First(&basicFile).Error
	if err != nil {
		//cacheData = []byte("")
		mycache.GetCache().Set(cacheKey, "", "GetPathByGuid")
		return "", err
	}
	path = basicFile.Path
	if !utils.IsEmpty(path) {
		path = global.CONFIG.Local.BaseUrl + path
		//path = "http://gozero.local.com/" + path
	}
	//cacheData = []byte(path)
	mycache.GetCache().Set(cacheKey, path, "GetPathByGuid")
	return path, err
}

func (m *BaseService) DelCache(table string, ids []int64) {
	for _, v := range ids {
		cacheKey := m.GetCacheKey(table, v)
		mycache.GetCache().Cache.Delete(cacheKey)
	}
}
func (m *BaseService) DelCache32(table string, ids []int) {
	for _, v := range ids {
		cacheKey := m.GetCacheKey(table, v)
		mycache.GetCache().Cache.Delete(cacheKey)
	}
}
func (m *BaseService) GetCacheKey(table string, id interface{}) (key string) {
	key = table + "_" + gconv.String(id)
	return key
}

// func (m *BaseModelSev) getTreeList(list []*interface{}, pid int64) []*interface{} {
// 	res := make([]*interface{}, 0)
// 	for _, v := range list {
// 		if v.Pid == pid {
// 			v.Children = m.getTreeList(list, v.Value)
// 			res = append(res, v)
// 		}
// 	}
// 	return res
// }

// func (m *BaseModelSev) getTreeList(list []*request.PidTreeData, pid int64) []*request.PidTreeData {
// 	res := make([]*request.PidTreeData, 0)
// 	for _, v := range list {
// 		if v.Pid == pid {
// 			v.Children = m.getTreeList(list, v.Value)
// 			res = append(res, v)
// 		}
// 	}
// 	return res
// }
