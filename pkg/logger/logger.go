package logger

import (
	"go.uber.org/zap"
)

func New(paths ...string) (*zap.Logger, error) {
	cfg := zap.NewDevelopmentConfig()
	if len(paths) > 0 {
		cfg.OutputPaths = paths
	}else{
		cfg.OutputPaths = []string{
			"stderr",
		}
	}

	l, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	// zap.S() 和 zap.L() 函数是zap提供给全局访问logger的方式
	// 此处是替换默认全局,使用自定义
	_ = zap.ReplaceGlobals(l)
	return l, err
}