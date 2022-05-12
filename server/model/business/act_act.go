// 自动生成模板ActAct
package business

import (
	"go-cms/global"
	"time"
)

// ActAct 结构体
type ActAct struct {
	ActActMini
}

// ActActMini 结构体
type ActActMini struct {
	global.GO_MODEL
	UserId       *int       `json:"userId" cn:"用户id" form:"userId" gorm:"column:user_id;comment:用户id;type:bigint"`
	CatId        *int       `json:"catId" cn:"系统栏目" form:"catId" gorm:"column:cat_id;comment:系统栏目;type:int"`
	CatIdUser    *int       `json:"catIdUser" cn:"用户栏目" form:"catIdUser" gorm:"column:cat_id_user;comment:用户栏目;type:int"`
	BeOnline     *bool      `json:"beOnline" cn:"类型" form:"beOnline" gorm:"column:be_online;comment:类型:0=线下 1=线上;type:tinyint"`
	ActType      string     `json:"actType" cn:"活动类型" form:"actType" gorm:"column:act_type;comment:活动类型:1室内活动,2户外活动,3视频直播,4视频点播,5群组聊天,6文本文档;type:varchar(500);"`
	BeMulti      *bool      `json:"beMulti" cn:"分期活动" form:"beMulti" gorm:"column:be_multi;comment:分期活动:0=否 1=是;type:tinyint"`
	Phase        *int       `json:"phase" cn:"分期" form:"phase" gorm:"column:phase;comment:分期;type:int"`
	Title        string     `json:"title" cn:"标题" form:"title" gorm:"column:title;comment:标题;type:varchar(150);"`
	Desc         string     `json:"desc" cn:"简介" form:"desc" gorm:"column:desc;comment:简介;type:varchar(500);"`
	Avater       string     `json:"avater" cn:"缩略图" form:"avater" gorm:"column:avater;comment:缩略图;type:varchar(100);"`
	MediaList    string     `json:"mediaList" cn:"媒体列表" form:"mediaList" gorm:"column:media_list;comment:媒体列表;type:varchar(500);"`
	Address      string     `json:"address" cn:"地址" form:"address" gorm:"column:address;comment:地址;type:varchar(500);"`
	AreaId       *int       `json:"areaId" cn:"地区id" form:"areaId" gorm:"column:area_id;comment:地区id;type:int"`
	Lng          *int       `json:"lng" cn:"经度" form:"lng" gorm:"column:lng;comment:经度:无6位小数;type:bigint"`
	Lat          *int       `json:"lat" cn:"纬度" form:"lat" gorm:"column:lat;comment:纬度:无6位小数;type:bigint"`
	BeginTime    *time.Time `json:"beginTime" cn:"开始时间" form:"beginTime" gorm:"column:begin_time;comment:开始时间;type:datetime"`
	EndTime      *time.Time `json:"endTime" cn:"结束时间" form:"endTime" gorm:"column:end_time;comment:结束时间;type:datetime"`
	LiveTime     *time.Time `json:"liveTime" cn:"直播开始" form:"liveTime" gorm:"column:live_time;comment:直播开始;type:datetime"`
	LiveEnd      *time.Time `json:"liveEnd" cn:"直播结束" form:"liveEnd" gorm:"column:live_end;comment:直播结束;type:datetime"`
	Presenter    string     `json:"presenter" cn:"主持/讲师" form:"presenter" gorm:"column:presenter;comment:主持/讲师;type:varchar(200);"`
	PhoneKefu    string     `json:"phoneKefu" cn:"客服电话" form:"phoneKefu" gorm:"column:phone_kefu;comment:客服电话;type:varchar(100);"`
	PhoneHezu    string     `json:"phoneHezu" cn:"合作电话" form:"phoneHezu" gorm:"column:phone_hezu;comment:合作电话;type:varchar(100);"`
	Wx           string     `json:"wx" cn:"微信" form:"wx" gorm:"column:wx;comment:微信;type:varchar(100);"`
	Qq           string     `json:"qq" cn:"QQ" form:"qq" gorm:"column:qq;comment:QQ;type:varchar(100);"`
	GroupId      string     `json:"groupId" cn:"群组id" form:"groupId" gorm:"column:group_id;comment:群组id;type:varchar(100);"`
	EndJoinTime  *time.Time `json:"endJoinTime" cn:"结束报名" form:"endJoinTime" gorm:"column:end_join_time;comment:结束报名;type:datetime"`
	Price        *int       `json:"price" cn:"票价" form:"price" gorm:"column:price;comment:票价:单位分;type:int"`
	PriceVip     *int       `json:"priceVip" cn:"vip票价" form:"priceVip" gorm:"column:price_vip;comment:vip票价:单位分;type:int"`
	PriceDesc    string     `json:"priceDesc" cn:"票价描述" form:"priceDesc" gorm:"column:price_desc;comment:票价描述;type:varchar(200);"`
	RefundType   *int       `json:"refundType" cn:"退费类型" form:"refundType" gorm:"column:refund_type;comment:退费类型:1=不可退款 2=活动开始前都可退款 3=活动开始前n天可退费;type:smallint"`
	RefundDays   *int       `json:"refundDays" cn:"退费天" form:"refundDays" gorm:"column:refund_days;comment:退费天:开始前n天退费;type:smallint"`
	BeShowjoin   *int       `json:"beShowjoin" cn:"是否显示人数" form:"beShowjoin" gorm:"column:be_showjoin;comment:是否显示人数:0=显示 1=不显示;type:smallint"`
	StatusAct    *int       `json:"statusAct" cn:"活动状态" form:"statusAct" gorm:"column:status_act;comment:活动状态:0未开始,1报名中,2停止报名,3进行中,4已结束,5评价期，6结束;type:smallint"`
	MaxJoin      *int       `json:"maxJoin" cn:"最大报名人数" form:"maxJoin" gorm:"column:max_join;comment:最大报名人数;type:int"`
	NowJoin      *int       `json:"nowJoin" cn:"当前报名人数" form:"nowJoin" gorm:"column:now_join;comment:当前报名人数;type:int"`
	BeChosen     *bool      `json:"beChosen" cn:"是否精选" form:"beChosen" gorm:"column:be_chosen;comment:是否精选:0=否 1=是;type:tinyint"`
	BeWeek       *bool      `json:"beWeek" cn:"是否周末" form:"beWeek" gorm:"column:be_week;comment:是否周末:0=否 1=是;type:tinyint"`
	BeVip        *bool      `json:"beVip" cn:"是否会员" form:"beVip" gorm:"column:be_vip;comment:是否会员:0=否 1=是;type:tinyint"`
	TotalWhole   *int       `json:"totalWhole" cn:"综合指数" form:"totalWhole" gorm:"column:total_whole;comment:综合指数;type:int"`
	TotalShare   *int       `json:"totalShare" cn:"总分享" form:"totalShare" gorm:"column:total_share;comment:总分享;type:int"`
	TotalFav     *int       `json:"totalFav" cn:"总收藏" form:"totalFav" gorm:"column:total_fav;comment:总收藏;type:int"`
	TotalJoin    *int       `json:"totalJoin" cn:"总报名数" form:"totalJoin" gorm:"column:total_join;comment:总报名数;type:int"`
	TotalDiscuss *int       `json:"totalDiscuss" cn:"总评论" form:"totalDiscuss" gorm:"column:total_discuss;comment:总评论;type:int"`
	TotalClick   *int       `json:"totalClick" cn:"总点击" form:"totalClick" gorm:"column:total_click;comment:总点击;type:int"`
	TotalStar    *int       `json:"totalStar" cn:"总评" form:"totalStar" gorm:"column:total_star;comment:总评;type:int"`
	TotalStar1   *int       `json:"totalStar1" cn:"总星评1" form:"totalStar1" gorm:"column:total_star_1;comment:总星评1:教学水平）;type:smallint"`
	TotalStar2   *int       `json:"totalStar2" cn:"总星评2" form:"totalStar2" gorm:"column:total_star_2;comment:总星评2:课程质量）;type:smallint"`
	TotalStar3   *int       `json:"totalStar3" cn:"总星评3" form:"totalStar3" gorm:"column:total_star_3;comment:总星评3:服务态度）;type:smallint"`
	Status       *int       `json:"status" cn:"状态" form:"status" gorm:"column:status;comment:状态:0未审核 1审核通过 2未通过审核 3草稿;type:smallint"`
	StatusMsg    string     `json:"statusMsg" cn:"审核原因" form:"statusMsg" gorm:"column:status_msg;comment:审核原因;type:varchar(500);"`
}

// TableName ActAct 表名
func (ActAct) TableName() string {
	return "act_act"
}
