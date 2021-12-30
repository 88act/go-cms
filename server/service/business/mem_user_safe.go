package business

import (
	"go-cms/global"
	"go-cms/model/business"
	bizReq "go-cms/model/business/request"
	"go-cms/model/common/request"
	"go-cms/utils"
	"sync"
)

type MemUserSafeService struct {
}

var once_MemUserSafe sync.Once = sync.Once{}
var obj_MemUserSafeService *MemUserSafeService

//获取单例
func GetMemUserSafeSev() *MemUserSafeService {
	once_MemUserSafe.Do(func() {
		obj_MemUserSafeService = new(MemUserSafeService)
		//instanse.init()
	})
	return obj_MemUserSafeService
}

// Create 创建MemUserSafe记录
// Author [88act](https://github.com/88act)
func (m *MemUserSafeService) Create(data business.MemUserSafe) (id uint, err error) {
	err = global.DB.Create(&data).Error
	if err != nil {
		return 0, err
	}
	return data.ID, err
}

// Delete 删除MemUserSafe记录
// Author [88act](https://github.com/88act)
func (m *MemUserSafeService) Delete(data business.MemUserSafe) (err error) {
	err = global.DB.Delete(&data).Error
	return err
}

// DeleteByIds 批量删除MemUserSafe记录
// Author [88act](https://github.com/88act)
func (m *MemUserSafeService) DeleteByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]business.MemUserSafe{}, "id in ?", ids.Ids).Error
	return err
}

// Update  更新MemUserSafe记录
// Author [88act](https://github.com/88act)
func (m *MemUserSafeService) Update(data business.MemUserSafe) (err error) {
	err = global.DB.Save(&data).Error
	return err
}

// UpdateByMap  更新MemUserSafe记录 by Map
// values := map[string]interface{}{
// 	"status":0,
// 	"from": hash,
// }
// Author [88act](https://github.com/88act)
func (m *MemUserSafeService) UpdateByMap(data business.MemUserSafe, mapData map[string]interface{}) (err error) {
	err = global.DB.Model(&data).Updates(mapData).Error
	return err
}

// Get 根据id获取MemUserSafe记录
// Author [88act](https://github.com/88act)
func (m *MemUserSafeService) Get(id uint, fields string) (obj business.MemUserSafe, err error) {

	if utils.IsEmpty(fields) {
		err = global.DB.Where("id = ?", id).First(&obj).Error
	} else {
		err = global.DB.Select(fields).Where("id = ?", id).First(&obj).Error
	}

	//如果有图片image类型，更新图片path
	obj.MapData = make(map[string]string)
	return obj, err
}

// GetList 分页获取MemUserSafe记录
// Author [88act](https://github.com/88act)
func (m *MemUserSafeService) GetList(info bizReq.MemUserSafeSearch, createdAtBetween []string, fields string) (list []business.MemUserSafeMini, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.MemUserSafe{})
	//var memUserSafes []business.MemUserSafe
	var memUserSafes []business.MemUserSafeMini

	//修改 by ljd
	if info.ID > 0 {
		db = db.Where("`id` = ?", info.ID)
	}
	if len(createdAtBetween) >= 2 {
		db = db.Where("`created_at` BETWEEN ? AND ?", createdAtBetween[0], createdAtBetween[1])
	}

	// 如果有条件搜索 下方会自动创建搜索语句
	if info.UserId != nil {
		db = db.Where("`user_id` = ?", info.UserId)
	}
	if info.Type != nil {
		db = db.Where("`type` = ?", info.Type)
	}
	if info.Status != nil {
		db = db.Where("`status` = ?", info.Status)
	}

	err = db.Count(&total).Error
	if err != nil {
		return memUserSafes, 0, err
	}
	//err = db.Limit(limit).Offset(offset).Find(&memUserSafes).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&memUserSafes).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&memUserSafes).Error
	}
	//如果有图片image类型，更新图片path
	for i, v := range memUserSafes {
		v.MapData = make(map[string]string)
		memUserSafes[i] = v
	}
	return memUserSafes, total, err
}

//GetListAll 分页获取MemUserSafe记录 (全部字段)
// Author [88act](https://github.com/88act)
func (m *MemUserSafeService) GetListAll(info bizReq.MemUserSafeSearch, createdAtBetween []string, fields string) (list []business.MemUserSafe, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.MemUserSafe{})
	var memUserSafes []business.MemUserSafe
	//var memUserSafes []business.MemUserSafeMini

	if info.ID > 0 {
		db = db.Where("`id` = ?", info.ID)
	}
	if len(createdAtBetween) >= 2 {
		db = db.Where("`created_at` BETWEEN ? AND ?", createdAtBetween[0], createdAtBetween[1])
	}

	// 如果有条件搜索 下方会自动创建搜索语句
	if info.UserId != nil {
		db = db.Where("`user_id` = ?", info.UserId)
	}
	if info.Type != nil {
		db = db.Where("`type` = ?", info.Type)
	}
	if info.Status != nil {
		db = db.Where("`status` = ?", info.Status)
	}

	err = db.Count(&total).Error
	if err != nil {
		return memUserSafes, 0, err
	}

	//err = db.Limit(limit).Offset(offset).Find(&memUserSafes).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&memUserSafes).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&memUserSafes).Error
	}
	//如果有图片image类型，更新图片path
	for i, v := range memUserSafes {
		v.MapData = make(map[string]string)
		memUserSafes[i] = v
	}
	return memUserSafes, total, err
}
