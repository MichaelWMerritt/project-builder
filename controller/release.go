package controller

import (
	"encoding/json"
	//"fmt"
	"github.com/gorilla/mux"
	"github.com/michaelwmerritt/project-builder/model"
	"net/http"
	"github.com/michaelwmerritt/project-builder/dao"
)

func CreateReleaseRoutes() []model.Route {
	return []model.Route{
		{
			"GetAllReleases",
			"GET",
			"/releases",
			GetAllReleases,
		},
		{
			"GetRelease",
			"GET",
			"/releases/{releaseId}",
			GetRelease,
		},
	}
}

func GetAllReleases(w http.ResponseWriter, r *http.Request) {
	//releases := []model.Release{
	//	{
	//		VersionInfo: model.VersionInfo{DisplayName: "release1"},
	//		RepoType:    model.GIT,
	//	},
	//	{
	//		VersionInfo: model.VersionInfo{DisplayName: "release2"},
	//		RepoType:    model.SVN,
	//	},
	//}
	releases := dao.GetAllReleases()

	if len(releases) == 0 {

	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(releases); err != nil {
		panic(err)
	}
}

func GetRelease(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	releaseId := vars["releaseId"]

	release := dao.GetRelease(releaseId)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(release); err != nil {
		panic(err)
	}
}
