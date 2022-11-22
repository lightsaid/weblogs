package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	mux := mux.NewRouter()

	// 中间件， 此写法过于粗暴
	// mux.Use(app.RecoverPanic)
	// mux.Use(app.Logging)
	// mux.Use(app.NoSurf)
	// mux.Use(app.SecureHeader)
	// mux.Use(app.LoadSesion)

	// 使用 alice 链式(组合)调用中间件
	// 基础中间件
	baseMw := alice.New(app.RecoverPanic, app.Logging, app.SecureHeader)
	// 动态中间件
	dynamicMw := alice.New(app.LoadSesion, app.NoSurf, app.Authenticate)
	// 组合认证中间件
	requireAuthMw := dynamicMw.Append(app.RequireAuth)

	mux.Handle("/", dynamicMw.ThenFunc(app.Controller.IndexPage))
	mux.Handle("/full", dynamicMw.ThenFunc(app.Controller.EditorMDFullExample))

	// 注册功能暂不对外开放
	if app.Config.Mode == "dev" {
		// Get: 提供CSRFToken, Post: 处理注册逻辑
		mux.Handle("/user/register", dynamicMw.ThenFunc(app.Controller.JSONRegister)).Methods(http.MethodPost, http.MethodGet)
	}

	// users 相关
	mux.Handle("/user/login", dynamicMw.ThenFunc(app.Controller.LoginPage)).Methods(http.MethodGet)
	mux.Handle("/user/login", dynamicMw.ThenFunc(app.Controller.PostLogin)).Methods(http.MethodPost)
	mux.Handle("/user/logout", dynamicMw.ThenFunc(app.Controller.Logout)).Methods(http.MethodGet)
	mux.Handle("/user/update/{id:[0-9]+}", requireAuthMw.ThenFunc(app.Controller.UpdateUser)).Methods(http.MethodPost)
	mux.Handle("/user/settings", requireAuthMw.ThenFunc(app.Controller.SettingsPage)).Methods(http.MethodGet)
	mux.Handle("/user/about", dynamicMw.ThenFunc(app.Controller.AboutPage)).Methods(http.MethodGet)

	// Tags 相关 // TODO:
	mux.Handle("/tag", dynamicMw.ThenFunc(app.Controller.TagPage)).Methods(http.MethodGet)
	mux.Handle("/tag/create", requireAuthMw.ThenFunc(app.Controller.CreateTag)).Methods(http.MethodPost)
	mux.Handle("/tag/delete/{id:[0-9]+}", requireAuthMw.ThenFunc(nil)).Methods(http.MethodGet)
	mux.Handle("/tag/update/{id:[0-9]+}", requireAuthMw.ThenFunc(nil)).Methods(http.MethodPost)

	// Posts 相关 // TODO:
	postRouter := mux.PathPrefix("/post").Subrouter()
	// JSON 接口
	postRouter.Handle("/create", requireAuthMw.ThenFunc(app.Controller.CreatePostPage)).Methods(http.MethodGet)
	postRouter.Handle("/create", requireAuthMw.ThenFunc(app.Controller.CreatePost)).Methods(http.MethodPost)
	postRouter.Handle("/detail/{id:[0-9]+}", dynamicMw.ThenFunc(app.Controller.PostDetail)).Methods(http.MethodGet)
	postRouter.Handle("/update/{id:[0-9]+}", requireAuthMw.ThenFunc(nil)).Methods(http.MethodPost)
	postRouter.Handle("/delete/{id:[0-9]+}", requireAuthMw.ThenFunc(nil)).Methods(http.MethodPost)

	// notifications 通知，比如有人评论文章就给用户发送一个通知 // TODO: 后续开发
	mux.Handle("/notification", requireAuthMw.ThenFunc(app.Controller.NotificationPage))

	// 用于检查应用是否正常运行
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	// 静态资源
	mux.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// 使用基础中间件
	return baseMw.Then(mux)
}
