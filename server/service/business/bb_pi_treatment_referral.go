package business

import (
	"go-cms/global"
	"go-cms/model/business"
	bizReq "go-cms/model/business/request"
	"go-cms/model/common/request"
	"go-cms/utils"
	"sync"
)

type BbPiTreatmentReferralService struct {
}

var once_BbPiTreatmentReferral sync.Once = sync.Once{}
var obj_BbPiTreatmentReferralService *BbPiTreatmentReferralService

//获取单例
func GetBbPiTreatmentReferralSev() *BbPiTreatmentReferralService {
	once_BbPiTreatmentReferral.Do(func() {
		obj_BbPiTreatmentReferralService = new(BbPiTreatmentReferralService)
		//instanse.init()
	})
	return obj_BbPiTreatmentReferralService
}

// Create 创建BbPiTreatmentReferral记录
// Author [88act](https://github.com/88act)
func (m *BbPiTreatmentReferralService) Create(data business.BbPiTreatmentReferral) (id uint, err error) {
	err = global.DB.Create(&data).Error
	if err != nil {
		return 0, err
	}
	return data.ID, err
}

// Delete 删除BbPiTreatmentReferral记录
// Author [88act](https://github.com/88act)
func (m *BbPiTreatmentReferralService) Delete(data business.BbPiTreatmentReferral) (err error) {
	err = global.DB.Delete(&data).Error
	return err
}

// DeleteByIds 批量删除BbPiTreatmentReferral记录
// Author [88act](https://github.com/88act)
func (m *BbPiTreatmentReferralService) DeleteByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]business.BbPiTreatmentReferral{}, "id in ?", ids.Ids).Error
	return err
}

// Update  更新BbPiTreatmentReferral记录
// Author [88act](https://github.com/88act)
func (m *BbPiTreatmentReferralService) Update(data business.BbPiTreatmentReferral) (err error) {
	err = global.DB.Save(&data).Error
	return err
}

// UpdateByMap  更新BbPiTreatmentReferral记录 by Map
// values := map[string]interface{}{
// 	"status":0,
// 	"from": hash,
// }
// Author [88act](https://github.com/88act)
func (m *BbPiTreatmentReferralService) UpdateByMap(data business.BbPiTreatmentReferral, mapData map[string]interface{}) (err error) {
	err = global.DB.Model(&data).Updates(mapData).Error
	return err
}

// Get 根据id获取BbPiTreatmentReferral记录
// Author [88act](https://github.com/88act)
func (m *BbPiTreatmentReferralService) Get(id uint, fields string) (obj business.BbPiTreatmentReferral, err error) {

	if utils.IsEmpty(fields) {
		err = global.DB.Where("id = ?", id).First(&obj).Error
	} else {
		err = global.DB.Select(fields).Where("id = ?", id).First(&obj).Error
	}

	//如果有图片image类型，更新图片path
	obj.MapData = make(map[string]string)
	return obj, err
}

// GetList 分页获取BbPiTreatmentReferral记录
// Author [88act](https://github.com/88act)
func (m *BbPiTreatmentReferralService) GetList(info bizReq.BbPiTreatmentReferralSearch, createdAtBetween []string, fields string) (list []business.BbPiTreatmentReferralMini, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.BbPiTreatmentReferral{})
	//var bbPiTreatmentReferrals []business.BbPiTreatmentReferral
	var bbPiTreatmentReferrals []business.BbPiTreatmentReferralMini

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
		return bbPiTreatmentReferrals, 0, err
	}
	//err = db.Limit(limit).Offset(offset).Find(&bbPiTreatmentReferrals).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&bbPiTreatmentReferrals).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&bbPiTreatmentReferrals).Error
	}
	//如果有图片image类型，更新图片path
	for i, v := range bbPiTreatmentReferrals {
		v.MapData = make(map[string]string)
		bbPiTreatmentReferrals[i] = v
	}
	return bbPiTreatmentReferrals, total, err
}

//GetListAll 分页获取BbPiTreatmentReferral记录 (全部字段)
// Author [88act](https://github.com/88act)
func (m *BbPiTreatmentReferralService) GetListAll(info bizReq.BbPiTreatmentReferralSearch, createdAtBetween []string, fields string) (list []business.BbPiTreatmentReferral, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.BbPiTreatmentReferral{})
	var bbPiTreatmentReferrals []business.BbPiTreatmentReferral
	//var bbPiTreatmentReferrals []business.BbPiTreatmentReferralMini

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
		return bbPiTreatmentReferrals, 0, err
	}

	//err = db.Limit(limit).Offset(offset).Find(&bbPiTreatmentReferrals).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&bbPiTreatmentReferrals).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&bbPiTreatmentReferrals).Error
	}
	//如果有图片image类型，更新图片path
	for i, v := range bbPiTreatmentReferrals {
		v.MapData = make(map[string]string)
		bbPiTreatmentReferrals[i] = v
	}
	return bbPiTreatmentReferrals, total, err
}
