package model

import (
	"context"
	"looklook/common/utils"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

type OrderOrderSev struct {
	db    *gorm.DB
	cache cache.ClusterConf
}

func NewOrderOrderSev(db *gorm.DB, cache cache.ClusterConf) *OrderOrderSev {
	instance := new(OrderOrderSev)
	instance.db = db
	instance.cache = cache
	return instance
}

// Create 创建记录
// Author [88act](https://github.com/88act)
func (m *OrderOrderSev) Create(ctx context.Context, data OrderOrder) (id int64, err error) {
	err = m.db.WithContext(ctx).Create(&data).Error
	if err != nil {
		return 0, err
	}
	return data.Id, err
}

// Delete 删除记录
// Author [88act](https://github.com/88act)
func (m *OrderOrderSev) Delete(ctx context.Context, data OrderOrder) (err error) {
	err = m.db.WithContext(ctx).Delete(&data).Error
	return err
}

// DeleteByIds 批量删除记录
// Author [88act](https://github.com/88act)
func (m *OrderOrderSev) DeleteByIds(ctx context.Context, ids IdsReq) (err error) {
	err = m.db.WithContext(ctx).Delete(&[]OrderOrder{}, "id in ?", ids.Ids).Error
	return err
}

// Update  更新记录
// Author [88act](https://github.com/88act)
func (m *OrderOrderSev) Update(ctx context.Context, data OrderOrder) (err error) {
	err = m.db.WithContext(ctx).Save(&data).Error
	return err
}

// UpdateByMap  更新记录 by Map
// Author [88act](https://github.com/88act)
func (m *OrderOrderSev) UpdateByMap(ctx context.Context, mapWhere map[string]interface{}, mapData map[string]interface{}) (rowsAffected int64, err error) {
	res := m.db.WithContext(ctx).Model(OrderOrder{}).Where(mapWhere).Updates(mapData)
	return res.RowsAffected, res.Error
}

// Get 根据id获取记录
// Author [88act](https://github.com/88act)
func (m *OrderOrderSev) Get(ctx context.Context, id int64, fields string) (obj *OrderOrder, err error) {

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
func (m *OrderOrderSev) GetByMap(ctx context.Context, mapData map[string]interface{}, fields string) (obj *OrderOrder, err error) {

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
func (m *OrderOrderSev) GetList(ctx context.Context, seq OrderOrderSearch, fields string) (list []OrderOrder, total int64, err error) {
	limit := seq.PageSize
	offset := seq.PageSize * (seq.Page - 1)
	order := seq.OrderKey
	desc := seq.OrderDesc

	db := m.db.WithContext(ctx).Model(&OrderOrder{})
	var orderOrders []OrderOrder

	if seq.Id > 0 {
		db = db.Where("`id` = ?", seq.Id)
	}
	if len(seq.CreatedAtBetween) >= 2 {
		db = db.Where("`created_at` BETWEEN ? AND ?", seq.CreatedAtBetween[0], seq.CreatedAtBetween[1])
	}
	// 如果有条件搜索
	if seq.Sn != "" {
		db = db.Where("`sn` = ?", seq.Sn)
	}
	if seq.UserId != nil {
		db = db.Where("`user_id` = ?", seq.UserId)
	}
	if seq.OrderType != nil {
		db = db.Where("`order_type` = ?", seq.OrderType)
	}
	if seq.ObjId != nil {
		db = db.Where("`obj_id` = ?", seq.ObjId)
	}
	if seq.OrderCode != "" {
		db = db.Where("`order_code` = ?", seq.OrderCode)
	}
	if seq.StatusPay != nil {
		db = db.Where("`status_pay` = ?", seq.StatusPay)
	}
	if seq.StatusOrder != nil {
		db = db.Where("`status_order` = ?", seq.StatusOrder)
	}
	if seq.Status != nil {
		db = db.Where("`status` = ?", seq.Status)
	}

	err = db.Count(&total).Error
	if err != nil {
		return orderOrders, 0, err
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&orderOrders).Error
	} else {
		err = db.Order(OrderStr).Find(&orderOrders).Error
	}
	//image类型,获取path
	for i, v := range orderOrders {
		v.MapData = make(map[string]string)
		orderOrders[i] = v
	}
	return orderOrders, total, err
}
