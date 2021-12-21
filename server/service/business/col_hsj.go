package business

import (
	"go-cms/global"
	"go-cms/model/business"
	businessReq "go-cms/model/business/request"
	"go-cms/model/common/request"
	"go-cms/utils"
)

type ColHsjService struct {
}

// CreateColHsj 创建ColHsj记录
// Author [88act](https://github.com/88act)
func (colHsjService *ColHsjService) CreateColHsj(colHsj business.ColHsj) (err error) {
	err = global.DB.Create(&colHsj).Error
	return err
}

// DeleteColHsj 删除ColHsj记录
// Author [88act](https://github.com/88act)
func (colHsjService *ColHsjService) DeleteColHsj(colHsj business.ColHsj) (err error) {
	err = global.DB.Delete(&colHsj).Error
	return err
}

// DeleteColHsjByIds 批量删除ColHsj记录
// Author [88act](https://github.com/88act)
func (colHsjService *ColHsjService) DeleteColHsjByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]business.ColHsj{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateColHsj 更新ColHsj记录
// Author [88act](https://github.com/88act)
func (colHsjService *ColHsjService) UpdateColHsj(colHsj business.ColHsj) (err error) {
	err = global.DB.Save(&colHsj).Error
	return err
}

// GetColHsj 根据id获取ColHsj记录
// Author [88act](https://github.com/88act)
func (colHsjService *ColHsjService) GetColHsj(id uint, fields string) (err error, obj business.ColHsj) {
	err = global.DB.Where("id = ?", id).First(&obj).Error
	if utils.IsEmpty(fields) {
		err = global.DB.Where("id = ?", id).First(&obj).Error
	} else {
		err = global.DB.Select(fields).Where("id = ?", id).First(&obj).Error
	}

	obj.MapData = make(map[string]string)
	return err, obj
}

// GetColHsjInfoList 分页获取ColHsj记录
// Author [88act](https://github.com/88act)
func (colHsjService *ColHsjService) GetColHsjInfoList(info businessReq.ColHsjSearch, createdAtBetween []string, fields string) (err error, list []business.ColHsjMini, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.ColHsj{})
	//var colHsjs []business.ColHsj
	var colHsjs []business.ColHsjMini

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
	if info.Title != "" {
		db = db.Where("`title` LIKE ?", "%"+info.Title+"%")
	}
	if info.Content != "" {
		db = db.Where("`content` LIKE ?", "%"+info.Content+"%")
	}
	if info.PubUnit != "" {
		db = db.Where("`pub_unit` LIKE ?", "%"+info.PubUnit+"%")
	}
	if info.Status != nil {
		db = db.Where("`status` = ?", info.Status)
	}
	err = db.Count(&total).Error
	//err = db.Limit(limit).Offset(offset).Find(&colHsjs).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&colHsjs).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&colHsjs).Error
	}
	//更新图片path
	for i, v := range colHsjs {
		v.MapData = make(map[string]string)
		colHsjs[i] = v
	}
	return err, colHsjs, total
}

// GetColHsjInfoList 分页获取ColHsj记录
// Author [88act](https://github.com/88act)
func (colHsjService *ColHsjService) GetColHsjInfoListAll(info businessReq.ColHsjSearch, createdAtBetween []string, fields string) (err error, list []business.ColHsj, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.ColHsj{})
	var colHsjs []business.ColHsj
	//var colHsjs []business.ColHsjMini

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
	if info.Title != "" {
		db = db.Where("`title` LIKE ?", "%"+info.Title+"%")
	}
	if info.Content != "" {
		db = db.Where("`content` LIKE ?", "%"+info.Content+"%")
	}
	if info.PubUnit != "" {
		db = db.Where("`pub_unit` LIKE ?", "%"+info.PubUnit+"%")
	}
	if info.Status != nil {
		db = db.Where("`status` = ?", info.Status)
	}
	err = db.Count(&total).Error
	//err = db.Limit(limit).Offset(offset).Find(&colHsjs).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&colHsjs).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&colHsjs).Error
	}
	//更新图片path
	for i, v := range colHsjs {
		v.MapData = make(map[string]string)
		colHsjs[i] = v
	}
	return err, colHsjs, total
}
