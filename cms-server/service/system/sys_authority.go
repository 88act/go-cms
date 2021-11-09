package system

import (
	"errors"
	"strconv"

	"github.com/88act/go-cms/server/global"
	"github.com/88act/go-cms/server/model/common/request"
	"github.com/88act/go-cms/server/model/system"
	"github.com/88act/go-cms/server/model/system/response"
	"gorm.io/gorm"
)

//@author: [88act](https://github.com/88act)
//@function: CreateAuthority
//@description: 创建一个角色
//@param: auth model.SysAuthority
//@return: err error, authority model.SysAuthority

type AuthorityService struct {
}

var AuthorityServiceApp = new(AuthorityService)

func (authorityService *AuthorityService) CreateAuthority(auth system.SysAuthority) (err error, authority system.SysAuthority) {
	var authorityBox system.SysAuthority
	if !errors.Is(global.DB.Where("authority_id = ?", auth.AuthorityId).First(&authorityBox).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同角色id"), auth
	}
	err = global.DB.Create(&auth).Error
	return err, auth
}

//@author: [88act](https://github.com/88act)
//@function: CopyAuthority
//@description: 复制一个角色
//@param: copyInfo response.SysAuthorityCopyResponse
//@return: err error, authority model.SysAuthority

func (authorityService *AuthorityService) CopyAuthority(copyInfo response.SysAuthorityCopyResponse) (err error, authority system.SysAuthority) {
	var authorityBox system.SysAuthority
	if !errors.Is(global.DB.Where("authority_id = ?", copyInfo.Authority.AuthorityId).First(&authorityBox).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同角色id"), authority
	}
	copyInfo.Authority.Children = []system.SysAuthority{}
	err, menus := MenuServiceApp.GetMenuAuthority(&request.GetAuthorityId{AuthorityId: copyInfo.OldAuthorityId})
	if err != nil {
		return
	}
	var baseMenu []system.SysBaseMenu
	for _, v := range menus {
		intNum, _ := strconv.Atoi(v.MenuId)
		v.SysBaseMenu.ID = uint(intNum)
		baseMenu = append(baseMenu, v.SysBaseMenu)
	}
	copyInfo.Authority.SysBaseMenus = baseMenu
	err = global.DB.Create(&copyInfo.Authority).Error
	if err != nil {
		return
	}
	paths := CasbinServiceApp.GetPolicyPathByAuthorityId(copyInfo.OldAuthorityId)
	err = CasbinServiceApp.UpdateCasbin(copyInfo.Authority.AuthorityId, paths)
	if err != nil {
		_ = authorityService.DeleteAuthority(&copyInfo.Authority)
	}
	return err, copyInfo.Authority
}

//@author: [88act](https://github.com/88act)
//@function: UpdateAuthority
//@description: 更改一个角色
//@param: auth model.SysAuthority
//@return: err error, authority model.SysAuthority

func (authorityService *AuthorityService) UpdateAuthority(auth system.SysAuthority) (err error, authority system.SysAuthority) {
	err = global.DB.Where("authority_id = ?", auth.AuthorityId).First(&system.SysAuthority{}).Updates(&auth).Error
	return err, auth
}

//@author: [88act](https://github.com/88act)
//@function: DeleteAuthority
//@description: 删除角色
//@param: auth *model.SysAuthority
//@return: err error

func (authorityService *AuthorityService) DeleteAuthority(auth *system.SysAuthority) (err error) {
	if !errors.Is(global.DB.Where("authority_id = ?", auth.AuthorityId).First(&system.SysUser{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("此角色有用户正在使用禁止删除")
	}
	if !errors.Is(global.DB.Where("parent_id = ?", auth.AuthorityId).First(&system.SysAuthority{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("此角色存在子角色不允许删除")
	}
	db := global.DB.Preload("SysBaseMenus").Where("authority_id = ?", auth.AuthorityId).First(auth)
	err = db.Unscoped().Delete(auth).Error
	if err != nil {
		return
	}
	if len(auth.SysBaseMenus) > 0 {
		err = global.DB.Model(auth).Association("SysBaseMenus").Delete(auth.SysBaseMenus)
		if err != nil {
			return
		}
		//err = db.Association("SysBaseMenus").Delete(&auth)
	} else {
		err = db.Error
		if err != nil {
			return
		}
	}
	err = global.DB.Delete(&[]system.SysUseAuthority{}, "sys_authority_authority_id = ?", auth.AuthorityId).Error
	CasbinServiceApp.ClearCasbin(0, auth.AuthorityId)
	return err
}

//@author: [88act](https://github.com/88act)
//@function: GetAuthorityInfoList
//@description: 分页获取数据
//@param: info request.PageInfo
//@return: err error, list interface{}, total int64

func (authorityService *AuthorityService) GetAuthorityInfoList(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.DB.Model(&system.SysAuthority{})
	err = db.Where("parent_id = 0").Count(&total).Error
	var authority []system.SysAuthority
	err = db.Limit(limit).Offset(offset).Preload("DataAuthorityId").Where("parent_id = 0").Find(&authority).Error
	if len(authority) > 0 {
		for k := range authority {
			err = authorityService.findChildrenAuthority(&authority[k])
		}
	}
	return err, authority, total
}

//@author: [88act](https://github.com/88act)
//@function: GetAuthorityInfo
//@description: 获取所有角色信息
//@param: auth model.SysAuthority
//@return: err error, sa model.SysAuthority

func (authorityService *AuthorityService) GetAuthorityInfo(auth system.SysAuthority) (err error, sa system.SysAuthority) {
	err = global.DB.Preload("DataAuthorityId").Where("authority_id = ?", auth.AuthorityId).First(&sa).Error
	return err, sa
}

//@author: [88act](https://github.com/88act)
//@function: SetDataAuthority
//@description: 设置角色资源权限
//@param: auth model.SysAuthority
//@return: error

func (authorityService *AuthorityService) SetDataAuthority(auth system.SysAuthority) error {
	var s system.SysAuthority
	global.DB.Preload("DataAuthorityId").First(&s, "authority_id = ?", auth.AuthorityId)
	err := global.DB.Model(&s).Association("DataAuthorityId").Replace(&auth.DataAuthorityId)
	return err
}

//@author: [88act](https://github.com/88act)
//@function: SetMenuAuthority
//@description: 菜单与角色绑定
//@param: auth *model.SysAuthority
//@return: error

func (authorityService *AuthorityService) SetMenuAuthority(auth *system.SysAuthority) error {
	var s system.SysAuthority
	global.DB.Preload("SysBaseMenus").First(&s, "authority_id = ?", auth.AuthorityId)
	err := global.DB.Model(&s).Association("SysBaseMenus").Replace(&auth.SysBaseMenus)
	return err
}

//@author: [88act](https://github.com/88act)
//@function: findChildrenAuthority
//@description: 查询子角色
//@param: authority *model.SysAuthority
//@return: err error

func (authorityService *AuthorityService) findChildrenAuthority(authority *system.SysAuthority) (err error) {
	err = global.DB.Preload("DataAuthorityId").Where("parent_id = ?", authority.AuthorityId).Find(&authority.Children).Error
	if len(authority.Children) > 0 {
		for k := range authority.Children {
			err = authorityService.findChildrenAuthority(&authority.Children[k])
		}
	}
	return err
}
