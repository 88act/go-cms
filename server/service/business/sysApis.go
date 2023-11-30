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

type SysApisService struct {
	comSev.BaseService
}

var once_SysApis sync.Once = sync.Once{}
var obj_SysApisService *SysApisService

// 获取单例
func GetSysApisSev() *SysApisService {
	once_SysApis.Do(func() {
		obj_SysApisService = new(SysApisService)
		obj_SysApisService.Db = global.DB
		//instanse.init()
	})
	return obj_SysApisService
}

// Create 创建SysApis记录
// Author 10512203@qq.com
func (m *SysApisService) Create(ctx context.Context, data business.SysApis) (id int64, err error) {
	err = m.Db.WithContext(ctx).Create(&data).Error
	if err != nil {
		return 0, err
	}
	return data.Id, err
}

// Delete 删除SysApis记录
// Author 10512203@qq.com
func (m *SysApisService) Delete(ctx context.Context, data business.SysApis) (err error) {
	m.DelCache("sys_apis", []int64{data.Id})
	err = m.Db.WithContext(ctx).Delete(&data).Error
	return err
}

// DeleteByIds 批量删除SysApis记录
// Author 10512203@qq.com
func (m *SysApisService) DeleteByIds(ctx context.Context, ids request.IdsReq) (err error) {
	m.DelCache("sys_apis", ids.Ids)
	err = m.Db.WithContext(ctx).Delete(&[]business.SysApis{}, "id in ?", ids.Ids).Error
	return err
}

// Update  更新SysApis记录
// Author 10512203@qq.com
func (m *SysApisService) Update(ctx context.Context, data business.SysApis) (err error) {
	m.DelCache("sys_apis", []int64{data.Id})
	err = m.Db.WithContext(ctx).Save(&data).Error
	return err
}

// UpdateByMap  更新SysApis记录 by Map
// Author 10512203@qq.com
func (m *SysApisService) UpdateByMap(ctx context.Context, data business.SysApis, mapData map[string]interface{}) (err error) {
	m.DelCache("sys_apis", []int64{data.Id})
	err = m.Db.WithContext(ctx).Model(&data).Updates(mapData).Error
	return err
}

// UpdateExpr 字段自增/自减
// Author 10512203@qq.com
func (m *SysApisService) UpdateExpr(ctx context.Context, data business.SysApis, column string, num int) (err error) {
	err = m.Db.WithContext(ctx).Model(&data).UpdateColumn(column, gorm.Expr(column+" + ?", num)).Error
	return err
}

// Get 根据id获取SysApis记录
// Author 10512203@qq.com
func (m *SysApisService) Get(ctx context.Context, id int64, fields string) (obj business.SysApis, err error) {

	if id <= 0 {
		return obj, errors.New("参数id错误")
	}
	cacheKey := m.GetCacheKey("sys_apis", id)
	if cacheData, err := mycache.GetCache().Get(cacheKey); err == nil {
		obj = cacheData.(business.SysApis)
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

// GetList 分页获取SysApis记录
// Author 10512203@qq.com
func (m *SysApisService) GetList(ctx context.Context, seq business.SysApisSearch, fields string) (list []business.SysApis, total int64, err error) {
	limit := seq.PageSize
	offset := seq.PageSize * (seq.Page - 1)
	//修改 by ljd  增加查询排序
	order := seq.OrderKey
	desc := seq.OrderDesc
	// 创建db
	db := m.Db.WithContext(ctx).Model(&business.SysApis{})

	//修改 by ljd
	if seq.Id > 0 {
		db = db.Where("id = ?", seq.Id)
	}

	if !utils.IsEmpty(seq.CreatedAtBegin) && !utils.IsEmpty(seq.CreatedAtEnd) {
		db = db.Where("created_at BETWEEN ? AND ?", seq.CreatedAtBegin, seq.CreatedAtEnd)
	}

	// 如果有条件搜索 下方会自动创建搜索语句
	if seq.Path != "" {
		db = db.Where("path = ?", seq.Path)
	}
	if seq.Desc != "" {
		db = db.Where("desc = ?", seq.Desc)
	}
	if seq.Model != "" {
		db = db.Where("model = ?", seq.Model)
	}
	if seq.ApiGroup != "" {
		db = db.Where("api_group = ?", seq.ApiGroup)
	}
	if seq.Method != "" {
		db = db.Where("method = ?", seq.Method)
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

// GetListAll 分页获取SysApis记录 (全部字段)
// Author 10512203@qq.com
func (m *SysApisService) GetListAll(ctx context.Context, seq business.SysApisSearch, fields string) (list []business.SysApis, total int64, err error) {
	return m.GetList(ctx, seq, "*")
}

// GetByMap 根据Map获取单一记录
// Author  [linjd] 10512203@qq.com
func (m *SysApisService) GetByMap(ctx context.Context, mapData map[string]interface{}, fields string) (obj business.SysApis, err error) {
	orderBy := "id desc"
	db := m.Db.WithContext(ctx).Model(&business.SysApis{})
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
func (m *SysApisService) GetListByMap(ctx context.Context, mapData map[string]interface{}, fields string, orderBy string) (list []business.SysApis, err error) {
	if utils.IsEmptyStr(orderBy) {
		orderBy = "id desc"
	}
	db := m.Db.WithContext(ctx).Model(&business.SysApis{})
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
func (m *SysApisService) Count(ctx context.Context, mapDataWhere map[string]interface{}) (total int64, err error) {
	db := m.Db.WithContext(ctx).Model(&business.SysApis{})
	err = db.Where(mapDataWhere).Count(&total).Error
	return total, err
}

// 如果有文件类型，更新文件 path
func (m *SysApisService) getFileList(ctx context.Context, list []business.SysApis) {
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
func (m *SysApisService) getFile(ctx context.Context, v *business.SysApis) {
	if v != nil {

		// if !utils.IsEmpty(v.FileList) {
		// 	objList := strings.Split(v.FileList, ",")
		// 	for _, guid := range objList {
		// 		if !utils.IsEmptyStr(guid) {
		// 			fileObj := global.FileObj{Guid: guid}
		// 			fileObj.Path, _ = m.GetPathByGuid(ctx, guid)
		// 			v.FileObjList = append(v.FileObjList, fileObj)
		// 		}
		// 	}
		// }
	}
}
