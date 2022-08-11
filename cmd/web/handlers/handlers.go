package handlers

import (
	"github.com/jmoiron/sqlx"
)

var AppH *AppHandler

// AppHandler 存储handlers包需要的数据
type AppHandler struct {
	DB *sqlx.DB
}

// New 创建一个 AppHandler 实例, 给整个handlers包使用
func New(db *sqlx.DB) *AppHandler {
	AppH = &AppHandler{
		DB: db,
	}
	return AppH
}
