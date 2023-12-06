package model

import (
	"context"
	"go-cms/common/mycache"
	"go-cms/common/utils"

	. "go-cms/common/baseModel"

	"gorm.io/gorm"
)

type CmsArtSev struct {
	BaseModelSev
}

func NewCmsArtSev(db *gorm.DB) *CmsArtSev {
	instance := new(CmsArtSev)
	instance.Db = db
	instance.CacheKeyPre = "CmsArt"
	instance.CacheKeyListPre = "CmsArtList"
	return instance
}

// Create 创建CmsArt记录
// Author 10512203@qq.com
func (m *CmsArtSev) Create(ctx context.Context, data CmsArt) (id int64, err error) {
	err = m.Db.WithContext(ctx).Create(&data).Error
	if err == nil {
		mycache.GetCache().DeleteKeyPre(m.CacheKeyListPre)
		return data.Id, err
	} else {
		return 0, err
	}
}

// GetActList
// Author 10512203@qq.com
func (m *CmsArtSev) GetActList(ctx context.Context, seq CmsArtSearch) (list []CmsArt, total int64, err error) {
	limit := seq.PageSize
	offset := seq.PageSize * (seq.Page - 1)
	// order := seq.OrderKey
	// desc := seq.OrderDesc
	db := m.Db.WithContext(ctx).Table("cms_cat_art a")
	db = db.Joins("left join cms_art b on a.art_id = b.id")
	db = db.Where("ISNULL(a.deleted_at)")
	db = db.Where("ISNULL(b.deleted_at)")
	db = db.Where("a.cat_id=?", *seq.CatId)
	fields := "b.*"
	OrderStr := "a.sort desc"

	err = db.Count(&total).Error
	if err != nil {
		return list, 0, err
	}
	err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Scan(&list).Error
	m.getFileList(ctx, list)
	return list, total, err
}

// Delete 删除CmsArt记录
// Author 10512203@qq.com
func (m *CmsArtSev) Delete(ctx context.Context, data CmsArt) (err error) {
	err = m.Db.WithContext(ctx).Delete(&data).Error
	if err == nil {
		m.DelCache(m.CacheKeyPre, []int64{data.Id})
	}
	return err
}

// DeleteByIds 批量删除CmsArt记录
// Author 10512203@qq.com
func (m *CmsArtSev) DeleteByIds(ctx context.Context, ids IdsReq) (err error) {
	err = m.Db.WithContext(ctx).Delete(&[]CmsArt{}, "id in ?", ids.Ids).Error
	if err == nil {
		m.DelCache(m.CacheKeyPre, ids.Ids)
	}
	return err
}

// Update  更新CmsArt记录
// Author 10512203@qq.com
func (m *CmsArtSev) Update(ctx context.Context, data CmsArt) (err error) {
	err = m.Db.WithContext(ctx).Save(&data).Error
	if err == nil {
		m.DelCache(m.CacheKeyPre, []int64{data.Id})
	}
	return err
}

// UpdateByMap  更新CmsArt记录 by Map
// Author 10512203@qq.com
func (m *CmsArtSev) UpdateByMap(ctx context.Context, data CmsArt, mapData map[string]interface{}) (err error) {
	err = m.Db.WithContext(ctx).Model(&data).Updates(mapData).Error
	if err == nil {
		m.DelCache(m.CacheKeyPre, []int64{data.Id})
	}
	return err
}

// UpdateExpr 字段自增/自减
// Author 10512203@qq.com
func (m *CmsArtSev) UpdateExpr(ctx context.Context, data CmsArt, column string, num int) (err error) {
	err = m.Db.WithContext(ctx).Model(&data).UpdateColumn(column, gorm.Expr(column+" + ?", num)).Error
	return err
}

// Get 根据id获取CmsArt记录
// Author 10512203@qq.com
func (m *CmsArtSev) Get(ctx context.Context, id int64, fields string) (obj CmsArt, err error) {
	if id <= 0 {
		return CmsArt{}, nil
	}
	cacheKey := m.GetCacheKey(m.CacheKeyPre, id)
	if cacheData, err := mycache.GetCache().GetObj(cacheKey); err == nil {
		obj = cacheData.(CmsArt)
		return obj, err
	}
	if utils.IsEmpty(fields) {
		err = m.Db.WithContext(ctx).Where("id = ?", id).First(&obj).Error
	} else {
		err = m.Db.WithContext(ctx).Select(fields).Where("id = ?", id).First(&obj).Error
	}
	if err == nil {
		m.getFile(ctx, &obj)
		mycache.GetCache().SetObj(cacheKey, obj, 0)
	}
	return obj, err
}

// GetList 分页获取CmsArt记录
// Author 10512203@qq.com
func (m *CmsArtSev) GetList(ctx context.Context, seq CmsArtSearch, fields string) (list []CmsArt, total int64, err error) {
	limit := seq.PageSize
	offset := seq.PageSize * (seq.Page - 1)

	order := seq.OrderKey
	desc := seq.OrderDesc
	// 创建db
	db := m.Db.WithContext(ctx).Model(&CmsArt{})
	//强制状态为1
	db = db.Where("status = 1")
	if seq.Id > 0 {
		db = db.Where("id = ?", seq.Id)
	}
	if seq.CreatedAtBegin != nil && seq.CreatedAtEnd != nil {
		db = db.Where("created_at BETWEEN ? AND ?", seq.CreatedAtBegin.Format("2006-01-02 15:04:05"), seq.CreatedAtEnd.Format("2006-01-02 15:04:05"))
	}

	// 如果有条件搜索 下方会自动创建搜索语句
	if seq.Pid != nil {
		db = db.Where("pid = ?", seq.Pid)
	}
	if seq.UserId != nil {
		db = db.Where("user_id = ?", seq.UserId)
	}
	if seq.CatId != nil {
		db = db.Where("cat_id = ?", seq.CatId)
	}
	if seq.CatIdSys != nil {
		db = db.Where("cat_id_sys = ?", seq.CatIdSys)
	}
	if seq.Type != nil {
		db = db.Where("type = ?", seq.Type)
	}
	if seq.Title != "" {
		db = db.Where("title LIKE ?", "%"+seq.Title+"%")
	}
	if seq.Desc != "" {
		db = db.Where("desc LIKE ?", "%"+seq.Desc+"%")
	}
	if seq.TagList != "" {
		db = db.Where("tag_list = ?", seq.TagList)
	}
	if seq.Source != "" {
		db = db.Where("source = ?", seq.Source)
	}
	if seq.Link != "" {
		db = db.Where("link = ?", seq.Link)
	}
	if seq.BeTop != nil {
		db = db.Where("be_top = ?", seq.BeTop)
	}
	if seq.TotalWhole != nil {
		db = db.Where("total_whole = ?", seq.TotalWhole)
	}
	if seq.TotalShare != nil {
		db = db.Where("total_share = ?", seq.TotalShare)
	}
	if seq.TotalFav != nil {
		db = db.Where("total_fav = ?", seq.TotalFav)
	}
	if seq.TotalDiscuss != nil {
		db = db.Where("total_discuss = ?", seq.TotalDiscuss)
	}
	if seq.TotalClick != nil {
		db = db.Where("total_click = ?", seq.TotalClick)
	}
	if seq.TotalStar != nil {
		db = db.Where("total_star = ?", seq.TotalStar)
	}
	if seq.TotalGood != nil {
		db = db.Where("total_good = ?", seq.TotalGood)
	}
	if seq.TotalPoor != nil {
		db = db.Where("total_poor = ?", seq.TotalPoor)
	}
	if seq.Status != nil {
		db = db.Where("status = ?", seq.Status)
	}
	if seq.VerifyMsg != "" {
		db = db.Where("verify_msg = ?", seq.VerifyMsg)
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
	if utils.IsEmpty(fields) {
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&list).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&list).Error
	}
	m.getFileList(ctx, list)
	return list, total, err
}

// GetListAll 分页获取CmsArt记录 (全部字段)
// Author 10512203@qq.com
func (m *CmsArtSev) GetListAll(ctx context.Context, seq CmsArtSearch, fields string) (list []CmsArt, total int64, err error) {
	return m.GetList(ctx, seq, "*")
}

// GetByMap 根据Map获取单一记录
// Author  [linjd] 10512203@qq.com
func (m *CmsArtSev) GetByMap(ctx context.Context, mapData map[string]interface{}, fields string) (obj CmsArt, err error) {
	orderBy := "id desc"
	db := m.Db.WithContext(ctx).Model(&CmsArt{})
	if utils.IsEmpty(fields) {
		err = db.WithContext(ctx).Where(mapData).Order(orderBy).First(&obj).Error
	} else {
		err = db.WithContext(ctx).Select(fields).Order(orderBy).Where(mapData).First(&obj).Error
	}
	if err != nil {
		return obj, err
	}
	m.getFile(ctx, &obj)
	return obj, err

}

// GetByMap 根据Map获取列表
// Author 10512203@qq.com
func (m *CmsArtSev) GetListByMap(ctx context.Context, mapData map[string]interface{}, fields string, orderBy string) (list []CmsArt, err error) {
	if utils.IsEmptyStr(orderBy) {
		orderBy = "id desc"
	}
	db := m.Db.WithContext(ctx).Model(&CmsArt{})
	if utils.IsEmpty(fields) {
		err = db.Where(mapData).Order(orderBy).Find(&list).Error
	} else {
		err = db.Select(fields).Where(mapData).Order(orderBy).Find(&list).Error
	}
	if err != nil {
		return nil, err
	}
	m.getFileList(ctx, list)
	return list, err
}

//	获取数量
//
// Author 10512203@qq.com
func (m *CmsArtSev) Count(ctx context.Context, mapDataWhere map[string]interface{}) (total int64, err error) {
	db := m.Db.WithContext(ctx).Model(&CmsArt{})
	err = db.Where(mapDataWhere).Count(&total).Error
	return total, err
}

// 如果有文件类型，更新文件 path
func (m *CmsArtSev) getFileList(ctx context.Context, list []CmsArt) {
	for i, v := range list {
		if !utils.IsEmpty(v.Image) {
			v.Image, _ = m.GetPathByGuid(ctx, v.Image)
		}
		list[i] = v
	}
}

// 如果有文件类型，更新文件 path
func (m *CmsArtSev) getFile(ctx context.Context, v *CmsArt) {
	if v != nil {
		if !utils.IsEmpty(v.Image) {
			v.Image, _ = m.GetPathByGuid(ctx, v.Image)
		}
	}
}
