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

type {{.StructName}}Service struct {
}

var once_{{.StructName}} sync.Once = sync.Once{}
var obj_{{.StructName}}Service *{{.StructName}}Service

//获取单例 
func Get{{.StructName}}Service() *{{.StructName}}Service {
	once_{{.StructName}}.Do(func() {
		obj_{{.StructName}}Service= new({{.StructName}}Service)
		//instanse.init()
	})
	return obj_{{.StructName}}Service
}



// Create{{.StructName}} 创建{{.StructName}}记录
// Author [88act](https://github.com/88act)
func (m *{{.StructName}}Service) Create{{.StructName}}({{.Abbreviation}} business.{{.StructName}}) (err error) {
	err = global.DB.Create(&{{.Abbreviation}}).Error
	return err
}

// Delete{{.StructName}} 删除{{.StructName}}记录
// Author [88act](https://github.com/88act)
func (m *{{.StructName}}Service)Delete{{.StructName}}({{.Abbreviation}} business.{{.StructName}}) (err error) {
	err = global.DB.Delete(&{{.Abbreviation}}).Error
	return err
}

// Delete{{.StructName}}ByIds 批量删除{{.StructName}}记录
// Author [88act](https://github.com/88act)
func (m *{{.StructName}}Service)Delete{{.StructName}}ByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]business.{{.StructName}}{},"id in ?",ids.Ids).Error
	return err
}

// Update{{.StructName}} 更新{{.StructName}}记录
// Author [88act](https://github.com/88act)
func (m *{{.StructName}}Service)Update{{.StructName}}({{.Abbreviation}} business.{{.StructName}}) (err error) {
	err = global.DB.Save(&{{.Abbreviation}}).Error
	return err
}

// Get{{.StructName}} 根据id获取{{.StructName}}记录
// Author [88act](https://github.com/88act)
func (m *{{.StructName}}Service)Get{{.StructName}}(id uint,fields string ) ( obj business.{{.StructName}},err error ) {
 
    if utils.IsEmpty(fields) {
        err = global.DB.Where("id = ?", id).First(&obj).Error 
        	} else {
        err = global.DB.Select(fields).Where("id = ?", id).First(&obj).Error  
	}

   //如果有图片image类型，更新图片path
    obj.MapData = make(map[string]string)
     {{- range .Fields}} 
         {{- if eq .FieldType "image"}} 
    if !utils.IsEmpty(obj.{{.FieldName}}) {
        _,obj.MapData[obj.{{.FieldName}}] = common.GetCommonFileService().GetPathByGuid(obj.{{.FieldName}})
    }     
        {{- end }}  
      {{- end }}  
    return  obj,err
}


// Get{{.StructName}}InfoList 分页获取{{.StructName}}记录
// Author [88act](https://github.com/88act)
func (m *{{.StructName}}Service)Get{{.StructName}}InfoList(info bizReq .{{.StructName}}Search, createdAtBetween []string,fields string) (list []business.{{.StructName}}Mini, total int64,err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    //修改 by ljd  增加查询排序 
    order := info.OrderKey
	desc := info.OrderDesc
    // 创建db
	db := global.DB.Model(&business.{{.StructName}}{})
    //var {{.Abbreviation}}s []business.{{.StructName}}
    var {{.Abbreviation}}s []business.{{.StructName}}Mini

    //修改 by ljd  
    if info.ID > 0 {
		db = db.Where("`id` = ?", info.ID)
	}
	if  len(createdAtBetween) >= 2 {
		db = db.Where("`created_at` BETWEEN ? AND ?", createdAtBetween[0], createdAtBetween[1])
	}

    // 如果有条件搜索 下方会自动创建搜索语句
        {{- range .Fields}}
            {{- if .FieldSearchType}}
                {{- if eq .FieldType "string" }}
    if info.{{.FieldName}} != "" {
        db = db.Where("`{{.ColumnName}}` {{.FieldSearchType}} ?",{{if eq .FieldSearchType "LIKE"}}"%"+ {{ end }}info.{{.FieldName}}{{if eq .FieldSearchType "LIKE"}}+"%"{{ end }})
    }
                {{- else if eq .FieldType "bool" }}
    if info.{{.FieldName}} != nil {
        db = db.Where("`{{.ColumnName}}` {{.FieldSearchType}} ?",{{if eq .FieldSearchType "LIKE"}}"%"+{{ end }}info.{{.FieldName}}{{if eq .FieldSearchType "LIKE"}}+"%"{{ end }})
    }
                {{- else if eq .FieldType "int" }}
    if info.{{.FieldName}} != nil {
        db = db.Where("`{{.ColumnName}}` {{.FieldSearchType}} ?",{{if eq .FieldSearchType "LIKE"}}"%"+{{ end }}info.{{.FieldName}}{{if eq .FieldSearchType "LIKE"}}+"%"{{ end }})
    }
                {{- else if eq .FieldType "float64" }}
    if info.{{.FieldName}} != nil {
        db = db.Where("`{{.ColumnName}}` {{.FieldSearchType}} ?",{{if eq .FieldSearchType "LIKE"}}"%"+{{ end }}info.{{.FieldName}}{{if eq .FieldSearchType "LIKE"}}+"%"{{ end }})
    }
                {{- else if eq .FieldType "time.Time" }}
    if info.{{.FieldName}} != nil {
          db = db.Where("`{{.ColumnName}}` {{.FieldSearchType}} > ?", info.{{.FieldName}})
    }
                {{- end }}
        {{- end }}
    {{- end }}
    
    err = db.Count(&total).Error
    if err != nil {
		return {{.Abbreviation}}s, 0, err
	}
	//err = db.Limit(limit).Offset(offset).Find(&{{.Abbreviation}}s).Error
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
      err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&{{.Abbreviation}}s).Error
    } else {
      err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&{{.Abbreviation}}s).Error
    }         
    //如果有图片image类型，更新图片path
	for i, v := range {{.Abbreviation}}s {
	 v.MapData = make(map[string]string)
     {{- range .Fields}} 
         {{- if eq .FieldType "image"}} 
        if !utils.IsEmpty(v.{{.FieldName}}) {
            _, v.MapData[v.{{.FieldName}}] = common.GetCommonFileService().GetPathByGuid(v.{{.FieldName}})
        }     
        {{- end}}  
      {{- end }}
	  {{.Abbreviation}}s[i] = v
	}
	return {{.Abbreviation}}s, total,err
}
 

// Get{{.StructName}}InfoListAll  分页获取{{.StructName}}记录 (全部字段)
// Author [88act](https://github.com/88act)
func (m *{{.StructName}}Service)Get{{.StructName}}InfoListAll(info bizReq .{{.StructName}}Search, createdAtBetween []string,fields string) (list []business.{{.StructName}}, total int64,err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    //修改 by ljd  增加查询排序 
    order := info.OrderKey
	desc := info.OrderDesc
    // 创建db
	db := global.DB.Model(&business.{{.StructName}}{})
    var {{.Abbreviation}}s []business.{{.StructName}}
    //var {{.Abbreviation}}s []business.{{.StructName}}Mini
   
    if info.ID > 0 {
		db = db.Where("`id` = ?", info.ID)
	}
	if len(createdAtBetween) >= 2 {
		db = db.Where("`created_at` BETWEEN ? AND ?", createdAtBetween[0], createdAtBetween[1])
	}

    // 如果有条件搜索 下方会自动创建搜索语句
        {{- range .Fields}}
            {{- if .FieldSearchType}}
                {{- if eq .FieldType "string" }}
    if info.{{.FieldName}} != "" {
        db = db.Where("`{{.ColumnName}}` {{.FieldSearchType}} ?",{{if eq .FieldSearchType "LIKE"}}"%"+ {{ end }}info.{{.FieldName}}{{if eq .FieldSearchType "LIKE"}}+"%"{{ end }})
    }
                {{- else if eq .FieldType "bool" }}
    if info.{{.FieldName}} != nil {
        db = db.Where("`{{.ColumnName}}` {{.FieldSearchType}} ?",{{if eq .FieldSearchType "LIKE"}}"%"+{{ end }}info.{{.FieldName}}{{if eq .FieldSearchType "LIKE"}}+"%"{{ end }})
    }
                {{- else if eq .FieldType "int" }}
    if info.{{.FieldName}} != nil {
        db = db.Where("`{{.ColumnName}}` {{.FieldSearchType}} ?",{{if eq .FieldSearchType "LIKE"}}"%"+{{ end }}info.{{.FieldName}}{{if eq .FieldSearchType "LIKE"}}+"%"{{ end }})
    }
                {{- else if eq .FieldType "float64" }}
    if info.{{.FieldName}} != nil {
        db = db.Where("`{{.ColumnName}}` {{.FieldSearchType}} ?",{{if eq .FieldSearchType "LIKE"}}"%"+{{ end }}info.{{.FieldName}}{{if eq .FieldSearchType "LIKE"}}+"%"{{ end }})
    }
                {{- else if eq .FieldType "time.Time" }}
    if info.{{.FieldName}} != nil {
          db = db.Where("`{{.ColumnName}}` {{.FieldSearchType}} > ?", info.{{.FieldName}})
    }
                {{- end }}
        {{- end }}
    {{- end }}

	err = db.Count(&total).Error
    if err != nil {
		return {{.Abbreviation}}s, 0, err
	}

	//err = db.Limit(limit).Offset(offset).Find(&{{.Abbreviation}}s).Error
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
      err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&{{.Abbreviation}}s).Error
    } else {
      err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&{{.Abbreviation}}s).Error
    }         
     //如果有图片image类型，更新图片path
	for i, v := range {{.Abbreviation}}s {
	 v.MapData = make(map[string]string)
     {{- range .Fields}} 
         {{- if eq .FieldType "image"}} 
        if !utils.IsEmpty(v.{{.FieldName}}) {
            _, v.MapData[v.{{.FieldName}}] = common.GetCommonFileService().GetPathByGuid(v.{{.FieldName}})
        }     
        {{- end}}  
      {{- end }}
	  {{.Abbreviation}}s[i] = v
	}
	return {{.Abbreviation}}s, total, err
}
