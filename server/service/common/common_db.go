package common

import (
	"context"
	"errors"
	"fmt"
	"go-cms/global"
	"go-cms/model/business"
	"go-cms/model/common/request"
	"go-cms/mycache"

	//bizSev "go-cms/service/business"
	"go-cms/utils"
	"sync"

	"github.com/gogf/gf/util/gconv"
)

type CommonDbService struct {
	BaseService
}

var once_CommonDb sync.Once = sync.Once{}
var obj_CommonDbService *CommonDbService

// 获取单例
func GetCommonDbSev() *CommonDbService {
	once_CommonDb.Do(func() {
		obj_CommonDbService = new(CommonDbService)
		//instanse.init()
	})
	return obj_CommonDbService
}

// QuickEdit 更新QuickEdit记录
// Author [Linjd]
// QuickEdit 更新QuickEdit记录
// Author [Linjd]
func (m *CommonDbService) QuickEdit(ctx context.Context, quickEdit request.QuickEdit) (err error) {
	// 清缓存
	m.DelCache(quickEdit.Table, []int64{quickEdit.Id})
	if quickEdit.Id == 0 {
		return errors.New("id不能为空")
	}
	err = global.DB.WithContext(ctx).Exec("UPDATE `"+quickEdit.Table+"` SET `"+quickEdit.Field+"`=? WHERE `id` = ?", quickEdit.Value, quickEdit.Id).Error
	//err = global.DB.(&quickEdit).Error
	return err
}

// 获取父子关系的数据，任意table表可适用
// Author [Linjd]
func (m *CommonDbService) GetPidData(ctx context.Context, req request.PidDataReq) (list []request.PidData, err error) {
	if !utils.IsEmpty(req.Table) {
		whereSql := "status=1 AND deleted_at IS NULL "
		if req.Table == "basic_area" {
			whereSql = " 1=1"
		}
		if req.Table == "ims_job" || req.Table == "ims_task" || req.Table == "mem_user" {
			whereSql += fmt.Sprintf(" AND cu_id= %d", *req.CuId)
		}
		if !utils.IsEmpty(req.Where) {
			whereSql += " AND " + req.Where
		}
		if req.PidValue != nil {
			whereSql += fmt.Sprintf(" AND  `"+req.PidField+"` = %v", req.PidValue)
		}

		global.DB.WithContext(ctx).Raw("select " + req.PidField + " as value," + req.NameField + " as label FROM  `" + req.Table + "`  WHERE " + whereSql).Scan(&list)

		return list, nil
	}
	return list, errors.New("参数错误")
	//err = global.DB.(&quickEdit).Error
}

// 更新表的 pid 和 sort
func (m *CommonDbService) UpdatePidSort(ctx context.Context, req request.UpdatePidSortReq) (err error) {
	if utils.IsEmptyStr(req.Table) || req.Id == 0 || req.Id2 == 0 || utils.IsEmptyStr(req.DropType) {
		return errors.New("参数错误")
	}
	db := global.DB.WithContext(ctx)

	//前后换顺序
	if req.DropType == "before" || req.DropType == "after" { //before、after、inner
		// 第一个

		upData := make(map[string]interface{})
		// 第二个
		upData2 := make(map[string]interface{})
		if req.DropType == "before" {
			//fmt.Print("before-----")
			//upData["sort"] = utils.Ternary(req.Sort>req.Sort2,req.Sort,req.Sort2).(int)
			upData["sort"] = utils.Ternary(req.Sort > req.Sort2, req.Sort2, req.Sort)
			upData2["sort"] = utils.Ternary(req.Sort > req.Sort2, req.Sort, req.Sort2)
		} else if req.DropType == "after" {
			//fmt.Print("after-----")
			upData["sort"] = utils.Ternary(req.Sort > req.Sort2, req.Sort, req.Sort2)
			upData2["sort"] = utils.Ternary(req.Sort > req.Sort2, req.Sort2, req.Sort)
		}
		//if req.Pid2 > 0 {
		// 同时要更改 父子关系
		upData["pid"] = req.Pid2
		//}
		result := db.Table(req.Table).Where("id", req.Id).Updates(upData)
		fmt.Print("11111111111111111-11=")
		fmt.Print(result.RowsAffected)
		//fmt.Print(result.Error.Error())
		//if result.RowsAffected <= 0 {
		//	return errors.New("更新失败1")
		//}

		result = db.Table(req.Table).Where("id", req.Id2).Updates(upData2)
		fmt.Print("222222222222-2222222=")
		fmt.Print(result.RowsAffected)
		//fmt.Print(result.Error.Error())
		//if result.RowsAffected <= 0 {
		//	return errors.New("更新失败2")
		//}

	} else if req.DropType == "inner" {
		// 父子关系
		upData := map[string]interface{}{"pid": req.Id2}
		result := db.Table(req.Table).Where("id", req.Id).Updates(upData)
		if result.RowsAffected <= 0 {
			return errors.New("更新失败3")
		}
	}
	return nil

}

// GetTreeData_jobTask 获取树状数据（ ImsJobTask ims_task ）
// Author [Linjd]
func (m *CommonDbService) GetTreeData_jobTask(ctx context.Context, req request.PidDataReq) (list []*request.PidTreeData) {

	// db := global.DB.WithContext(ctx).Model(&business.ImsJobTask{})
	// db = db.Joins("left join ims_task on ims_job_task.task_id = ims_task.id")
	// OrderStr := "ims_job_task.sort desc"
	// fields := "ims_job_task.id as `value` ,ims_job_task.pid ,ims_task.title as label"

	// err := db.Select(fields).Where("ims_task.status=1 and ims_task.deleted_at IS NULL and ims_job_task.status=1 and ims_job_task.deleted_at IS NULL").Where("ims_job_task.job_id", req.PidValue).Order(OrderStr).Scan(&list).Error
	// if err == nil {
	// 	list2 := m.getTreeList(list, 0)
	// 	return list2
	// }

	return nil
	//err = global.DB.(&quickEdit).Error
}

// GetTreeData 获取树状数据  memBranch
// Author [Linjd]
func (m *CommonDbService) GetTreeData_memDepart(ctx context.Context, req request.PidDataReq) (list []*request.PidTreeData) {

	OrderStr := "sort asc"
	fields := "id as `value`,name as label,pid"
	db := global.DB.WithContext(ctx).Model(&business.MemDepart{})
	err := db.Select(fields).Where("cu_id =? and  deleted_at IS NULL", *req.CuId).Order(OrderStr).Scan(&list).Error
	if err == nil {
		list2 := m.getTreeList(list, 0)
		return list2
	}
	return nil
}

func (m *CommonDbService) GetTreeData_devCustomer(ctx context.Context, req request.PidDataReq) (list []*request.PidTreeData) {

	// OrderStr := "id desc"
	// fields := "id as `value`,name as label,pid"
	// db := global.DB.WithContext(ctx).Model(&business.DevCustomer{})
	// err := db.Select(fields).Where("status=1 and  deleted_at IS NULL").Order(OrderStr).Scan(&list).Error
	// if err == nil {
	// 	list2 := m.getTreeList(list, 0)
	// 	return list2
	// }
	return nil
}

func (m *CommonDbService) GetTreeData_basicFileCat(ctx context.Context, req request.PidDataReq) (list []*request.PidTreeData) {

	OrderStr := "sort asc"
	fields := "id as `value`,name as label,pid"
	db := global.DB.WithContext(ctx).Model(&business.BasicFileCat{})
	err := db.Select(fields).Where("status=1 and  deleted_at IS NULL").Order(OrderStr).Scan(&list).Error
	if err == nil {
		list2 := m.getTreeList(list, 0)
		return list2
	}
	return nil
}

// func (m *CommonDbService) GetTreeData3(ctx context.Context, req request.PidDataReq) (list []*request.PidTreeData) {

// 	OrderStr := req.PidField + " asc"
// 	fields := ""
// 	if utils.IsEmptyStr(req.Pid) {
// 		fields = req.PidField + " as `value`," + req.NameField + " as label,pid"
// 	} else {
// 		fields = req.PidField + " as `value`," + req.NameField + " as label," + req.Pid + " as pid"
// 	}
// 	db := global.DB.WithContext(ctx).Table(req.Table)
// 	err := db.Select(fields).Where("status=1 and  deleted_at IS NULL").Order(OrderStr).Scan(&list).Error
// 	if err == nil {
// 		list2 := m.getTreeList(list, 0)
// 		return list2
// 	}
// 	return nil
// }

func (m *CommonDbService) GetTreeData(ctx context.Context, req request.PidDataReq) (list []*request.PidTreeData) {

	OrderStr := req.PidField + " asc"
	fields := ""
	pidFields := req.Pid
	if utils.IsEmptyStr(req.Pid) {
		pidFields = "pid"
	}
	fields = req.PidField + " as `value`," + req.NameField + " as label," + pidFields + " as pid"
	db := global.DB.WithContext(ctx).Table(req.Table)
	whereSql := " status=1 AND deleted_at IS NULL "
	if req.Table == "basic_area" {
		fields = "id as `value`, areaname as label,parentid as pid,id"
		whereSql = ""
		if !utils.IsEmpty(req.Where) {
			whereSql = req.Where
		}
	} else {
		// if !utils.IsEmpty(req.PidValue) {
		// 	db = db.Where(pidFields+"=?", req.PidValue)
		// }
		if req.Table == "sys_menu" {
			if req.Where != "be_sys=1" {
				whereSql += " AND  be_sys=0"
			}
		} else if req.Table == "sys_role" {
			whereSql += " AND cu_id=" + gconv.String(req.CuId)
		} else {
			if !utils.IsEmpty(req.Where) {
				whereSql += " AND " + req.Where
			}
		}
	}

	// if req.PidValue != nil {
	// 	whereSql += fmt.Sprintf(" AND  `"+req.PidField+"` = %v", req.PidValue)
	// }

	err := db.Select(fields).Where(whereSql).Order(OrderStr).Scan(&list).Error
	if err == nil {
		list2 := m.getTreeList(list, 0)
		return list2
	}

	return nil
}

func (m *CommonDbService) getTreeList(list []*request.PidTreeData, pid int64) []*request.PidTreeData {
	res := make([]*request.PidTreeData, 0)
	for _, v := range list {
		if v.Pid == pid {
			v.Children = m.getTreeList(list, v.Value)
			res = append(res, v)
		}
	}
	return res
}

// 地区使用
// func (m *CommonDbService) getTreeList(list []*request.PidTreeData, pid int64) []*request.PidTreeData {
// 	res := make([]*request.PidTreeData, 0)
// 	for _, v := range list {
// 		if v.Pid == pid {
// 			v.Text = v.Label
// 			//val, _ := strconv.ParseInt(v.Value, 10, 64)
// 			//v.Pid = val
// 			v.Children = m.getTreeList(list, v.Id)
// 			res = append(res, v)
// 		}
// 	}
// 	return res
// }

// 从数据库获取子节点
// Author [Linjd]
func (commonDbService *CommonDbService) getTreeData(ctx context.Context, req request.PidDataReq) (list []request.PidTreeData) {
	if !utils.IsEmpty(req.Table) {

		if req.PidValue != nil {
			global.DB.WithContext(ctx).Raw("select "+req.PidField+" as value,"+req.NameField+" as label  FROM  `"+req.Table+"` WHERE `"+req.PidField+"` = ?", req.PidValue).Scan(&list)
		} else {
			global.DB.WithContext(ctx).Raw("select " + req.PidField + " as value," + req.NameField + " as label  FROM  `" + req.Table + "`").Scan(&list)
		}
		return list
	}
	return nil
	//err = global.DB.(&quickEdit).Error
}

// GetDict 获取字典
// Author [Linjd]
func (m *CommonDbService) GetDict(ctx context.Context, req request.GetDictReq) (list []request.PidData) {

	db := global.DB.WithContext(ctx).Model(&business.BasicDict{})
	field := "value,label,id,pid,module,type"

	if req.Module != "" {
		db.Where("module", req.Module)
		db.Where("pid > 0")
	} else {
		if req.ValueField != "" {
			field = req.ValueField + " as value,label,id,pid,module,type"
		}
		if req.Pid != nil {
			db.Where("pid", req.Pid)
		}
	}
	err := db.Select(field).Order("sort asc,value asc").Scan(&list).Error
	if err != nil {
		return nil
	}
	return list
	//err = global.DB.(&quickEdit).Error
}

// GetDict 获取字典 多了 type module
// Author [Linjd]
func (m *CommonDbService) GetDict2(ctx context.Context, req request.GetDictReq) (list []request.PidData2) {

	db := global.DB.WithContext(ctx).Model(&business.BasicDict{})
	field := "value,label,id,pid,module,type"
	if req.ValueField != "" {
		field = req.ValueField + " as value,label,id,pid,module,type"
	}
	if req.Pid != nil {
		db.Where("pid", req.Pid)
	}
	if req.Type != nil {
		db.Where("type", req.Type)
	}
	if req.Module != "" {
		db.Where("module", req.Module)
	}
	err := db.Select(field).Order("sort asc,value asc").Scan(&list).Error
	if err != nil {
		return nil
	}
	return list
	//err = global.DB.(&quickEdit).Error
}

// GetDict 获取字典 树状多级结构
// Author [Linjd]
func (m *CommonDbService) GetDictTree(ctx context.Context, req request.GetDictReq) (list []*request.PidTreeData) {

	db := global.DB.WithContext(ctx).Model(&business.BasicDict{})
	if req.Module == "" {
		return nil
	}
	field := "value,label,id,pid,module"
	if req.ValueField != "" {
		field = req.ValueField + " as value,label,id,pid,module"
	}
	if req.Pid != nil {
		db.Where("pid", req.Pid)
	}
	if req.Module != "" {
		db.Where("module", req.Module)
	}
	err := db.Select(field).Order("value asc ,id desc").Scan(&list).Error
	if err != nil {
		return nil
	}

	list2 := m.getTreeList(list, 0)
	return list2
}

// 获取个人memUser realname
func (m *CommonFileService) GetRealName(id int64) (name string, err error) {
	//basicFS := new(businessSev.BasicFileService)
	// 先从 redis 获取 ，没有则读取数据库 ，缓存2小时
	if id == 0 {
		return "", nil
	}
	cacheKey := m.GetCacheKey("GetRealName", id)
	cacheData, err := mycache.GetCache().Get(cacheKey)
	if err == nil {
		//global.LOG.Info("from cacheKey=" + cacheKey)
		name = cacheData.(string)
		//global.LOG.Info("from path=" + path)
		return name, err
	}
	user := business.MemUser{}
	err = global.DB.Where("id = ?", id).Select("id,username,realname").First(&user).Error
	if err != nil {
		//cacheData = []byte("")
		mycache.GetCache().Set(cacheKey, "", "GetRealName")
		return "", err
	}
	name = user.Realname
	//cacheData = []byte(path)
	mycache.GetCache().Set(cacheKey, name, "GetRealName")
	return name, err
}

// 获取系统用户 sysUser nickname
func (m *CommonFileService) GetNickNameSys(id int64) (name string, err error) {
	if id == 0 {
		return "", nil
	}
	cacheKey := m.GetCacheKey("GetNickNameSys", id)
	cacheData, err := mycache.GetCache().Get(cacheKey)
	if err == nil {
		global.LOG.Info("from cacheKey=" + cacheKey)
		name = cacheData.(string)
		global.LOG.Info("from name=" + name)
		return name, err
	}
	user := business.MemUser{}
	err = global.DB.Where("id = ?", id).Select("id,username,nick_name").First(&user).Error
	if err != nil {
		//cacheData = []byte("")
		mycache.GetCache().Set(cacheKey, "", "GetNickNameSys")
		return "", err
	}
	name = user.Nickname
	//cacheData = []byte(path)
	mycache.GetCache().Set(cacheKey, name, "GetNickNameSys")
	return name, err
}
