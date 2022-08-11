package main

import (
	"log"
	"net/http"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/zap"
	"lightsaid.com/weblogs/cmd/web/routes"
	"lightsaid.com/weblogs/internal/models"
	"lightsaid.com/weblogs/pkg/logger"
)


func main(){
	var err error
	err = godotenv.Load(".env")
	if err != nil {
		log.Panic(err)
	}

	l, err := logger.New("./logs.log", "stderr")
	if err != nil {
		log.Panic(err)
	}
	defer l.Sync()

	// sqlx.Connect() 带有 ping 功能. 
	db, err := sqlx.Connect("sqlite3", "file:./resources/database/weblogs.db")
	// db, err := sqlx.Open("sqlite3", "./resources/database/weblogs.db")
	if err != nil {
		log.Panic(err)
		return
	}

	// Test
	var user models.User
	query := "SELECT id, username, email, active, created_at, updated_at from users limit 1;"
	err = db.Get(&user, query)
	if err != nil {
		log.Println(err)
		return
	}
	// log.Println(user)
	
	zap.S().Error(user)

	r := routes.New()

	srv := &http.Server{
		Addr: os.Getenv("WEBPORT"),
		Handler: r,
	}

	log.Println("Starting server on port ", os.Getenv("WEBPORT"))
	if err := srv.ListenAndServe(); err != nil {
		log.Panic(err)
	}
}
