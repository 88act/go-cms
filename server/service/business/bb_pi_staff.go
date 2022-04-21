package business

import (
	"go-cms/global"
	"go-cms/model/business"
	bizReq "go-cms/model/business/request"
	"go-cms/model/common/request"
	"go-cms/utils"
	"sync"
)

type BbPiStaffService struct {
}

var once_BbPiStaff sync.Once = sync.Once{}
var obj_BbPiStaffService *BbPiStaffService

//获取单例
func GetBbPiStaffSev() *BbPiStaffService {
	once_BbPiStaff.Do(func() {
		obj_BbPiStaffService = new(BbPiStaffService)
		//instanse.init()
	})
	return obj_BbPiStaffService
}

// Create 创建BbPiStaff记录
// Author [88act](https://github.com/88act)
func (m *BbPiStaffService) Create(data business.BbPiStaff) (id uint, err error) {
	err = global.DB.Create(&data).Error
	if err != nil {
		return 0, err
	}
	return data.ID, err
}

// Delete 删除BbPiStaff记录
// Author [88act](https://github.com/88act)
func (m *BbPiStaffService) Delete(data business.BbPiStaff) (err error) {
	err = global.DB.Delete(&data).Error
	return err
}

// DeleteByIds 批量删除BbPiStaff记录
// Author [88act](https://github.com/88act)
func (m *BbPiStaffService) DeleteByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]business.BbPiStaff{}, "id in ?", ids.Ids).Error
	return err
}

// Update  更新BbPiStaff记录
// Author [88act](https://github.com/88act)
func (m *BbPiStaffService) Update(data business.BbPiStaff) (err error) {
	err = global.DB.Save(&data).Error
	return err
}

// UpdateByMap  更新BbPiStaff记录 by Map
// values := map[string]interface{}{
// 	"status":0,
// 	"from": hash,
// }
// Author [88act](https://github.com/88act)
func (m *BbPiStaffService) UpdateByMap(data business.BbPiStaff, mapData map[string]interface{}) (err error) {
	err = global.DB.Model(&data).Updates(mapData).Error
	return err
}

// Get 根据id获取BbPiStaff记录
// Author [88act](https://github.com/88act)
func (m *BbPiStaffService) Get(id uint, fields string) (obj business.BbPiStaff, err error) {

	if utils.IsEmpty(fields) {
		err = global.DB.Where("id = ?", id).First(&obj).Error
	} else {
		err = global.DB.Select(fields).Where("id = ?", id).First(&obj).Error
	}

	//如果有图片image类型，更新图片path
	obj.MapData = make(map[string]string)
	return obj, err
}

// GetList 分页获取BbPiStaff记录
// Author [88act](https://github.com/88act)
func (m *BbPiStaffService) GetList(info bizReq.BbPiStaffSearch, createdAtBetween []string, fields string) (list []business.BbPiStaffMini, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.BbPiStaff{})
	//var bbPiStaffs []business.BbPiStaff
	var bbPiStaffs []business.BbPiStaffMini

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
		return bbPiStaffs, 0, err
	}
	//err = db.Limit(limit).Offset(offset).Find(&bbPiStaffs).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&bbPiStaffs).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&bbPiStaffs).Error
	}
	//如果有图片image类型，更新图片path
	for i, v := range bbPiStaffs {
		v.MapData = make(map[string]string)
		bbPiStaffs[i] = v
	}
	return bbPiStaffs, total, err
}

//GetListAll 分页获取BbPiStaff记录 (全部字段)
// Author [88act](https://github.com/88act)
func (m *BbPiStaffService) GetListAll(info bizReq.BbPiStaffSearch, createdAtBetween []string, fields string) (list []business.BbPiStaff, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.BbPiStaff{})
	var bbPiStaffs []business.BbPiStaff
	//var bbPiStaffs []business.BbPiStaffMini

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
		return bbPiStaffs, 0, err
	}

	//err = db.Limit(limit).Offset(offset).Find(&bbPiStaffs).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&bbPiStaffs).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&bbPiStaffs).Error
	}
	//如果有图片image类型，更新图片path
	for i, v := range bbPiStaffs {
		v.MapData = make(map[string]string)
		bbPiStaffs[i] = v
	}
	return bbPiStaffs, total, err
}
