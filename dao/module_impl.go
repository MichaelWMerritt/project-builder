package dao

import (
	"github.com/michaelwmerritt/project-builder/model"
	"github.com/michaelwmerritt/project-builder/database"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
)

func GetAllModules() ([]model.Module, error) {
	modules := []model.Module{}
	err := getModulesCollection().Find(bson.M{}).All(&modules)
	return modules, err
	//if err != nil {
	//	panic(err)
	//}
	//return modules
}

func GetModule(moduleId string) (model.Module, error) {
	module := model.Module{}
	err := getModulesCollection().FindId(moduleId).One(&module)
	//err := getModulesCollection().FindId(moduleId).One(&module)
	//if err != nil {
	//	panic(err)
	//}
	//return module
	return module, err
}

func DeleteModule(moduleId string) error {
	return getModulesCollection().RemoveId(moduleId)
}

func getModulesCollection() *mgo.Collection {
	return database.PROJECT.DB().C("modules")
}