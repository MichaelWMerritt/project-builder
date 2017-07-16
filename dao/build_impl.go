package dao

//func GetAllBuilds(r *http.Request) (*[]model.Build, error) {
//	//TODO: get all builds from database
//	return &[]model.Build{}, nil
//}
//
//func GetBuild(buildId string, r *http.Request) (*model.Build, error) {
//	//TODO: get build based on id from database
//	return &model.Build{}, nil
//}
//
//func DeleteBuild(buildId string, r *http.Request) error {
//	return getCollection(r).RemoveId(buildId)
//}
//
//func CreateBuild(build model.Build, r *http.Request) (*model.Build, error) {
//	//TODO: create build in database
//	return &model.Build{}, nil
//}
//
//func getCollection(r *http.Request) *mgo.GridFS {
//	db := context.Get(r, server.CONTEXT_DB_KEY).(*mgo.Session)
//	return database.BUILD.DB(db).GridFS("builds")
//}