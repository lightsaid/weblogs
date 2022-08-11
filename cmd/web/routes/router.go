package routes

import (
	"net/http"

	"lightsaid.com/weblogs/cmd/web/handlers"
)

type Router struct {
	Path string
	Handler func(w http.ResponseWriter, r *http.Request)
	Method string
	AuthRequired bool
}

// 用户路由
var userRoutes = []Router{
	{
		Path: "/users",
		Handler: handlers.AppH.GetUsers,
		Method: http.MethodGet,
		AuthRequired: true,
	},
	{
		Path: "/users",
		Handler: handlers.AppH.CreateUser,
		Method: http.MethodPost,
		AuthRequired: false,
	},
	{
		Path: "/users/{id:[0-9]+}",
		Handler: handlers.AppH.GetUser,
		Method: http.MethodGet,
		AuthRequired: false,
	},
	{
		Path: "/users/{id:[0-9]+}",
		Handler: handlers.AppH.UpdateUser,
		Method: http.MethodPut,
		AuthRequired: true,
	},
	{
		Path: "/users/{id:[0-9]+}",
		Handler: handlers.AppH.DeleteUser,
		Method: http.MethodDelete,
		AuthRequired: true,
	},
}

// 文章路由
var postRoutes = []Router{
	{
		Path: "/posts",
		Handler: handlers.AppH.GetPosts,
		Method: http.MethodGet,
		AuthRequired: true,
	},
	{
		Path: "/posts",
		Handler: handlers.AppH.CreatePost,
		Method: http.MethodPost,
		AuthRequired: false,
	},
	{
		Path: "/posts/{id:[0-9]+}",
		Handler: handlers.AppH.GetPost,
		Method: http.MethodGet,
		AuthRequired: false,
	},
	{
		Path: "/posts/{id:[0-9]+}",
		Handler: handlers.AppH.UpdatePost,
		Method: http.MethodPut,
		AuthRequired: true,
	},
	{
		Path: "/posts/{id:[0-9]+}",
		Handler: handlers.AppH.DeletePost,
		Method: http.MethodDelete,
		AuthRequired: true,
	},
}


// 分类路由
var cateRoutes = []Router{
	{
		Path: "/categories",
		Handler: handlers.AppH.GetCategories,
		Method: http.MethodGet,
		AuthRequired: true,
	},
	{
		Path: "/categories",
		Handler: handlers.AppH.CreateCategory,
		Method: http.MethodPost,
		AuthRequired: false,
	},
	{
		Path: "/categories/{id:[0-9]+}",
		Handler: handlers.AppH.GetCategory,
		Method: http.MethodGet,
		AuthRequired: false,
	},
	{
		Path: "/categories/{id:[0-9]+}",
		Handler: handlers.AppH.UpdateCategory,
		Method: http.MethodPut,
		AuthRequired: true,
	},
	{
		Path: "/categories/{id:[0-9]+}",
		Handler: handlers.AppH.DeleteCategory,
		Method: http.MethodDelete,
		AuthRequired: true,
	},
}

// 属性路由
var attrRoutes = []Router{
	{
		Path: "/attrs",
		Handler: handlers.AppH.GetAttributes,
		Method: http.MethodGet,
		AuthRequired: true,
	},
	{
		Path: "/attrs",
		Handler: handlers.AppH.CreateAttribute,
		Method: http.MethodPost,
		AuthRequired: false,
	},
	{
		Path: "/attrs/{id:[0-9]+}",
		Handler: handlers.AppH.GetAttributes,
		Method: http.MethodGet,
		AuthRequired: false,
	},
	{
		Path: "/attrs/{id:[0-9]+}",
		Handler: handlers.AppH.UpdateAttribute,
		Method: http.MethodPut,
		AuthRequired: true,
	},
	{
		Path: "/categories/{id:[0-9]+}",
		Handler: handlers.AppH.DeleteCategory,
		Method: http.MethodDelete,
		AuthRequired: true,
	},
}


// 评论路由
