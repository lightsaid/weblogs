package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"lightsaid.com/weblogs/internal/models"
	"lightsaid.com/weblogs/internal/service"
)

var (
	pageIndex int = 0
	pageSize  int = 10
)

func ShowAdminPosts(w http.ResponseWriter, r *http.Request) {
	td := models.NewTemplateData()
	td.Menubar.PostList = true

	vars := mux.Vars(r)
	pageIndexS := vars["page_index"]
	pageSizeS := vars["page_size"]

	index, err := strconv.Atoi(pageIndexS)
	if err != nil {
		index = pageIndex
	}
	size, err := strconv.Atoi(pageSizeS)
	if err != nil {
		size = pageSize
	}

	posts, err := H.Repo.GetPosts(size, index)
	if err != nil {
		zap.S().Error(err)
		ServerError(w, err)
		return
	}

	td.Data["posts"] = posts

	H.Template.Render(w, r, "posts.page.tmpl", &td)
}

func ShowPublishPost(w http.ResponseWriter, r *http.Request) {
	td := models.NewTemplateData()

	attrs, err := H.Repo.GetAttributes()
	if err != nil {
		td.Error = "获取属性出错"
	}
	// td.Success = "获取属性成功"
	td.Data["attrs"] = attrs

	// 获取分类
	categories, err := H.Repo.GetLevelCategories()
	if err != nil {
		td.Error = "获取分类出错"
	}
	td.Data["categories"] = categories
	zap.S().Info("分类获取>>> ", categories)
	H.Template.Render(w, r, "publish.page.tmpl", &td)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	session := GetSession(w, r)
	userinfo := r.Context().Value("userinfo")
	info, ok := userinfo.(service.SessionUser)
	if !ok {
		session.AddFlash("登录失效,请重新登录", "Error")
		http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		return
	}
	p := models.Post{
		UserId:  info.UserID,
		Author:  info.Username,
		Title:   "post",
		Content: "测试测试测试",
	}
	post, err := H.Repo.InsertPost(p)
	if err != nil {
		session.AddFlash("添加失败", "Error")
		session.Save(r, w)
		http.Redirect(w, r, "/admin/posts", http.StatusSeeOther)
		return
	}
	fmt.Println("new post >>> ", post)

	session.AddFlash("添加成功", "Success")
	session.Save(r, w)
	http.Redirect(w, r, "/admin/posts", http.StatusSeeOther)
}

func CreateJsonPost(w http.ResponseWriter, r *http.Request) {

	// H.readJSON()
}

func GetPosts(w http.ResponseWriter, r *http.Request) {

}

func GetPost(w http.ResponseWriter, r *http.Request) {

}

func UpdatePost(w http.ResponseWriter, r *http.Request) {

}

func DeletePost(w http.ResponseWriter, r *http.Request) {

}

func ShowDemo(w http.ResponseWriter, r *http.Request) {
	td := models.NewTemplateData()
	H.Template.Render(w, r, "example.page.tmpl", &td)
}
