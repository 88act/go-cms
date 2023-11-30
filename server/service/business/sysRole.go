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

type SysRoleService struct {
	comSev.BaseService
}

var once_SysRole sync.Once = sync.Once{}
var obj_SysRoleService *SysRoleService

// 获取单例
func GetSysRoleSev() *SysRoleService {
	once_SysRole.Do(func() {
		obj_SysRoleService = new(SysRoleService)
		obj_SysRoleService.Db = global.DB
		//instanse.init()
	})
	return obj_SysRoleService
}

// Create 创建SysRole记录
// Author 10512203@qq.com
func (m *SysRoleService) Create(ctx context.Context, data business.SysRole) (id int64, err error) {
	err = m.Db.WithContext(ctx).Create(&data).Error
	if err != nil {
		return 0, err
	}
	return data.Id, err
}

// Delete 删除SysRole记录
// Author 10512203@qq.com
func (m *SysRoleService) Delete(ctx context.Context, data business.SysRole) (err error) {
	m.DelCache("sys_role", []int64{data.Id})
	err = m.Db.WithContext(ctx).Delete(&data).Error
	return err
}

// DeleteByIds 批量删除SysRole记录
// Author 10512203@qq.com
func (m *SysRoleService) DeleteByIds(ctx context.Context, ids request.IdsReq) (err error) {
	m.DelCache("sys_role", ids.Ids)
	err = m.Db.WithContext(ctx).Delete(&[]business.SysRole{}, "id in ?", ids.Ids).Error
	return err
}

// Update  更新SysRole记录
// Author 10512203@qq.com
func (m *SysRoleService) Update(ctx context.Context, data business.SysRole) (err error) {
	m.DelCache("sys_role", []int64{data.Id})
	err = m.Db.WithContext(ctx).Save(&data).Error
	return err
}

// UpdateByMap  更新SysRole记录 by Map
// Author 10512203@qq.com
func (m *SysRoleService) UpdateByMap(ctx context.Context, data business.SysRole, mapData map[string]interface{}) (err error) {
	m.DelCache("sys_role", []int64{data.Id})
	err = m.Db.WithContext(ctx).Model(&data).Updates(mapData).Error
	return err
}

// UpdateExpr 字段自增/自减
// Author 10512203@qq.com
func (m *SysRoleService) UpdateExpr(ctx context.Context, data business.SysRole, column string, num int) (err error) {
	err = m.Db.WithContext(ctx).Model(&data).UpdateColumn(column, gorm.Expr(column+" + ?", num)).Error
	return err
}

// Get 根据id获取SysRole记录
// Author 10512203@qq.com
func (m *SysRoleService) Get(ctx context.Context, id int64, fields string) (obj business.SysRole, err error) {

	if id <= 0 {
		return obj, errors.New("参数id错误")
	}
	cacheKey := m.GetCacheKey("sys_role", id)
	if cacheData, err := mycache.GetCache().Get(cacheKey); err == nil {
		obj = cacheData.(business.SysRole)
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

// GetList 分页获取SysRole记录
// Author 10512203@qq.com
func (m *SysRoleService) GetList(ctx context.Context, seq business.SysRoleSearch, fields string) (list []business.SysRole, total int64, err error) {
	limit := seq.PageSize
	offset := seq.PageSize * (seq.Page - 1)
	//修改 by ljd  增加查询排序
	order := seq.OrderKey
	desc := seq.OrderDesc
	// 创建db
	db := m.Db.WithContext(ctx).Model(&business.SysRole{})

	//修改 by ljd
	if seq.Id > 0 {
		db = db.Where("id = ?", seq.Id)
	}

	if !utils.IsEmpty(seq.CreatedAtBegin) && !utils.IsEmpty(seq.CreatedAtEnd) {
		db = db.Where("created_at BETWEEN ? AND ?", seq.CreatedAtBegin, seq.CreatedAtEnd)
	}

	if seq.Pid != nil {
		db = db.Where("pid = ?", seq.Pid)
	}
	if seq.RoleName != "" {
		db = db.Where("role_name  LIKE ?", "%"+seq.RoleName+"%")
	}
	if seq.RoleCode != "" {
		db = db.Where("role_code = ?", seq.RoleCode)
	}
	if seq.DefaultMenu != nil {
		db = db.Where("default_menu = ?", seq.DefaultMenu)
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

// GetListAll 分页获取SysRole记录 (全部字段)
// Author 10512203@qq.com
func (m *SysRoleService) GetListAll(ctx context.Context, seq business.SysRoleSearch, fields string) (list []business.SysRole, total int64, err error) {
	return m.GetList(ctx, seq, "*")
}

// GetByMap 根据Map获取单一记录
// Author  [linjd] 10512203@qq.com
func (m *SysRoleService) GetByMap(ctx context.Context, mapData map[string]interface{}, fields string) (obj business.SysRole, err error) {
	orderBy := "id desc"
	db := m.Db.WithContext(ctx).Model(&business.SysRole{})
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
func (m *SysRoleService) GetListByMap(ctx context.Context, mapData map[string]interface{}, fields string, orderBy string) (list []business.SysRole, err error) {
	if utils.IsEmptyStr(orderBy) {
		orderBy = "id desc"
	}
	db := m.Db.WithContext(ctx).Model(&business.SysRole{})
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
func (m *SysRoleService) Count(ctx context.Context, mapDataWhere map[string]interface{}) (total int64, err error) {
	db := m.Db.WithContext(ctx).Model(&business.SysRole{})
	err = db.Where(mapDataWhere).Count(&total).Error
	return total, err
}

// 如果有文件类型，更新文件 path
func (m *SysRoleService) getFileList(ctx context.Context, list []business.SysRole) {
	// for i, v := range list {

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
	// 	list[i] = v
	// }
}

// 如果有文件类型，更新文件 path
func (m *SysRoleService) getFile(ctx context.Context, v *business.SysRole) {
	// if v != nil {

	// 	if !utils.IsEmpty(v.FileList) {
	// 		objList := strings.Split(v.FileList, ",")
	// 		for _, guid := range objList {
	// 			if !utils.IsEmptyStr(guid) {
	// 				fileObj := global.FileObj{Guid: guid}
	// 				fileObj.Path, _ = m.GetPathByGuid(ctx, guid)
	// 				v.FileObjList = append(v.FileObjList, fileObj)
	// 			}
	// 		}
	// 	}
	// }
}
