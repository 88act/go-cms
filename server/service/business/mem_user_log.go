package business

import (
	"go-cms/global"
	"go-cms/model/business"
	bizReq "go-cms/model/business/request"
	"go-cms/model/common/request"
	"go-cms/utils"
	"sync"
)

type MemUserLogService struct {
}

var once_MemUserLog sync.Once = sync.Once{}
var obj_MemUserLogService *MemUserLogService

//获取单例
func GetMemUserLogSev() *MemUserLogService {
	once_MemUserLog.Do(func() {
		obj_MemUserLogService = new(MemUserLogService)
		//instanse.init()
	})
	return obj_MemUserLogService
}

// Create 创建MemUserLog记录
// Author [88act](https://github.com/88act)
func (m *MemUserLogService) Create(data business.MemUserLog) (id uint, err error) {
	err = global.DB.Create(&data).Error
	if err != nil {
		return 0, err
	}
	return data.ID, err
}

// Delete 删除MemUserLog记录
// Author [88act](https://github.com/88act)
func (m *MemUserLogService) Delete(data business.MemUserLog) (err error) {
	err = global.DB.Delete(&data).Error
	return err
}

// DeleteByIds 批量删除MemUserLog记录
// Author [88act](https://github.com/88act)
func (m *MemUserLogService) DeleteByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]business.MemUserLog{}, "id in ?", ids.Ids).Error
	return err
}

// Update  更新MemUserLog记录
// Author [88act](https://github.com/88act)
func (m *MemUserLogService) Update(data business.MemUserLog) (err error) {
	err = global.DB.Save(&data).Error
	return err
}

// UpdateByMap  更新MemUserLog记录 by Map
// values := map[string]interface{}{
// 	"status":0,
// 	"from": hash,
// }
// Author [88act](https://github.com/88act)
func (m *MemUserLogService) UpdateByMap(data business.MemUserLog, mapData map[string]interface{}) (err error) {
	err = global.DB.Model(&data).Updates(mapData).Error
	return err
}

// Get 根据id获取MemUserLog记录
// Author [88act](https://github.com/88act)
func (m *MemUserLogService) Get(id uint, fields string) (obj business.MemUserLog, err error) {

	if utils.IsEmpty(fields) {
		err = global.DB.Where("id = ?", id).First(&obj).Error
	} else {
		err = global.DB.Select(fields).Where("id = ?", id).First(&obj).Error
	}

	//如果有图片image类型，更新图片path
	obj.MapData = make(map[string]string)
	return obj, err
}

// GetList 分页获取MemUserLog记录
// Author [88act](https://github.com/88act)
func (m *MemUserLogService) GetList(info bizReq.MemUserLogSearch, createdAtBetween []string, fields string) (list []business.MemUserLogMini, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.MemUserLog{})
	//var memUserLogs []business.MemUserLog
	var memUserLogs []business.MemUserLogMini

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
	if info.Ip != "" {
		db = db.Where("`ip` LIKE ?", "%"+info.Ip+"%")
	}

	err = db.Count(&total).Error
	if err != nil {
		return memUserLogs, 0, err
	}
	//err = db.Limit(limit).Offset(offset).Find(&memUserLogs).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&memUserLogs).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&memUserLogs).Error
	}
	//如果有图片image类型，更新图片path
	for i, v := range memUserLogs {
		v.MapData = make(map[string]string)
		memUserLogs[i] = v
	}
	return memUserLogs, total, err
}

//GetListAll 分页获取MemUserLog记录 (全部字段)
// Author [88act](https://github.com/88act)
func (m *MemUserLogService) GetListAll(info bizReq.MemUserLogSearch, createdAtBetween []string, fields string) (list []business.MemUserLog, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.MemUserLog{})
	var memUserLogs []business.MemUserLog
	//var memUserLogs []business.MemUserLogMini

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
	if info.Ip != "" {
		db = db.Where("`ip` LIKE ?", "%"+info.Ip+"%")
	}

	err = db.Count(&total).Error
	if err != nil {
		return memUserLogs, 0, err
	}

	//err = db.Limit(limit).Offset(offset).Find(&memUserLogs).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&memUserLogs).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&memUserLogs).Error
	}
	//如果有图片image类型，更新图片path
	for i, v := range memUserLogs {
		v.MapData = make(map[string]string)
		memUserLogs[i] = v
	}
	return memUserLogs, total, err
}
