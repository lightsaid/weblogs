package app

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/zap"

	"lightsaid.com/weblogs/cmd/web/config"
	"lightsaid.com/weblogs/cmd/web/handlers"
	"lightsaid.com/weblogs/cmd/web/routes"
	"lightsaid.com/weblogs/pkg/logger"
)

var app config.AppConfig

func initial() {
	setupEnv()
	db := setupDB()
	app = config.AppConfig{
		DB:       db,
		Looger:   setupLogger(),
		UseCache: false,
	}
	app.CookieStore = setupSession()

	mode := os.Getenv("RUNMODE")
	if mode == "prod" {
		app.UseCache = true
	}

}

func handleError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func setupEnv() {
	err := godotenv.Load(".env")
	handleError(err)
}

func setupDB() *sqlx.DB {
	db, err := sqlx.Connect("sqlite3", "file:./resources/database/weblogs.db")
	handleError(err)
	return db
}

func setupLogger() *zap.Logger {
	l, err := logger.New("./logs.log", "stderr")
	handleError(err)
	return l
}

func setupSession() *sessions.CookieStore {
	store := sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY)")))
	// NOTE: CSRF 的本质实际上是利用了 Cookie 会自动在请求中携带的特性，诱使用户在第三方站点发起请求的行为。
	// 因此 Cookie 增加了 SameSite 属性，用来规避该问题。
	// SameSite=None：无论是否跨站都会发送 Cookie
	// SameSite=Lax：允许部分第三方请求携带 Cookie
	// SameSite=Strict：仅允许同站请求携带 Cookie，即当前网页 URL 与请求目标 URL 完全一致
	var secure bool
	if os.Getenv("RUNMODE") == "prod" {
		secure = true
	}
	store.Options.SameSite = http.SameSiteLaxMode
	// Secure cookie 仅通过 HTTPS 协议加密发送到服务器。请注意，不安全站点（http:）无法使用 Secure 指令设置 cookies。
	store.Options.Secure = secure
	store.Options.MaxAge = 3 * 60 * 60 // 单位秒
	// store.Options.MaxAge = 10 // 10 秒测试
	return store
}

func Serve() {
	initial()

	// 刷新日志缓存
	defer app.Looger.Sync()

	// defer app.DB.Close()

	// 先实例化 handlers.AppHandler 再创建路由
	handlers.New(app.DB, &app)

	// 创建路由
	r := routes.New()

	// 配置 http server
	srv := &http.Server{
		Addr:    os.Getenv("WEBPORT"),
		Handler: r,
	}

	// 启动server
	log.Println("Starting server on port ", os.Getenv("WEBPORT"))
	if err := srv.ListenAndServe(); err != nil {
		log.Panic(err)
	}
}
