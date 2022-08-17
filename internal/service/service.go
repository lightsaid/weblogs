package service

import (
	"github.com/jmoiron/sqlx"
	"lightsaid.com/weblogs/internal/repository"
	"lightsaid.com/weblogs/internal/repository/dbrepo"
)

// Service 连接handler和repository的中间层
type Service struct {
	repository.Repository
}

// New 创建一个 Service 实例，使用service层调用db层
func New(db *sqlx.DB) *Service {
	return &Service{
		dbrepo.NewDatabaseRepo(db),
	}
}
