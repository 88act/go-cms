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
func GetK8sDeploymentsService() *K8sDeploymentsService {
	once_K8sDeployments.Do(func() {
		obj_K8sDeploymentsService = new(K8sDeploymentsService)
		//instanse.init()
	})
	return obj_K8sDeploymentsService
}

// CreateK8sDeployments 创建K8sDeployments记录
// Author [88act](https://github.com/88act)
func (m *K8sDeploymentsService) CreateK8sDeployments(k8sDeployments business.K8sDeployments) (err error) {
	err = global.DB.Create(&k8sDeployments).Error
	return err
}

// DeleteK8sDeployments 删除K8sDeployments记录
// Author [88act](https://github.com/88act)
func (m *K8sDeploymentsService) DeleteK8sDeployments(k8sDeployments business.K8sDeployments) (err error) {
	err = global.DB.Delete(&k8sDeployments).Error
	return err
}

// DeleteK8sDeploymentsByIds 批量删除K8sDeployments记录
// Author [88act](https://github.com/88act)
func (m *K8sDeploymentsService) DeleteK8sDeploymentsByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]business.K8sDeployments{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateK8sDeployments 更新K8sDeployments记录
// Author [88act](https://github.com/88act)
func (m *K8sDeploymentsService) UpdateK8sDeployments(k8sDeployments business.K8sDeployments) (err error) {
	err = global.DB.Save(&k8sDeployments).Error
	return err
}

// GetK8sDeployments 根据id获取K8sDeployments记录
// Author [88act](https://github.com/88act)
func (m *K8sDeploymentsService) GetK8sDeployments(id uint, fields string) (obj business.K8sDeployments, err error) {

	if utils.IsEmpty(fields) {
		err = global.DB.Where("id = ?", id).First(&obj).Error
	} else {
		err = global.DB.Select(fields).Where("id = ?", id).First(&obj).Error
	}

	//如果有图片image类型，更新图片path
	obj.MapData = make(map[string]string)
	return obj, err
}

// GetK8sDeploymentsInfoList 分页获取K8sDeployments记录
// Author [88act](https://github.com/88act)
func (m *K8sDeploymentsService) GetK8sDeploymentsInfoList(info bizReq.K8sDeploymentsSearch, createdAtBetween []string, fields string) (list []business.K8sDeploymentsMini, total int64, err error) {
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

// GetK8sDeploymentsInfoListAll  分页获取K8sDeployments记录 (全部字段)
// Author [88act](https://github.com/88act)
func (m *K8sDeploymentsService) GetK8sDeploymentsInfoListAll(info bizReq.K8sDeploymentsSearch, createdAtBetween []string, fields string) (list []business.K8sDeployments, total int64, err error) {
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
