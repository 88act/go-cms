package business

import (
	"go-cms/global"
	"go-cms/model/business"
	bizReq "go-cms/model/business/request"
	"go-cms/model/common/request"
	"go-cms/utils"
	"sync"
)

type K8sDeploymentsService struct {
}

var once_K8sDeployments sync.Once = sync.Once{}
var obj_K8sDeploymentsService *K8sDeploymentsService

//获取单例
func GetK8sDeploymentsSev() *K8sDeploymentsService {
	once_K8sDeployments.Do(func() {
		obj_K8sDeploymentsService = new(K8sDeploymentsService)
		//instanse.init()
	})
	return obj_K8sDeploymentsService
}

// Create 创建K8sDeployments记录
// Author [88act](https://github.com/88act)
func (m *K8sDeploymentsService) Create(data business.K8sDeployments) (id uint, err error) {
	err = global.DB.Create(&data).Error
	if err != nil {
		return 0, err
	}
	return data.ID, err
}

// Delete 删除K8sDeployments记录
// Author [88act](https://github.com/88act)
func (m *K8sDeploymentsService) Delete(data business.K8sDeployments) (err error) {
	err = global.DB.Delete(&data).Error
	return err
}

// DeleteByIds 批量删除K8sDeployments记录
// Author [88act](https://github.com/88act)
func (m *K8sDeploymentsService) DeleteByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]business.K8sDeployments{}, "id in ?", ids.Ids).Error
	return err
}

// Update  更新K8sDeployments记录
// Author [88act](https://github.com/88act)
func (m *K8sDeploymentsService) Update(data business.K8sDeployments) (err error) {
	err = global.DB.Save(&data).Error
	return err
}

// UpdateByMap  更新K8sDeployments记录 by Map
// values := map[string]interface{}{
// 	"status":0,
// 	"from": hash,
// }
// Author [88act](https://github.com/88act)
func (m *K8sDeploymentsService) UpdateByMap(data business.K8sDeployments, mapData map[string]interface{}) (err error) {
	err = global.DB.Model(&data).Updates(mapData).Error
	return err
}

// Get 根据id获取K8sDeployments记录
// Author [88act](https://github.com/88act)
func (m *K8sDeploymentsService) Get(id uint, fields string) (obj business.K8sDeployments, err error) {

	if utils.IsEmpty(fields) {
		err = global.DB.Where("id = ?", id).First(&obj).Error
	} else {
		err = global.DB.Select(fields).Where("id = ?", id).First(&obj).Error
	}

	//如果有图片image类型，更新图片path
	obj.MapData = make(map[string]string)
	return obj, err
}

// GetList 分页获取K8sDeployments记录
// Author [88act](https://github.com/88act)
func (m *K8sDeploymentsService) GetList(info bizReq.K8sDeploymentsSearch, createdAtBetween []string, fields string) (list []business.K8sDeploymentsMini, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.K8sDeployments{})
	//var k8sDeploymentss []business.K8sDeployments
	var k8sDeploymentss []business.K8sDeploymentsMini

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
	if info.Deployment != "" {
		db = db.Where("`deployment` = ?", info.Deployment)
	}

	err = db.Count(&total).Error
	if err != nil {
		return k8sDeploymentss, 0, err
	}
	//err = db.Limit(limit).Offset(offset).Find(&k8sDeploymentss).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&k8sDeploymentss).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&k8sDeploymentss).Error
	}
	//如果有图片image类型，更新图片path
	for i, v := range k8sDeploymentss {
		v.MapData = make(map[string]string)
		k8sDeploymentss[i] = v
	}
	return k8sDeploymentss, total, err
}

//GetListAll 分页获取K8sDeployments记录 (全部字段)
// Author [88act](https://github.com/88act)
func (m *K8sDeploymentsService) GetListAll(info bizReq.K8sDeploymentsSearch, createdAtBetween []string, fields string) (list []business.K8sDeployments, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.K8sDeployments{})
	var k8sDeploymentss []business.K8sDeployments
	//var k8sDeploymentss []business.K8sDeploymentsMini

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
	if info.Deployment != "" {
		db = db.Where("`deployment` = ?", info.Deployment)
	}

	err = db.Count(&total).Error
	if err != nil {
		return k8sDeploymentss, 0, err
	}

	//err = db.Limit(limit).Offset(offset).Find(&k8sDeploymentss).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&k8sDeploymentss).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&k8sDeploymentss).Error
	}
	//如果有图片image类型，更新图片path
	for i, v := range k8sDeploymentss {
		v.MapData = make(map[string]string)
		k8sDeploymentss[i] = v
	}
	return k8sDeploymentss, total, err
}
