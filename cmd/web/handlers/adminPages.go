package handlers

import (
	"net/http"

	"lightsaid.com/weblogs/internal/models"
	"lightsaid.com/weblogs/internal/validator"
)

func ShowAdminLogin(w http.ResponseWriter, r *http.Request) {
	data := models.NewTemplateData()
	data.JsonValidator, _ = validator.NewJsonValidator(nil)
	H.Template.Render(w, r, "login.page.tmpl", &data)
}

func ShowAdminIndex(w http.ResponseWriter, r *http.Request) {
	data := models.NewTemplateData()
	H.Template.Render(w, r, "dashboard.page.tmpl", &data)
}

func ShowAdminUsers(w http.ResponseWriter, r *http.Request) {
	var data = models.NewTemplateData()
	users, err := H.Repo.GetUsers()
	if err != nil {
		data.Error = err.Error()
	}
	data.Data["users"] = users
	H.Template.Render(w, r, "users.page.tmpl", &data)
}

func ShowAdminPosts(w http.ResponseWriter, r *http.Request) {

}
