package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"lightsaid.com/weblogs/internal/models"
	"lightsaid.com/weblogs/internal/service"
	"lightsaid.com/weblogs/internal/validator"
)

func ShowAdminAttrs(w http.ResponseWriter, r *http.Request) {
	td := models.NewTemplateData()
	td.Menubar.AttributeList = true

	session := GetSession(w, r)

	attrs, err := H.Repo.GetAttributes()
	if err != nil {
		zap.S().Error(err)
		session.AddFlash("获取属性列表出错", "Error")
		session.Save(r, w)
		http.Redirect(w, r, "/admin/attrs", http.StatusSeeOther)
		return
	}
	td.Data["attrs"] = attrs
	H.Template.Render(w, r, "attrs.page.tmpl", &td)
}

func CreateAttribute(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		ServerError(w, err)
		return
	}

	session := GetSession(w, r)

	var req = models.Attribute{}

	req.Kind = r.Form.Get("kind")
	req.Name = r.Form.Get("name")
	userinfo := r.Context().Value("userinfo")
	info, ok := userinfo.(service.SessionUser)
	if !ok || info.UserID <= 0 {
		zap.S().Info("获取不到ID", info)
		http.Redirect(w, r, "/admin/logout", http.StatusSeeOther)
		return
	}
	req.UserID = info.UserID

	formValidator := validator.NewFormValidator(r.PostForm)
	formValidator.Required("kind", "name")
	if !formValidator.Valid() {
		zap.S().Info("验证不通过", formValidator.Errors)
		session.AddFlash(formValidator.Errors, "Error")
		session.Save(r, w)
		http.Redirect(w, r, "/admin/attrs", http.StatusSeeOther)
		return
	}

	_, err := H.Repo.InsertAttrs(&req)
	if err != nil {
		zap.S().Error(err)
		session.AddFlash("添加出错", "Error")
		session.Save(r, w)
		http.Redirect(w, r, "/admin/attrs", http.StatusSeeOther)
		return
	}

	session.AddFlash("添加成功", "Success")
	session.Save(r, w)
	http.Redirect(w, r, "/admin/attrs", http.StatusSeeOther)
}

func UpdateAttribute(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		ServerError(w, err)
		return
	}

	var vars = mux.Vars(r)

	session := GetSession(w, r)

	var req = models.Attribute{}
	id := vars["id"]
	idint, err := strconv.Atoi(id)
	if err != nil {
		zap.S().Info("ID 转换失败: ", id)
		session.AddFlash("ID转换失败", "Error")
		session.Save(r, w)
		http.Redirect(w, r, "/admin/attrs", http.StatusSeeOther)
		return
	}
	req.ID = idint

	req.Kind = r.Form.Get("kind")
	req.Name = r.Form.Get("name")
	userinfo := r.Context().Value("userinfo")
	info, ok := userinfo.(service.SessionUser)
	if !ok || info.UserID <= 0 {
		zap.S().Info("获取不到ID", info)
		session.AddFlash("登录失效", "Error")
		session.Save(r, w)
		http.Redirect(w, r, "/admin/logout", http.StatusSeeOther)
		return
	}
	req.UserID = info.UserID

	formValidator := validator.NewFormValidator(r.PostForm)
	formValidator.Required("kind", "name")
	if !formValidator.Valid() {
		zap.S().Info("验证不通过", formValidator.Errors)
		session.AddFlash(formValidator.Errors, "Error")
		session.Save(r, w)
		http.Redirect(w, r, "/admin/attrs", http.StatusSeeOther)
		return
	}

	err = H.Repo.UpdateAttributes(&req)
	if err != nil {
		zap.S().Error(err)
		session.AddFlash("修改出错", "Error")
		session.Save(r, w)
		http.Redirect(w, r, "/admin/attrs", http.StatusSeeOther)
		return
	}

	session.AddFlash("修改成功", "Success")
	session.Save(r, w)
	http.Redirect(w, r, "/admin/attrs", http.StatusSeeOther)
}

func DeleteAttribute(w http.ResponseWriter, r *http.Request) {
	session := GetSession(w, r)
	var vars = mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		session.AddFlash("id不正确", "Error")
		session.Save(r, w)
		http.Redirect(w, r, "/admin/attrs", http.StatusSeeOther)
		return
	}
	err = H.Repo.DeleteAttribute(id)
	if err != nil {
		zap.S().Error(err)
		session.AddFlash("删除出错", "Error")
		session.Save(r, w)
		http.Redirect(w, r, "/admin/attrs", http.StatusSeeOther)
		return
	}

	session.AddFlash("删除成功", "Success")
	session.Save(r, w)
	http.Redirect(w, r, "/admin/attrs", http.StatusSeeOther)
}
