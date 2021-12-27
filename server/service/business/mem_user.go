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

type MemUserService struct {
}

var once_MemUser sync.Once = sync.Once{}
var obj_MemUserService *MemUserService

//获取单例
func GetMemUserService() *MemUserService {
	once_MemUser.Do(func() {
		obj_MemUserService = new(MemUserService)
		//instanse.init()
	})
	return obj_MemUserService
}

// CreateMemUser 创建MemUser记录
// Author [88act](https://github.com/88act)
func (m *MemUserService) CreateMemUser(memUser business.MemUser) (err error) {
	err = global.DB.Create(&memUser).Error
	return err
}

// DeleteMemUser 删除MemUser记录
// Author [88act](https://github.com/88act)
func (m *MemUserService) DeleteMemUser(memUser business.MemUser) (err error) {
	err = global.DB.Delete(&memUser).Error
	return err
}

// DeleteMemUserByIds 批量删除MemUser记录
// Author [88act](https://github.com/88act)
func (m *MemUserService) DeleteMemUserByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]business.MemUser{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateMemUser 更新MemUser记录
// Author [88act](https://github.com/88act)
func (m *MemUserService) UpdateMemUser(memUser business.MemUser) (err error) {
	err = global.DB.Save(&memUser).Error
	return err
}

// GetMemUser 根据id获取MemUser记录
// Author [88act](https://github.com/88act)
func (m *MemUserService) GetMemUser(id uint, fields string) (obj business.MemUser, err error) {

	if utils.IsEmpty(fields) {
		err = global.DB.Where("id = ?", id).First(&obj).Error
	} else {
		err = global.DB.Select(fields).Where("id = ?", id).First(&obj).Error
	}

	//如果有图片image类型，更新图片path
	obj.MapData = make(map[string]string)
	if !utils.IsEmpty(obj.Avatar) {
		_, obj.MapData[obj.Avatar] = common.GetCommonFileService().GetPathByGuid(obj.Avatar)
	}
	return obj, err
}

// GetMemUserInfoList 分页获取MemUser记录
// Author [88act](https://github.com/88act)
func (m *MemUserService) GetMemUserInfoList(info bizReq.MemUserSearch, createdAtBetween []string, fields string) (list []business.MemUserMini, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.MemUser{})
	//var memUsers []business.MemUser
	var memUsers []business.MemUserMini

	//修改 by ljd
	if info.ID > 0 {
		db = db.Where("`id` = ?", info.ID)
	}
	if len(createdAtBetween) >= 2 {
		db = db.Where("`created_at` BETWEEN ? AND ?", createdAtBetween[0], createdAtBetween[1])
	}

	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Sex != nil {
		db = db.Where("`sex` = ?", info.Sex)
	}
	if info.Status != nil {
		db = db.Where("`status` = ?", info.Status)
	}

	err = db.Count(&total).Error
	if err != nil {
		return memUsers, 0, err
	}
	//err = db.Limit(limit).Offset(offset).Find(&memUsers).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&memUsers).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&memUsers).Error
	}
	//如果有图片image类型，更新图片path
	// for i, v := range memUsers {
	// 	v.MapData = make(map[string]string)
	// 	if !utils.IsEmpty(v.) {
	// 		_, v.MapData[v.Avatar] = common.GetCommonFileService().GetPathByGuid(v.Avatar)
	// 	}
	// 	memUsers[i] = v
	// }
	return memUsers, total, err
}

// GetMemUserInfoListAll  分页获取MemUser记录 (全部字段)
// Author [88act](https://github.com/88act)
func (m *MemUserService) GetMemUserInfoListAll(info bizReq.MemUserSearch, createdAtBetween []string, fields string) (list []business.MemUser, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.MemUser{})
	var memUsers []business.MemUser
	//var memUsers []business.MemUserMini

	if info.ID > 0 {
		db = db.Where("`id` = ?", info.ID)
	}
	if len(createdAtBetween) >= 2 {
		db = db.Where("`created_at` BETWEEN ? AND ?", createdAtBetween[0], createdAtBetween[1])
	}

	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Sex != nil {
		db = db.Where("`sex` = ?", info.Sex)
	}
	if info.Status != nil {
		db = db.Where("`status` = ?", info.Status)
	}

	err = db.Count(&total).Error
	if err != nil {
		return memUsers, 0, err
	}

	//err = db.Limit(limit).Offset(offset).Find(&memUsers).Error
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
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&memUsers).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&memUsers).Error
	}
	//如果有图片image类型，更新图片path
	for i, v := range memUsers {
		v.MapData = make(map[string]string)
		if !utils.IsEmpty(v.Avatar) {
			_, v.MapData[v.Avatar] = common.GetCommonFileService().GetPathByGuid(v.Avatar)
		}
		memUsers[i] = v
	}
	return memUsers, total, err
}
