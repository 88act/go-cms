package business

import (
	"go-cms/global"
	"go-cms/model/business"
	bizReq "go-cms/model/business/request"
	"go-cms/model/common/request"
	"go-cms/utils"
	"sync"
)

type SysSuperBuilderHistoriesService struct {
}

var once_SysSuperBuilderHistories sync.Once = sync.Once{}
var obj_SysSuperBuilderHistoriesService *SysSuperBuilderHistoriesService

//获取单例
func GetSysSuperBuilderHistoriesSev() *SysSuperBuilderHistoriesService {
	once_SysSuperBuilderHistories.Do(func() {
		obj_SysSuperBuilderHistoriesService = new(SysSuperBuilderHistoriesService)
		//instanse.init()
	})
	return obj_SysSuperBuilderHistoriesService
}

// Create 创建SysSuperBuilderHistories记录
// Author [88act](https://github.com/88act)
func (m *SysSuperBuilderHistoriesService) Create(data business.SysSuperBuilderHistories) (id uint, err error) {
	err = global.DB.Create(&data).Error
	if err != nil {
		return 0, err
	}
	return data.ID, err
}

// Delete 删除SysSuperBuilderHistories记录
// Author [88act](https://github.com/88act)
func (m *SysSuperBuilderHistoriesService) Delete(data business.SysSuperBuilderHistories) (err error) {
	err = global.DB.Delete(&data).Error
	return err
}

// DeleteByIds 批量删除SysSuperBuilderHistories记录
// Author [88act](https://github.com/88act)
func (m *SysSuperBuilderHistoriesService) DeleteByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]business.SysSuperBuilderHistories{}, "id in ?", ids.Ids).Error
	return err
}

// Update  更新SysSuperBuilderHistories记录
// Author [88act](https://github.com/88act)
func (m *SysSuperBuilderHistoriesService) Update(data business.SysSuperBuilderHistories) (err error) {
	err = global.DB.Save(&data).Error
	return err
}

// UpdateByMap  更新SysSuperBuilderHistories记录 by Map
// values := map[string]interface{}{
// 	"status":0,
// 	"from": hash,
// }
// Author [88act](https://github.com/88act)
func (m *SysSuperBuilderHistoriesService) UpdateByMap(data business.SysSuperBuilderHistories, mapData map[string]interface{}) (err error) {
	err = global.DB.Model(&data).Updates(mapData).Error
	return err
}

// Get 根据id获取SysSuperBuilderHistories记录
// Author [88act](https://github.com/88act)
func (m *SysSuperBuilderHistoriesService) Get(id uint, fields string) (obj business.SysSuperBuilderHistories, err error) {

	if utils.IsEmpty(fields) {
		err = global.DB.Where("id = ?", id).First(&obj).Error
	} else {
		err = global.DB.Select(fields).Where("id = ?", id).First(&obj).Error
	}

	//如果有图片image类型，更新图片path
	obj.MapData = make(map[string]string)
	return obj, err
}

// GetList 分页获取SysSuperBuilderHistories记录
// Author [88act](https://github.com/88act)
func (m *SysSuperBuilderHistoriesService) GetList(info bizReq.SysSuperBuilderHistoriesSearch, createdAtBetween []string, fields string) (list []business.SysSuperBuilderHistoriesMini, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.SysSuperBuilderHistories{})
	//var sysSuperBuilderHistoriess []business.SysSuperBuilderHistories
	var sysSuperBuilderHistoriess []business.SysSuperBuilderHistoriesMini

	//修改 by ljd
	if info.ID > 0 {
		db = db.Where("`id` = ?", info.ID)
	}
	if len(createdAtBetween) >= 2 {
		db = db.Where("`created_at` BETWEEN ? AND ?", createdAtBetween[0], createdAtBetween[1])
	}

	// 如果有条件搜索 下方会自动创建搜索语句
	if info.TableName != "" {
		db = db.Where("`table_name` LIKE ?", "%"+info.TableName+"%")
	}
	if info.StructName != "" {
		db = db.Where("`struct_name` LIKE ?", "%"+info.StructName+"%")
	}
	if info.StructCnName != "" {
		db = db.Where("`struct_cn_name` LIKE ?", "%"+info.StructCnName+"%")
	}

	err = db.Count(&total).Error
	if err != nil {
		return sysSuperBuilderHistoriess, 0, err
	}
	//err = db.Limit(limit).Offset(offset).Find(&sysSuperBuilderHistoriess).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&sysSuperBuilderHistoriess).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&sysSuperBuilderHistoriess).Error
	}
	//如果有图片image类型，更新图片path
	for i, v := range sysSuperBuilderHistoriess {
		v.MapData = make(map[string]string)
		sysSuperBuilderHistoriess[i] = v
	}
	return sysSuperBuilderHistoriess, total, err
}

//GetListAll 分页获取SysSuperBuilderHistories记录 (全部字段)
// Author [88act](https://github.com/88act)
func (m *SysSuperBuilderHistoriesService) GetListAll(info bizReq.SysSuperBuilderHistoriesSearch, createdAtBetween []string, fields string) (list []business.SysSuperBuilderHistories, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.SysSuperBuilderHistories{})
	var sysSuperBuilderHistoriess []business.SysSuperBuilderHistories
	//var sysSuperBuilderHistoriess []business.SysSuperBuilderHistoriesMini

	if info.ID > 0 {
		db = db.Where("`id` = ?", info.ID)
	}
	if len(createdAtBetween) >= 2 {
		db = db.Where("`created_at` BETWEEN ? AND ?", createdAtBetween[0], createdAtBetween[1])
	}

	// 如果有条件搜索 下方会自动创建搜索语句
	if info.TableName != "" {
		db = db.Where("`table_name` LIKE ?", "%"+info.TableName+"%")
	}
	if info.StructName != "" {
		db = db.Where("`struct_name` LIKE ?", "%"+info.StructName+"%")
	}
	if info.StructCnName != "" {
		db = db.Where("`struct_cn_name` LIKE ?", "%"+info.StructCnName+"%")
	}

	err = db.Count(&total).Error
	if err != nil {
		return sysSuperBuilderHistoriess, 0, err
	}

	//err = db.Limit(limit).Offset(offset).Find(&sysSuperBuilderHistoriess).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&sysSuperBuilderHistoriess).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&sysSuperBuilderHistoriess).Error
	}
	//如果有图片image类型，更新图片path
	for i, v := range sysSuperBuilderHistoriess {
		v.MapData = make(map[string]string)
		sysSuperBuilderHistoriess[i] = v
	}
	return sysSuperBuilderHistoriess, total, err
}
