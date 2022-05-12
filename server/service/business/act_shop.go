package business

import (
	"go-cms/global"
	"go-cms/model/business"
	bizReq "go-cms/model/business/request"
	"go-cms/model/common/request"
	"go-cms/utils"
	"sync"
)

type ActShopService struct {
}

var once_ActShop sync.Once = sync.Once{}
var obj_ActShopService *ActShopService

//获取单例
func GetActShopSev() *ActShopService {
	once_ActShop.Do(func() {
		obj_ActShopService = new(ActShopService)
		//instanse.init()
	})
	return obj_ActShopService
}

// Create 创建ActShop记录
// Author [88act](https://github.com/88act)
func (m *ActShopService) Create(data business.ActShop) (id uint, err error) {
	err = global.DB.Create(&data).Error
	if err != nil {
		return 0, err
	}
	return data.ID, err
}

// Delete 删除ActShop记录
// Author [88act](https://github.com/88act)
func (m *ActShopService) Delete(data business.ActShop) (err error) {
	err = global.DB.Delete(&data).Error
	return err
}

// DeleteByIds 批量删除ActShop记录
// Author [88act](https://github.com/88act)
func (m *ActShopService) DeleteByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]business.ActShop{}, "id in ?", ids.Ids).Error
	return err
}

// Update  更新ActShop记录
// Author [88act](https://github.com/88act)
func (m *ActShopService) Update(data business.ActShop) (err error) {
	err = global.DB.Save(&data).Error
	return err
}

// UpdateByMap  更新ActShop记录 by Map
// values := map[string]interface{}{
// 	"status":0,
// 	"from": hash,
// }
// Author [88act](https://github.com/88act)
func (m *ActShopService) UpdateByMap(data business.ActShop, mapData map[string]interface{}) (err error) {
	err = global.DB.Model(&data).Updates(mapData).Error
	return err
}

// Get 根据id获取ActShop记录
// Author [88act](https://github.com/88act)
func (m *ActShopService) Get(id uint, fields string) (obj business.ActShop, err error) {

	if utils.IsEmpty(fields) {
		err = global.DB.Where("id = ?", id).First(&obj).Error
	} else {
		err = global.DB.Select(fields).Where("id = ?", id).First(&obj).Error
	}

	//如果有图片image类型，更新图片path
	obj.MapData = make(map[string]string)
	return obj, err
}

// GetList 分页获取ActShop记录
// Author [88act](https://github.com/88act)
func (m *ActShopService) GetList(info bizReq.ActShopSearch, createdAtBetween []string, fields string) (list []business.ActShopMini, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.ActShop{})
	//var actShops []business.ActShop
	var actShops []business.ActShopMini

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
	if info.Name != "" {
		db = db.Where("`name` LIKE ?", "%"+info.Name+"%")
	}
	if info.Desc != "" {
		db = db.Where("`desc` LIKE ?", "%"+info.Desc+"%")
	}
	if info.Mobile != "" {
		db = db.Where("`mobile` = ?", info.Mobile)
	}
	if info.VipLev != nil {
		db = db.Where("`vip_lev` = ?", info.VipLev)
	}
	if info.Status != nil {
		db = db.Where("`status` = ?", info.Status)
	}

	err = db.Count(&total).Error
	if err != nil {
		return actShops, 0, err
	}
	//err = db.Limit(limit).Offset(offset).Find(&actShops).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&actShops).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&actShops).Error
	}
	//如果有图片image类型，更新图片path
	for i, v := range actShops {
		v.MapData = make(map[string]string)
		actShops[i] = v
	}
	return actShops, total, err
}

//GetListAll 分页获取ActShop记录 (全部字段)
// Author [88act](https://github.com/88act)
func (m *ActShopService) GetListAll(info bizReq.ActShopSearch, createdAtBetween []string, fields string) (list []business.ActShop, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.ActShop{})
	var actShops []business.ActShop
	//var actShops []business.ActShopMini

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
	if info.Name != "" {
		db = db.Where("`name` LIKE ?", "%"+info.Name+"%")
	}
	if info.Desc != "" {
		db = db.Where("`desc` LIKE ?", "%"+info.Desc+"%")
	}
	if info.Mobile != "" {
		db = db.Where("`mobile` = ?", info.Mobile)
	}
	if info.VipLev != nil {
		db = db.Where("`vip_lev` = ?", info.VipLev)
	}
	if info.Status != nil {
		db = db.Where("`status` = ?", info.Status)
	}

	err = db.Count(&total).Error
	if err != nil {
		return actShops, 0, err
	}

	//err = db.Limit(limit).Offset(offset).Find(&actShops).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&actShops).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&actShops).Error
	}
	//如果有图片image类型，更新图片path
	for i, v := range actShops {
		v.MapData = make(map[string]string)
		actShops[i] = v
	}
	return actShops, total, err
}
