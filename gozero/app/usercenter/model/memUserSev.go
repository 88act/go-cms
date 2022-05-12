package model

import (
	"context"
	"go-cms/common/utils"

	"github.com/zeromicro/go-zero/core/stores/cache"

	"gorm.io/gorm"
)

type MemUserSev struct {
	db    *gorm.DB
	cache cache.ClusterConf
}

func NewMemUserSev(db *gorm.DB, cache cache.ClusterConf) *MemUserSev {
	instance := new(MemUserSev)
	instance.db = db
	instance.cache = cache
	return instance
}

// Create 创建记录
// Author [88act](https://github.com/88act)
func (m *MemUserSev) Create(ctx context.Context, data MemUser) (id int64, err error) {
	// fmt.Println("tracespec.TracingKey 1= ")
	// if v := ctx.Value(tracespec.TracingKey); v != nil {
	// 	fmt.Println("tracespec.TracingKey 2= ", v)
	// }

	err = m.db.WithContext(ctx).Create(&data).Error
	if err != nil {
		return 0, err
	}
	return data.Id, err
}

// Delete 删除记录
// Author [88act](https://github.com/88act)
func (m *MemUserSev) Delete(ctx context.Context, data MemUser) (err error) {
	err = m.db.WithContext(ctx).Delete(&data).Error
	return err
}

// DeleteByIds 批量删除记录
// Author [88act](https://github.com/88act)
func (m *MemUserSev) DeleteByIds(ctx context.Context, ids IdsReq) (err error) {
	err = m.db.WithContext(ctx).Delete(&[]MemUser{}, "id in ?", ids.Ids).Error
	return err
}

// Update  更新记录
// Author [88act](https://github.com/88act)
func (m *MemUserSev) Update(ctx context.Context, data MemUser) (err error) {
	err = m.db.WithContext(ctx).Save(&data).Error
	return err
}

// UpdateByMap  更新记录 by Map
// Author [88act](https://github.com/88act)
func (m *MemUserSev) UpdateByMap(ctx context.Context, data MemUser, mapData map[string]interface{}) (err error) {
	err = m.db.WithContext(ctx).Model(&data).Updates(mapData).Error
	return err
}

// Get 根据id获取记录
// Author [88act](https://github.com/88act)
func (m *MemUserSev) Get(ctx context.Context, id int64, fields string) (obj *MemUser, err error) {

	if utils.IsEmpty(fields) {
		err = m.db.WithContext(ctx).Where("id = ?", id).First(&obj).Error
	} else {
		err = m.db.WithContext(ctx).Select(fields).Where("id = ?", id).First(&obj).Error
	}
	if err != nil {
		return nil, err
	}

	// image类型 图片path
	obj.MapData = make(map[string]string)
	return obj, err
}

// GetByMap 根据Map获取记录
// Author [88act](https://github.com/88act)
func (m *MemUserSev) GetByMap(ctx context.Context, mapData map[string]interface{}, fields string) (obj *MemUser, err error) {

	if utils.IsEmpty(fields) {
		err = m.db.WithContext(ctx).Where(mapData).First(&obj).Error
	} else {
		err = m.db.WithContext(ctx).Select(fields).Where(mapData).First(&obj).Error
	}
	if err != nil {
		return nil, err
	}

	// image类型 图片path
	obj.MapData = make(map[string]string)
	return obj, err
}

// GetList 分页获取记录
// Author [88act](https://github.com/88act)
func (m *MemUserSev) GetList(ctx context.Context, seq MemUserSearch, fields string) (list []MemUser, total int64, err error) {
	limit := seq.PageSize
	offset := seq.PageSize * (seq.Page - 1)
	order := seq.OrderKey
	desc := seq.OrderDesc

	db := m.db.WithContext(ctx).Model(&MemUser{})
	var memUsers []MemUser

	if seq.Id > 0 {
		db = db.Where("`id` = ?", seq.Id)
	}
	if len(seq.CreatedAtBetween) >= 2 {
		db = db.Where("`created_at` BETWEEN ? AND ?", seq.CreatedAtBetween[0], seq.CreatedAtBetween[1])
	}
	// 如果有条件搜索
	if seq.Username != "" {
		db = db.Where("`username` = ?", seq.Username)
	}
	if seq.Email != "" {
		db = db.Where("`email` = ?", seq.Email)
	}
	if seq.Mobile != "" {
		db = db.Where("`mobile` = ?", seq.Mobile)
	}
	if seq.Nickname != "" {
		db = db.Where("`nickname` = ?", seq.Nickname)
	}
	if seq.Realname != "" {
		db = db.Where("`realname` = ?", seq.Realname)
	}
	if seq.CardId != "" {
		db = db.Where("`card_id` = ?", seq.CardId)
	}
	if seq.Sex != nil {
		db = db.Where("`sex` = ?", seq.Sex)
	}
	if seq.MobileValidated != nil {
		db = db.Where("`mobile_validated` = ?", seq.MobileValidated)
	}
	if seq.EmailValidated != nil {
		db = db.Where("`email_validated` = ?", seq.EmailValidated)
	}
	if seq.CardidValidated != nil {
		db = db.Where("`cardid_validated` = ?", seq.CardidValidated)
	}
	if seq.RecommendCode != "" {
		db = db.Where("`recommend_code` = ?", seq.RecommendCode)
	}
	if seq.Status != nil {
		db = db.Where("`status` = ?", seq.Status)
	}
	if seq.StatusSafe != nil {
		db = db.Where("`status_safe` = ?", seq.StatusSafe)
	}
	if seq.RegIp != nil {
		db = db.Where("`reg_ip` = ?", seq.RegIp)
	}
	if seq.LoginIp != nil {
		db = db.Where("`login_ip` = ?", seq.LoginIp)
	}
	if seq.LoginTime != nil {
		db = db.Where("`login_time` > > ?", seq.LoginTime)
	}

	err = db.Count(&total).Error
	if err != nil {
		return memUsers, 0, err
	}
	OrderStr := "id desc"
	if !utils.IsEmpty(order) {
		if desc {
			OrderStr = order + " desc"
		} else {
			OrderStr = order
		}
	}
	if !utils.IsEmpty(fields) {
		db.Select(fields)
	}
	if limit > 0 {
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&memUsers).Error
	} else {
		err = db.Order(OrderStr).Find(&memUsers).Error
	}
	//image类型,获取path
	for i, v := range memUsers {
		v.MapData = make(map[string]string)
		memUsers[i] = v
	}
	return memUsers, total, err
}
