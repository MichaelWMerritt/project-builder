package main

import (
	"log"
	"net/http"

	"github.com/michaelwmerritt/project-builder/controller"
	"github.com/michaelwmerritt/project-builder/model"

	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"os"
)

var (
	port = ":8080"
)

func main() {
	log.Fatal(http.ListenAndServe(port, handlers.CORS()(handlers.LoggingHandler(os.Stdout, handlers.RecoveryHandler()(NewRouter())))))
}

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	for _, route := range InitializeRoutes() {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = CreateLogger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func InitializeRoutes() []model.Route {
	routes := controller.CreateIndexRoutes()
	routes = append(routes, controller.CreateReleaseRoutes()...)
	routes = append(routes, controller.CreateModuleRoutes()...)
	routes = append(routes, controller.CreateBuilderRoutes()...)
	return routes
}
