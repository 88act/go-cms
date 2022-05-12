// 自动生成模板ActShop
package business

import (
	"go-cms/global"
	"time"
)

// ActShop 结构体
type ActShop struct {
	ActShopMini
	Desc         string     `json:"desc" cn:"简介" form:"desc" gorm:"column:desc;comment:简介;type:varchar(500);"`
	Detail       string     `json:"detail" cn:"详细内容" form:"detail" gorm:"column:detail;comment:详细内容;type:mediumtext;"`
	Avater       string     `json:"avater" cn:"缩略图" form:"avater" gorm:"column:avater;comment:缩略图;type:varchar(100);"`
	MediaList    string     `json:"mediaList" cn:"媒体列表" form:"mediaList" gorm:"column:media_list;comment:媒体列表;type:varchar(500);"`
	Lng          *int       `json:"lng" cn:"经度" form:"lng" gorm:"column:lng;comment:经度:无6位小数;type:bigint"`
	Lat          *int       `json:"lat" cn:"纬度" form:"lat" gorm:"column:lat;comment:纬度:无6位小数;type:bigint"`
	VipLev       *int       `json:"vipLev" cn:"vip等级" form:"vipLev" gorm:"column:vip_lev;comment:vip等级:0,1,2,3;type:int"`
	VipTime      *time.Time `json:"vipTime" cn:"vip结束时间" form:"vipTime" gorm:"column:vip_time;comment:vip结束时间;type:datetime"`
	TotalWhole   *int       `json:"totalWhole" cn:"综合指数" form:"totalWhole" gorm:"column:total_whole;comment:综合指数;type:int"`
	TotalShare   *int       `json:"totalShare" cn:"总分享" form:"totalShare" gorm:"column:total_share;comment:总分享;type:int"`
	TotalFav     *int       `json:"totalFav" cn:"总收藏" form:"totalFav" gorm:"column:total_fav;comment:总收藏;type:int"`
	TotalJoin    *int       `json:"totalJoin" cn:"总报名人数" form:"totalJoin" gorm:"column:total_join;comment:总报名人数;type:int"`
	TotalDiscuss *int       `json:"totalDiscuss" cn:"总评论" form:"totalDiscuss" gorm:"column:total_discuss;comment:总评论;type:int"`
	TotalClick   *int       `json:"totalClick" cn:"总点击" form:"totalClick" gorm:"column:total_click;comment:总点击;type:int"`
	TotalStar    *int       `json:"totalStar" cn:"总评" form:"totalStar" gorm:"column:total_star;comment:总评;type:int"`
}

// ActShopMini 结构体
type ActShopMini struct {
	global.GO_MODEL
	UserId  *int   `json:"userId" cn:"用户id" form:"userId" gorm:"column:user_id;comment:用户id;type:bigint"`
	Name    string `json:"name" cn:"标题" form:"name" gorm:"column:name;comment:标题;type:varchar(150);"`
	Address string `json:"address" cn:"地址" form:"address" gorm:"column:address;comment:地址;type:varchar(500);"`
	AreaId  *int   `json:"areaId" cn:"地区id" form:"areaId" gorm:"column:area_id;comment:地区id;type:int"`
	Email   string `json:"email" cn:"邮件" form:"email" gorm:"column:email;comment:邮件;type:varchar(60);"`
	Mobile  string `json:"mobile" cn:"手机" form:"mobile" gorm:"column:mobile;comment:手机;type:varchar(20);"`
	Status  *int   `json:"status" cn:"状态" form:"status" gorm:"column:status;comment:状态:0未审核 1审核通过 2未通过审核 3黑名单;type:smallint"`
}

// TableName ActShop 表名
func (ActShop) TableName() string {
	return "act_shop"
}
