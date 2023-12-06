package model

import (
	"context"
	. "go-cms/common/baseModel"
	"go-cms/common/utils"

	"gorm.io/gorm"
)

type BasicSmsSev struct {
	db *gorm.DB
	//cache cache.ClusterConf
}

func NewBasicSmsSev(db *gorm.DB) *BasicSmsSev {
	instance := new(BasicSmsSev)
	instance.db = db
	//nstance.cache = cache
	return instance
}

// Create 创建记录
// Author 10512203@qq.com
func (m *BasicSmsSev) Create(ctx context.Context, data BasicSms) (id int64, err error) {
	err = m.db.WithContext(ctx).Create(&data).Error
	if err != nil {
		return 0, err
	}
	return data.Id, err
}

// Delete 删除记录
// Author 10512203@qq.com
func (m *BasicSmsSev) Delete(ctx context.Context, data BasicSms) (err error) {
	err = m.db.WithContext(ctx).Delete(&data).Error
	return err
}

// DeleteByIds 批量删除记录
// Author 10512203@qq.com
func (m *BasicSmsSev) DeleteByIds(ctx context.Context, ids IdsReq) (err error) {
	err = m.db.WithContext(ctx).Delete(&[]BasicSms{}, "id in ?", ids.Ids).Error
	return err
}

// Update  更新记录
// Author 10512203@qq.com
func (m *BasicSmsSev) Update(ctx context.Context, data BasicSms) (err error) {
	err = m.db.WithContext(ctx).Save(&data).Error
	return err
}

// UpdateByMap  更新记录 by Map
// Author 10512203@qq.com
func (m *BasicSmsSev) UpdateByMap(ctx context.Context, data BasicSms, mapData map[string]interface{}) (err error) {
	err = m.db.WithContext(ctx).Model(&data).Updates(mapData).Error
	return err
}

// Get 根据id获取记录
// Author 10512203@qq.com
func (m *BasicSmsSev) Get(ctx context.Context, id int64, fields string) (obj *BasicSms, err error) {

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
// Author 10512203@qq.com
func (m *BasicSmsSev) GetByMap(ctx context.Context, mapData map[string]interface{}, fields string) (obj *BasicSms, err error) {

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
// Author 10512203@qq.com
func (m *BasicSmsSev) GetList(ctx context.Context, seq BasicSmsSearch, fields string) (list []BasicSms, total int64, err error) {
	limit := seq.PageSize
	offset := seq.PageSize * (seq.Page - 1)
	order := seq.OrderKey
	desc := seq.OrderDesc

	db := m.db.WithContext(ctx).Model(&BasicSms{})
	var basicSmss []BasicSms

	if seq.Id > 0 {
		db = db.Where("`id` = ?", seq.Id)
	}

	// 如果有条件搜索
	if seq.UserId != nil {
		db = db.Where("`user_id` = ?", seq.UserId)
	}
	if seq.UserIdSend != nil {
		db = db.Where("`user_id_send` = ?", seq.UserIdSend)
	}
	if seq.Phone != "" {
		db = db.Where("`phone` = ?", seq.Phone)
	}
	if seq.Title != "" {
		db = db.Where("`title` LIKE ?", "%"+seq.Title+"%")
	}
	if seq.Content != "" {
		db = db.Where("`content` LIKE ?", "%"+seq.Content+"%")
	}
	if seq.TempletId != "" {
		db = db.Where("`templet_id` = ?", seq.TempletId)
	}
	if seq.VerilyCode != "" {
		db = db.Where("`verily_code` = ?", seq.VerilyCode)
	}
	if seq.Status != nil {
		db = db.Where("`status` = ?", seq.Status)
	}

	err = db.Count(&total).Error
	if err != nil {
		return basicSmss, 0, err
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&basicSmss).Error
	} else {
		err = db.Order(OrderStr).Find(&basicSmss).Error
	}
	//image类型,获取path
	for i, v := range basicSmss {
		v.MapData = make(map[string]string)
		basicSmss[i] = v
	}
	return basicSmss, total, err
}
