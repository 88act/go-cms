package system

import (
	"go-cms/global"
	"go-cms/model/system"
	"go-cms/model/system/request"
)

//@author: [88act](https://github.com/88act)
//@function: CreateSysDictionaryDetail
//@description: 创建字典详情数据
//@param: sysDictionaryDetail model.SysDictionaryDetail
//@return: err error

type DictionaryDetailService struct {
}

func (dictionaryDetailService *DictionaryDetailService) CreateSysDictionaryDetail(sysDictionaryDetail system.SysDictionaryDetail) (err error) {
	err = global.DB.Create(&sysDictionaryDetail).Error
	return err
}

//@author: [88act](https://github.com/88act)
//@function: DeleteSysDictionaryDetail
//@description: 删除字典详情数据
//@param: sysDictionaryDetail model.SysDictionaryDetail
//@return: err error

func (dictionaryDetailService *DictionaryDetailService) DeleteSysDictionaryDetail(sysDictionaryDetail system.SysDictionaryDetail) (err error) {
	err = global.DB.Delete(&sysDictionaryDetail).Error
	return err
}

//@author: [88act](https://github.com/88act)
//@function: UpdateSysDictionaryDetail
//@description: 更新字典详情数据
//@param: sysDictionaryDetail *model.SysDictionaryDetail
//@return: err error

func (dictionaryDetailService *DictionaryDetailService) UpdateSysDictionaryDetail(sysDictionaryDetail *system.SysDictionaryDetail) (err error) {
	err = global.DB.Save(sysDictionaryDetail).Error
	return err
}

//@author: [88act](https://github.com/88act)
//@function: GetSysDictionaryDetail
//@description: 根据id获取字典详情单条数据
//@param: id uint
//@return: err error, sysDictionaryDetail model.SysDictionaryDetail

func (dictionaryDetailService *DictionaryDetailService) GetSysDictionaryDetail(id uint) (err error, sysDictionaryDetail system.SysDictionaryDetail) {
	err = global.DB.Where("id = ?", id).First(&sysDictionaryDetail).Error
	return
}

//@author: [88act](https://github.com/88act)
//@function: GetSysDictionaryDetailInfoList
//@description: 分页获取字典详情列表
//@param: info request.SysDictionaryDetailSearch
//@return: err error, list interface{}, total int64

func (dictionaryDetailService *DictionaryDetailService) GetSysDictionaryDetailInfoList(info request.SysDictionaryDetailSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.DB.Model(&system.SysDictionaryDetail{})
	var sysDictionaryDetails []system.SysDictionaryDetail
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Label != "" {
		db = db.Where("label LIKE ?", "%"+info.Label+"%")
	}
	if info.Value != 0 {
		db = db.Where("value = ?", info.Value)
	}
	if info.Status != nil {
		db = db.Where("status = ?", info.Status)
	}
	if info.SysDictionaryID != 0 {
		db = db.Where("sys_dictionary_id = ?", info.SysDictionaryID)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Order("sort asc,id asc").Limit(limit).Offset(offset).Find(&sysDictionaryDetails).Error
	return err, sysDictionaryDetails, total
}

//@author: ljd
//@function: GetNameByValue
//@description: 根据val获取字典 label
//@param: dictType int, val int
//@return: err error, sysDictionaryDetail model.SysDictionaryDetail
func (dictionaryDetailService *DictionaryDetailService) GetNameByValue(dictType int, val int) (string, error) {
	var sysDictionaryDetail system.SysDictionaryDetail
	err := global.DB.Where("sys_dictionary_id = ? AND value=?", dictType, val).First(&sysDictionaryDetail).Error
	return sysDictionaryDetail.Label, err
}
