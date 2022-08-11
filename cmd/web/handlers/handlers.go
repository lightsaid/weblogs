package handlers

import (
	"github.com/jmoiron/sqlx"
	"lightsaid.com/weblogs/internal/service"
)

var AppH *AppHandler

// AppHandler 存储handlers包需要的数据
type AppHandler struct {
	DB   *sqlx.DB
	Repo *service.Service
}

// New 创建一个 AppHandler 实例, 给整个handlers包使用
func New(db *sqlx.DB) *AppHandler {
	AppH = &AppHandler{
		DB:   db,
		Repo: service.New(db),
	}
	return AppH
}
