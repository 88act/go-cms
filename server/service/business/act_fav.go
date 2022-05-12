package business

import (
	"go-cms/global"
	"go-cms/model/business"
	bizReq "go-cms/model/business/request"
	"go-cms/model/common/request"
	"go-cms/utils"
	"sync"
)

type ActFavService struct {
}

var once_ActFav sync.Once = sync.Once{}
var obj_ActFavService *ActFavService

//获取单例
func GetActFavSev() *ActFavService {
	once_ActFav.Do(func() {
		obj_ActFavService = new(ActFavService)
		//instanse.init()
	})
	return obj_ActFavService
}

// Create 创建ActFav记录
// Author [88act](https://github.com/88act)
func (m *ActFavService) Create(data business.ActFav) (id uint, err error) {
	err = global.DB.Create(&data).Error
	if err != nil {
		return 0, err
	}
	return data.ID, err
}

// Delete 删除ActFav记录
// Author [88act](https://github.com/88act)
func (m *ActFavService) Delete(data business.ActFav) (err error) {
	err = global.DB.Delete(&data).Error
	return err
}

// DeleteByIds 批量删除ActFav记录
// Author [88act](https://github.com/88act)
func (m *ActFavService) DeleteByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]business.ActFav{}, "id in ?", ids.Ids).Error
	return err
}

// Update  更新ActFav记录
// Author [88act](https://github.com/88act)
func (m *ActFavService) Update(data business.ActFav) (err error) {
	err = global.DB.Save(&data).Error
	return err
}

// UpdateByMap  更新ActFav记录 by Map
// values := map[string]interface{}{
// 	"status":0,
// 	"from": hash,
// }
// Author [88act](https://github.com/88act)
func (m *ActFavService) UpdateByMap(data business.ActFav, mapData map[string]interface{}) (err error) {
	err = global.DB.Model(&data).Updates(mapData).Error
	return err
}

// Get 根据id获取ActFav记录
// Author [88act](https://github.com/88act)
func (m *ActFavService) Get(id uint, fields string) (obj business.ActFav, err error) {

	if utils.IsEmpty(fields) {
		err = global.DB.Where("id = ?", id).First(&obj).Error
	} else {
		err = global.DB.Select(fields).Where("id = ?", id).First(&obj).Error
	}

	//如果有图片image类型，更新图片path
	obj.MapData = make(map[string]string)
	return obj, err
}

// GetList 分页获取ActFav记录
// Author [88act](https://github.com/88act)
func (m *ActFavService) GetList(info bizReq.ActFavSearch, createdAtBetween []string, fields string) (list []business.ActFavMini, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.ActFav{})
	//var actFavs []business.ActFav
	var actFavs []business.ActFavMini

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
	if info.ObjId != nil {
		db = db.Where("`obj_id` = ?", info.ObjId)
	}
	if info.Status != nil {
		db = db.Where("`status` = ?", info.Status)
	}

	err = db.Count(&total).Error
	if err != nil {
		return actFavs, 0, err
	}
	//err = db.Limit(limit).Offset(offset).Find(&actFavs).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&actFavs).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&actFavs).Error
	}
	//如果有图片image类型，更新图片path
	for i, v := range actFavs {
		v.MapData = make(map[string]string)
		actFavs[i] = v
	}
	return actFavs, total, err
}

//GetListAll 分页获取ActFav记录 (全部字段)
// Author [88act](https://github.com/88act)
func (m *ActFavService) GetListAll(info bizReq.ActFavSearch, createdAtBetween []string, fields string) (list []business.ActFav, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.ActFav{})
	var actFavs []business.ActFav
	//var actFavs []business.ActFavMini

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
	if info.ObjId != nil {
		db = db.Where("`obj_id` = ?", info.ObjId)
	}
	if info.Status != nil {
		db = db.Where("`status` = ?", info.Status)
	}

	err = db.Count(&total).Error
	if err != nil {
		return actFavs, 0, err
	}

	//err = db.Limit(limit).Offset(offset).Find(&actFavs).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&actFavs).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&actFavs).Error
	}
	//如果有图片image类型，更新图片path
	for i, v := range actFavs {
		v.MapData = make(map[string]string)
		actFavs[i] = v
	}
	return actFavs, total, err
}
