package business

import (
	"go-cms/global"
	"go-cms/model/business"
	bizReq "go-cms/model/business/request"
	"go-cms/model/common/request"
	comSev "go-cms/service/common"
	"go-cms/utils"
	"sync"
)

type CmsCatService struct {
}

var once_CmsCat sync.Once = sync.Once{}
var obj_CmsCatService *CmsCatService

//获取单例
func GetCmsCatSev() *CmsCatService {
	once_CmsCat.Do(func() {
		obj_CmsCatService = new(CmsCatService)
		//instanse.init()
	})
	return obj_CmsCatService
}

// Create 创建CmsCat记录
// Author [88act](https://github.com/88act)
func (m *CmsCatService) Create(data business.CmsCat) (id uint, err error) {
	err = global.DB.Create(&data).Error
	if err != nil {
		return 0, err
	}
	return data.ID, err
}

// Delete 删除CmsCat记录
// Author [88act](https://github.com/88act)
func (m *CmsCatService) Delete(data business.CmsCat) (err error) {
	err = global.DB.Delete(&data).Error
	return err
}

// DeleteByIds 批量删除CmsCat记录
// Author [88act](https://github.com/88act)
func (m *CmsCatService) DeleteByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]business.CmsCat{}, "id in ?", ids.Ids).Error
	return err
}

// Update  更新CmsCat记录
// Author [88act](https://github.com/88act)
func (m *CmsCatService) Update(data business.CmsCat) (err error) {
	err = global.DB.Save(&data).Error
	return err
}

// UpdateByMap  更新CmsCat记录 by Map
// values := map[string]interface{}{
// 	"status":0,
// 	"from": hash,
// }
// Author [88act](https://github.com/88act)
func (m *CmsCatService) UpdateByMap(data business.CmsCat, mapData map[string]interface{}) (err error) {
	err = global.DB.Model(&data).Updates(mapData).Error
	return err
}

// Get 根据id获取CmsCat记录
// Author [88act](https://github.com/88act)
func (m *CmsCatService) Get(id uint, fields string) (obj business.CmsCat, err error) {

	if utils.IsEmpty(fields) {
		err = global.DB.Where("id = ?", id).First(&obj).Error
	} else {
		err = global.DB.Select(fields).Where("id = ?", id).First(&obj).Error
	}

	//如果有图片image类型，更新图片path
	obj.MapData = make(map[string]string)
	if !utils.IsEmpty(obj.Thumb) {
		_, obj.MapData[obj.Thumb] = comSev.GetCommonFileSev().GetPathByGuid(obj.Thumb)
	}
	return obj, err
}

// GetList 分页获取CmsCat记录
// Author [88act](https://github.com/88act)
func (m *CmsCatService) GetList(info bizReq.CmsCatSearch, createdAtBetween []string, fields string) (list []business.CmsCatMini, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.CmsCat{})
	//var cmsCats []business.CmsCat
	var cmsCats []business.CmsCatMini

	//修改 by ljd
	if info.ID > 0 {
		db = db.Where("`id` = ?", info.ID)
	}
	if len(createdAtBetween) >= 2 {
		db = db.Where("`created_at` BETWEEN ? AND ?", createdAtBetween[0], createdAtBetween[1])
	}

	// 如果有条件搜索 下方会自动创建搜索语句
	if info.BeSys != nil {
		db = db.Where("`be_sys` = ?", info.BeSys)
	}
	if info.GroupId != nil {
		db = db.Where("`group_id` = ?", info.GroupId)
	}
	if info.MediaType != nil {
		db = db.Where("`media_type` = ?", info.MediaType)
	}
	if info.Name != "" {
		db = db.Where("`name` LIKE ?", "%"+info.Name+"%")
	}
	if info.Desc != "" {
		db = db.Where("`desc` LIKE ?", "%"+info.Desc+"%")
	}
	if info.Keywords != "" {
		db = db.Where("`keywords` LIKE ?", "%"+info.Keywords+"%")
	}
	if info.Status != nil {
		db = db.Where("`status` = ?", info.Status)
	}

	err = db.Count(&total).Error
	if err != nil {
		return cmsCats, 0, err
	}
	//err = db.Limit(limit).Offset(offset).Find(&cmsCats).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&cmsCats).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&cmsCats).Error
	}
	//如果有图片image类型，更新图片path
	for i, v := range cmsCats {
		v.MapData = make(map[string]string)
		if !utils.IsEmpty(v.Thumb) {
			_, v.MapData[v.Thumb] = comSev.GetCommonFileSev().GetPathByGuid(v.Thumb)
		}
		cmsCats[i] = v
	}
	return cmsCats, total, err
}

//GetListAll 分页获取CmsCat记录 (全部字段)
// Author [88act](https://github.com/88act)
func (m *CmsCatService) GetListAll(info bizReq.CmsCatSearch, createdAtBetween []string, fields string) (list []business.CmsCat, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.CmsCat{})
	var cmsCats []business.CmsCat
	//var cmsCats []business.CmsCatMini

	if info.ID > 0 {
		db = db.Where("`id` = ?", info.ID)
	}
	if len(createdAtBetween) >= 2 {
		db = db.Where("`created_at` BETWEEN ? AND ?", createdAtBetween[0], createdAtBetween[1])
	}

	// 如果有条件搜索 下方会自动创建搜索语句
	if info.BeSys != nil {
		db = db.Where("`be_sys` = ?", info.BeSys)
	}
	if info.GroupId != nil {
		db = db.Where("`group_id` = ?", info.GroupId)
	}
	if info.MediaType != nil {
		db = db.Where("`media_type` = ?", info.MediaType)
	}
	if info.Name != "" {
		db = db.Where("`name` LIKE ?", "%"+info.Name+"%")
	}
	if info.Desc != "" {
		db = db.Where("`desc` LIKE ?", "%"+info.Desc+"%")
	}
	if info.Keywords != "" {
		db = db.Where("`keywords` LIKE ?", "%"+info.Keywords+"%")
	}
	if info.Status != nil {
		db = db.Where("`status` = ?", info.Status)
	}

	err = db.Count(&total).Error
	if err != nil {
		return cmsCats, 0, err
	}

	//err = db.Limit(limit).Offset(offset).Find(&cmsCats).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&cmsCats).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&cmsCats).Error
	}
	//如果有图片image类型，更新图片path
	for i, v := range cmsCats {
		v.MapData = make(map[string]string)
		if !utils.IsEmpty(v.Thumb) {
			_, v.MapData[v.Thumb] = comSev.GetCommonFileSev().GetPathByGuid(v.Thumb)
		}
		cmsCats[i] = v
	}
	return cmsCats, total, err
}
