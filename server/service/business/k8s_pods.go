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
func GetK8sPodsService() *K8sPodsService {
	once_K8sPods.Do(func() {
		obj_K8sPodsService = new(K8sPodsService)
		//instanse.init()
	})
	return obj_K8sPodsService
}

// CreateK8sPods 创建K8sPods记录
// Author [88act](https://github.com/88act)
func (m *K8sPodsService) CreateK8sPods(k8sPods business.K8sPods) (err error) {
	err = global.DB.Create(&k8sPods).Error
	return err
}

// DeleteK8sPods 删除K8sPods记录
// Author [88act](https://github.com/88act)
func (m *K8sPodsService) DeleteK8sPods(k8sPods business.K8sPods) (err error) {
	err = global.DB.Delete(&k8sPods).Error
	return err
}

// DeleteK8sPodsByIds 批量删除K8sPods记录
// Author [88act](https://github.com/88act)
func (m *K8sPodsService) DeleteK8sPodsByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]business.K8sPods{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateK8sPods 更新K8sPods记录
// Author [88act](https://github.com/88act)
func (m *K8sPodsService) UpdateK8sPods(k8sPods business.K8sPods) (err error) {
	err = global.DB.Save(&k8sPods).Error
	return err
}

// GetK8sPods 根据id获取K8sPods记录
// Author [88act](https://github.com/88act)
func (m *K8sPodsService) GetK8sPods(id uint, fields string) (obj business.K8sPods, err error) {

	if utils.IsEmpty(fields) {
		err = global.DB.Where("id = ?", id).First(&obj).Error
	} else {
		err = global.DB.Select(fields).Where("id = ?", id).First(&obj).Error
	}

	//如果有图片image类型，更新图片path
	obj.MapData = make(map[string]string)
	return obj, err
}

// GetK8sPodsInfoList 分页获取K8sPods记录
// Author [88act](https://github.com/88act)
func (m *K8sPodsService) GetK8sPodsInfoList(info bizReq.K8sPodsSearch, createdAtBetween []string, fields string) (list []business.K8sPodsMini, total int64, err error) {
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

// GetK8sPodsInfoListAll  分页获取K8sPods记录 (全部字段)
// Author [88act](https://github.com/88act)
func (m *K8sPodsService) GetK8sPodsInfoListAll(info bizReq.K8sPodsSearch, createdAtBetween []string, fields string) (list []business.K8sPods, total int64, err error) {
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
