package business

import (
	"go-cms/global"
	"go-cms/model/business"
	bizReq "go-cms/model/business/request"
	"go-cms/model/common/request"
	"go-cms/utils"
	"sync"
)

type BasicFileService struct {
}

var once_BasicFile sync.Once = sync.Once{}
var obj_BasicFileService *BasicFileService

//获取单例
func GetBasicFileService() *BasicFileService {
	once_BasicFile.Do(func() {
		obj_BasicFileService = new(BasicFileService)
		//instanse.init()
	})
	return obj_BasicFileService
}

// CreateBasicFile 创建BasicFile记录
// Author [88act](https://github.com/88act)
func (m *BasicFileService) CreateBasicFile(basicFile business.BasicFile) (err error) {
	err = global.DB.Create(&basicFile).Error
	return err
}

// DeleteBasicFile 删除BasicFile记录
// Author [88act](https://github.com/88act)
func (m *BasicFileService) DeleteBasicFile(basicFile business.BasicFile) (err error) {
	err = global.DB.Delete(&basicFile).Error
	return err
}

// DeleteBasicFileByIds 批量删除BasicFile记录
// Author [88act](https://github.com/88act)
func (m *BasicFileService) DeleteBasicFileByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]business.BasicFile{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateBasicFile 更新BasicFile记录
// Author [88act](https://github.com/88act)
func (m *BasicFileService) UpdateBasicFile(basicFile business.BasicFile) (err error) {
	err = global.DB.Save(&basicFile).Error
	return err
}

// GetBasicFile 根据id获取BasicFile记录
// Author [88act](https://github.com/88act)
func (m *BasicFileService) GetBasicFile(id uint, fields string) (obj business.BasicFile, err error) {

	if utils.IsEmpty(fields) {
		err = global.DB.Where("id = ?", id).First(&obj).Error
	} else {
		err = global.DB.Select(fields).Where("id = ?", id).First(&obj).Error
	}

	//如果有图片image类型，更新图片path
	obj.MapData = make(map[string]string)
	return obj, err
}

// GetBasicFileInfoList 分页获取BasicFile记录
// Author [88act](https://github.com/88act)
func (m *BasicFileService) GetBasicFileInfoList(info bizReq.BasicFileSearch, createdAtBetween []string, fields string) (list []business.BasicFileMini, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.BasicFile{})
	//var basicFiles []business.BasicFile
	var basicFiles []business.BasicFileMini

	//修改 by ljd
	if info.ID > 0 {
		db = db.Where("`id` = ?", info.ID)
	}
	if len(createdAtBetween) >= 2 {
		db = db.Where("`created_at` BETWEEN ? AND ?", createdAtBetween[0], createdAtBetween[1])
	}

	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Guid != "" {
		db = db.Where("`guid` = ?", info.Guid)
	}
	if info.UserId != nil {
		db = db.Where("`user_id` = ?", info.UserId)
	}
	if info.Name != "" {
		db = db.Where("`name` LIKE ?", "%"+info.Name+"%")
	}
	if info.Module != nil {
		db = db.Where("`module` = ?", info.Module)
	}
	if info.MediaType != nil {
		db = db.Where("`media_type` = ?", info.MediaType)
	}
	if info.Driver != nil {
		db = db.Where("`driver` = ?", info.Driver)
	}

	err = db.Count(&total).Error
	if err != nil {
		return basicFiles, 0, err
	}
	//err = db.Limit(limit).Offset(offset).Find(&basicFiles).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&basicFiles).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&basicFiles).Error
	}
	//如果有图片image类型，更新图片path
	for i, v := range basicFiles {
		if !utils.IsEmpty(v.Path) {
			v.Path = global.CONFIG.Local.BaseUrl + v.Path
		}
		basicFiles[i] = v
	}
	return basicFiles, total, err
}

// GetBasicFileInfoListAll  分页获取BasicFile记录 (全部字段)
// Author [88act](https://github.com/88act)
func (m *BasicFileService) GetBasicFileInfoListAll(info bizReq.BasicFileSearch, createdAtBetween []string, fields string) (list []business.BasicFile, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.BasicFile{})
	var basicFiles []business.BasicFile
	//var basicFiles []business.BasicFileMini

	if info.ID > 0 {
		db = db.Where("`id` = ?", info.ID)
	}
	if len(createdAtBetween) >= 2 {
		db = db.Where("`created_at` BETWEEN ? AND ?", createdAtBetween[0], createdAtBetween[1])
	}

	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Guid != "" {
		db = db.Where("`guid` = ?", info.Guid)
	}
	if info.UserId != nil {
		db = db.Where("`user_id` = ?", info.UserId)
	}
	if info.Name != "" {
		db = db.Where("`name` LIKE ?", "%"+info.Name+"%")
	}
	if info.Module != nil {
		db = db.Where("`module` = ?", info.Module)
	}
	if info.MediaType != nil {
		db = db.Where("`media_type` = ?", info.MediaType)
	}
	if info.Driver != nil {
		db = db.Where("`driver` = ?", info.Driver)
	}

	err = db.Count(&total).Error
	if err != nil {
		return basicFiles, 0, err
	}

	//err = db.Limit(limit).Offset(offset).Find(&basicFiles).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&basicFiles).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&basicFiles).Error
	}
	//如果有图片image类型，更新图片path
	for i, v := range basicFiles {
		if !utils.IsEmpty(v.Path) {
			v.Path = global.CONFIG.Local.BaseUrl + v.Path
		}
		basicFiles[i] = v
	}

	return basicFiles, total, err
}

// GetBasicFile 根据id获取BasicFile记录
// Author  有修改的 LJD
func (m *BasicFileService) GetBasicFileByKey(key string, val string) (basicFile business.BasicFile, err error) {
	err = global.DB.Where(key, val).First(&basicFile).Error
	if !utils.IsEmpty(basicFile.Path) {
		basicFile.Path = global.CONFIG.Local.BaseUrl + basicFile.Path
	}
	return basicFile, err
}
