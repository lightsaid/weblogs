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

// 访问管理员模板文件路由
var adminPageRoutes = []Router{
	{
		Path:         "/admin/register",
		Handler:      handlers.ShowAdminRegister,
		Method:       http.MethodGet,
		AuthRequired: false,
	},
	{
		Path:         "/admin/register",
		Handler:      handlers.PostAdminRegister,
		Method:       http.MethodPost,
		AuthRequired: false,
	},
	{
		Path:         "/admin/login",
		Handler:      handlers.ShowAdminLogin,
		Method:       http.MethodGet,
		AuthRequired: false,
	},
	{
		Path:         "/admin/login",
		Handler:      handlers.PostLogin,
		Method:       http.MethodPost,
		AuthRequired: false,
	},
	{
		Path:         "/admin/logout",
		Handler:      handlers.Logout,
		Method:       http.MethodGet,
		AuthRequired: true,
	},
	{
		Path:         "/admin/index",
		Handler:      handlers.ShowAdminIndex,
		Method:       http.MethodGet,
		AuthRequired: true,
	},
	{
		Path:         "/admin/users",
		Handler:      handlers.ShowAdminUsers,
		Method:       http.MethodGet,
		AuthRequired: true,
	},
	{
		Path:         "/admin/users/{id:[0-9]+}",
		Handler:      handlers.UpdateUser,
		Method:       http.MethodPost,
		AuthRequired: true,
	},
	{
		Path:         "/admin/users/{id:[0-9]+}",
		Handler:      handlers.DeleteUser,
		Method:       http.MethodGet,
		AuthRequired: true,
	},
	{
		Path:         "/admin/posts",
		Handler:      handlers.ShowAdminPosts,
		Method:       http.MethodGet,
		AuthRequired: true,
	},
	{
		Path:         "/admin/attrs",
		Handler:      handlers.ShowAdminAttrs,
		Method:       http.MethodGet,
		AuthRequired: true,
	},
}

var blogRoutes = []Router{
	{
		Path:         "/",
		Handler:      handlers.ShowBlogFront,
		Method:       http.MethodGet,
		AuthRequired: true,
	},
}

// 用户路由
var userRoutes = []Router{}

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
