package business

import (
	"go-cms/global"
	"go-cms/model/business"
	bizReq "go-cms/model/business/request"
	"go-cms/model/common/request"
	"go-cms/utils"
	"sync"
)

type OrderOrderService struct {
}

var once_OrderOrder sync.Once = sync.Once{}
var obj_OrderOrderService *OrderOrderService

//获取单例
func GetOrderOrderSev() *OrderOrderService {
	once_OrderOrder.Do(func() {
		obj_OrderOrderService = new(OrderOrderService)
		//instanse.init()
	})
	return obj_OrderOrderService
}

// Create 创建OrderOrder记录
// Author [88act](https://github.com/88act)
func (m *OrderOrderService) Create(data business.OrderOrder) (id uint, err error) {
	err = global.DB.Create(&data).Error
	if err != nil {
		return 0, err
	}
	return data.ID, err
}

// Delete 删除OrderOrder记录
// Author [88act](https://github.com/88act)
func (m *OrderOrderService) Delete(data business.OrderOrder) (err error) {
	err = global.DB.Delete(&data).Error
	return err
}

// DeleteByIds 批量删除OrderOrder记录
// Author [88act](https://github.com/88act)
func (m *OrderOrderService) DeleteByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]business.OrderOrder{}, "id in ?", ids.Ids).Error
	return err
}

// Update  更新OrderOrder记录
// Author [88act](https://github.com/88act)
func (m *OrderOrderService) Update(data business.OrderOrder) (err error) {
	err = global.DB.Save(&data).Error
	return err
}

// UpdateByMap  更新OrderOrder记录 by Map
// values := map[string]interface{}{
// 	"status":0,
// 	"from": hash,
// }
// Author [88act](https://github.com/88act)
func (m *OrderOrderService) UpdateByMap(data business.OrderOrder, mapData map[string]interface{}) (err error) {
	err = global.DB.Model(&data).Updates(mapData).Error
	return err
}

// Get 根据id获取OrderOrder记录
// Author [88act](https://github.com/88act)
func (m *OrderOrderService) Get(id uint, fields string) (obj business.OrderOrder, err error) {

	if utils.IsEmpty(fields) {
		err = global.DB.Where("id = ?", id).First(&obj).Error
	} else {
		err = global.DB.Select(fields).Where("id = ?", id).First(&obj).Error
	}

	//如果有图片image类型，更新图片path
	obj.MapData = make(map[string]string)
	return obj, err
}

// GetList 分页获取OrderOrder记录
// Author [88act](https://github.com/88act)
func (m *OrderOrderService) GetList(info bizReq.OrderOrderSearch, createdAtBetween []string, fields string) (list []business.OrderOrderMini, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.OrderOrder{})
	//var orderOrders []business.OrderOrder
	var orderOrders []business.OrderOrderMini

	//修改 by ljd
	if info.ID > 0 {
		db = db.Where("`id` = ?", info.ID)
	}
	if len(createdAtBetween) >= 2 {
		db = db.Where("`created_at` BETWEEN ? AND ?", createdAtBetween[0], createdAtBetween[1])
	}

	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Sn != "" {
		db = db.Where("`sn` = ?", info.Sn)
	}
	if info.UserId != nil {
		db = db.Where("`user_id` = ?", info.UserId)
	}
	if info.OrderType != nil {
		db = db.Where("`order_type` = ?", info.OrderType)
	}
	if info.ObjId != nil {
		db = db.Where("`obj_id` = ?", info.ObjId)
	}
	if info.OrderCode != "" {
		db = db.Where("`order_code` = ?", info.OrderCode)
	}
	if info.StatusPay != nil {
		db = db.Where("`status_pay` = ?", info.StatusPay)
	}
	if info.StatusOrder != nil {
		db = db.Where("`status_order` = ?", info.StatusOrder)
	}
	if info.Status != nil {
		db = db.Where("`status` = ?", info.Status)
	}

	err = db.Count(&total).Error
	if err != nil {
		return orderOrders, 0, err
	}
	//err = db.Limit(limit).Offset(offset).Find(&orderOrders).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&orderOrders).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&orderOrders).Error
	}
	//如果有图片image类型，更新图片path
	for i, v := range orderOrders {
		v.MapData = make(map[string]string)
		orderOrders[i] = v
	}
	return orderOrders, total, err
}

//GetListAll 分页获取OrderOrder记录 (全部字段)
// Author [88act](https://github.com/88act)
func (m *OrderOrderService) GetListAll(info bizReq.OrderOrderSearch, createdAtBetween []string, fields string) (list []business.OrderOrder, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.OrderOrder{})
	var orderOrders []business.OrderOrder
	//var orderOrders []business.OrderOrderMini

	if info.ID > 0 {
		db = db.Where("`id` = ?", info.ID)
	}
	if len(createdAtBetween) >= 2 {
		db = db.Where("`created_at` BETWEEN ? AND ?", createdAtBetween[0], createdAtBetween[1])
	}

	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Sn != "" {
		db = db.Where("`sn` = ?", info.Sn)
	}
	if info.UserId != nil {
		db = db.Where("`user_id` = ?", info.UserId)
	}
	if info.OrderType != nil {
		db = db.Where("`order_type` = ?", info.OrderType)
	}
	if info.ObjId != nil {
		db = db.Where("`obj_id` = ?", info.ObjId)
	}
	if info.OrderCode != "" {
		db = db.Where("`order_code` = ?", info.OrderCode)
	}
	if info.StatusPay != nil {
		db = db.Where("`status_pay` = ?", info.StatusPay)
	}
	if info.StatusOrder != nil {
		db = db.Where("`status_order` = ?", info.StatusOrder)
	}
	if info.Status != nil {
		db = db.Where("`status` = ?", info.Status)
	}

	err = db.Count(&total).Error
	if err != nil {
		return orderOrders, 0, err
	}

	//err = db.Limit(limit).Offset(offset).Find(&orderOrders).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&orderOrders).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&orderOrders).Error
	}
	//如果有图片image类型，更新图片path
	for i, v := range orderOrders {
		v.MapData = make(map[string]string)
		orderOrders[i] = v
	}
	return orderOrders, total, err
}
