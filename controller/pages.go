package controller

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
	"lightsaid.com/weblogs/forms"
)

func (_this *Controller) IndexPage(w http.ResponseWriter, r *http.Request) {
	td := TemplateData{}
	posts, err := _this.Models.Posts.GetList(DefaultPageSize, 0)
	if err != nil {
		// TODO: 提示error
		log.Error(err)
		return
	}
	td.DataMap = map[string]interface{}{"posts": posts}
	_this.Render(w, r, "index.page.gtpl", &td)
}

func (_this *Controller) EditorMDFullExample(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second * 1)
	_this.Render(w, r, "full.page.gtpl", nil)
}

func (_this *Controller) LoginPage(w http.ResponseWriter, r *http.Request) {
	td := TemplateData{Form: forms.New(nil)}
	_this.Render(w, r, "login.page.gtpl", &td)
}

func (_this *Controller) TagPage(w http.ResponseWriter, r *http.Request) {
	tags, err := _this.Models.Tags.Statistics()
	if err != nil {
		// TODO:
		log.Error(err)
		return
	}
	td := TemplateData{Form: forms.New(nil)}
	td.DataMap = map[string]interface{}{"tags": tags}
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
	id := _this.session.GetInt(r.Context(), "user_id")
	if id <= 0 {
		// TODO: message
		http.Redirect(w, r, "/user/login", http.StatusSeeOther)
		return
	}
	user, err := _this.Models.Users.GetById(id)
	if err != nil {
		// TODO: message
		http.Redirect(w, r, "/user/login", http.StatusSeeOther)
		return
	}
	td := TemplateData{Form: forms.New(nil)}
	td.DataMap = map[string]interface{}{"user": user}
	_this.Render(w, r, "settings.page.gtpl", &td)
}
