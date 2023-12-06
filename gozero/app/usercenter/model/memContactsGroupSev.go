package model

import (
	"context"
	"go-cms/common/mycache"
	"go-cms/common/utils"

	. "go-cms/common/baseModel"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

type MemContactsGroupSev struct {
	BaseModelSev
}

func NewMemContactsGroupSev(db *gorm.DB, cache cache.ClusterConf) *MemContactsGroupSev {
	instance := new(MemContactsGroupSev)
	instance.Db = db
	instance.Cache = cache
	instance.CacheKeyPre = "MemContactsGroup"
	instance.CacheKeyListPre = "MemContactsGroupList"
	return instance
}

// Create 创建MemContactsGroup记录
// Author 10512203@qq.com
func (m *MemContactsGroupSev) Create(ctx context.Context, data MemContactsGroup) (id int64, err error) {
	err = m.Db.WithContext(ctx).Create(&data).Error
	if err == nil {
		mycache.GetCache().DeleteKeyPre(m.CacheKeyListPre)
		return data.Id, err
	} else {
		return 0, err
	}
}

// Delete 删除MemContactsGroup记录
// Author 10512203@qq.com
func (m *MemContactsGroupSev) Delete(ctx context.Context, data MemContactsGroup) (err error) {
	err = m.Db.WithContext(ctx).Delete(&data).Error
	if err == nil {
		m.DelCache(m.CacheKeyPre, []int64{data.Id})
	}
	return err
}

// DeleteByIds 批量删除MemContactsGroup记录
// Author 10512203@qq.com
func (m *MemContactsGroupSev) DeleteByIds(ctx context.Context, ids IdsReq) (err error) {
	err = m.Db.WithContext(ctx).Delete(&[]MemContactsGroup{}, "id in ?", ids.Ids).Error
	if err == nil {
		m.DelCache(m.CacheKeyPre, ids.Ids)
	}
	return err
}

// Update  更新MemContactsGroup记录
// Author 10512203@qq.com
func (m *MemContactsGroupSev) Update(ctx context.Context, data MemContactsGroup) (err error) {
	err = m.Db.WithContext(ctx).Save(&data).Error
	if err == nil {
		m.DelCache(m.CacheKeyPre, []int64{data.Id})
	}
	return err
}

// UpdateByMap  更新MemContactsGroup记录 by Map
// Author 10512203@qq.com
func (m *MemContactsGroupSev) UpdateByMap(ctx context.Context, data MemContactsGroup, mapData map[string]interface{}) (err error) {
	err = m.Db.WithContext(ctx).Model(&data).Updates(mapData).Error
	if err == nil {
		m.DelCache(m.CacheKeyPre, []int64{data.Id})
	}
	return err
}

// UpdateExpr 字段自增/自减
// Author 10512203@qq.com
func (m *MemContactsGroupSev) UpdateExpr(ctx context.Context, data MemContactsGroup, column string, num int) (err error) {
	err = m.Db.WithContext(ctx).Model(&data).UpdateColumn(column, gorm.Expr(column+" + ?", num)).Error
	return err
}

// Get 根据id获取MemContactsGroup记录
// Author 10512203@qq.com
func (m *MemContactsGroupSev) Get(ctx context.Context, id int64, fields string) (obj MemContactsGroup, err error) {

	if id <= 0 {
		return MemContactsGroup{}, nil
	}
	cacheKey := m.GetCacheKey(m.CacheKeyPre, id)
	if cacheData, err := mycache.GetCache().GetObj(cacheKey); err == nil {
		obj = cacheData.(MemContactsGroup)
		return obj, err
	}
	if utils.IsEmpty(fields) {
		err = m.Db.WithContext(ctx).Where("id = ?", id).First(&obj).Error
	} else {
		err = m.Db.WithContext(ctx).Select(fields).Where("id = ?", id).First(&obj).Error
	}
	if err == nil {
		m.getFile(ctx, &obj)
		mycache.GetCache().SetObj(cacheKey, obj, 0)
	}
	return obj, err
}

// GetList 分页获取MemContactsGroup记录
// Author 10512203@qq.com
func (m *MemContactsGroupSev) GetList(ctx context.Context, seq MemContactsGroupSearch, fields string) (list []MemContactsGroup, total int64, err error) {
	limit := seq.PageSize
	offset := seq.PageSize * (seq.Page - 1)
	//修改 by ljd  增加查询排序
	order := seq.OrderKey
	desc := seq.OrderDesc
	// 创建db
	db := m.Db.WithContext(ctx).Model(&MemContactsGroup{})
	//强制状态为1
	db = db.Where("status = 1")
	if seq.Id > 0 {
		db = db.Where("id = ?", seq.Id)
	}
	if seq.CreatedAtBegin != nil && seq.CreatedAtEnd != nil {
		db = db.Where("created_at BETWEEN ? AND ?", seq.CreatedAtBegin.Format("2006-01-02 15:04:05"), seq.CreatedAtEnd.Format("2006-01-02 15:04:05"))
	}

	// 如果有条件搜索 下方会自动创建搜索语句
	if seq.UserId != nil {
		db = db.Where("user_id = ?", seq.UserId)
	}
	if seq.Name != "" {
		db = db.Where("name LIKE ?", "%"+seq.Name+"%")
	}
	if seq.Star != nil {
		db = db.Where("star = ?", seq.Star)
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
	if utils.IsEmpty(fields) {
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&list).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&list).Error
	}
	m.getFileList(ctx, list)
	return list, total, err
}

// GetListAll 分页获取MemContactsGroup记录 (全部字段)
// Author 10512203@qq.com
func (m *MemContactsGroupSev) GetListAll(ctx context.Context, seq MemContactsGroupSearch, fields string) (list []MemContactsGroup, total int64, err error) {
	return m.GetList(ctx, seq, "*")
}

// GetByMap 根据Map获取单一记录
// Author  [linjd] 10512203@qq.com
func (m *MemContactsGroupSev) GetByMap(ctx context.Context, mapData map[string]interface{}, fields string) (obj MemContactsGroup, err error) {
	orderBy := "id desc"
	db := m.Db.WithContext(ctx).Model(&MemContactsGroup{})
	if utils.IsEmpty(fields) {
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
func (m *MemContactsGroupSev) GetListByMap(ctx context.Context, mapData map[string]interface{}, fields string, orderBy string) (list []MemContactsGroup, err error) {
	if utils.IsEmptyStr(orderBy) {
		orderBy = "id desc"
	}
	db := m.Db.WithContext(ctx).Model(&MemContactsGroup{})
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
func (m *MemContactsGroupSev) Count(ctx context.Context, mapDataWhere map[string]interface{}) (total int64, err error) {
	db := m.Db.WithContext(ctx).Model(&MemContactsGroup{})
	err = db.Where(mapDataWhere).Count(&total).Error
	return total, err
}

// 如果有文件类型，更新文件 path
func (m *MemContactsGroupSev) getFileList(ctx context.Context, list []MemContactsGroup) {
	for i, v := range list {
		list[i] = v
	}
}

// 如果有文件类型，更新文件 path
func (m *MemContactsGroupSev) getFile(ctx context.Context, v *MemContactsGroup) {
	if v != nil {
	}
}
