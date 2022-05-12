

## go-zero -gocms

 https://github.com/88act/looklook/gozero/README.md) | 简体中文



## 前言

本项目是 gozero 的 gorm v2 版本 ,欢迎 Star
超过 500 star,将放出配套的  uniapp 版本 和 缓存版本

能根据数据库表, 

自动生成 gozero 项目需要的 gorm版本 model数据库操作层, proto文件, api文件, 进一步解放生产力 

自动生成 管理后台的go代码 ,和vue 代码

使用 gormv2 版本操作数据库 ,可能更合适大多数学习了gorm又不想转换sqlx的朋友,来使用gozero框架

没有了gozero原来的缓存设计, 后面,本项目会增加居于 api层的缓存和rpc层缓存的双层缓存设计,减少rpc的调用耗时 , 
根据api接口和参数, 作为缓存的key ,根据具体业务需要,来设计缓存的时间 

超过 500 star,将更新配套的 uniapp 前端版本 和 缓存版本 

 

####  最简单的启动模式

1  还原数据库,到本机的mysql 数据库,修改etc文件夹的 DB配置项, 使之正常连接数据库 

2 







####  
 

 
## 感谢


go-zero 微服务: https://github.com/zeromicro/go-zero

go-zero-looklook   https://github.com/Mikaelemmmm/go-zero-looklook

dtm分布式事务：https://github.com/dtm-labs/dtm

gva项目  ：https://github.com/flipped-aurora/gin-vue-admin

  
 
