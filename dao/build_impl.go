package dao

import (
	"github.com/michaelwmerritt/project-builder/model"
	"gopkg.in/mgo.v2"
)

func GetAllBuilds() *[]model.Build {
	//TODO: get all builds from database
	return &[]model.Build{}
}

func GetBuild(buildId string) *model.Build {
	//TODO: get build based on id from database
	return &model.Build{}
}

func DeleteBuild(buildId string) {
	err := getCollection().RemoveId(buildId)
	if err != nil {
		panic(err)
	}
}

func CreateBuild(build model.Build) *model.Build {
	//TODO: create build in database
	return &model.Build{}
}

func getCollection() *mgo.GridFS {
	return model.BUILD.DB().GridFS("builds")
}