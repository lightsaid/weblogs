package handlers

import (
	"database/sql"
	"errors"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"go.uber.org/zap"
	"lightsaid.com/weblogs/internal/models"
	"lightsaid.com/weblogs/internal/repository/dbrepo"
	"lightsaid.com/weblogs/internal/security"
	"lightsaid.com/weblogs/internal/service"
	"lightsaid.com/weblogs/internal/validator"
)

var session *sessions.Session

func ShowAdminRegister(w http.ResponseWriter, r *http.Request) {
	var td = models.NewTemplateData()
	td.FormValidator = validator.NewFormValidator(nil)
	H.Template.Render(w, r, "register.page.tmpl", &td)
}

func PostAdminRegister(w http.ResponseWriter, r *http.Request) {
	// 解析 Form
	err := r.ParseForm()
	if err != nil {
		ServerError(w, err)
		return
	}

	var td = models.NewTemplateData()
	var req = service.RegisterRequest{}

	req.Email = r.Form.Get("email")
	req.Password = r.Form.Get("password")
	req.AckPassword = r.Form.Get("ack_password")

	formValidator := validator.NewFormValidator(r.PostForm)

	// 验证必填项
	formValidator.Required("email", "password", "ack_password")
	formValidator.IsEmail("email")
	formValidator.MinLength("password", 6)
	formValidator.MaxLength("password", 16)
	if req.AckPassword != req.Password {
		formValidator.Errors.Add("ack_password", "两次密码不一致")
	}

	td.Data["request"] = req

	td.FormValidator = formValidator

	if !formValidator.Valid() {
		H.Template.Render(w, r, "register.page.tmpl", &td)
		return
	}

	// 创建请求参数
	hassPwd, err := security.Hash(req.Password)
	if err != nil {
		zap.S().Error("hash 密码错误", err)
		td.Error = "注册失败"
		H.Template.Render(w, r, "register.page.tmpl", &td)
		return
	}

	arg := service.CreateUserRequest{
		Email:    req.Email,
		Password: string(hassPwd),
		Username: service.CreateDefaultUsername(),
		Avatar:   service.CreateDefaultAvatar(),
	}

	// 添加到数据库
	user, err := H.Repo.Register(arg)
	if err != nil && strings.Contains(strings.ToLower(err.Error()), "unique") {
		zap.S().Error("邮箱已被使用", err)
		td.Error = "邮箱已被使用"
		H.Template.Render(w, r, "register.page.tmpl", &td)
		return
	}
	// 其他类型错误
	if err != nil {
		ServerError(w, err)
		return
	}

	zap.S().Info("创建用户成功：", user)

	if session, err = H.CookieStore.Get(r, os.Getenv("SESSION")); err != nil {
		ServerError(w, err)
		return
	}

	// 往Session添加Flsh消息
	session.AddFlash("注册成功，跳转到登录页!", "Success")
	if err = session.Save(r, w); err != nil {
		ServerError(w, err)
		return
	}

	http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
}

func ShowAdminLogin(w http.ResponseWriter, r *http.Request) {
	data := models.NewTemplateData()
	data.FormValidator = validator.NewFormValidator(nil)
	H.Template.Render(w, r, "login.page.tmpl", &data)
}

func PostLogin(w http.ResponseWriter, r *http.Request) {
	// 解析 Form
	err := r.ParseForm()
	if err != nil {
		ServerError(w, err)
		return
	}

	var formValidator = validator.NewFormValidator(r.PostForm)
	var td = models.NewTemplateData()
	var req = service.LoginRequest{}

	req.Email = r.Form.Get("email")
	req.Password = r.Form.Get("password")
	if len(r.Form.Get("remember")) > 0 {
		req.Remember = true
	}

	// 验证必填项
	formValidator.Required("email", "password")
	formValidator.IsEmail("email")
	formValidator.MinLength("password", 6)
	formValidator.MaxLength("password", 16)

	td.Data["request"] = req

	td.FormValidator = formValidator

	if !formValidator.Valid() {
		H.Template.Render(w, r, "login.page.tmpl", &td)
		return
	}

	user, err := H.Repo.GetUserByEmial(req.Email)
	if err != nil {
		td.Error = "服务内部错误"
		if err == sql.ErrNoRows {
			td.Error = "查询不到用户"
		}
		zap.S().Info(">>> GetUserByEmial error: ", err)
		H.Template.Render(w, r, "login.page.tmpl", &td)
		return
	}

	err = security.VerifyPassword(user.Password, req.Password)
	if err != nil {
		zap.S().Info(">>> VerifyPassword error: ", err)
		td.Error = "密码不正确"
		H.Template.Render(w, r, "login.page.tmpl", &td)
		return
	}

	if user.Active < 0 {
		td.Error = "查询不到用户或用户已删除"
		zap.S().Info(">>> GetUserByEmial error: ", err)
		H.Template.Render(w, r, "login.page.tmpl", &td)
		return
	}

	// 登录成功
	if session, err = H.CookieStore.Get(r, os.Getenv("SESSION")); err != nil {
		H.errorResponse(w)
		return
	}

	userinfo := service.SessionUser{
		UserID:   user.ID,
		Username: user.Username,
		Avatar:   *user.Avatar,
	}
	session.Values["userinfo"] = userinfo
	session.AddFlash("登录成功！", "Success")
	if err = session.Save(r, w); err != nil {
		ServerError(w, err)
		return
	}

	// TODO: 记住我
	if req.Remember {

	}

	// NOTE: 此处不合理，不应该是后台重定向，而是前端做跳转
	if user.IfAdmin == 1 {
		http.Redirect(w, r, "/admin/index", http.StatusSeeOther)
	} else {
		// TODO:
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	var session *sessions.Session
	var err error
	if session, err = H.CookieStore.Get(r, os.Getenv("SESSION")); err != nil {
		H.errorResponse(w)
		return
	}
	session.Values = make(map[interface{}]interface{})
	session.Save(r, w)
	http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
}

func ShowAdminIndex(w http.ResponseWriter, r *http.Request) {
	data := models.NewTemplateData()
	data.Menubar.Home = true
	H.Template.Render(w, r, "dashboard.page.tmpl", &data)
}

func ShowAdminUsers(w http.ResponseWriter, r *http.Request) {
	var data = models.NewTemplateData()
	users, err := H.Repo.GetUsers()
	if err != nil {
		data.Error = err.Error()
	}
	// 改用 template.FuncMap 方式
	// for i, u := range users {
	// 	avatar := *u.Avatar
	// 	if len(avatar) > 0 && avatar[0] == '.' {
	// 		prefix := os.Getenv("ASSETS_PREFIX")
	// 		url := fmt.Sprintf("%s%s", prefix, avatar[1:])
	// 		users[i].Avatar = &url
	// 	}
	// }
	data.Menubar.UserList = true
	data.Data["users"] = users
	H.Template.Render(w, r, "users.page.tmpl", &data)
}

// func GetUsers(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte(fmt.Sprintf("Hello, %s", r.RemoteAddr)))
// }

// func CreateUser(w http.ResponseWriter, r *http.Request) {
// 	var req service.CreateUserRequest
// 	err := H.readJSON(w, r, &req)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		w.Write([]byte("请求参数错误."))
// 		return
// 	}
// 	user, err := H.Repo.InsertUser(req.Email, req.Username, req.Password, req.Avatar)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte("请求发生错误."))
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	b, _ := json.MarshalIndent(&user, "", "\t")
// 	w.Write(b)
// }

// func GetUser(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte(fmt.Sprintf("GetUser, %s", r.RemoteAddr)))
// }

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	session, err := H.CookieStore.Get(r, os.Getenv("SESSION"))
	if err != nil {
		H.errorResponse(w)
		return
	}

	req := models.User{}
	id := vars["id"]
	req.ID, err = strconv.Atoi(id)
	if err != nil {
		zap.S().Error("用户id转换失败, user_id=", id)
		session.AddFlash("用户ID不存在", "Error")
		SaveSession(session, w, r)
		http.Redirect(w, r, "/admin/users", http.StatusSeeOther)
		return
	}

	req.Email = r.FormValue("email")
	avatar := r.FormValue("avatar")
	req.Avatar = &avatar
	req.Username = r.FormValue("username")
	if len(r.FormValue("if_admin")) > 0 {
		req.IfAdmin = 1
	}

	form := validator.NewFormValidator(r.PostForm)
	form.Required("email", "username")

	if !form.Valid() {
		zap.S().Info("验证不通过：", form.Errors)
		session.AddFlash(form.Errors.Get("username"), "Error")
		SaveSession(session, w, r)
		http.Redirect(w, r, "/admin/users", http.StatusSeeOther)
		return
	}

	fileUrl, err := UploadFile(r, os.Getenv("AVATAR_PATH"))

	// 检查是否上传文件， 忽略 err == http.ErrMissingFile 的处理
	if err != nil && !errors.Is(err, http.ErrMissingFile) {
		zap.S().Error("上传文件错误", err)
		ServerError(w, err)
		return
	}

	// 使用上传头像覆盖原有头像
	if len(fileUrl) > 0 {
		req.Avatar = &fileUrl
	}
	err = H.Repo.UpdateUser(req)
	if err != nil {
		zap.S().Error("更新user数据错误", err)
		if errors.Is(err, dbrepo.ErrUpdateNoRows) {
			session.AddFlash("用户ID不存在", "Error")
			SaveSession(session, w, r)
			http.Redirect(w, r, "/admin/users", http.StatusSeeOther)
			return
		}
		session.AddFlash("更新操作出错", "Error")
		SaveSession(session, w, r)
		http.Redirect(w, r, "/admin/users", http.StatusSeeOther)
		return
	}
	session.AddFlash("更新成功", "Success")
	SaveSession(session, w, r)
	http.Redirect(w, r, "/admin/users", http.StatusSeeOther)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	session, err := H.CookieStore.Get(r, os.Getenv("SESSION"))
	if err != nil {
		H.errorResponse(w)
		return
	}

	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		session.AddFlash("ID不存在", "Error")
		SaveSession(session, w, r)
		http.Redirect(w, r, "/admin/users", http.StatusSeeOther)
		return
	}

	err = H.Repo.DeleteUser(id)
	if err != nil {
		zap.S().Error(err)
		if errors.Is(err, dbrepo.ErrUpdateNoRows) {
			session.AddFlash("ID不存在", "Error")
			SaveSession(session, w, r)
			http.Redirect(w, r, "/admin/users", http.StatusSeeOther)
			return
		}
		session.AddFlash("删除出错", "Error")
		SaveSession(session, w, r)
		http.Redirect(w, r, "/admin/users", http.StatusSeeOther)
		return
	}

	session.AddFlash("删除成功", "Success")
	SaveSession(session, w, r)
	http.Redirect(w, r, "/admin/users", http.StatusSeeOther)
}
