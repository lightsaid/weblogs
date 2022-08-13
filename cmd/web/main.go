package main

import (
	_ "github.com/mattn/go-sqlite3"
	"lightsaid.com/weblogs/cmd/web/app"
)

func main() {
	// db, err := sqlx.Connect("sqlite3", "file:./resources/database/weblogs.db")
	// log.Println("amin.go>> ", db, err)
	app.Serve()
}
