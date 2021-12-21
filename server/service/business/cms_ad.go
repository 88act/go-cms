package business

import (
	"go-cms/global"
	"go-cms/model/business" 
     bizReq "go-cms/model/business/request"
    "go-cms/model/common/request"
    "go-cms/service/common"
    "go-cms/utils"
    "sync"
)

type CmsAdService struct {
}

var once_CmsAd sync.Once = sync.Once{}
var obj_CmsAdService *CmsAdService

//获取单例 
func GetCmsAdService() *CmsAdService {
	once_CmsAd.Do(func() {
		obj_CmsAdService= new(CmsAdService)
		//instanse.init()
	})
	return obj_CmsAdService
}



// CreateCmsAd 创建CmsAd记录
// Author [88act](https://github.com/88act)
func (m *CmsAdService) CreateCmsAd(cmsAd business.CmsAd) (err error) {
	err = global.DB.Create(&cmsAd).Error
	return err
}

// DeleteCmsAd 删除CmsAd记录
// Author [88act](https://github.com/88act)
func (m *CmsAdService)DeleteCmsAd(cmsAd business.CmsAd) (err error) {
	err = global.DB.Delete(&cmsAd).Error
	return err
}

// DeleteCmsAdByIds 批量删除CmsAd记录
// Author [88act](https://github.com/88act)
func (m *CmsAdService)DeleteCmsAdByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]business.CmsAd{},"id in ?",ids.Ids).Error
	return err
}

// UpdateCmsAd 更新CmsAd记录
// Author [88act](https://github.com/88act)
func (m *CmsAdService)UpdateCmsAd(cmsAd business.CmsAd) (err error) {
	err = global.DB.Save(&cmsAd).Error
	return err
}

// GetCmsAd 根据id获取CmsAd记录
// Author [88act](https://github.com/88act)
func (m *CmsAdService)GetCmsAd(id uint,fields string ) ( obj business.CmsAd,err error ) {
 
    if utils.IsEmpty(fields) {
        err = global.DB.Where("id = ?", id).First(&obj).Error 
        	} else {
        err = global.DB.Select(fields).Where("id = ?", id).First(&obj).Error  
	}

   //如果有图片image类型，更新图片path
    obj.MapData = make(map[string]string) 
    if !utils.IsEmpty(obj.Image) {
        _,obj.MapData[obj.Image] = common.GetCommonFileService().GetPathByGuid(obj.Image)
    }  
    return  obj,err
}


// GetCmsAdInfoList 分页获取CmsAd记录
// Author [88act](https://github.com/88act)
func (m *CmsAdService)GetCmsAdInfoList(info bizReq .CmsAdSearch, createdAtBetween []string,fields string) (list []business.CmsAdMini, total int64,err error) {
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
	if  len(createdAtBetween) >= 2 {
		db = db.Where("`created_at` BETWEEN ? AND ?", createdAtBetween[0], createdAtBetween[1])
	}

    // 如果有条件搜索 下方会自动创建搜索语句
    if info.MediaType != nil {
        db = db.Where("`media_type` = ?",info.MediaType)
    }
    if info.Status != nil {
        db = db.Where("`status` = ?",info.Status)
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
            _, v.MapData[v.Image] = common.GetCommonFileService().GetPathByGuid(v.Image)
        }
	  cmsAds[i] = v
	}
	return cmsAds, total,err
}
 

// GetCmsAdInfoListAll  分页获取CmsAd记录 (全部字段)
// Author [88act](https://github.com/88act)
func (m *CmsAdService)GetCmsAdInfoListAll(info bizReq .CmsAdSearch, createdAtBetween []string,fields string) (list []business.CmsAd, total int64,err error) {
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
        db = db.Where("`media_type` = ?",info.MediaType)
    }
    if info.Status != nil {
        db = db.Where("`status` = ?",info.Status)
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
            _, v.MapData[v.Image] = common.GetCommonFileService().GetPathByGuid(v.Image)
        }
	  cmsAds[i] = v
	}
	return cmsAds, total, err
}
