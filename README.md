<div align=center>
<img src="https://cms.88act.com/res/img/go-cms.png" width=300" height="300" />
</div>
<div align=center>
<img src="https://img.shields.io/badge/golang-1.21-blue"/>
<img src="https://img.shields.io/badge/go-zero-blue"/>
<img src="https://img.shields.io/badge/gin-1.6.3-blue"/>
<img src="https://img.shields.io/badge/vue-3.3.7-blue"/> 
<img src="https://img.shields.io/badge/vite-4.5.0-blue"/> 
<img src="https://img.shields.io/badge/typescript-5.2.2-blue"/> 
 
</div>

[English](./README-en.md) | 简体中文

 
[github地址](https://github.com/88act/go-cms): https://github.com/88act/go-cms
 
go-cms v2.0 新版 升级到 go 1.21    
本项目是一个 cms/企业官网/文章/新闻/群组/圈子 通用内容管理系统 
演示golang gin  结合 vue3 的管理后台，以及演示居于gozero+gorm的微服务前端api项目 
vue3 使用最新的TypeScript、Pinia、Tailwindcss，Vite、Element-Plus技术。

> 在线体验： 
- 管理后台: https://120.79.55.204/admin/  账号admin  密码123456    

 
## 1. 基本介绍

> 本项目 目录介绍
-  /gozero/ 居于最新版 gozero v1.60，结合gorm v2 微服务项目,作为前端api
-  /uniapp/ 配套的前端的uniapp项目,可发布成小程序或h5，hybird程序 
-  /server/ 居于gin的 gorm v2 管理后台项目
-  /web/ 配套的管理后台vue3，居于 Vue3、TypeScript、Pinia、Tailwindcss，Vite、Element-Plus、 等主流技术栈开发，支持多语言国际化。
-  /gokit/ gokit微服务器测试项目
 


## 2.  管理后台介绍
go-cms, golang的cms内容管理系统,
  管理后台web端， 居于 vue3 element-plus ，TypeScript、Pinia,
  客户端居于uniapp,
  后端居于 golang gin，  
  集成jwt鉴权，动态路由，动态菜单，casbin鉴权功能. 

-  集成文件库功能, 统一管理图片/视频/文件等附件,自动对比文件哈希码,避免相同的文件多次上传

-  集成 markdown 编辑器 ，生成静态md文件和html 文件，方便搜索引擎收集

-  集成html模板功能，直接服务器端生成html 页面 方便搜索引擎收集

-  集成居于celly的数据采集,网络爬虫系统
                                                      
                                                      
## 3.  /gozero/ 客户端API微服务后台 

>  本项目居于最新版 gozero v1.6 + gorm v2 版本 
-  使用 gormv2 版本操作数据库 ,可能更合适熟悉了gorm v2不想换sqlx,又想享受高效率的的gozero微服务框架的朋友   
-  本项目没使用 gozero原来的缓存设计, 本项目使用 api层的传入参数作为缓存key ，根据具体业务需要,灵活设计缓存的时间 ,直接从缓存返回结果。 相比较gozero的单记录缓存设计， 居于api传入参数作为key的缓存方式，性能更好，速度更快。  
 
-  待开源部分   
- [x] 1、自动生成gozero gorm v2版本 model数据库操作层代码，proto文件, api文件, 进一步解放生产力 
- [x] 2、自动生成 gin管理后台的go代码, vue3,ts前端代码 
  

####  最简单的启动模式
- 1、  克隆本项目
    git clone https://github.com/88act/go-cms.git

- 2、 数据库文件/server/文档/go-cms.sql, 还原数据库go-cms.sql, 到本机的mysql 数据库

- 3、 启动gin服务器 
      cd /server/
      go run main.go 

- 4、 启动web vue
      cd /web/ 
      pnpm install 
      pnpm dev 

- 5、 启动gozero  

    cd gozero

    ./modd  
 




####  gocms 管理后台截图
![gocms](https://120.79.55.204/res/img/cms1.jpg)
![gocms](https://120.79.55.204/res/img/cms2.jpg)

<img src="https://120.79.55.204/res/img/cms1.jpg" width = "70%" />
<img src="https://120.79.55.204/res/img/cms2.jpg" width = "70%" />
<img src="https://120.79.55.204/res/img/cms3.jpg" width = "70%" />
<img src="https://120.79.55.204/res/img/cms4.jpg" width = "70%" />
<img src="https://120.79.55.204/res/img/cms5.jpg" width = "70%" />
<img src="https://120.79.55.204/res/img/cms6.jpg" width = "70%" />

