package model

import (
	"context"
	"looklook/common/utils"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

type PayPaymentSev struct {
	db    *gorm.DB
	cache cache.ClusterConf
}

func NewPayPaymentSev(db *gorm.DB, cache cache.ClusterConf) *PayPaymentSev {
	instance := new(PayPaymentSev)
	instance.db = db
	instance.cache = cache
	return instance
}

// Create 创建记录
// Author [88act](https://github.com/88act)
func (m *PayPaymentSev) Create(ctx context.Context, data PayPayment) (id int64, err error) {
	err = m.db.WithContext(ctx).Create(&data).Error
	if err != nil {
		return 0, err
	}
	return data.Id, err
}

// Delete 删除记录
// Author [88act](https://github.com/88act)
func (m *PayPaymentSev) Delete(ctx context.Context, data PayPayment) (err error) {
	err = m.db.WithContext(ctx).Delete(&data).Error
	return err
}

// DeleteByIds 批量删除记录
// Author [88act](https://github.com/88act)
func (m *PayPaymentSev) DeleteByIds(ctx context.Context, ids IdsReq) (err error) {
	err = m.db.WithContext(ctx).Delete(&[]PayPayment{}, "id in ?", ids.Ids).Error
	return err
}

// Update  更新记录
// Author [88act](https://github.com/88act)
func (m *PayPaymentSev) Update(ctx context.Context, data PayPayment) (err error) {
	err = m.db.WithContext(ctx).Save(&data).Error
	return err
}

// UpdateByMap  更新记录 by Map
// Author [88act](https://github.com/88act)
func (m *PayPaymentSev) UpdateByMap(ctx context.Context, data PayPayment, mapData map[string]interface{}) (err error) {
	err = m.db.WithContext(ctx).Model(&data).Updates(mapData).Error
	return err
}

// Get 根据id获取记录
// Author [88act](https://github.com/88act)
func (m *PayPaymentSev) Get(ctx context.Context, id int64, fields string) (obj *PayPayment, err error) {

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
func (m *PayPaymentSev) GetByMap(ctx context.Context, mapData map[string]interface{}, fields string) (obj *PayPayment, err error) {

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
func (m *PayPaymentSev) GetList(ctx context.Context, seq PayPaymentSearch, fields string) (list []PayPayment, total int64, err error) {
	limit := seq.PageSize
	offset := seq.PageSize * (seq.Page - 1)
	order := seq.OrderKey
	desc := seq.OrderDesc

	db := m.db.WithContext(ctx).Model(&PayPayment{})
	var payPayments []PayPayment

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
	if seq.PayMode != "" {
		db = db.Where("`pay_mode` = ?", seq.PayMode)
	}
	if seq.TradeType != "" {
		db = db.Where("`trade_type` = ?", seq.TradeType)
	}
	if seq.TradeState != "" {
		db = db.Where("`trade_state` = ?", seq.TradeState)
	}
	if seq.PayTotal != nil {
		db = db.Where("`pay_total` = ?", seq.PayTotal)
	}
	if seq.TransactionId != "" {
		db = db.Where("`transaction_id` = ?", seq.TransactionId)
	}
	if seq.OrderSn != "" {
		db = db.Where("`order_sn` = ?", seq.OrderSn)
	}
	if seq.ServiceType != "" {
		db = db.Where("`service_type` = ?", seq.ServiceType)
	}
	if seq.PayStatus != nil {
		db = db.Where("`pay_status` = ?", seq.PayStatus)
	}
	if seq.PayTime != nil {
		db = db.Where("`pay_time` = > ?", seq.PayTime)
	}
	if seq.Status != nil {
		db = db.Where("`status` = ?", seq.Status)
	}

	err = db.Count(&total).Error
	if err != nil {
		return payPayments, 0, err
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&payPayments).Error
	} else {
		err = db.Order(OrderStr).Find(&payPayments).Error
	}
	//image类型,获取path
	for i, v := range payPayments {
		v.MapData = make(map[string]string)
		payPayments[i] = v
	}
	return payPayments, total, err
}
