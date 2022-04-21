package business

import (
	"go-cms/global"
	"go-cms/model/business"
	bizReq "go-cms/model/business/request"
	"go-cms/model/common/request"
	"go-cms/utils"
	"sync"
)

type BbPiDepartmentService struct {
}

var once_BbPiDepartment sync.Once = sync.Once{}
var obj_BbPiDepartmentService *BbPiDepartmentService

//获取单例
func GetBbPiDepartmentSev() *BbPiDepartmentService {
	once_BbPiDepartment.Do(func() {
		obj_BbPiDepartmentService = new(BbPiDepartmentService)
		//instanse.init()
	})
	return obj_BbPiDepartmentService
}

// Create 创建BbPiDepartment记录
// Author [88act](https://github.com/88act)
func (m *BbPiDepartmentService) Create(data business.BbPiDepartment) (id uint, err error) {
	err = global.DB.Create(&data).Error
	if err != nil {
		return 0, err
	}
	return data.ID, err
}

// Delete 删除BbPiDepartment记录
// Author [88act](https://github.com/88act)
func (m *BbPiDepartmentService) Delete(data business.BbPiDepartment) (err error) {
	err = global.DB.Delete(&data).Error
	return err
}

// DeleteByIds 批量删除BbPiDepartment记录
// Author [88act](https://github.com/88act)
func (m *BbPiDepartmentService) DeleteByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]business.BbPiDepartment{}, "id in ?", ids.Ids).Error
	return err
}

// Update  更新BbPiDepartment记录
// Author [88act](https://github.com/88act)
func (m *BbPiDepartmentService) Update(data business.BbPiDepartment) (err error) {
	err = global.DB.Save(&data).Error
	return err
}

// UpdateByMap  更新BbPiDepartment记录 by Map
// values := map[string]interface{}{
// 	"status":0,
// 	"from": hash,
// }
// Author [88act](https://github.com/88act)
func (m *BbPiDepartmentService) UpdateByMap(data business.BbPiDepartment, mapData map[string]interface{}) (err error) {
	err = global.DB.Model(&data).Updates(mapData).Error
	return err
}

// Get 根据id获取BbPiDepartment记录
// Author [88act](https://github.com/88act)
func (m *BbPiDepartmentService) Get(id uint, fields string) (obj business.BbPiDepartment, err error) {

	if utils.IsEmpty(fields) {
		err = global.DB.Where("id = ?", id).First(&obj).Error
	} else {
		err = global.DB.Select(fields).Where("id = ?", id).First(&obj).Error
	}

	//如果有图片image类型，更新图片path
	obj.MapData = make(map[string]string)
	return obj, err
}

// GetList 分页获取BbPiDepartment记录
// Author [88act](https://github.com/88act)
func (m *BbPiDepartmentService) GetList(info bizReq.BbPiDepartmentSearch, createdAtBetween []string, fields string) (list []business.BbPiDepartmentMini, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.BbPiDepartment{})
	//var bbPiDepartments []business.BbPiDepartment
	var bbPiDepartments []business.BbPiDepartmentMini

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
		return bbPiDepartments, 0, err
	}
	//err = db.Limit(limit).Offset(offset).Find(&bbPiDepartments).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&bbPiDepartments).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&bbPiDepartments).Error
	}
	//如果有图片image类型，更新图片path
	for i, v := range bbPiDepartments {
		v.MapData = make(map[string]string)
		bbPiDepartments[i] = v
	}
	return bbPiDepartments, total, err
}

//GetListAll 分页获取BbPiDepartment记录 (全部字段)
// Author [88act](https://github.com/88act)
func (m *BbPiDepartmentService) GetListAll(info bizReq.BbPiDepartmentSearch, createdAtBetween []string, fields string) (list []business.BbPiDepartment, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.BbPiDepartment{})
	var bbPiDepartments []business.BbPiDepartment
	//var bbPiDepartments []business.BbPiDepartmentMini

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
		return bbPiDepartments, 0, err
	}

	//err = db.Limit(limit).Offset(offset).Find(&bbPiDepartments).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&bbPiDepartments).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&bbPiDepartments).Error
	}
	//如果有图片image类型，更新图片path
	for i, v := range bbPiDepartments {
		v.MapData = make(map[string]string)
		bbPiDepartments[i] = v
	}
	return bbPiDepartments, total, err
}
