package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"lightsaid.com/weblogs/internal/models"
	"lightsaid.com/weblogs/internal/service"
	"lightsaid.com/weblogs/internal/validator"
)

func ShowAdminCategories(w http.ResponseWriter, r *http.Request) {
	td := models.NewTemplateData()
	td.Menubar.Categories = true

	session := GetSession(w, r)
	vars := mux.Vars(r)
	parent_id := vars["parent_id"]
	id, err := strconv.Atoi(parent_id)
	if err != nil {
		id = 0
	}
	cates, err := H.Repo.GetCategories(id)
	if err != nil {
		zap.S().Error(err)
		session.AddFlash("获取分类列表出错", "Error")
		session.Save(r, w)
		http.Redirect(w, r, fmt.Sprintf("/admin/categories/%d", id), http.StatusSeeOther)
		return
	}

	// 获取 cates 的父类
	parent_cate, _ := H.Repo.GetCategoriesById(id)

	td.Data["categories"] = cates
	td.Data["parent_id"] = parent_id
	td.Data["parent"] = parent_cate
	H.Template.Render(w, r, "categories.page.tmpl", &td)
}

func CreateCategories(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		ServerError(w, err)
		return
	}

	session := GetSession(w, r)

	var req = models.Category{}
	req.Name = r.Form.Get("name")
	parent_id_str := r.Form.Get("parent_id")
	parent_id, err := strconv.Atoi(parent_id_str)
	if err != nil {
		parent_id = 0
	}
	req.ParentID = &parent_id
	thumb := r.Form.Get("thumb")
	req.Thumb = &thumb

	userinfo := r.Context().Value("userinfo")
	info, ok := userinfo.(service.SessionUser)
	if !ok || info.UserID <= 0 {
		zap.S().Info("获取不到ID", info)
		http.Redirect(w, r, "/admin/logout", http.StatusSeeOther)
		return
	}
	req.UserID = info.UserID

	formValidator := validator.NewFormValidator(r.PostForm)
	formValidator.Required("name")
	formValidator.MaxLength("name", 32)
	if !formValidator.Valid() {
		zap.S().Info("验证不通过", formValidator.Errors)
		session.AddFlash(formValidator.Errors, "Error")
		session.Save(r, w)
		http.Redirect(w, r, fmt.Sprintf("/admin/categories/%d", parent_id), http.StatusSeeOther)
		return
	}

	// if *req.ParentID > 0 {
	// 	// NOTE: 不是顶级父类, 因此需要查找父类，更改父类IfParent=1
	// 交由 InsertCategories 实现

	// }

	// 添加资
	_, err = H.Repo.InsertCategories(&req, parent_id)
	if err != nil {
		zap.S().Error(err)
		if err != nil && strings.Contains(strings.ToLower(err.Error()), "unique") {
			zap.S().Error("分类名字已被使用", err)
			session.AddFlash("分类名字已被使用", "Error")
			session.Save(r, w)
			http.Redirect(w, r, fmt.Sprintf("/admin/categories/%d", parent_id), http.StatusSeeOther)
			return
		}
		session.AddFlash("添加出错", "Error")
		session.Save(r, w)
		http.Redirect(w, r, fmt.Sprintf("/admin/categories/%d", parent_id), http.StatusSeeOther)
		return
	}

	session.AddFlash("添加成功", "Success")
	session.Save(r, w)
	http.Redirect(w, r, fmt.Sprintf("/admin/categories/%d", parent_id), http.StatusSeeOther)
}

func UpdateCategories(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		ServerError(w, err)
		return
	}

	session := GetSession(w, r)

	var req = models.Category{}
	req.Name = r.Form.Get("name")
	parent_id_str := r.Form.Get("parent_id")
	parent_id, err := strconv.Atoi(parent_id_str)
	if err != nil {
		parent_id = 0
	}
	req.ParentID = &parent_id
	thumb := r.Form.Get("thumb")
	req.Thumb = &thumb

	if *req.ParentID > 0 {
		req.IfParent = 1
	}

	vars := mux.Vars(r)
	cid := vars["id"]
	id, err := strconv.Atoi(cid)
	if err != nil {
		ServerError(w, err)
		return
	}
	req.ID = id

	userinfo := r.Context().Value("userinfo")
	info, ok := userinfo.(service.SessionUser)
	if !ok || info.UserID <= 0 {
		zap.S().Info("获取不到ID", info)
		http.Redirect(w, r, "/admin/logout", http.StatusSeeOther)
		return
	}
	req.UserID = info.UserID

	formValidator := validator.NewFormValidator(r.PostForm)
	formValidator.Required("name")
	formValidator.MaxLength("name", 32)
	if !formValidator.Valid() {
		zap.S().Info("验证不通过", formValidator.Errors)
		session.AddFlash(formValidator.Errors, "Error")
		session.Save(r, w)
		http.Redirect(w, r, fmt.Sprintf("/admin/categories/%d", parent_id), http.StatusSeeOther)
		return
	}

	err = H.Repo.UpdateCategories(&req)
	if err != nil {
		zap.S().Error(err)
		session.AddFlash("修改出错", "Error")
		session.Save(r, w)
		http.Redirect(w, r, fmt.Sprintf("/admin/categories/%d", parent_id), http.StatusSeeOther)
		return
	}

	session.AddFlash("修改成功", "Success")
	session.Save(r, w)
	http.Redirect(w, r, fmt.Sprintf("/admin/categories/%d", parent_id), http.StatusSeeOther)

}

func DeleteCategories(w http.ResponseWriter, r *http.Request) {
	session := GetSession(w, r)
	var vars = mux.Vars(r)

	p_idStr := vars["parent_id"]
	parent_id, err := strconv.Atoi(p_idStr)
	if err != nil {
		ServerError(w, err)
		return
	}

	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		session.AddFlash("id不正确", "Error")
		session.Save(r, w)
		http.Redirect(w, r, fmt.Sprintf("/admin/categories/%d", parent_id), http.StatusSeeOther)
		return
	}

	// NOTE: 检测是否有子类，有则不允许删除
	cates, err := H.Repo.GetCategories(id)
	if err != nil {
		zap.S().Error(err)
		session.AddFlash("删除出错", "Error")
		session.Save(r, w)
		http.Redirect(w, r, fmt.Sprintf("/admin/categories/%d", parent_id), http.StatusSeeOther)
		return
	}

	if len(cates) > 0 {
		zap.S().Error(err)
		session.AddFlash("存在子类不允许删除，请先删除子类", "Error")
		session.Save(r, w)
		http.Redirect(w, r, fmt.Sprintf("/admin/categories/%d", parent_id), http.StatusSeeOther)
		return
	}

	err = H.Repo.DeleteCategories(id)
	if err != nil {
		zap.S().Error(err)
		session.AddFlash("删除出错", "Error")
		session.Save(r, w)
		http.Redirect(w, r, fmt.Sprintf("/admin/categories/%d", parent_id), http.StatusSeeOther)
		return
	}

	session.AddFlash("删除成功", "Success")
	session.Save(r, w)
	http.Redirect(w, r, fmt.Sprintf("/admin/categories/%d", parent_id), http.StatusSeeOther)
}
