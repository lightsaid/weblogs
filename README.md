# 博客
这是个人博客项目，开发环境， window10 + wsl2 + docker；
采用了前后端不分离开发模式。

### 启动
- 启动要求
	1. 数据库使用 sqlite3，因此需要安装sqlite3
	1. 修改 .env 文件密钥相关的 CSRF_SECRET SESSION_KEY

- 启动项目
	- 1. 如果是有安装make, 则可以使用 `make start` 命令
	- 1. 或者 `go run ./cmd/web` 

- 所有页面内容必须是管理员才能看见, 关于数据库设计，请查看 resources/database 目录内容

- 项目还在开发完善中...

### 主要技术栈
	- Golang	
	- Gorilla Toolkit
	- Sqlite3
	- sqlx
	- Bootstrap5
	- editor.md

### 目录结构设计
``` js
weblogs
	cmd/web
		routes
			router.go // 定义路由
			routes.go // 实现路由
		middleware
			middleware.go // 实现所有中间件
		handlers // 处理接口请求，与service层对接
			handlers.go
			// ...
		errors  // 请求响应公共方法
			errors.go
		main.go // 程序启动入口
	internal
		models
			models.go // 定义所有 model
		// 服务, 与handlers对接，调用crud方法处理业务返回给handlers层, ->>> 弱化 (只有再特别复杂的情况才在service层处理)
		service 
			request.go // 请求入参定义
			response.go // 请求出参定义 
			// ...
		repository
			repository.go // 
			dbrepo
				// ... 实现 models crud
	pkg
		validator
		logger // logger 配置全局使用zap
	resources // 存放资源
		docker
			docker-compose.yml
		database
			database.sql // 建库建表 sql
			001_modify_user_up.sql
			001_modify_user_down.sql
			// ...
		docs // 文档	
```