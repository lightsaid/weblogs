package dbrepo

import (
	"fmt"
	"time"

	"lightsaid.com/weblogs/internal/models"
)

// users 表CRUD操作

func (repo *databaseRepo) InsertUser(email, username, password, avatar string) (models.User, error) {
	query := `insert into users(email, username, password, avatar)
				values($1, $2, $3, $4) returning *`
	var u models.User
	err := repo.DB.Get(&u, query, email, username, password, avatar)
	return u, err
}

func (repo *databaseRepo) GetUser(id int) (models.User, error) {
	var user models.User
	return user, nil
}

func (repo *databaseRepo) GetUserByEmial(email string) (models.User, error) {
	query := `select id, email, password, username, avatar, if_admin, active from users where email=$1 limit 1;`
	user := models.User{}
	err := repo.DB.Get(&user, query, email)
	return user, err
}

func (repo *databaseRepo) GetUsers() ([]models.User, error) {
	var users []models.User
	query := `select id, email, username, avatar, if_admin, active from users 
		order by created_at, active desc limit 10 offset 0;`

	err := repo.DB.Select(&users, query)

	return users, err
}

var ErrUpdateNoRows = fmt.Errorf("影响 0 行数据")

func (repo *databaseRepo) UpdateUser(user models.User) error {

	query := `update users set username=$1, avatar=$2, if_admin=$3, updated_at=$4 where id=$5;`

	result, err := repo.DB.Exec(query, user.Username, user.Avatar, user.IfAdmin, time.Now(), user.ID)
	if err != nil {
		return err
	}

	var n int64
	if n, err = result.RowsAffected(); err != nil {
		return err
	}

	if n <= 0 {
		return ErrUpdateNoRows
	}
	return nil
}

func (repo *databaseRepo) DeleteUser(id int) error {
	query := "delete from users where id = $1"
	ret, err := repo.DB.Exec(query, id)
	if err != nil {
		return err
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		return err
	}
	if n <= 0 {
		return ErrUpdateNoRows
	}
	return nil
}
