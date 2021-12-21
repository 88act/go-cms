package business

import (
	"go-cms/global"
	"go-cms/model/business"
	bizReq "go-cms/model/business/request"
	"go-cms/model/common/request"
	"go-cms/utils"
	"sync"
)

type K8sNamespacesService struct {
}

var once_K8sNamespaces sync.Once = sync.Once{}
var obj_K8sNamespacesService *K8sNamespacesService

//获取单例
func GetK8sNamespacesService() *K8sNamespacesService {
	once_K8sNamespaces.Do(func() {
		obj_K8sNamespacesService = new(K8sNamespacesService)
		//instanse.init()
	})
	return obj_K8sNamespacesService
}

// CreateK8sNamespaces 创建K8sNamespaces记录
// Author [88act](https://github.com/88act)
func (m *K8sNamespacesService) CreateK8sNamespaces(k8sNamespaces business.K8sNamespaces) (err error) {
	err = global.DB.Create(&k8sNamespaces).Error
	return err
}

// DeleteK8sNamespaces 删除K8sNamespaces记录
// Author [88act](https://github.com/88act)
func (m *K8sNamespacesService) DeleteK8sNamespaces(k8sNamespaces business.K8sNamespaces) (err error) {
	err = global.DB.Delete(&k8sNamespaces).Error
	return err
}

// DeleteK8sNamespacesByIds 批量删除K8sNamespaces记录
// Author [88act](https://github.com/88act)
func (m *K8sNamespacesService) DeleteK8sNamespacesByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]business.K8sNamespaces{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateK8sNamespaces 更新K8sNamespaces记录
// Author [88act](https://github.com/88act)
func (m *K8sNamespacesService) UpdateK8sNamespaces(k8sNamespaces business.K8sNamespaces) (err error) {
	err = global.DB.Save(&k8sNamespaces).Error
	return err
}

// GetK8sNamespaces 根据id获取K8sNamespaces记录
// Author [88act](https://github.com/88act)
func (m *K8sNamespacesService) GetK8sNamespaces(id uint, fields string) (obj business.K8sNamespaces, err error) {

	if utils.IsEmpty(fields) {
		err = global.DB.Where("id = ?", id).First(&obj).Error
	} else {
		err = global.DB.Select(fields).Where("id = ?", id).First(&obj).Error
	}

	//如果有图片image类型，更新图片path
	obj.MapData = make(map[string]string)
	return obj, err
}

// GetK8sNamespacesInfoList 分页获取K8sNamespaces记录
// Author [88act](https://github.com/88act)
func (m *K8sNamespacesService) GetK8sNamespacesInfoList(info bizReq.K8sNamespacesSearch, createdAtBetween []string, fields string) (list []business.K8sNamespacesMini, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.K8sNamespaces{})
	//var k8sNamespacess []business.K8sNamespaces
	var k8sNamespacess []business.K8sNamespacesMini

	//修改 by ljd
	if info.ID > 0 {
		db = db.Where("`id` = ?", info.ID)
	}
	if len(createdAtBetween) >= 2 {
		db = db.Where("`created_at` BETWEEN ? AND ?", createdAtBetween[0], createdAtBetween[1])
	}

	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Namespace != "" {
		db = db.Where("`namespace` = ?", info.Namespace)
	}
	if info.Status != "" {
		db = db.Where("`status` = ?", info.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return k8sNamespacess, 0, err
	}
	//err = db.Limit(limit).Offset(offset).Find(&k8sNamespacess).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&k8sNamespacess).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&k8sNamespacess).Error
	}
	//如果有图片image类型，更新图片path
	for i, v := range k8sNamespacess {
		v.MapData = make(map[string]string)
		k8sNamespacess[i] = v
	}
	return k8sNamespacess, total, err
}

// GetK8sNamespacesInfoListAll  分页获取K8sNamespaces记录 (全部字段)
// Author [88act](https://github.com/88act)
func (m *K8sNamespacesService) GetK8sNamespacesInfoListAll(info bizReq.K8sNamespacesSearch, createdAtBetween []string, fields string) (list []business.K8sNamespaces, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.K8sNamespaces{})
	var k8sNamespacess []business.K8sNamespaces
	//var k8sNamespacess []business.K8sNamespacesMini

	if info.ID > 0 {
		db = db.Where("`id` = ?", info.ID)
	}
	if len(createdAtBetween) >= 2 {
		db = db.Where("`created_at` BETWEEN ? AND ?", createdAtBetween[0], createdAtBetween[1])
	}

	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Namespace != "" {
		db = db.Where("`namespace` = ?", info.Namespace)
	}
	if info.Status != "" {
		db = db.Where("`status` = ?", info.Status)
	}

	err = db.Count(&total).Error
	if err != nil {
		return k8sNamespacess, 0, err
	}

	//err = db.Limit(limit).Offset(offset).Find(&k8sNamespacess).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&k8sNamespacess).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&k8sNamespacess).Error
	}
	//如果有图片image类型，更新图片path
	for i, v := range k8sNamespacess {
		v.MapData = make(map[string]string)
		k8sNamespacess[i] = v
	}
	return k8sNamespacess, total, err
}
