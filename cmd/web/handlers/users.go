package handlers

import (
	"fmt"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("Hello, %s", r.RemoteAddr)))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {

}

func GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println(AppH.DB)
	w.Write([]byte(fmt.Sprintf("GetUser, %s", r.RemoteAddr)))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

}
