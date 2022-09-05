package service

import (
	"go.uber.org/zap"
	"lightsaid.com/weblogs/internal/models"
	"lightsaid.com/weblogs/internal/repository"
)

type LevelCategories struct {
	Category *models.Category
	Children []*LevelCategories
}

// GetLevelCategories 获取等级分类
func (s *Service) GetLevelCategories() ([]*LevelCategories, error) {
	var categories []*LevelCategories

	// NOTE: 此处使用递归，没有返分类列表返回值，因此需要传地址
	err := getLevelCategories(s.Repository, 0, &categories)

	return categories, err
}

func getLevelCategories(repo repository.Repository, id int, categories *[]*LevelCategories) error {
	cates, err := repo.GetCategories(id)
	zap.S().Info("cates >>> ", cates)
	if err != nil {
		return err
	}
	for _, c := range cates {
		var children []*LevelCategories
		var item = &LevelCategories{Category: c, Children: children}
		// NOTE:
		*categories = append(*categories, item)

		if c.IfParent > 0 {
			// 递归获取子分类
			getLevelCategories(repo, c.ID, &item.Children)

			// 获取二级分类
			// child_cates, err := repo.GetCategories(c.ID)
			// zap.S().Info("child_cates >>> ", child_cates)
			// if err != nil {
			// 	return err
			// }
			// for _, c := range child_cates {
			// 	var children []*LevelCategories
			// 	item.Children = append(item.Children, &LevelCategories{Category: c, Children: children})
			// }
		}
	}

	return nil
}
