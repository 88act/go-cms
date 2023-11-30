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

type SysRoleMenuService struct {
	comSev.BaseService
}

var once_SysRoleMenu sync.Once = sync.Once{}
var obj_SysRoleMenuService *SysRoleMenuService

// 获取单例
func GetSysRoleMenuSev() *SysRoleMenuService {
	once_SysRoleMenu.Do(func() {
		obj_SysRoleMenuService = new(SysRoleMenuService)
		obj_SysRoleMenuService.Db = global.DB
		//instanse.init()
	})
	return obj_SysRoleMenuService
}

// Create 创建SysRoleMenu记录
// Author 10512203@qq.com
func (m *SysRoleMenuService) Create(ctx context.Context, data business.SysRoleMenu) (id int64, err error) {
	err = m.Db.WithContext(ctx).Create(&data).Error
	if err != nil {
		return 0, err
	}
	return data.Id, err
}

// Delete 删除SysRoleMenu记录
// Author 10512203@qq.com
func (m *SysRoleMenuService) Delete(ctx context.Context, data business.SysRoleMenu) (err error) {
	m.DelCache("sys_role_menu", []int64{data.Id})
	err = m.Db.WithContext(ctx).Delete(&data).Error
	return err
}

// DeleteByIds 批量删除SysRoleMenu记录
// Author 10512203@qq.com
func (m *SysRoleMenuService) DeleteByIds(ctx context.Context, ids request.IdsReq) (err error) {
	m.DelCache("sys_role_menu", ids.Ids)
	err = m.Db.WithContext(ctx).Delete(&[]business.SysRoleMenu{}, "id in ?", ids.Ids).Error
	return err
}

// Update  更新SysRoleMenu记录
// Author 10512203@qq.com
func (m *SysRoleMenuService) Update(ctx context.Context, data business.SysRoleMenu) (err error) {
	m.DelCache("sys_role_menu", []int64{data.Id})
	err = m.Db.WithContext(ctx).Save(&data).Error
	return err
}

// UpdateByMap  更新SysRoleMenu记录 by Map
// Author 10512203@qq.com
func (m *SysRoleMenuService) UpdateByMap(ctx context.Context, data business.SysRoleMenu, mapData map[string]interface{}) (err error) {
	m.DelCache("sys_role_menu", []int64{data.Id})
	err = m.Db.WithContext(ctx).Model(&data).Updates(mapData).Error
	return err
}

// UpdateExpr 字段自增/自减
// Author 10512203@qq.com
func (m *SysRoleMenuService) UpdateExpr(ctx context.Context, data business.SysRoleMenu, column string, num int) (err error) {
	err = m.Db.WithContext(ctx).Model(&data).UpdateColumn(column, gorm.Expr(column+" + ?", num)).Error
	return err
}

// Get 根据id获取SysRoleMenu记录
// Author 10512203@qq.com
func (m *SysRoleMenuService) Get(ctx context.Context, id int64, fields string) (obj business.SysRoleMenu, err error) {

	if id <= 0 {
		return obj, errors.New("参数id错误")
	}
	cacheKey := m.GetCacheKey("sys_role_menu", id)
	if cacheData, err := mycache.GetCache().Get(cacheKey); err == nil {
		obj = cacheData.(business.SysRoleMenu)
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

// GetList 分页获取SysRoleMenu记录
// Author 10512203@qq.com
func (m *SysRoleMenuService) GetList(ctx context.Context, seq business.SysRoleMenuSearch, fields string) (list []business.SysRoleMenu, total int64, err error) {
	limit := seq.PageSize
	offset := seq.PageSize * (seq.Page - 1)
	//修改 by ljd  增加查询排序
	order := seq.OrderKey
	desc := seq.OrderDesc
	// 创建db
	db := m.Db.WithContext(ctx).Model(&business.SysRoleMenu{})

	//修改 by ljd
	if seq.Id > 0 {
		db = db.Where("id = ?", seq.Id)
	}

	if !utils.IsEmpty(seq.CreatedAtBegin) && !utils.IsEmpty(seq.CreatedAtEnd) {
		db = db.Where("created_at BETWEEN ? AND ?", seq.CreatedAtBegin, seq.CreatedAtEnd)
	}

	// 如果有条件搜索 下方会自动创建搜索语句
	if seq.RoleId != nil {
		db = db.Where("role_id = ?", seq.RoleId)
	}
	if seq.MenuId != nil {
		db = db.Where("menu_id = ?", seq.MenuId)
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

// GetListAll 分页获取SysRoleMenu记录 (全部字段)
// Author 10512203@qq.com
func (m *SysRoleMenuService) GetListAll(ctx context.Context, seq business.SysRoleMenuSearch, fields string) (list []business.SysRoleMenu, total int64, err error) {
	return m.GetList(ctx, seq, "*")
}

// GetByMap 根据Map获取单一记录
// Author  [linjd] 10512203@qq.com
func (m *SysRoleMenuService) GetByMap(ctx context.Context, mapData map[string]interface{}, fields string) (obj business.SysRoleMenu, err error) {
	orderBy := "id desc"
	db := m.Db.WithContext(ctx).Model(&business.SysRoleMenu{})
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
func (m *SysRoleMenuService) GetListByMap(ctx context.Context, mapData map[string]interface{}, fields string, orderBy string) (list []business.SysRoleMenu, err error) {
	if utils.IsEmptyStr(orderBy) {
		orderBy = "id desc"
	}
	db := m.Db.WithContext(ctx).Model(&business.SysRoleMenu{})
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
func (m *SysRoleMenuService) Count(ctx context.Context, mapDataWhere map[string]interface{}) (total int64, err error) {
	db := m.Db.WithContext(ctx).Model(&business.SysRoleMenu{})
	err = db.Where(mapDataWhere).Count(&total).Error
	return total, err
}

// 如果有文件类型，更新文件 path
func (m *SysRoleMenuService) getFileList(ctx context.Context, list []business.SysRoleMenu) {
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
func (m *SysRoleMenuService) getFile(ctx context.Context, v *business.SysRoleMenu) {
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
