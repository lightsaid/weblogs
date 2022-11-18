package data

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Tag struct {
	ID     int    `db:"id" json:"id"`
	Name   string `db:"name" json:"name"`
	UserID int    `db:"user_id" json:"user_id"`
}

type TagModel struct {
	DB *sqlx.DB
}

// Insert 新增tag
func (m *TagModel) Insert(userId int, name string) error {
	query := `insert into tags(user_id, name) values($1, $2)`

	result, err := m.DB.Exec(query, userId, name)
	if err != nil {
		return fmt.Errorf("insert tag error: %w", err)
	}
	if id, err := result.LastInsertId(); err != nil || id <= 0 {
		if err != nil {
			return fmt.Errorf("insert tag with lastInsertId error: %w", err)
		}
		return errors.New("insert tag lastInsertId <= 0")
	}
	return nil
}

// GetPostsByTagId 根据 TagId 获取 post 列表
func (m *TagModel) GetPostsByTagId(tagId int, pageSize, pageIndex int) ([]*Post, error) {
	query := `select * from posts p join post_tags pt on p.id = pt.post_id where pt.tag_id = $1 limit $2 offset $3;`

	var posts []*Post
	err := m.DB.Select(&posts, query, tagId, pageSize, pageIndex)
	return posts, err
}
