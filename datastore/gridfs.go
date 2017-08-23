package datastore

import (
	"github.com/michaelwmerritt/project-builder/database"
	"gopkg.in/mgo.v2"
)

type GridFS struct {}

func (gridFs GridFS) Delete(collectionProvider database.CollectionProvider, objectId string) (err error) {
	search := func() error {
		return database.WithGridFS(collectionProvider.DbProvider, collectionProvider.CollectionName, func(gridFs *mgo.GridFS) error {
			fn := gridFs.RemoveId(objectId)
			return fn
		})
	}
	err = search()
	return
}

