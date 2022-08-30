package routes

import (
	"net/http"
	"os"

	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"lightsaid.com/weblogs/cmd/web/middleware"
)

func New() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	// TODO:
	// r.Host("www.example.com")

	csrfMiddleware := csrf.Protect([]byte(os.Getenv("CSRF_SECRET")), csrf.Secure(false), csrf.Path("/"))

	// r.Use(middleware.EditorRedirect)

	// 静态资源访问
	fileHandler := http.StripPrefix("/static/", http.FileServer(http.Dir(os.Getenv("STATIC_PATH"))))
	r.PathPrefix("/static/").Handler(fileHandler)

	r.Use(middleware.LogMiddlewate)
	r.Use(csrfMiddleware)
	r.Use(middleware.RateLimit(30))

	return setupRoutes(r)
}

func load() []Router {
	routes := userRoutes
	routes = append(routes, postRoutes...)
	routes = append(routes, cateRoutes...)
	routes = append(routes, attrRoutes...)
	routes = append(routes, blogRoutes...)
	routes = append(routes, adminPageRoutes...)
	return routes
}

func setupRoutes(r *mux.Router) *mux.Router {
	for _, route := range load() {
		if route.AuthRequired {
			r.HandleFunc(route.Path, middleware.MultipleMiddleware(route.Handler, middleware.AuthMiddleware)).Methods(route.Method)
		} else {
			r.HandleFunc(route.Path, middleware.MultipleMiddleware(route.Handler, middleware.SettingUserinfo)).Methods(route.Method)
		}
	}
	return r
}
