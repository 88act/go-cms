package common

import (
	"github.com/88act/go-cms/server/global"
	"github.com/88act/go-cms/server/model/common/request"
)

type CommonDbService struct {
}

// QuickEdit 更新QuickEdit记录
// Author [Linjd]
func (commonDbService *CommonDbService) QuickEdit(quickEdit request.QuickEdit) (err error) {
	err = global.DB.Exec("UPDATE `"+quickEdit.Table+"` SET `"+quickEdit.Field+"`=? WHERE `id` = ?", quickEdit.Value, quickEdit.Id).Error
	//err = global.DB.(&quickEdit).Error
	return err
}
