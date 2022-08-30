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
	{
		Path:         "/admin/attrs",
		Handler:      handlers.CreateAttribute,
		Method:       http.MethodPost,
		AuthRequired: true,
	},
	{
		Path:         "/admin/attrs/{id:[0-9]+}",
		Handler:      handlers.UpdateAttribute,
		Method:       http.MethodPost,
		AuthRequired: true,
	},
	{
		Path:         "/admin/attrs/{id:[0-9]+}",
		Handler:      handlers.DeleteAttribute,
		Method:       http.MethodGet,
		AuthRequired: true,
	},
	{
		Path:         "/admin/category",
		Handler:      handlers.CreateCategories,
		Method:       http.MethodPost,
		AuthRequired: true,
	},
	{
		Path:         "/admin/category/{id:[0-9]+}",
		Handler:      handlers.UpdateCategories,
		Method:       http.MethodPost,
		AuthRequired: true,
	},
	{
		Path:         "/admin/category/{parent_id:[0-9]+}/{id:[0-9]+}",
		Handler:      handlers.DeleteCategories,
		Method:       http.MethodGet,
		AuthRequired: true,
	},
	{
		Path:         "/admin/categories/{parent_id:[0-9]+}",
		Handler:      handlers.ShowAdminCategories,
		Method:       http.MethodGet,
		AuthRequired: true,
	},
	{
		Path:         "/admin/posts",
		Handler:      handlers.CreatePost,
		Method:       http.MethodPost,
		AuthRequired: true,
	},
	{
		Path:         "/admin/json_post",
		Handler:      handlers.CreateJsonPost,
		Method:       http.MethodPost,
		AuthRequired: true,
	},
	{
		Path:         "/admin/posts",
		Handler:      handlers.GetPosts,
		Method:       http.MethodGet,
		AuthRequired: true,
	},

	{
		Path:         "/admin/demo",
		Handler:      handlers.ShowDemo,
		Method:       http.MethodGet,
		AuthRequired: true,
	},
	{
		Path:         "/admin/publish",
		Handler:      handlers.ShowPublishPost,
		Method:       http.MethodGet,
		AuthRequired: true,
	},
}

var blogRoutes = []Router{
	{
		Path:         "/",
		Handler:      handlers.ShowDesktop,
		Method:       http.MethodGet,
		AuthRequired: false,
	},
	{
		Path:         "/column",
		Handler:      handlers.ShowColumn,
		Method:       http.MethodGet,
		AuthRequired: false,
	},
	{
		Path:         "/about",
		Handler:      handlers.ShowAbout,
		Method:       http.MethodGet,
		AuthRequired: false,
	},
}

// 用户路由
var userRoutes = []Router{}

// 文章路由
var postRoutes = []Router{}

// 分类路由
var cateRoutes = []Router{}

// 属性路由
var attrRoutes = []Router{}

// 评论路由
