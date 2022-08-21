package handlers

import (
	"github.com/gorilla/sessions"
	"github.com/jmoiron/sqlx"
	"lightsaid.com/weblogs/cmd/web/config"
	"lightsaid.com/weblogs/internal/render"
	"lightsaid.com/weblogs/internal/service"
)

var App *config.AppConfig
var H *AppHandler

// AppHandler 存储handlers包需要的数据
type AppHandler struct {
	DB          *sqlx.DB
	Repo        *service.Service
	Template    *render.TemplateData
	CookieStore *sessions.CookieStore
}

// New 创建一个 AppHandler 实例, 给整个handlers包使用, 同时在 handlers 包引用 App
func New(db *sqlx.DB, cfg *config.AppConfig) *AppHandler {
	App = cfg

	H = &AppHandler{
		DB:          db,
		Repo:        service.New(db),
		Template:    render.New(cfg.UseCache, cfg.CookieStore),
		CookieStore: cfg.CookieStore,
	}

	return H
}
