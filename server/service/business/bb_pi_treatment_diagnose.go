package business

import (
	"go-cms/global"
	"go-cms/model/business"
	bizReq "go-cms/model/business/request"
	"go-cms/model/common/request"
	"go-cms/utils"
	"sync"
)

type BbPiTreatmentDiagnoseService struct {
}

var once_BbPiTreatmentDiagnose sync.Once = sync.Once{}
var obj_BbPiTreatmentDiagnoseService *BbPiTreatmentDiagnoseService

//获取单例
func GetBbPiTreatmentDiagnoseSev() *BbPiTreatmentDiagnoseService {
	once_BbPiTreatmentDiagnose.Do(func() {
		obj_BbPiTreatmentDiagnoseService = new(BbPiTreatmentDiagnoseService)
		//instanse.init()
	})
	return obj_BbPiTreatmentDiagnoseService
}

// Create 创建BbPiTreatmentDiagnose记录
// Author [88act](https://github.com/88act)
func (m *BbPiTreatmentDiagnoseService) Create(data business.BbPiTreatmentDiagnose) (id uint, err error) {
	err = global.DB.Create(&data).Error
	if err != nil {
		return 0, err
	}
	return data.ID, err
}

// Delete 删除BbPiTreatmentDiagnose记录
// Author [88act](https://github.com/88act)
func (m *BbPiTreatmentDiagnoseService) Delete(data business.BbPiTreatmentDiagnose) (err error) {
	err = global.DB.Delete(&data).Error
	return err
}

// DeleteByIds 批量删除BbPiTreatmentDiagnose记录
// Author [88act](https://github.com/88act)
func (m *BbPiTreatmentDiagnoseService) DeleteByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]business.BbPiTreatmentDiagnose{}, "id in ?", ids.Ids).Error
	return err
}

// Update  更新BbPiTreatmentDiagnose记录
// Author [88act](https://github.com/88act)
func (m *BbPiTreatmentDiagnoseService) Update(data business.BbPiTreatmentDiagnose) (err error) {
	err = global.DB.Save(&data).Error
	return err
}

// UpdateByMap  更新BbPiTreatmentDiagnose记录 by Map
// values := map[string]interface{}{
// 	"status":0,
// 	"from": hash,
// }
// Author [88act](https://github.com/88act)
func (m *BbPiTreatmentDiagnoseService) UpdateByMap(data business.BbPiTreatmentDiagnose, mapData map[string]interface{}) (err error) {
	err = global.DB.Model(&data).Updates(mapData).Error
	return err
}

// Get 根据id获取BbPiTreatmentDiagnose记录
// Author [88act](https://github.com/88act)
func (m *BbPiTreatmentDiagnoseService) Get(id uint, fields string) (obj business.BbPiTreatmentDiagnose, err error) {

	if utils.IsEmpty(fields) {
		err = global.DB.Where("id = ?", id).First(&obj).Error
	} else {
		err = global.DB.Select(fields).Where("id = ?", id).First(&obj).Error
	}

	//如果有图片image类型，更新图片path
	obj.MapData = make(map[string]string)
	return obj, err
}

// GetList 分页获取BbPiTreatmentDiagnose记录
// Author [88act](https://github.com/88act)
func (m *BbPiTreatmentDiagnoseService) GetList(info bizReq.BbPiTreatmentDiagnoseSearch, createdAtBetween []string, fields string) (list []business.BbPiTreatmentDiagnoseMini, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.BbPiTreatmentDiagnose{})
	//var bbPiTreatmentDiagnoses []business.BbPiTreatmentDiagnose
	var bbPiTreatmentDiagnoses []business.BbPiTreatmentDiagnoseMini

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
		return bbPiTreatmentDiagnoses, 0, err
	}
	//err = db.Limit(limit).Offset(offset).Find(&bbPiTreatmentDiagnoses).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&bbPiTreatmentDiagnoses).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&bbPiTreatmentDiagnoses).Error
	}
	//如果有图片image类型，更新图片path
	for i, v := range bbPiTreatmentDiagnoses {
		v.MapData = make(map[string]string)
		bbPiTreatmentDiagnoses[i] = v
	}
	return bbPiTreatmentDiagnoses, total, err
}

//GetListAll 分页获取BbPiTreatmentDiagnose记录 (全部字段)
// Author [88act](https://github.com/88act)
func (m *BbPiTreatmentDiagnoseService) GetListAll(info bizReq.BbPiTreatmentDiagnoseSearch, createdAtBetween []string, fields string) (list []business.BbPiTreatmentDiagnose, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.BbPiTreatmentDiagnose{})
	var bbPiTreatmentDiagnoses []business.BbPiTreatmentDiagnose
	//var bbPiTreatmentDiagnoses []business.BbPiTreatmentDiagnoseMini

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
		return bbPiTreatmentDiagnoses, 0, err
	}

	//err = db.Limit(limit).Offset(offset).Find(&bbPiTreatmentDiagnoses).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&bbPiTreatmentDiagnoses).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&bbPiTreatmentDiagnoses).Error
	}
	//如果有图片image类型，更新图片path
	for i, v := range bbPiTreatmentDiagnoses {
		v.MapData = make(map[string]string)
		bbPiTreatmentDiagnoses[i] = v
	}
	return bbPiTreatmentDiagnoses, total, err
}
