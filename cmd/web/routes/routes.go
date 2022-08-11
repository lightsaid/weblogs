package routes

import (
	"github.com/gorilla/mux"
	"lightsaid.com/weblogs/cmd/web/middleware"
)

func New() *mux.Router{
	r := mux.NewRouter().StrictSlash(true)
	// TODO:
	// r.Host("www.example.com")

	r.Use(middleware.LogMiddlewate)
	return setupRoutes(r)
}

func load() []Router {
	routes := userRoutes
	routes = append(routes, postRoutes...)
	routes = append(routes, cateRoutes...)
	routes = append(routes, attrRoutes...)
	return routes
}

func setupRoutes(r *mux.Router) *mux.Router{
	for _, route := range load(){
		r.HandleFunc(route.Path, route.Handler).Methods(route.Method)
	}
	return r
}