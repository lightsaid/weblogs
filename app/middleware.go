package app

import (
	"net/http"
	"runtime/debug"
	"time"

	"github.com/ory/nosurf"
	log "github.com/sirupsen/logrus"
)

func (app *application) RecoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				// 记录一下调用堆栈信息
				log.Panicf("panic recover err: %s \n %s", err, debug.Stack())
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
		log.Infof("%s %s %s", r.Method, r.URL, time.Since(start))
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
