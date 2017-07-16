package main

import (
	"log"
	"net/http"

	"github.com/michaelwmerritt/project-builder/controller"
	"github.com/michaelwmerritt/project-builder/model"

	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"os"
	"github.com/michaelwmerritt/project-builder/server"
)

const  (
	serverAddress = ":8080"
)

func main() {
	log.Fatal(http.ListenAndServe(serverAddress, createHandlers()))
}

func createHandlers() http.Handler {
	return handlers.CORS()(
		handlers.LoggingHandler(
			os.Stdout, handlers.RecoveryHandler()(
				createRouter())))
}

func createRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range initializeRoutes() {
		var handler http.Handler

		handler = route.HandlerFunc
		loggerAdapter := server.CreateLoggerAdapter(route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(server.Adapt(handler, loggerAdapter))
	}

	return router
}

func initializeRoutes() []model.Route {
	routes := controller.CreateIndexRoutes()
	routes = append(routes, controller.CreateReleaseRoutes()...)
	routes = append(routes, controller.CreateModuleRoutes()...)
	routes = append(routes, controller.CreateBuilderRoutes()...)
	return routes
}
