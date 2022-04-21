package business

import (
	"go-cms/global"
	"go-cms/model/business"
	bizReq "go-cms/model/business/request"
	"go-cms/model/common/request"
	"go-cms/utils"
	"sync"
)

type BbPiServicePointService struct {
}

var once_BbPiServicePoint sync.Once = sync.Once{}
var obj_BbPiServicePointService *BbPiServicePointService

//获取单例
func GetBbPiServicePointSev() *BbPiServicePointService {
	once_BbPiServicePoint.Do(func() {
		obj_BbPiServicePointService = new(BbPiServicePointService)
		//instanse.init()
	})
	return obj_BbPiServicePointService
}

// Create 创建BbPiServicePoint记录
// Author [88act](https://github.com/88act)
func (m *BbPiServicePointService) Create(data business.BbPiServicePoint) (id uint, err error) {
	err = global.DB.Create(&data).Error
	if err != nil {
		return 0, err
	}
	return data.ID, err
}

// Delete 删除BbPiServicePoint记录
// Author [88act](https://github.com/88act)
func (m *BbPiServicePointService) Delete(data business.BbPiServicePoint) (err error) {
	err = global.DB.Delete(&data).Error
	return err
}

// DeleteByIds 批量删除BbPiServicePoint记录
// Author [88act](https://github.com/88act)
func (m *BbPiServicePointService) DeleteByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]business.BbPiServicePoint{}, "id in ?", ids.Ids).Error
	return err
}

// Update  更新BbPiServicePoint记录
// Author [88act](https://github.com/88act)
func (m *BbPiServicePointService) Update(data business.BbPiServicePoint) (err error) {
	err = global.DB.Save(&data).Error
	return err
}

// UpdateByMap  更新BbPiServicePoint记录 by Map
// values := map[string]interface{}{
// 	"status":0,
// 	"from": hash,
// }
// Author [88act](https://github.com/88act)
func (m *BbPiServicePointService) UpdateByMap(data business.BbPiServicePoint, mapData map[string]interface{}) (err error) {
	err = global.DB.Model(&data).Updates(mapData).Error
	return err
}

// Get 根据id获取BbPiServicePoint记录
// Author [88act](https://github.com/88act)
func (m *BbPiServicePointService) Get(id uint, fields string) (obj business.BbPiServicePoint, err error) {

	if utils.IsEmpty(fields) {
		err = global.DB.Where("id = ?", id).First(&obj).Error
	} else {
		err = global.DB.Select(fields).Where("id = ?", id).First(&obj).Error
	}

	//如果有图片image类型，更新图片path
	obj.MapData = make(map[string]string)
	return obj, err
}

// GetList 分页获取BbPiServicePoint记录
// Author [88act](https://github.com/88act)
func (m *BbPiServicePointService) GetList(info bizReq.BbPiServicePointSearch, createdAtBetween []string, fields string) (list []business.BbPiServicePointMini, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.BbPiServicePoint{})
	//var bbPiServicePoints []business.BbPiServicePoint
	var bbPiServicePoints []business.BbPiServicePointMini

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
		return bbPiServicePoints, 0, err
	}
	//err = db.Limit(limit).Offset(offset).Find(&bbPiServicePoints).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&bbPiServicePoints).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&bbPiServicePoints).Error
	}
	//如果有图片image类型，更新图片path
	for i, v := range bbPiServicePoints {
		v.MapData = make(map[string]string)
		bbPiServicePoints[i] = v
	}
	return bbPiServicePoints, total, err
}

//GetListAll 分页获取BbPiServicePoint记录 (全部字段)
// Author [88act](https://github.com/88act)
func (m *BbPiServicePointService) GetListAll(info bizReq.BbPiServicePointSearch, createdAtBetween []string, fields string) (list []business.BbPiServicePoint, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.BbPiServicePoint{})
	var bbPiServicePoints []business.BbPiServicePoint
	//var bbPiServicePoints []business.BbPiServicePointMini

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
		return bbPiServicePoints, 0, err
	}

	//err = db.Limit(limit).Offset(offset).Find(&bbPiServicePoints).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&bbPiServicePoints).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&bbPiServicePoints).Error
	}
	//如果有图片image类型，更新图片path
	for i, v := range bbPiServicePoints {
		v.MapData = make(map[string]string)
		bbPiServicePoints[i] = v
	}
	return bbPiServicePoints, total, err
}
