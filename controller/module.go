package controller

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/michaelwmerritt/project-builder/model"
	"github.com/michaelwmerritt/project-builder/dao"
	"github.com/michaelwmerritt/project-builder/server"
)

var (
	moduleDao = dao.NewModuleDao()
)

func CreateModuleRoutes() []model.Route {
	return []model.Route{
		{
			"GetAllModules",
			"GET",
			server.API_ENDPOINT + "/modules",
			GetAllModules,
		},
		{
			"GetModule",
			"GET",
			server.API_ENDPOINT + "/modules/{moduleId}",
			GetModule,
		},
		{
			"DeleteModule",
			"DELETE",
			server.API_ENDPOINT + "/modules/{moduleId}",
			DeleteModule,
		},
	}
}

func GetAllModules(w http.ResponseWriter, r *http.Request) {
	modules, err := moduleDao.GetAllModules()
	if err != nil {
		HandleError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(modules); err != nil {
		HandleServerError(w, "module.GetAllModules: Failed to convert modules")
	}
}

func GetModule(w http.ResponseWriter, r *http.Request) {
	moduleId := getModuleId(r)
	module, err := moduleDao.GetModule(moduleId)
	if err != nil {
		HandleNotFoundError(w, err, moduleId)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(module); err != nil {
		HandleServerError(w, "module.GetModule: Failed to convert module " + moduleId)
	}
}

func DeleteModule(w http.ResponseWriter, r *http.Request) {
	moduleId := getModuleId(r)
	if err := moduleDao.DeleteModule(moduleId); err != nil {
		HandleNotFoundError(w, err, moduleId)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func getModuleId(r *http.Request) string {
	return mux.Vars(r)["moduleId"]
}