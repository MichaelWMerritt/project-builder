package dao

import (
	"github.com/michaelwmerritt/project-builder/model"
	"github.com/michaelwmerritt/project-builder/database"
	"gopkg.in/mgo.v2/bson"
	"github.com/michaelwmerritt/project-builder/datastore"
)

type ModuleDao struct {
	moduleDatastore datastore.Module
}

func NewModuleDao() ModuleDao {
	return ModuleDao{moduleDatastore:datastore.NewModuleDatastore()}
}

func (moduleDao ModuleDao) GetAllModules() (modules []model.Module, err error) {
	results, err := moduleDao.moduleDatastore.Find(getModuleCollectionProvider(), bson.M{}, 0, 0)
	modules = make([]model.Module, len(results))
	for i, module := range results {
		var m model.Module
		bsonBytes, _ := bson.Marshal(module)
		bson.Unmarshal(bsonBytes, &m)
		modules[i] = m
	}
	return
}

func (moduleDao ModuleDao) GetModule(moduleId string) (module model.Module, err error) {
	result, err := moduleDao.moduleDatastore.FindOne(getModuleCollectionProvider(), bson.M{"_id":moduleId})
	bsonBytes, _ := bson.Marshal(result)
	bson.Unmarshal(bsonBytes, &module)
	return
}

func (moduleDao ModuleDao) DeleteModule(moduleId string) (err error) {
	err = moduleDao.moduleDatastore.Delete(getModuleCollectionProvider(), moduleId)
	return
}

func getModuleCollectionProvider() (collectionProvider database.CollectionProvider) {
	collectionProvider = database.CollectionProvider{
		DbProvider:database.PROJECT,
		CollectionName:"modules",
	}
	return
}