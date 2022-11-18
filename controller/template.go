package controller

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/ory/nosurf"
	"github.com/sirupsen/logrus"
	"lightsaid.com/weblogs/forms"
	"lightsaid.com/weblogs/global"
)

type HTMLTemplate struct {
	cache map[string]*template.Template
}

type TemplateData struct {
	StringMap map[string]string
	DataMap   map[string]interface{}
	Form      forms.Form
	CSRFToken string
	Error     string
	Success   string
	IsLogin   int
}

func NewHTMLTemplate() (*HTMLTemplate, error) {
	data, err := createCache()
	if err != nil {
		return nil, err
	}
	t := &HTMLTemplate{
		data,
	}

	return t, nil
}

var funcMap = template.FuncMap{}

func createCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob(PagePathPattern)
	if err != nil {
		err = fmt.Errorf("matching page template error: %w", err)
		return cache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(funcMap).ParseFiles(page)
		if err != nil {
			err = fmt.Errorf("new template with parseFiles error: %w", err)
			return cache, err
		}

		// 添加布局模版
		ts, err = ts.ParseGlob(LayoutPathPattern)
		if err != nil {
			err = fmt.Errorf("template parse layout components error: %w", err)
			return cache, err
		}

		// 添加其他公共组建
		ts, err = ts.ParseGlob(PartialPathPattern)
		if err != nil {
			err = fmt.Errorf("template parse partial components error: %w", err)
			return cache, err
		}

		cache[name] = ts
	}

	return cache, nil
}

func (t *HTMLTemplate) Render(w http.ResponseWriter, r *http.Request, name string, data *TemplateData) {
	var cache map[string]*template.Template
	var err error

	// 追加基础数据
	data = t.appendTemplateData(w, r, data)

	// 获取模板
	if global.Config.Mode == DevModeValue {
		cache, err = createCache()
		if err != nil {
			t.ServerError(w, err)
		}
	} else {
		cache = t.cache
	}

	template, ok := cache[name]
	if !ok {
		err = errors.New("template not found")
		t.ServerError(w, err)
		return
	}

	buf := new(bytes.Buffer)

	// 执行模板到指定 bytes.Buffer，感知错误
	err = template.Execute(buf, data)
	if err != nil {
		logrus.Error("exectue template error: " + err.Error())
		w.Write([]byte("模板渲染错误"))
		return
	}

	// 写入 http.ResponseWriter 流
	_, err = buf.WriteTo(w)
	if err != nil {
		logrus.Error("write template error: " + err.Error())
		w.Write([]byte("写入模板错误"))
		return
	}
}

func (t *HTMLTemplate) ServerError(w http.ResponseWriter, err error) {
	logrus.Errorf("server error: %v", err)
	http.Error(w, "服务器发生错误，无法做出响应", http.StatusInternalServerError)
}

// appendTemplateData 追加默认数据
func (t *HTMLTemplate) appendTemplateData(w http.ResponseWriter, r *http.Request, td *TemplateData) *TemplateData {
	td.CSRFToken = nosurf.Token(r)

	return td
}
