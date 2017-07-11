package controller

import (
	"fmt"
	"github.com/michaelwmerritt/project-builder/model"
	"html"
	"net/http"
)

func CreateIndexRoutes() []model.Route {
	return []model.Route{
		{
			"Index",
			"GET",
			"/",
			Index,
		},
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
