package model

import (
	"database/sql"

	"gorm.io/gorm"
)

//统一model 执行接口
type Executable interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
}

var ErrNotFound = gorm.ErrRecordNotFound // sqlx.ErrNotFound

var UserAuthTypeSystem string = "system"  //平台内部
var UserAuthTypeSmallWX string = "wxMini" //微信小程序
