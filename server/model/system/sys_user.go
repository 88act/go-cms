package system

import (
	"go-cms/global"

	uuid "github.com/satori/go.uuid"
)

type SysUser struct {
	global.GO_MODEL
	UUID        uuid.UUID      `json:"uuid" gorm:"comment:用户UUID"`                                                       // 用户UUID
	Username    string         `json:"userName" gorm:"comment:用户登录名"`                                                    // 用户登录名
	Password    string         `json:"-"  gorm:"comment:用户登录密码"`                                                         // 用户登录密码
	NickName    string         `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`                                        // 用户昵称
	SideMode    string         `json:"sideMode" gorm:"default:dark;comment:用户侧边主题"`                                      // 用户侧边主题
	HeaderImg   string         `json:"headerImg" gorm:"default:https://cms.88act.com/res/sys/0/avatar.jpg;comment:用户头像"` // 用户头像
	BaseColor   string         `json:"baseColor" gorm:"default:#fff;comment:基础颜色"`                                       // 基础颜色
	ActiveColor string         `json:"activeColor" gorm:"default:#1890ff;comment:活跃颜色"`                                  // 活跃颜色
	AuthorityId string         `json:"authorityId" gorm:"default:888;comment:用户角色ID"`                                    // 用户角色ID
	Authority   SysAuthority   `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
	Authorities []SysAuthority `json:"authorities" gorm:"many2many:sys_user_authority;"`
}
