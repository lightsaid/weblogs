package middleware

import (
	"context"
	"net/http"

	"go.uber.org/zap"
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

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "auth", "abc123")
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
