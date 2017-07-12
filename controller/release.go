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
		{
			"DeleteRelease",
			"DELETE",
			"/releases/{releaseId}",
			DeleteRelease,
		},
	}
}

func GetAllReleases(w http.ResponseWriter, r *http.Request) {
	releases, err := dao.GetAllReleases()

	if err != nil {

	}
	if len(releases) == 0 {

	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(releases); err != nil {
		panic(err)
	}
}

func GetRelease(w http.ResponseWriter, r *http.Request) {
	releaseId := getReleaseId(r)

	release, err := dao.GetRelease(releaseId)
	if err != nil {

	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(release); err != nil {
		panic(err)
	}
}

func DeleteRelease(w http.ResponseWriter, r *http.Request) {
	releaseId := getReleaseId(r)

	dao.DeleteRelease(releaseId)
}

func getReleaseId(r *http.Request) string {
	return mux.Vars(r)["releaseId"]
}
