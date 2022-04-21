package business

import (
	"go-cms/global"
	"go-cms/model/business"
	bizReq "go-cms/model/business/request"
	"go-cms/model/common/request"
	"go-cms/utils"
	"sync"
)

type BbPiInstitutionService struct {
}

var once_BbPiInstitution sync.Once = sync.Once{}
var obj_BbPiInstitutionService *BbPiInstitutionService

//获取单例
func GetBbPiInstitutionSev() *BbPiInstitutionService {
	once_BbPiInstitution.Do(func() {
		obj_BbPiInstitutionService = new(BbPiInstitutionService)
		//instanse.init()
	})
	return obj_BbPiInstitutionService
}

// Create 创建BbPiInstitution记录
// Author [88act](https://github.com/88act)
func (m *BbPiInstitutionService) Create(data business.BbPiInstitution) (id uint, err error) {
	err = global.DB.Create(&data).Error
	if err != nil {
		return 0, err
	}
	return data.ID, err
}

// Delete 删除BbPiInstitution记录
// Author [88act](https://github.com/88act)
func (m *BbPiInstitutionService) Delete(data business.BbPiInstitution) (err error) {
	err = global.DB.Delete(&data).Error
	return err
}

// DeleteByIds 批量删除BbPiInstitution记录
// Author [88act](https://github.com/88act)
func (m *BbPiInstitutionService) DeleteByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]business.BbPiInstitution{}, "id in ?", ids.Ids).Error
	return err
}

// Update  更新BbPiInstitution记录
// Author [88act](https://github.com/88act)
func (m *BbPiInstitutionService) Update(data business.BbPiInstitution) (err error) {
	err = global.DB.Save(&data).Error
	return err
}

// UpdateByMap  更新BbPiInstitution记录 by Map
// values := map[string]interface{}{
// 	"status":0,
// 	"from": hash,
// }
// Author [88act](https://github.com/88act)
func (m *BbPiInstitutionService) UpdateByMap(data business.BbPiInstitution, mapData map[string]interface{}) (err error) {
	err = global.DB.Model(&data).Updates(mapData).Error
	return err
}

// Get 根据id获取BbPiInstitution记录
// Author [88act](https://github.com/88act)
func (m *BbPiInstitutionService) Get(id uint, fields string) (obj business.BbPiInstitution, err error) {

	if utils.IsEmpty(fields) {
		err = global.DB.Where("id = ?", id).First(&obj).Error
	} else {
		err = global.DB.Select(fields).Where("id = ?", id).First(&obj).Error
	}

	//如果有图片image类型，更新图片path
	obj.MapData = make(map[string]string)
	return obj, err
}

// GetList 分页获取BbPiInstitution记录
// Author [88act](https://github.com/88act)
func (m *BbPiInstitutionService) GetList(info bizReq.BbPiInstitutionSearch, createdAtBetween []string, fields string) (list []business.BbPiInstitutionMini, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.BbPiInstitution{})
	//var bbPiInstitutions []business.BbPiInstitution
	var bbPiInstitutions []business.BbPiInstitutionMini

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
		return bbPiInstitutions, 0, err
	}
	//err = db.Limit(limit).Offset(offset).Find(&bbPiInstitutions).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&bbPiInstitutions).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&bbPiInstitutions).Error
	}
	//如果有图片image类型，更新图片path
	for i, v := range bbPiInstitutions {
		v.MapData = make(map[string]string)
		bbPiInstitutions[i] = v
	}
	return bbPiInstitutions, total, err
}

//GetListAll 分页获取BbPiInstitution记录 (全部字段)
// Author [88act](https://github.com/88act)
func (m *BbPiInstitutionService) GetListAll(info bizReq.BbPiInstitutionSearch, createdAtBetween []string, fields string) (list []business.BbPiInstitution, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.BbPiInstitution{})
	var bbPiInstitutions []business.BbPiInstitution
	//var bbPiInstitutions []business.BbPiInstitutionMini

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
		return bbPiInstitutions, 0, err
	}

	//err = db.Limit(limit).Offset(offset).Find(&bbPiInstitutions).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&bbPiInstitutions).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&bbPiInstitutions).Error
	}
	//如果有图片image类型，更新图片path
	for i, v := range bbPiInstitutions {
		v.MapData = make(map[string]string)
		bbPiInstitutions[i] = v
	}
	return bbPiInstitutions, total, err
}
