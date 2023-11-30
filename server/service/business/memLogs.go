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

type MemLogsService struct {
	comSev.BaseService
}

var once_MemLogs sync.Once = sync.Once{}
var obj_MemLogsService *MemLogsService

// 获取单例
func GetMemLogsSev() *MemLogsService {
	once_MemLogs.Do(func() {
		obj_MemLogsService = new(MemLogsService)
		obj_MemLogsService.Db = global.DB
		//instanse.init()
	})
	return obj_MemLogsService
}

// Create 创建MemLogs记录
// Author 10512203@qq.com
func (m *MemLogsService) Create(ctx context.Context, data business.MemLogs) (id int64, err error) {
	err = m.Db.WithContext(ctx).Create(&data).Error
	if err != nil {
		return 0, err
	}
	return data.Id, err
}

// Delete 删除MemLogs记录
// Author 10512203@qq.com
func (m *MemLogsService) Delete(ctx context.Context, data business.MemLogs) (err error) {
	m.DelCache("mem_logs", []int64{data.Id})
	err = m.Db.WithContext(ctx).Delete(&data).Error
	return err
}

// DeleteByIds 批量删除MemLogs记录
// Author 10512203@qq.com
func (m *MemLogsService) DeleteByIds(ctx context.Context, ids request.IdsReq) (err error) {
	m.DelCache("mem_logs", ids.Ids)
	err = m.Db.WithContext(ctx).Delete(&[]business.MemLogs{}, "id in ?", ids.Ids).Error
	return err
}

// Update  更新MemLogs记录
// Author 10512203@qq.com
func (m *MemLogsService) Update(ctx context.Context, data business.MemLogs) (err error) {
	m.DelCache("mem_logs", []int64{data.Id})
	err = m.Db.WithContext(ctx).Save(&data).Error
	return err
}

// UpdateByMap  更新MemLogs记录 by Map
// Author 10512203@qq.com
func (m *MemLogsService) UpdateByMap(ctx context.Context, data business.MemLogs, mapData map[string]interface{}) (err error) {
	m.DelCache("mem_logs", []int64{data.Id})
	err = m.Db.WithContext(ctx).Model(&data).Updates(mapData).Error
	return err
}

// UpdateExpr 字段自增/自减
// Author 10512203@qq.com
func (m *MemLogsService) UpdateExpr(ctx context.Context, data business.MemLogs, column string, num int) (err error) {
	err = m.Db.WithContext(ctx).Model(&data).UpdateColumn(column, gorm.Expr(column+" + ?", num)).Error
	return err
}

// Get 根据id获取MemLogs记录
// Author 10512203@qq.com
func (m *MemLogsService) Get(ctx context.Context, id int64, fields string) (obj business.MemLogs, err error) {

	if id <= 0 {
		return obj, errors.New("参数id错误")
	}
	cacheKey := m.GetCacheKey("mem_logs", id)
	if cacheData, err := mycache.GetCache().Get(cacheKey); err == nil {
		obj = cacheData.(business.MemLogs)
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

// GetList 分页获取MemLogs记录
// Author 10512203@qq.com
func (m *MemLogsService) GetList(ctx context.Context, seq business.MemLogsSearch, fields string) (list []business.MemLogs, total int64, err error) {
	limit := seq.PageSize
	offset := seq.PageSize * (seq.Page - 1)
	//修改 by ljd  增加查询排序
	order := seq.OrderKey
	desc := seq.OrderDesc
	// 创建db
	db := m.Db.WithContext(ctx).Model(&business.MemLogs{})

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
	if seq.Type != nil {
		db = db.Where("type = ?", seq.Type)
	}
	if seq.Remark != "" {
		db = db.Where("remark = ?", seq.Remark)
	}
	if seq.Ip != nil {
		db = db.Where("ip = ?", seq.Ip)
	}
	if seq.IpAddr != "" {
		db = db.Where("ip_addr = ?", seq.IpAddr)
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

// GetListAll 分页获取MemLogs记录 (全部字段)
// Author 10512203@qq.com
func (m *MemLogsService) GetListAll(ctx context.Context, seq business.MemLogsSearch, fields string) (list []business.MemLogs, total int64, err error) {
	return m.GetList(ctx, seq, "*")
}

// GetByMap 根据Map获取单一记录
// Author  [linjd] 10512203@qq.com
func (m *MemLogsService) GetByMap(ctx context.Context, mapData map[string]interface{}, fields string) (obj business.MemLogs, err error) {
	orderBy := "id desc"
	db := m.Db.WithContext(ctx).Model(&business.MemLogs{})
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
func (m *MemLogsService) GetListByMap(ctx context.Context, mapData map[string]interface{}, fields string, orderBy string) (list []business.MemLogs, err error) {
	if utils.IsEmptyStr(orderBy) {
		orderBy = "id desc"
	}
	db := m.Db.WithContext(ctx).Model(&business.MemLogs{})
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
func (m *MemLogsService) Count(ctx context.Context, mapDataWhere map[string]interface{}) (total int64, err error) {
	db := m.Db.WithContext(ctx).Model(&business.MemLogs{})
	err = db.Where(mapDataWhere).Count(&total).Error
	return total, err
}

// 如果有文件类型，更新文件 path
func (m *MemLogsService) getFileList(ctx context.Context, list []business.MemLogs) {
	for i, v := range list {

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
func (m *MemLogsService) getFile(ctx context.Context, v *business.MemLogs) {
	if v != nil {

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
