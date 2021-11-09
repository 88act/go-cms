package business

import (
	"github.com/88act/go-cms/server/global"
	"github.com/88act/go-cms/server/model/business"
	"github.com/88act/go-cms/server/model/common/request"
    businessReq "github.com/88act/go-cms/server/model/business/request"
    "github.com/88act/go-cms/server/utils"
)

type CmsAdSeatService struct {
}

// CreateCmsAdSeat 创建CmsAdSeat记录
// Author [88act](https://github.com/88act)
func (cmsAdSeatService *CmsAdSeatService) CreateCmsAdSeat(cmsAdSeat business.CmsAdSeat) (err error) {
	err = global.DB.Create(&cmsAdSeat).Error
	return err
}

// DeleteCmsAdSeat 删除CmsAdSeat记录
// Author [88act](https://github.com/88act)
func (cmsAdSeatService *CmsAdSeatService)DeleteCmsAdSeat(cmsAdSeat business.CmsAdSeat) (err error) {
	err = global.DB.Delete(&cmsAdSeat).Error
	return err
}

// DeleteCmsAdSeatByIds 批量删除CmsAdSeat记录
// Author [88act](https://github.com/88act)
func (cmsAdSeatService *CmsAdSeatService)DeleteCmsAdSeatByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]business.CmsAdSeat{},"id in ?",ids.Ids).Error
	return err
}

// UpdateCmsAdSeat 更新CmsAdSeat记录
// Author [88act](https://github.com/88act)
func (cmsAdSeatService *CmsAdSeatService)UpdateCmsAdSeat(cmsAdSeat business.CmsAdSeat) (err error) {
	err = global.DB.Save(&cmsAdSeat).Error
	return err
}

// GetCmsAdSeat 根据id获取CmsAdSeat记录
// Author [88act](https://github.com/88act)
func (cmsAdSeatService *CmsAdSeatService)GetCmsAdSeat(id uint,fields string ) (err error, obj business.CmsAdSeat) {
	err = global.DB.Where("id = ?", id).First(&obj).Error 
    if utils.IsEmpty(fields) {
        err = global.DB.Where("id = ?", id).First(&obj).Error 
        	} else {
        err = global.DB.Select(fields).Where("id = ?", id).First(&obj).Error  
	}

    obj.MapData = make(map[string]string)  
    return err, obj
}

// GetCmsAdSeatInfoList 分页获取CmsAdSeat记录
// Author [88act](https://github.com/88act)
func (cmsAdSeatService *CmsAdSeatService)GetCmsAdSeatInfoList(info businessReq.CmsAdSeatSearch, createdAtBetween []string,fields string) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    //修改 by ljd  增加查询排序 
    order := info.OrderKey
	desc := info.OrderDesc
    // 创建db
	db := global.DB.Model(&business.CmsAdSeat{})
    //var cmsAdSeats []business.CmsAdSeat
    var cmsAdSeats []business.CmsAdSeatMini

    //修改 by ljd  
    if info.ID > 0 {
		db = db.Where("`id` = ?", info.ID)
	}
	if createdAtBetween != nil && len(createdAtBetween) > 0 {
		db = db.Where("`created_at` BETWEEN ? AND ?", createdAtBetween[0], createdAtBetween[1])
	}

    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Status != nil {
        db = db.Where("`status` = ?",info.Status)
    }
	err = db.Count(&total).Error
	//err = db.Limit(limit).Offset(offset).Find(&cmsAdSeats).Error
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
      err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&cmsAdSeats).Error
    } else {
      err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&cmsAdSeats).Error
    }         
     //更新图片path
	for i, v := range cmsAdSeats {
	 v.MapData = make(map[string]string)
	  cmsAdSeats[i] = v
	}
	return err, cmsAdSeats, total
}
