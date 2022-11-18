package controller

import (
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"lightsaid.com/weblogs/data"
	"lightsaid.com/weblogs/global"
)

const (
	DevModeValue       = "dev"
	ProdModeValue      = "prod"
	PagePathPattern    = "./templates/*.page.gtpl"
	LayoutPathPattern  = "./templates/*.layout.gtpl"
	PartialPathPattern = "./templates/*.partial.gtpl"
)

type Controller struct {
	DB      *sqlx.DB
	Models  data.Models
	Session *scs.SessionManager
	*HTMLTemplate
	*Toolkit
}

func NewController(db *sqlx.DB) *Controller {
	control := &Controller{
		DB:      db,
		Models:  data.NewModels(db),
		Session: initSession(),
	}
	var err error
	control.HTMLTemplate, err = NewHTMLTemplate()
	if err != nil {
		log.Fatal(err)
	}
	return control
}

func initSession() *scs.SessionManager {
	// session设置
	// 参考：
	//   https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Cookies
	//	 https://github.com/alexedwards/scs
	sessionManager := scs.New()
	sessionManager.Lifetime = 3 * time.Hour
	// sessionManager.IdleTimeout = 20 * time.Minute
	// sessionManager.Cookie.Name = "session_id"
	// sessionManager.Cookie.Domain = "example"
	sessionManager.Cookie.HttpOnly = true
	sessionManager.Cookie.Path = "/"
	sessionManager.Cookie.Persist = true // 是否持久
	// SameSite 属性允许服务器指定是否/何时通过跨站点请求发送
	// 它采用三个可能的值：Strict、Lax 和 None。
	// SameSite=None 的 cookie 还必须指定 Secure 属性（它们需要安全上下文）。
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode // 使用默认
	sessionManager.Cookie.Secure = false

	if global.Config.Mode == ProdModeValue {
		sessionManager.Cookie.Secure = true
		sessionManager.Cookie.Domain = global.Config.Domain
	}

	return sessionManager
}
