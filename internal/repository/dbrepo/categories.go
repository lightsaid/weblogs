package dbrepo

import (
	"fmt"

	"lightsaid.com/weblogs/internal/models"
)

func (repo *databaseRepo) InsertCategories(cate *models.Category) (*models.Category, error) {
	query := `insert into categories(user_id, parent_id, if_parent, name, thumb)
				values($1, $2, $3, $4, $5) returning *`
	var c models.Category
	err := repo.DB.Get(&c, query, cate.UserID, cate.ParentID, cate.IfParent, cate.Name, cate.Thumb)
	return cate, err
}

func (repo *databaseRepo) GetCategories(parent_id int) ([]*models.Category, error) {
	var cates []*models.Category
	query := `select id, user_id, parent_id, if_parent, name, thumb from categories where parent_id = $1;`

	err := repo.DB.Select(&cates, query, parent_id)

	return cates, err
}

func (repo *databaseRepo) UpdateCategories(a *models.Category) error {
	query := `update categories set name=$1, thumb=$2 where id=$3`

	result, err := repo.DB.Exec(query, a.Name, a.Thumb, a.ID)
	if err != nil {
		return err
	}
	row, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if row <= 0 {
		return fmt.Errorf("更新失败, ID 不存在或其他原因，影响行：%d, 入参：%v", row, a)
	}

	return nil
}

func (repo *databaseRepo) DeleteCategories(id int) error {
	// query := `delete from attributes where id =$1;`
	// result, err := repo.DB.Exec(query, id)
	// if err != nil {
	// 	return err
	// }
	// row, err := result.RowsAffected()
	// if err != nil {
	// 	return err
	// }
	// if row <= 0 {
	// 	return errors.New("数据不存在，删除失败")
	// }
	return nil
}
