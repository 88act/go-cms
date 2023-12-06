package model

import (
	"context"
	"go-cms/common/mycache"
	"go-cms/common/utils"

	. "go-cms/common/baseModel"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

type MemContactsSev struct {
	BaseModelSev
}

func NewMemContactsSev(db *gorm.DB, cache cache.ClusterConf) *MemContactsSev {
	instance := new(MemContactsSev)
	instance.Db = db
	instance.Cache = cache
	instance.CacheKeyPre = "MemContacts"
	instance.CacheKeyListPre = "MemContactsList"
	return instance
}

// Create 创建MemContacts记录
// Author 10512203@qq.com
func (m *MemContactsSev) Create(ctx context.Context, data MemContacts) (id int64, err error) {
	err = m.Db.WithContext(ctx).Create(&data).Error
	if err == nil {
		mycache.GetCache().DeleteKeyPre(m.CacheKeyListPre)
		return data.Id, err
	} else {
		return 0, err
	}
}

// Delete 删除MemContacts记录
// Author 10512203@qq.com
func (m *MemContactsSev) Delete(ctx context.Context, data MemContacts) (err error) {
	err = m.Db.WithContext(ctx).Delete(&data).Error
	if err == nil {
		m.DelCache(m.CacheKeyPre, []int64{data.Id})
	}
	return err
}

// DeleteByIds 批量删除MemContacts记录
// Author 10512203@qq.com
func (m *MemContactsSev) DeleteByIds(ctx context.Context, ids IdsReq) (err error) {
	err = m.Db.WithContext(ctx).Delete(&[]MemContacts{}, "id in ?", ids.Ids).Error
	if err == nil {
		m.DelCache(m.CacheKeyPre, ids.Ids)
	}
	return err
}

// Update  更新MemContacts记录
// Author 10512203@qq.com
func (m *MemContactsSev) Update(ctx context.Context, data MemContacts) (err error) {
	err = m.Db.WithContext(ctx).Save(&data).Error
	if err == nil {
		m.DelCache(m.CacheKeyPre, []int64{data.Id})
	}
	return err
}

// UpdateByMap  更新MemContacts记录 by Map
// Author 10512203@qq.com
func (m *MemContactsSev) UpdateByMap(ctx context.Context, data MemContacts, mapData map[string]interface{}) (err error) {
	err = m.Db.WithContext(ctx).Model(&data).Updates(mapData).Error
	if err == nil {
		m.DelCache(m.CacheKeyPre, []int64{data.Id})
	}
	return err
}

// UpdateExpr 字段自增/自减
// Author 10512203@qq.com
func (m *MemContactsSev) UpdateExpr(ctx context.Context, data MemContacts, column string, num int) (err error) {
	err = m.Db.WithContext(ctx).Model(&data).UpdateColumn(column, gorm.Expr(column+" + ?", num)).Error
	return err
}

// Get 根据id获取MemContacts记录
// Author 10512203@qq.com
func (m *MemContactsSev) Get(ctx context.Context, id int64, fields string) (obj MemContacts, err error) {

	if id <= 0 {
		return MemContacts{}, nil
	}
	cacheKey := m.GetCacheKey(m.CacheKeyPre, id)
	if cacheData, err := mycache.GetCache().GetObj(cacheKey); err == nil {
		obj = cacheData.(MemContacts)
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

// GetList 分页获取MemContacts记录
// Author 10512203@qq.com
func (m *MemContactsSev) GetList(ctx context.Context, seq MemContactsSearch, fields string) (list []MemContacts, total int64, err error) {
	limit := seq.PageSize
	offset := seq.PageSize * (seq.Page - 1)
	//修改 by ljd  增加查询排序
	order := seq.OrderKey
	desc := seq.OrderDesc
	// 创建db
	db := m.Db.WithContext(ctx).Model(&MemContacts{})
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
	if seq.Realname != "" {
		db = db.Where("realname LIKE ?", "%"+seq.Realname+"%")
	}
	if seq.Area != "" {
		db = db.Where("area = ?", seq.Area)
	}
	if seq.Mobile != "" {
		db = db.Where("mobile LIKE ?", "%"+seq.Mobile+"%")
	}
	if seq.Phone != "" {
		db = db.Where("phone LIKE ?", "%"+seq.Phone+"%")
	}
	if seq.Email != "" {
		db = db.Where("email = ?", seq.Email)
	}
	if seq.Group != nil {
		db = db.Where("group = ?", seq.Group)
	}
	if seq.Star != nil {
		db = db.Where("star = ?", seq.Star)
	}
	if seq.LastTime != nil {
		db = db.Where("last_time = > ?", seq.LastTime)
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

// GetListAll 分页获取MemContacts记录 (全部字段)
// Author 10512203@qq.com
func (m *MemContactsSev) GetListAll(ctx context.Context, seq MemContactsSearch, fields string) (list []MemContacts, total int64, err error) {
	return m.GetList(ctx, seq, "*")
}

// GetByMap 根据Map获取单一记录
// Author  [linjd] 10512203@qq.com
func (m *MemContactsSev) GetByMap(ctx context.Context, mapData map[string]interface{}, fields string) (obj MemContacts, err error) {
	orderBy := "id desc"
	db := m.Db.WithContext(ctx).Model(&MemContacts{})
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
func (m *MemContactsSev) GetListByMap(ctx context.Context, mapData map[string]interface{}, fields string, orderBy string) (list []MemContacts, err error) {
	if utils.IsEmptyStr(orderBy) {
		orderBy = "id desc"
	}
	db := m.Db.WithContext(ctx).Model(&MemContacts{})
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
func (m *MemContactsSev) Count(ctx context.Context, mapDataWhere map[string]interface{}) (total int64, err error) {
	db := m.Db.WithContext(ctx).Model(&MemContacts{})
	err = db.Where(mapDataWhere).Count(&total).Error
	return total, err
}

// 如果有文件类型，更新文件 path
func (m *MemContactsSev) getFileList(ctx context.Context, list []MemContacts) {
	for i, v := range list {
		list[i] = v
	}
}

// 如果有文件类型，更新文件 path
func (m *MemContactsSev) getFile(ctx context.Context, v *MemContacts) {
	if v != nil {
	}
}
