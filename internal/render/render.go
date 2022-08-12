package render

import (
	"errors"
	"fmt"
	"html/template"
	"path/filepath"

	"go.uber.org/zap"
)

const tplPath = "./templates"

type TemplateData struct {
	Cache map[string]*template.Template
}

// New 实例化TemplateData，里面包含 模板缓存，提供给handlers包使用
func New() *TemplateData {
	cache, err := CreateTemplateCache()
	if err != nil {
		zap.S().Panic(err)
		return nil
	}

	return &TemplateData{
		Cache: cache,
	}
}

func Render() {

}

// CreateTemplateCache 创建所有模板缓存
func CreateTemplateCache() (map[string]*template.Template, error) {

	cache := make(map[string]*template.Template)

	matches, err := filepath.Glob(fmt.Sprintf("%s/*.page.html", tplPath))
	if err != nil {
		return nil, err
	}

	zap.S().Info("matchs=>> ", matches)

	if len(matches) <= 0 {
		return nil, errors.New("没有匹配到模板文件")
	}

	for _, page := range matches {
		name := filepath.Base(page)

		// 使用 ParseGlob 创建模板，它可能包含其他组件（比如：布局layout组件）
		t, err := template.New(name).ParseGlob(page)
		if err != nil {
			return nil, err
		}

		// 匹配布局模板，添加到name模板上
		layouts, err := filepath.Glob(fmt.Sprintf("%s/*.layout.html", tplPath))
		if err != nil {
			return nil, err
		}

		// 将组件添加到 page 模板上
		t, err = t.ParseFiles(layouts...)
		if err != nil {
			return nil, err
		}

		cache[name] = t
	}

	return cache, nil

}
