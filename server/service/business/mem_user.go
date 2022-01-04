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

type MemUserService struct {
}

var once_MemUser sync.Once = sync.Once{}
var obj_MemUserService *MemUserService

//获取单例
func GetMemUserSev() *MemUserService {
	once_MemUser.Do(func() {
		obj_MemUserService = new(MemUserService)
		//instanse.init()
	})
	return obj_MemUserService
}

// Create 创建MemUser记录
// Author [88act](https://github.com/88act)
func (m *MemUserService) Create(data business.MemUser) (id uint, err error) {
	err = global.DB.Create(&data).Error
	if err != nil {
		return 0, err
	}
	return data.ID, err
}

// Delete 删除MemUser记录
// Author [88act](https://github.com/88act)
func (m *MemUserService) Delete(data business.MemUser) (err error) {
	err = global.DB.Delete(&data).Error
	return err
}

// DeleteByIds 批量删除MemUser记录
// Author [88act](https://github.com/88act)
func (m *MemUserService) DeleteByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]business.MemUser{}, "id in ?", ids.Ids).Error
	return err
}

// Update  更新MemUser记录
// Author [88act](https://github.com/88act)
func (m *MemUserService) Update(data business.MemUser) (err error) {
	err = global.DB.Save(&data).Error
	return err
}

// UpdateByMap  更新MemUser记录 by Map
// values := map[string]interface{}{
// 	"status":0,
// 	"from": hash,
// }
// Author [88act](https://github.com/88act)
func (m *MemUserService) UpdateByMap(data business.MemUser, mapData map[string]interface{}) (err error) {
	err = global.DB.Model(&data).Updates(mapData).Error
	return err
}

// Get 根据id获取MemUser记录
// Author [88act](https://github.com/88act)
func (m *MemUserService) Get(id uint, fields string) (obj business.MemUser, err error) {

	if utils.IsEmpty(fields) {
		err = global.DB.Where("id = ?", id).First(&obj).Error
	} else {
		err = global.DB.Select(fields).Where("id = ?", id).First(&obj).Error
	}

	//如果有图片image类型，更新图片path
	obj.MapData = make(map[string]string)
	if !utils.IsEmpty(obj.Avatar) {
		_, obj.MapData[obj.Avatar] = comSev.GetCommonFileSev().GetPathByGuid(obj.Avatar)
	}
	return obj, err
}

// GetList 分页获取MemUser记录
// Author [88act](https://github.com/88act)
func (m *MemUserService) GetList(info bizReq.MemUserSearch, createdAtBetween []string, fields string) (list []business.MemUserMini, total int64, err error) {
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
	if info.Username != "" {
		db = db.Where("`username` LIKE ?", "%"+info.Username+"%")
	}
	if info.Pws != "" {
		db = db.Where("`pws` = ?", info.Pws)
	}
	if info.Email != "" {
		db = db.Where("`email` LIKE ?", "%"+info.Email+"%")
	}
	if info.Mobile != "" {
		db = db.Where("`mobile` LIKE ?", "%"+info.Mobile+"%")
	}
	if info.Nickname != "" {
		db = db.Where("`nickname` LIKE ?", "%"+info.Nickname+"%")
	}
	if info.Realname != "" {
		db = db.Where("`realname` LIKE ?", "%"+info.Realname+"%")
	}
	if info.CardId != "" {
		db = db.Where("`card_id` = ?", info.CardId)
	}
	if info.Sex != nil {
		db = db.Where("`sex` = ?", info.Sex)
	}
	if info.Status != nil {
		db = db.Where("`status` = ?", info.Status)
	}
	if info.StatusSafe != nil {
		db = db.Where("`status_safe` = ?", info.StatusSafe)
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
	//  v.MapData = make(map[string]string)
	//     if !utils.IsEmpty(v.Avatar) {
	//         _, v.MapData[v.Avatar] = comSev.GetCommonFileSev().GetPathByGuid(v.Avatar)
	//     }
	//   memUsers[i] = v
	// }
	return memUsers, total, err
}

//GetListAll 分页获取MemUser记录 (全部字段)
// Author [88act](https://github.com/88act)
func (m *MemUserService) GetListAll(info bizReq.MemUserSearch, createdAtBetween []string, fields string) (list []business.MemUser, total int64, err error) {
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
	if info.Username != "" {
		db = db.Where("`username` LIKE ?", "%"+info.Username+"%")
	}
	if info.Pws != "" {
		db = db.Where("`pws` = ?", info.Pws)
	}
	if info.Email != "" {
		db = db.Where("`email` LIKE ?", "%"+info.Email+"%")
	}
	if info.Mobile != "" {
		db = db.Where("`mobile` LIKE ?", "%"+info.Mobile+"%")
	}
	if info.Nickname != "" {
		db = db.Where("`nickname` LIKE ?", "%"+info.Nickname+"%")
	}
	if info.Realname != "" {
		db = db.Where("`realname` LIKE ?", "%"+info.Realname+"%")
	}
	if info.CardId != "" {
		db = db.Where("`card_id` = ?", info.CardId)
	}
	if info.Sex != nil {
		db = db.Where("`sex` = ?", info.Sex)
	}
	if info.Status != nil {
		db = db.Where("`status` = ?", info.Status)
	}
	if info.StatusSafe != nil {
		db = db.Where("`status_safe` = ?", info.StatusSafe)
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
			_, v.MapData[v.Avatar] = comSev.GetCommonFileSev().GetPathByGuid(v.Avatar)
		}
		memUsers[i] = v
	}
	return memUsers, total, err
}
