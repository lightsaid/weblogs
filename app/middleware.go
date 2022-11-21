package app

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"runtime/debug"
	"strings"
	"time"

	"github.com/ory/nosurf"
	log "github.com/sirupsen/logrus"
	"lightsaid.com/weblogs/global"
)

// RecoverPanic recover
func (app *application) RecoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				// 记录调用堆栈信息
				log.Panicf("panic recover err: %v \n %s", err, debug.Stack())
				w.Header().Set("Connection", "close")
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}

// Logging 日志
func (app *application) Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		if !strings.Contains(r.URL.RequestURI(), "/static/") {
			log.Infof("%s %s %s", r.Method, r.URL, time.Since(start))
		}
	})
}

// LoadSesion session 管理
func (app *application) LoadSesion(next http.Handler) http.Handler {
	return app.Controller.Session.LoadAndSave(next)
}

// NoSurf 防止 CSRF 攻击
func (app *application) NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.ExemptPaths(app.excludeURIs...)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.IsProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// SecureHeader 设置安全的请求头，防止 XSS 攻击
func (app *application) SecureHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("X-Frame-Options", "deny")

		next.ServeHTTP(w, r)
	})
}

// NOTE: 关于认证的中间件有两个，请求进来先经过 Authenticate 查询数据库鉴权，权限没问题，则设置上下文，
//       再交由 RequireAuth 获取上下文，鉴别是否有权限访问资源

// Authenticate 进行认证
func (app *application) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 判断 Session 中是否包含 "user_id" 字段，登录过后的用户必须是存在的
		exists := app.Controller.Session.Exists(r.Context(), "user_id")
		if !exists {
			// 不存在，放行，交由 RequireAuth 判刑是否需要登录才能访问
			next.ServeHTTP(w, r)
		}
		// 如果存在，则说明用户应该已经登录，为了更安全的，查询数据库，该用户是否还有资源访问的合法性
		// （防止：登录后，用户被删了，权限更改了，Session 没有及时更新，还能访问）
		userID := app.Controller.Session.GetInt(r.Context(), "user_id")
		user, err := app.Controller.Models.Users.GetById(userID)
		if errors.Is(errors.Unwrap(err), sql.ErrNoRows) {
			app.Controller.Session.Remove(r.Context(), "user_id")
			next.ServeHTTP(w, r)
			return
		} else if err != nil {
			app.Controller.ServerError(w, err)
		}
		// 如果用户被删除
		if user.Active == -1 {
			app.Controller.Session.Remove(r.Context(), "user_id")
			next.ServeHTTP(w, r)
			return
		}

		// 以上是通过查询数据库，鉴别是否有足够的权限访问

		// 当用户有足够的权限，设置用户已经登录的上下文，交由 RequireAuth 获取使用
		ctx := context.WithValue(r.Context(), global.KeyIsAuthenticated, true)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// requireAuth 身份认证, 只有登录有才有权限访问与操作资源
func (app *application) RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 没有登录
		if !app.Controller.IsAuthenticated(r) {
			// 重定向到登录页
			http.Redirect(w, r, "/user/login", http.StatusSeeOther)
			return
		}
		// 不缓存请求
		w.Header().Add("Cache-Control", "no-store")
		next.ServeHTTP(w, r)
	})
}
