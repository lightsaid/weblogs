package main

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"lightsaid.com/weblogs/app"
	"lightsaid.com/weblogs/configs"
	"lightsaid.com/weblogs/global"
)

// initConfig 初始化并设置全局 config
func initConfig() *configs.Config {
	conf, err := configs.ReadConfig("./config.json")
	if err != nil {
		log.Fatal(err)
	}
	global.Config = conf
	return conf
}

// 初始化 database 连接并检查链接是否正常
func initSQLite(source string) *sqlx.DB {
	db, err := sqlx.Open("sqlite3", source)
	if err != nil {
		log.Fatal("open SQLite error: ", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("ping error: ", err)
	}
	return db
}

// main 运行程序入口
func main() {
	conf := initConfig()
	log.Printf("%+v", conf)

	db := initSQLite(conf.DBSource)
	defer db.Close()

	app := app.New(db, conf)

	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
