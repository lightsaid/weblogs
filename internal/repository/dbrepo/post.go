package dbrepo

import (
	"lightsaid.com/weblogs/internal/models"
)

func (repo *databaseRepo) InsertPost(post models.Post) (*models.Post, error) {
	query := `insert into posts(user_id, author, title, content, thumb) values($1, $2, $3, $4, $5) returning *`
	var p models.Post

	err := repo.DB.Get(&p, query, post.UserId, post.Author, post.Title, post.Content, post.Thumb)
	return &p, err
}

func (repo *databaseRepo) GetPost(id int) (*models.Post, error) {

	query := `select 
		id, 
		user_id, 
		author, 
		title, 
		content, 
		thumb, 
		readings, 
		comments, 
		likes, 
		active, 
		created_at, 
		updated_at 
		from posts 
		where id = $1; `

	var p models.Post
	err := repo.DB.Get(&p, query, id)

	return &p, err
}
func (repo *databaseRepo) GetPosts(pageSize, pageIndex int) ([]*models.Post, error) {
	query := `select 
		id, 
		user_id, 
		author, 
		title, 
		content, 
		thumb, 
		readings, 
		comments, 
		likes, 
		active, 
		created_at, 
		updated_at 
		from posts 
		limit $1 offset $2;`

	var posts []*models.Post

	err := repo.DB.Select(&posts, query, pageSize, pageIndex)

	return posts, err
}
func (repo *databaseRepo) UpdatePost(models.Post) error {
	return nil
}
func (repo *databaseRepo) DeletePost(id int) error {
	return nil
}
