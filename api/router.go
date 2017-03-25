package main

import (
	"net/http"

	"github.com/timkellogg/todo/api/config"
	"github.com/timkellogg/todo/api/controllers"

	"github.com/gorilla/mux"
)

// Route is a REST endpoint that maps common methods
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes is a collection of Route
type Routes []Route

var routes = Routes{
	Route{"Index", "GET", "/todos", controllers.TodoIndex},
	Route{"Show", "GET", "/todos/{id}", controllers.TodoShow},
	Route{"Create", "POST", "/todos", controllers.TodoCreate},
	Route{"Update", "PATCH", "/todos/{id}", controllers.TodoUpdate},
	Route{"Destroy", "DELETE", "/todos/{id}", controllers.TodoDestroy},
}

// NewRouter establishes the root application router
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = config.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}

	return router
}
