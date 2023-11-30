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
	"strings"
	"sync"

	"gorm.io/gorm"
)

type SysMenuService struct {
	comSev.BaseService
}

var once_SysMenu sync.Once = sync.Once{}
var obj_SysMenuService *SysMenuService

// 获取单例
func GetSysMenuSev() *SysMenuService {
	once_SysMenu.Do(func() {
		obj_SysMenuService = new(SysMenuService)
		obj_SysMenuService.Db = global.DB
		//instanse.init()
	})
	return obj_SysMenuService
}

func (m *SysMenuService) GetMenu(ctx context.Context, menuList string) (list []business.SysMenu, err error) {
	db := m.Db.WithContext(ctx).Model(&business.SysMenu{})
	OrderStr := "sort asc"
	fields := "*" //"id,pid,name,path,title,level,hidden,path_vue,param,sort,keep_alive,icon"
	idList := strings.Split(menuList, ",")
	// 确保 不能让普通用户获取 系统菜单
	err = db.Select(fields).Where("status=1  AND id in ?", idList).Order(OrderStr).Find(&list).Error
	return list, err
}

// Create 创建SysMenu记录
// Author 10512203@qq.com
func (m *SysMenuService) Create(ctx context.Context, data business.SysMenu) (id int64, err error) {
	err = m.Db.WithContext(ctx).Create(&data).Error
	if err != nil {
		return 0, err
	}
	return data.Id, err
}

// Delete 删除SysMenu记录
// Author 10512203@qq.com
func (m *SysMenuService) Delete(ctx context.Context, data business.SysMenu) (err error) {
	m.DelCache("sys_menu", []int64{data.Id})
	err = m.Db.WithContext(ctx).Delete(&data).Error
	return err
}

// DeleteByIds 批量删除SysMenu记录
// Author 10512203@qq.com
func (m *SysMenuService) DeleteByIds(ctx context.Context, ids request.IdsReq) (err error) {
	m.DelCache("sys_menu", ids.Ids)
	err = m.Db.WithContext(ctx).Delete(&[]business.SysMenu{}, "id in ?", ids.Ids).Error
	return err
}

// Update  更新SysMenu记录
// Author 10512203@qq.com
func (m *SysMenuService) Update(ctx context.Context, data business.SysMenu) (err error) {
	m.DelCache("sys_menu", []int64{data.Id})

	err = m.Db.WithContext(ctx).Save(&data).Error
	return err
}

// UpdateByMap  更新SysMenu记录 by Map
// Author 10512203@qq.com
func (m *SysMenuService) UpdateByMap(ctx context.Context, data business.SysMenu, mapData map[string]interface{}) (err error) {
	m.DelCache("sys_menu", []int64{data.Id})
	err = m.Db.WithContext(ctx).Model(&data).Updates(mapData).Error
	return err
}

// UpdateExpr 字段自增/自减
// Author 10512203@qq.com
func (m *SysMenuService) UpdateExpr(ctx context.Context, data business.SysMenu, column string, num int) (err error) {
	err = m.Db.WithContext(ctx).Model(&data).UpdateColumn(column, gorm.Expr(column+" + ?", num)).Error
	return err
}

// Get 根据id获取SysMenu记录
// Author 10512203@qq.com
func (m *SysMenuService) Get(ctx context.Context, id int64, fields string) (obj business.SysMenu, err error) {

	if id <= 0 {
		return obj, errors.New("参数id错误")
	}
	cacheKey := m.GetCacheKey("sys_menu", id)
	if cacheData, err := mycache.GetCache().Get(cacheKey); err == nil {
		obj = cacheData.(business.SysMenu)
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

// GetList 分页获取SysMenu记录
// Author 10512203@qq.com
func (m *SysMenuService) GetList(ctx context.Context, seq business.SysMenuSearch, fields string) (list []*business.SysMenu, total int64, err error) {
	limit := seq.PageSize
	if limit == 0 {
		limit = 1000
	}
	offset := seq.PageSize * (seq.Page - 1)

	//修改 by ljd  增加查询排序
	order := seq.OrderKey
	desc := seq.OrderDesc
	// 创建db
	db := m.Db.WithContext(ctx).Model(&business.SysMenu{})

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
	if seq.Name != "" {
		db = db.Where("name = ?", seq.Name)
	}
	if seq.Path != "" {
		db = db.Where("path = ?", seq.Path)
	}
	if seq.Title != "" {
		db = db.Where("title   LIKE ?", "%"+seq.Title+"%")
	}
	if seq.Level != nil {
		db = db.Where("level = ?", seq.Level)
	}
	if seq.Auths != "" {
		db = db.Where("auths = ?", seq.Auths)
	}
	if seq.ShowLink != nil {
		db = db.Where("show_link = ?", seq.ShowLink)
	}
	if seq.ActivePath != "" {
		db = db.Where("active_path = ?", seq.ActivePath)
	}
	if seq.FrameSrc != "" {
		db = db.Where("frame_src = ?", seq.FrameSrc)
	}
	if seq.Component != "" {
		db = db.Where("component = ?", seq.Component)
	}
	if seq.Param != "" {
		db = db.Where("param = ?", seq.Param)
	}
	if seq.Sort != nil {
		db = db.Where("sort = ?", seq.Sort)
	}
	if seq.KeepAlive != nil {
		db = db.Where("keep_alive = ?", seq.KeepAlive)
	}
	if seq.Icon != "" {
		db = db.Where("icon = ?", seq.Icon)
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
	//m.getFileList(ctx, list)
	return list, total, err
}

// GetListAll 分页获取SysMenu记录 (全部字段)
// Author 10512203@qq.com
func (m *SysMenuService) GetListAll(ctx context.Context, seq business.SysMenuSearch, fields string) (list []*business.SysMenu, total int64, err error) {
	return m.GetList(ctx, seq, "*")
}

// GetByMap 根据Map获取单一记录
// Author  [linjd] 10512203@qq.com
func (m *SysMenuService) GetByMap(ctx context.Context, mapData map[string]interface{}, fields string) (obj business.SysMenu, err error) {
	orderBy := "id desc"
	db := m.Db.WithContext(ctx).Model(&business.SysMenu{})
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
func (m *SysMenuService) GetListByMap(ctx context.Context, mapData map[string]interface{}, fields string, orderBy string) (list []business.SysMenu, err error) {
	if utils.IsEmptyStr(orderBy) {
		orderBy = "id desc"
	}
	db := m.Db.WithContext(ctx).Model(&business.SysMenu{})
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
func (m *SysMenuService) Count(ctx context.Context, mapDataWhere map[string]interface{}) (total int64, err error) {
	db := m.Db.WithContext(ctx).Model(&business.SysMenu{})
	err = db.Where(mapDataWhere).Count(&total).Error
	return total, err
}

// 如果有文件类型，更新文件 path
func (m *SysMenuService) getFileList(ctx context.Context, list []business.SysMenu) {
	// for i, v := range list {

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
	// 	list[i] = v
	// }
}

// 如果有文件类型，更新文件 path
func (m *SysMenuService) getFile(ctx context.Context, v *business.SysMenu) {
	if v != nil {

		// if !utils.IsEmpty(v.FileList) {
		// 	objList := strings.Split(v.FileList, ",")
		// 	for _, guid := range objList {
		// 		if !utils.IsEmptyStr(guid) {
		// 			fileObj := global.FileObj{Guid: guid}
		// 			fileObj.Path, _ = m.GetPathByGuid(ctx, guid)
		// 			v.FileObjList = append(v.FileObjList, fileObj)
		// 		}
		// 	}
		// }
	}
}
