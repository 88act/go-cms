package business

import (
	"go-cms/global"
	"go-cms/model/business"
	bizReq "go-cms/model/business/request"
	"go-cms/model/common/request"
	"go-cms/utils"
	"sync"
)

type K8sPodsService struct {
}

var once_K8sPods sync.Once = sync.Once{}
var obj_K8sPodsService *K8sPodsService

//获取单例
func GetK8sPodsSev() *K8sPodsService {
	once_K8sPods.Do(func() {
		obj_K8sPodsService = new(K8sPodsService)
		//instanse.init()
	})
	return obj_K8sPodsService
}

// Create 创建K8sPods记录
// Author [88act](https://github.com/88act)
func (m *K8sPodsService) Create(data business.K8sPods) (id uint, err error) {
	err = global.DB.Create(&data).Error
	if err != nil {
		return 0, err
	}
	return data.ID, err
}

// Delete 删除K8sPods记录
// Author [88act](https://github.com/88act)
func (m *K8sPodsService) Delete(data business.K8sPods) (err error) {
	err = global.DB.Delete(&data).Error
	return err
}

// DeleteByIds 批量删除K8sPods记录
// Author [88act](https://github.com/88act)
func (m *K8sPodsService) DeleteByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]business.K8sPods{}, "id in ?", ids.Ids).Error
	return err
}

// Update  更新K8sPods记录
// Author [88act](https://github.com/88act)
func (m *K8sPodsService) Update(data business.K8sPods) (err error) {
	err = global.DB.Save(&data).Error
	return err
}

// UpdateByMap  更新K8sPods记录 by Map
// values := map[string]interface{}{
// 	"status":0,
// 	"from": hash,
// }
// Author [88act](https://github.com/88act)
func (m *K8sPodsService) UpdateByMap(data business.K8sPods, mapData map[string]interface{}) (err error) {
	err = global.DB.Model(&data).Updates(mapData).Error
	return err
}

// Get 根据id获取K8sPods记录
// Author [88act](https://github.com/88act)
func (m *K8sPodsService) Get(id uint, fields string) (obj business.K8sPods, err error) {

	if utils.IsEmpty(fields) {
		err = global.DB.Where("id = ?", id).First(&obj).Error
	} else {
		err = global.DB.Select(fields).Where("id = ?", id).First(&obj).Error
	}

	//如果有图片image类型，更新图片path
	obj.MapData = make(map[string]string)
	return obj, err
}

// GetList 分页获取K8sPods记录
// Author [88act](https://github.com/88act)
func (m *K8sPodsService) GetList(info bizReq.K8sPodsSearch, createdAtBetween []string, fields string) (list []business.K8sPodsMini, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.K8sPods{})
	//var k8sPodss []business.K8sPods
	var k8sPodss []business.K8sPodsMini

	//修改 by ljd
	if info.ID > 0 {
		db = db.Where("`id` = ?", info.ID)
	}
	if len(createdAtBetween) >= 2 {
		db = db.Where("`created_at` BETWEEN ? AND ?", createdAtBetween[0], createdAtBetween[1])
	}

	// 如果有条件搜索 下方会自动创建搜索语句
	if info.PodName != "" {
		db = db.Where("`pod_name` = ?", info.PodName)
	}
	if info.Status != "" {
		db = db.Where("`status` = ?", info.Status)
	}
	if info.NameSpace != "" {
		db = db.Where("`name_space` = ?", info.NameSpace)
	}

	err = db.Count(&total).Error
	if err != nil {
		return k8sPodss, 0, err
	}
	//err = db.Limit(limit).Offset(offset).Find(&k8sPodss).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&k8sPodss).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&k8sPodss).Error
	}
	//如果有图片image类型，更新图片path
	for i, v := range k8sPodss {
		v.MapData = make(map[string]string)
		k8sPodss[i] = v
	}
	return k8sPodss, total, err
}

//GetListAll 分页获取K8sPods记录 (全部字段)
// Author [88act](https://github.com/88act)
func (m *K8sPodsService) GetListAll(info bizReq.K8sPodsSearch, createdAtBetween []string, fields string) (list []business.K8sPods, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.K8sPods{})
	var k8sPodss []business.K8sPods
	//var k8sPodss []business.K8sPodsMini

	if info.ID > 0 {
		db = db.Where("`id` = ?", info.ID)
	}
	if len(createdAtBetween) >= 2 {
		db = db.Where("`created_at` BETWEEN ? AND ?", createdAtBetween[0], createdAtBetween[1])
	}

	// 如果有条件搜索 下方会自动创建搜索语句
	if info.PodName != "" {
		db = db.Where("`pod_name` = ?", info.PodName)
	}
	if info.Status != "" {
		db = db.Where("`status` = ?", info.Status)
	}
	if info.NameSpace != "" {
		db = db.Where("`name_space` = ?", info.NameSpace)
	}

	err = db.Count(&total).Error
	if err != nil {
		return k8sPodss, 0, err
	}

	//err = db.Limit(limit).Offset(offset).Find(&k8sPodss).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&k8sPodss).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&k8sPodss).Error
	}
	//如果有图片image类型，更新图片path
	for i, v := range k8sPodss {
		v.MapData = make(map[string]string)
		k8sPodss[i] = v
	}
	return k8sPodss, total, err
}
