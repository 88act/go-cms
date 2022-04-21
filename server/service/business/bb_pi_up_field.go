package business

import (
	"go-cms/global"
	"go-cms/model/business" 
     bizReq "go-cms/model/business/request"
    "go-cms/model/common/request" 
    "go-cms/utils"
    "sync"
)

type BbPiUpFieldService struct {
}

var once_BbPiUpField sync.Once = sync.Once{}
var obj_BbPiUpFieldService *BbPiUpFieldService

//获取单例 
func GetBbPiUpFieldSev() *BbPiUpFieldService {
	once_BbPiUpField.Do(func() {
		obj_BbPiUpFieldService= new(BbPiUpFieldService)
		//instanse.init()
	})
	return obj_BbPiUpFieldService
}



// Create 创建BbPiUpField记录
// Author [88act](https://github.com/88act)
func (m *BbPiUpFieldService) Create(data business.BbPiUpField) (id uint,err error) {
	err = global.DB.Create(&data).Error
	if err != nil {
		return 0, err
	}
	return data.ID, err
}

// Delete 删除BbPiUpField记录
// Author [88act](https://github.com/88act)
func (m *BbPiUpFieldService)Delete(data business.BbPiUpField) (err error) {
	err = global.DB.Delete(&data).Error
	return err
}

// DeleteByIds 批量删除BbPiUpField记录
// Author [88act](https://github.com/88act)
func (m *BbPiUpFieldService)DeleteByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]business.BbPiUpField{},"id in ?",ids.Ids).Error
	return err
}

// Update  更新BbPiUpField记录
// Author [88act](https://github.com/88act)
func (m *BbPiUpFieldService)Update(data business.BbPiUpField) (err error) {
	err = global.DB.Save(&data).Error
	return err
}


// UpdateByMap  更新BbPiUpField记录 by Map
// values := map[string]interface{}{
// 	"status":0,
// 	"from": hash,
// }
// Author [88act](https://github.com/88act)
func (m *BbPiUpFieldService)UpdateByMap(data business.BbPiUpField, mapData map[string]interface{}) (err error) {
    err = global.DB.Model(&data).Updates(mapData).Error
	return err
}


// Get 根据id获取BbPiUpField记录
// Author [88act](https://github.com/88act)
func (m *BbPiUpFieldService)Get(id uint,fields string ) ( obj business.BbPiUpField,err error ) {
 
    if utils.IsEmpty(fields) {
        err = global.DB.Where("id = ?", id).First(&obj).Error 
    } else {
        err = global.DB.Select(fields).Where("id = ?", id).First(&obj).Error  
	}

   //如果有图片image类型，更新图片path
    obj.MapData = make(map[string]string)  
    return  obj,err
}


// GetList 分页获取BbPiUpField记录
// Author [88act](https://github.com/88act)
func (m *BbPiUpFieldService)GetList(info bizReq .BbPiUpFieldSearch, createdAtBetween []string,fields string) (list []business.BbPiUpFieldMini, total int64,err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    //修改 by ljd  增加查询排序 
    order := info.OrderKey
	desc := info.OrderDesc
    // 创建db
	db := global.DB.Model(&business.BbPiUpField{})
    //var bbPiUpFields []business.BbPiUpField
    var bbPiUpFields []business.BbPiUpFieldMini

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
    if info.TabName != "" {
        db = db.Where("`tab_name` = ?",info.TabName)
    }
    if info.TabField != "" {
        db = db.Where("`tab_field` = ?",info.TabField)
    }
    
    err = db.Count(&total).Error
    if err != nil {
		return bbPiUpFields, 0, err
	}
	//err = db.Limit(limit).Offset(offset).Find(&bbPiUpFields).Error
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
      err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&bbPiUpFields).Error
    } else {
      err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&bbPiUpFields).Error
    }         
    //如果有图片image类型，更新图片path
	for i, v := range bbPiUpFields {
	 v.MapData = make(map[string]string)
	  bbPiUpFields[i] = v
	}
	return bbPiUpFields, total,err
}
 

//GetListAll 分页获取BbPiUpField记录 (全部字段)
// Author [88act](https://github.com/88act)
func (m *BbPiUpFieldService)GetListAll(info bizReq .BbPiUpFieldSearch, createdAtBetween []string,fields string) (list []business.BbPiUpField, total int64,err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    //修改 by ljd  增加查询排序 
    order := info.OrderKey
	desc := info.OrderDesc
    // 创建db
	db := global.DB.Model(&business.BbPiUpField{})
    var bbPiUpFields []business.BbPiUpField
    //var bbPiUpFields []business.BbPiUpFieldMini
   
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
    if info.TabName != "" {
        db = db.Where("`tab_name` = ?",info.TabName)
    }
    if info.TabField != "" {
        db = db.Where("`tab_field` = ?",info.TabField)
    }

	err = db.Count(&total).Error
    if err != nil {
		return bbPiUpFields, 0, err
	}

	//err = db.Limit(limit).Offset(offset).Find(&bbPiUpFields).Error
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
      err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&bbPiUpFields).Error
    } else {
      err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&bbPiUpFields).Error
    }         
     //如果有图片image类型，更新图片path
	for i, v := range bbPiUpFields {
	 v.MapData = make(map[string]string)
	  bbPiUpFields[i] = v
	}
	return bbPiUpFields, total, err
}
