package dao

import (
	"github.com/michaelwmerritt/project-builder/model"
	"github.com/michaelwmerritt/project-builder/database"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func GetAllReleases() ([]model.Release, error) {
	releases := []model.Release{}
	err := getReleaseCollection().Find(bson.M{}).All(&releases)
	return releases, err
}

func GetRelease(releaseId string) (model.Release, error) {
	release := model.Release{}
	err := getReleaseCollection().FindId(releaseId).One(&release)
	return release, err
}

func DeleteRelease(releaseId string) error {
	return getReleaseCollection().RemoveId(releaseId)
}

func getReleaseCollection() *mgo.Collection {
	return database.PROJECT.DB().C("release")
}