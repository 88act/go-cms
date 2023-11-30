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

type MemDepartService struct {
	comSev.BaseService
}

var once_MemDepart sync.Once = sync.Once{}
var obj_MemDepartService *MemDepartService

// 获取单例
func GetMemDepartSev() *MemDepartService {
	once_MemDepart.Do(func() {
		obj_MemDepartService = new(MemDepartService)
		obj_MemDepartService.Db = global.DB
		//instanse.init()
	})
	return obj_MemDepartService
}

// Create 创建MemDepart记录
// Author 10512203@qq.com
func (m *MemDepartService) Create(ctx context.Context, data business.MemDepart) (id int64, err error) {
	err = m.Db.WithContext(ctx).Create(&data).Error
	if err != nil {
		return 0, err
	}
	return data.Id, err
}

// Delete 删除MemDepart记录
// Author 10512203@qq.com
func (m *MemDepartService) Delete(ctx context.Context, data business.MemDepart) (err error) {
	m.DelCache("mem_depart", []int64{data.Id})
	err = m.Db.WithContext(ctx).Delete(&data).Error
	return err
}

// DeleteByIds 批量删除MemDepart记录
// Author 10512203@qq.com
func (m *MemDepartService) DeleteByIds(ctx context.Context, ids request.IdsReq) (err error) {
	m.DelCache("mem_depart", ids.Ids)
	err = m.Db.WithContext(ctx).Delete(&[]business.MemDepart{}, "id in ?", ids.Ids).Error
	return err
}

// Update  更新MemDepart记录
// Author 10512203@qq.com
func (m *MemDepartService) Update(ctx context.Context, data business.MemDepart) (err error) {
	m.DelCache("mem_depart", []int64{data.Id})
	err = m.Db.WithContext(ctx).Save(&data).Error
	return err
}

// UpdateByMap  更新MemDepart记录 by Map
// Author 10512203@qq.com
func (m *MemDepartService) UpdateByMap(ctx context.Context, data business.MemDepart, mapData map[string]interface{}) (err error) {
	m.DelCache("mem_depart", []int64{data.Id})
	err = m.Db.WithContext(ctx).Model(&data).Updates(mapData).Error
	return err
}

// UpdateExpr 字段自增/自减
// Author 10512203@qq.com
func (m *MemDepartService) UpdateExpr(ctx context.Context, data business.MemDepart, column string, num int) (err error) {
	err = m.Db.WithContext(ctx).Model(&data).UpdateColumn(column, gorm.Expr(column+" + ?", num)).Error
	return err
}

// Get 根据id获取MemDepart记录
// Author 10512203@qq.com
func (m *MemDepartService) Get(ctx context.Context, id int64, fields string) (obj business.MemDepart, err error) {

	if id <= 0 {
		return obj, errors.New("参数id错误")
	}
	cacheKey := m.GetCacheKey("mem_depart", id)
	if cacheData, err := mycache.GetCache().Get(cacheKey); err == nil {
		obj = cacheData.(business.MemDepart)
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

// GetList 分页获取MemDepart记录
// Author 10512203@qq.com
func (m *MemDepartService) GetList(ctx context.Context, seq business.MemDepartSearch, fields string) (list []*business.MemDepart, total int64, err error) {
	limit := seq.PageSize
	offset := seq.PageSize * (seq.Page - 1)
	//修改 by ljd  增加查询排序
	order := seq.OrderKey
	desc := seq.OrderDesc
	// 创建db
	db := m.Db.WithContext(ctx).Model(&business.MemDepart{})

	//修改 by ljd
	if seq.Id > 0 {
		db = db.Where("id = ?", seq.Id)
	}

	if !utils.IsEmpty(seq.CreatedAtBegin) && !utils.IsEmpty(seq.CreatedAtEnd) {
		db = db.Where("created_at BETWEEN ? AND ?", seq.CreatedAtBegin, seq.CreatedAtEnd)
	}

	if seq.Pid != nil {
		db = db.Where("pid = ?", seq.Pid)
	}
	if seq.UserId != nil {
		db = db.Where("user_id = ?", seq.UserId)
	}
	if seq.Name != "" {
		db = db.Where("name  LIKE ?", "%"+seq.Name+"%")
	}
	if seq.Encode != "" {
		db = db.Where("encode = ?", seq.Encode)
	}
	if seq.Desc != "" {
		db = db.Where("desc = ?", seq.Desc)
	}
	if seq.Address != "" {
		db = db.Where("address = ?", seq.Address)
	}
	if seq.Phone != "" {
		db = db.Where("phone = ?", seq.Phone)
	}
	if seq.Email != "" {
		db = db.Where("email = ?", seq.Email)
	}
	if seq.Sort != nil {
		db = db.Where("sort = ?", seq.Sort)
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
	//m.getFileList(ctx, list)
	return list, total, err
}

// GetListAll 分页获取MemDepart记录 (全部字段)
// Author 10512203@qq.com
func (m *MemDepartService) GetListAll(ctx context.Context, seq business.MemDepartSearch, fields string) (list []*business.MemDepart, total int64, err error) {
	return m.GetList(ctx, seq, "*")
}

// GetByMap 根据Map获取单一记录
// Author  [linjd] 10512203@qq.com
func (m *MemDepartService) GetByMap(ctx context.Context, mapData map[string]interface{}, fields string) (obj business.MemDepart, err error) {
	orderBy := "id desc"
	db := m.Db.WithContext(ctx).Model(&business.MemDepart{})
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
func (m *MemDepartService) GetListByMap(ctx context.Context, mapData map[string]interface{}, fields string, orderBy string) (list []business.MemDepart, err error) {
	if utils.IsEmptyStr(orderBy) {
		orderBy = "id desc"
	}
	db := m.Db.WithContext(ctx).Model(&business.MemDepart{})
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
func (m *MemDepartService) Count(ctx context.Context, mapDataWhere map[string]interface{}) (total int64, err error) {
	db := m.Db.WithContext(ctx).Model(&business.MemDepart{})
	err = db.Where(mapDataWhere).Count(&total).Error
	return total, err
}

// 如果有文件类型，更新文件 path
func (m *MemDepartService) getFileList(ctx context.Context, list []business.MemDepart) {
	for i, v := range list {
		if !utils.IsEmpty(v.Image) {
			fileObjImage := global.FileObj{Guid: v.Image}
			fileObjImage.Path, _ = m.GetPathByGuid(ctx, v.Image)
			v.FileObjList = append(v.FileObjList, fileObjImage)
		}

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
func (m *MemDepartService) getFile(ctx context.Context, v *business.MemDepart) {
	if v != nil {
		if !utils.IsEmpty(v.Image) {
			fileObjImage := global.FileObj{Guid: v.Image}
			fileObjImage.Path, _ = m.GetPathByGuid(ctx, v.Image)
			v.FileObjList = append(v.FileObjList, fileObjImage)
		}

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
