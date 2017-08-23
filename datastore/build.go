package datastore

import (
	"github.com/michaelwmerritt/project-builder/database"
	"gopkg.in/mgo.v2"
)

type Build struct {
	Collection
	GridFS
}

func NewBuildDatastore() Build {
	return Build{Collection{}, GridFS{}}
}

func (buildDatastore Build) CreateFile(collectionProvider database.CollectionProvider, filePath string) (file *mgo.GridFile, err error) {
	create := func() (*mgo.GridFile, error) {
		return database.WithGridFSFile(collectionProvider.DbProvider, collectionProvider.CollectionName, func(fs *mgo.GridFS) (file *mgo.GridFile, err error) {
			file, err = fs.Create(filePath)
			return
		})
	}
	file, err = create()
	return
}

