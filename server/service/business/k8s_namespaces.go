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
func GetK8sNamespacesSev() *K8sNamespacesService {
	once_K8sNamespaces.Do(func() {
		obj_K8sNamespacesService = new(K8sNamespacesService)
		//instanse.init()
	})
	return obj_K8sNamespacesService
}

// Create 创建K8sNamespaces记录
// Author [88act](https://github.com/88act)
func (m *K8sNamespacesService) Create(data business.K8sNamespaces) (id uint, err error) {
	err = global.DB.Create(&data).Error
	if err != nil {
		return 0, err
	}
	return data.ID, err
}

// Delete 删除K8sNamespaces记录
// Author [88act](https://github.com/88act)
func (m *K8sNamespacesService) Delete(data business.K8sNamespaces) (err error) {
	err = global.DB.Delete(&data).Error
	return err
}

// DeleteByIds 批量删除K8sNamespaces记录
// Author [88act](https://github.com/88act)
func (m *K8sNamespacesService) DeleteByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]business.K8sNamespaces{}, "id in ?", ids.Ids).Error
	return err
}

// Update  更新K8sNamespaces记录
// Author [88act](https://github.com/88act)
func (m *K8sNamespacesService) Update(data business.K8sNamespaces) (err error) {
	err = global.DB.Save(&data).Error
	return err
}

// UpdateByMap  更新K8sNamespaces记录 by Map
// values := map[string]interface{}{
// 	"status":0,
// 	"from": hash,
// }
// Author [88act](https://github.com/88act)
func (m *K8sNamespacesService) UpdateByMap(data business.K8sNamespaces, mapData map[string]interface{}) (err error) {
	err = global.DB.Model(&data).Updates(mapData).Error
	return err
}

// Get 根据id获取K8sNamespaces记录
// Author [88act](https://github.com/88act)
func (m *K8sNamespacesService) Get(id uint, fields string) (obj business.K8sNamespaces, err error) {

	if utils.IsEmpty(fields) {
		err = global.DB.Where("id = ?", id).First(&obj).Error
	} else {
		err = global.DB.Select(fields).Where("id = ?", id).First(&obj).Error
	}

	//如果有图片image类型，更新图片path
	obj.MapData = make(map[string]string)
	return obj, err
}

// GetList 分页获取K8sNamespaces记录
// Author [88act](https://github.com/88act)
func (m *K8sNamespacesService) GetList(info bizReq.K8sNamespacesSearch, createdAtBetween []string, fields string) (list []business.K8sNamespacesMini, total int64, err error) {
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

//GetListAll 分页获取K8sNamespaces记录 (全部字段)
// Author [88act](https://github.com/88act)
func (m *K8sNamespacesService) GetListAll(info bizReq.K8sNamespacesSearch, createdAtBetween []string, fields string) (list []business.K8sNamespaces, total int64, err error) {
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
