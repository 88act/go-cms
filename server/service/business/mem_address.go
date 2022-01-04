package business

import (
	"go-cms/global"
	"go-cms/model/business"
	bizReq "go-cms/model/business/request"
	"go-cms/model/common/request"
	"go-cms/utils"
	"sync"
)

type MemAddressService struct {
}

var once_MemAddress sync.Once = sync.Once{}
var obj_MemAddressService *MemAddressService

//获取单例
func GetMemAddressSev() *MemAddressService {
	once_MemAddress.Do(func() {
		obj_MemAddressService = new(MemAddressService)
		//instanse.init()
	})
	return obj_MemAddressService
}

// Create 创建MemAddress记录
// Author [88act](https://github.com/88act)
func (m *MemAddressService) Create(data business.MemAddress) (id uint, err error) {
	err = global.DB.Create(&data).Error
	if err != nil {
		return 0, err
	}
	return data.ID, err
}

// Delete 删除MemAddress记录
// Author [88act](https://github.com/88act)
func (m *MemAddressService) Delete(data business.MemAddress) (err error) {
	err = global.DB.Delete(&data).Error
	return err
}

// DeleteByIds 批量删除MemAddress记录
// Author [88act](https://github.com/88act)
func (m *MemAddressService) DeleteByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]business.MemAddress{}, "id in ?", ids.Ids).Error
	return err
}

// Update  更新MemAddress记录
// Author [88act](https://github.com/88act)
func (m *MemAddressService) Update(data business.MemAddress) (err error) {
	err = global.DB.Save(&data).Error
	return err
}

// UpdateByMap  更新MemAddress记录 by Map
// values := map[string]interface{}{
// 	"status":0,
// 	"from": hash,
// }
// Author [88act](https://github.com/88act)
func (m *MemAddressService) UpdateByMap(data business.MemAddress, mapData map[string]interface{}) (err error) {
	err = global.DB.Model(&data).Updates(mapData).Error
	return err
}

// Get 根据id获取MemAddress记录
// Author [88act](https://github.com/88act)
func (m *MemAddressService) Get(id uint, fields string) (obj business.MemAddress, err error) {

	if utils.IsEmpty(fields) {
		err = global.DB.Where("id = ?", id).First(&obj).Error
	} else {
		err = global.DB.Select(fields).Where("id = ?", id).First(&obj).Error
	}

	//如果有图片image类型，更新图片path
	obj.MapData = make(map[string]string)
	return obj, err
}

// GetList 分页获取MemAddress记录
// Author [88act](https://github.com/88act)
func (m *MemAddressService) GetList(info bizReq.MemAddressSearch, createdAtBetween []string, fields string) (list []business.MemAddressMini, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.MemAddress{})
	//var memAddresss []business.MemAddress
	var memAddresss []business.MemAddressMini

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
	if info.Realname != "" {
		db = db.Where("`realname` LIKE ?", "%"+info.Realname+"%")
	}
	if info.Phone != "" {
		db = db.Where("`phone` LIKE ?", "%"+info.Phone+"%")
	}
	if info.Email != "" {
		db = db.Where("`email` LIKE ?", "%"+info.Email+"%")
	}
	if info.AreaCode != "" {
		db = db.Where("`area_code` = ?", info.AreaCode)
	}
	if info.Address != "" {
		db = db.Where("`address` = ?", info.Address)
	}
	if info.ZipCode != "" {
		db = db.Where("`zip_code` LIKE ?", "%"+info.ZipCode+"%")
	}
	if info.IsDefault != nil {
		db = db.Where("`is_default` = ?", info.IsDefault)
	}
	if info.Status != nil {
		db = db.Where("`status` = ?", info.Status)
	}

	err = db.Count(&total).Error
	if err != nil {
		return memAddresss, 0, err
	}
	//err = db.Limit(limit).Offset(offset).Find(&memAddresss).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&memAddresss).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&memAddresss).Error
	}
	//如果有图片image类型，更新图片path
	for i, v := range memAddresss {
		v.MapData = make(map[string]string)
		memAddresss[i] = v
	}
	return memAddresss, total, err
}

//GetListAll 分页获取MemAddress记录 (全部字段)
// Author [88act](https://github.com/88act)
func (m *MemAddressService) GetListAll(info bizReq.MemAddressSearch, createdAtBetween []string, fields string) (list []business.MemAddress, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.MemAddress{})
	var memAddresss []business.MemAddress
	//var memAddresss []business.MemAddressMini

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
	if info.Realname != "" {
		db = db.Where("`realname` LIKE ?", "%"+info.Realname+"%")
	}
	if info.Phone != "" {
		db = db.Where("`phone` LIKE ?", "%"+info.Phone+"%")
	}
	if info.Email != "" {
		db = db.Where("`email` LIKE ?", "%"+info.Email+"%")
	}
	if info.AreaCode != "" {
		db = db.Where("`area_code` = ?", info.AreaCode)
	}
	if info.Address != "" {
		db = db.Where("`address` = ?", info.Address)
	}
	if info.ZipCode != "" {
		db = db.Where("`zip_code` LIKE ?", "%"+info.ZipCode+"%")
	}
	if info.IsDefault != nil {
		db = db.Where("`is_default` = ?", info.IsDefault)
	}
	if info.Status != nil {
		db = db.Where("`status` = ?", info.Status)
	}

	err = db.Count(&total).Error
	if err != nil {
		return memAddresss, 0, err
	}

	//err = db.Limit(limit).Offset(offset).Find(&memAddresss).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&memAddresss).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&memAddresss).Error
	}
	//如果有图片image类型，更新图片path
	for i, v := range memAddresss {
		v.MapData = make(map[string]string)
		memAddresss[i] = v
	}
	return memAddresss, total, err
}
