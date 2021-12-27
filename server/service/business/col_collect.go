package business

import (
	"go-cms/global"
	"go-cms/model/business"
	bizReq "go-cms/model/business/request"
	"go-cms/model/common/request"
	"go-cms/utils"
	"sync"
)

type ColCollectService struct {
}

var once_ColCollect sync.Once = sync.Once{}
var obj_ColCollectService *ColCollectService

//获取单例
func GetColCollectService() *ColCollectService {
	once_ColCollect.Do(func() {
		obj_ColCollectService = new(ColCollectService)
		//instanse.init()
	})
	return obj_ColCollectService
}

// CreateColCollect 创建ColCollect记录
// Author [88act](https://github.com/88act)
func (m *ColCollectService) CreateColCollect(colCollect business.ColCollect) (err error) {
	err = global.DB.Create(&colCollect).Error
	return err
}

// DeleteColCollect 删除ColCollect记录
// Author [88act](https://github.com/88act)
func (m *ColCollectService) DeleteColCollect(colCollect business.ColCollect) (err error) {
	err = global.DB.Delete(&colCollect).Error
	return err
}

// DeleteColCollectByIds 批量删除ColCollect记录
// Author [88act](https://github.com/88act)
func (m *ColCollectService) DeleteColCollectByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]business.ColCollect{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateColCollect 更新ColCollect记录
// Author [88act](https://github.com/88act)
func (m *ColCollectService) UpdateColCollect(colCollect business.ColCollect) (err error) {
	err = global.DB.Save(&colCollect).Error
	return err
}

// GetColCollect 根据id获取ColCollect记录
// Author [88act](https://github.com/88act)
func (m *ColCollectService) GetColCollect(id uint, fields string) (obj business.ColCollect, err error) {

	if utils.IsEmpty(fields) {
		err = global.DB.Where("id = ?", id).First(&obj).Error
	} else {
		err = global.DB.Select(fields).Where("id = ?", id).First(&obj).Error
	}

	//如果有图片image类型，更新图片path
	obj.MapData = make(map[string]string)
	return obj, err
}

// GetColCollectInfoList 分页获取ColCollect记录
// Author [88act](https://github.com/88act)
func (m *ColCollectService) GetColCollectInfoList(info bizReq.ColCollectSearch, createdAtBetween []string, fields string) (list []business.ColCollectMini, total int64, err error) {
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
	if len(createdAtBetween) >= 2 {
		db = db.Where("`created_at` BETWEEN ? AND ?", createdAtBetween[0], createdAtBetween[1])
	}

	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Name != "" {
		db = db.Where("`name` LIKE ?", "%"+info.Name+"%")
	}
	if info.ToTable != "" {
		db = db.Where("`to_table` = ?", info.ToTable)
	}
	if info.NowId != "" {
		db = db.Where("`now_id` = ?", info.NowId)
	}
	if info.PageNow != nil {
		db = db.Where("`page_now` = ?", info.PageNow)
	}
	if info.StatusRun != nil {
		db = db.Where("`status_run` = ?", info.StatusRun)
	}
	if info.Status != nil {
		db = db.Where("`status` = ?", info.Status)
	}

	err = db.Count(&total).Error
	if err != nil {
		return colCollects, 0, err
	}
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
	//如果有图片image类型，更新图片path
	for i, v := range colCollects {
		v.MapData = make(map[string]string)
		colCollects[i] = v
	}
	return colCollects, total, err
}

// GetColCollectInfoListAll  分页获取ColCollect记录 (全部字段)
// Author [88act](https://github.com/88act)
func (m *ColCollectService) GetColCollectInfoListAll(info bizReq.ColCollectSearch, createdAtBetween []string, fields string) (list []business.ColCollect, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.ColCollect{})
	var colCollects []business.ColCollect
	//var colCollects []business.ColCollectMini

	if info.ID > 0 {
		db = db.Where("`id` = ?", info.ID)
	}
	if len(createdAtBetween) >= 2 {
		db = db.Where("`created_at` BETWEEN ? AND ?", createdAtBetween[0], createdAtBetween[1])
	}

	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Name != "" {
		db = db.Where("`name` LIKE ?", "%"+info.Name+"%")
	}
	if info.ToTable != "" {
		db = db.Where("`to_table` = ?", info.ToTable)
	}
	if info.NowId != "" {
		db = db.Where("`now_id` = ?", info.NowId)
	}
	if info.PageNow != nil {
		db = db.Where("`page_now` = ?", info.PageNow)
	}
	if info.StatusRun != nil {
		db = db.Where("`status_run` = ?", info.StatusRun)
	}
	if info.Status != nil {
		db = db.Where("`status` = ?", info.Status)
	}

	err = db.Count(&total).Error
	if err != nil {
		return colCollects, 0, err
	}

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
	//如果有图片image类型，更新图片path
	for i, v := range colCollects {
		v.MapData = make(map[string]string)
		colCollects[i] = v
	}
	return colCollects, total, err
}
