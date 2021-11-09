package business

import (
	"github.com/88act/go-cms/server/global"
	"github.com/88act/go-cms/server/model/business"
	"github.com/88act/go-cms/server/model/common/request"
    businessReq "github.com/88act/go-cms/server/model/business/request"
    "github.com/88act/go-cms/server/utils"
)

type CmsArticleService struct {
}

// CreateCmsArticle 创建CmsArticle记录
// Author [88act](https://github.com/88act)
func (cmsArticleService *CmsArticleService) CreateCmsArticle(cmsArticle business.CmsArticle) (err error) {
	err = global.DB.Create(&cmsArticle).Error
	return err
}

// DeleteCmsArticle 删除CmsArticle记录
// Author [88act](https://github.com/88act)
func (cmsArticleService *CmsArticleService)DeleteCmsArticle(cmsArticle business.CmsArticle) (err error) {
	err = global.DB.Delete(&cmsArticle).Error
	return err
}

// DeleteCmsArticleByIds 批量删除CmsArticle记录
// Author [88act](https://github.com/88act)
func (cmsArticleService *CmsArticleService)DeleteCmsArticleByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]business.CmsArticle{},"id in ?",ids.Ids).Error
	return err
}

// UpdateCmsArticle 更新CmsArticle记录
// Author [88act](https://github.com/88act)
func (cmsArticleService *CmsArticleService)UpdateCmsArticle(cmsArticle business.CmsArticle) (err error) {
	err = global.DB.Save(&cmsArticle).Error
	return err
}

// GetCmsArticle 根据id获取CmsArticle记录
// Author [88act](https://github.com/88act)
func (cmsArticleService *CmsArticleService)GetCmsArticle(id uint,fields string ) (err error, obj business.CmsArticle) {
	err = global.DB.Where("id = ?", id).First(&obj).Error 
    if utils.IsEmpty(fields) {
        err = global.DB.Where("id = ?", id).First(&obj).Error 
        	} else {
        err = global.DB.Select(fields).Where("id = ?", id).First(&obj).Error  
	}

    obj.MapData = make(map[string]string) 
    if !utils.IsEmpty(obj.Thumb) {
        _,obj.MapData[obj.Thumb] = commFileService.GetPathByGuid(obj.Thumb)
    }  
    return err, obj
}

// GetCmsArticleInfoList 分页获取CmsArticle记录
// Author [88act](https://github.com/88act)
func (cmsArticleService *CmsArticleService)GetCmsArticleInfoList(info businessReq.CmsArticleSearch, createdAtBetween []string,fields string) (err error, list interface{}, total int64) {
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
	if createdAtBetween != nil && len(createdAtBetween) > 0 {
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
     //更新图片path
	for i, v := range cmsArticles {
	 v.MapData = make(map[string]string) 
        if !utils.IsEmpty(v.Thumb) {
            _, v.MapData[v.Thumb] = commFileService.GetPathByGuid(v.Thumb)
        }
	  cmsArticles[i] = v
	}
	return err, cmsArticles, total
}
