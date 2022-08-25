package main

import (
	"encoding/gob"

	_ "github.com/mattn/go-sqlite3"
	"lightsaid.com/weblogs/cmd/web/app"
	"lightsaid.com/weblogs/internal/service"
)

func main() {
	// db, err := sqlx.Connect("sqlite3", "file:./resources/database/weblogs.db")
	// log.Println("amin.go>> ", db, err)

	gob.Register(service.SessionUser{})

	app.Serve()
}
