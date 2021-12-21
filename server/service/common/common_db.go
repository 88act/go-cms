package common

import (
	"go-cms/global"
	"go-cms/model/common/request"
	"sync"
)

type CommonDbService struct {
}

var once_CommonDb sync.Once = sync.Once{}
var obj_CommonDbService *CommonDbService

//获取单例
func GetCommonDbService() *CommonDbService {
	once_CommonDb.Do(func() {
		obj_CommonDbService = new(CommonDbService)
		//instanse.init()
	})
	return obj_CommonDbService
}

// QuickEdit 更新QuickEdit记录
// Author [Linjd]
func (commonDbService *CommonDbService) QuickEdit(quickEdit request.QuickEdit) (err error) {
	err = global.DB.Exec("UPDATE `"+quickEdit.Table+"` SET `"+quickEdit.Field+"`=? WHERE `id` = ?", quickEdit.Value, quickEdit.Id).Error
	//err = global.DB.(&quickEdit).Error
	return err
}
