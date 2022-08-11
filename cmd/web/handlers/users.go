package handlers

import (
	"fmt"
	"net/http"
)

func (app *AppHandler) GetUsers(w http.ResponseWriter, r *http.Request){
	w.Write([]byte(fmt.Sprintf("Hello, %s", r.RemoteAddr)))
}

func (app *AppHandler) CreateUser(w http.ResponseWriter, r *http.Request){
	
}

func (app *AppHandler) GetUser(w http.ResponseWriter, r *http.Request){
	
}

func (app *AppHandler) UpdateUser(w http.ResponseWriter, r *http.Request){
	
}

func (app *AppHandler) DeleteUser(w http.ResponseWriter, r *http.Request){
	
}