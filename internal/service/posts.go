package service

import (
	"fmt"

	"lightsaid.com/weblogs/internal/models"
	"lightsaid.com/weblogs/pkg/utils"
)

type DesktopPost struct {
	Post       models.Post
	Categories []*models.Category
	Attributes []*models.Attribute
}

type DesktopResponse struct {
	Posts         []DesktopPost
	PageIndex     int
	PageSize      int
	CategoryList  []*models.Category
	AttributeList []*models.Attribute
}

func GetDedaultPostThumb() string {

	index := utils.RandomInt(1, 12)

	url := fmt.Sprintf("./static/images/pexels-post-cover-%d.jpg", index)

	return url
}
