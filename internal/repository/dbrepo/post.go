package dbrepo

import "lightsaid.com/weblogs/internal/models"

func (repo *databaseRepo) InsertPost(Post models.Post) (int, error) {
	return 0, nil
}
func (repo *databaseRepo) GetPost(id int) (models.Post, error) {
	return models.Post{}, nil
}
func (repo *databaseRepo) GetPosts() ([]models.Post, error) {
	return []models.Post{}, nil
}
func (repo *databaseRepo) UpdatePost(models.Post) error {
	return nil
}
func (repo *databaseRepo) DeletePost(id int) error {
	return nil
}
