package dao

import (
	"github.com/michaelwmerritt/project-builder/model"
	"github.com/michaelwmerritt/project-builder/database"
	"gopkg.in/mgo.v2"
)

func GetAllBuilds() (*[]model.Build, error) {
	//TODO: get all builds from database
	return &[]model.Build{}, nil
}

func GetBuild(buildId string) (*model.Build, error) {
	//TODO: get build based on id from database
	return &model.Build{}, nil
}

func DeleteBuild(buildId string) error {
	return getCollection().RemoveId(buildId)
}

func CreateBuild(build model.Build) (*model.Build, error) {
	//TODO: create build in database
	return &model.Build{}, nil
}

func getCollection() *mgo.GridFS {
	return database.BUILD.DB().GridFS("builds")
}