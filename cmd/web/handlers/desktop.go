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

func ShowDesktop(w http.ResponseWriter, r *http.Request) {
	td := models.NewTemplateData()
	td.Menubar.Desktop = true
	session := GetSession(w, r)

	user := models.User{}
	userinfo := r.Context().Value("userinfo")
	if info, ok := userinfo.(service.SessionUser); ok {
		user, _ = H.Repo.GetUser(info.UserID)
	}
	td.Data["user"] = user

	var (
		pageIndex = 0
		pageSize  = 10
		response  = service.DesktopResponse{}
	)

	posts, err := H.Repo.GetPosts(pageSize, pageIndex)
	if err != nil {
		zap.S().Error(err)
		session.AddFlash("获取文章列表错误", "Error")
		session.Save(r, w)
		return
	}
	if len(posts) > 0 {
		// 根据post查分类
		for i, _ := range posts {
			var desktop service.DesktopPost
			desktop.Categories, err = H.Repo.GetCategoriesByPostID(posts[i].ID)
			if err != nil {
				zap.S().Error(err)
				session.AddFlash("获取文章分类出错", "Error")
				session.Save(r, w)
			}
			desktop.Attributes, err = H.Repo.GetAttributesByPostID(posts[i].ID)
			if err != nil {
				zap.S().Error(err)
				session.AddFlash("获取文章属性出错", "Error")
				session.Save(r, w)
			}
			fmt.Printf(">>>>>>>>>>>>>> desktop.Attributes :%#v\n", desktop.Attributes)
			desktop.Post = *posts[i]
			response.Posts = append(response.Posts, desktop)
		}
	}

	// 获取右边栏的分类和tags
	response.AttributeList, err = H.Repo.GetAttributes()
	if err != nil {
		zap.S().Error(err)
		session.AddFlash("获取文章属性出错", "Error")
		session.Save(r, w)
	}

	response.CategoryList, err = H.Repo.GetCategoriesNoParentID(10, 0) // 默认取10条
	if err != nil {
		zap.S().Error(err)
		session.AddFlash("获取文章分类出错", "Error")
		session.Save(r, w)
	}

	response.PageIndex = pageIndex
	response.PageSize = pageSize
	td.Data["response"] = response

	// fmt.Println(">>> Posts: ", len(response.Posts), response.Posts)
	// fmt.Println(">>> AttributeList: ", len(response.AttributeList), response.AttributeList)
	// fmt.Println(">>> CategoryList: ", len(response.CategoryList), response.CategoryList)

	// if len(response.Posts) > 0 {
	// 	fmt.Println(">>> Posts 0 C: ", response.Posts[0].Categories)
	// 	fmt.Println(">>> Posts 0 A: ", response.Posts[0].Attributes)
	// 	fmt.Println(">>> Posts 0 Title: ", response.Posts[0].Attributes)
	// }

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

func ShowDetails(w http.ResponseWriter, r *http.Request) {
	td := models.NewTemplateData()
	td.Menubar.Desktop = true
	session := GetSession(w, r)

	vars := mux.Vars(r)
	postID := vars["id"]
	id, err := strconv.Atoi(postID)
	if err != nil {
		session.AddFlash("获取文章详情出错", "Error")
		SaveSession(session, w, r)
		return
	}

	post, err := H.Repo.GetPost(id)
	if err != nil {
		session.AddFlash("获取文章详情出错", "Error")
		SaveSession(session, w, r)
		return
	}
	td.Data["post"] = post
	H.Template.Render(w, r, "details.page.tmpl", &td)

}
