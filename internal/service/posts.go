package service

import (
	"fmt"

	"lightsaid.com/weblogs/pkg/utils"
)

func GetDedaultPostThumb() string {

	index := utils.RandomInt(1, 12)

	url := fmt.Sprintf("./static/images/pexels-post-cover-%d.jpg", index)

	return url
}
