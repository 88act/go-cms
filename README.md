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

 

## 1. 基本介绍

### 1.1 项目介绍

> go-cms, golang的cms内容管理系统, 管理后台web-admin居于 vue3 element-plus ,客户端居于uniapp,后端居于 golang gin  后端集成jwt鉴权，动态路由，动态菜单，casbin鉴权功能.

- 前端api接口，居于gozero 微服务框架，主要考虑性能和服务可用性，应对高并发的client请求。
本项目提供了gozero 和gokit的微服务demo代码. 

- 后端api接口，使用的是gin框架的单体服务模式，使用代码生成器，基本能生成全功能代码，兼顾开发效率
，后端单体服务api已可以部署到k8s容器，实现负载动态扩展。

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
go generate 
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
  
#### 2.3.2 生成API文档

```` shell
cd server
swag init
````

> 执行上面的命令后，server目录下会出现docs文件夹里的 `docs.go`, `swagger.json`, `swagger.yaml` 三个文件更新，启动go服务之后, 在浏览器输入 [http://localhost:40040/swagger/index.html](http://localhost:40040/swagger/index.html) 即可查看swagger文档


## 3. 技术选型

- 前端：用基于vue3 element-plus构建基础页面。
- 后端：golang gin gormv2  也使用了部分其他开源项目模块,如 gofarme  gav celly
- 开发热更新: air 。
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
 
## 6. 商用注意事项

如果您将此项目用于商业用途，请遵守Apache2.0协议 


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
 