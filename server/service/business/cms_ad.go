package business

import (
	"go-cms/global"
	"go-cms/model/business"
	bizReq "go-cms/model/business/request"
	"go-cms/model/common/request"
	comSev "go-cms/service/common"
	"go-cms/utils"
	"sync"
)

type CmsAdService struct {
}

var once_CmsAd sync.Once = sync.Once{}
var obj_CmsAdService *CmsAdService

//获取单例
func GetCmsAdSev() *CmsAdService {
	once_CmsAd.Do(func() {
		obj_CmsAdService = new(CmsAdService)
		//instanse.init()
	})
	return obj_CmsAdService
}

// Create 创建CmsAd记录
// Author [88act](https://github.com/88act)
func (m *CmsAdService) Create(data business.CmsAd) (id uint, err error) {
	err = global.DB.Create(&data).Error
	if err != nil {
		return 0, err
	}
	return data.ID, err
}

// Delete 删除CmsAd记录
// Author [88act](https://github.com/88act)
func (m *CmsAdService) Delete(data business.CmsAd) (err error) {
	err = global.DB.Delete(&data).Error
	return err
}

// DeleteByIds 批量删除CmsAd记录
// Author [88act](https://github.com/88act)
func (m *CmsAdService) DeleteByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]business.CmsAd{}, "id in ?", ids.Ids).Error
	return err
}

// Update  更新CmsAd记录
// Author [88act](https://github.com/88act)
func (m *CmsAdService) Update(data business.CmsAd) (err error) {
	err = global.DB.Save(&data).Error
	return err
}

// UpdateByMap  更新CmsAd记录 by Map
// values := map[string]interface{}{
// 	"status":0,
// 	"from": hash,
// }
// Author [88act](https://github.com/88act)
func (m *CmsAdService) UpdateByMap(data business.CmsAd, mapData map[string]interface{}) (err error) {
	err = global.DB.Model(&data).Updates(mapData).Error
	return err
}

// Get 根据id获取CmsAd记录
// Author [88act](https://github.com/88act)
func (m *CmsAdService) Get(id uint, fields string) (obj business.CmsAd, err error) {

	if utils.IsEmpty(fields) {
		err = global.DB.Where("id = ?", id).First(&obj).Error
	} else {
		err = global.DB.Select(fields).Where("id = ?", id).First(&obj).Error
	}

	//如果有图片image类型，更新图片path
	obj.MapData = make(map[string]string)
	if !utils.IsEmpty(obj.Image) {
		_, obj.MapData[obj.Image] = comSev.GetCommonFileSev().GetPathByGuid(obj.Image)
	}
	return obj, err
}

// GetList 分页获取CmsAd记录
// Author [88act](https://github.com/88act)
func (m *CmsAdService) GetList(info bizReq.CmsAdSearch, createdAtBetween []string, fields string) (list []business.CmsAdMini, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.CmsAd{})
	//var cmsAds []business.CmsAd
	var cmsAds []business.CmsAdMini

	//修改 by ljd
	if info.ID > 0 {
		db = db.Where("`id` = ?", info.ID)
	}
	if len(createdAtBetween) >= 2 {
		db = db.Where("`created_at` BETWEEN ? AND ?", createdAtBetween[0], createdAtBetween[1])
	}

	// 如果有条件搜索 下方会自动创建搜索语句
	if info.MediaType != nil {
		db = db.Where("`media_type` = ?", info.MediaType)
	}
	if info.Status != nil {
		db = db.Where("`status` = ?", info.Status)
	}

	err = db.Count(&total).Error
	if err != nil {
		return cmsAds, 0, err
	}
	//err = db.Limit(limit).Offset(offset).Find(&cmsAds).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&cmsAds).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&cmsAds).Error
	}
	//如果有图片image类型，更新图片path
	for i, v := range cmsAds {
		v.MapData = make(map[string]string)
		if !utils.IsEmpty(v.Image) {
			_, v.MapData[v.Image] = comSev.GetCommonFileSev().GetPathByGuid(v.Image)
		}
		cmsAds[i] = v
	}
	return cmsAds, total, err
}

//GetListAll 分页获取CmsAd记录 (全部字段)
// Author [88act](https://github.com/88act)
func (m *CmsAdService) GetListAll(info bizReq.CmsAdSearch, createdAtBetween []string, fields string) (list []business.CmsAd, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.CmsAd{})
	var cmsAds []business.CmsAd
	//var cmsAds []business.CmsAdMini

	if info.ID > 0 {
		db = db.Where("`id` = ?", info.ID)
	}
	if len(createdAtBetween) >= 2 {
		db = db.Where("`created_at` BETWEEN ? AND ?", createdAtBetween[0], createdAtBetween[1])
	}

	// 如果有条件搜索 下方会自动创建搜索语句
	if info.MediaType != nil {
		db = db.Where("`media_type` = ?", info.MediaType)
	}
	if info.Status != nil {
		db = db.Where("`status` = ?", info.Status)
	}

	err = db.Count(&total).Error
	if err != nil {
		return cmsAds, 0, err
	}

	//err = db.Limit(limit).Offset(offset).Find(&cmsAds).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&cmsAds).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&cmsAds).Error
	}
	//如果有图片image类型，更新图片path
	for i, v := range cmsAds {
		v.MapData = make(map[string]string)
		if !utils.IsEmpty(v.Image) {
			_, v.MapData[v.Image] = comSev.GetCommonFileSev().GetPathByGuid(v.Image)
		}
		cmsAds[i] = v
	}
	return cmsAds, total, err
}
