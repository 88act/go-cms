package system

import (
	"github.com/88act/go-cms/server/config"
)

// 配置文件结构体
type System struct {
	Config config.Server `json:"config"`
}
