package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"lightsaid.com/weblogs/configs"
	"lightsaid.com/weblogs/controller"
	"lightsaid.com/weblogs/logger"
)

type application struct {
	Config       *configs.Config
	Controller   *controller.Controller
	IsProduction bool     // 是否生产环境
	excludeURIs  []string // 不检查 CSRF 的路由
}

func New(db *sqlx.DB, conf *configs.Config) *application {
	// 初始化日志输出
	logger.InitLogger()
	control := controller.NewController(db)

	var isProd bool
	if conf.Mode == "prod" {
		isProd = true
	}

	app := application{
		Config:       conf,
		Controller:   control,
		IsProduction: isProd,
		// “/register” 是JSON请求，不需要 NoSurf 中间处理，交由对应的controller处理
		excludeURIs: []string{"/register", "/post/create"},
	}

	return &app
}

func (app *application) Serve() error {
	var address = fmt.Sprintf("0.0.0.0:%d", app.Config.Port)
	srv := &http.Server{
		Addr:         address,
		Handler:      app.routes(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 20 * time.Second,
		IdleTimeout:  time.Minute,
	}

	shutdownError := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		s := <-quit

		logrus.Println("recve shutting down server signal: ", s.String())

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		err := srv.Shutdown(ctx)
		if err != nil {
			shutdownError <- err
		}

		shutdownError <- nil
	}()

	logrus.Info("app start on: ", address)
	err := srv.ListenAndServe()
	if err != nil {
		return err
	}

	err = <-shutdownError
	if err != nil {
		return err
	}

	log.Println("stopped server")

	return nil
}
