

## go-zero -gocms

 https://github.com/88act/go-cms/gozero/README.md) | 简体中文



# 重要提示

> 项目特点: 
 -  自动生成 gozero 项目需要的 gorm v2版本 model数据库操作层, proto文件, api文件,
 -  自动生成 管理后台的go代码 ,和vue 代码 
 -  超过 500 star,将更新配套的 uniapp 前端版本 和 缓存版本。  欢迎star来鼓励。。。。

## 1. 基本介绍

> 本项目 目录介绍
-  /gozero/ 的 gorm v2 版本的微服务项目,作为项目的前端api
-  /uniapp/ 前端的uniapp项目,发布小程序或h5 
-  /server/的gin gorm v2 管理后台项目
-  /web/ 是gin gorm v2管理后台项目配套的vue的项目
-  /gokit/ gokit微服务器测试项目


## /gozero/微服务项目介绍 

>  本项目是 gozero 的 gorm v2 版本 ,欢迎 Star
-  使用 gormv2 版本操作数据库 ,可能更合适熟悉了gorm v2不想换sqlx,又想享受高效率的的gozero微服务框架的朋友 
 
-  项目背景:居于互联网PC和小程序的活动圈项目,包括活动发布,活动报名,活动下单/支付等功能

-  本项目特点，能根据数据库表, 

-  1、自动生成 gozero 项目需要的 gorm版本 model数据库操作层, proto文件, api文件,  进一步解放生产力 

-  2、自动生成 管理后台的go代码 ,和vue 代码 

-  3、api和rpc 双层redis缓存设计，根据业务需求灵活设置缓存时长和键值（开发中....） 


-  没有gozero原来的缓存设计, 本项目后继将增加居于 api层的缓存和rpc层缓存的双层缓存设计,减少调用rpc的耗时, 
根据api接口和参数, 作为缓存的key ,根据具体业务需要,灵活设计缓存的时间 

-  超过 500 star,将更新配套的 uniapp 前端版本 和 缓存版本。  欢迎star来鼓励。。。。

 

####  最简单的启动模式
- 1、  克隆本项目
    git clone https://github.com/88act/go-cms.git

- 2、  还原数据库,到本机的mysql 数据库,  数据库文件在文件夹  go-cms/doc/sql/  ,四个模块的表放在同一个数据库,方便测试,实际部署时可以分库
    -  微服务器项目包括四个模块: usercenter 用户中心 ,act 活动模块 ,order 订单模块 pay 支付模块 
    -  对应的 数据库前缀分别是 : mem_ , act_ , order_, pay_   
 
    - 修改app目录的四个etc文件夹的 DB配置项, 使之正常连接数据库 ,  比如  act.yaml 文件 ，修改自己的数据库配置 ，比如 ： 
       DataSource: root:123456@tcp(mysql:53306)/go-cms?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai

- 3、 启动4个微服务器,本项目使用 modd 启动和管理 ,  关于modd更多用法可以去这里了解 ： https://github.com/cortesi/modd
    cd gozero
    ./modd  

- 4、 正常可以看到 api 和 rpc 已经正常运行

- 5、 测试 ,启动 postman, 运行下面post, 终端可以观察到gorm sql语句的执行情况

post  http://127.0.0.1:1004/usercenter/v1/user/register 

参数选body json 格式 :   {"mobile":"13088886666","password":"123456"}
 
post  http://127.0.0.1:1004/usercenter/v1/user/login 

参数选body json 格式 :  {"mobile":"13088886666","password":"123456"}
 

post  http://127.0.0.1:1001/order/v1/order/addOrder

参数选body json 格式 : {"orderType":2,"objId":3,"title":"3","price":1100,"priceSrc":1210,"couponId":121,"remark":"remar3333"}
 
6、 完成测试

####  比较复杂的方式 

1、 先启动 docker , 然后  
$ docker network create go-zero-looklook_looklook_net
$ docker-compose -f docker-compose-env.yml up -d   

2、  然后,按照最简单的方式启动

3、 本项目docker只启动了基本的 jaeger链路追踪, asynq延迟队列 redis,elasticsearch,mysql 等
  如果想启动所有的服务 ,可以参考 https://github.com/Mikaelemmmm/go-zero-looklook  项目里面的详细说明

4、  访问 http://127.0.0.1:16686/search  可以查看链路追踪的情况
####  
 

 






## 感谢以下优秀的项目 


go-zero 微服务: https://github.com/zeromicro/go-zero

go-zero-looklook   https://github.com/Mikaelemmmm/go-zero-looklook

dtm分布式事务：https://github.com/dtm-labs/dtm

gva项目  ：https://github.com/flipped-aurora/gin-vue-admin

  
 
