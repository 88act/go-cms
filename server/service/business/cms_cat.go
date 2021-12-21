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

type CmsCatService struct {
}

var once_CmsCat sync.Once = sync.Once{}
var obj_CmsCatService *CmsCatService

//获取单例 
func GetCmsCatService() *CmsCatService {
	once_CmsCat.Do(func() {
		obj_CmsCatService= new(CmsCatService)
		//instanse.init()
	})
	return obj_CmsCatService
}



// CreateCmsCat 创建CmsCat记录
// Author [88act](https://github.com/88act)
func (m *CmsCatService) CreateCmsCat(cmsCat business.CmsCat) (err error) {
	err = global.DB.Create(&cmsCat).Error
	return err
}

// DeleteCmsCat 删除CmsCat记录
// Author [88act](https://github.com/88act)
func (m *CmsCatService)DeleteCmsCat(cmsCat business.CmsCat) (err error) {
	err = global.DB.Delete(&cmsCat).Error
	return err
}

// DeleteCmsCatByIds 批量删除CmsCat记录
// Author [88act](https://github.com/88act)
func (m *CmsCatService)DeleteCmsCatByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]business.CmsCat{},"id in ?",ids.Ids).Error
	return err
}

// UpdateCmsCat 更新CmsCat记录
// Author [88act](https://github.com/88act)
func (m *CmsCatService)UpdateCmsCat(cmsCat business.CmsCat) (err error) {
	err = global.DB.Save(&cmsCat).Error
	return err
}

// GetCmsCat 根据id获取CmsCat记录
// Author [88act](https://github.com/88act)
func (m *CmsCatService)GetCmsCat(id uint,fields string ) ( obj business.CmsCat,err error ) {
 
    if utils.IsEmpty(fields) {
        err = global.DB.Where("id = ?", id).First(&obj).Error 
        	} else {
        err = global.DB.Select(fields).Where("id = ?", id).First(&obj).Error  
	}

   //如果有图片image类型，更新图片path
    obj.MapData = make(map[string]string) 
    if !utils.IsEmpty(obj.Thumb) {
        _,obj.MapData[obj.Thumb] = common.GetCommonFileService().GetPathByGuid(obj.Thumb)
    }  
    return  obj,err
}


// GetCmsCatInfoList 分页获取CmsCat记录
// Author [88act](https://github.com/88act)
func (m *CmsCatService)GetCmsCatInfoList(info bizReq .CmsCatSearch, createdAtBetween []string,fields string) (list []business.CmsCatMini, total int64,err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    //修改 by ljd  增加查询排序 
    order := info.OrderKey
	desc := info.OrderDesc
    // 创建db
	db := global.DB.Model(&business.CmsCat{})
    //var cmsCats []business.CmsCat
    var cmsCats []business.CmsCatMini

    //修改 by ljd  
    if info.ID > 0 {
		db = db.Where("`id` = ?", info.ID)
	}
	if  len(createdAtBetween) >= 2 {
		db = db.Where("`created_at` BETWEEN ? AND ?", createdAtBetween[0], createdAtBetween[1])
	}

    // 如果有条件搜索 下方会自动创建搜索语句
    if info.BeSys != nil {
        db = db.Where("`be_sys` = ?",info.BeSys)
    }
    if info.GroupId != nil {
        db = db.Where("`group_id` = ?",info.GroupId)
    }
    if info.MediaType != nil {
        db = db.Where("`media_type` = ?",info.MediaType)
    }
    if info.Name != "" {
        db = db.Where("`name` LIKE ?","%"+ info.Name+"%")
    }
    if info.Desc != "" {
        db = db.Where("`desc` LIKE ?","%"+ info.Desc+"%")
    }
    if info.Keywords != "" {
        db = db.Where("`keywords` LIKE ?","%"+ info.Keywords+"%")
    }
    if info.Status != nil {
        db = db.Where("`status` = ?",info.Status)
    }
    
    err = db.Count(&total).Error
    if err != nil {
		return cmsCats, 0, err
	}
	//err = db.Limit(limit).Offset(offset).Find(&cmsCats).Error
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
      err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&cmsCats).Error
    } else {
      err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&cmsCats).Error
    }         
    //如果有图片image类型，更新图片path
	for i, v := range cmsCats {
	 v.MapData = make(map[string]string) 
        if !utils.IsEmpty(v.Thumb) {
            _, v.MapData[v.Thumb] = common.GetCommonFileService().GetPathByGuid(v.Thumb)
        }
	  cmsCats[i] = v
	}
	return cmsCats, total,err
}
 

// GetCmsCatInfoListAll  分页获取CmsCat记录 (全部字段)
// Author [88act](https://github.com/88act)
func (m *CmsCatService)GetCmsCatInfoListAll(info bizReq .CmsCatSearch, createdAtBetween []string,fields string) (list []business.CmsCat, total int64,err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    //修改 by ljd  增加查询排序 
    order := info.OrderKey
	desc := info.OrderDesc
    // 创建db
	db := global.DB.Model(&business.CmsCat{})
    var cmsCats []business.CmsCat
    //var cmsCats []business.CmsCatMini
   
    if info.ID > 0 {
		db = db.Where("`id` = ?", info.ID)
	}
	if len(createdAtBetween) >= 2 {
		db = db.Where("`created_at` BETWEEN ? AND ?", createdAtBetween[0], createdAtBetween[1])
	}

    // 如果有条件搜索 下方会自动创建搜索语句
    if info.BeSys != nil {
        db = db.Where("`be_sys` = ?",info.BeSys)
    }
    if info.GroupId != nil {
        db = db.Where("`group_id` = ?",info.GroupId)
    }
    if info.MediaType != nil {
        db = db.Where("`media_type` = ?",info.MediaType)
    }
    if info.Name != "" {
        db = db.Where("`name` LIKE ?","%"+ info.Name+"%")
    }
    if info.Desc != "" {
        db = db.Where("`desc` LIKE ?","%"+ info.Desc+"%")
    }
    if info.Keywords != "" {
        db = db.Where("`keywords` LIKE ?","%"+ info.Keywords+"%")
    }
    if info.Status != nil {
        db = db.Where("`status` = ?",info.Status)
    }

	err = db.Count(&total).Error
    if err != nil {
		return cmsCats, 0, err
	}

	//err = db.Limit(limit).Offset(offset).Find(&cmsCats).Error
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
      err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&cmsCats).Error
    } else {
      err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&cmsCats).Error
    }         
     //如果有图片image类型，更新图片path
	for i, v := range cmsCats {
	 v.MapData = make(map[string]string) 
        if !utils.IsEmpty(v.Thumb) {
            _, v.MapData[v.Thumb] = common.GetCommonFileService().GetPathByGuid(v.Thumb)
        }
	  cmsCats[i] = v
	}
	return cmsCats, total, err
}
