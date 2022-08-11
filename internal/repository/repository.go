package repository

import "lightsaid.com/weblogs/internal/models"

// Repository 定义 Database 操作，需要dbrepo包实现
type Repository interface {
	// User 模块
	InsertUser(email, username, password, avatar string) (models.User, error)
	GetUser(id int) (models.User, error)
	GetUsers() ([]models.User, error)
	UpdateUser(models.User) error
	DeleteUser(id int) error

	// Post 模块
	InsertPost(Post models.Post) (int, error)
	GetPost(id int) (models.Post, error)
	GetPosts() ([]models.Post, error)
	UpdatePost(models.Post) error
	DeletePost(id int) error
}
