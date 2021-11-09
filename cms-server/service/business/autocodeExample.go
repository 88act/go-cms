package business

import (
	"github.com/88act/go-cms/server/global"
	"github.com/88act/go-cms/server/model/business"
	"github.com/88act/go-cms/server/model/business/request"
)

type BusinessExampleService struct {
}

//@author: [88act](https://github.com/88act)
//@function: CreateBusinessExample
//@description: 创建自动化示例数据
//@param: businessExampleService *BusinessExampleService
//@return: err error

func (businessExampleService *BusinessExampleService) CreateBusinessExample(businessExample business.BusinessExample) (err error) {

	err = global.DB.Create(&businessExample).Error
	return err
}

//@author: [88act](https://github.com/88act)
//@function: DeleteBusinessExample
//@description: 删除自动化示例数据
//@param: businessExample business.BusinessExample
//@return: err error

func (businessExampleService *BusinessExampleService) DeleteBusinessExample(businessExample business.BusinessExample) (err error) {
	err = global.DB.Delete(&businessExample).Error
	return err
}

//@author: [88act](https://github.com/88act)
//@function: UpdateBusinessExample
//@description: 更新自动化示例数据
//@param: businessExample *business.BusinessExample
//@return: err error

func (businessExampleService *BusinessExampleService) UpdateBusinessExample(businessExample *business.BusinessExample) (err error) {
	err = global.DB.Save(businessExample).Error
	return err
}

//@author: [88act](https://github.com/88act)
//@function: GetBusinessExampleDetail
//@description: 根据id获取自动化示例数据
//@param: id uint
//@return: err error, businessExample business.BusinessExample

func (businessExampleService *BusinessExampleService) GetBusinessExample(id uint) (err error, businessExample business.BusinessExample) {
	err = global.DB.Where("id = ?", id).First(&businessExample).Error
	return
}

//@author: [88act](https://github.com/88act)
//@function: GetBusinessExampleInfoList
//@description: 分页获取自动化示例数据
//@param: info request.BusinessExampleSearch
//@return: err error, list interface{}, total int64

func (businessExampleService *BusinessExampleService) GetBusinessExampleInfoList(info request.BusinessExampleSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.DB.Model(&business.BusinessExample{})
	var businessExamples []business.BusinessExample
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.BusinessExampleField != "" {
		db = db.Where("label LIKE ?", "%"+info.BusinessExampleField+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&businessExamples).Error
	return err, businessExamples, total
}
