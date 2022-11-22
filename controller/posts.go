package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/ory/nosurf"
	log "github.com/sirupsen/logrus"
	"lightsaid.com/weblogs/data"
)

type CreatePostRequest struct {
	Title   string `json:"title"`
	TagIDs  string `jsno:"tag_ids"`
	Content string `json:"content"`
	Active  int    `josn:"active"`
	Token   string `json:"token"`
}

type CreatePostResponse struct {
	OK    bool   `json:"ok"`
	Error string `json:"error"`
}

func (_this *Controller) CreatePost(w http.ResponseWriter, r *http.Request) {
	var response CreatePostResponse
	var req CreatePostRequest
	defer func() {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(response)
	}()

	buf, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error("CreatePost io.ReadAll error: ", err)
		response.OK = false
		response.Error = "参数不正确"
		return
	}
	err = json.Unmarshal(buf, &req)
	if err != nil {
		log.Error("CreatePost json.Unmarshal error: ", err)
		response.OK = false
		response.Error = "无法解释参数"
		return
	}

	if !nosurf.VerifyToken(nosurf.Token(r), req.Token) {
		err = fmt.Errorf("CSRF token incorrect")
		log.Error(err)
		response.OK = false
		response.Error = "token 无效"
		return
	}

	if strings.TrimSpace(req.Title) == "" || strings.TrimSpace(req.Content) == "" {
		log.Error("title,  content 不能为空")
		response.OK = false
		response.Error = "title 和 content 不能为空"
		return
	}

	post := data.Post{
		Title:   req.Title,
		Active:  req.Active,
		Content: req.Content,
	}

	err = _this.Models.Posts.Insert(post)
	if err != nil {
		log.Error("CreatePost _this.Models.Posts.Insert: ", err)
		response.OK = false
		response.Error = "发布错误"
		return
	}
	response.OK = true
	response.Error = "发布成功"
}

func (_this *Controller) PostDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Error(err)
		// TODO:
		return
	}
	post, err := _this.Models.Posts.GetById(id)
	if err != nil {
		log.Error(err)
		// TODO:
		return
	}
	td := new(TemplateData)
	td.DataMap = map[string]interface{}{"post": post}
	_this.Render(w, r, "detail.page.gtpl", td)
}
