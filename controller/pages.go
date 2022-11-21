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
	time.Sleep(time.Second * 1)
	c.Render(w, r, "full.page.gtpl", nil)
}

func (_this *Controller) LoginPage(w http.ResponseWriter, r *http.Request) {
	td := TemplateData{Form: forms.New(nil)}
	_this.Render(w, r, "login.page.gtpl", &td)
}

func (_this *Controller) TagPage(w http.ResponseWriter, r *http.Request) {
	td := TemplateData{Form: forms.New(nil)}
	_this.Render(w, r, "tags.page.gtpl", &td)
}

func (_this *Controller) NotificationPage(w http.ResponseWriter, r *http.Request) {
	td := TemplateData{Form: forms.New(nil)}
	_this.Render(w, r, "notifications.page.gtpl", &td)
}

func (_this *Controller) AboutPage(w http.ResponseWriter, r *http.Request) {
	td := TemplateData{Form: forms.New(nil)}
	_this.Render(w, r, "about.page.gtpl", &td)
}

func (_this *Controller) CreatePostPage(w http.ResponseWriter, r *http.Request) {
	td := TemplateData{Form: forms.New(nil)}
	_this.Render(w, r, "create.post.page.gtpl", &td)
}

func (_this *Controller) SettingsPage(w http.ResponseWriter, r *http.Request) {
	td := TemplateData{Form: forms.New(nil)}
	_this.Render(w, r, "settings.page.gtpl", &td)
}
