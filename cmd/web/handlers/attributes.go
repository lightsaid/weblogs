package handlers

import (
	"net/http"

	"lightsaid.com/weblogs/internal/models"
)

func ShowAdminAttrs(w http.ResponseWriter, r *http.Request) {
	td := models.NewTemplateData()
	td.Menubar.AttributeList = true
	H.Template.Render(w, r, "attrs.page.tmpl", &td)
}

func GetAttributes(w http.ResponseWriter, r *http.Request) {

}

func CreateAttribute(w http.ResponseWriter, r *http.Request) {

}

func GetAttribute(w http.ResponseWriter, r *http.Request) {

}

func UpdateAttribute(w http.ResponseWriter, r *http.Request) {

}

func DeleteAttribute(w http.ResponseWriter, r *http.Request) {

}
