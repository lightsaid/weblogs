package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/csrf"
	"github.com/gorilla/sessions"
	"go.uber.org/zap"
	"lightsaid.com/weblogs/internal/models"
	"lightsaid.com/weblogs/internal/security"
	"lightsaid.com/weblogs/internal/service"
	"lightsaid.com/weblogs/internal/validator"
)

var session *sessions.Session

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("Hello, %s", r.RemoteAddr)))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var req service.CreateUserRequest
	err := H.readJSON(w, r, &req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("请求参数错误."))
		return
	}
	user, err := H.Repo.InsertUser(req.Email, req.Username, req.Password, req.Avatar)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("请求发生错误."))
		return
	}

	w.WriteHeader(http.StatusOK)
	b, _ := json.MarshalIndent(&user, "", "\t")
	w.Write(b)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println(H.DB)
	w.Write([]byte(fmt.Sprintf("GetUser, %s", r.RemoteAddr)))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

}

//  LoginWithRegister 登录和注册，
func LoginWithRegister(w http.ResponseWriter, r *http.Request) {
	var req = service.LoginWithRegisterRequest{}
	var td = models.NewTemplateData()
	w.Header().Set("X-CSRF-Token", csrf.Token(r))

	err := H.readJSON(w, r, &req)
	if err != nil {
		zap.S().Error(err)
		td.Data["user"] = req
		td.Error = "解析参数错误"
		H.Template.Render(w, r, "login.page.tmpl", &td)
		return
	}
	td.Data["user"] = req
	// TODO: 存储 Session
	jvalid, err := validator.NewJsonValidator(req)
	if err != nil {
		zap.S().Error(err)
		td.Error = err.Error()
		H.Template.Render(w, r, "login.page.tmpl", &td)
		return
	}
	// 校验 formType
	jvalid.Check(jvalid.Includes(req.FormType, "Register", "Login"), "formType", "formType 选项必须是Register或Login")
	if !jvalid.Valid() {
		zap.S().Info("验证不通过：", jvalid.Errors)
		td.JsonValidator = jvalid
		H.Template.Render(w, r, "login.page.tmpl", &td)
		return
	}
	if req.FormType == "Register" {
		Register(w, r, req, jvalid, &td)
		return
	}
	Login(w, r, req, jvalid, &td)
}

// 注册
func Register(w http.ResponseWriter, r *http.Request, req service.LoginWithRegisterRequest,
	jvalid *validator.JsonValidator, td *models.TemplateData) {

	// 验证参数
	jvalid.Required("email", "password", "ack_password")

	if req.AckPassword != req.Password {
		jvalid.AddError("ack_password", "两次密码不一致")
	}

	jvalid.MinLength("password", 6)
	jvalid.MaxLength("password", 16)
	jvalid.IsEmail("email")

	td.JsonValidator = jvalid

	if !jvalid.Valid() {
		zap.S().Info("验证不通过：", jvalid.Errors)
		H.Template.Render(w, r, "login.page.tmpl", td)
		return
	}

	// 创建请求参数
	hassPwd, err := security.Hash(req.Password)
	if err != nil {
		zap.S().Error("hash 密码错误", err)
		td.Error = "注册失败"
		H.Template.Render(w, r, "login.page.tmpl", td)
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
		H.Template.Render(w, r, "login.page.tmpl", td)
		return
	}

	if err != nil {
		zap.S().Error("注册服务故障：", err)
		td.Error = "注册服务故障"
		H.Template.Render(w, r, "login.page.tmpl", td)
		return
	}

	zap.S().Info("创建用户成功：", user)

	// 往Session添加Flsh消息
	req.FormType = "Login"
	req.Password = ""
	req.AckPassword = ""
	td.Data["user"] = req
	td.Success = "注册成功，转到登录"
	H.Template.Render(w, r, "login.page.tmpl", td)
}

func Login(w http.ResponseWriter, r *http.Request, req service.LoginWithRegisterRequest,
	jvalid *validator.JsonValidator, td *models.TemplateData) {

	// 验证参数
	jvalid.Required("email", "password")

	// 传递验证，page可以使用
	td.JsonValidator = jvalid

	if !jvalid.Valid() {
		zap.S().Info("验证不通过：", jvalid.Errors)
		H.Template.Render(w, r, "login.page.tmpl", td)
		return
	}

	user, err := H.Repo.GetUserByEmial(req.Email)
	if err != nil {
		td.Error = "服务内部错误"
		if err == sql.ErrNoRows {
			td.Error = "查询不到用户"
		}
		zap.S().Info(">>> GetUserByEmial error：", err)
		H.Template.Render(w, r, "login.page.tmpl", td)
		return
	}

	err = security.VerifyPassword(user.Password, req.Password)
	if err != nil {
		zap.S().Info(">>> VerifyPassword error：", err)
		td.Error = "密码不正确"
		H.Template.Render(w, r, "login.page.tmpl", td)
		return
	}

	// 登录成功  TODO: 区分管理员和非管理员
	if session, err = H.CookieStore.Get(r, os.Getenv("SESSION")); err != nil {
		H.errorResponse(w)
		return
	}
	// session.Values["Flash"] = "登录成功"
	session.AddFlash("登录成功！", "Success")
	if err = session.Save(r, w); err != nil {
		H.errorResponse(w)
		return
	}

	// NOTE: 此处不合理，不应该是后台重定向，而是前端做跳转
	if user.IfAdmin == 1 {
		http.Redirect(w, r, "/admin/index", http.StatusSeeOther)
	} else {

	}
	// resp := service.LoginWithRegisterResponse{
	// 	IfAdmin: user.IfAdmin,
	// }
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// _ = json.NewEncoder(w).Encode(&resp)

}
