package business

import (
	"go-cms/global"
	"go-cms/model/business"
	bizReq "go-cms/model/business/request"
	"go-cms/model/common/request"
	"go-cms/utils"
	"sync"
)

type BbPiInstitutionBusinessService struct {
}

var once_BbPiInstitutionBusiness sync.Once = sync.Once{}
var obj_BbPiInstitutionBusinessService *BbPiInstitutionBusinessService

//获取单例
func GetBbPiInstitutionBusinessSev() *BbPiInstitutionBusinessService {
	once_BbPiInstitutionBusiness.Do(func() {
		obj_BbPiInstitutionBusinessService = new(BbPiInstitutionBusinessService)
		//instanse.init()
	})
	return obj_BbPiInstitutionBusinessService
}

// Create 创建BbPiInstitutionBusiness记录
// Author [88act](https://github.com/88act)
func (m *BbPiInstitutionBusinessService) Create(data business.BbPiInstitutionBusiness) (id uint, err error) {
	err = global.DB.Create(&data).Error
	if err != nil {
		return 0, err
	}
	return data.ID, err
}

// Delete 删除BbPiInstitutionBusiness记录
// Author [88act](https://github.com/88act)
func (m *BbPiInstitutionBusinessService) Delete(data business.BbPiInstitutionBusiness) (err error) {
	err = global.DB.Delete(&data).Error
	return err
}

// DeleteByIds 批量删除BbPiInstitutionBusiness记录
// Author [88act](https://github.com/88act)
func (m *BbPiInstitutionBusinessService) DeleteByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]business.BbPiInstitutionBusiness{}, "id in ?", ids.Ids).Error
	return err
}

// Update  更新BbPiInstitutionBusiness记录
// Author [88act](https://github.com/88act)
func (m *BbPiInstitutionBusinessService) Update(data business.BbPiInstitutionBusiness) (err error) {
	err = global.DB.Save(&data).Error
	return err
}

// UpdateByMap  更新BbPiInstitutionBusiness记录 by Map
// values := map[string]interface{}{
// 	"status":0,
// 	"from": hash,
// }
// Author [88act](https://github.com/88act)
func (m *BbPiInstitutionBusinessService) UpdateByMap(data business.BbPiInstitutionBusiness, mapData map[string]interface{}) (err error) {
	err = global.DB.Model(&data).Updates(mapData).Error
	return err
}

// Get 根据id获取BbPiInstitutionBusiness记录
// Author [88act](https://github.com/88act)
func (m *BbPiInstitutionBusinessService) Get(id uint, fields string) (obj business.BbPiInstitutionBusiness, err error) {

	if utils.IsEmpty(fields) {
		err = global.DB.Where("id = ?", id).First(&obj).Error
	} else {
		err = global.DB.Select(fields).Where("id = ?", id).First(&obj).Error
	}

	//如果有图片image类型，更新图片path
	obj.MapData = make(map[string]string)
	return obj, err
}

// GetList 分页获取BbPiInstitutionBusiness记录
// Author [88act](https://github.com/88act)
func (m *BbPiInstitutionBusinessService) GetList(info bizReq.BbPiInstitutionBusinessSearch, createdAtBetween []string, fields string) (list []business.BbPiInstitutionBusinessMini, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.BbPiInstitutionBusiness{})
	//var bbPiInstitutionBusinesss []business.BbPiInstitutionBusiness
	var bbPiInstitutionBusinesss []business.BbPiInstitutionBusinessMini

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
		return bbPiInstitutionBusinesss, 0, err
	}
	//err = db.Limit(limit).Offset(offset).Find(&bbPiInstitutionBusinesss).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&bbPiInstitutionBusinesss).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&bbPiInstitutionBusinesss).Error
	}
	//如果有图片image类型，更新图片path
	for i, v := range bbPiInstitutionBusinesss {
		v.MapData = make(map[string]string)
		bbPiInstitutionBusinesss[i] = v
	}
	return bbPiInstitutionBusinesss, total, err
}

//GetListAll 分页获取BbPiInstitutionBusiness记录 (全部字段)
// Author [88act](https://github.com/88act)
func (m *BbPiInstitutionBusinessService) GetListAll(info bizReq.BbPiInstitutionBusinessSearch, createdAtBetween []string, fields string) (list []business.BbPiInstitutionBusiness, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.BbPiInstitutionBusiness{})
	var bbPiInstitutionBusinesss []business.BbPiInstitutionBusiness
	//var bbPiInstitutionBusinesss []business.BbPiInstitutionBusinessMini

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
		return bbPiInstitutionBusinesss, 0, err
	}

	//err = db.Limit(limit).Offset(offset).Find(&bbPiInstitutionBusinesss).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&bbPiInstitutionBusinesss).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&bbPiInstitutionBusinesss).Error
	}
	//如果有图片image类型，更新图片path
	for i, v := range bbPiInstitutionBusinesss {
		v.MapData = make(map[string]string)
		bbPiInstitutionBusinesss[i] = v
	}
	return bbPiInstitutionBusinesss, total, err
}
