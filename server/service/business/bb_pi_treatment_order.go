package business

import (
	"go-cms/global"
	"go-cms/model/business"
	bizReq "go-cms/model/business/request"
	"go-cms/model/common/request"
	"go-cms/utils"
	"sync"
)

type BbPiTreatmentOrderService struct {
}

var once_BbPiTreatmentOrder sync.Once = sync.Once{}
var obj_BbPiTreatmentOrderService *BbPiTreatmentOrderService

//获取单例
func GetBbPiTreatmentOrderSev() *BbPiTreatmentOrderService {
	once_BbPiTreatmentOrder.Do(func() {
		obj_BbPiTreatmentOrderService = new(BbPiTreatmentOrderService)
		//instanse.init()
	})
	return obj_BbPiTreatmentOrderService
}

// Create 创建BbPiTreatmentOrder记录
// Author [88act](https://github.com/88act)
func (m *BbPiTreatmentOrderService) Create(data business.BbPiTreatmentOrder) (id uint, err error) {
	err = global.DB.Create(&data).Error
	if err != nil {
		return 0, err
	}
	return data.ID, err
}

// Delete 删除BbPiTreatmentOrder记录
// Author [88act](https://github.com/88act)
func (m *BbPiTreatmentOrderService) Delete(data business.BbPiTreatmentOrder) (err error) {
	err = global.DB.Delete(&data).Error
	return err
}

// DeleteByIds 批量删除BbPiTreatmentOrder记录
// Author [88act](https://github.com/88act)
func (m *BbPiTreatmentOrderService) DeleteByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]business.BbPiTreatmentOrder{}, "id in ?", ids.Ids).Error
	return err
}

// Update  更新BbPiTreatmentOrder记录
// Author [88act](https://github.com/88act)
func (m *BbPiTreatmentOrderService) Update(data business.BbPiTreatmentOrder) (err error) {
	err = global.DB.Save(&data).Error
	return err
}

// UpdateByMap  更新BbPiTreatmentOrder记录 by Map
// values := map[string]interface{}{
// 	"status":0,
// 	"from": hash,
// }
// Author [88act](https://github.com/88act)
func (m *BbPiTreatmentOrderService) UpdateByMap(data business.BbPiTreatmentOrder, mapData map[string]interface{}) (err error) {
	err = global.DB.Model(&data).Updates(mapData).Error
	return err
}

// Get 根据id获取BbPiTreatmentOrder记录
// Author [88act](https://github.com/88act)
func (m *BbPiTreatmentOrderService) Get(id uint, fields string) (obj business.BbPiTreatmentOrder, err error) {

	if utils.IsEmpty(fields) {
		err = global.DB.Where("id = ?", id).First(&obj).Error
	} else {
		err = global.DB.Select(fields).Where("id = ?", id).First(&obj).Error
	}

	//如果有图片image类型，更新图片path
	obj.MapData = make(map[string]string)
	return obj, err
}

// GetList 分页获取BbPiTreatmentOrder记录
// Author [88act](https://github.com/88act)
func (m *BbPiTreatmentOrderService) GetList(info bizReq.BbPiTreatmentOrderSearch, createdAtBetween []string, fields string) (list []business.BbPiTreatmentOrderMini, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.BbPiTreatmentOrder{})
	//var bbPiTreatmentOrders []business.BbPiTreatmentOrder
	var bbPiTreatmentOrders []business.BbPiTreatmentOrderMini

	//修改 by ljd
	if info.ID > 0 {
		db = db.Where("`id` = ?", info.ID)
	}
	if len(createdAtBetween) >= 2 {
		db = db.Where("`created_at` BETWEEN ? AND ?", createdAtBetween[0], createdAtBetween[1])
	}

	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Status != nil {
		db = db.Where("`status` = ?", info.Status)
	}

	err = db.Count(&total).Error
	if err != nil {
		return bbPiTreatmentOrders, 0, err
	}
	//err = db.Limit(limit).Offset(offset).Find(&bbPiTreatmentOrders).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&bbPiTreatmentOrders).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&bbPiTreatmentOrders).Error
	}
	//如果有图片image类型，更新图片path
	for i, v := range bbPiTreatmentOrders {
		v.MapData = make(map[string]string)
		bbPiTreatmentOrders[i] = v
	}
	return bbPiTreatmentOrders, total, err
}

//GetListAll 分页获取BbPiTreatmentOrder记录 (全部字段)
// Author [88act](https://github.com/88act)
func (m *BbPiTreatmentOrderService) GetListAll(info bizReq.BbPiTreatmentOrderSearch, createdAtBetween []string, fields string) (list []business.BbPiTreatmentOrder, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.BbPiTreatmentOrder{})
	var bbPiTreatmentOrders []business.BbPiTreatmentOrder
	//var bbPiTreatmentOrders []business.BbPiTreatmentOrderMini

	if info.ID > 0 {
		db = db.Where("`id` = ?", info.ID)
	}
	if len(createdAtBetween) >= 2 {
		db = db.Where("`created_at` BETWEEN ? AND ?", createdAtBetween[0], createdAtBetween[1])
	}

	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Status != nil {
		db = db.Where("`status` = ?", info.Status)
	}

	err = db.Count(&total).Error
	if err != nil {
		return bbPiTreatmentOrders, 0, err
	}

	//err = db.Limit(limit).Offset(offset).Find(&bbPiTreatmentOrders).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&bbPiTreatmentOrders).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&bbPiTreatmentOrders).Error
	}
	//如果有图片image类型，更新图片path
	for i, v := range bbPiTreatmentOrders {
		v.MapData = make(map[string]string)
		bbPiTreatmentOrders[i] = v
	}
	return bbPiTreatmentOrders, total, err
}
