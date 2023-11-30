package model

import (
	"context"
	"go-cms/common/mycache"
	"go-cms/common/utils"

	. "go-cms/common/baseModel"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

type {{.StructName}}Sev struct {
    BaseModelSev
}

func New{{.StructName}}Sev(db *gorm.DB, cache cache.ClusterConf) *{{.StructName}}Sev {
	instance := new({{.StructName}}Sev)
	instance.Db = db
	instance.Cache = cache
   	instance.CacheKeyPre = "{{.StructName}}"
	instance.CacheKeyListPre = "{{.StructName}}List"
	return instance
}


// Create 创建{{.StructName}}记录
// Author 10512203@qq.com
func (m *{{.StructName}}Sev) Create(ctx context.Context, data {{.StructName}}) (id int64,err error) {
   {{- range .Fields}}  
   {{- if eq .DataType "json" }} 
    if data.{{.FieldName}} == "" {
		data.{{.FieldName}} = "{}"
	} 
    {{- end }} 
    {{- end }} 
	err =  m.Db.WithContext(ctx).Create(&data).Error
    if err == nil {
		mycache.GetCache().DeleteKeyPre(m.CacheKeyListPre)
		return data.Id, err
	} else {
		return 0, err
	}
}

// Delete 删除{{.StructName}}记录
// Author 10512203@qq.com
func (m *{{.StructName}}Sev) Delete(ctx context.Context, data {{.StructName}}) (err error) { 
	err =  m.Db.WithContext(ctx).Delete(&data).Error
    if err == nil {
        m.DelCache(m.CacheKeyPre, []int64{data.Id})		 
	}
	return err
}

// DeleteByIds 批量删除{{.StructName}}记录
// Author 10512203@qq.com
func (m *{{.StructName}}Sev) DeleteByIds(ctx context.Context, ids IdsReq) (err error) {   
	err =  m.Db.WithContext(ctx).Delete(&[]{{.StructName}}{},"id in ?",ids.Ids).Error
    if err == nil {
		m.DelCache(m.CacheKeyPre, ids.Ids)		 
	}
	return err
}

// Update  更新{{.StructName}}记录
// Author 10512203@qq.com
func (m *{{.StructName}}Sev) Update(ctx context.Context, data {{.StructName}}) (err error) {    
	err =  m.Db.WithContext(ctx).Save(&data).Error
    if err == nil {
		m.DelCache(m.CacheKeyPre, []int64{data.Id}) 
	}
	return err
}


// UpdateByMap  更新{{.StructName}}记录 by Map
// Author 10512203@qq.com
func (m *{{.StructName}}Sev) UpdateByMap(ctx context.Context, data {{.StructName}}, mapData map[string]interface{}) (err error) {
    err =  m.Db.WithContext(ctx).Model(&data).Updates(mapData).Error
    if err == nil {
		m.DelCache(m.CacheKeyPre, []int64{data.Id}) 
	}
	return err
}


// UpdateExpr 字段自增/自减
// Author 10512203@qq.com
func (m *{{.StructName}}Sev) UpdateExpr(ctx context.Context, data {{.StructName}},column string, num int) (err error) {
       err = m.Db.WithContext(ctx).Model(&data).UpdateColumn(column, gorm.Expr(column+" + ?", num)).Error
	return err
}




// Get 根据id获取{{.StructName}}记录
// Author 10512203@qq.com
func (m *{{.StructName}}Sev) Get(ctx context.Context, id int64,fields string ) (obj {{.StructName}},err error ) {
 
 	if id <= 0 {
		return {{.StructName}}{}, nil
	}
    cacheKey := m.GetCacheKey(m.CacheKeyPre, id)
	if cacheData, err := mycache.GetCache().GetObj(cacheKey); err == nil {
		obj = cacheData.({{.StructName}})
		return obj, err
	}
    if utils.IsEmpty(fields) {
        err =  m.Db.WithContext(ctx).Where("id = ?", id).First(&obj).Error 
    } else {
        err =  m.Db.WithContext(ctx).Select(fields).Where("id = ?", id).First(&obj).Error  
	}     
  	if err == nil {
		m.getFile(ctx, &obj)
		mycache.GetCache().SetObj(cacheKey, obj, 0)
	}
    return  obj,err
}


// GetList 分页获取{{.StructName}}记录
// Author 10512203@qq.com
func (m *{{.StructName}}Sev) GetList(ctx context.Context, seq {{.StructName}}Search, fields string) (list []{{.StructName}}, total int64,err error) {
	limit := seq.PageSize
	offset := seq.PageSize * (seq.Page - 1)
    //修改 by ljd  增加查询排序 
    order := seq.OrderKey
	desc := seq.OrderDesc
    // 创建db
	db :=  m.Db.WithContext(ctx).Model(&{{.StructName}}{})  
	//强制状态为1 
  	db = db.Where("status = 1") 
    if seq.Id > 0 {
		db = db.Where("id = ?", seq.Id)
	}
	if seq.CreatedAtBegin != nil && seq.CreatedAtEnd != nil {
		db = db.Where("created_at BETWEEN ? AND ?", seq.CreatedAtBegin.Format("2006-01-02 15:04:05"),seq.CreatedAtEnd.Format("2006-01-02 15:04:05"))
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
    if seq.{{.FieldName}} != nil {
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
func (m *{{.StructName}}Sev) GetListAll(ctx context.Context, seq {{.StructName}}Search, fields string) (list []{{.StructName}}, total int64,err error) {
	return m.GetList(ctx, seq, "*")
}


// GetByMap 根据Map获取单一记录
// Author  [linjd] 10512203@qq.com
func (m *{{.StructName}}Sev) GetByMap(ctx context.Context, mapData map[string]interface{}, fields string) (obj {{.StructName}}, err error) {
    orderBy := "id desc"
	db := m.Db.WithContext(ctx).Model(&{{.StructName}}{})
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
func (m *{{.StructName}}Sev) GetListByMap(ctx context.Context,mapData map[string]interface{}, fields string ,orderBy string) (list []{{.StructName}}, err error) {
	if utils.IsEmptyStr(orderBy) {
		orderBy = "id desc"
	}
    db :=  m.Db.WithContext(ctx).Model(&{{.StructName}}{})
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
 func (m *{{.StructName}}Sev) Count(ctx context.Context,mapDataWhere map[string]interface{})  (total int64, err error) {
	db :=  m.Db.WithContext(ctx).Model(&{{.StructName}}{})
	err = db.Where(mapDataWhere).Count(&total).Error
	return total, err
}

//如果有文件类型，更新文件 path
func (m *{{.StructName}}Sev) getFileList(ctx context.Context,list []{{.StructName}}) {
    for i, v := range list { 
     {{- range .Fields}} 
         {{- if eq .FieldType "image"}} 
            if !utils.IsEmpty(v.{{.FieldName}}) {  
                fileObj{{.FieldName}} := FileObj{Guid: v.{{.FieldName}}}
                fileObj{{.FieldName}}.Path, _ = m.GetPathByGuid(ctx,v.{{.FieldName}})
                v.FileObjList = append(v.FileObjList, fileObj{{.FieldName}}) 
            } 
        {{- end}}  
      {{- end }}
	  list[i] = v
	} 
}

//如果有文件类型，更新文件 path
func (m *{{.StructName}}Sev) getFile(ctx context.Context,v *{{.StructName}}) { 
    if v !=nil { 
     {{- range .Fields}} 
         {{- if eq .FieldType "image"}} 
            if !utils.IsEmpty(v.{{.FieldName}}) {  
                fileObj{{.FieldName}} := FileObj{Guid: v.{{.FieldName}}}
                fileObj{{.FieldName}}.Path, _ = m.GetPathByGuid(ctx,v.{{.FieldName}})
                v.FileObjList = append(v.FileObjList, fileObj{{.FieldName}}) 
            }     
        {{- end}}  
      {{- end }} 
    }   
} 