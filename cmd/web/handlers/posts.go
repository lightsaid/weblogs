package handlers

import (
	"net/http"

	"lightsaid.com/weblogs/internal/models"
)

func ShowAdminPosts(w http.ResponseWriter, r *http.Request) {
	td := models.NewTemplateData()
	td.Menubar.PostList = true
	H.Template.Render(w, r, "posts.page.tmpl", &td)
}

func GetPosts(w http.ResponseWriter, r *http.Request) {

}

func CreatePost(w http.ResponseWriter, r *http.Request) {

}

func GetPost(w http.ResponseWriter, r *http.Request) {

}

func UpdatePost(w http.ResponseWriter, r *http.Request) {

}

func DeletePost(w http.ResponseWriter, r *http.Request) {

}
