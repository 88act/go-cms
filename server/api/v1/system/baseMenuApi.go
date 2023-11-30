package system

import (
	"go-cms/model/business"
	"go-cms/model/common/request"
	bizSev "go-cms/service/business"
	commSev "go-cms/service/common"
	"go-cms/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gconv"
	"github.com/jinzhu/copier"
)

type BaseMenuApi struct {
	commSev.BaseApi
}
type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

// 平台用户
var UserType_Plat = 9

// @Tags AuthorityMenu
// @Summary 获取用户动态路由
func (m *BaseMenuApi) GetMenu(c *gin.Context) {
	var idReq request.IdReq
	_ = c.ShouldBindJSON(&idReq)
	userId := utils.GetUserID(c)

	memUser, err := bizSev.GetMemUserSev().Get(c, userId, "")
	if err != nil || memUser.Status != 1 {
		m.FailWithMessage("用户不存在或状态异常", c)
		return
	}

	roleId := idReq.Id
	if roleId <= 0 {
		roleId = memUser.Role
	}
	roleList := strings.Split(memUser.RoleList, ",")

	if !utils.ArrayHave(roleList, gconv.String(memUser.Role)) {
		m.FailWithMessage("用户没分配该权限", c)
		return
	}
	if !utils.ArrayHave(roleList, gconv.String(memUser.Role)) {
		m.FailWithMessage("用户没分配该权限", c)
		return
	}
	// 菜单
	role, err := bizSev.GetSysRoleSev().Get(c, roleId, "")
	if err != nil {
		m.FailWithMessage("权限不存在", c)
		return
	}

	menus, err := bizSev.GetSysMenuSev().GetMenu(c, role.MenuList)
	var mList []*business.SysMenuMini
	for _, v := range menus {
		var mini business.SysMenuMini
		_ = copier.Copy(&mini, v)
		mini.Meta = business.MenuMeta{}
		mini.Meta.Icon = v.Icon
		mini.Meta.Title = v.Title
		mini.Meta.Rank = v.Sort
		mini.Meta.ShowLink = v.ShowLink // 是否隐藏页面
		mini.Path = "/" + mini.Path
		mList = append(mList, &mini)
	}
	// 转换成树状结构
	menuTree := m.getTreeList(mList, 0)
	menuResp := business.SysMenuMiniResp{List: menuTree}
	// 默认菜单
	if role.DefaultMenu > 0 {
		menuDef, err := bizSev.GetSysMenuSev().Get(c, role.DefaultMenu, "")
		if err == nil {
			menuResp.DefaultMenu = menuDef.Path
		}
	}

	if err != nil {
		m.FailWithMessage("获取失败", c)
	} else {
		m.OkWithDetailed(menuResp, "获取成功", c)
	}
}

func (m *BaseMenuApi) getTreeList(list []*business.SysMenuMini, pid int64) []*business.SysMenuMini {
	res := make([]*business.SysMenuMini, 0)
	for _, v := range list {
		if v.Pid == pid {
			v.Children = m.getTreeList(list, v.Id)
			res = append(res, v)
		}
	}
	return res
}
