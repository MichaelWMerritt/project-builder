package controller

import (
	"github.com/michaelwmerritt/project-builder/model"
	"net/http"
	//"encoding/json"
	//"github.com/gorilla/mux"
	//"fmt"
)

func CreateBuilderRoutes() []model.Route {
	return []model.Route{
		{
			"Build",
			"POST",
			"/builds",
			Build,
		},
		{
			"GetBuildStatus",
			"GET",
			"/builds/{buildId}",
			GetBuildStatus,
		},
		{
			"CancelBuild",
			"PATCH",
			"/builds/{buildId}",
			CancelBuild,
		},
		{
			"DeleteBuild",
			"DELETE",
			"/builds/{buildId}",
			DeleteBuild,
		},
		{
			"GetBuildResult",
			"GET",
			"/builds/{buildId}/result",
			GetBuildResult,
		},
	}
}

func Build(w http.ResponseWriter, r *http.Request) {
	//TODO: get list of modules and build project, then store results in mongodb and update status
}

func GetBuildStatus(w http.ResponseWriter, r *http.Request) {

}

func CancelBuild(w http.ResponseWriter, r *http.Request) {

}

func DeleteBuild(w http.ResponseWriter, r *http.Request) {

}

func GetBuildResult(w http.ResponseWriter, r *http.Request) {

}
