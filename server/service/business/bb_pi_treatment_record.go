package business

import (
	"go-cms/global"
	"go-cms/model/business"
	bizReq "go-cms/model/business/request"
	"go-cms/model/common/request"
	"go-cms/utils"
	"sync"
)

type BbPiTreatmentRecordService struct {
}

var once_BbPiTreatmentRecord sync.Once = sync.Once{}
var obj_BbPiTreatmentRecordService *BbPiTreatmentRecordService

//获取单例
func GetBbPiTreatmentRecordSev() *BbPiTreatmentRecordService {
	once_BbPiTreatmentRecord.Do(func() {
		obj_BbPiTreatmentRecordService = new(BbPiTreatmentRecordService)
		//instanse.init()
	})
	return obj_BbPiTreatmentRecordService
}

// Create 创建BbPiTreatmentRecord记录
// Author [88act](https://github.com/88act)
func (m *BbPiTreatmentRecordService) Create(data business.BbPiTreatmentRecord) (id uint, err error) {
	err = global.DB.Create(&data).Error
	if err != nil {
		return 0, err
	}
	return data.ID, err
}

// Delete 删除BbPiTreatmentRecord记录
// Author [88act](https://github.com/88act)
func (m *BbPiTreatmentRecordService) Delete(data business.BbPiTreatmentRecord) (err error) {
	err = global.DB.Delete(&data).Error
	return err
}

// DeleteByIds 批量删除BbPiTreatmentRecord记录
// Author [88act](https://github.com/88act)
func (m *BbPiTreatmentRecordService) DeleteByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]business.BbPiTreatmentRecord{}, "id in ?", ids.Ids).Error
	return err
}

// Update  更新BbPiTreatmentRecord记录
// Author [88act](https://github.com/88act)
func (m *BbPiTreatmentRecordService) Update(data business.BbPiTreatmentRecord) (err error) {
	err = global.DB.Save(&data).Error
	return err
}

// UpdateByMap  更新BbPiTreatmentRecord记录 by Map
// values := map[string]interface{}{
// 	"status":0,
// 	"from": hash,
// }
// Author [88act](https://github.com/88act)
func (m *BbPiTreatmentRecordService) UpdateByMap(data business.BbPiTreatmentRecord, mapData map[string]interface{}) (err error) {
	err = global.DB.Model(&data).Updates(mapData).Error
	return err
}

// Get 根据id获取BbPiTreatmentRecord记录
// Author [88act](https://github.com/88act)
func (m *BbPiTreatmentRecordService) Get(id uint, fields string) (obj business.BbPiTreatmentRecord, err error) {

	if utils.IsEmpty(fields) {
		err = global.DB.Where("id = ?", id).First(&obj).Error
	} else {
		err = global.DB.Select(fields).Where("id = ?", id).First(&obj).Error
	}

	//如果有图片image类型，更新图片path
	obj.MapData = make(map[string]string)
	return obj, err
}

// GetList 分页获取BbPiTreatmentRecord记录
// Author [88act](https://github.com/88act)
func (m *BbPiTreatmentRecordService) GetList(info bizReq.BbPiTreatmentRecordSearch, createdAtBetween []string, fields string) (list []business.BbPiTreatmentRecordMini, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.BbPiTreatmentRecord{})
	//var bbPiTreatmentRecords []business.BbPiTreatmentRecord
	var bbPiTreatmentRecords []business.BbPiTreatmentRecordMini

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
		return bbPiTreatmentRecords, 0, err
	}
	//err = db.Limit(limit).Offset(offset).Find(&bbPiTreatmentRecords).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&bbPiTreatmentRecords).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&bbPiTreatmentRecords).Error
	}
	//如果有图片image类型，更新图片path
	for i, v := range bbPiTreatmentRecords {
		v.MapData = make(map[string]string)
		bbPiTreatmentRecords[i] = v
	}
	return bbPiTreatmentRecords, total, err
}

//GetListAll 分页获取BbPiTreatmentRecord记录 (全部字段)
// Author [88act](https://github.com/88act)
func (m *BbPiTreatmentRecordService) GetListAll(info bizReq.BbPiTreatmentRecordSearch, createdAtBetween []string, fields string) (list []business.BbPiTreatmentRecord, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.BbPiTreatmentRecord{})
	var bbPiTreatmentRecords []business.BbPiTreatmentRecord
	//var bbPiTreatmentRecords []business.BbPiTreatmentRecordMini

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
		return bbPiTreatmentRecords, 0, err
	}

	//err = db.Limit(limit).Offset(offset).Find(&bbPiTreatmentRecords).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&bbPiTreatmentRecords).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&bbPiTreatmentRecords).Error
	}
	//如果有图片image类型，更新图片path
	for i, v := range bbPiTreatmentRecords {
		v.MapData = make(map[string]string)
		bbPiTreatmentRecords[i] = v
	}
	return bbPiTreatmentRecords, total, err
}
