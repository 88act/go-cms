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

> go-cms, golang的cms内容管理系统, 前端web-admin居于 vue3 element-plus ,客户端居于uniapp,后端居于 golang gin  后端集成jwt鉴权，动态路由，动态菜单，casbin鉴权功能.
- 同时,系统集成 客服聊天websocket系统, 
- 集成居于celly的数据采集,网络爬虫系统,
- 集成强大的前后端代码生成器.
- go-cms管理后台,集成文件库功能,统一管理图片/视频/文件等附件,自动对比文件哈希码,避免相同的文件多次上传
- go-cms管理后台,集成tinyEditor 富媒体编辑器
- golang后台可运行于 微信云托管容器,阿里云托管容器


 
[在线预览](http://go-cms.88act.com/admin)   
http://go-cms.88act.com/admin
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

### 微信交流群
| 微信 |
|  :---:  | 
| <img width="150" src="https://cms.88act.com/res/img/wx.png"> 

添加微信，备注"加入go-cms交流群"

### [关于我们](https://cms.88act.com/about/)

## 5. 贡献者
 
## 6. 商用注意事项

如果您将此项目用于商业用途，请遵守Apache2.0协议 
