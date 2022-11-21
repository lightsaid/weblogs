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

// Controller 控制器，处理请求
type Controller struct {
	DB      *sqlx.DB
	Models  data.Models
	Session *scs.SessionManager
	*HTMLTemplate
	*Toolkit
}

// initSession 初始化 Session
func initSession() *scs.SessionManager {
	// 选择使用 alexedwards/scs/v2 而不是  Gorilla Sessions 的原因
	// SCS 操作起来更方便，有比 Gorilla Sessions 更良好的接口
	// SCS RenewToken() 可以更新会话，减少 session fixation attacks （https://owasp.org/www-community/attacks/Session_fixation）

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

// NewController 实例化一个控制器，提供给routes使用
func NewController(db *sqlx.DB) *Controller {
	control := &Controller{
		DB:      db,
		Models:  data.NewModels(db),
		Session: initSession(),
	}
	var err error
	control.HTMLTemplate, err = NewHTMLTemplate(control.Session)
	if err != nil {
		log.Fatal(err)
	}
	return control
}

// IsAuthenticated 检查是否已登录(是否已认证，认证的说法更合理，认证包含了登录，还有用户有足够的权限)
func (_this *Controller) IsAuthenticated(r *http.Request) bool {
	isAuth, ok := r.Context().Value(global.KeyIsAuthenticated).(bool)
	if !ok {
		return false
	}
	return isAuth
}
