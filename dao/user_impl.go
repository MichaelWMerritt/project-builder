package dao

import (
	"github.com/michaelwmerritt/project-builder/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func GetAllUsers() []model.User {
	users := []model.User{}
	getUserCollection().Find(bson.M{}).All(users)
	return users
}

func GetUser(userName string) model.User {
	user := model.User{}
	getUserCollection().FindId(userName).One(user)
	return user
}

func DeleteUser(userName string) {
	getUserCollection().RemoveId(userName)
}

func UpdateUser(user model.User) {
	getUserCollection().UpdateId(user.Id, user)
}

func getUserCollection() *mgo.Collection {
	return model.PROJECT.DB().C("users")
}