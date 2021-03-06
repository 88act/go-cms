package business

import (
	"go-cms/global"
	"go-cms/model/business"
	bizReq "go-cms/model/business/request"
	"go-cms/model/common/request"
	"go-cms/utils"
	"sync"
)

type K8sClustersService struct {
}

var once_K8sClusters sync.Once = sync.Once{}
var obj_K8sClustersService *K8sClustersService

//获取单例
func GetK8sClustersSev() *K8sClustersService {
	once_K8sClusters.Do(func() {
		obj_K8sClustersService = new(K8sClustersService)
		//instanse.init()
	})
	return obj_K8sClustersService
}

// Create 创建K8sClusters记录
// Author [88act](https://github.com/88act)
func (m *K8sClustersService) Create(data business.K8sClusters) (id uint, err error) {
	err = global.DB.Create(&data).Error
	if err != nil {
		return 0, err
	}
	return data.ID, err
}

// Delete 删除K8sClusters记录
// Author [88act](https://github.com/88act)
func (m *K8sClustersService) Delete(data business.K8sClusters) (err error) {
	err = global.DB.Delete(&data).Error
	return err
}

// DeleteByIds 批量删除K8sClusters记录
// Author [88act](https://github.com/88act)
func (m *K8sClustersService) DeleteByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]business.K8sClusters{}, "id in ?", ids.Ids).Error
	return err
}

// Update  更新K8sClusters记录
// Author [88act](https://github.com/88act)
func (m *K8sClustersService) Update(data business.K8sClusters) (err error) {
	err = global.DB.Save(&data).Error
	return err
}

// UpdateByMap  更新K8sClusters记录 by Map
// values := map[string]interface{}{
// 	"status":0,
// 	"from": hash,
// }
// Author [88act](https://github.com/88act)
func (m *K8sClustersService) UpdateByMap(data business.K8sClusters, mapData map[string]interface{}) (err error) {
	err = global.DB.Model(&data).Updates(mapData).Error
	return err
}

// Get 根据id获取K8sClusters记录
// Author [88act](https://github.com/88act)
func (m *K8sClustersService) Get(id uint, fields string) (obj business.K8sClusters, err error) {

	if utils.IsEmpty(fields) {
		err = global.DB.Where("id = ?", id).First(&obj).Error
	} else {
		err = global.DB.Select(fields).Where("id = ?", id).First(&obj).Error
	}

	//如果有图片image类型，更新图片path
	obj.MapData = make(map[string]string)
	return obj, err
}

// GetList 分页获取K8sClusters记录
// Author [88act](https://github.com/88act)
func (m *K8sClustersService) GetList(info bizReq.K8sClustersSearch, createdAtBetween []string, fields string) (list []business.K8sClustersMini, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.K8sClusters{})
	//var k8sClusterss []business.K8sClusters
	var k8sClusterss []business.K8sClustersMini

	//修改 by ljd
	if info.ID > 0 {
		db = db.Where("`id` = ?", info.ID)
	}
	if len(createdAtBetween) >= 2 {
		db = db.Where("`created_at` BETWEEN ? AND ?", createdAtBetween[0], createdAtBetween[1])
	}

	// 如果有条件搜索 下方会自动创建搜索语句
	if info.ClusterName != "" {
		db = db.Where("`cluster_name` = ?", info.ClusterName)
	}

	err = db.Count(&total).Error
	if err != nil {
		return k8sClusterss, 0, err
	}
	//err = db.Limit(limit).Offset(offset).Find(&k8sClusterss).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&k8sClusterss).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&k8sClusterss).Error
	}
	//如果有图片image类型，更新图片path
	for i, v := range k8sClusterss {
		v.MapData = make(map[string]string)
		k8sClusterss[i] = v
	}
	return k8sClusterss, total, err
}

//GetListAll 分页获取K8sClusters记录 (全部字段)
// Author [88act](https://github.com/88act)
func (m *K8sClustersService) GetListAll(info bizReq.K8sClustersSearch, createdAtBetween []string, fields string) (list []business.K8sClusters, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.K8sClusters{})
	var k8sClusterss []business.K8sClusters
	//var k8sClusterss []business.K8sClustersMini

	if info.ID > 0 {
		db = db.Where("`id` = ?", info.ID)
	}
	if len(createdAtBetween) >= 2 {
		db = db.Where("`created_at` BETWEEN ? AND ?", createdAtBetween[0], createdAtBetween[1])
	}

	// 如果有条件搜索 下方会自动创建搜索语句
	if info.ClusterName != "" {
		db = db.Where("`cluster_name` = ?", info.ClusterName)
	}

	err = db.Count(&total).Error
	if err != nil {
		return k8sClusterss, 0, err
	}

	//err = db.Limit(limit).Offset(offset).Find(&k8sClusterss).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&k8sClusterss).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&k8sClusterss).Error
	}
	//如果有图片image类型，更新图片path
	for i, v := range k8sClusterss {
		v.MapData = make(map[string]string)
		k8sClusterss[i] = v
	}
	return k8sClusterss, total, err
}
