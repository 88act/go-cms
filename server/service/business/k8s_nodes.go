package business

import (
	"go-cms/global"
	"go-cms/model/business"
	bizReq "go-cms/model/business/request"
	"go-cms/model/common/request"
	"go-cms/utils"
	"sync"
)

type K8sNodesService struct {
}

var once_K8sNodes sync.Once = sync.Once{}
var obj_K8sNodesService *K8sNodesService

//获取单例
func GetK8sNodesSev() *K8sNodesService {
	once_K8sNodes.Do(func() {
		obj_K8sNodesService = new(K8sNodesService)
		//instanse.init()
	})
	return obj_K8sNodesService
}

// Create 创建K8sNodes记录
// Author [88act](https://github.com/88act)
func (m *K8sNodesService) Create(data business.K8sNodes) (id uint, err error) {
	err = global.DB.Create(&data).Error
	if err != nil {
		return 0, err
	}
	return data.ID, err
}

// Delete 删除K8sNodes记录
// Author [88act](https://github.com/88act)
func (m *K8sNodesService) Delete(data business.K8sNodes) (err error) {
	err = global.DB.Delete(&data).Error
	return err
}

// DeleteByIds 批量删除K8sNodes记录
// Author [88act](https://github.com/88act)
func (m *K8sNodesService) DeleteByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]business.K8sNodes{}, "id in ?", ids.Ids).Error
	return err
}

// Update  更新K8sNodes记录
// Author [88act](https://github.com/88act)
func (m *K8sNodesService) Update(data business.K8sNodes) (err error) {
	err = global.DB.Save(&data).Error
	return err
}

// UpdateByMap  更新K8sNodes记录 by Map
// values := map[string]interface{}{
// 	"status":0,
// 	"from": hash,
// }
// Author [88act](https://github.com/88act)
func (m *K8sNodesService) UpdateByMap(data business.K8sNodes, mapData map[string]interface{}) (err error) {
	err = global.DB.Model(&data).Updates(mapData).Error
	return err
}

// Get 根据id获取K8sNodes记录
// Author [88act](https://github.com/88act)
func (m *K8sNodesService) Get(id uint, fields string) (obj business.K8sNodes, err error) {

	if utils.IsEmpty(fields) {
		err = global.DB.Where("id = ?", id).First(&obj).Error
	} else {
		err = global.DB.Select(fields).Where("id = ?", id).First(&obj).Error
	}

	//如果有图片image类型，更新图片path
	obj.MapData = make(map[string]string)
	return obj, err
}

// GetList 分页获取K8sNodes记录
// Author [88act](https://github.com/88act)
func (m *K8sNodesService) GetList(info bizReq.K8sNodesSearch, createdAtBetween []string, fields string) (list []business.K8sNodesMini, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.K8sNodes{})
	//var k8sNodess []business.K8sNodes
	var k8sNodess []business.K8sNodesMini

	//修改 by ljd
	if info.ID > 0 {
		db = db.Where("`id` = ?", info.ID)
	}
	if len(createdAtBetween) >= 2 {
		db = db.Where("`created_at` BETWEEN ? AND ?", createdAtBetween[0], createdAtBetween[1])
	}

	// 如果有条件搜索 下方会自动创建搜索语句
	if info.NodeName != "" {
		db = db.Where("`node_name` = ?", info.NodeName)
	}
	if info.Ip != "" {
		db = db.Where("`ip` = ?", info.Ip)
	}
	if info.Status != "" {
		db = db.Where("`status` = ?", info.Status)
	}
	if info.Roles != "" {
		db = db.Where("`roles` = ?", info.Roles)
	}
	if info.Label != "" {
		db = db.Where("`label` = ?", info.Label)
	}

	err = db.Count(&total).Error
	if err != nil {
		return k8sNodess, 0, err
	}
	//err = db.Limit(limit).Offset(offset).Find(&k8sNodess).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&k8sNodess).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&k8sNodess).Error
	}
	//如果有图片image类型，更新图片path
	for i, v := range k8sNodess {
		v.MapData = make(map[string]string)
		k8sNodess[i] = v
	}
	return k8sNodess, total, err
}

//GetListAll 分页获取K8sNodes记录 (全部字段)
// Author [88act](https://github.com/88act)
func (m *K8sNodesService) GetListAll(info bizReq.K8sNodesSearch, createdAtBetween []string, fields string) (list []business.K8sNodes, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.K8sNodes{})
	var k8sNodess []business.K8sNodes
	//var k8sNodess []business.K8sNodesMini

	if info.ID > 0 {
		db = db.Where("`id` = ?", info.ID)
	}
	if len(createdAtBetween) >= 2 {
		db = db.Where("`created_at` BETWEEN ? AND ?", createdAtBetween[0], createdAtBetween[1])
	}

	// 如果有条件搜索 下方会自动创建搜索语句
	if info.NodeName != "" {
		db = db.Where("`node_name` = ?", info.NodeName)
	}
	if info.Ip != "" {
		db = db.Where("`ip` = ?", info.Ip)
	}
	if info.Status != "" {
		db = db.Where("`status` = ?", info.Status)
	}
	if info.Roles != "" {
		db = db.Where("`roles` = ?", info.Roles)
	}
	if info.Label != "" {
		db = db.Where("`label` = ?", info.Label)
	}

	err = db.Count(&total).Error
	if err != nil {
		return k8sNodess, 0, err
	}

	//err = db.Limit(limit).Offset(offset).Find(&k8sNodess).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&k8sNodess).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&k8sNodess).Error
	}
	//如果有图片image类型，更新图片path
	for i, v := range k8sNodess {
		v.MapData = make(map[string]string)
		k8sNodess[i] = v
	}
	return k8sNodess, total, err
}
