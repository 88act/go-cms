package business

import (
	"go-cms/global"
	"go-cms/model/business"
	bizReq "go-cms/model/business/request"
	"go-cms/model/common/request"
	"go-cms/utils"
	"sync"
)

type ImTximService struct {
}

var once_ImTxim sync.Once = sync.Once{}
var obj_ImTximService *ImTximService

//获取单例
func GetImTximService() *ImTximService {
	once_ImTxim.Do(func() {
		obj_ImTximService = new(ImTximService)
		//instanse.init()
	})
	return obj_ImTximService
}

// CreateImTxim 创建ImTxim记录
// Author [88act](https://github.com/88act)
func (m *ImTximService) CreateImTxim(imTxim business.ImTxim) (err error) {
	err = global.DB.Create(&imTxim).Error
	return err
}

// DeleteImTxim 删除ImTxim记录
// Author [88act](https://github.com/88act)
func (m *ImTximService) DeleteImTxim(imTxim business.ImTxim) (err error) {
	err = global.DB.Delete(&imTxim).Error
	return err
}

// DeleteImTximByIds 批量删除ImTxim记录
// Author [88act](https://github.com/88act)
func (m *ImTximService) DeleteImTximByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]business.ImTxim{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateImTxim 更新ImTxim记录
// Author [88act](https://github.com/88act)
func (m *ImTximService) UpdateImTxim(imTxim business.ImTxim) (err error) {
	err = global.DB.Save(&imTxim).Error
	return err
}

// GetImTxim 根据id获取ImTxim记录
// Author [88act](https://github.com/88act)
func (m *ImTximService) GetImTxim(id uint, fields string) (obj business.ImTxim, err error) {

	if utils.IsEmpty(fields) {
		err = global.DB.Where("id = ?", id).First(&obj).Error
	} else {
		err = global.DB.Select(fields).Where("id = ?", id).First(&obj).Error
	}

	//如果有图片image类型，更新图片path
	obj.MapData = make(map[string]string)
	return obj, err
}

// GetImTximInfoList 分页获取ImTxim记录
// Author [88act](https://github.com/88act)
func (m *ImTximService) GetImTximInfoList(info bizReq.ImTximSearch, createdAtBetween []string, fields string) (list []business.ImTximMini, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.ImTxim{})
	//var imTxims []business.ImTxim
	var imTxims []business.ImTximMini

	//修改 by ljd
	if info.ID > 0 {
		db = db.Where("`id` = ?", info.ID)
	}
	if len(createdAtBetween) >= 2 {
		db = db.Where("`created_at` BETWEEN ? AND ?", createdAtBetween[0], createdAtBetween[1])
	}

	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Name != "" {
		db = db.Where("`name` = ?", info.Name)
	}
	if info.AppId != "" {
		db = db.Where("`app_id` = ?", info.AppId)
	}
	if info.Identifier != "" {
		db = db.Where("`identifier` = ?", info.Identifier)
	}
	if info.RunTimes != nil {
		db = db.Where("`run_times` = ?", info.RunTimes)
	}
	if info.Status != nil {
		db = db.Where("`status` = ?", info.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return imTxims, 0, err
	}
	//err = db.Limit(limit).Offset(offset).Find(&imTxims).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&imTxims).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&imTxims).Error
	}
	//如果有图片image类型，更新图片path
	for i, v := range imTxims {
		v.MapData = make(map[string]string)
		imTxims[i] = v
	}
	return imTxims, total, err
}

// GetImTximInfoListAll  分页获取ImTxim记录 (全部字段)
// Author [88act](https://github.com/88act)
func (m *ImTximService) GetImTximInfoListAll(info bizReq.ImTximSearch, createdAtBetween []string, fields string) (list []business.ImTxim, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.ImTxim{})
	var imTxims []business.ImTxim
	//var imTxims []business.ImTximMini

	if info.ID > 0 {
		db = db.Where("`id` = ?", info.ID)
	}
	if len(createdAtBetween) >= 2 {
		db = db.Where("`created_at` BETWEEN ? AND ?", createdAtBetween[0], createdAtBetween[1])
	}

	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Name != "" {
		db = db.Where("`name` = ?", info.Name)
	}
	if info.AppId != "" {
		db = db.Where("`app_id` = ?", info.AppId)
	}
	if info.Identifier != "" {
		db = db.Where("`identifier` = ?", info.Identifier)
	}
	if info.RunTimes != nil {
		db = db.Where("`run_times` = ?", info.RunTimes)
	}
	if info.Status != nil {
		db = db.Where("`status` = ?", info.Status)
	}

	err = db.Count(&total).Error
	if err != nil {
		return imTxims, 0, err
	}

	//err = db.Limit(limit).Offset(offset).Find(&imTxims).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&imTxims).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&imTxims).Error
	}
	//如果有图片image类型，更新图片path
	for i, v := range imTxims {
		v.MapData = make(map[string]string)
		imTxims[i] = v
	}
	return imTxims, total, err
}
