package business

import (
	"go-cms/global"
	"go-cms/model/business"
	bizReq "go-cms/model/business/request"
	"go-cms/model/common/request"
	"go-cms/utils"
	"sync"
)

type BasicAreaService struct {
}

var once_BasicArea sync.Once = sync.Once{}
var obj_BasicAreaService *BasicAreaService

//获取单例
func GetBasicAreaSev() *BasicAreaService {
	once_BasicArea.Do(func() {
		obj_BasicAreaService = new(BasicAreaService)
		//instanse.init()
	})
	return obj_BasicAreaService
}

// Create 创建BasicArea记录
// Author [88act](https://github.com/88act)
func (m *BasicAreaService) Create(data business.BasicArea) (id uint, err error) {
	err = global.DB.Create(&data).Error
	if err != nil {
		return 0, err
	}
	return data.ID, err
}

// Delete 删除BasicArea记录
// Author [88act](https://github.com/88act)
func (m *BasicAreaService) Delete(data business.BasicArea) (err error) {
	err = global.DB.Delete(&data).Error
	return err
}

// DeleteByIds 批量删除BasicArea记录
// Author [88act](https://github.com/88act)
func (m *BasicAreaService) DeleteByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]business.BasicArea{}, "id in ?", ids.Ids).Error
	return err
}

// Update  更新BasicArea记录
// Author [88act](https://github.com/88act)
func (m *BasicAreaService) Update(data business.BasicArea) (err error) {
	err = global.DB.Save(&data).Error
	return err
}

// UpdateByMap  更新BasicArea记录 by Map
// values := map[string]interface{}{
// 	"status":0,
// 	"from": hash,
// }
// Author [88act](https://github.com/88act)
func (m *BasicAreaService) UpdateByMap(data business.BasicArea, mapData map[string]interface{}) (err error) {
	err = global.DB.Model(&data).Updates(mapData).Error
	return err
}

// Get 根据id获取BasicArea记录
// Author [88act](https://github.com/88act)
func (m *BasicAreaService) Get(id uint, fields string) (obj business.BasicArea, err error) {

	if utils.IsEmpty(fields) {
		err = global.DB.Where("id = ?", id).First(&obj).Error
	} else {
		err = global.DB.Select(fields).Where("id = ?", id).First(&obj).Error
	}

	//如果有图片image类型，更新图片path
	obj.MapData = make(map[string]string)
	return obj, err
}

// GetList 分页获取BasicArea记录
// Author [88act](https://github.com/88act)
func (m *BasicAreaService) GetList(info bizReq.BasicAreaSearch, createdAtBetween []string, fields string) (list []business.BasicAreaMini, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.BasicArea{})
	//var basicAreas []business.BasicArea
	var basicAreas []business.BasicAreaMini

	//修改 by ljd
	if info.ID > 0 {
		db = db.Where("`id` = ?", info.ID)
	}
	if len(createdAtBetween) >= 2 {
		db = db.Where("`created_at` BETWEEN ? AND ?", createdAtBetween[0], createdAtBetween[1])
	}

	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Areaname != "" {
		db = db.Where("`areaname` = ?", info.Areaname)
	}
	if info.Parentid != nil {
		db = db.Where("`parentid` = ?", info.Parentid)
	}
	if info.Level != nil {
		db = db.Where("`level` = ?", info.Level)
	}
	if info.Sort != nil {
		db = db.Where("`sort` = ?", info.Sort)
	}

	err = db.Count(&total).Error
	if err != nil {
		return basicAreas, 0, err
	}
	//err = db.Limit(limit).Offset(offset).Find(&basicAreas).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&basicAreas).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&basicAreas).Error
	}
	//如果有图片image类型，更新图片path
	for i, v := range basicAreas {
		v.MapData = make(map[string]string)
		basicAreas[i] = v
	}
	return basicAreas, total, err
}

//GetListAll 分页获取BasicArea记录 (全部字段)
// Author [88act](https://github.com/88act)
func (m *BasicAreaService) GetListAll(info bizReq.BasicAreaSearch, createdAtBetween []string, fields string) (list []business.BasicArea, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.BasicArea{})
	var basicAreas []business.BasicArea
	//var basicAreas []business.BasicAreaMini

	if info.ID > 0 {
		db = db.Where("`id` = ?", info.ID)
	}
	if len(createdAtBetween) >= 2 {
		db = db.Where("`created_at` BETWEEN ? AND ?", createdAtBetween[0], createdAtBetween[1])
	}

	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Areaname != "" {
		db = db.Where("`areaname` = ?", info.Areaname)
	}
	if info.Parentid != nil {
		db = db.Where("`parentid` = ?", info.Parentid)
	}
	if info.Level != nil {
		db = db.Where("`level` = ?", info.Level)
	}
	if info.Sort != nil {
		db = db.Where("`sort` = ?", info.Sort)
	}

	err = db.Count(&total).Error
	if err != nil {
		return basicAreas, 0, err
	}

	//err = db.Limit(limit).Offset(offset).Find(&basicAreas).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&basicAreas).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&basicAreas).Error
	}
	//如果有图片image类型，更新图片path
	for i, v := range basicAreas {
		v.MapData = make(map[string]string)
		basicAreas[i] = v
	}
	return basicAreas, total, err
}
