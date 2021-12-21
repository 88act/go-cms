package business

import (
	"go-cms/global"
	"go-cms/model/business"
	businessReq "go-cms/model/business/request"
	"go-cms/model/common/request"
	"go-cms/utils"
)

type ColCollectService struct {
}

// CreateColCollect 创建ColCollect记录
// Author [88act](https://github.com/88act)
func (colCollectService *ColCollectService) CreateColCollect(colCollect business.ColCollect) (err error) {
	err = global.DB.Create(&colCollect).Error
	return err
}

// DeleteColCollect 删除ColCollect记录
// Author [88act](https://github.com/88act)
func (colCollectService *ColCollectService) DeleteColCollect(colCollect business.ColCollect) (err error) {
	err = global.DB.Delete(&colCollect).Error
	return err
}

// DeleteColCollectByIds 批量删除ColCollect记录
// Author [88act](https://github.com/88act)
func (colCollectService *ColCollectService) DeleteColCollectByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]business.ColCollect{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateColCollect 更新ColCollect记录
// Author [88act](https://github.com/88act)
func (colCollectService *ColCollectService) UpdateColCollect(colCollect business.ColCollect) (err error) {
	err = global.DB.Save(&colCollect).Error
	return err
}

// GetColCollect 根据id获取ColCollect记录
// Author [88act](https://github.com/88act)
func (colCollectService *ColCollectService) GetColCollect(id uint, fields string) (err error, obj business.ColCollect) {
	err = global.DB.Where("id = ?", id).First(&obj).Error
	if utils.IsEmpty(fields) {
		err = global.DB.Where("id = ?", id).First(&obj).Error
	} else {
		err = global.DB.Select(fields).Where("id = ?", id).First(&obj).Error
	}

	obj.MapData = make(map[string]string)
	return err, obj
}

// GetColCollectInfoList 分页获取ColCollect记录
// Author [88act](https://github.com/88act)
func (colCollectService *ColCollectService) GetColCollectInfoList(info businessReq.ColCollectSearch, createdAtBetween []string, fields string) (err error, list []business.ColCollectMini, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.ColCollect{})
	//var colCollects []business.ColCollect
	var colCollects []business.ColCollectMini

	//修改 by ljd
	if info.ID > 0 {
		db = db.Where("`id` = ?", info.ID)
	}
	if createdAtBetween != nil && len(createdAtBetween) > 0 {
		db = db.Where("`created_at` BETWEEN ? AND ?", createdAtBetween[0], createdAtBetween[1])
	}

	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Name != "" {
		db = db.Where("`name` LIKE ?", "%"+info.Name+"%")
	}
	if info.Url != "" {
		db = db.Where("`url` LIKE ?", "%"+info.Url+"%")
	}
	if info.StatusRun != nil {
		db = db.Where("`status_run` = ?", info.StatusRun)
	}
	if info.Status != nil {
		db = db.Where("`status` = ?", info.Status)
	}
	err = db.Count(&total).Error
	//err = db.Limit(limit).Offset(offset).Find(&colCollects).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&colCollects).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&colCollects).Error
	}
	//更新图片path
	for i, v := range colCollects {
		v.MapData = make(map[string]string)
		colCollects[i] = v
	}
	return err, colCollects, total
}

// GetColCollectInfoList 分页获取ColCollect记录
// Author [88act](https://github.com/88act)
func (colCollectService *ColCollectService) GetColCollectInfoListAll(info businessReq.ColCollectSearch, createdAtBetween []string, fields string) (err error, list []business.ColCollect, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.ColCollect{})
	var colCollects []business.ColCollect
	//var colCollects []business.ColCollectMini

	//修改 by ljd
	if info.ID > 0 {
		db = db.Where("`id` = ?", info.ID)
	}
	if createdAtBetween != nil && len(createdAtBetween) > 0 {
		db = db.Where("`created_at` BETWEEN ? AND ?", createdAtBetween[0], createdAtBetween[1])
	}

	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Name != "" {
		db = db.Where("`name` LIKE ?", "%"+info.Name+"%")
	}
	if info.Url != "" {
		db = db.Where("`url` LIKE ?", "%"+info.Url+"%")
	}
	if info.StatusRun != nil {
		db = db.Where("`status_run` = ?", info.StatusRun)
	}
	if info.Status != nil {
		db = db.Where("`status` = ?", info.Status)
	}
	err = db.Count(&total).Error
	//err = db.Limit(limit).Offset(offset).Find(&colCollects).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&colCollects).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&colCollects).Error
	}
	//更新图片path
	for i, v := range colCollects {
		v.MapData = make(map[string]string)
		colCollects[i] = v
	}
	return err, colCollects, total
}
