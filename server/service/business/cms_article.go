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

type CmsArticleService struct {
}

var once_CmsArticle sync.Once = sync.Once{}
var obj_CmsArticleService *CmsArticleService

//获取单例 
func GetCmsArticleService() *CmsArticleService {
	once_CmsArticle.Do(func() {
		obj_CmsArticleService= new(CmsArticleService)
		//instanse.init()
	})
	return obj_CmsArticleService
}



// CreateCmsArticle 创建CmsArticle记录
// Author [88act](https://github.com/88act)
func (m *CmsArticleService) CreateCmsArticle(cmsArticle business.CmsArticle) (err error) {
	err = global.DB.Create(&cmsArticle).Error
	return err
}

// DeleteCmsArticle 删除CmsArticle记录
// Author [88act](https://github.com/88act)
func (m *CmsArticleService)DeleteCmsArticle(cmsArticle business.CmsArticle) (err error) {
	err = global.DB.Delete(&cmsArticle).Error
	return err
}

// DeleteCmsArticleByIds 批量删除CmsArticle记录
// Author [88act](https://github.com/88act)
func (m *CmsArticleService)DeleteCmsArticleByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]business.CmsArticle{},"id in ?",ids.Ids).Error
	return err
}

// UpdateCmsArticle 更新CmsArticle记录
// Author [88act](https://github.com/88act)
func (m *CmsArticleService)UpdateCmsArticle(cmsArticle business.CmsArticle) (err error) {
	err = global.DB.Save(&cmsArticle).Error
	return err
}

// GetCmsArticle 根据id获取CmsArticle记录
// Author [88act](https://github.com/88act)
func (m *CmsArticleService)GetCmsArticle(id uint,fields string ) ( obj business.CmsArticle,err error ) {
 
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


// GetCmsArticleInfoList 分页获取CmsArticle记录
// Author [88act](https://github.com/88act)
func (m *CmsArticleService)GetCmsArticleInfoList(info bizReq .CmsArticleSearch, createdAtBetween []string,fields string) (list []business.CmsArticleMini, total int64,err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    //修改 by ljd  增加查询排序 
    order := info.OrderKey
	desc := info.OrderDesc
    // 创建db
	db := global.DB.Model(&business.CmsArticle{})
    //var cmsArticles []business.CmsArticle
    var cmsArticles []business.CmsArticleMini

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
    if info.Name != "" {
        db = db.Where("`name` LIKE ?","%"+ info.Name+"%")
    }
    if info.Sketch != "" {
        db = db.Where("`sketch` LIKE ?","%"+ info.Sketch+"%")
    }
    if info.Detail != "" {
        db = db.Where("`detail` LIKE ?","%"+ info.Detail+"%")
    }
    if info.Status != nil {
        db = db.Where("`status` = ?",info.Status)
    }
    
    err = db.Count(&total).Error
    if err != nil {
		return cmsArticles, 0, err
	}
	//err = db.Limit(limit).Offset(offset).Find(&cmsArticles).Error
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
      err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&cmsArticles).Error
    } else {
      err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&cmsArticles).Error
    }         
    //如果有图片image类型，更新图片path
	for i, v := range cmsArticles {
	 v.MapData = make(map[string]string) 
        if !utils.IsEmpty(v.Thumb) {
            _, v.MapData[v.Thumb] = common.GetCommonFileService().GetPathByGuid(v.Thumb)
        }
	  cmsArticles[i] = v
	}
	return cmsArticles, total,err
}
 

// GetCmsArticleInfoListAll  分页获取CmsArticle记录 (全部字段)
// Author [88act](https://github.com/88act)
func (m *CmsArticleService)GetCmsArticleInfoListAll(info bizReq .CmsArticleSearch, createdAtBetween []string,fields string) (list []business.CmsArticle, total int64,err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    //修改 by ljd  增加查询排序 
    order := info.OrderKey
	desc := info.OrderDesc
    // 创建db
	db := global.DB.Model(&business.CmsArticle{})
    var cmsArticles []business.CmsArticle
    //var cmsArticles []business.CmsArticleMini
   
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
    if info.Name != "" {
        db = db.Where("`name` LIKE ?","%"+ info.Name+"%")
    }
    if info.Sketch != "" {
        db = db.Where("`sketch` LIKE ?","%"+ info.Sketch+"%")
    }
    if info.Detail != "" {
        db = db.Where("`detail` LIKE ?","%"+ info.Detail+"%")
    }
    if info.Status != nil {
        db = db.Where("`status` = ?",info.Status)
    }

	err = db.Count(&total).Error
    if err != nil {
		return cmsArticles, 0, err
	}

	//err = db.Limit(limit).Offset(offset).Find(&cmsArticles).Error
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
      err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&cmsArticles).Error
    } else {
      err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&cmsArticles).Error
    }         
     //如果有图片image类型，更新图片path
	for i, v := range cmsArticles {
	 v.MapData = make(map[string]string) 
        if !utils.IsEmpty(v.Thumb) {
            _, v.MapData[v.Thumb] = common.GetCommonFileService().GetPathByGuid(v.Thumb)
        }
	  cmsArticles[i] = v
	}
	return cmsArticles, total, err
}
