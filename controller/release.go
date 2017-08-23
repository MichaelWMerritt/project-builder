package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/michaelwmerritt/project-builder/model"
	"net/http"
	"github.com/michaelwmerritt/project-builder/dao"
	"github.com/michaelwmerritt/project-builder/server"
)

const (
	RELEASES_ENDPOINT = server.API_ENDPOINT + "/releases"
	RELEASE_ID  = "releaseId"
)

var (
	releaseDao = dao.NewReleaseDao()
)

func CreateReleaseRoutes() []model.Route {
	return []model.Route{
		{
			"GetAllReleases",
			"GET",
			RELEASES_ENDPOINT,
			GetAllReleases,
		},
		{
			"GetRelease",
			"GET",
			RELEASES_ENDPOINT + "/{releaseId}",
			GetRelease,
		},
		{
			"DeleteRelease",
			"DELETE",
			RELEASES_ENDPOINT + "/{releaseId}",
			DeleteRelease,
		},
	}
}

func GetAllReleases(w http.ResponseWriter, r *http.Request) {
	releases, err := releaseDao.GetAllReleases()
	if err != nil {
		HandleError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(releases); err != nil {
		HandleServerError(w, "release.GetAllReleases: Failed to convert releases")
	}
}

func GetRelease(w http.ResponseWriter, r *http.Request) {
	releaseId := getReleaseId(r)
	release, err := releaseDao.GetRelease(releaseId)
	if err != nil {
		HandleNotFoundError(w, err, releaseId)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(release); err != nil {
		HandleServerError(w, "release.GetRelease: Failed to convert release " + releaseId)
	}
}

func DeleteRelease(w http.ResponseWriter, r *http.Request) {
	releaseId := getReleaseId(r)
	if err := releaseDao.DeleteRelease(releaseId); err != nil {
		HandleNotFoundError(w, err, releaseId)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func getReleaseId(r *http.Request) string {
	return mux.Vars(r)[RELEASE_ID]
}
