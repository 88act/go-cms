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
func GetMemUserSafeService() *MemUserSafeService {
	once_MemUserSafe.Do(func() {
		obj_MemUserSafeService = new(MemUserSafeService)
		//instanse.init()
	})
	return obj_MemUserSafeService
}

// CreateMemUserSafe 创建MemUserSafe记录
// Author [88act](https://github.com/88act)
func (m *MemUserSafeService) CreateMemUserSafe(memUserSafe business.MemUserSafe) (err error) {
	err = global.DB.Create(&memUserSafe).Error
	return err
}

// DeleteMemUserSafe 删除MemUserSafe记录
// Author [88act](https://github.com/88act)
func (m *MemUserSafeService) DeleteMemUserSafe(memUserSafe business.MemUserSafe) (err error) {
	err = global.DB.Delete(&memUserSafe).Error
	return err
}

// DeleteMemUserSafeByIds 批量删除MemUserSafe记录
// Author [88act](https://github.com/88act)
func (m *MemUserSafeService) DeleteMemUserSafeByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]business.MemUserSafe{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateMemUserSafe 更新MemUserSafe记录
// Author [88act](https://github.com/88act)
func (m *MemUserSafeService) UpdateMemUserSafe(memUserSafe business.MemUserSafe) (err error) {
	err = global.DB.Save(&memUserSafe).Error
	return err
}

// GetMemUserSafe 根据id获取MemUserSafe记录
// Author [88act](https://github.com/88act)
func (m *MemUserSafeService) GetMemUserSafe(id uint, fields string) (obj business.MemUserSafe, err error) {

	if utils.IsEmpty(fields) {
		err = global.DB.Where("id = ?", id).First(&obj).Error
	} else {
		err = global.DB.Select(fields).Where("id = ?", id).First(&obj).Error
	}

	//如果有图片image类型，更新图片path
	obj.MapData = make(map[string]string)
	return obj, err
}

// GetMemUserSafeInfoList 分页获取MemUserSafe记录
// Author [88act](https://github.com/88act)
func (m *MemUserSafeService) GetMemUserSafeInfoList(info bizReq.MemUserSafeSearch, createdAtBetween []string, fields string) (list []business.MemUserSafeMini, total int64, err error) {
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

// GetMemUserSafeInfoListAll  分页获取MemUserSafe记录 (全部字段)
// Author [88act](https://github.com/88act)
func (m *MemUserSafeService) GetMemUserSafeInfoListAll(info bizReq.MemUserSafeSearch, createdAtBetween []string, fields string) (list []business.MemUserSafe, total int64, err error) {
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
