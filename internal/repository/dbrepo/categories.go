package dbrepo

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.uber.org/zap"
	"lightsaid.com/weblogs/internal/models"
)

func (repo *databaseRepo) InsertCategories(cate *models.Category, parentID int) (*models.Category, error) {
	// 插入分类 SQL
	insertSQL := `insert into categories(user_id, parent_id, if_parent, name, thumb)
				values($1, $2, $3, $4, $5) returning *`

	if parentID <= 0 {
		var c models.Category
		err := repo.DB.Get(&c, insertSQL, cate.UserID, cate.ParentID, cate.IfParent, cate.Name, cate.Thumb)
		if err != nil {
			return nil, err
		}
		return &c, nil
	}

	// 查找父类 SQL
	querySQL := `select id, user_id, parent_id, if_parent, name, thumb from categories where id = $1;`

	// 更新父类if_parent=1
	updateSQL := `update categories set if_parent=1 where id=$1`

	var c models.Category
	var parent_cate models.Category
	var err error

	// NOTE: 开启sqlx事务
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 事务开始
	// NOTE: BeginTx() 返回的时 sql.Tx 非 sqlx.TX
	// tx, err := repo.DB.BeginTx(ctx, nil)

	// 使用 sqlx.Tx 操作更加简便
	tx := repo.DB.MustBeginTx(ctx, nil)
	// 插入分类
	if err = tx.Get(&c, insertSQL, cate.UserID, cate.ParentID, cate.IfParent, cate.Name, cate.Thumb); err != nil {
		zap.S().Error("tx.Get error: ", err)
		err = tx.Rollback()
		if err != nil {
			zap.S().Panic("tx.Rollback error: ", err)
			return nil, err
		}
		return nil, err
	}

	if err = tx.Get(&parent_cate, querySQL, parentID); err != nil {
		zap.S().Error("tx.Get error: ", err)
		err = tx.Rollback()
		if err != nil {
			zap.S().Error("tx.Rollback error: ", err)
			return nil, err
		}
		return nil, err
	}

	// 父类 if_parent 字段已修改
	if parent_cate.IfParent > 0 {
		err = tx.Commit()
		if err != nil {
			zap.S().Error("tx.Commit error: ", err)
			return nil, err
		}
		return &c, nil
	}

	// 父类 if_parent 字段未修改，下面开始修改
	result, err := tx.Exec(updateSQL, parentID)
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			zap.S().Error("tx.Rollback error: ", err)
			return nil, err
		}
		return nil, err
	}

	if row, err := result.RowsAffected(); err != nil {
		err = tx.Rollback()
		if err != nil {
			zap.S().Error("tx.Rollback error: ", err)
			return nil, err
		}
		return nil, err
	} else {
		if row > 0 {
			err = tx.Commit()
			if err != nil {
				zap.S().Error("tx.Commit error: ", err)
				return nil, err
			}
			return &c, nil
		}
		return &c, errors.New("更新父类 if_parent 字段影响行数为0")
	}
}

func (repo *databaseRepo) GetCategoriesNoParentID(limit int, offese int) ([]*models.Category, error) {
	var cates []*models.Category
	query := `select id, user_id, parent_id, if_parent, name, thumb from categories limit $1 offset $2;`

	err := repo.DB.Select(&cates, query, limit, offese)

	return cates, err
}

func (repo *databaseRepo) GetCategories(parent_id int) ([]*models.Category, error) {
	var cates []*models.Category
	query := `select id, user_id, parent_id, if_parent, name, thumb from categories where parent_id = $1;`

	err := repo.DB.Select(&cates, query, parent_id)

	return cates, err
}

func (repo *databaseRepo) GetCategoriesByIds(ids []int) ([]*models.Category, error) {
	var cates []*models.Category
	query := `select id, user_id, parent_id, if_parent, name, thumb from categories where id in (`
	var idstr string
	for i := 0; i < len(ids); i++ {
		idstr += fmt.Sprintf("%d", ids[i])
		if i != len(ids)-1 {
			idstr += ","
		}
	}
	idstr = idstr + ")"
	query += idstr

	fmt.Println("get cate query >>>> ", query)

	err := repo.DB.Select(&cates, query)

	return cates, err
}

func (repo *databaseRepo) GetCategoriesByUserID(parent_id int, userID int) ([]*models.Category, error) {
	var cates []*models.Category
	query := `select id, user_id, parent_id, if_parent, name, thumb from categories where user_id =$1 and parent_id = $2;`

	err := repo.DB.Select(&cates, query, userID, parent_id)

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

func (repo *databaseRepo) GetCategoriesById(id int) (*models.Category, error) {
	var cate models.Category
	query := `select id, user_id, parent_id, if_parent, name, thumb from categories where id = $1;`

	err := repo.DB.Get(&cate, query, id)

	return &cate, err
}

func (repo *databaseRepo) DeleteCategories(id int) error {
	query := `delete from categories where id =$1;`
	result, err := repo.DB.Exec(query, id)
	if err != nil {
		return err
	}
	row, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if row <= 0 {
		return errors.New("数据不存在，删除失败")
	}
	return nil
}

func (repo *databaseRepo) GetCategoriesByPostID(postId int) ([]*models.Category, error) {

	query := `select id, post_id, cate_id from pc_mapping where post_id=$1;`
	pcmappings := []models.PCMapping{}

	err := repo.DB.Select(&pcmappings, query, postId)
	if err != nil {
		return nil, err
	}
	var cates []*models.Category
	if len(pcmappings) > 0 {
		ids := []int{}
		for _, v := range pcmappings {
			ids = append(ids, v.CateID)
		}

		cates, err = repo.GetCategoriesByIds(ids)
		if err != nil {
			return nil, err
		}
	}

	return cates, nil
}
