package controller

import (
	"net/http"
	"time"

	"lightsaid.com/weblogs/forms"
)

func (c *Controller) IndexPage(w http.ResponseWriter, r *http.Request) {
	td := TemplateData{}
	c.Render(w, r, "index.page.gtpl", &td)
}

func (c *Controller) EditorMDFullExample(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second * 2)
	var arr [2]int
	var index = 3
	arr[index] = 100
	c.Render(w, r, "full.page.gtpl", nil)
}

func (_this *Controller) LoginPage(w http.ResponseWriter, r *http.Request) {
	td := TemplateData{Form: *forms.New(nil)}
	_this.Render(w, r, "login.page.gtpl", &td)
}
