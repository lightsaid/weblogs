package config

import (
	"github.com/gorilla/sessions"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type AppConfig struct {
	DB          *sqlx.DB
	Looger      *zap.Logger
	UseCache    bool
	CookieStore *sessions.CookieStore
}
