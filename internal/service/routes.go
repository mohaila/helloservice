package service

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route defines a signle web route
type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write([]byte(`{"status" : "ok"}`))
}

// Routes is an array of web routes
type Routes []Route

var (
	routes = Routes{
		Route{
			Name:        "Status",
			Method:      "GET",
			Path:        "/status",
			HandlerFunc: statusHandler,
		},
	}
)

// NewRouter returns a router
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.Name(route.Name).
			Path(route.Path).
			Methods(route.Method).
			Handler(route.HandlerFunc)
	}

	return router
}
