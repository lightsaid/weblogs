package global

import (
	"lightsaid.com/weblogs/configs"
)

var Config *configs.Config

type ContextKey string

const KeyIsAuthenticated = ContextKey("isAuthenticated")
