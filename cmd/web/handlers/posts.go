package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"lightsaid.com/weblogs/internal/models"
	"lightsaid.com/weblogs/internal/service"
	"lightsaid.com/weblogs/internal/validator"
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

	var jsonResponse service.JSONResponse
	var req service.CreatePostRequest
	err := H.readJSON(w, r, &req)
	if err != nil {
		zap.S().Error(err)
		H.errorJSONResponse(w)
		return
	}
	zap.S().Info("创建文章参数：", req)
	jsvalid, err := validator.NewJsonValidator(req)
	if err != nil {
		zap.S().Error(err)
		H.errorJSONResponse(w)
		return
	}
	jsvalid.Required("title", "content")
	if !jsvalid.Valid() {
		msg := fmt.Sprintf("%s, %s", jsvalid.Errors.Get("title"), jsvalid.Errors.Get("content"))
		if len(msg) < 5 {
			msg = "验证步通过"
		}
		jsonResponse.Message = msg
		jsonResponse.Error = true
		H.writeJSON(w, http.StatusBadRequest, jsonResponse)
		return
	}

	req.Thumb = service.GetDedaultPostThumb()

	// 组织数据
	post := models.Post{
		UserId:  info.UserID,
		Author:  info.Username,
		Title:   req.Title,
		Content: req.Content,
		Thumb:   &req.Thumb,
	}

	// TODO: 设置分类，设置属性

	newPost, err := H.Repo.InsertPost(post)
	if err != nil {
		zap.S().Error(err)
		jsonResponse.Error = true
		jsonResponse.Message = "添加出错"
		return
	}

	jsonResponse.Error = false
	jsonResponse.Data = newPost
	_ = H.writeJSON(w, http.StatusOK, jsonResponse)

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
