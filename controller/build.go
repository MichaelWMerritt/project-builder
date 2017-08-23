package controller

import (
	"github.com/michaelwmerritt/project-builder/model"
	"net/http"
	"github.com/michaelwmerritt/project-builder/server"
	"github.com/michaelwmerritt/project-builder/dao"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/michaelwmerritt/project-builder/builder"
	"time"
	"strconv"
	"fmt"
)

const (
	BUILDS_ENDPOINT = server.API_ENDPOINT + "/builds"
	BUILD_ID = "buildId"
)

var (
	buildDao = dao.NewBuildDao()
)

func CreateBuildRoutes() []model.Route {
	return []model.Route{
		{
			"GetAllBuilds",
			"GET",
			BUILDS_ENDPOINT,
			GetAllBuildReferences,
		},
		{
			"Build",
			"POST",
			BUILDS_ENDPOINT,
			Build,
		},
		{
			"GetBuildStatus",
			"GET",
			BUILDS_ENDPOINT + "/{buildId}",
			GetBuildReference,
		},
		{
			"CancelBuild",
			"PATCH",
			BUILDS_ENDPOINT + "/{buildId}",
			CancelBuild,
		},
		{
			"DeleteBuild",
			"DELETE",
			BUILDS_ENDPOINT + "/{buildId}",
			DeleteBuild,
		},
		{
			"GetBuildResult",
			"GET",
			BUILDS_ENDPOINT + "/{buildId}/result",
			GetBuildResult,
		},
	}
}

func GetAllBuildReferences(w http.ResponseWriter, r *http.Request) {
	builds, err := buildDao.GetAllBuildReferences()
	if err != nil {
		HandleError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(builds); err != nil {
		HandleServerError(w, "build.GetAllBuilds: Failed to convert builds")
	}
}

func Build(w http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var buildReference model.BuildReference
	err := decoder.Decode(&buildReference)
	if err != nil {
		HandleServerError(w, "Unable to parse request")
		return
	}
	t := time.Now()
	buildReference.DateCreated = t.Format("2006-01-02 15:04:05")
	if buildReference.DisplayName == "" {
		buildReference.DisplayName = "New Build @ " + buildReference.DateCreated
	}

	id := fmt.Sprintf("%s", buildReference.BuildType)+ "." + strconv.FormatInt(t.Unix(), 10)
	build := model.Build{
		Id : id,
		BuildReference : buildReference,
		Status : model.CREATED,
		Link : BUILDS_ENDPOINT + "/builds/" + id,
	}
	buildManager := builder.NewBuildManager()
	go buildManager.ExecuteBuild(build)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(buildReference); err != nil {
		HandleServerError(w, "buildDao.build: Failed to encode object buildReference")
	}
}

func GetBuildReference(w http.ResponseWriter, r *http.Request) {
	buildId := getBuildId(r)
	buildReference, err := buildDao.GetBuildReference(buildId)
	if err != nil {
		HandleNotFoundError(w, err, buildId)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(buildReference); err != nil {
		HandleServerError(w, "buildDao.getBuildReference: Failed to convert module " + buildId)
	}
}

func CancelBuild(w http.ResponseWriter, r *http.Request) {

}

func DeleteBuild(w http.ResponseWriter, r *http.Request) {
	buildId := getBuildId(r)
	if err := moduleDao.DeleteModule(buildId); err != nil {
		HandleNotFoundError(w, err, buildId)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetBuildResult(w http.ResponseWriter, r *http.Request) {
	//TODO: get result from gridfs if status is complete, otherwise return message not done or something like status
}

func getBuildId(r *http.Request) string {
	return mux.Vars(r)[BUILD_ID]
}
