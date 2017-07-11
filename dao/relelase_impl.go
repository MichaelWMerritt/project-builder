package dao

import (
	"github.com/michaelwmerritt/project-builder/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func GetAllReleases() *[]model.Release {
	releases := &[]model.Release{}
	getReleaseCollection().Find(bson.M{}).All(releases)
	return releases
}

func GetRelease(releaseId string) *model.Release {
	release := &model.Release{}
	getReleaseCollection().FindId(releaseId).One(release)
	return release
}

func DeleteRelease(releaseId string) {
	getReleaseCollection().RemoveId(releaseId)
}

func getReleaseCollection() *mgo.Collection {
	return model.PROJECT.DB().C("release")
}