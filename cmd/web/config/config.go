package config

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type AppConfig struct {
	DB       *sqlx.DB
	Looger   *zap.Logger
	UseCache bool
}
