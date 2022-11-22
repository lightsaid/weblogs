package controller

import (
	"fmt"

	"lightsaid.com/weblogs/global"
	"lightsaid.com/weblogs/utils"
)

// Toolkit 一个额外的工具结构，嵌套在 Controller，提供辅助方法
type Toolkit struct {
}

// GenDefaultUserName 返回一个随机用户名
func (t *Toolkit) GenDefaultUserName() string {
	return global.Config.DefaultUserNamePrefix + utils.RandomString(4)
}

// GenDefaultAvatar 返回一个随机默认头像
func (t *Toolkit) GenDefaultAvatar() string {
	return fmt.Sprintf("/static/images/default/avatar%d.png", utils.RandomInt(1, 3))
}
