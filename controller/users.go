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
	"lightsaid.com/weblogs/data"
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
	_this.Session.RenewToken(r.Context()) // 更新会话，有利于减少 session fixation attacks
	var err error
	var user = new(data.User)
	var td = new(TemplateData)
	defer func() {
		if err != nil {
			log.Error(err)
			_this.Render(w, r, "login.page.gtpl", td)
			return
		}
		_this.Session.Put(r.Context(), "user_id", user.ID)
		_this.Session.Put(r.Context(), "success", "登录成功")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}()

	err = r.ParseForm()
	if err != nil {
		err = fmt.Errorf("r.ParseForm() error: %w", err)
		return
	}

	v := forms.New(r.PostForm)
	td.Form = v

	email := r.PostForm.Get("email")
	password := r.PostForm.Get("password")

	email = strings.TrimSpace(email)
	password = strings.TrimSpace(password)

	td.StringMap = make(map[string]string)
	td.StringMap["email"] = email

	v.Required("email", "password")
	v.MatchesPattern(forms.EmailRX, "email", "邮箱地址不正确")
	v.MinLength("password", 6)
	v.MaxLength("password", 16)
	if !v.Valid() {
		err = fmt.Errorf("入参不正确: email: %v | password: %v, %v", email, password, v.Errors)
		return
	}

	log.Debugf("errors: %v, %v, %v, %d", v.Errors, email, password, len(password))

	user, err = _this.Models.Users.GetByEmail(email)
	if err != nil {
		return
	}

	err = utils.VerifyPassword(password, user.Password)
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			err = fmt.Errorf("邮箱和密码不匹配 %w", err)
		} else if errors.Is(err, bcrypt.ErrHashTooShort) {
			err = fmt.Errorf("密码过短: %w", err)
		} else {
			err = fmt.Errorf("密码未知错误 %w", err)
		}
		v.Errors.Append("password", "邮箱或密码不正确")
		return
	}
}

// Logout 注销
func (_this *Controller) Logout(w http.ResponseWriter, r *http.Request) {
	_this.Session.Destroy(r.Context())
	_this.Session.RenewToken(r.Context())

	http.Redirect(w, r, "/user/login", http.StatusSeeOther)
}

func (_this *Controller) UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "更新")
}
