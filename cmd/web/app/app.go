package app

import (
	"log"
	"net/http"
	"os"

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
	log.Println("setupDB>> ", db, err)
	return db
}

func setupLogger() *zap.Logger {
	l, err := logger.New("./logs.log", "stderr")
	handleError(err)
	return l
}

func Serve() {
	initial()

	// 刷新日志缓存
	defer app.Looger.Sync()

	// defer app.DB.Close()

	zap.S().Info("Serve>> ", app.DB)
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
