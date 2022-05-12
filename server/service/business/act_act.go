package business

import (
	"go-cms/global"
	"go-cms/model/business"
	bizReq "go-cms/model/business/request"
	"go-cms/model/common/request"
	"go-cms/utils"
	"sync"
)

type ActActService struct {
}

var once_ActAct sync.Once = sync.Once{}
var obj_ActActService *ActActService

//获取单例
func GetActActSev() *ActActService {
	once_ActAct.Do(func() {
		obj_ActActService = new(ActActService)
		//instanse.init()
	})
	return obj_ActActService
}

// Create 创建ActAct记录
// Author [88act](https://github.com/88act)
func (m *ActActService) Create(data business.ActAct) (id uint, err error) {
	err = global.DB.Create(&data).Error
	if err != nil {
		return 0, err
	}
	return data.ID, err
}

// Delete 删除ActAct记录
// Author [88act](https://github.com/88act)
func (m *ActActService) Delete(data business.ActAct) (err error) {
	err = global.DB.Delete(&data).Error
	return err
}

// DeleteByIds 批量删除ActAct记录
// Author [88act](https://github.com/88act)
func (m *ActActService) DeleteByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]business.ActAct{}, "id in ?", ids.Ids).Error
	return err
}

// Update  更新ActAct记录
// Author [88act](https://github.com/88act)
func (m *ActActService) Update(data business.ActAct) (err error) {
	err = global.DB.Save(&data).Error
	return err
}

// UpdateByMap  更新ActAct记录 by Map
// values := map[string]interface{}{
// 	"status":0,
// 	"from": hash,
// }
// Author [88act](https://github.com/88act)
func (m *ActActService) UpdateByMap(data business.ActAct, mapData map[string]interface{}) (err error) {
	err = global.DB.Model(&data).Updates(mapData).Error
	return err
}

// Get 根据id获取ActAct记录
// Author [88act](https://github.com/88act)
func (m *ActActService) Get(id uint, fields string) (obj business.ActAct, err error) {

	if utils.IsEmpty(fields) {
		err = global.DB.Where("id = ?", id).First(&obj).Error
	} else {
		err = global.DB.Select(fields).Where("id = ?", id).First(&obj).Error
	}

	//如果有图片image类型，更新图片path
	obj.MapData = make(map[string]string)
	return obj, err
}

// GetList 分页获取ActAct记录
// Author [88act](https://github.com/88act)
func (m *ActActService) GetList(info bizReq.ActActSearch, createdAtBetween []string, fields string) (list []business.ActActMini, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.ActAct{})
	//var actActs []business.ActAct
	var actActs []business.ActActMini

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
		return actActs, 0, err
	}
	//err = db.Limit(limit).Offset(offset).Find(&actActs).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&actActs).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&actActs).Error
	}
	//如果有图片image类型，更新图片path
	for i, v := range actActs {
		v.MapData = make(map[string]string)
		actActs[i] = v
	}
	return actActs, total, err
}

//GetListAll 分页获取ActAct记录 (全部字段)
// Author [88act](https://github.com/88act)
func (m *ActActService) GetListAll(info bizReq.ActActSearch, createdAtBetween []string, fields string) (list []business.ActAct, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.ActAct{})
	var actActs []business.ActAct
	//var actActs []business.ActActMini

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
		return actActs, 0, err
	}

	//err = db.Limit(limit).Offset(offset).Find(&actActs).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&actActs).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&actActs).Error
	}
	//如果有图片image类型，更新图片path
	for i, v := range actActs {
		v.MapData = make(map[string]string)
		actActs[i] = v
	}
	return actActs, total, err
}
