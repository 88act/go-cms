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

type CmsVisitorService struct {
   	comSev.BaseService
}

var once_CmsVisitor sync.Once = sync.Once{}
var obj_CmsVisitorService *CmsVisitorService

//获取单例 
func GetCmsVisitorSev() *CmsVisitorService {
	once_CmsVisitor.Do(func() {
		obj_CmsVisitorService= new(CmsVisitorService)
		obj_CmsVisitorService.Db = global.DB
        //instanse.init()
	})
	return obj_CmsVisitorService
}



// Create 创建CmsVisitor记录
// Author 10512203@qq.com
func (m *CmsVisitorService) Create(ctx context.Context, data business.CmsVisitor) (id int64,err error) { 
	err = m.Db.WithContext(ctx).Create(&data).Error
	if err != nil {
		return 0, err
	}
	return data.Id, err
}

// Delete 删除CmsVisitor记录
// Author 10512203@qq.com
func (m *CmsVisitorService) Delete(ctx context.Context, data business.CmsVisitor) (err error) {
    m.DelCache("cms_visitor", []int64{data.Id})
	err = m.Db.WithContext(ctx).Delete(&data).Error
	return err
}

// DeleteByIds 批量删除CmsVisitor记录
// Author 10512203@qq.com
func (m *CmsVisitorService) DeleteByIds(ctx context.Context, ids request.IdsReq) (err error) {
    m.DelCache("cms_visitor", ids.Ids)
	err = m.Db.WithContext(ctx).Delete(&[]business.CmsVisitor{},"id in ?",ids.Ids).Error
	return err
}

// map删除 ，beUnscoped=true 硬删除
func (m *CmsVisitorService) DeleteByMap(ctx context.Context, mapData map[string]interface{}, beUnscoped bool) (err error) {
	if beUnscoped {
		err = m.Db.WithContext(ctx).Unscoped().Where(mapData).Delete(business.CmsVisitor{}).Error
	} else {
		err = m.Db.WithContext(ctx).Where(mapData).Delete(business.CmsVisitor{}).Error
	}
	return err
}


// Update  更新CmsVisitor记录
// Author 10512203@qq.com
func (m *CmsVisitorService) Update(ctx context.Context, data business.CmsVisitor) (err error) {
    m.DelCache("cms_visitor", []int64{data.Id})
	err = m.Db.WithContext(ctx).Save(&data).Error
	return err
}


// UpdateByMap  更新CmsVisitor记录 by Map
// Author 10512203@qq.com
func (m *CmsVisitorService) UpdateByMap(ctx context.Context, data business.CmsVisitor, mapData map[string]interface{}) (err error) {
    m.DelCache("cms_visitor", []int64{data.Id})
    err = m.Db.WithContext(ctx).Model(&data).Updates(mapData).Error
	return err
}


// UpdateExpr 字段自增/自减
// Author 10512203@qq.com
func (m *CmsVisitorService) UpdateExpr(ctx context.Context, data business.CmsVisitor,column string, num int) (err error) {
       err = m.Db.WithContext(ctx).Model(&data).UpdateColumn(column, gorm.Expr(column+" + ?", num)).Error
	return err
}
 

// Get 根据id获取CmsVisitor记录
// Author 10512203@qq.com
func (m *CmsVisitorService) Get(ctx context.Context, id int64,fields string ) (obj business.CmsVisitor,err error ) {
 
 	if id <= 0 {
		return obj, errors.New("参数id错误")
	}
    cacheKey := m.GetCacheKey("cms_visitor", id)
	if cacheData, err := mycache.GetCache().Get(cacheKey); err == nil {
		obj = cacheData.(business.CmsVisitor)
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


// GetList 分页获取CmsVisitor记录
// Author 10512203@qq.com
func (m *CmsVisitorService) GetList(ctx context.Context, seq business .CmsVisitorSearch, fields string) (list []business.CmsVisitor, total int64,err error) {
	limit := seq.PageSize
	offset := seq.PageSize * (seq.Page - 1)
    //修改 by ljd  增加查询排序 
    order := seq.OrderKey
	desc := seq.OrderDesc
    // 创建db
	db := m.Db.WithContext(ctx).Model(&business.CmsVisitor{})
 

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
    if seq.Type != nil  {
        db = db.Where("type = ?",seq.Type)
    }
    if seq.ObjId != nil  {
        db = db.Where("obj_id = ?",seq.ObjId)
    }
    if seq.PlatType != nil  {
        db = db.Where("plat_type = ?",seq.PlatType)
    }
    if seq.ClientIp != nil  {
        db = db.Where("client_ip = ?",seq.ClientIp)
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
 

//GetListAll 分页获取CmsVisitor记录 (全部字段)
// Author 10512203@qq.com
func (m *CmsVisitorService) GetListAll(ctx context.Context, seq business .CmsVisitorSearch, fields string) (list []business.CmsVisitor, total int64,err error) {
	return m.GetList(ctx, seq, "*")
}


// GetByMap 根据Map获取单一记录
// Author  [linjd] 10512203@qq.com
func (m *CmsVisitorService) GetByMap(ctx context.Context, mapData map[string]interface{}, fields string) (obj business.CmsVisitor, err error) {
    orderBy := "id desc"
	db := m.Db.WithContext(ctx).Model(&business.CmsVisitor{})
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
func (m *CmsVisitorService) GetListByMap(ctx context.Context,mapData map[string]interface{}, fields string ,orderBy string) (list []business.CmsVisitor, err error) {
	if utils.IsEmptyStr(orderBy) {
		orderBy = "id desc"
	}
	db := m.Db.WithContext(ctx).Model(&business.CmsVisitor{})
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
 func (m *CmsVisitorService) Count(ctx context.Context,mapDataWhere map[string]interface{})  (total int64, err error) {
	db := m.Db.WithContext(ctx).Model(&business.CmsVisitor{})
	err = db.Where(mapDataWhere).Count(&total).Error
	return total, err
}

//如果有文件类型，更新文件 path
func (m *CmsVisitorService) getFileList(ctx context.Context,list []business.CmsVisitor) {
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
func (m *CmsVisitorService) getFile(ctx context.Context,v *business.CmsVisitor) { 
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