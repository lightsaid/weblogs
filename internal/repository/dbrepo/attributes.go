package dbrepo

import (
	"errors"
	"fmt"

	"lightsaid.com/weblogs/internal/models"
)

func (repo *databaseRepo) InsertAttrs(attr *models.Attribute) (*models.Attribute, error) {
	query := `insert into attributes(user_id, kind, name)
				values($1, $2, $3) returning *`
	var a models.Attribute
	err := repo.DB.Get(&a, query, attr.UserID, attr.Kind, attr.Name)
	return &a, err
}

func (repo *databaseRepo) GetAttributes() ([]*models.Attribute, error) {
	var attrs []*models.Attribute
	query := `select id, user_id, kind, name from attributes;`

	err := repo.DB.Select(&attrs, query)

	return attrs, err
}

func (repo *databaseRepo) GetAttributesByIds(ids []int) ([]*models.Attribute, error) {
	var attrs []*models.Attribute

	query := `select id, user_id, kind, name from attributes where id in (`
	var idstr string
	for i := 0; i < len(ids); i++ {
		idstr += fmt.Sprintf("%d", ids[i])
		if i != len(ids)-1 {
			idstr += ","
		}
	}
	fmt.Println(">>>>>>> idstr 1", idstr)
	idstr = idstr + ")"
	fmt.Println(">>>>>>> idstr 2", idstr)

	query += idstr

	fmt.Println("get attrs query >>>> ", query)

	err := repo.DB.Select(&attrs, query)

	return attrs, err
}

func (repo *databaseRepo) GetAttributesByUserID(id int) ([]*models.Attribute, error) {
	var attrs []*models.Attribute
	query := `select id, user_id, kind, name from attributes where user_id=$1;`

	err := repo.DB.Select(&attrs, query, id)

	return attrs, err
}

func (repo *databaseRepo) UpdateAttributes(a *models.Attribute) error {
	query := `update attributes set name=$1, kind=$2, user_id=$3 where id=$4`

	result, err := repo.DB.Exec(query, a.Name, a.Kind, a.UserID, a.ID)
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

func (repo *databaseRepo) DeleteAttribute(id int) error {
	query := `delete from attributes where id =$1;`
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

func (repo *databaseRepo) GetAttributesByPostID(postId int) ([]*models.Attribute, error) {

	query := `select id, post_id, attr_id from pa_mapping where post_id=$1;`
	pamappings := []models.PAMapping{}

	err := repo.DB.Select(&pamappings, query, postId)
	if err != nil {
		return nil, err
	}

	var attrs []*models.Attribute
	if len(pamappings) > 0 {
		ids := []int{}
		for _, v := range pamappings {
			ids = append(ids, v.AttrID)
		}

		attrs, err = repo.GetAttributesByIds(ids)
		if err != nil {
			return nil, err
		}
	}

	return attrs, err
}
