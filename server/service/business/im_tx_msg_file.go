package business

import (
	"go-cms/global"
	"go-cms/model/business"
	bizReq "go-cms/model/business/request"
	"go-cms/model/common/request"
	"go-cms/utils"
	"sync"
)

type ImTxMsgFileService struct {
}

var once_ImTxMsgFile sync.Once = sync.Once{}
var obj_ImTxMsgFileService *ImTxMsgFileService

//获取单例
func GetImTxMsgFileService() *ImTxMsgFileService {
	once_ImTxMsgFile.Do(func() {
		obj_ImTxMsgFileService = new(ImTxMsgFileService)
		//instanse.init()
	})
	return obj_ImTxMsgFileService
}

// CreateImTxMsgFile 创建ImTxMsgFile记录
// Author [88act](https://github.com/88act)
func (m *ImTxMsgFileService) CreateImTxMsgFile(imTxMsgFile business.ImTxMsgFile) (err error) {
	err = global.DB.Create(&imTxMsgFile).Error
	return err
}

// DeleteImTxMsgFile 删除ImTxMsgFile记录
// Author [88act](https://github.com/88act)
func (m *ImTxMsgFileService) DeleteImTxMsgFile(imTxMsgFile business.ImTxMsgFile) (err error) {
	err = global.DB.Delete(&imTxMsgFile).Error
	return err
}

// DeleteImTxMsgFileByIds 批量删除ImTxMsgFile记录
// Author [88act](https://github.com/88act)
func (m *ImTxMsgFileService) DeleteImTxMsgFileByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]business.ImTxMsgFile{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateImTxMsgFile 更新ImTxMsgFile记录
// Author [88act](https://github.com/88act)
func (m *ImTxMsgFileService) UpdateImTxMsgFile(imTxMsgFile business.ImTxMsgFile) (err error) {
	err = global.DB.Save(&imTxMsgFile).Error
	return err
}

// GetImTxMsgFile 根据id获取ImTxMsgFile记录
// Author [88act](https://github.com/88act)
func (m *ImTxMsgFileService) GetImTxMsgFile(id uint, fields string) (obj business.ImTxMsgFile, err error) {

	if utils.IsEmpty(fields) {
		err = global.DB.Where("id = ?", id).First(&obj).Error
	} else {
		err = global.DB.Select(fields).Where("id = ?", id).First(&obj).Error
	}

	//如果有图片image类型，更新图片path
	obj.MapData = make(map[string]string)
	return obj, err
}

// GetImTxMsgFileInfoList 分页获取ImTxMsgFile记录
// Author [88act](https://github.com/88act)
func (m *ImTxMsgFileService) GetImTxMsgFileInfoList(info bizReq.ImTxMsgFileSearch, createdAtBetween []string, fields string) (list []business.ImTxMsgFileMini, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.ImTxMsgFile{})
	//var imTxMsgFiles []business.ImTxMsgFile
	var imTxMsgFiles []business.ImTxMsgFileMini

	//修改 by ljd
	if info.ID > 0 {
		db = db.Where("`id` = ?", info.ID)
	}
	if len(createdAtBetween) >= 2 {
		db = db.Where("`created_at` BETWEEN ? AND ?", createdAtBetween[0], createdAtBetween[1])
	}

	// 如果有条件搜索 下方会自动创建搜索语句
	if info.ChatType != "" {
		db = db.Where("`chat_type` = ?", info.ChatType)
	}
	if info.MsgTime != "" {
		db = db.Where("`msg_time` = ?", info.MsgTime)
	}
	if info.Status != nil {
		db = db.Where("`status` = ?", info.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return imTxMsgFiles, 0, err
	}
	//err = db.Limit(limit).Offset(offset).Find(&imTxMsgFiles).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&imTxMsgFiles).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&imTxMsgFiles).Error
	}
	//如果有图片image类型，更新图片path
	for i, v := range imTxMsgFiles {
		v.MapData = make(map[string]string)
		imTxMsgFiles[i] = v
	}
	return imTxMsgFiles, total, err
}

// GetImTxMsgFileInfoListAll  分页获取ImTxMsgFile记录 (全部字段)
// Author [88act](https://github.com/88act)
func (m *ImTxMsgFileService) GetImTxMsgFileInfoListAll(info bizReq.ImTxMsgFileSearch, createdAtBetween []string, fields string) (list []business.ImTxMsgFile, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.ImTxMsgFile{})
	var imTxMsgFiles []business.ImTxMsgFile
	//var imTxMsgFiles []business.ImTxMsgFileMini

	if info.ID > 0 {
		db = db.Where("`id` = ?", info.ID)
	}
	if len(createdAtBetween) >= 2 {
		db = db.Where("`created_at` BETWEEN ? AND ?", createdAtBetween[0], createdAtBetween[1])
	}

	// 如果有条件搜索 下方会自动创建搜索语句
	if info.ChatType != "" {
		db = db.Where("`chat_type` = ?", info.ChatType)
	}
	if info.MsgTime != "" {
		db = db.Where("`msg_time` = ?", info.MsgTime)
	}
	if info.Status != nil {
		db = db.Where("`status` = ?", info.Status)
	}

	err = db.Count(&total).Error
	if err != nil {
		return imTxMsgFiles, 0, err
	}

	//err = db.Limit(limit).Offset(offset).Find(&imTxMsgFiles).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&imTxMsgFiles).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&imTxMsgFiles).Error
	}
	//如果有图片image类型，更新图片path
	for i, v := range imTxMsgFiles {
		v.MapData = make(map[string]string)
		imTxMsgFiles[i] = v
	}
	return imTxMsgFiles, total, err
}
