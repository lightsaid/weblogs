package handlers

import (
	"html/template"
	"net/http"

	"go.uber.org/zap"
)

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
	H.Template.Render(w, r, "admin.page.tmpl")
}

func ShowAdminPosts(w http.ResponseWriter, r *http.Request) {

}
