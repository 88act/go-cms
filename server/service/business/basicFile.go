package business

import (
	"context"
	"errors"
	"go-cms/global"
	"go-cms/model/business"
	"go-cms/model/common/request"
	"go-cms/mycache"
	comSev "go-cms/service/common"
	"go-cms/utils"
	"sync"

	"gorm.io/gorm"
)

type BasicFileService struct {
	comSev.BaseService
}

var once_BasicFile sync.Once = sync.Once{}
var obj_BasicFileService *BasicFileService

// 获取单例
func GetBasicFileSev() *BasicFileService {
	once_BasicFile.Do(func() {
		obj_BasicFileService = new(BasicFileService)
		obj_BasicFileService.Db = global.DB
		//instanse.init()
	})
	return obj_BasicFileService
}

// Get 根据id获取BasicFile记录
// Author 10512203@qq.com
func (m *BasicFileService) GetByMd5(ctx context.Context, md5 string) (obj business.BasicFile, err error) {

	if utils.IsEmptyStr(md5) {
		return obj, errors.New("参数错误")
	}
	cacheKey := m.GetCacheKey("basic_file_", md5)
	if cacheData, err := mycache.GetCache().Get(cacheKey); err == nil {
		obj = cacheData.(business.BasicFile)
		return obj, err
	}
	err = m.Db.WithContext(ctx).Where("md5 = ?", md5).First(&obj).Error
	m.getFile(ctx, &obj)
	mycache.GetCache().Set(cacheKey, obj, obj.TableName())
	return obj, err
}

// Create 创建BasicFile记录
// Author 10512203@qq.com
func (m *BasicFileService) Create(ctx context.Context, data business.BasicFile) (id int64, err error) {
	err = m.Db.WithContext(ctx).Create(&data).Error
	if err != nil {
		return 0, err
	}
	return data.Id, err
}

// Delete 删除BasicFile记录
// Author 10512203@qq.com
func (m *BasicFileService) Delete(ctx context.Context, data business.BasicFile) (err error) {
	m.DelCache("basic_file", []int64{data.Id})
	err = m.Db.WithContext(ctx).Delete(&data).Error
	return err
}

// DeleteByIds 批量删除BasicFile记录
// Author 10512203@qq.com
func (m *BasicFileService) DeleteByIds(ctx context.Context, ids request.IdsReq) (err error) {
	m.DelCache("basic_file", ids.Ids)
	err = m.Db.WithContext(ctx).Delete(&[]business.BasicFile{}, "id in ?", ids.Ids).Error
	return err
}

// Update  更新BasicFile记录
// Author 10512203@qq.com
func (m *BasicFileService) Update(ctx context.Context, data business.BasicFile) (err error) {
	m.DelCache("basic_file", []int64{data.Id})
	err = m.Db.WithContext(ctx).Save(&data).Error
	return err
}

// UpdateByMap  更新BasicFile记录 by Map
// Author 10512203@qq.com
func (m *BasicFileService) UpdateByMap(ctx context.Context, data business.BasicFile, mapData map[string]interface{}) (err error) {
	m.DelCache("basic_file", []int64{data.Id})
	err = m.Db.WithContext(ctx).Model(&data).Updates(mapData).Error
	return err
}

// UpdateExpr 字段自增/自减
// Author 10512203@qq.com
func (m *BasicFileService) UpdateExpr(ctx context.Context, data business.BasicFile, column string, num int) (err error) {
	err = m.Db.WithContext(ctx).Model(&data).UpdateColumn(column, gorm.Expr(column+" + ?", num)).Error
	return err
}

// Get 根据id获取BasicFile记录
// Author 10512203@qq.com
func (m *BasicFileService) Get(ctx context.Context, id int64, fields string) (obj business.BasicFile, err error) {

	if id <= 0 {
		return obj, errors.New("参数id错误")
	}
	cacheKey := m.GetCacheKey("basic_file", id)
	if cacheData, err := mycache.GetCache().Get(cacheKey); err == nil {
		obj = cacheData.(business.BasicFile)
		return obj, err
	}
	if utils.IsEmptyStr(fields) {
		err = m.Db.WithContext(ctx).Where("id = ?", id).First(&obj).Error
	} else {
		err = m.Db.WithContext(ctx).Select(fields).Where("id = ?", id).First(&obj).Error
	}
	m.getFile(ctx, &obj)
	mycache.GetCache().Set(cacheKey, obj, obj.TableName())
	return obj, err
}

// GetList 分页获取BasicFile记录
// Author 10512203@qq.com
func (m *BasicFileService) GetList(ctx context.Context, seq business.BasicFileSearch, fields string) (list []business.BasicFile, total int64, err error) {
	limit := seq.PageSize
	offset := seq.PageSize * (seq.Page - 1)
	//修改 by ljd  增加查询排序
	order := seq.OrderKey
	desc := seq.OrderDesc
	// 创建db
	db := m.Db.WithContext(ctx).Model(&business.BasicFile{})

	//修改 by ljd
	if seq.Id > 0 {
		db = db.Where("id = ?", seq.Id)
	}

	if !utils.IsEmpty(seq.CreatedAtBegin) && !utils.IsEmpty(seq.CreatedAtEnd) {
		db = db.Where("created_at BETWEEN ? AND ?", seq.CreatedAtBegin, seq.CreatedAtEnd)
	}

	// 如果有条件搜索 下方会自动创建搜索语句
	if seq.Guid != "" {
		db = db.Where("guid = ?", seq.Guid)
	}
	if seq.UserId != nil {
		db = db.Where("user_id = ?", seq.UserId)
	}

	if seq.Name != "" {
		db = db.Where("name = ?", seq.Name)
	}
	if seq.CatId != nil {
		db = db.Where("cat_id = ?", seq.CatId)
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
	if seq.Path != "" {
		db = db.Where("path = ?", seq.Path)
	}
	if seq.FileType != "" {
		db = db.Where("file_type = ?", seq.FileType)
	}
	if seq.Ext != "" {
		db = db.Where("ext = ?", seq.Ext)
	}
	if seq.Size != nil {
		db = db.Where("size = ?", seq.Size)
	}
	if seq.Md5 != "" {
		db = db.Where("md5 = ?", seq.Md5)
	}
	if seq.Sha1 != "" {
		db = db.Where("sha1 = ?", seq.Sha1)
	}
	if seq.Sort != nil {
		db = db.Where("sort = ?", seq.Sort)
	}
	if seq.Download != nil {
		db = db.Where("download = ?", seq.Download)
	}
	if seq.UsedTime != nil {
		db = db.Where("used_time = ?", seq.UsedTime)
	}
	if seq.Status != nil {
		db = db.Where("status = ?", seq.Status)
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
	if utils.IsEmptyStr(fields) {
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&list).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&list).Error
	}
	m.getFileList(ctx, list)
	return list, total, err
}

// GetListAll 分页获取BasicFile记录 (全部字段)
// Author 10512203@qq.com
func (m *BasicFileService) GetListAll(ctx context.Context, seq business.BasicFileSearch, fields string) (list []business.BasicFile, total int64, err error) {
	return m.GetList(ctx, seq, "*")
}

// GetByMap 根据Map获取单一记录
// Author  [linjd] 10512203@qq.com
func (m *BasicFileService) GetByMap(ctx context.Context, mapData map[string]interface{}, fields string) (obj business.BasicFile, err error) {
	orderBy := "id desc"
	db := m.Db.WithContext(ctx).Model(&business.BasicFile{})
	if utils.IsEmptyStr(fields) {
		err = db.WithContext(ctx).Where(mapData).Order(orderBy).First(&obj).Error
	} else {
		err = db.WithContext(ctx).Select(fields).Order(orderBy).Where(mapData).First(&obj).Error
	}
	if err != nil {
		return obj, err
	}
	m.getFile(ctx, &obj)
	return obj, err
}

// GetByMap 根据Map获取列表
// Author 10512203@qq.com
func (m *BasicFileService) GetListByMap(ctx context.Context, mapData map[string]interface{}, fields string, orderBy string) (list []business.BasicFile, err error) {
	if utils.IsEmptyStr(orderBy) {
		orderBy = "id desc"
	}
	db := m.Db.WithContext(ctx).Model(&business.BasicFile{})
	if utils.IsEmptyStr(fields) {
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
func (m *BasicFileService) Count(ctx context.Context, mapDataWhere map[string]interface{}) (total int64, err error) {
	db := m.Db.WithContext(ctx).Model(&business.BasicFile{})
	err = db.Where(mapDataWhere).Count(&total).Error
	return total, err
}

// 如果有文件类型，更新文件 path
func (m *BasicFileService) getFileList(ctx context.Context, list []business.BasicFile) {
	for i, v := range list {

		//    if !utils.IsEmpty(v.FileList) {
		// 		objList := strings.Split(v.FileList, ",")
		// 		for _, guid := range objList {
		// 			if !utils.IsEmptyStr(guid) {
		// 				fileObj := global.FileObj{Guid: guid}
		// 				fileObj.Path, _ = m.GetPathByGuid(ctx, guid)
		// 				v.FileObjList = append(v.FileObjList, fileObj)
		// 			}
		// 		}
		// 	}
		list[i] = v
	}
}

// 如果有文件类型，更新文件 path
func (m *BasicFileService) getFile(ctx context.Context, v *business.BasicFile) {
	if v != nil {

		//   if !utils.IsEmpty(v.FileList) {
		// 		objList := strings.Split(v.FileList, ",")
		// 		for _, guid := range objList {
		// 			if !utils.IsEmptyStr(guid) {
		// 				fileObj := global.FileObj{Guid: guid}
		// 				fileObj.Path, _ = m.GetPathByGuid(ctx, guid)
		// 				v.FileObjList = append(v.FileObjList, fileObj)
		// 			}
		// 		}
		// 	}
	}
}
