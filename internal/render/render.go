package render

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/csrf"
	"github.com/gorilla/sessions"
	"go.uber.org/zap"
	"lightsaid.com/weblogs/internal/models"
)

const tplPath = "./templates"

var session *sessions.Session

type TemplateData struct {
	Cache       map[string]*template.Template
	UseCache    bool // 是否使用缓存
	CookieStore *sessions.CookieStore
}

// New 实例化TemplateData，里面包含 模板缓存，提供给handlers包使用
func New(use bool, store *sessions.CookieStore) *TemplateData {
	cache, err := CreateTemplateCache()
	if err != nil {
		return nil
	}
	return &TemplateData{
		Cache:       cache,
		UseCache:    use,
		CookieStore: store,
	}
}

func (t TemplateData) AddBaseData(td *models.TemplateData, r *http.Request, w http.ResponseWriter) *models.TemplateData {

	// data.Flash = "成功提示"
	// data.Error = "错误提示"
	// data.Warning = "警告提示"
	td.RunMode = os.Getenv("RUNMODE")

	// csrfField 字段是 csrf.TemplateTag 提供，约定名字
	td.Data["csrfField"] = csrf.TemplateField(r)

	session, _ = t.CookieStore.Get(r, os.Getenv("SESSION"))
	fmt.Println(">>>>>>>>>>>>>>>>>>>>before", session)
	msgs := session.Flashes("Success")
	if len(msgs) > 0 {
		td.Success = fmt.Sprint(msgs[0])
	}
	session.Save(r, w)
	fmt.Println(">>>>>>>>>>>>>>>>>>>>after", session)

	return td
}

// Render 获取模板并渲染
func (t TemplateData) Render(w http.ResponseWriter, r *http.Request, tmpl string, data *models.TemplateData) error {
	var err error
	cache := t.Cache
	if !t.UseCache {
		cache, err = CreateTemplateCache()
		if err != nil {
			return err
		}
	}

	tt, ok := cache[tmpl]
	if !ok {
		return fmt.Errorf("模板名字不存在：%s", tmpl)
	}

	// w.Header().Add("Content-Type", "text/html; charset=utf-8")

	// Execute 如果内存存在错误（例如：一个页面由header和content组成，加入header加载正确，而content错误）也会渲染到页面上了。
	// 因此需要中转一下
	// err = tt.Execute(w, nil)
	// if err != nil {
	// 	zap.S().Error("解析模板发生错误", err)
	// }

	data = t.AddBaseData(data, r, w)

	buf := new(bytes.Buffer)
	err = tt.Execute(buf, data)
	if err != nil {
		zap.S().Error("解析模板发生错误", err)
		return err
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		zap.S().Error("写入模板发生错误", err)
		return err
	}

	return nil
}

// CreateTemplateCache 创建所有模板缓存
func CreateTemplateCache() (map[string]*template.Template, error) {

	cache := make(map[string]*template.Template)

	matches, err := filepath.Glob(fmt.Sprintf("%s/*.page.tmpl", tplPath))
	if err != nil {
		zap.S().Panic(err)
		return nil, err
	}

	zap.S().Info("matchs=>> ", matches)

	if len(matches) <= 0 {
		err = errors.New("没有匹配到模板文件")
		zap.S().Panic(err)
		return nil, err
	}

	for _, page := range matches {
		name := filepath.Base(page)

		// 使用 ParseGlob 创建模板，它可能包含其他组件（比如：布局layout组件）
		t, err := template.New(name).ParseGlob(page)
		if err != nil {
			zap.S().Panic(err)
			return nil, err
		}

		// 匹配布局模板，添加到name模板上
		layouts, err := filepath.Glob(fmt.Sprintf("%s/*.layout.tmpl", tplPath))
		if err != nil {
			zap.S().Panic(err)
			return nil, err
		}

		// 将组件添加到 page 模板上
		t, err = t.ParseFiles(layouts...)
		if err != nil {
			zap.S().Panic(err)
			return nil, err
		}

		cache[name] = t
	}

	return cache, nil
}
