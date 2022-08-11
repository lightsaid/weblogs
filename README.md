# 博客


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
		service // 服务, 与handlers对接，调用crud方法处理业务返回给handlers层
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