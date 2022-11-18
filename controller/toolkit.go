package controller

import (
	"fmt"

	"lightsaid.com/weblogs/global"
	"lightsaid.com/weblogs/utils"
)

type Toolkit struct {
}

func (t *Toolkit) GenDefaultUserName() string {
	return global.Config.DefaultUserNamePrefix + utils.RandomString(4)
}

func (t *Toolkit) GenDefaultAvatar() string {
	return fmt.Sprintf("/static/images/default/%d.png", utils.RandomInt(1, 3))
}
