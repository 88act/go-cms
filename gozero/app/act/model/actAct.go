// 自动生成模板ActAct
package model

import (
	"time"
)

// ActAct 结构体
type ActAct struct {
	BaseModel
	UserId       int64      `db:"user_id" json:"userId" gorm:"column:user_id"`                   //用户id
	CatId        int        `db:"cat_id" json:"catId" gorm:"column:cat_id"`                      //系统栏目
	CatIdUser    int        `db:"cat_id_user" json:"catIdUser" gorm:"column:cat_id_user"`        //用户栏目
	BeOnline     bool       `db:"be_online" json:"beOnline" gorm:"column:be_online"`             //类型
	ActType      string     `db:"act_type" json:"actType" gorm:"column:act_type"`                //活动类型
	BeMulti      bool       `db:"be_multi" json:"beMulti" gorm:"column:be_multi"`                //分期活动
	Phase        int        `db:"phase" json:"phase" gorm:"column:phase"`                        //分期
	Title        string     `db:"title" json:"title" gorm:"column:title"`                        //标题
	Desc         string     `db:"desc" json:"desc" gorm:"column:desc"`                           //简介
	Avater       string     `db:"avater" json:"avater" gorm:"column:avater"`                     //缩略图
	MediaList    string     `db:"media_list" json:"mediaList" gorm:"column:media_list"`          //媒体列表
	Address      string     `db:"address" json:"address" gorm:"column:address"`                  //地址
	AreaId       int        `db:"area_id" json:"areaId" gorm:"column:area_id"`                   //地区id
	Lng          int64      `db:"lng" json:"lng" gorm:"column:lng"`                              //经度
	Lat          int64      `db:"lat" json:"lat" gorm:"column:lat"`                              //纬度
	BeginTime    *time.Time `db:"begin_time" json:"beginTime" gorm:"column:begin_time"`          //开始时间
	EndTime      *time.Time `db:"end_time" json:"endTime" gorm:"column:end_time"`                //结束时间
	LiveTime     *time.Time `db:"live_time" json:"liveTime" gorm:"column:live_time"`             //直播开始
	LiveEnd      *time.Time `db:"live_end" json:"liveEnd" gorm:"column:live_end"`                //直播结束
	Presenter    string     `db:"presenter" json:"presenter" gorm:"column:presenter"`            //主持/讲师
	PhoneKefu    string     `db:"phone_kefu" json:"phoneKefu" gorm:"column:phone_kefu"`          //客服电话
	PhoneHezu    string     `db:"phone_hezu" json:"phoneHezu" gorm:"column:phone_hezu"`          //合作电话
	Wx           string     `db:"wx" json:"wx" gorm:"column:wx"`                                 //微信
	Qq           string     `db:"qq" json:"qq" gorm:"column:qq"`                                 //QQ
	GroupId      string     `db:"group_id" json:"groupId" gorm:"column:group_id"`                //群组id
	EndJoinTime  *time.Time `db:"end_join_time" json:"endJoinTime" gorm:"column:end_join_time"`  //结束报名
	Price        int        `db:"price" json:"price" gorm:"column:price"`                        //票价
	PriceVip     int        `db:"price_vip" json:"priceVip" gorm:"column:price_vip"`             //vip票价
	PriceDesc    string     `db:"price_desc" json:"priceDesc" gorm:"column:price_desc"`          //票价描述
	RefundType   int        `db:"refund_type" json:"refundType" gorm:"column:refund_type"`       //退费类型
	RefundDays   int        `db:"refund_days" json:"refundDays" gorm:"column:refund_days"`       //退费天
	BeShowjoin   int        `db:"be_showjoin" json:"beShowjoin" gorm:"column:be_showjoin"`       //是否显示人数
	StatusAct    int        `db:"status_act" json:"statusAct" gorm:"column:status_act"`          //活动状态
	MaxJoin      int        `db:"max_join" json:"maxJoin" gorm:"column:max_join"`                //最大报名人数
	NowJoin      int        `db:"now_join" json:"nowJoin" gorm:"column:now_join"`                //当前报名人数
	BeChosen     bool       `db:"be_chosen" json:"beChosen" gorm:"column:be_chosen"`             //是否精选
	BeWeek       bool       `db:"be_week" json:"beWeek" gorm:"column:be_week"`                   //是否周末
	BeVip        bool       `db:"be_vip" json:"beVip" gorm:"column:be_vip"`                      //是否会员
	TotalWhole   int        `db:"total_whole" json:"totalWhole" gorm:"column:total_whole"`       //综合指数
	TotalShare   int        `db:"total_share" json:"totalShare" gorm:"column:total_share"`       //总分享
	TotalFav     int        `db:"total_fav" json:"totalFav" gorm:"column:total_fav"`             //总收藏
	TotalJoin    int        `db:"total_join" json:"totalJoin" gorm:"column:total_join"`          //总报名数
	TotalDiscuss int        `db:"total_discuss" json:"totalDiscuss" gorm:"column:total_discuss"` //总评论
	TotalClick   int        `db:"total_click" json:"totalClick" gorm:"column:total_click"`       //总点击
	TotalStar    int        `db:"total_star" json:"totalStar" gorm:"column:total_star"`          //总评
	TotalStar1   int        `db:"total_star_1" json:"totalStar1" gorm:"column:total_star_1"`     //总星评1
	TotalStar2   int        `db:"total_star_2" json:"totalStar2" gorm:"column:total_star_2"`     //总星评2
	TotalStar3   int        `db:"total_star_3" json:"totalStar3" gorm:"column:total_star_3"`     //总星评3
	Status       int        `db:"status" json:"status" gorm:"column:status"`                     //状态
	StatusMsg    string     `db:"status_msg" json:"statusMsg" gorm:"column:status_msg"`          //审核原因

}

// TableName ActAct 表名
func (ActAct) TableName() string {
	return "act_act"
}

// ActActSearch  查询
type ActActSearch struct {
	ActAct
	PageInfo
}
