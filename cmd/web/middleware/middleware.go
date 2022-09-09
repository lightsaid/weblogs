package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/gorilla/sessions"
	"go.uber.org/ratelimit"
	"go.uber.org/zap"
	"lightsaid.com/weblogs/cmd/web/handlers"
	"lightsaid.com/weblogs/internal/models"
	"lightsaid.com/weblogs/internal/service"
)

// TODO:

type Middleware func(http.HandlerFunc) http.HandlerFunc

// MultipleMiddleware 解决多层中间件嵌套书写格式问题, 让多个中间件书写更加优雅
func MultipleMiddleware(h http.HandlerFunc, m ...Middleware) http.HandlerFunc {
	if len(m) < 1 {
		return h
	}
	wrapped := h
	// 多个中间件,
	for i := len(m) - 1; i >= 0; i-- {
		wrapped = m[i](wrapped)
	}
	return wrapped
}

func LogMiddlewate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		zap.S().Info(r.Method, r.URL)

		next.ServeHTTP(w, r)
	})
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var session *sessions.Session
		var err error
		if session, err = handlers.H.CookieStore.Get(r, os.Getenv("SESSION")); err != nil {
			handlers.ServerError(w, errors.New("get session error"))
			return
		}
		userinfo, exists := session.Values["userinfo"]
		if !exists {
			// 没有登录
			session.AddFlash("请先登录", "Warning")
			http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
			return
		}

		var user models.User
		info, ok := userinfo.(service.SessionUser)
		if ok && info.UserID > 0 {
			user, err = handlers.H.Repo.GetUser(info.UserID)
			if err != nil {
				zap.S().Error(err)
				session.Options.MaxAge = -1
				session.AddFlash("获取用户信息服务错误", "Error")
				session.Save(r, w)
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}
			ub, _ := json.Marshal(user)
			fmt.Println("当前访问后台用户：", string(ub))
		}

		// 如果查询到用户，将用户信息写入上下文
		if user.ID > 0 {
			sessionUser := service.SessionUser{
				UserID:   user.ID,
				Username: user.Username,
				Avatar:   *user.Avatar,
			}
			ctx := context.WithValue(r.Context(), "userinfo", sessionUser)
			r = r.WithContext(ctx)
		} else {
			// 清理用户cookie
			session.Options.MaxAge = -1
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		// NOTE: 访问的是否是管理员资源，是否登陆? 是否管理员？
		// 访问管理员理由
		if strings.Contains(r.URL.Path, "/admin/") {
			if user.ID > 0 && user.IfAdmin == 1 {
				next.ServeHTTP(w, r)
				return
			}
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		// 普通用户
		if user.ID > 0 {
			next.ServeHTTP(w, r)
		} else {
			// 没有登录
			session.AddFlash("您还没登陆！", "Warning")
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	})
}

func SettingUserinfo(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var session *sessions.Session
		var err error
		if session, err = handlers.H.CookieStore.Get(r, os.Getenv("SESSION")); err != nil {
			handlers.ServerError(w, errors.New("get session error"))
			return
		}
		userinfo, exists := session.Values["userinfo"]
		if exists {
			// 存到 context 更方便取
			ctx := context.WithValue(r.Context(), "userinfo", userinfo.(service.SessionUser))
			r = r.WithContext(ctx)

			fmt.Println("SettingUserinfo >>> ", userinfo.(service.SessionUser))
		}
		next.ServeHTTP(w, r)
	})
}

// EditorRedirect editor.md编辑器内部资源引用重定向
// func EditorRedirect(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		var path = r.URL.Path
// 		if strings.HasPrefix(path, "/lib/") || strings.HasPrefix(path, "/plugins/") {
// 			url := fmt.Sprintf("%s%s", "/static/editor.md", path)
// 			fmt.Println("editor.md >>> ", url)
// 			http.Redirect(w, r, url, http.StatusMovedPermanently)
// 		} else {
// 			next.ServeHTTP(w, r)
// 		}
// 	})
// }

func RateLimit(rate int) func(next http.Handler) http.Handler {
	var lmap sync.Map

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			host, _, err := net.SplitHostPort(r.RemoteAddr)
			if err != nil {
				http.Error(w, fmt.Sprintf("invalid RemoteAddr: %s", err), http.StatusInternalServerError)

				return
			}

			lif, ok := lmap.Load(host)
			if !ok {
				lif = ratelimit.New(rate)
			}

			lm, ok := lif.(ratelimit.Limiter)
			if !ok {
				http.Error(w, "internal middleware error: typecast failed", http.StatusInternalServerError)
				return
			}

			lm.Take()
			lmap.Store(host, lm)

			next.ServeHTTP(w, r)
		})
	}
}
