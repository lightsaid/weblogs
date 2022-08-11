package dbrepo

import (
	"github.com/jmoiron/sqlx"
	"lightsaid.com/weblogs/internal/repository"
)

// databaseRepo 数据操作结构体
type databaseRepo struct {
	DB *sqlx.DB
}

// NewDatabaseRepo
func NewDatabaseRepo(db *sqlx.DB) repository.Repository {
	return &databaseRepo{
		DB: db,
	}
}
