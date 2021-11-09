package business

import (
	"github.com/88act/go-cms/server/global"
	"github.com/88act/go-cms/server/model/business"
	businessReq "github.com/88act/go-cms/server/model/business/request"
	"github.com/88act/go-cms/server/model/common/request"
	"github.com/88act/go-cms/server/utils"
)

type BasicFileService struct {
}

var basicFileFields string = "aaa"

// CreateBasicFile 创建BasicFile记录
// Author [88act](https://github.com/88act)
func (basicFileService *BasicFileService) CreateBasicFile(basicFile business.BasicFile) (err error) {
	err = global.DB.Create(&basicFile).Error
	return err
}

// DeleteBasicFile 删除BasicFile记录
// Author [88act](https://github.com/88act)
func (basicFileService *BasicFileService) DeleteBasicFile(basicFile business.BasicFile) (err error) {
	err = global.DB.Delete(&basicFile).Error
	return err
}

// DeleteBasicFileByIds 批量删除BasicFile记录
// Author [88act](https://github.com/88act)
func (basicFileService *BasicFileService) DeleteBasicFileByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]business.BasicFile{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateBasicFile 更新BasicFile记录
// Author [88act](https://github.com/88act)
func (basicFileService *BasicFileService) UpdateBasicFile(basicFile business.BasicFile) (err error) {
	err = global.DB.Save(&basicFile).Error
	return err
}

// GetBasicFile 根据id获取BasicFile记录
// Author  有修改的 LJD
func (basicFileService *BasicFileService) GetBasicFile(key string, val string) (err error, basicFile business.BasicFile) {
	err = global.DB.Where(key, val).First(&basicFile).Error
	if !utils.IsEmpty(basicFile.Path) {
		basicFile.Path = global.CONFIG.Local.BaseUrl + basicFile.Path
	}
	return err, basicFile
}

// GetBasicFileInfoList 分页获取BasicFile记录
// Author [88act](https://github.com/88act)
func (basicFileService *BasicFileService) GetBasicFileInfoList(info businessReq.BasicFileSearch, createdAtBetween []string) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.BasicFile{})
	var basicFiles []business.BasicFile

	//修改 by ljd
	if info.ID > 0 {
		db = db.Where("`id` = ?", info.ID)
	}
	if createdAtBetween != nil && len(createdAtBetween) > 0 {
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
	err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&basicFiles).Error
	//更新图片path
	for i, v := range basicFiles {
		if !utils.IsEmpty(v.Path) {
			v.Path = global.CONFIG.Local.BaseUrl + v.Path
		}
		basicFiles[i] = v
	}

	return err, basicFiles, total
}
