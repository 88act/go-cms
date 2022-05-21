package model

import (
	"context"
	"go-cms/common/utils"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

type BasicFileSev struct {
	db    *gorm.DB
	cache cache.ClusterConf
}

func NewBasicFileSev(db *gorm.DB, cache cache.ClusterConf) *BasicFileSev {
	instance := new(BasicFileSev)
	instance.db = db
	instance.cache = cache
	return instance
}

// Create 创建记录
// Author [88act](https://github.com/88act)
func (m *BasicFileSev) Create(ctx context.Context, data BasicFile) (id int64, err error) {
	err = m.db.WithContext(ctx).Create(&data).Error
	if err != nil {
		return 0, err
	}
	return data.Id, err
}

// Delete 删除记录
// Author [88act](https://github.com/88act)
func (m *BasicFileSev) Delete(ctx context.Context, data BasicFile) (err error) {
	err = m.db.WithContext(ctx).Delete(&data).Error
	return err
}

// DeleteByIds 批量删除记录
// Author [88act](https://github.com/88act)
func (m *BasicFileSev) DeleteByIds(ctx context.Context, ids IdsReq) (err error) {
	err = m.db.WithContext(ctx).Delete(&[]BasicFile{}, "id in ?", ids.Ids).Error
	return err
}

// Update  更新记录
// Author [88act](https://github.com/88act)
func (m *BasicFileSev) Update(ctx context.Context, data BasicFile) (err error) {
	err = m.db.WithContext(ctx).Save(&data).Error
	return err
}

// UpdateByMap  更新记录 by Map
// Author [88act](https://github.com/88act)
func (m *BasicFileSev) UpdateByMap(ctx context.Context, data BasicFile, mapData map[string]interface{}) (err error) {
	err = m.db.WithContext(ctx).Model(&data).Updates(mapData).Error
	return err
}

// Get 根据id获取记录
// Author [88act](https://github.com/88act)
func (m *BasicFileSev) Get(ctx context.Context, id int64, fields string) (obj *BasicFile, err error) {

	if utils.IsEmpty(fields) {
		err = m.db.WithContext(ctx).Where("id = ?", id).First(&obj).Error
	} else {
		err = m.db.WithContext(ctx).Select(fields).Where("id = ?", id).First(&obj).Error
	}
	if err != nil {
		return nil, err
	}

	// image类型 图片path
	obj.MapData = make(map[string]string)
	return obj, err
}

// GetByMap 根据Map获取记录
// Author [88act](https://github.com/88act)
func (m *BasicFileSev) GetByMap(ctx context.Context, mapData map[string]interface{}, fields string) (obj *BasicFile, err error) {

	if utils.IsEmpty(fields) {
		err = m.db.WithContext(ctx).Where(mapData).First(&obj).Error
	} else {
		err = m.db.WithContext(ctx).Select(fields).Where(mapData).First(&obj).Error
	}
	if err != nil {
		return nil, err
	}

	// image类型 图片path
	obj.MapData = make(map[string]string)
	return obj, err
}

// GetList 分页获取记录
// Author [88act](https://github.com/88act)
func (m *BasicFileSev) GetList(ctx context.Context, seq BasicFileSearch, fields string) (list []BasicFile, total int64, err error) {
	limit := seq.PageSize
	offset := seq.PageSize * (seq.Page - 1)
	order := seq.OrderKey
	desc := seq.OrderDesc

	db := m.db.WithContext(ctx).Model(&BasicFile{})
	var basicFiles []BasicFile

	if seq.Id > 0 {
		db = db.Where("`id` = ?", seq.Id)
	}
	if len(seq.CreatedAtBetween) >= 2 {
		db = db.Where("`created_at` BETWEEN ? AND ?", seq.CreatedAtBetween[0], seq.CreatedAtBetween[1])
	}
	// 如果有条件搜索
	if seq.Guid != "" {
		db = db.Where("`guid` = ?", seq.Guid)
	}
	if seq.UserId != nil {
		db = db.Where("`user_id` = ?", seq.UserId)
	}
	if seq.Name != "" {
		db = db.Where("`name` LIKE ?", "%"+seq.Name+"%")
	}
	if seq.Module != nil {
		db = db.Where("`module` = ?", seq.Module)
	}
	if seq.MediaType != nil {
		db = db.Where("`media_type` = ?", seq.MediaType)
	}
	if seq.Driver != nil {
		db = db.Where("`driver` = ?", seq.Driver)
	}
	if seq.Ext != "" {
		db = db.Where("`ext` = ?", seq.Ext)
	}
	if seq.Size != nil {
		db = db.Where("`size` > ?", seq.Size)
	}
	if seq.Download != nil {
		db = db.Where("`download` > ?", seq.Download)
	}
	if seq.UsedTime != nil {
		db = db.Where("`used_time` > ?", seq.UsedTime)
	}

	err = db.Count(&total).Error
	if err != nil {
		return basicFiles, 0, err
	}
	OrderStr := "id desc"
	if !utils.IsEmpty(order) {
		if desc {
			OrderStr = order + " desc"
		} else {
			OrderStr = order
		}
	}
	if !utils.IsEmpty(fields) {
		db.Select(fields)
	}
	if limit > 0 {
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&basicFiles).Error
	} else {
		err = db.Order(OrderStr).Find(&basicFiles).Error
	}
	//image类型,获取path
	for i, v := range basicFiles {
		v.MapData = make(map[string]string)
		basicFiles[i] = v
	}
	return basicFiles, total, err
}
