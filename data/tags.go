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

// 统计 tag 列表
type TagList struct {
	TagID    int    `db:"tag_id"`
	Name     string `db:"name" json:"name"`
	TagCount string `db:"tag_count" json:"tag_count"`
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
	if err != nil {
		return nil, fmt.Errorf("GetPostsByTagId error: %w", err)
	}
	return posts, nil
}

func (m *TagModel) Statistics() ([]*TagList, error) {
	query := `
		select distinct 
			pt.tag_id, 
			t.name,
			count(tag_id) as tag_count 
		from tags t 
		join post_tags pt
		where pt.tag_id = t.id
		group by pt.tag_id, t.name
		union
		select id as tag_id, name, 0 as tag_count from tags where tag_id not in (
			select tag_id from post_tags
		);`

	tags := []*TagList{}
	err := m.DB.Select(&tags, query)
	if err != nil {
		return tags, fmt.Errorf("Statistics() error: %w", err)
	}
	return tags, nil
}
