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
func GetImTximSev() *ImTximService {
	once_ImTxim.Do(func() {
		obj_ImTximService = new(ImTximService)
		//instanse.init()
	})
	return obj_ImTximService
}

// Create 创建ImTxim记录
// Author [88act](https://github.com/88act)
func (m *ImTximService) Create(data business.ImTxim) (id uint, err error) {
	err = global.DB.Create(&data).Error
	if err != nil {
		return 0, err
	}
	return data.ID, err
}

// Delete 删除ImTxim记录
// Author [88act](https://github.com/88act)
func (m *ImTximService) Delete(data business.ImTxim) (err error) {
	err = global.DB.Delete(&data).Error
	return err
}

// DeleteByIds 批量删除ImTxim记录
// Author [88act](https://github.com/88act)
func (m *ImTximService) DeleteByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]business.ImTxim{}, "id in ?", ids.Ids).Error
	return err
}

// Update  更新ImTxim记录
// Author [88act](https://github.com/88act)
func (m *ImTximService) Update(data business.ImTxim) (err error) {
	err = global.DB.Save(&data).Error
	return err
}

// UpdateByMap  更新ImTxim记录 by Map
// values := map[string]interface{}{
// 	"status":0,
// 	"from": hash,
// }
// Author [88act](https://github.com/88act)
func (m *ImTximService) UpdateByMap(data business.ImTxim, mapData map[string]interface{}) (err error) {
	err = global.DB.Model(&data).Updates(mapData).Error
	return err
}

// Get 根据id获取ImTxim记录
// Author [88act](https://github.com/88act)
func (m *ImTximService) Get(id uint, fields string) (obj business.ImTxim, err error) {

	if utils.IsEmpty(fields) {
		err = global.DB.Where("id = ?", id).First(&obj).Error
	} else {
		err = global.DB.Select(fields).Where("id = ?", id).First(&obj).Error
	}

	//如果有图片image类型，更新图片path
	obj.MapData = make(map[string]string)
	return obj, err
}

// GetList 分页获取ImTxim记录
// Author [88act](https://github.com/88act)
func (m *ImTximService) GetList(info bizReq.ImTximSearch, createdAtBetween []string, fields string) (list []business.ImTximMini, total int64, err error) {
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
	if info.AppId != "" {
		db = db.Where("`app_id` = ?", info.AppId)
	}
	if info.BeginTime != nil {
		db = db.Where("`begin_time` > > ?", info.BeginTime)
	}
	if info.NowTime != nil {
		db = db.Where("`now_time` > > ?", info.NowTime)
	}
	if info.Status != nil {
		db = db.Where("`status` = ?", info.Status)
	}
	if info.StatusRun != nil {
		db = db.Where("`status_run` = ?", info.StatusRun)
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

//GetListAll 分页获取ImTxim记录 (全部字段)
// Author [88act](https://github.com/88act)
func (m *ImTximService) GetListAll(info bizReq.ImTximSearch, createdAtBetween []string, fields string) (list []business.ImTxim, total int64, err error) {
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
	if info.AppId != "" {
		db = db.Where("`app_id` = ?", info.AppId)
	}
	if info.BeginTime != nil {
		db = db.Where("`begin_time` > > ?", info.BeginTime)
	}
	if info.NowTime != nil {
		db = db.Where("`now_time` > > ?", info.NowTime)
	}
	if info.Status != nil {
		db = db.Where("`status` = ?", info.Status)
	}
	if info.StatusRun != nil {
		db = db.Where("`status_run` = ?", info.StatusRun)
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
