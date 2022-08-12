package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"lightsaid.com/weblogs/internal/service"
)

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
