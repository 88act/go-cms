package example

import (
	"github.com/88act/go-cms/server/global"
	"github.com/88act/go-cms/server/model/common/request"
	"github.com/88act/go-cms/server/model/example"
	"github.com/88act/go-cms/server/model/system"
	systemService "github.com/88act/go-cms/server/service/system"
)

type CustomerService struct {
}

//@author: [88act](https://github.com/88act)
//@function: CreateExaCustomer
//@description: 创建客户
//@param: e model.ExaCustomer
//@return: err error

func (exa *CustomerService) CreateExaCustomer(e example.ExaCustomer) (err error) {
	err = global.DB.Create(&e).Error
	return err
}

//@author: [88act](https://github.com/88act)
//@function: DeleteFileChunk
//@description: 删除客户
//@param: e model.ExaCustomer
//@return: err error

func (exa *CustomerService) DeleteExaCustomer(e example.ExaCustomer) (err error) {
	err = global.DB.Delete(&e).Error
	return err
}

//@author: [88act](https://github.com/88act)
//@function: UpdateExaCustomer
//@description: 更新客户
//@param: e *model.ExaCustomer
//@return: err error

func (exa *CustomerService) UpdateExaCustomer(e *example.ExaCustomer) (err error) {
	err = global.DB.Save(e).Error
	return err
}

//@author: [88act](https://github.com/88act)
//@function: GetExaCustomer
//@description: 获取客户信息
//@param: id uint
//@return: err error, customer model.ExaCustomer

func (exa *CustomerService) GetExaCustomer(id uint) (err error, customer example.ExaCustomer) {
	err = global.DB.Where("id = ?", id).First(&customer).Error
	return
}

//@author: [88act](https://github.com/88act)
//@function: GetCustomerInfoList
//@description: 分页获取客户列表
//@param: sysUserAuthorityID string, info request.PageInfo
//@return: err error, list interface{}, total int64

func (exa *CustomerService) GetCustomerInfoList(sysUserAuthorityID string, info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.DB.Model(&example.ExaCustomer{})
	var a system.SysAuthority
	a.AuthorityId = sysUserAuthorityID
	err, auth := systemService.AuthorityServiceApp.GetAuthorityInfo(a)
	if err != nil {
		return
	}
	var dataId []string
	for _, v := range auth.DataAuthorityId {
		dataId = append(dataId, v.AuthorityId)
	}
	var CustomerList []example.ExaCustomer
	err = db.Where("sys_user_authority_id in ?", dataId).Count(&total).Error
	if err != nil {
		return err, CustomerList, total
	} else {
		err = db.Limit(limit).Offset(offset).Preload("SysUser").Where("sys_user_authority_id in ?", dataId).Find(&CustomerList).Error
	}
	return err, CustomerList, total
}
