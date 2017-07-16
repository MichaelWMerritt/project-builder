package datastore

type Module struct {
	Base
}

func NewModuleDatastore() Module {
	return Module{Base{}}
}
//
//func NewModuleDatastore(session *mgo.Session) *ModuleDatastore {
//	return &ModuleDatastore{collection:database.PROJECT.DB(session).C("modules")}
//}
//
//func (moduleDatastore ModuleDatastore) GetAll(r *http.Request) ([]model.Module, error) {
//	modules := []model.Module{}
//	err := moduleDatastore.collection.Find(bson.M{}).All(&modules)
//	return modules, err
//}
//
//func (moduleDatastore ModuleDatastore) Get(moduleId string, r *http.Request) (model.Module, error) {
//	module := model.Module{}
//	err := moduleDatastore.collection.FindId(moduleId).One(&module)
//	return module, err
//}
//
//func (moduleDatastore ModuleDatastore) Delete(moduleId string, r *http.Request) error {
//	return moduleDatastore.collection.RemoveId(moduleId)
//}

//func getModulesCollection(r *http.Request) *mgo.Collection {
//	db := context.Get(r, server.CONTEXT_DB_KEY).(*mgo.Session)
//	return database.PROJECT.DB(db).C("modules")
//}