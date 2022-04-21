package business

import (
	"go-cms/global"
	"go-cms/model/business"
	bizReq "go-cms/model/business/request"
	"go-cms/model/common/request"
	"go-cms/utils"
	"sync"
)

type ImTxMsgService struct {
}

var once_ImTxMsg sync.Once = sync.Once{}
var obj_ImTxMsgService *ImTxMsgService

//获取单例
func GetImTxMsgSev() *ImTxMsgService {
	once_ImTxMsg.Do(func() {
		obj_ImTxMsgService = new(ImTxMsgService)
		//instanse.init()
	})
	return obj_ImTxMsgService
}

// Create 创建ImTxMsg记录
// Author [88act](https://github.com/88act)
func (m *ImTxMsgService) Create(data business.ImTxMsg) (id uint, err error) {
	err = global.DB.Create(&data).Error
	if err != nil {
		return 0, err
	}
	return data.ID, err
}

// Delete 删除ImTxMsg记录
// Author [88act](https://github.com/88act)
func (m *ImTxMsgService) Delete(data business.ImTxMsg) (err error) {
	err = global.DB.Delete(&data).Error
	return err
}

// DeleteByIds 批量删除ImTxMsg记录
// Author [88act](https://github.com/88act)
func (m *ImTxMsgService) DeleteByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]business.ImTxMsg{}, "id in ?", ids.Ids).Error
	return err
}

// Update  更新ImTxMsg记录
// Author [88act](https://github.com/88act)
func (m *ImTxMsgService) Update(data business.ImTxMsg) (err error) {
	err = global.DB.Save(&data).Error
	return err
}

// UpdateByMap  更新ImTxMsg记录 by Map
// values := map[string]interface{}{
// 	"status":0,
// 	"from": hash,
// }
// Author [88act](https://github.com/88act)
func (m *ImTxMsgService) UpdateByMap(data business.ImTxMsg, mapData map[string]interface{}) (err error) {
	err = global.DB.Model(&data).Updates(mapData).Error
	return err
}

// Get 根据id获取ImTxMsg记录
// Author [88act](https://github.com/88act)
func (m *ImTxMsgService) Get(id uint, fields string) (obj business.ImTxMsg, err error) {

	if utils.IsEmpty(fields) {
		err = global.DB.Where("id = ?", id).First(&obj).Error
	} else {
		err = global.DB.Select(fields).Where("id = ?", id).First(&obj).Error
	}

	//如果有图片image类型，更新图片path
	obj.MapData = make(map[string]string)
	return obj, err
}

// GetList 分页获取ImTxMsg记录
// Author [88act](https://github.com/88act)
func (m *ImTxMsgService) GetList(info bizReq.ImTxMsgSearch, createdAtBetween []string, fields string) (list []business.ImTxMsgMini, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.ImTxMsg{})
	//var imTxMsgs []business.ImTxMsg
	var imTxMsgs []business.ImTxMsgMini

	//修改 by ljd
	if info.ID > 0 {
		db = db.Where("`id` = ?", info.ID)
	}
	if len(createdAtBetween) >= 2 {
		db = db.Where("`created_at` BETWEEN ? AND ?", createdAtBetween[0], createdAtBetween[1])
	}

	// 如果有条件搜索 下方会自动创建搜索语句
	if info.OrderId != "" {
		db = db.Where("`order_id` = ?", info.OrderId)
	}
	if info.ChatType != "" {
		db = db.Where("`chat_type` = ?", info.ChatType)
	}
	if info.MsgTime != "" {
		db = db.Where("`msg_time` = ?", info.MsgTime)
	}
	if info.FromAccount != "" {
		db = db.Where("`from_account` = ?", info.FromAccount)
	}
	if info.ToAccount != "" {
		db = db.Where("`to_account` = ?", info.ToAccount)
	}
	if info.MsgTimestamp != nil {
		db = db.Where("`msg_timestamp` > > ?", info.MsgTimestamp)
	}
	if info.MsgType != "" {
		db = db.Where("`msg_type` = ?", info.MsgType)
	}
	if info.MsgText != "" {
		db = db.Where("`msg_text` LIKE ?", "%"+info.MsgText+"%")
	}
	if info.StatusMedia != nil {
		db = db.Where("`status_media` = ?", info.StatusMedia)
	}
	if info.ClientIp != "" {
		db = db.Where("`client_ip` LIKE ?", "%"+info.ClientIp+"%")
	}
	if info.MsgFromPlatform != "" {
		db = db.Where("`msg_from_platform` = ?", info.MsgFromPlatform)
	}
	if info.Status != nil {
		db = db.Where("`status` = ?", info.Status)
	}
	if info.OrderStatus != nil {
		db = db.Where("`order_status` = ?", info.OrderStatus)
	}

	err = db.Count(&total).Error
	if err != nil {
		return imTxMsgs, 0, err
	}
	//err = db.Limit(limit).Offset(offset).Find(&imTxMsgs).Error
	//修改 by ljd  增加查询排序
	OrderStr := "id desc"
	if !utils.IsEmpty(order) {
		if desc {
			OrderStr = order + " desc"
		} else {
			OrderStr = order
		}
	}
	if utils.IsEmpty(fields) {
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&imTxMsgs).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&imTxMsgs).Error
	}
	//如果有图片image类型，更新图片path
	for i, v := range imTxMsgs {
		v.MapData = make(map[string]string)
		imTxMsgs[i] = v
	}
	return imTxMsgs, total, err
}

//GetListAll 分页获取ImTxMsg记录 (全部字段)
// Author [88act](https://github.com/88act)
func (m *ImTxMsgService) GetListAll(info bizReq.ImTxMsgSearch, createdAtBetween []string, fields string) (list []business.ImTxMsg, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//修改 by ljd  增加查询排序
	order := info.OrderKey
	desc := info.OrderDesc
	// 创建db
	db := global.DB.Model(&business.ImTxMsg{})
	var imTxMsgs []business.ImTxMsg
	//var imTxMsgs []business.ImTxMsgMini

	if info.ID > 0 {
		db = db.Where("`id` = ?", info.ID)
	}
	if len(createdAtBetween) >= 2 {
		db = db.Where("`created_at` BETWEEN ? AND ?", createdAtBetween[0], createdAtBetween[1])
	}

	// 如果有条件搜索 下方会自动创建搜索语句
	if info.OrderId != "" {
		db = db.Where("`order_id` = ?", info.OrderId)
	}
	if info.ChatType != "" {
		db = db.Where("`chat_type` = ?", info.ChatType)
	}
	if info.MsgTime != "" {
		db = db.Where("`msg_time` = ?", info.MsgTime)
	}
	if info.FromAccount != "" {
		db = db.Where("`from_account` = ?", info.FromAccount)
	}
	if info.ToAccount != "" {
		db = db.Where("`to_account` = ?", info.ToAccount)
	}
	if info.MsgTimestamp != nil {
		db = db.Where("`msg_timestamp` > > ?", info.MsgTimestamp)
	}
	if info.MsgType != "" {
		db = db.Where("`msg_type` = ?", info.MsgType)
	}
	if info.MsgText != "" {
		db = db.Where("`msg_text` LIKE ?", "%"+info.MsgText+"%")
	}
	if info.StatusMedia != nil {
		db = db.Where("`status_media` = ?", info.StatusMedia)
	}
	if info.ClientIp != "" {
		db = db.Where("`client_ip` LIKE ?", "%"+info.ClientIp+"%")
	}
	if info.MsgFromPlatform != "" {
		db = db.Where("`msg_from_platform` = ?", info.MsgFromPlatform)
	}
	if info.Status != nil {
		db = db.Where("`status` = ?", info.Status)
	}
	if info.OrderStatus != nil {
		db = db.Where("`order_status` = ?", info.OrderStatus)
	}

	err = db.Count(&total).Error
	if err != nil {
		return imTxMsgs, 0, err
	}

	//err = db.Limit(limit).Offset(offset).Find(&imTxMsgs).Error
	//修改 by ljd  增加查询排序
	OrderStr := "id desc"
	if !utils.IsEmpty(order) {
		if desc {
			OrderStr = order + " desc"
		} else {
			OrderStr = order
		}
	}
	if utils.IsEmpty(fields) {
		err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&imTxMsgs).Error
	} else {
		err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&imTxMsgs).Error
	}
	//如果有图片image类型，更新图片path
	for i, v := range imTxMsgs {
		v.MapData = make(map[string]string)
		imTxMsgs[i] = v
	}
	return imTxMsgs, total, err
}
