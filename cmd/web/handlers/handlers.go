package handlers

import (
	"github.com/jmoiron/sqlx"
)

var AppH *AppHandler

// AppHandler 提供一个请求入口结构体,并处理所有请求
type AppHandler struct {
	DB *sqlx.DB
}

// New 创建一个 AppHandler 实例, 并且赋值给 AppH 提供给 routes 使用
func New(db *sqlx.DB) *AppHandler{
	AppH = &AppHandler{
		DB: db,
	}
	return AppH
}