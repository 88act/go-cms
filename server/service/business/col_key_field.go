package business

import (
	"go-cms/global"
	"go-cms/model/business"
	bizReq "go-cms/model/business/request"
	"go-cms/model/common/request"
	"go-cms/utils"
	"sync"
)

type ColKeyFieldService struct {
}

var once_ColKeyField sync.Once = sync.Once{}
var obj_ColKeyFieldService *ColKeyFieldService

//获取单例
func GetColKeyFieldService() *ColKeyFieldService {
	once_ColKeyField.Do(func() {
		obj_ColKeyFieldService = new(ColKeyFieldService)
		//instanse.init()
	})
	return obj_ColKeyFieldService
}

// CreateColKeyField 创建ColKeyField记录
// Author [88act](https://github.com/88act)
func (m *ColKeyFieldService) CreateColKeyField(colKeyField business.ColKeyField) (err error) {
	err = global.DB.Create(&colKeyField).Error
	return err
}

// DeleteColKeyField 删除ColKeyField记录
// Author [88act](https://github.com/88act)
func (m *ColKeyFieldService) DeleteColKeyField(colKeyField business.ColKeyField) (err error) {
	err = global.DB.Delete(&colKeyField).Error
	return err
}

// DeleteColKeyFieldByIds 批量删除ColKeyField记录
// Author [88act](https://github.com/88act)
func (m *ColKeyFieldService) DeleteColKeyFieldByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]business.ColKeyField{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateColKeyField 更新ColKeyField记录
// Author [88act](https://github.com/88act)
func (m *ColKeyFieldService) UpdateColKeyField(colKeyField business.ColKeyField) (err error) {
	err = global.DB.Save(&colKeyField).Error
	return err
}

// GetColKeyField 根据id获取ColKeyField记录
// Author [88act](https://github.com/88act)
func (m *ColKeyFieldService) GetColKeyField(id uint, fields string) (obj business.ColKeyField, err error) {

	if utils.IsEmpty(fields) {
		err = global.DB.Where("id = ?", id).First(&obj).Error
	} else {
		err = global.DB.Select(fields).Where("id = ?", id).First(&obj).Error
	}

	//如果有图片image类型，更新图片path
	obj.MapData = make(map[string]string)
	return obj, err
}

// GetColKeyFieldInfoList 分页获取ColKeyField记录
// Author [88act](https://github.com/88act)
func (m *ColKeyFieldService) GetColKeyFieldInfoList(info bizReq.ColKeyFieldSearch, createdAtBetween []string, fields string) (list []business.ColKeyFieldMini, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.ColKeyField{})
	//var colKeyFields []business.ColKeyField
	var colKeyFields []business.ColKeyFieldMini

	//修改 by ljd
	if info.ID > 0 {
		db = db.Where("`id` = ?", info.ID)
	}
	if len(createdAtBetween) >= 2 {
		db = db.Where("`created_at` BETWEEN ? AND ?", createdAtBetween[0], createdAtBetween[1])
	}

	// 如果有条件搜索 下方会自动创建搜索语句
	if info.ColId != nil {
		db = db.Where("`col_id` = ?", info.ColId)
	}
	if info.ToKey != "" {
		db = db.Where("`to_key` = ?", info.ToKey)
	}
	if info.ToField != "" {
		db = db.Where("`to_field` = ?", info.ToField)
	}
	if info.Status != nil {
		db = db.Where("`status` = ?", info.Status)
	}

	err = db.Count(&total).Error
	if err != nil {
		return colKeyFields, 0, err
	}
	//err = db.Limit(limit).Offset(offset).Find(&colKeyFields).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&colKeyFields).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&colKeyFields).Error
	}
	//如果有图片image类型，更新图片path
	for i, v := range colKeyFields {
		v.MapData = make(map[string]string)
		colKeyFields[i] = v
	}
	return colKeyFields, total, err
}

// GetColKeyFieldInfoListAll  分页获取ColKeyField记录 (全部字段)
// Author [88act](https://github.com/88act)
func (m *ColKeyFieldService) GetColKeyFieldInfoListAll(info bizReq.ColKeyFieldSearch, createdAtBetween []string, fields string) (list []business.ColKeyField, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.ColKeyField{})
	var colKeyFields []business.ColKeyField
	//var colKeyFields []business.ColKeyFieldMini

	if info.ID > 0 {
		db = db.Where("`id` = ?", info.ID)
	}
	if len(createdAtBetween) >= 2 {
		db = db.Where("`created_at` BETWEEN ? AND ?", createdAtBetween[0], createdAtBetween[1])
	}

	// 如果有条件搜索 下方会自动创建搜索语句
	if info.ColId != nil {
		db = db.Where("`col_id` = ?", info.ColId)
	}
	if info.ToKey != "" {
		db = db.Where("`to_key` = ?", info.ToKey)
	}
	if info.ToField != "" {
		db = db.Where("`to_field` = ?", info.ToField)
	}
	if info.Status != nil {
		db = db.Where("`status` = ?", info.Status)
	}

	err = db.Count(&total).Error
	if err != nil {
		return colKeyFields, 0, err
	}

	//err = db.Limit(limit).Offset(offset).Find(&colKeyFields).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&colKeyFields).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&colKeyFields).Error
	}
	//如果有图片image类型，更新图片path
	for i, v := range colKeyFields {
		v.MapData = make(map[string]string)
		colKeyFields[i] = v
	}
	return colKeyFields, total, err
}
