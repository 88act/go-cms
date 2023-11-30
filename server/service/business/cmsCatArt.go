package business

import (
    "context"
	"errors"
    "go-cms/mycache"
	"go-cms/global"
	"go-cms/model/business"  
    "go-cms/model/common/request"
     comSev "go-cms/service/common" 
    "go-cms/utils"
    "sync" 
    "gorm.io/gorm" 
	
)

type CmsCatArtService struct {
   	comSev.BaseService
}

var once_CmsCatArt sync.Once = sync.Once{}
var obj_CmsCatArtService *CmsCatArtService

//获取单例 
func GetCmsCatArtSev() *CmsCatArtService {
	once_CmsCatArt.Do(func() {
		obj_CmsCatArtService= new(CmsCatArtService)
		obj_CmsCatArtService.Db = global.DB
        //instanse.init()
	})
	return obj_CmsCatArtService
}



// Create 创建CmsCatArt记录
// Author 10512203@qq.com
func (m *CmsCatArtService) Create(ctx context.Context, data business.CmsCatArt) (id int64,err error) { 
	err = m.Db.WithContext(ctx).Create(&data).Error
	if err != nil {
		return 0, err
	}
	return data.Id, err
}

// Delete 删除CmsCatArt记录
// Author 10512203@qq.com
func (m *CmsCatArtService) Delete(ctx context.Context, data business.CmsCatArt) (err error) {
    m.DelCache("cms_cat_art", []int64{data.Id})
	err = m.Db.WithContext(ctx).Delete(&data).Error
	return err
}

// DeleteByIds 批量删除CmsCatArt记录
// Author 10512203@qq.com
func (m *CmsCatArtService) DeleteByIds(ctx context.Context, ids request.IdsReq) (err error) {
    m.DelCache("cms_cat_art", ids.Ids)
	err = m.Db.WithContext(ctx).Delete(&[]business.CmsCatArt{},"id in ?",ids.Ids).Error
	return err
}

// map删除 ，beUnscoped=true 硬删除
func (m *CmsCatArtService) DeleteByMap(ctx context.Context, mapData map[string]interface{}, beUnscoped bool) (err error) {
	if beUnscoped {
		err = m.Db.WithContext(ctx).Unscoped().Where(mapData).Delete(business.CmsCatArt{}).Error
	} else {
		err = m.Db.WithContext(ctx).Where(mapData).Delete(business.CmsCatArt{}).Error
	}
	return err
}


// Update  更新CmsCatArt记录
// Author 10512203@qq.com
func (m *CmsCatArtService) Update(ctx context.Context, data business.CmsCatArt) (err error) {
    m.DelCache("cms_cat_art", []int64{data.Id})
	err = m.Db.WithContext(ctx).Save(&data).Error
	return err
}


// UpdateByMap  更新CmsCatArt记录 by Map
// Author 10512203@qq.com
func (m *CmsCatArtService) UpdateByMap(ctx context.Context, data business.CmsCatArt, mapData map[string]interface{}) (err error) {
    m.DelCache("cms_cat_art", []int64{data.Id})
    err = m.Db.WithContext(ctx).Model(&data).Updates(mapData).Error
	return err
}


// UpdateExpr 字段自增/自减
// Author 10512203@qq.com
func (m *CmsCatArtService) UpdateExpr(ctx context.Context, data business.CmsCatArt,column string, num int) (err error) {
       err = m.Db.WithContext(ctx).Model(&data).UpdateColumn(column, gorm.Expr(column+" + ?", num)).Error
	return err
}
 

// Get 根据id获取CmsCatArt记录
// Author 10512203@qq.com
func (m *CmsCatArtService) Get(ctx context.Context, id int64,fields string ) (obj business.CmsCatArt,err error ) {
 
 	if id <= 0 {
		return obj, errors.New("参数id错误")
	}
    cacheKey := m.GetCacheKey("cms_cat_art", id)
	if cacheData, err := mycache.GetCache().Get(cacheKey); err == nil {
		obj = cacheData.(business.CmsCatArt)
		return obj, err
	}
    if utils.IsEmptyStr(fields) {
        err = m.Db.WithContext(ctx).Where("id = ?", id).First(&obj).Error 
    } else {
        err = m.Db.WithContext(ctx).Select(fields).Where("id = ?", id).First(&obj).Error  
	} 
    m.getFile(ctx,&obj)
    mycache.GetCache().Set(cacheKey, obj,obj.TableName())
    return  obj, err
}


// GetList 分页获取CmsCatArt记录
// Author 10512203@qq.com
func (m *CmsCatArtService) GetList(ctx context.Context, seq business .CmsCatArtSearch, fields string) (list []business.CmsCatArt, total int64,err error) {
	limit := seq.PageSize
	offset := seq.PageSize * (seq.Page - 1)
    //修改 by ljd  增加查询排序 
    order := seq.OrderKey
	desc := seq.OrderDesc
    // 创建db
	db := m.Db.WithContext(ctx).Model(&business.CmsCatArt{})
 

    //修改 by ljd  
    if seq.Id > 0 {
		db = db.Where("id = ?", seq.Id)
	}
    
	if !utils.IsEmpty(seq.CreatedAtBegin) && !utils.IsEmpty(seq.CreatedAtEnd) {
		db = db.Where("created_at BETWEEN ? AND ?", seq.CreatedAtBegin, seq.CreatedAtEnd)
	}

    // 如果有条件搜索 下方会自动创建搜索语句
    if seq.UserId != nil  {
        db = db.Where("user_id = ?",seq.UserId)
    }
    if seq.CatId != nil  {
        db = db.Where("cat_id = ?",seq.CatId)
    }
    if seq.ArtId != nil  {
        db = db.Where("art_id = ?",seq.ArtId)
    }
    if seq.Status != nil  {
        db = db.Where("status = ?",seq.Status)
    }
    
    err = db.Count(&total).Error
    if err != nil {
		return list, 0, err
	}
	//err = db.Limit(limit).Offset(offset).Find(&list).Error
    //修改 by ljd  增加查询排序 
     OrderStr := "id desc"
     if !utils.IsEmpty(order) { 
		if desc {
			OrderStr = order + " desc"
		} else {
			OrderStr = order
		} 
	}  
    if utils.IsEmptyStr(fields) {
      err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&list).Error
    } else {
      err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&list).Error
    }         
    m.getFileList(ctx,list)
	return list, total,err
}
 

//GetListAll 分页获取CmsCatArt记录 (全部字段)
// Author 10512203@qq.com
func (m *CmsCatArtService) GetListAll(ctx context.Context, seq business .CmsCatArtSearch, fields string) (list []business.CmsCatArt, total int64,err error) {
	return m.GetList(ctx, seq, "*")
}


// GetByMap 根据Map获取单一记录
// Author  [linjd] 10512203@qq.com
func (m *CmsCatArtService) GetByMap(ctx context.Context, mapData map[string]interface{}, fields string) (obj business.CmsCatArt, err error) {
    orderBy := "id desc"
	db := m.Db.WithContext(ctx).Model(&business.CmsCatArt{})
	if utils.IsEmptyStr(fields) {
		err = db.WithContext(ctx).Where(mapData).Order(orderBy).First(&obj).Error
	} else {
		err = db.WithContext(ctx).Select(fields).Order(orderBy).Where(mapData).First(&obj).Error
	}
	if err != nil {
		return obj, err
	}
	m.getFile(ctx,&obj)
	return obj, err
}

// GetByMap 根据Map获取列表
// Author 10512203@qq.com
func (m *CmsCatArtService) GetListByMap(ctx context.Context,mapData map[string]interface{}, fields string ,orderBy string) (list []business.CmsCatArt, err error) {
	if utils.IsEmptyStr(orderBy) {
		orderBy = "id desc"
	}
	db := m.Db.WithContext(ctx).Model(&business.CmsCatArt{})
	if utils.IsEmptyStr(fields) {
		err = db.Where(mapData).Order(orderBy).Find(&list).Error
	} else {
		err = db.Select(fields).Where(mapData).Order(orderBy).Find(&list).Error
	}
	if err != nil {
		return nil, err
	}	 
    m.getFileList(ctx,list)
	return list, err
} 

//	获取数量
//
// Author 10512203@qq.com
 func (m *CmsCatArtService) Count(ctx context.Context,mapDataWhere map[string]interface{})  (total int64, err error) {
	db := m.Db.WithContext(ctx).Model(&business.CmsCatArt{})
	err = db.Where(mapDataWhere).Count(&total).Error
	return total, err
}

//如果有文件类型，更新文件 path
func (m *CmsCatArtService) getFileList(ctx context.Context,list []business.CmsCatArt) {
    for i, v := range list {

    //    if !utils.IsEmpty(v.FileList) {			 
	// 		objList := strings.Split(v.FileList, ",")
	// 		for _, guid := range objList {
	// 			if !utils.IsEmptyStr(guid) {
	// 				fileObj := global.FileObj{Guid: guid}
	// 				fileObj.Path, _ = m.GetPathByGuid(ctx, guid)
	// 				v.FileObjList = append(v.FileObjList, fileObj)
	// 			}
	// 		}
	// 	}
		list[i] = v 
	} 
} 
 

//如果有文件类型，更新文件 path
func (m *CmsCatArtService) getFile(ctx context.Context,v *business.CmsCatArt) { 
    if v !=nil { 

    //   if !utils.IsEmpty(v.FileList) {
	// 		objList := strings.Split(v.FileList, ",")  
	// 		for _, guid := range objList {
	// 			if !utils.IsEmptyStr(guid) {
	// 				fileObj := global.FileObj{Guid: guid}
	// 				fileObj.Path, _ = m.GetPathByGuid(ctx, guid)
	// 				v.FileObjList = append(v.FileObjList, fileObj)
	// 			}
	// 		}
	// 	}
    }   
} 