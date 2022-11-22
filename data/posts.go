package data

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Post struct {
	ID        int     `db:"id" json:"id"`
	UserId    int     `db:"user_id" json:"user_id"`
	Author    string  `db:"author" json:"author"`
	Title     string  `db:"title" json:"title"`
	Content   string  `db:"content" json:"content"`
	Thumb     *string `db:"thumb" json:"thumb"`
	Readings  int     `db:"readings" json:"readings"`
	Comments  int     `db:"comments" json:"comments"`
	Likes     int     `db:"likes" json:"likes"`
	Active    int     `db:"active" json:"active"`
	CreatedAt string  `db:"created_at" json:"created_at"`
	UpdatedAt string  `db:"updated_at" json:"updated_at"`
}

type PostModel struct {
	DB *sqlx.DB
}

func (repo *PostModel) Insert(post Post) error {
	query := `insert into posts(user_id, author, title, content, thumb) values($1, $2, $3, $4, $5)`

	result, err := repo.DB.Exec(query, post.UserId, post.Author, post.Title, post.Content, post.Thumb)
	if err != nil {
		return fmt.Errorf("insert post error: %w", err)
	}
	if id, err := result.LastInsertId(); err != nil || id <= 0 {
		if err != nil {
			return fmt.Errorf("insert post with lastInsertId error: %w", err)
		}
		return errors.New("insert post lastInsertId <= 0")
	}
	return nil
}

func (repo *PostModel) GetById(id int) (*Post, error) {
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
	where id = $1 and active != -1;`

	var p Post
	err := repo.DB.Get(&p, query, id)
	return &p, err
}
func (repo *PostModel) GetList(pageSize, pageIndex int) ([]*Post, error) {
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
	where active != -1
	order by created_at desc 
	limit $1 offset $2; `

	var posts []*Post
	err := repo.DB.Select(&posts, query, pageSize, pageIndex)
	if err != nil {
		return nil, fmt.Errorf("PostModel GetList error: %w", err)
	}
	return posts, nil
}
