package middleware

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"sync"

	"github.com/gorilla/sessions"
	"go.uber.org/ratelimit"
	"go.uber.org/zap"
	"lightsaid.com/weblogs/cmd/web/handlers"
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
		// 存到 context 更方便取
		ctx := context.WithValue(r.Context(), "userinfo", userinfo.(service.SessionUser))
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

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
