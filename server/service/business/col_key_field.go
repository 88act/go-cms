package business

import (
	"go-cms/global"
	"go-cms/model/business"
	businessReq "go-cms/model/business/request"
	"go-cms/model/common/request"
	"go-cms/utils"
)

type ColKeyFieldService struct {
}

// CreateColKeyField 创建ColKeyField记录
// Author [88act](https://github.com/88act)
func (colKeyFieldService *ColKeyFieldService) CreateColKeyField(colKeyField business.ColKeyField) (err error) {
	err = global.DB.Create(&colKeyField).Error
	return err
}

// DeleteColKeyField 删除ColKeyField记录
// Author [88act](https://github.com/88act)
func (colKeyFieldService *ColKeyFieldService) DeleteColKeyField(colKeyField business.ColKeyField) (err error) {
	err = global.DB.Delete(&colKeyField).Error
	return err
}

// DeleteColKeyFieldByIds 批量删除ColKeyField记录
// Author [88act](https://github.com/88act)
func (colKeyFieldService *ColKeyFieldService) DeleteColKeyFieldByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]business.ColKeyField{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateColKeyField 更新ColKeyField记录
// Author [88act](https://github.com/88act)
func (colKeyFieldService *ColKeyFieldService) UpdateColKeyField(colKeyField business.ColKeyField) (err error) {
	err = global.DB.Save(&colKeyField).Error
	return err
}

// GetColKeyField 根据id获取ColKeyField记录
// Author [88act](https://github.com/88act)
func (colKeyFieldService *ColKeyFieldService) GetColKeyField(id uint, fields string) (err error, obj business.ColKeyField) {
	err = global.DB.Where("id = ?", id).First(&obj).Error
	if utils.IsEmpty(fields) {
		err = global.DB.Where("id = ?", id).First(&obj).Error
	} else {
		err = global.DB.Select(fields).Where("id = ?", id).First(&obj).Error
	}

	obj.MapData = make(map[string]string)
	return err, obj
}

// GetColKeyFieldInfoList 分页获取ColKeyField记录
// Author [88act](https://github.com/88act)
func (colKeyFieldService *ColKeyFieldService) GetColKeyFieldInfoList(info businessReq.ColKeyFieldSearch, createdAtBetween []string, fields string) (err error, list []business.ColKeyFieldMini, total int64) {
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
	if createdAtBetween != nil && len(createdAtBetween) > 0 {
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
	//更新图片path
	for i, v := range colKeyFields {
		v.MapData = make(map[string]string)
		colKeyFields[i] = v
	}
	return err, colKeyFields, total
}

// GetColKeyFieldInfoList 分页获取ColKeyField记录
// Author [88act](https://github.com/88act)
func (colKeyFieldService *ColKeyFieldService) GetColKeyFieldInfoListAll(info businessReq.ColKeyFieldSearch, createdAtBetween []string, fields string) (err error, list []business.ColKeyField, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.ColKeyField{})
	var colKeyFields []business.ColKeyField
	//var colKeyFields []business.ColKeyFieldMini

	//修改 by ljd
	if info.ID > 0 {
		db = db.Where("`id` = ?", info.ID)
	}
	if createdAtBetween != nil && len(createdAtBetween) > 0 {
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
	//更新图片path
	for i, v := range colKeyFields {
		v.MapData = make(map[string]string)
		colKeyFields[i] = v
	}
	return err, colKeyFields, total
}
