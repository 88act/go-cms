package business

import (
	"context"
	"errors"
	"go-cms/global"
	"go-cms/model/business"
	"go-cms/model/common/request"
	"go-cms/mycache"
	comSev "go-cms/service/common"
	"go-cms/utils"
	"sync"

	"gorm.io/gorm"
)

type CmsArtService struct {
	comSev.BaseService
}

var once_CmsArt sync.Once = sync.Once{}
var obj_CmsArtService *CmsArtService

// 获取单例
func GetCmsArtSev() *CmsArtService {
	once_CmsArt.Do(func() {
		obj_CmsArtService = new(CmsArtService)
		obj_CmsArtService.Db = global.DB
		//instanse.init()
	})
	return obj_CmsArtService
}

// Create 创建CmsArt记录
// Author 10512203@qq.com
func (m *CmsArtService) Create(ctx context.Context, data business.CmsArt) (id int64, err error) {
	err = m.Db.WithContext(ctx).Create(&data).Error
	if err != nil {
		return 0, err
	}
	return data.Id, err
}

// Delete 删除CmsArt记录
// Author 10512203@qq.com
func (m *CmsArtService) Delete(ctx context.Context, data business.CmsArt) (err error) {
	m.DelCache("cms_art", []int64{data.Id})
	err = m.Db.WithContext(ctx).Delete(&data).Error
	return err
}

// DeleteByIds 批量删除CmsArt记录
// Author 10512203@qq.com
func (m *CmsArtService) DeleteByIds(ctx context.Context, ids request.IdsReq) (err error) {
	m.DelCache("cms_art", ids.Ids)
	err = m.Db.WithContext(ctx).Delete(&[]business.CmsArt{}, "id in ?", ids.Ids).Error
	return err
}

// map删除 ，beUnscoped=true 硬删除
func (m *CmsArtService) DeleteByMap(ctx context.Context, mapData map[string]interface{}, beUnscoped bool) (err error) {
	if beUnscoped {
		err = m.Db.WithContext(ctx).Unscoped().Where(mapData).Delete(business.CmsArt{}).Error
	} else {
		err = m.Db.WithContext(ctx).Where(mapData).Delete(business.CmsArt{}).Error
	}
	return err
}

// Update  更新CmsArt记录
// Author 10512203@qq.com
func (m *CmsArtService) Update(ctx context.Context, data business.CmsArt) (err error) {
	m.DelCache("cms_art", []int64{data.Id})
	err = m.Db.WithContext(ctx).Save(&data).Error
	return err
}

// UpdateByMap  更新CmsArt记录 by Map
// Author 10512203@qq.com
func (m *CmsArtService) UpdateByMap(ctx context.Context, data business.CmsArt, mapData map[string]interface{}) (err error) {
	m.DelCache("cms_art", []int64{data.Id})
	err = m.Db.WithContext(ctx).Model(&data).Updates(mapData).Error
	return err
}

// UpdateExpr 字段自增/自减
// Author 10512203@qq.com
func (m *CmsArtService) UpdateExpr(ctx context.Context, data business.CmsArt, column string, num int) (err error) {
	err = m.Db.WithContext(ctx).Model(&data).UpdateColumn(column, gorm.Expr(column+" + ?", num)).Error
	return err
}

// Get 根据id获取CmsArt记录
// Author 10512203@qq.com
func (m *CmsArtService) Get(ctx context.Context, id int64, fields string) (obj business.CmsArt, err error) {

	if id <= 0 {
		return obj, errors.New("参数id错误")
	}
	cacheKey := m.GetCacheKey("cms_art", id)
	if cacheData, err := mycache.GetCache().Get(cacheKey); err == nil {
		obj = cacheData.(business.CmsArt)
		return obj, err
	}
	if utils.IsEmptyStr(fields) {
		err = m.Db.WithContext(ctx).Where("id = ?", id).First(&obj).Error
	} else {
		err = m.Db.WithContext(ctx).Select(fields).Where("id = ?", id).First(&obj).Error
	}
	m.getFile(ctx, &obj)
	mycache.GetCache().Set(cacheKey, obj, obj.TableName())
	return obj, err
}

// GetList 分页获取CmsArt记录
// Author 10512203@qq.com
func (m *CmsArtService) GetList(ctx context.Context, seq business.CmsArtSearch, fields string) (list []business.CmsArt, total int64, err error) {
	limit := seq.PageSize
	offset := seq.PageSize * (seq.Page - 1)
	//修改 by ljd  增加查询排序
	order := seq.OrderKey
	desc := seq.OrderDesc
	// 创建db
	db := m.Db.WithContext(ctx).Model(&business.CmsArt{})

	//修改 by ljd
	if seq.Id > 0 {
		db = db.Where("id = ?", seq.Id)
	}

	if !utils.IsEmpty(seq.CreatedAtBegin) && !utils.IsEmpty(seq.CreatedAtEnd) {
		db = db.Where("created_at BETWEEN ? AND ?", seq.CreatedAtBegin, seq.CreatedAtEnd)
	}

	// 如果有条件搜索 下方会自动创建搜索语句
	if seq.Pid != nil {
		db = db.Where("pid = ?", seq.Pid)
	}
	if seq.UserId != nil {
		db = db.Where("user_id = ?", seq.UserId)
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
	if utils.IsEmptyStr(fields) {
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&list).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&list).Error
	}
	m.getFileList(ctx, list)
	return list, total, err
}

// GetListAll 分页获取CmsArt记录 (全部字段)
// Author 10512203@qq.com
func (m *CmsArtService) GetListAll(ctx context.Context, seq business.CmsArtSearch, fields string) (list []business.CmsArt, total int64, err error) {
	return m.GetList(ctx, seq, "*")
}

// GetByMap 根据Map获取单一记录
// Author  [linjd] 10512203@qq.com
func (m *CmsArtService) GetByMap(ctx context.Context, mapData map[string]interface{}, fields string) (obj business.CmsArt, err error) {
	orderBy := "id desc"
	db := m.Db.WithContext(ctx).Model(&business.CmsArt{})
	if utils.IsEmptyStr(fields) {
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
func (m *CmsArtService) GetListByMap(ctx context.Context, mapData map[string]interface{}, fields string, orderBy string) (list []business.CmsArt, err error) {
	if utils.IsEmptyStr(orderBy) {
		orderBy = "id desc"
	}
	db := m.Db.WithContext(ctx).Model(&business.CmsArt{})
	if utils.IsEmptyStr(fields) {
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
func (m *CmsArtService) Count(ctx context.Context, mapDataWhere map[string]interface{}) (total int64, err error) {
	db := m.Db.WithContext(ctx).Model(&business.CmsArt{})
	err = db.Where(mapDataWhere).Count(&total).Error
	return total, err
}

// 如果有文件类型，更新文件 path
func (m *CmsArtService) getFileList(ctx context.Context, list []business.CmsArt) {
	for i, v := range list {
		if !utils.IsEmptyStr(v.Image) {
			fileObjImage := global.FileObj{Guid: v.Image}
			fileObjImage.Path, _ = m.GetPathByGuid(ctx, v.Image)
			v.FileObjList = append(v.FileObjList, fileObjImage)
		}
		if !utils.IsEmptyStr(v.FileList) {
			fileObjFileList := global.FileObj{Guid: v.FileList}
			fileObjFileList.Path, _ = m.GetPathByGuid(ctx, v.FileList)
			v.FileObjList = append(v.FileObjList, fileObjFileList)
		}

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

// 如果有文件类型，更新文件 path
func (m *CmsArtService) getFile(ctx context.Context, v *business.CmsArt) {
	if v != nil {
		if !utils.IsEmptyStr(v.Image) {
			fileObjImage := global.FileObj{Guid: v.Image}
			fileObjImage.Path, _ = m.GetPathByGuid(ctx, v.Image)
			v.FileObjList = append(v.FileObjList, fileObjImage)
		}
		if !utils.IsEmptyStr(v.FileList) {
			fileObjFileList := global.FileObj{Guid: v.FileList}
			fileObjFileList.Path, _ = m.GetPathByGuid(ctx, v.FileList)
			v.FileObjList = append(v.FileObjList, fileObjFileList)
		}

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
