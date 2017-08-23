package dao

//func GetAllUsers(r *http.Request) (*[]model.User, error) {
//	users := &[]model.User{}
//	err := getUserCollection(r).Find(bson.M{}).All(users)
//	return users, err
//}
//
//func GetUser(userName string, r *http.Request) (*model.User, error) {
//	user := &model.User{}
//	err := getUserCollection(r).FindId(userName).One(user)
//	return user, err
//}
//
//func DeleteUser(userName string, r *http.Request) error{
//	return getUserCollection(r).RemoveId(userName)
//}
//
//func UpdateUser(user model.User, r *http.Request) error {
//	return getUserCollection(r).UpdateId(user.Id, user)
//}
//
//func getUserCollection(r *http.Request) *mgo.Collection {
//	session := context.Get(r, server.CONTEXT_DB_KEY).(*mgo.Session)
//	return database.PROJECT.DB(session).C("users")
//}