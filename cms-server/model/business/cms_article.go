// 自动生成模板CmsArticle
package business

import (
	"github.com/88act/go-cms/server/global"
)

// CmsArticle 结构体
// 如果含有time.Time 请自行import time包
type CmsArticle struct {
     CmsArticleMini
      Sketch  string `json:"sketch" form:"sketch" gorm:"column:sketch;comment:文章摘要;type:varchar(500);"`
      Detail  string `json:"detail" form:"detail" gorm:"column:detail;comment:文章内容:(以后需要分表);type:mediumtext;"`
      Qrcode  string `json:"qrcode" form:"qrcode" gorm:"column:qrcode;comment:二维码图片;type:varchar(255);"`
      ImageList  string `json:"imageList" form:"imageList" gorm:"column:image_list;comment:图片列表:id列 1,2,3;type:varchar(255);"`
      MediaList  string `json:"mediaList" form:"mediaList" gorm:"column:media_list;comment:媒体列表:id列 1,2,3;type:varchar(255);"`
      Link  string `json:"link" form:"link" gorm:"column:link;comment:链接地址;type:varchar(255);"`
      TotalStar1  *int `json:"totalStar1" form:"totalStar1" gorm:"column:total_star_1;comment:总星评1;type:bigint"`
      TotalStar2  *int `json:"totalStar2" form:"totalStar2" gorm:"column:total_star_2;comment:总星评2;type:bigint"`
      TotalStar3  *int `json:"totalStar3" form:"totalStar3" gorm:"column:total_star_3;comment:总星评3;type:bigint"`
      TotalStar4  *int `json:"totalStar4" form:"totalStar4" gorm:"column:total_star_4;comment:总星评4;type:bigint"`
      TotalStar5  *int `json:"totalStar5" form:"totalStar5" gorm:"column:total_star_5;comment:总星评5;type:bigint"`
      Chapter  string `json:"chapter" form:"chapter" gorm:"column:chapter;comment:章节集合:json_id;type:varchar(500);"`
      VerifyMsg  string `json:"verifyMsg" form:"verifyMsg" gorm:"column:verify_msg;comment:审核信息;type:varchar(255);"`
}

// CmsArticleMini 结构体
type CmsArticleMini struct {
      global.GO_MODEL
      UserId  *int `json:"userId" form:"userId" gorm:"column:user_id;comment:用户id;type:bigint"`
      CatId  *int `json:"catId" form:"catId" gorm:"column:cat_id;comment:类别ID;type:int"`
      CatIdSys  *int `json:"catIdSys" form:"catIdSys" gorm:"column:cat_id_sys;comment:系统类别ID;type:int"`
      MediaType  *int `json:"mediaType" form:"mediaType" gorm:"column:media_type;comment:文章类型;type:smallint"`
      Name  string `json:"name" form:"name" gorm:"column:name;comment:文章标题;type:varchar(150);"`
      Author  string `json:"author" form:"author" gorm:"column:author;comment:文章作者;type:varchar(100);"`
      TagList  string `json:"tagList" form:"tagList" gorm:"column:tag_list;comment:标签列表:id1,id2,id3,;type:varchar(500);"`
      AuthorEmail  string `json:"authorEmail" form:"authorEmail" gorm:"column:author_email;comment:作者邮箱;type:varchar(60);"`
      Referer  string `json:"referer" form:"referer" gorm:"column:referer;comment:来源;type:varchar(255);"`
      Thumb  string `json:"thumb" form:"thumb" gorm:"column:thumb;comment:插图;type:单图片"`
      IsTop  *bool `json:"isTop" form:"isTop" gorm:"column:is_top;comment:置顶:0否 1是）;type:tinyint"`
      IsHot  *bool `json:"isHot" form:"isHot" gorm:"column:is_hot;comment:热门:(0否  1是);type:tinyint"`
      TotalDiscuss  *int `json:"totalDiscuss" form:"totalDiscuss" gorm:"column:total_discuss;comment:总评论;type:int"`
      TotalClick  *int `json:"totalClick" form:"totalClick" gorm:"column:total_click;comment:总点击;type:int"`
      TotalStar  *int `json:"totalStar" form:"totalStar" gorm:"column:total_star;comment:总评;type:int"`
      Sort  *int `json:"sort" form:"sort" gorm:"column:sort;comment:排序;type:int"`
      Pid  *int `json:"pid" form:"pid" gorm:"column:pid;comment:父id:章节集合;type:bigint"`
      Status  *int `json:"status" form:"status" gorm:"column:status;comment:状态:0未审核 1审核  2未通过审核 3 草稿;type:smallint"`

}


// TableName CmsArticle 表名
func (CmsArticle) TableName() string {
  return "cms_article"
}

