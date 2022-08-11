package dbrepo

import "lightsaid.com/weblogs/internal/models"

// users 表CRUD操作

func (repo *databaseRepo) InsertUser(user models.User) (int, error) {
	return 0, nil
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
