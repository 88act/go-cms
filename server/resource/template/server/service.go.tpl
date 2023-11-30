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
	"strings"
    "gorm.io/gorm" 
	
)

type {{.StructName}}Service struct {
   	comSev.BaseService
}

var once_{{.StructName}} sync.Once = sync.Once{}
var obj_{{.StructName}}Service *{{.StructName}}Service

//获取单例 
func Get{{.StructName}}Sev() *{{.StructName}}Service {
	once_{{.StructName}}.Do(func() {
		obj_{{.StructName}}Service= new({{.StructName}}Service)
		obj_{{.StructName}}Service.Db = global.DB
        //instanse.init()
	})
	return obj_{{.StructName}}Service
}



// Create 创建{{.StructName}}记录
// Author 10512203@qq.com
func (m *{{.StructName}}Service) Create(ctx context.Context, data business.{{.StructName}}) (id int64,err error) {
   {{- range .Fields}}  
   {{- if eq .DataType "json" }} 
    if data.{{.FieldName}} == "" {
		data.{{.FieldName}} = "{}"
	} 
    {{- end }} 
    {{- end }} 
	err = m.Db.WithContext(ctx).Create(&data).Error
	if err != nil {
		return 0, err
	}
	return data.Id, err
}

// Delete 删除{{.StructName}}记录
// Author 10512203@qq.com
func (m *{{.StructName}}Service) Delete(ctx context.Context, data business.{{.StructName}}) (err error) {
    m.DelCache("{{.TableName}}", []int64{data.Id})
	err = m.Db.WithContext(ctx).Delete(&data).Error
	return err
}

// DeleteByIds 批量删除{{.StructName}}记录
// Author 10512203@qq.com
func (m *{{.StructName}}Service) DeleteByIds(ctx context.Context, ids request.IdsReq) (err error) {
    m.DelCache("{{.TableName}}", ids.Ids)
	err = m.Db.WithContext(ctx).Delete(&[]business.{{.StructName}}{},"id in ?",ids.Ids).Error
	return err
}

// Update  更新{{.StructName}}记录
// Author 10512203@qq.com
func (m *{{.StructName}}Service) Update(ctx context.Context, data business.{{.StructName}}) (err error) {
    m.DelCache("{{.TableName}}", []int64{data.Id})
	err = m.Db.WithContext(ctx).Save(&data).Error
	return err
}


// UpdateByMap  更新{{.StructName}}记录 by Map
// Author 10512203@qq.com
func (m *{{.StructName}}Service) UpdateByMap(ctx context.Context, data business.{{.StructName}}, mapData map[string]interface{}) (err error) {
    m.DelCache("{{.TableName}}", []int64{data.Id})
    err = m.Db.WithContext(ctx).Model(&data).Updates(mapData).Error
	return err
}


// UpdateExpr 字段自增/自减
// Author 10512203@qq.com
func (m *{{.StructName}}Service) UpdateExpr(ctx context.Context, data business.{{.StructName}},column string, num int) (err error) {
       err = m.Db.WithContext(ctx).Model(&data).UpdateColumn(column, gorm.Expr(column+" + ?", num)).Error
	return err
}
 

// Get 根据id获取{{.StructName}}记录
// Author 10512203@qq.com
func (m *{{.StructName}}Service) Get(ctx context.Context, id int64,fields string ) (obj business.{{.StructName}},err error ) {
 
 	if id <= 0 {
		return obj, errors.New("参数id错误")
	}
    cacheKey := m.GetCacheKey("{{.TableName}}", id)
	if cacheData, err := mycache.GetCache().Get(cacheKey); err == nil {
		obj = cacheData.(business.{{.StructName}})
		return obj, err
	}
    if utils.IsEmpty(fields) {
        err = m.Db.WithContext(ctx).Where("id = ?", id).First(&obj).Error 
    } else {
        err = m.Db.WithContext(ctx).Select(fields).Where("id = ?", id).First(&obj).Error  
	} 
    m.getFile(ctx,&obj)
    mycache.GetCache().Set(cacheKey, obj,obj.TableName())
    return  obj, err
}


// GetList 分页获取{{.StructName}}记录
// Author 10512203@qq.com
func (m *{{.StructName}}Service) GetList(ctx context.Context, seq business .{{.StructName}}Search, fields string) (list []business.{{.StructName}}, total int64,err error) {
	limit := seq.PageSize
	offset := seq.PageSize * (seq.Page - 1)
    //修改 by ljd  增加查询排序 
    order := seq.OrderKey
	desc := seq.OrderDesc
    // 创建db
	db := m.Db.WithContext(ctx).Model(&business.{{.StructName}}{})
 

    //修改 by ljd  
    if seq.Id > 0 {
		db = db.Where("id = ?", seq.Id)
	}
    
	if !utils.IsEmpty(seq.CreatedAtBegin) && !utils.IsEmpty(seq.CreatedAtEnd) {
		db = db.Where("created_at BETWEEN ? AND ?", seq.CreatedAtBegin, seq.CreatedAtEnd)
	}

    // 如果有条件搜索 下方会自动创建搜索语句
        {{- range .Fields}}
            {{- if .FieldSearchType}}
                {{- if eq .FieldType "string" }}
    if seq.{{.FieldName}} != "" {
        db = db.Where("{{.ColumnName}} {{.FieldSearchType}} ?",{{if eq .FieldSearchType "LIKE"}}"%"+ {{ end }}seq.{{.FieldName}}{{if eq .FieldSearchType "LIKE"}}+"%"{{ end }})
    }
                {{- else if eq .FieldType "bool" }}
    if seq.{{.FieldName}} != nil {
        db = db.Where("{{.ColumnName}} {{.FieldSearchType}} ?",{{if eq .FieldSearchType "LIKE"}}"%"+{{ end }}seq.{{.FieldName}}{{if eq .FieldSearchType "LIKE"}}+"%"{{ end }})
    }
                {{- else if eq .FieldType "int" }}
    if seq.{{.FieldName}} != nil  {
        db = db.Where("{{.ColumnName}} {{.FieldSearchType}} ?",{{if eq .FieldSearchType "LIKE"}}"%"+{{ end }}seq.{{.FieldName}}{{if eq .FieldSearchType "LIKE"}}+"%"{{ end }})
    }
                {{- else if eq .FieldType "float64" }}
    if seq.{{.FieldName}} != nil {
        db = db.Where("{{.ColumnName}} {{.FieldSearchType}} ?",{{if eq .FieldSearchType "LIKE"}}"%"+{{ end }}seq.{{.FieldName}}{{if eq .FieldSearchType "LIKE"}}+"%"{{ end }})
    }
                {{- else if eq .FieldType "time.Time" }}
    if seq.{{.FieldName}} != nil {
          db = db.Where("{{.ColumnName}} {{.FieldSearchType}} > ?", seq.{{.FieldName}})
    }
                {{- end }}
        {{- end }}
    {{- end }}
    
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
    if utils.IsEmpty(fields) {
      err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&list).Error
    } else {
      err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&list).Error
    }         
    m.getFileList(ctx,list)
	return list, total,err
}
 

//GetListAll 分页获取{{.StructName}}记录 (全部字段)
// Author 10512203@qq.com
func (m *{{.StructName}}Service) GetListAll(ctx context.Context, seq business .{{.StructName}}Search, fields string) (list []business.{{.StructName}}, total int64,err error) {
	return m.GetList(ctx, seq, "*")
}


// GetByMap 根据Map获取单一记录
// Author  [linjd] 10512203@qq.com
func (m *{{.StructName}}Service) GetByMap(ctx context.Context, mapData map[string]interface{}, fields string) (obj business.{{.StructName}}, err error) {
    orderBy := "id desc"
	db := m.Db.WithContext(ctx).Model(&business.{{.StructName}}{})
	if utils.IsEmpty(fields) {
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
func (m *{{.StructName}}Service) GetListByMap(ctx context.Context,mapData map[string]interface{}, fields string ,orderBy string) (list []business.{{.StructName}}, err error) {
	if utils.IsEmptyStr(orderBy) {
		orderBy = "id desc"
	}
	db := m.Db.WithContext(ctx).Model(&business.{{.StructName}}{})
	if utils.IsEmpty(fields) {
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
 func (m *{{.StructName}}Service) Count(ctx context.Context,mapDataWhere map[string]interface{})  (total int64, err error) {
	db := m.Db.WithContext(ctx).Model(&business.{{.StructName}}{})
	err = db.Where(mapDataWhere).Count(&total).Error
	return total, err
}

//如果有文件类型，更新文件 path
func (m *{{.StructName}}Service) getFileList(ctx context.Context,list []business.{{.StructName}}) {
    for i, v := range list { 
     {{- range .Fields}} 
         {{- if eq .FieldType "image"}} 
            if !utils.IsEmpty(v.{{.FieldName}}) {  
                fileObj{{.FieldName}} := global.FileObj{Guid: v.{{.FieldName}}}
                fileObj{{.FieldName}}.Path, _ =  m.GetPathByGuid(ctx,v.{{.FieldName}})
                v.FileObjList = append(v.FileObjList, fileObj{{.FieldName}}) 
            } 
        {{- end}}  
      {{- end }}

       if !utils.IsEmpty(v.FileList) {			 
			objList := strings.Split(v.FileList, ",")
			for _, guid := range objList {
				if !utils.IsEmptyStr(guid) {
					fileObj := global.FileObj{Guid: guid}
					fileObj.Path, _ = m.GetPathByGuid(ctx, guid)
					v.FileObjList = append(v.FileObjList, fileObj)
				}
			}
		}
		list[i] = v 
	} 
} 
 

//如果有文件类型，更新文件 path
func (m *{{.StructName}}Service) getFile(ctx context.Context,v *business.{{.StructName}}) { 
    if v !=nil { 
     {{- range .Fields}} 
         {{- if eq .FieldType "image"}} 
            if !utils.IsEmpty(v.{{.FieldName}}) {  
                fileObj{{.FieldName}} := global.FileObj{Guid: v.{{.FieldName}}}
                fileObj{{.FieldName}}.Path, _ =  m.GetPathByGuid(ctx,v.{{.FieldName}})
                v.FileObjList = append(v.FileObjList, fileObj{{.FieldName}}) 
            }     
        {{- end}}  
      {{- end }} 

      if !utils.IsEmpty(v.FileList) {
			objList := strings.Split(v.FileList, ",")  
			for _, guid := range objList {
				if !utils.IsEmptyStr(guid) {
					fileObj := global.FileObj{Guid: guid}
					fileObj.Path, _ = m.GetPathByGuid(ctx, guid)
					v.FileObjList = append(v.FileObjList, fileObj)
				}
			}
		}
    }   
} 