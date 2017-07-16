package controller

import (
	"github.com/michaelwmerritt/project-builder/model"
	"net/http"
	//"encoding/json"
	//"github.com/gorilla/mux"
	//"fmt"
	"github.com/michaelwmerritt/project-builder/server"
)

func CreateBuilderRoutes() []model.Route {
	return []model.Route{
		{
			"Build",
			"POST",
			server.API_ENDPOINT + "/builds",
			Build,
		},
		{
			"GetBuildStatus",
			"GET",
			server.API_ENDPOINT + "/builds/{buildId}",
			GetBuildStatus,
		},
		{
			"CancelBuild",
			"PATCH",
			server.API_ENDPOINT + "/builds/{buildId}",
			CancelBuild,
		},
		{
			"DeleteBuild",
			"DELETE",
			server.API_ENDPOINT + "/builds/{buildId}",
			DeleteBuild,
		},
		{
			"GetBuildResult",
			"GET",
			server.API_ENDPOINT + "/builds/{buildId}/result",
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
