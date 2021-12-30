package business

import (
	"go-cms/global"
	"go-cms/model/business"
	bizReq "go-cms/model/business/request"
	"go-cms/model/common/request"
	"go-cms/utils"
	"sync"
)

type CmsAdSeatService struct {
}

var once_CmsAdSeat sync.Once = sync.Once{}
var obj_CmsAdSeatService *CmsAdSeatService

//获取单例
func GetCmsAdSeatSev() *CmsAdSeatService {
	once_CmsAdSeat.Do(func() {
		obj_CmsAdSeatService = new(CmsAdSeatService)
		//instanse.init()
	})
	return obj_CmsAdSeatService
}

// Create 创建CmsAdSeat记录
// Author [88act](https://github.com/88act)
func (m *CmsAdSeatService) Create(data business.CmsAdSeat) (id uint, err error) {
	err = global.DB.Create(&data).Error
	if err != nil {
		return 0, err
	}
	return data.ID, err
}

// Delete 删除CmsAdSeat记录
// Author [88act](https://github.com/88act)
func (m *CmsAdSeatService) Delete(data business.CmsAdSeat) (err error) {
	err = global.DB.Delete(&data).Error
	return err
}

// DeleteByIds 批量删除CmsAdSeat记录
// Author [88act](https://github.com/88act)
func (m *CmsAdSeatService) DeleteByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]business.CmsAdSeat{}, "id in ?", ids.Ids).Error
	return err
}

// Update  更新CmsAdSeat记录
// Author [88act](https://github.com/88act)
func (m *CmsAdSeatService) Update(data business.CmsAdSeat) (err error) {
	err = global.DB.Save(&data).Error
	return err
}

// UpdateByMap  更新CmsAdSeat记录 by Map
// values := map[string]interface{}{
// 	"status":0,
// 	"from": hash,
// }
// Author [88act](https://github.com/88act)
func (m *CmsAdSeatService) UpdateByMap(data business.CmsAdSeat, mapData map[string]interface{}) (err error) {
	err = global.DB.Model(&data).Updates(mapData).Error
	return err
}

// Get 根据id获取CmsAdSeat记录
// Author [88act](https://github.com/88act)
func (m *CmsAdSeatService) Get(id uint, fields string) (obj business.CmsAdSeat, err error) {

	if utils.IsEmpty(fields) {
		err = global.DB.Where("id = ?", id).First(&obj).Error
	} else {
		err = global.DB.Select(fields).Where("id = ?", id).First(&obj).Error
	}

	//如果有图片image类型，更新图片path
	obj.MapData = make(map[string]string)
	return obj, err
}

// GetList 分页获取CmsAdSeat记录
// Author [88act](https://github.com/88act)
func (m *CmsAdSeatService) GetList(info bizReq.CmsAdSeatSearch, createdAtBetween []string, fields string) (list []business.CmsAdSeatMini, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.CmsAdSeat{})
	//var cmsAdSeats []business.CmsAdSeat
	var cmsAdSeats []business.CmsAdSeatMini

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
		return cmsAdSeats, 0, err
	}
	//err = db.Limit(limit).Offset(offset).Find(&cmsAdSeats).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&cmsAdSeats).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&cmsAdSeats).Error
	}
	//如果有图片image类型，更新图片path
	for i, v := range cmsAdSeats {
		v.MapData = make(map[string]string)
		cmsAdSeats[i] = v
	}
	return cmsAdSeats, total, err
}

//GetListAll 分页获取CmsAdSeat记录 (全部字段)
// Author [88act](https://github.com/88act)
func (m *CmsAdSeatService) GetListAll(info bizReq.CmsAdSeatSearch, createdAtBetween []string, fields string) (list []business.CmsAdSeat, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.CmsAdSeat{})
	var cmsAdSeats []business.CmsAdSeat
	//var cmsAdSeats []business.CmsAdSeatMini

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
		return cmsAdSeats, 0, err
	}

	//err = db.Limit(limit).Offset(offset).Find(&cmsAdSeats).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&cmsAdSeats).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&cmsAdSeats).Error
	}
	//如果有图片image类型，更新图片path
	for i, v := range cmsAdSeats {
		v.MapData = make(map[string]string)
		cmsAdSeats[i] = v
	}
	return cmsAdSeats, total, err
}
