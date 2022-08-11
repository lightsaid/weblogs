package routes

import (
	"net/http"

	"lightsaid.com/weblogs/cmd/web/handlers"
)

type Router struct {
	Path         string
	Handler      func(w http.ResponseWriter, r *http.Request)
	Method       string
	AuthRequired bool
}

// 用户路由
var userRoutes = []Router{
	{
		Path:         "/users",
		Handler:      handlers.GetUsers,
		Method:       http.MethodGet,
		AuthRequired: true,
	},
	{
		Path:         "/users",
		Handler:      handlers.CreateUser,
		Method:       http.MethodPost,
		AuthRequired: false,
	},
	{
		Path:         "/users/{id:[0-9]+}",
		Handler:      handlers.GetUser,
		Method:       http.MethodGet,
		AuthRequired: false,
	},
	{
		Path:         "/users/{id:[0-9]+}",
		Handler:      handlers.UpdateUser,
		Method:       http.MethodPut,
		AuthRequired: true,
	},
	{
		Path:         "/users/{id:[0-9]+}",
		Handler:      handlers.DeleteUser,
		Method:       http.MethodDelete,
		AuthRequired: true,
	},
}

// 文章路由
var postRoutes = []Router{
	{
		Path:         "/posts",
		Handler:      handlers.GetPosts,
		Method:       http.MethodGet,
		AuthRequired: true,
	},
	{
		Path:         "/posts",
		Handler:      handlers.CreatePost,
		Method:       http.MethodPost,
		AuthRequired: false,
	},
	{
		Path:         "/posts/{id:[0-9]+}",
		Handler:      handlers.GetPost,
		Method:       http.MethodGet,
		AuthRequired: false,
	},
	{
		Path:         "/posts/{id:[0-9]+}",
		Handler:      handlers.UpdatePost,
		Method:       http.MethodPut,
		AuthRequired: true,
	},
	{
		Path:         "/posts/{id:[0-9]+}",
		Handler:      handlers.DeletePost,
		Method:       http.MethodDelete,
		AuthRequired: true,
	},
}

// 分类路由
var cateRoutes = []Router{
	{
		Path:         "/categories",
		Handler:      handlers.GetCategories,
		Method:       http.MethodGet,
		AuthRequired: true,
	},
	{
		Path:         "/categories",
		Handler:      handlers.CreateCategory,
		Method:       http.MethodPost,
		AuthRequired: false,
	},
	{
		Path:         "/categories/{id:[0-9]+}",
		Handler:      handlers.GetCategory,
		Method:       http.MethodGet,
		AuthRequired: false,
	},
	{
		Path:         "/categories/{id:[0-9]+}",
		Handler:      handlers.UpdateCategory,
		Method:       http.MethodPut,
		AuthRequired: true,
	},
	{
		Path:         "/categories/{id:[0-9]+}",
		Handler:      handlers.DeleteCategory,
		Method:       http.MethodDelete,
		AuthRequired: true,
	},
}

// 属性路由
var attrRoutes = []Router{
	{
		Path:         "/attrs",
		Handler:      handlers.GetAttributes,
		Method:       http.MethodGet,
		AuthRequired: true,
	},
	{
		Path:         "/attrs",
		Handler:      handlers.CreateAttribute,
		Method:       http.MethodPost,
		AuthRequired: false,
	},
	{
		Path:         "/attrs/{id:[0-9]+}",
		Handler:      handlers.GetAttributes,
		Method:       http.MethodGet,
		AuthRequired: false,
	},
	{
		Path:         "/attrs/{id:[0-9]+}",
		Handler:      handlers.UpdateAttribute,
		Method:       http.MethodPut,
		AuthRequired: true,
	},
	{
		Path:         "/categories/{id:[0-9]+}",
		Handler:      handlers.DeleteCategory,
		Method:       http.MethodDelete,
		AuthRequired: true,
	},
}

// 评论路由
