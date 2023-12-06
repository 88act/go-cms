package model

import (
	"context"
	"go-cms/common/mycache"
	"go-cms/common/utils"

	. "go-cms/common/baseModel"

	"gorm.io/gorm"
)

type CmsCatSev struct {
	BaseModelSev
}

func NewCmsCatSev(db *gorm.DB) *CmsCatSev {
	instance := new(CmsCatSev)
	instance.Db = db
	instance.CacheKeyPre = "CmsCat"
	instance.CacheKeyListPre = "CmsCatList"
	return instance
}

// Create 创建CmsCat记录
// Author 10512203@qq.com
func (m *CmsCatSev) Create(ctx context.Context, data CmsCat) (id int64, err error) {
	err = m.Db.WithContext(ctx).Create(&data).Error
	if err == nil {
		mycache.GetCache().DeleteKeyPre(m.CacheKeyListPre)
		return data.Id, err
	} else {
		return 0, err
	}
}

// Delete 删除CmsCat记录
// Author 10512203@qq.com
func (m *CmsCatSev) Delete(ctx context.Context, data CmsCat) (err error) {
	err = m.Db.WithContext(ctx).Delete(&data).Error
	if err == nil {
		m.DelCache(m.CacheKeyPre, []int64{data.Id})
	}
	return err
}

// DeleteByIds 批量删除CmsCat记录
// Author 10512203@qq.com
func (m *CmsCatSev) DeleteByIds(ctx context.Context, ids IdsReq) (err error) {
	err = m.Db.WithContext(ctx).Delete(&[]CmsCat{}, "id in ?", ids.Ids).Error
	if err == nil {
		m.DelCache(m.CacheKeyPre, ids.Ids)
	}
	return err
}

// Update  更新CmsCat记录
// Author 10512203@qq.com
func (m *CmsCatSev) Update(ctx context.Context, data CmsCat) (err error) {
	err = m.Db.WithContext(ctx).Save(&data).Error
	if err == nil {
		m.DelCache(m.CacheKeyPre, []int64{data.Id})
	}
	return err
}

// UpdateByMap  更新CmsCat记录 by Map
// Author 10512203@qq.com
func (m *CmsCatSev) UpdateByMap(ctx context.Context, data CmsCat, mapData map[string]interface{}) (err error) {
	err = m.Db.WithContext(ctx).Model(&data).Updates(mapData).Error
	if err == nil {
		m.DelCache(m.CacheKeyPre, []int64{data.Id})
	}
	return err
}

// UpdateExpr 字段自增/自减
// Author 10512203@qq.com
func (m *CmsCatSev) UpdateExpr(ctx context.Context, data CmsCat, column string, num int) (err error) {
	err = m.Db.WithContext(ctx).Model(&data).UpdateColumn(column, gorm.Expr(column+" + ?", num)).Error
	return err
}

// Get 根据id获取CmsCat记录
// Author 10512203@qq.com
func (m *CmsCatSev) Get(ctx context.Context, id int64, fields string) (obj CmsCat, err error) {

	if id <= 0 {
		return CmsCat{}, nil
	}
	cacheKey := m.GetCacheKey(m.CacheKeyPre, id)
	if cacheData, err := mycache.GetCache().GetObj(cacheKey); err == nil {
		obj = cacheData.(CmsCat)
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

// GetList 分页获取CmsCat记录
// Author 10512203@qq.com
func (m *CmsCatSev) GetList(ctx context.Context, seq CmsCatSearch, fields string) (list []CmsCat, total int64, err error) {
	limit := seq.PageSize
	offset := seq.PageSize * (seq.Page - 1)
	//修改 by ljd  增加查询排序
	order := seq.OrderKey
	desc := seq.OrderDesc
	// 创建db
	db := m.Db.WithContext(ctx).Model(&CmsCat{})
	//强制状态为1
	db = db.Where("status = 1")
	if seq.Id > 0 {
		db = db.Where("id = ?", seq.Id)
	}
	if seq.CreatedAtBegin != nil && seq.CreatedAtEnd != nil {
		db = db.Where("created_at BETWEEN ? AND ?", seq.CreatedAtBegin.Format("2006-01-02 15:04:05"), seq.CreatedAtEnd.Format("2006-01-02 15:04:05"))
	}

	// 如果有条件搜索 下方会自动创建搜索语句
	if seq.Pid != nil {
		db = db.Where("pid = ?", seq.Pid)
	}
	if seq.UserId != nil {
		db = db.Where("user_id = ?", seq.UserId)
	}
	if seq.BeSys != nil {
		db = db.Where("be_sys = ?", seq.BeSys)
	}
	if seq.ObjId != nil {
		db = db.Where("obj_id = ?", seq.ObjId)
	}
	if seq.Type != nil {
		db = db.Where("type = ?", seq.Type)
	}
	if seq.Title != "" {
		db = db.Where("title LIKE ?", "%"+seq.Title+"%")
	}
	if seq.Desc != "" {
		db = db.Where("desc LIKE ?", "%"+seq.Desc+"%")
	}
	if seq.TagList != "" {
		db = db.Where("tag_list LIKE ?", "%"+seq.TagList+"%")
	}
	if seq.BeTop != nil {
		db = db.Where("be_top = ?", seq.BeTop)
	}
	if seq.BeNav != nil {
		db = db.Where("be_nav = ?", seq.BeNav)
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
	if utils.IsEmpty(fields) {
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&list).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&list).Error
	}
	m.getFileList(ctx, list)
	return list, total, err
}

// GetListAll 分页获取CmsCat记录 (全部字段)
// Author 10512203@qq.com
func (m *CmsCatSev) GetListAll(ctx context.Context, seq CmsCatSearch, fields string) (list []CmsCat, total int64, err error) {
	return m.GetList(ctx, seq, "*")
}

// GetByMap 根据Map获取单一记录
// Author  [linjd] 10512203@qq.com
func (m *CmsCatSev) GetByMap(ctx context.Context, mapData map[string]interface{}, fields string) (obj CmsCat, err error) {
	orderBy := "id desc"
	db := m.Db.WithContext(ctx).Model(&CmsCat{})
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
func (m *CmsCatSev) GetListByMap(ctx context.Context, mapData map[string]interface{}, fields string, orderBy string) (list []CmsCat, err error) {
	if utils.IsEmptyStr(orderBy) {
		orderBy = "id desc"
	}
	db := m.Db.WithContext(ctx).Model(&CmsCat{})
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
func (m *CmsCatSev) Count(ctx context.Context, mapDataWhere map[string]interface{}) (total int64, err error) {
	db := m.Db.WithContext(ctx).Model(&CmsCat{})
	err = db.Where(mapDataWhere).Count(&total).Error
	return total, err
}

// 如果有文件类型，更新文件 path
func (m *CmsCatSev) getFileList(ctx context.Context, list []CmsCat) {
	for i, v := range list {
		if !utils.IsEmpty(v.Image) {
			v.Image, _ = m.GetPathByGuid(ctx, v.Image)
		}
		list[i] = v
	}
}

// 如果有文件类型，更新文件 path
func (m *CmsCatSev) getFile(ctx context.Context, v *CmsCat) {
	if v != nil {
		if !utils.IsEmpty(v.Image) {
			v.Image, _ = m.GetPathByGuid(ctx, v.Image)
		}

	}
}
