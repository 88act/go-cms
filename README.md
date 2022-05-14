<div align=center>
<img src="https://cms.88act.com/res/img/go-cms.png" width=300" height="300" />
</div>
<div align=center>
<img src="https://img.shields.io/badge/golang-1.16-blue"/>
<img src="https://img.shields.io/badge/gin-1.6.3-blue"/>
<img src="https://img.shields.io/badge/vue-3.0.0-blue"/> 
 
</div>

[English](./README-en.md) | 简体中文

 
[github地址](https://github.com/88act/go-cms): https://github.com/88act/go-cms
 
 

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

[在线预览](http://120.24.150.47/admin)   
http://120.24.150.47/admin
测试用户名：test123
测试密码：test123 
                                                      
                                                      
## /gozero/微服务项目介绍 

>  本项目是 gozero 的 gorm v2 版本 ,欢迎 Star
-  使用 gormv2 版本操作数据库 ,可能更合适熟悉了gorm v2不想换sqlx,又想享受高效率的的gozero微服务框架的朋友 
 
-  项目背景:居于互联网PC和小程序的活动项目,包括活动发布,活动报名,活动下单/支付等功能

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
 
- 6、 完成测试

####  比较复杂的启动方式 

- 1、 先启动 docker , 然后  
  cd /go-cms/gozero/    
                                                                                                                 
$ docker network create go-zero-looklook_looklook_net                                                                                                                      
$ docker-compose -f docker-compose-env.yml up -d   

- 2、  然后,按照最简单的方式启动

- 3、 本项目docker只启动了基本的 jaeger链路追踪, asynq延迟队列 redis,elasticsearch,mysql 等
  如果想启动所有的服务 ,可以参考 https://github.com/Mikaelemmmm/go-zero-looklook  项目里面的详细说明

- 4、  访问 http://127.0.0.1:16686/search  可以查看链路追踪的情况
####  
 
####  =============================================================================================================================================


### 1.1 gin gorm v2 管理后台项目

> go-cms, golang的cms内容管理系统, 管理后台web-admin居于 vue3 element-plus ,客户端居于uniapp,后端居于 golang gin  后端集成jwt鉴权，动态路由，动态菜单，casbin鉴权功能. 

- 后端api接口，使用的是gin框架的单体服务模式，使用代码生成器，基本能生成全功能代码，兼顾开发效率
，后端单体服务api也可以部署到k8s容器，实现负载动态扩展。

- 集成强大的前后端代码生成器,可以生成前后端所有代码，包括 api vue js 。可以完成复杂的查询，排序，增删改操作

-  集成文件库功能,统一管理图片/视频/文件等附件,自动对比文件哈希码,避免相同的文件多次上传

-  集成tinyEditor 富媒体编辑器

- 集成居于celly的数据采集,网络爬虫系统
 
-  集成腾讯IM消息记录定时下载功能

-  集成k8s容器管理器（进行中...）

 
[在线预览](http://120.24.150.47/admin)   
http://120.24.150.47/admin
测试用户名：test123
测试密码：test123 
 

### 2.1 server项目

使用 `Goland`或 vscode 编辑工具，打开server目录 
```bash

# 克隆项目
git clone https://github.com/88act/go-cms.git
# 进入server文件夹
cd server 
# 使用 go mod 并安装go依赖包
go mod tidy  
# 编译 
go build -o server main.go (windows编译命令为go build -o server.exe main.go ) 
# 运行二进制
./server (windows运行命令为 server.exe)
```

### 2.2 web项目

```bash
# 进入web文件夹
cd web

# 安装依赖
cnpm install || npm install

# 启动web项目
npm run serve
```

### 2.3 swagger自动化API文档

#### 2.3.1 安装 swagger

 
````
go get -u github.com/swaggo/swag/cmd/swag
````
 

## 3. 技术  
- 前端：用基于vue3 element-plus构建基础页面。
- 后端：golang gin gormv2  也使用了部分其他开源项目模块,如 gofarme  gav celly 
- 数据库： MySql`(8.0) 
- 缓存： Redis` 
- API文档：使用`Swagger`构建自动化文档。
 

### 4. 技术群

### QQ交流群：555475796
| QQ 群 |
|  :---:  |
| <img src="https://cms.88act.com/res/img/qq.jpg" width="180"/> |
 

### [关于我们] 

## 5. 贡献者
 
 

## 7. 附系统截图
  代码生成与字段配置  
| <img src="https://cms.88act.com/res/img/gocms/g1.jpg"  /> |

集成媒体库管理 ,集中管理上传文件 ,文件分模块管理,文件上传自动对比哈希值,避免重复
| <img src="https://cms.88act.com/res/img/gocms/g2.jpg"  /> |

集成富文本编辑器 ,集成媒体库选择文件
| <img src="https://cms.88act.com/res/img/gocms/g3.jpg"  /> |

新闻采集器
| <img src="https://cms.88act.com/res/img/gocms/g4.jpg"  /> |

k8s集群管理(进行中....)
| <img src="https://cms.88act.com/res/img/gocms/g5.jpg"  /> |

腾讯IM历史消息下载器
| <img src="https://cms.88act.com/res/img/gocms/g6.jpg"  /> |
 



## 感谢以下优秀的项目 


go-zero 微服务: https://github.com/zeromicro/go-zero

go-zero-looklook   https://github.com/Mikaelemmmm/go-zero-looklook

dtm分布式事务：https://github.com/dtm-labs/dtm

gva项目  ：https://github.com/flipped-aurora/gin-vue-admin

  
 
