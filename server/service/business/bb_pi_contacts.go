package business

import (
	"go-cms/global"
	"go-cms/model/business"
	bizReq "go-cms/model/business/request"
	"go-cms/model/common/request"
	"go-cms/utils"
	"sync"
)

type BbPiContactsService struct {
}

var once_BbPiContacts sync.Once = sync.Once{}
var obj_BbPiContactsService *BbPiContactsService

//获取单例 
func GetBbPiContactsSev() *BbPiContactsService {
	once_BbPiContacts.Do(func() {
		obj_BbPiContactsService= new(BbPiContactsService)
		//instanse.init()
	})
	return obj_BbPiContactsService
}



// Create 创建BbPiContacts记录
// Author [88act](https://github.com/88act)
func (m *BbPiContactsService) Create(data business.BbPiContacts) (id uint,err error) {
	err = global.DB.Create(&data).Error
	if err != nil {
		return 0, err
	}
	return data.ID, err
}

// Delete 删除BbPiContacts记录
// Author [88act](https://github.com/88act)
func (m *BbPiContactsService)Delete(data business.BbPiContacts) (err error) {
	err = global.DB.Delete(&data).Error
	return err
}

// DeleteByIds 批量删除BbPiContacts记录
// Author [88act](https://github.com/88act)
func (m *BbPiContactsService)DeleteByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]business.BbPiContacts{},"id in ?",ids.Ids).Error
	return err
}

// Update  更新BbPiContacts记录
// Author [88act](https://github.com/88act)
func (m *BbPiContactsService)Update(data business.BbPiContacts) (err error) {
	err = global.DB.Save(&data).Error
	return err
}


// UpdateByMap  更新BbPiContacts记录 by Map
// values := map[string]interface{}{
// 	"status":0,
// 	"from": hash,
// }
// Author [88act](https://github.com/88act)
func (m *BbPiContactsService)UpdateByMap(data business.BbPiContacts, mapData map[string]interface{}) (err error) {
    err = global.DB.Model(&data).Updates(mapData).Error
	return err
}


// Get 根据id获取BbPiContacts记录
// Author [88act](https://github.com/88act)
func (m *BbPiContactsService)Get(id uint,fields string ) ( obj business.BbPiContacts,err error ) {
 
    if utils.IsEmpty(fields) {
        err = global.DB.Where("id = ?", id).First(&obj).Error 
    } else {
        err = global.DB.Select(fields).Where("id = ?", id).First(&obj).Error  
	}

   //如果有图片image类型，更新图片path
    obj.MapData = make(map[string]string)  
    return  obj,err
}


// GetList 分页获取BbPiContacts记录
// Author [88act](https://github.com/88act)
func (m *BbPiContactsService)GetList(info bizReq .BbPiContactsSearch, createdAtBetween []string,fields string) (list []business.BbPiContactsMini, total int64,err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    //修改 by ljd  增加查询排序 
    order := info.OrderKey
	desc := info.OrderDesc
    // 创建db
	db := global.DB.Model(&business.BbPiContacts{})
    //var bbPiContactss []business.BbPiContacts
    var bbPiContactss []business.BbPiContactsMini

    //修改 by ljd  
    if info.ID > 0 {
		db = db.Where("`id` = ?", info.ID)
	}
	if  len(createdAtBetween) >= 2 {
		db = db.Where("`created_at` BETWEEN ? AND ?", createdAtBetween[0], createdAtBetween[1])
	}

    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Status != nil {
        db = db.Where("`status` = ?",info.Status)
    }
    
    err = db.Count(&total).Error
    if err != nil {
		return bbPiContactss, 0, err
	}
	//err = db.Limit(limit).Offset(offset).Find(&bbPiContactss).Error
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
      err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&bbPiContactss).Error
    } else {
      err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&bbPiContactss).Error
    }         
    //如果有图片image类型，更新图片path
	for i, v := range bbPiContactss {
	 v.MapData = make(map[string]string)
	  bbPiContactss[i] = v
	}
	return bbPiContactss, total,err
}
 

//GetListAll 分页获取BbPiContacts记录 (全部字段)
// Author [88act](https://github.com/88act)
func (m *BbPiContactsService)GetListAll(info bizReq .BbPiContactsSearch, createdAtBetween []string,fields string) (list []business.BbPiContacts, total int64,err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    //修改 by ljd  增加查询排序 
    order := info.OrderKey
	desc := info.OrderDesc
    // 创建db
	db := global.DB.Model(&business.BbPiContacts{})
    var bbPiContactss []business.BbPiContacts
    //var bbPiContactss []business.BbPiContactsMini
   
    if info.ID > 0 {
		db = db.Where("`id` = ?", info.ID)
	}
	if len(createdAtBetween) >= 2 {
		db = db.Where("`created_at` BETWEEN ? AND ?", createdAtBetween[0], createdAtBetween[1])
	}

    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Status != nil {
        db = db.Where("`status` = ?",info.Status)
    }

	err = db.Count(&total).Error
    if err != nil {
		return bbPiContactss, 0, err
	}

	//err = db.Limit(limit).Offset(offset).Find(&bbPiContactss).Error
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
      err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&bbPiContactss).Error
    } else {
      err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&bbPiContactss).Error
    }         
     //如果有图片image类型，更新图片path
	for i, v := range bbPiContactss {
	 v.MapData = make(map[string]string)
	  bbPiContactss[i] = v
	}
	return bbPiContactss, total, err
}
