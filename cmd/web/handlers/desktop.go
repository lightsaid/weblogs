package handlers

import (
	"net/http"

	"lightsaid.com/weblogs/internal/models"
	"lightsaid.com/weblogs/internal/service"
)

func ShowDesktop(w http.ResponseWriter, r *http.Request) {
	td := models.NewTemplateData()
	td.Menubar.Desktop = true

	user := models.User{}
	userinfo := r.Context().Value("userinfo")
	if info, ok := userinfo.(service.SessionUser); ok {
		user, _ = H.Repo.GetUser(info.UserID)
	}
	td.Data["user"] = user
	H.Template.Render(w, r, "index.page.tmpl", &td)
}

func ShowColumn(w http.ResponseWriter, r *http.Request) {
	td := models.NewTemplateData()
	td.Menubar.About = true
	H.Template.Render(w, r, "column.page.tmpl", &td)
}

func ShowAbout(w http.ResponseWriter, r *http.Request) {
	td := models.NewTemplateData()
	td.Menubar.About = true
	H.Template.Render(w, r, "about.page.tmpl", &td)
}
