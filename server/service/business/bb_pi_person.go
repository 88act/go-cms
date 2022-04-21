package business

import (
	"go-cms/global"
	"go-cms/model/business"
	bizReq "go-cms/model/business/request"
	"go-cms/model/common/request"
	"go-cms/utils"
	"sync"
)

type BbPiPersonService struct {
}

var once_BbPiPerson sync.Once = sync.Once{}
var obj_BbPiPersonService *BbPiPersonService

//获取单例
func GetBbPiPersonSev() *BbPiPersonService {
	once_BbPiPerson.Do(func() {
		obj_BbPiPersonService = new(BbPiPersonService)
		//instanse.init()
	})
	return obj_BbPiPersonService
}

// Create 创建BbPiPerson记录
// Author [88act](https://github.com/88act)
func (m *BbPiPersonService) Create(data business.BbPiPerson) (id uint, err error) {
	err = global.DB.Create(&data).Error
	if err != nil {
		return 0, err
	}
	return data.ID, err
}

// Delete 删除BbPiPerson记录
// Author [88act](https://github.com/88act)
func (m *BbPiPersonService) Delete(data business.BbPiPerson) (err error) {
	err = global.DB.Delete(&data).Error
	return err
}

// DeleteByIds 批量删除BbPiPerson记录
// Author [88act](https://github.com/88act)
func (m *BbPiPersonService) DeleteByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]business.BbPiPerson{}, "id in ?", ids.Ids).Error
	return err
}

// Update  更新BbPiPerson记录
// Author [88act](https://github.com/88act)
func (m *BbPiPersonService) Update(data business.BbPiPerson) (err error) {
	err = global.DB.Save(&data).Error
	return err
}

// UpdateByMap  更新BbPiPerson记录 by Map
// values := map[string]interface{}{
// 	"status":0,
// 	"from": hash,
// }
// Author [88act](https://github.com/88act)
func (m *BbPiPersonService) UpdateByMap(data business.BbPiPerson, mapData map[string]interface{}) (err error) {
	err = global.DB.Model(&data).Updates(mapData).Error
	return err
}

// Get 根据id获取BbPiPerson记录
// Author [88act](https://github.com/88act)
func (m *BbPiPersonService) Get(id uint, fields string) (obj business.BbPiPerson, err error) {

	if utils.IsEmpty(fields) {
		err = global.DB.Where("id = ?", id).First(&obj).Error
	} else {
		err = global.DB.Select(fields).Where("id = ?", id).First(&obj).Error
	}

	//如果有图片image类型，更新图片path
	obj.MapData = make(map[string]string)
	return obj, err
}

// GetList 分页获取BbPiPerson记录
// Author [88act](https://github.com/88act)
func (m *BbPiPersonService) GetList(info bizReq.BbPiPersonSearch, createdAtBetween []string, fields string) (list []business.BbPiPersonMini, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.BbPiPerson{})
	//var bbPiPersons []business.BbPiPerson
	var bbPiPersons []business.BbPiPersonMini

	//修改 by ljd
	if info.ID > 0 {
		db = db.Where("`id` = ?", info.ID)
	}
	if len(createdAtBetween) >= 2 {
		db = db.Where("`created_at` BETWEEN ? AND ?", createdAtBetween[0], createdAtBetween[1])
	}

	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Userid != nil {
		db = db.Where("`userid` = ?", info.Userid)
	}
	if info.Status != nil {
		db = db.Where("`status` = ?", info.Status)
	}

	err = db.Count(&total).Error
	if err != nil {
		return bbPiPersons, 0, err
	}
	//err = db.Limit(limit).Offset(offset).Find(&bbPiPersons).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&bbPiPersons).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&bbPiPersons).Error
	}
	//如果有图片image类型，更新图片path
	for i, v := range bbPiPersons {
		v.MapData = make(map[string]string)
		bbPiPersons[i] = v
	}
	return bbPiPersons, total, err
}

//GetListAll 分页获取BbPiPerson记录 (全部字段)
// Author [88act](https://github.com/88act)
func (m *BbPiPersonService) GetListAll(info bizReq.BbPiPersonSearch, createdAtBetween []string, fields string) (list []business.BbPiPerson, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.BbPiPerson{})
	var bbPiPersons []business.BbPiPerson
	//var bbPiPersons []business.BbPiPersonMini

	if info.ID > 0 {
		db = db.Where("`id` = ?", info.ID)
	}
	if len(createdAtBetween) >= 2 {
		db = db.Where("`created_at` BETWEEN ? AND ?", createdAtBetween[0], createdAtBetween[1])
	}

	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Userid != nil {
		db = db.Where("`userid` = ?", info.Userid)
	}
	if info.Status != nil {
		db = db.Where("`status` = ?", info.Status)
	}

	err = db.Count(&total).Error
	if err != nil {
		return bbPiPersons, 0, err
	}

	//err = db.Limit(limit).Offset(offset).Find(&bbPiPersons).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&bbPiPersons).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&bbPiPersons).Error
	}
	//如果有图片image类型，更新图片path
	for i, v := range bbPiPersons {
		v.MapData = make(map[string]string)
		bbPiPersons[i] = v
	}
	return bbPiPersons, total, err
}
