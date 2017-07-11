package dao

import (
	"github.com/michaelwmerritt/project-builder/model"
	"github.com/michaelwmerritt/project-builder/database"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func GetAllUsers() (*[]model.User, error) {
	users := &[]model.User{}
	err := getUserCollection().Find(bson.M{}).All(users)
	return users, err
}

func GetUser(userName string) (*model.User, error) {
	user := &model.User{}
	err := getUserCollection().FindId(userName).One(user)
	return user, err
}

func DeleteUser(userName string) error{
	return getUserCollection().RemoveId(userName)
}

func UpdateUser(user model.User) error {
	return getUserCollection().UpdateId(user.Id, user)
}

func getUserCollection() *mgo.Collection {
	return database.PROJECT.DB().C("users")
}