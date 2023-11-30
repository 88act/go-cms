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

type MemUserService struct {
	comSev.BaseService
}

var once_MemUser sync.Once = sync.Once{}
var obj_MemUserService *MemUserService

// 获取单例
func GetMemUserSev() *MemUserService {
	once_MemUser.Do(func() {
		obj_MemUserService = new(MemUserService)
		obj_MemUserService.Db = global.DB
		//instanse.init()
	})
	return obj_MemUserService
}

//	检查账号是否存在
//
// Author 10512203@qq.com
func (m *MemUserService) ChechUserExist(ctx context.Context, username string, userId int64) (exist bool, err error) {

	mapData := make(map[string]interface{})
	mapData["username"] = username
	data, err := m.GetByMap(ctx, mapData, "")
	if err == nil && data.Id > 0 {
		return true, errors.New("员工已存在")
	}
	// data2, err2 := GetMemUserSev().GetByMap(ctx, mapData, "")
	// if err2 == nil && data2.Id > 0 {
	// 	return true, errors.New("系统用户已存在")
	// }

	return false, nil
}

//	改密码
//
// Author 10512203@qq.com
func (m *MemUserService) ChangePasswordCu(ctx context.Context, userCPW business.ChangePasswordStruct, cuUserId int64) (err error) {
	memUser, err := m.Get(ctx, userCPW.UserId, "")
	if err != nil {
		return errors.New("用户不存在")
	}

	memUser.PasswordSlat = utils.Krand(8, utils.KC_RAND_KIND_ALL)
	//global.LOG.Error("user.PasswordSlat = " + user.PasswordSlat)
	pwd := userCPW.NewPassword + memUser.PasswordSlat
	mapData := make(map[string]interface{})
	mapData["password"] = utils.Md5ByString(pwd)
	mapData["password_slat"] = memUser.PasswordSlat
	//err = m.UpdateByMap(ctx, memUser, mapData)
	err = m.Db.Table("mem_user").Where("id=?", memUser.Id).Updates(mapData).Error
	return err

}

//	改密码
//
// Author 10512203@qq.com
func (m *MemUserService) ChangePassword(ctx context.Context, userCPW business.ChangePasswordStruct, bePlatAdmin bool) (err error) {
	memUser, err := m.Get(ctx, userCPW.UserId, "")
	if err != nil {
		return errors.New("用户不存在")
	}
	if !bePlatAdmin {
		pwd := userCPW.Password + memUser.PasswordSlat
		if memUser.Password != utils.Md5ByString(pwd) {
			return errors.New("原密码错误")
		}
	}
	memUser.PasswordSlat = utils.Krand(8, utils.KC_RAND_KIND_ALL)
	//global.LOG.Error("user.PasswordSlat = " + user.PasswordSlat)
	pwd := userCPW.NewPassword + memUser.PasswordSlat
	mapData := make(map[string]interface{})
	mapData["password"] = utils.Md5ByString(pwd)
	mapData["password_slat"] = memUser.PasswordSlat
	err = m.Db.Table("mem_user").Where("id=?", memUser.Id).Updates(mapData).Error
	//err = m.UpdateByMap(ctx, memUser, mapData)
	return err

}

// Get 根据username获取
// Author 10512203@qq.com
func (m *MemUserService) GetByUsername(ctx context.Context, username string, fields string) (obj business.MemUser, err error) {
	// cacheKey := m.GetCacheKey("MemUser_name", username)
	// if cacheData, err := mycache.GetCache().Get(cacheKey); err == nil {
	// 	obj = cacheData.(business.MemUser)
	// 	return obj, err
	// }
	if utils.IsEmptyStr(fields) {
		err = m.Db.WithContext(ctx).Where("  username = ?", username).First(&obj).Error
	} else {
		err = m.Db.WithContext(ctx).Select(fields).Where("  username = ?", username).First(&obj).Error
	}
	if err == nil {
		m.getFile(ctx, &obj)
		//mycache.GetCache().Set(cacheKey, obj, "")
	}

	return obj, err
}

func (m *MemUserService) GetByOpenId(ctx context.Context, id string, fields string) (obj business.MemUser, err error) {
	cacheKey := m.GetCacheKey("MemUser_open", id)
	if cacheData, err := mycache.GetCache().Get(cacheKey); err == nil {
		obj = cacheData.(business.MemUser)
		return obj, err
	}
	if utils.IsEmptyStr(fields) {
		err = m.Db.WithContext(ctx).Where("openid = ?", id).First(&obj).Error
	} else {
		err = m.Db.WithContext(ctx).Select(fields).Where("openid = ?", id).First(&obj).Error
	}
	if err == nil {
		m.getFile(ctx, &obj)
		mycache.GetCache().Set(cacheKey, obj, "")
	}
	return obj, err
}

// Create 创建MemUser记录
// Author 10512203@qq.com
func (m *MemUserService) Create(ctx context.Context, data business.MemUser) (id int64, err error) {
	data.Guid = utils.GUID()
	err = m.Db.WithContext(ctx).Create(&data).Error
	if err != nil {
		return 0, err
	}
	return data.Id, err
}

// Delete 删除MemUser记录
// Author 10512203@qq.com
func (m *MemUserService) Delete(ctx context.Context, data business.MemUser) (err error) {
	m.DelCache("mem_user", []int64{data.Id})
	err = m.Db.WithContext(ctx).Delete(&data).Error
	return err
}

// DeleteByIds 批量删除MemUser记录
// Author 10512203@qq.com
func (m *MemUserService) DeleteByIds(ctx context.Context, ids request.IdsReq) (err error) {
	m.DelCache("mem_user", ids.Ids)
	err = m.Db.WithContext(ctx).Delete(&[]business.MemUser{}, "id in ?", ids.Ids).Error
	return err
}

// Update  更新MemUser记录
// Author 10512203@qq.com
func (m *MemUserService) Update(ctx context.Context, data business.MemUser) (err error) {
	m.DelCache("mem_user", []int64{data.Id})
	err = m.Db.WithContext(ctx).Save(&data).Error
	return err
}

// UpdateByMap  更新MemUser记录 by Map
// Author 10512203@qq.com
func (m *MemUserService) UpdateByMap(ctx context.Context, data business.MemUser, mapData map[string]interface{}) (err error) {
	m.DelCache("mem_user", []int64{data.Id})
	err = m.Db.WithContext(ctx).Model(&data).Updates(mapData).Error
	return err
}

// UpdateByMap  更新MemUser记录 by Map
// Author 10512203@qq.com
func (m *MemUserService) UpdateByMap2(ctx context.Context, mapWhere map[string]interface{}, mapData map[string]interface{}) (err error) {
	//m.DelCache("mem_user", []int64{data.Id})
	err = m.Db.WithContext(ctx).Model(business.MemUser{}).Where(mapWhere).Updates(mapData).Error
	return err
}

// UpdateExpr 字段自增/自减
// Author 10512203@qq.com
func (m *MemUserService) UpdateExpr(ctx context.Context, data business.MemUser, column string, num int) (err error) {
	err = m.Db.WithContext(ctx).Model(&data).UpdateColumn(column, gorm.Expr(column+" + ?", num)).Error
	return err
}

// Get 根据id获取MemUser记录
// Author 10512203@qq.com
func (m *MemUserService) Get(ctx context.Context, id int64, fields string) (obj business.MemUser, err error) {

	if id <= 0 {
		return obj, errors.New("参数id错误")
	}
	cacheKey := m.GetCacheKey("mem_user", id)
	if cacheData, err := mycache.GetCache().Get(cacheKey); err == nil {
		obj = cacheData.(business.MemUser)
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

// GetList 分页获取MemUser记录
// Author 10512203@qq.com
func (m *MemUserService) GetList(ctx context.Context, seq business.MemUserSearch, fields string) (list []business.MemUser, total int64, err error) {
	limit := seq.PageSize
	offset := seq.PageSize * (seq.Page - 1)
	//修改 by ljd  增加查询排序
	order := seq.OrderKey
	desc := seq.OrderDesc
	// 创建db
	db := m.Db.WithContext(ctx).Model(&business.MemUser{})

	//修改 by ljd
	if seq.Id > 0 {
		db = db.Where("id = ?", seq.Id)
	}

	if !utils.IsEmpty(seq.CreatedAtBegin) && !utils.IsEmpty(seq.CreatedAtEnd) {
		db = db.Where("created_at BETWEEN ? AND ?", seq.CreatedAtBegin, seq.CreatedAtEnd)
	}

	if seq.UserId != nil {
		db = db.Where("user_id = ?", seq.UserId)
	}
	if seq.Guid != "" {
		db = db.Where("guid = ?", seq.Guid)
	}
	if seq.Username != "" {
		db = db.Where("username =  LIKE ?", "%"+seq.Username+"%")
	}
	if seq.Password != "" {
		db = db.Where("password = ?", seq.Password)
	}
	if seq.PasswordSlat != "" {
		db = db.Where("password_slat = ?", seq.PasswordSlat)
	}
	if seq.Nickname != "" {
		db = db.Where("nickname  LIKE ?", "%"+seq.Nickname+"%")
	}
	if seq.Realname != "" {
		db = db.Where("realname  LIKE ?", "%"+seq.Realname+"%")
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
	if utils.IsEmptyStr(fields) {
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&list).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&list).Error
	}
	m.getFileList(ctx, list)
	return list, total, err
}

// GetListAll 分页获取MemUser记录 (全部字段)
// Author 10512203@qq.com
func (m *MemUserService) GetListAll(ctx context.Context, seq business.MemUserSearch, fields string) (list []business.MemUser, total int64, err error) {
	return m.GetList(ctx, seq, "*")
}

// GetByMap 根据Map获取单一记录
// Author  [linjd] 10512203@qq.com
func (m *MemUserService) GetByMap(ctx context.Context, mapData map[string]interface{}, fields string) (obj business.MemUser, err error) {
	orderBy := "id desc"
	db := m.Db.WithContext(ctx).Model(&business.MemUser{})
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
func (m *MemUserService) GetListByMap(ctx context.Context, mapData map[string]interface{}, fields string, orderBy string) (list []business.MemUser, err error) {
	if utils.IsEmptyStr(orderBy) {
		orderBy = "id desc"
	}
	db := m.Db.WithContext(ctx).Model(&business.MemUser{})
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
func (m *MemUserService) Count(ctx context.Context, mapDataWhere map[string]interface{}) (total int64, err error) {
	db := m.Db.WithContext(ctx).Model(&business.MemUser{})
	err = db.Where(mapDataWhere).Count(&total).Error
	return total, err
}

// 如果有文件类型，更新文件 path
func (m *MemUserService) getFileList(ctx context.Context, list []business.MemUser) {
	for i, v := range list {
		if !utils.IsEmpty(v.Avatar) {
			fileObjAvatar := global.FileObj{Guid: v.Avatar}
			fileObjAvatar.Path, _ = m.GetPathByGuid(ctx, v.Avatar)
			v.FileObjList = append(v.FileObjList, fileObjAvatar)
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
func (m *MemUserService) getFile(ctx context.Context, v *business.MemUser) {
	if v != nil {
		if !utils.IsEmpty(v.Avatar) {
			fileObjAvatar := global.FileObj{Guid: v.Avatar}
			fileObjAvatar.Path, _ = m.GetPathByGuid(ctx, v.Avatar)
			v.FileObjList = append(v.FileObjList, fileObjAvatar)
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
