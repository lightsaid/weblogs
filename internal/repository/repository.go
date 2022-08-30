package repository

import "lightsaid.com/weblogs/internal/models"

// Repository 定义 Database 操作，需要dbrepo包实现
type Repository interface {
	// User 模块
	InsertUser(email, username, password, avatar string) (models.User, error)
	GetUser(id int) (models.User, error)
	GetUsers() ([]models.User, error)
	GetUserByEmial(email string) (models.User, error)
	UpdateUser(models.User) error
	DeleteUser(id int) error

	// Attributes 模块
	InsertAttrs(attr *models.Attribute) (*models.Attribute, error)
	GetAttributes() ([]*models.Attribute, error)
	UpdateAttributes(a *models.Attribute) error
	DeleteAttribute(id int) error

	// Categories 模块
	InsertCategories(cate *models.Category) (*models.Category, error)
	GetCategories(parent_id int) ([]*models.Category, error)
	UpdateCategories(a *models.Category) error
	DeleteCategories(id int) error
	GetCategoriesById(id int) (*models.Category, error)

	// Post 模块
	InsertPost(post models.Post) (*models.Post, error)
	GetPost(id int) (*models.Post, error)
	GetPosts(pageSize, pageInt int) ([]*models.Post, error)
	UpdatePost(models.Post) error
	DeletePost(id int) error
}
