package configs

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	// dev|prod
	Mode      string `json:"mode"`
	Port      int    `json:"port"`
	DBSource  string `json:"db_source"`
	LogOutput string `json:"log_output"`
	// 根据 logrus 库日志等级设置： panic fatal error warn info debug trace
	LogLevel              string `json:"log_level"`
	DefaultUserNamePrefix string `json:"default_username_prefix"`
	Domain                string `json:"domain"`
	ImageBaseUrl          string `json:"image_base_url"`
}

func ReadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("open config file error: %w", err)
	}
	var conf = new(Config)
	err = json.Unmarshal(data, conf)
	if err != nil {
		return nil, fmt.Errorf("json unmarshal config file error: %w", err)
	}

	return conf, nil
}
