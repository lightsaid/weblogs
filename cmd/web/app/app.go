package app

import (
	"log"
	"net/http"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/zap"

	"lightsaid.com/weblogs/cmd/web/handlers"
	"lightsaid.com/weblogs/cmd/web/routes"
	"lightsaid.com/weblogs/pkg/logger"
)

type AppConfig struct {
	DB     *sqlx.DB
	AppH   *handlers.AppHandler
	Looger *zap.Logger
}

func New() *AppConfig {
	setupEnv()
	return &AppConfig{
		DB:     setupDB(),
		Looger: setupLogger(),
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
	var err error
	err = godotenv.Load(".env")
	handleError(err)

	db, err := sqlx.Connect("sqlite3", "file:./resources/database/weblogs.db")
	handleError(err)
	return db
}

func setupLogger() *zap.Logger {
	l, err := logger.New("./logs.log", "stderr")
	handleError(err)
	return l
}

func (app *AppConfig) Serve() {
	// 刷新日志缓存
	defer app.Looger.Sync()

	// 先实例化 handlers.AppHandler 再创建路由
	app.AppH = handlers.New(app.DB)

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
