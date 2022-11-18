package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (app *application) routes() http.Handler {
	mux := mux.NewRouter()

	// 中间件
	mux.Use(app.RecoverPanic)
	mux.Use(app.NoSurf)
	mux.Use(app.LoadSesion)
	mux.Use(app.Logging)

	mux.HandleFunc("/", app.Controller.IndexPage)
	mux.HandleFunc("/full", app.Controller.EditorMDFullExample)

	// 注册功能暂不对外开放
	if app.Config.Mode == "dev" {
		// Get: 提供CSRFToken, Post: 处理注册逻辑
		mux.HandleFunc("/register", app.Controller.JSONRegister).Methods(http.MethodPost, http.MethodGet)
	}

	mux.HandleFunc("/login", app.Controller.LoginPage).Methods(http.MethodGet)
	mux.HandleFunc("/login", app.Controller.PostLogin).Methods(http.MethodPost)

	// 静态资源
	mux.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	return mux
}
