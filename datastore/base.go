package datastore

import (
	"gopkg.in/mgo.v2"
	"github.com/michaelwmerritt/project-builder/database"
)

type Base struct {}

func (base Base) Find(collectionProvider database.CollectionProvider, q interface{}, skip int, limit int) (results []interface{}, err error) {
	search := func() error {
		return database.WithCollection(collectionProvider.DbProvider, collectionProvider.CollectionName, func(collection *mgo.Collection) error {
			fn := collection.Find(q).Skip(skip).Limit(limit).All(&results)
			if limit < 0 {
				fn = collection.Find(q).Skip(skip).All(&results)
			}
			return fn
		})
	}
	err = search()
	return
}

func (base Base) Delete(collectionProvider database.CollectionProvider, id string) (err error) {
	search := func() error {
		return database.WithCollection(collectionProvider.DbProvider, collectionProvider.CollectionName, func(collection *mgo.Collection) error {
			fn := collection.RemoveId(id)
			return fn
		})
	}
	err = search()
	return
}

func (base Base) FindOne(collectionProvider database.CollectionProvider, q interface{}) (result interface{}, err error) {
	search := func() error {
		return database.WithCollection(collectionProvider.DbProvider, collectionProvider.CollectionName, func(collection *mgo.Collection) error {
			fn := collection.Find(q).One(&result)
			return fn
		})
	}
	err = search()
	return
}

func (base Base) UpdateOne(collectionProvider database.CollectionProvider, q interface{}, object interface{}) (err error) {
	update := func() error {
		return database.WithCollection(collectionProvider.DbProvider, collectionProvider.CollectionName, func(collection *mgo.Collection) error {
			fn := collection.Update(q, object)
			return fn
		})
	}
	err = update()
	return
}

func (base Base) CreateOne(collectionProvider database.CollectionProvider, object interface{}) (err error) {
	create := func() error {
		return database.WithCollection(collectionProvider.DbProvider, collectionProvider.CollectionName, func(collection *mgo.Collection) error {
			fn := collection.Insert(object)
			return fn
		})
	}
	err = create()
	return
}
