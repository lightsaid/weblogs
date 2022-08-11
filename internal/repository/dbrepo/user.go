package dbrepo

import "lightsaid.com/weblogs/internal/models"

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

func (repo *databaseRepo) GetUsers() ([]models.User, error) {
	var users []models.User
	return users, nil
}

func (repo *databaseRepo) UpdateUser(models.User) error {
	return nil
}

func (repo *databaseRepo) DeleteUser(id int) error {
	return nil
}
