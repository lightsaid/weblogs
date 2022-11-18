package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/ory/nosurf"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"lightsaid.com/weblogs/forms"
	"lightsaid.com/weblogs/utils"
)

// JSONRegister Get: 提供CSRFToken, Post: 处理内部注册，不公开 入参和出参是JSON数据
func (_this *Controller) JSONRegister(w http.ResponseWriter, r *http.Request) {
	var err error

	// Get 请求响应
	if r.Method == http.MethodGet {
		fmt.Fprintln(w, nosurf.Token(r))
		return
	}

	// POST请求 响应
	defer func() {
		var data = make(map[string]interface{}, 2)
		data["ok"] = true
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err != nil {
			data["ok"] = false
			data["error"] = err.Error()
			log.Error(err)
		}
		json.NewEncoder(w).Encode(data)
	}()

	var req struct {
		Email     string `json:"email"`
		Password  string `json:"password"`
		CSRFToken string `json:"csrf_token"`
	}

	dec := json.NewDecoder(r.Body)
	err = dec.Decode(&req)
	if err != nil {
		err = fmt.Errorf("入参不正确: %w", err)
		return
	}

	if !nosurf.VerifyToken(nosurf.Token(r), req.CSRFToken) {
		err = fmt.Errorf("CSRF token incorrect")
		return
	}

	req.Email = strings.TrimSpace(req.Email)
	req.Password = strings.TrimSpace(req.Password)
	if req.Email == "" || req.Password == "" {
		err = fmt.Errorf("邮箱和密码必填")
		return
	}

	req.Password, err = utils.GenHashedPwsd(req.Password)
	if err != nil {
		err = fmt.Errorf("密码加密错误 password: %v, error: %w", req.Password, err)
		return
	}

	username := _this.GenDefaultUserName()
	avatar := _this.GenDefaultAvatar()

	err = _this.Models.Users.Insert(req.Email, username, req.Password, avatar)
}

// PostLogin 登录
func (_this *Controller) PostLogin(w http.ResponseWriter, r *http.Request) {
	var err error
	log.Info("登录ing～")
	defer func() {
		if err != nil {
			log.Error(err)
		}
	}()

	r.ParseForm()
	v := forms.New(r.PostForm)

	email := r.Form.Get("email")
	password := r.Form.Get("password")

	email = strings.TrimSpace(email)
	password = strings.TrimSpace(password)

	v.Required("email", "password")
	if !v.Valid() {
		err = fmt.Errorf("入参不正确: email: %v | password: %v", email, password)
		return
	}

	user, err := _this.Models.Users.GetByEmail(email)
	if err != nil {
		return
	}
	err = utils.VerifyPassword(password, user.Password)
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			err = fmt.Errorf("密码和邮箱不匹配 %w", err)
		} else if errors.Is(err, bcrypt.ErrHashTooShort) {
			err = fmt.Errorf("密码过短: %w", err)
		} else {
			err = fmt.Errorf("密码未知错误 %w", err)
		}
		return
	}

	_this.Session.Put(r.Context(), "user_id", user.ID)
	_this.Session.Put(r.Context(), "success", "登录成功")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
