package handlers

import (
	"html/template"
	"net/http"

	"go.uber.org/zap"
	"lightsaid.com/weblogs/internal/models"
)

func ShowAdminLogin(w http.ResponseWriter, r *http.Request) {
	data := models.NewTemplateData()
	data.Title = "登录"
	H.Template.Render(w, r, "login.page.tmpl", &data)
}

func ShowAdminIndex(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/admin.page.tmpl")
	if err != nil {
		zap.S().Error("解析模板发生错误", err)
	}
	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	PageTitle := "Admin"
	err = t.Execute(w, PageTitle)
	if err != nil {
		zap.S().Error("解析模板发生错误", err)
	}
}

func ShowAdminUsers(w http.ResponseWriter, r *http.Request) {
	var data = models.NewTemplateData()
	users, err := H.Repo.GetUsers()
	if err != nil {
		data.Error = err.Error()
	}
	data.Data["Users"] = users
	H.Template.Render(w, r, "users.page.tmpl", &data)
}

func ShowAdminPosts(w http.ResponseWriter, r *http.Request) {

}
