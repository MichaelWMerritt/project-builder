package controller

import (
	"encoding/json"
	//"fmt"
	"github.com/gorilla/mux"
	"github.com/michaelwmerritt/project-builder/model"
	"net/http"
	"github.com/michaelwmerritt/project-builder/dao"
)

func CreateModuleRoutes() []model.Route {
	return []model.Route{
		{
			"GetAllModules",
			"GET",
			"/modules",
			GetAllModules,
		},
		{
			"GetModule",
			"GET",
			"/modules/{moduleId}",
			GetModule,
		},
	}
}

func GetAllModules(w http.ResponseWriter, r *http.Request) {
	//modules := []model.Module{
	//	{
	//		Id:			"module1",
	//		DisplayName: "Module 1",
	//		VersionInfo: model.VersionInfo{},
	//		Group:       "? don't know yet 1",
	//	},
	//	{
	//		Id:			"module2",
	//		DisplayName: "Module 2",
	//		VersionInfo: model.VersionInfo{},
	//		Group:       "? don't know yet 2",
	//	},
	//}
	modules, err := dao.GetAllModules()
	if err != nil {

	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(modules); err != nil {
		panic(err)
	}
}

func GetModule(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	moduleId := vars["moduleId"]

	module, err := dao.GetModule(moduleId)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	status := http.StatusOK
	if err != nil {
		status = http.StatusBadRequest
	}
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(module); err != nil {
		panic(err)
	}
}

func DeleteModule(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	moduleId := vars["moduleId"]

	dao.DeleteModule(moduleId)
	//w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//w.WriteHeader(http.StatusOK)
	//if err != nil {
	//	w.WriteHeader(http.StatusBadRequest)
	//} else {
	//	w.WriteHeader(http.StatusOK)
	//}
	//json.NewEncoder(w).Encode("Deleted Module")
}
