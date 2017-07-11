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
		{
			"DeleteModule",
			"DELETE",
			"/modules/{moduleId}",
			DeleteModule,
		},
	}
}

func GetAllModules(w http.ResponseWriter, r *http.Request) {
	panic("testing")
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
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(module); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}
}

func DeleteModule(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	moduleId := vars["moduleId"]

	if err := dao.DeleteModule(moduleId); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	//w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//w.WriteHeader(http.StatusOK)
	//if err != nil {
	//	w.WriteHeader(http.StatusBadRequest)
	//} else {
	//	w.WriteHeader(http.StatusOK)
	//}
	//json.NewEncoder(w).Encode("Deleted Module")
}

//func NewUserController(s *mgo.Session) *UserController {
//	return &UserController{s}
//}

//// CreateUser creates a new user resource
//func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
//	// Stub an user to be populated from the body
//	u := models.User{}
//
//	// Populate the user data
//	json.NewDecoder(r.Body).Decode(&u)
//
//	// Add an Id
//	u.Id = bson.NewObjectId()
//
//	// Write the user to mongo
//	uc.session.DB("go_rest_tutorial").C("users").Insert(u)
//
//	// Marshal provided interface into JSON structure
//	uj, _ := json.Marshal(u)
//
//	// Write content-type, statuscode, payload
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(201)
//	fmt.Fprintf(w, "%s", uj)
//}

//
//// Grab id
//id := p.ByName("id")
//
//// Verify id is ObjectId, otherwise bail
//if !bson.IsObjectIdHex(id) {
//w.WriteHeader(404)
//return
//}