package main

import (
	_ "github.com/mattn/go-sqlite3"
	"lightsaid.com/weblogs/cmd/web/app"
)

func main() {
	app := app.New()
	app.Serve()
}
