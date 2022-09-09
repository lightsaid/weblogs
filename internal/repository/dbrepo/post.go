package dbrepo

import (
	"fmt"

	"github.com/jmoiron/sqlx"
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

// InserManyPC 批量添加文章分类 mapping
func (repo *databaseRepo) InserManyPC(pcs []interface{}) error {
	var manyValues string
	for i := 0; i < len(pcs); i++ {
		manyValues += " (?),"
	}
	manyValues = manyValues[0 : len(manyValues)-1]
	var query = fmt.Sprintf("insert into pc_mapping(post_id,cate_id) values %s", manyValues)
	query, args, err := sqlx.In(query, pcs...)
	if err != nil {
		return err
	}
	_, err = repo.DB.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
	// TODO: 待用事务
}

// InserManyPA 批量添加文章分类 mapping
func (repo *databaseRepo) InserManyPA(pas []interface{}) error {
	var manyValues string
	for i := 0; i < len(pas); i++ {
		manyValues += " (?),"
	}
	manyValues = manyValues[0 : len(manyValues)-1]
	var query = fmt.Sprintf("insert into pa_mapping(post_id,attr_id) values %s", manyValues)
	query, args, err := sqlx.In(query, pas...)
	if err != nil {
		return err
	}
	_, err = repo.DB.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
	// TODO: 待用事务

}

// func BatchInsertUsers2(users []interface{}) error {
// 	query, args, _ := sqlx.In(
// 		"INSERT INTO user (name, age) VALUES (?), (?), (?)",
// 		users..., // 如果arg实现了 driver.Valuer, sqlx.In 会通过调用 Value()来展开它
// 	)
// 	fmt.Println(query) // 查看生成的querystring
// 	fmt.Println(args)  // 查看生成的args
// 	_, err := DB.Exec(query, args...)
// 	return err
// }
