package controller

import (
	"net/http"

	log "github.com/sirupsen/logrus"
	"lightsaid.com/weblogs/forms"
)

// CreateTag 创建tag
func (_this *Controller) CreateTag(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	form := forms.New(r.PostForm)
	tagName := form.Get("tag")

	form.Required("tag")
	if !form.Valid() {
		log.Error("create tag 验证不通过: tag 参数必填")
		// TODO:
		return
	}
	userID := _this.Session.GetInt(r.Context(), "user_id")
	if userID <= 0 {
		log.Error("session 中 userid <= 0")
		return
	}
	err := _this.Models.Tags.Insert(userID, tagName)
	if err != nil {
		log.Error(err)
		return
	}

	http.Redirect(w, r, "/tag", http.StatusSeeOther)
	log.Info("create tag success")
}
