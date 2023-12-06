package model

import (
	"context"
	"go-cms/common/mycache"
	"go-cms/common/utils"

	. "go-cms/common/baseModel"

	"gorm.io/gorm"
)

type MemUserSev struct {
	BaseModelSev
}

func NewMemUserSev(db *gorm.DB) *MemUserSev {
	instance := new(MemUserSev)
	instance.Db = db
	instance.CacheKeyPre = "MemUser"
	instance.CacheKeyListPre = "MemUserList"
	return instance
}

// Get 根据username获取
// Author 10512203@qq.com
func (m *MemUserSev) GetByUsername(ctx context.Context, cuId int64, username string) (obj MemUser, err error) {
	// cacheKey := m.GetCacheKey("MemUser_name", username)
	// if cacheData, err := mycache.GetCache().Get(cacheKey); err == nil {
	// 	obj = cacheData.(business.MemUser)
	// 	return obj, err
	// }
	err = m.Db.WithContext(ctx).Where("username = ?", username).First(&obj).Error
	// if err == nil {
	// 	m.getFile(ctx, &obj)
	// 	//mycache.GetCache().Set(cacheKey, obj, "")
	// }
	return obj, err
}

// Create 创建MemUser记录
// Author 10512203@qq.com
func (m *MemUserSev) Create(ctx context.Context, data MemUser) (id int64, err error) {
	err = m.Db.WithContext(ctx).Create(&data).Error
	if err == nil {
		mycache.GetCache().DeleteKeyPre(m.CacheKeyListPre)
		return data.Id, err
	} else {
		return 0, err
	}
}

// Delete 删除MemUser记录
// Author 10512203@qq.com
func (m *MemUserSev) Delete(ctx context.Context, data MemUser) (err error) {
	err = m.Db.WithContext(ctx).Delete(&data).Error
	if err == nil {
		m.DelCache(m.CacheKeyPre, []int64{data.Id})
	}
	return err
}

// DeleteByIds 批量删除MemUser记录
// Author 10512203@qq.com
func (m *MemUserSev) DeleteByIds(ctx context.Context, ids IdsReq) (err error) {
	err = m.Db.WithContext(ctx).Delete(&[]MemUser{}, "id in ?", ids.Ids).Error
	if err == nil {
		m.DelCache(m.CacheKeyPre, ids.Ids)
	}
	return err
}

// Update  更新MemUser记录
// Author 10512203@qq.com
func (m *MemUserSev) Update(ctx context.Context, data MemUser) (err error) {
	err = m.Db.WithContext(ctx).Save(&data).Error
	if err == nil {
		m.DelCache(m.CacheKeyPre, []int64{data.Id})
	}
	return err
}

// UpdateByMap  更新MemUser记录 by Map
// Author 10512203@qq.com
func (m *MemUserSev) UpdateByMap(ctx context.Context, data MemUser, mapData map[string]interface{}) (err error) {
	err = m.Db.WithContext(ctx).Model(&data).Updates(mapData).Error
	if err == nil {
		m.DelCache(m.CacheKeyPre, []int64{data.Id})
	}
	return err
}

// UpdateExpr 字段自增/自减
// Author 10512203@qq.com
func (m *MemUserSev) UpdateExpr(ctx context.Context, data MemUser, column string, num int) (err error) {
	err = m.Db.WithContext(ctx).Model(&data).UpdateColumn(column, gorm.Expr(column+" + ?", num)).Error
	return err
}

// Get 根据id获取MemUser记录
// Author 10512203@qq.com
func (m *MemUserSev) Get(ctx context.Context, id int64, fields string) (obj MemUser, err error) {

	if id <= 0 {
		return MemUser{}, nil
	}
	cacheKey := m.GetCacheKey(m.CacheKeyPre, id)
	if cacheData, err := mycache.GetCache().GetObj(cacheKey); err == nil {
		obj = cacheData.(MemUser)
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

// GetList 分页获取MemUser记录
// Author 10512203@qq.com
func (m *MemUserSev) GetList(ctx context.Context, seq MemUserSearch, fields string) (list []MemUser, total int64, err error) {
	limit := seq.PageSize
	offset := seq.PageSize * (seq.Page - 1)
	//修改 by ljd  增加查询排序
	order := seq.OrderKey
	desc := seq.OrderDesc
	// 创建db
	db := m.Db.WithContext(ctx).Model(&MemUser{})
	//强制状态为1
	db = db.Where("status = 1")
	if seq.Id > 0 {
		db = db.Where("id = ?", seq.Id)
	}
	if seq.CreatedAtBegin != nil && seq.CreatedAtEnd != nil {
		db = db.Where("created_at BETWEEN ? AND ?", seq.CreatedAtBegin.Format("2006-01-02 15:04:05"), seq.CreatedAtEnd.Format("2006-01-02 15:04:05"))
	}

	// 如果有条件搜索 下方会自动创建搜索语句
	if seq.CuId != nil {
		db = db.Where("cu_id = ?", seq.CuId)
	}
	if seq.UserId != nil {
		db = db.Where("user_id = ?", seq.UserId)
	}
	if seq.Guid != "" {
		db = db.Where("guid = ?", seq.Guid)
	}
	if seq.Username != "" {
		db = db.Where("username = ?", seq.Username)
	}
	if seq.Password != "" {
		db = db.Where("password = ?", seq.Password)
	}
	if seq.PasswordSlat != "" {
		db = db.Where("password_slat = ?", seq.PasswordSlat)
	}
	if seq.Nickname != "" {
		db = db.Where("nickname = ?", seq.Nickname)
	}
	if seq.Realname != "" {
		db = db.Where("realname = ?", seq.Realname)
	}
	if seq.RoleList != "" {
		db = db.Where("role_list = ?", seq.RoleList)
	}
	if seq.Role != nil {
		db = db.Where("role = ?", seq.Role)
	}
	if seq.UserType != nil {
		db = db.Where("user_type = ?", seq.UserType)
	}
	if seq.Email != "" {
		db = db.Where("email = ?", seq.Email)
	}
	if seq.Mobile != "" {
		db = db.Where("mobile = ?", seq.Mobile)
	}
	if seq.CardId != "" {
		db = db.Where("card_id = ?", seq.CardId)
	}
	if seq.Sex != nil {
		db = db.Where("sex = ?", seq.Sex)
	}
	if seq.Birthday != nil {
		db = db.Where("birthday = > ?", seq.Birthday)
	}
	if seq.JobId != nil {
		db = db.Where("job_id = ?", seq.JobId)
	}
	if seq.DepartId != nil {
		db = db.Where("depart_id = ?", seq.DepartId)
	}
	if seq.MobileValidated != nil {
		db = db.Where("mobile_validated = ?", seq.MobileValidated)
	}
	if seq.EmailValidated != nil {
		db = db.Where("email_validated = ?", seq.EmailValidated)
	}
	if seq.CardidValidated != nil {
		db = db.Where("cardid_validated = ?", seq.CardidValidated)
	}
	if seq.Remark != "" {
		db = db.Where("remark = ?", seq.Remark)
	}
	if seq.StatusSafe != nil {
		db = db.Where("status_safe = ?", seq.StatusSafe)
	}
	if seq.RegIp != nil {
		db = db.Where("reg_ip = ?", seq.RegIp)
	}
	if seq.LoginIp != nil {
		db = db.Where("login_ip = ?", seq.LoginIp)
	}
	if seq.LoginTotal != nil {
		db = db.Where("login_total = ?", seq.LoginTotal)
	}
	if seq.LoginTime != nil {
		db = db.Where("login_time = > ?", seq.LoginTime)
	}
	if seq.Scan != "" {
		db = db.Where("scan = ?", seq.Scan)
	}
	if seq.ScanTime != nil {
		db = db.Where("scan_time = > ?", seq.ScanTime)
	}
	if seq.Setting != "" {
		db = db.Where("setting = ?", seq.Setting)
	}
	if seq.RtcModel != nil {
		db = db.Where("rtc_model = ?", seq.RtcModel)
	}
	if seq.Status != nil {
		db = db.Where("status = ?", seq.Status)
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

// GetListAll 分页获取MemUser记录 (全部字段)
// Author 10512203@qq.com
func (m *MemUserSev) GetListAll(ctx context.Context, seq MemUserSearch, fields string) (list []MemUser, total int64, err error) {
	return m.GetList(ctx, seq, "*")
}

// GetByMap 根据Map获取单一记录
// Author  [linjd] 10512203@qq.com
func (m *MemUserSev) GetByMap(ctx context.Context, mapData map[string]interface{}, fields string) (obj MemUser, err error) {
	orderBy := "id desc"
	db := m.Db.WithContext(ctx).Model(&MemUser{})
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
func (m *MemUserSev) GetListByMap(ctx context.Context, mapData map[string]interface{}, fields string, orderBy string) (list []MemUser, err error) {
	if utils.IsEmptyStr(orderBy) {
		orderBy = "id desc"
	}
	db := m.Db.WithContext(ctx).Model(&MemUser{})
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
func (m *MemUserSev) Count(ctx context.Context, mapDataWhere map[string]interface{}) (total int64, err error) {
	db := m.Db.WithContext(ctx).Model(&MemUser{})
	err = db.Where(mapDataWhere).Count(&total).Error
	return total, err
}

// 如果有文件类型，更新文件 path
func (m *MemUserSev) getFileList(ctx context.Context, list []MemUser) {
	for i, v := range list {
		if !utils.IsEmpty(v.Avatar) {
			fileObjAvatar := FileObj{Guid: v.Avatar}
			fileObjAvatar.Path, _ = m.GetPathByGuid(ctx, v.Avatar)
			v.FileObjList = append(v.FileObjList, fileObjAvatar)
		}
		list[i] = v
	}
}

// 如果有文件类型，更新文件 path
func (m *MemUserSev) getFile(ctx context.Context, v *MemUser) {
	if v != nil {
		if !utils.IsEmpty(v.Avatar) {
			fileObjAvatar := FileObj{Guid: v.Avatar}
			fileObjAvatar.Path, _ = m.GetPathByGuid(ctx, v.Avatar)
			v.FileObjList = append(v.FileObjList, fileObjAvatar)
		}
	}
}
