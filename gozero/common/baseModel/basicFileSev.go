package baseModel

import (
	"context"
	"go-cms/common/mycache"
	"go-cms/common/myconfig"
	"go-cms/common/utils"
	"strings"

	"gorm.io/gorm"
)

type BasicFileSev struct {
	BaseModelSev
}

func NewBasicFileSev(db *gorm.DB) *BasicFileSev {
	instance := new(BasicFileSev)
	instance.Db = db
	//instance.Cache = cache
	return instance
}

// Create 创建BasicFile记录
// Author 10512203@qq.com
func (m *BasicFileSev) Create(ctx context.Context, data BasicFile) (id int64, err error) {
	err = m.Db.WithContext(ctx).Create(&data).Error
	if err != nil {
		return 0, err
	}
	return data.Id, err
}

// Delete 删除BasicFile记录
// Author 10512203@qq.com
func (m *BasicFileSev) Delete(ctx context.Context, data BasicFile) (err error) {
	m.DelCache("basic_file", []int64{data.Id})
	err = m.Db.WithContext(ctx).Delete(&data).Error
	return err
}

// DeleteByIds 批量删除BasicFile记录
// Author 10512203@qq.com
func (m *BasicFileSev) DeleteByIds(ctx context.Context, ids IdsReq) (err error) {
	m.DelCache("basic_file", ids.Ids)
	err = m.Db.WithContext(ctx).Delete(&[]BasicFile{}, "id in ?", ids.Ids).Error
	return err
}

// Update  更新BasicFile记录
// Author 10512203@qq.com
func (m *BasicFileSev) Update(ctx context.Context, data BasicFile) (err error) {
	m.DelCache("basic_file", []int64{data.Id})
	err = m.Db.WithContext(ctx).Save(&data).Error
	return err
}

// UpdateByMap  更新BasicFile记录 by Map
//
//	values := map[string]interface{}{
//		"status":0,
//		"from": hash,
//	}
//
// Author 10512203@qq.com
func (m *BasicFileSev) UpdateByMap(ctx context.Context, data BasicFile, mapData map[string]interface{}) (err error) {
	m.DelCache("basic_file", []int64{data.Id})
	err = m.Db.WithContext(ctx).Model(&data).Updates(mapData).Error
	return err
}

// Get 根据id获取BasicFile记录
// Author 10512203@qq.com
func (m *BasicFileSev) Get(ctx context.Context, id int64, fields string) (obj BasicFile, err error) {

	if id <= 0 {
		return BasicFile{}, nil
	}
	cacheKey := m.GetCacheKey("basic_file", id)
	if cacheData, err := mycache.GetCache().Get(cacheKey); err == nil {
		obj = cacheData.(BasicFile)
		return obj, err
	}
	if utils.IsEmpty(fields) {
		err = m.Db.WithContext(ctx).Where("id = ?", id).First(&obj).Error
	} else {
		err = m.Db.WithContext(ctx).Select(fields).Where("id = ?", id).First(&obj).Error
	}
	m.getFile(ctx, &obj)
	mycache.GetCache().Set(cacheKey, obj, 0)
	return obj, err
}

// GetList 分页获取BasicFile记录
// Author 10512203@qq.com
func (m *BasicFileSev) GetList2(ctx context.Context, seq BasicFileSearch, fields string) (list []BasicFile, total int64, err error) {
	limit := seq.PageSize
	offset := seq.PageSize * (seq.Page - 1)
	//修改 by ljd  增加查询排序
	order := seq.OrderKey
	desc := seq.OrderDesc
	// 创建db
	db := m.Db.WithContext(ctx).Model(&BasicFile{})
	//强制状态为1
	db = db.Where("status = 1")
	if seq.Id > 0 {
		db = db.Where("id = ?", seq.Id)
	}
	if seq.CreatedAtBegin != nil && seq.CreatedAtEnd != nil {
		db = db.Where("created_at BETWEEN ? AND ?", seq.CreatedAtBegin.Format("2006-01-02 15:04:05"), seq.CreatedAtEnd.Format("2006-01-02 15:04:05"))
	}

	// 如果有条件搜索 下方会自动创建搜索语句
	if seq.Guid != "" {
		db = db.Where("guid = ?", seq.Guid)
	}
	if seq.UserIdSys != nil {
		db = db.Where("user_id_sys = ?", seq.UserIdSys)
	}
	if seq.UserId != nil {
		db = db.Where("user_id = ?", seq.UserId)
	}
	if seq.CatId != nil {
		db = db.Where("cat_id = ?", seq.CatId)
	}
	if seq.Name != "" {
		db = db.Where("name LIKE ?", "%"+seq.Name+"%")
	}
	if seq.Module != nil {
		db = db.Where("module = ?", seq.Module)
	}
	if seq.MediaType != nil {
		db = db.Where("media_type = ?", seq.MediaType)
	}
	if seq.Driver != nil {
		db = db.Where("driver = ?", seq.Driver)
	}
	if seq.Ext != "" {
		db = db.Where("ext = ?", seq.Ext)
	}
	if seq.FileType != "" {
		db = db.Where("file_type = ?", seq.FileType)
	}
	if seq.Size != nil {
		db = db.Where("size > ?", seq.Size)
	}
	if seq.Download != nil {
		db = db.Where("download > ?", seq.Download)
	}
	if seq.UsedTime != nil {
		db = db.Where("used_time > ?", seq.UsedTime)
	}
	if seq.Status != nil {
		db = db.Where("sttus = ?", seq.Status)
	}

	err = db.Count(&total).Error
	if err != nil {
		return list, 0, err
	}
	//err = db.Limit(limit).Offset(offset).Find(&list).Error
	//修改 by ljd  增加查询排序
	OrderStr := "id desc"
	if !utils.IsEmpty(order) {
		if desc {
			OrderStr = order + " desc"
		} else {
			OrderStr = order
		}
	}
	if utils.IsEmpty(fields) {
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&list).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&list).Error
	}
	m.getFileList(ctx, list)
	return list, total, err
}

// GetListAll 分页获取BasicFile记录 (全部字段)
// Author 10512203@qq.com
func (m *BasicFileSev) GetListAll(ctx context.Context, seq BasicFileSearch, fields string) (list []BasicFile, total int64, err error) {
	return m.GetList(ctx, seq, "*")
}

// GetByMap 根据Map获取单一记录
// Author  [linjd] 10512203@qq.com
func (m *BasicFileSev) GetByMap(ctx context.Context, mapData map[string]interface{}, fields string) (obj BasicFile, err error) {
	// 创建db
	db := m.Db.WithContext(ctx).Model(&BasicFile{})
	if utils.IsEmpty(fields) {
		err = db.WithContext(ctx).Where(mapData).First(&obj).Error
	} else {
		err = db.WithContext(ctx).Select(fields).Where(mapData).First(&obj).Error
	}
	if err != nil {
		return obj, err
	}
	m.getFile(ctx, &obj)
	return obj, err
}

// GetByMap 根据Map获取列表
// Author 10512203@qq.com
func (m *BasicFileSev) GetListByMap(ctx context.Context, mapData map[string]interface{}, fields string, orderBy string) (list []BasicFile, err error) {
	db := m.Db.WithContext(ctx).Model(&BasicFile{})
	if utils.IsEmpty(fields) {
		err = db.Where(mapData).Order(orderBy).Find(&list).Error
	} else {
		err = db.Select(fields).Where(mapData).Order(orderBy).Find(&list).Error
	}
	if err != nil {
		return nil, err
	}
	m.getFileList(ctx, list)
	return list, err
}

//	获取数量
//
// Author 10512203@qq.com
func (m *BasicFileSev) Count(ctx context.Context, mapDataWhere map[string]interface{}) (total int64, err error) {
	db := m.Db.WithContext(ctx).Model(&BasicFile{})
	err = db.Where(mapDataWhere).Count(&total).Error
	return total, err
}

// GetList 分页获取记录
// Author linjd  10512203@qq.com
func (m *BasicFileSev) GetList(ctx context.Context, seq BasicFileSearch, fields string) (list []BasicFile, total int64, err error) {
	limit := seq.PageSize
	offset := seq.PageSize * (seq.Page - 1)
	order := seq.OrderKey
	desc := seq.OrderDesc

	db := m.Db.WithContext(ctx).Model(&BasicFile{})
	db = db.Where("status = 1")
	if seq.Id > 0 {
		db = db.Where("id = ?", seq.Id)
	}

	if seq.Guid != "" {
		db = db.Where("guid = ?", seq.Guid)
	}
	// if seq.UserId != nil {
	// 	db = db.Where("user_id = ?", seq.UserId)
	// }
	// if seq.CatId != nil {
	// 	db = db.Where("cat_id = ?", seq.CatId)
	// }
	if seq.Name != "" {
		db = db.Where("name LIKE ?", "%"+seq.Name+"%")
	}
	// if seq.Module != nil {
	// 	db = db.Where("module = ?", seq.Module)
	// }
	// if seq.MediaType != nil {
	// 	db = db.Where("media_type = ?", seq.MediaType)
	// }
	// if seq.Driver != nil {
	// 	db = db.Where("driver = ?", seq.Driver)
	// }
	if seq.Ext != "" {
		db = db.Where("ext = ?", seq.Ext)
	}
	if seq.FileType != "" {
		db = db.Where("file_type = ?", seq.FileType)
	}
	// if seq.Size != nil {
	// 	db = db.Where("size > ?", seq.Size)
	// }
	// if seq.Download != nil {
	// 	db = db.Where("download > ?", seq.Download)
	// }
	// if seq.UsedTime != nil {
	// 	db = db.Where("used_time > ?", seq.UsedTime)
	// }

	err = db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	OrderStr := "id desc"
	if !utils.IsEmpty(order) {
		if desc {
			OrderStr = order + " desc"
		} else {
			OrderStr = order + " asc"
		}
	}
	if !utils.IsEmpty(fields) {
		db.Select(fields)
	}
	if limit > 0 {
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&list).Error
	} else {
		err = db.Order(OrderStr).Find(&list).Error
	}

	//如果有图片image类型，更新图片path
	for i, v := range list {
		list[i] = v
	}
	return list, total, err
}

// 如果有文件类型，更新文件 path
func (m *BasicFileSev) getFileList(ctx context.Context, list []BasicFile) {
	for i, v := range list {

		if !utils.IsEmpty(v.Guid) {
			v.Path, _ = m.GetPathByGuid(ctx, v.Guid)
		}
		list[i] = v
	}
}

// 如果有文件类型，更新文件 path
func (m *BasicFileSev) getFile(ctx context.Context, v *BasicFile) {
	if v != nil {
		if !utils.IsEmpty(v.Guid) {
			v.Path, _ = m.GetPathByGuid(ctx, v.Guid)
		}
	}
}

// Author linjd  10512203@qq.com
func (m *BasicFileSev) DeleteByGuids(ctx context.Context, where map[string]interface{}, guids []string) (err error) {
	err = m.Db.WithContext(ctx).Delete(&[]BasicFile{}, "guid in ?", guids).Error
	// err = m.Db.WithContext(ctx).Where(where).Delete(&[]BasicFile{}, "guid in ?", guids).Error
	return err
}

func (m *BasicFileSev) DeleteByGuidsSoft(ctx context.Context, where map[string]interface{}, guids []string) (err error) {
	// 文件删除 设置状态为 3
	err = m.Db.WithContext(ctx).Where("guid in ?", guids).Update("status", 3).Error
	// err = m.Db.WithContext(ctx).Where(where).Delete(&[]BasicFile{}, "guid in ?", guids).Error
	return err
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

// func (m *BasicFileSev) GetPathByGuid(ctx context.Context, guid string) (path string, err error) {
// 	//basicFS := new(businessSev.BasicFileService)
// 	// 先从 redis 获取 ，没有则读取数据库 ，缓存2小时
// 	if utils.IsEmpty(guid) {
// 		return "", nil
// 	}

// 	basicFile := BasicFile{}
// 	err = m.Db.WithContext(ctx).Where("guid = ? AND status =1", guid).First(&basicFile).Error
// 	if err != nil {
// 		return "", err
// 	}
// 	path = basicFile.Path
// 	if !utils.IsEmpty(path) {
// 		path = " global.CONFIG.Local.BaseUrl" + path
// 	}
// 	return path, err
// }

func (m *BasicFileSev) GetPathByGuid(ctx context.Context, guid string) (path string, err error) {
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
	err = m.Db.WithContext(ctx).Where("guid = ? ", guid).Select("path,guid").First(&basicFile).Error
	if err != nil {
		//cacheData = []byte("")
		mycache.GetCache().Set(cacheKey, "", 0)
		return "", err
	}
	path = basicFile.Path
	if !utils.IsEmpty(path) {
		path = myconfig.HttpRoot + path
		//path = "http://gozero.local.com/" + path
	}
	//cacheData = []byte(path)
	mycache.GetCache().Set(cacheKey, path, 0)
	//logx.Errorf("path =%s", path)
	return path, err
}
