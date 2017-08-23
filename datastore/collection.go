package datastore

import (
	"gopkg.in/mgo.v2"
	"github.com/michaelwmerritt/project-builder/database"
	"gopkg.in/mgo.v2/bson"
)

type Collection struct {}

func (collection Collection) Find(collectionProvider database.CollectionProvider, query interface{}, skip int, limit int) (results []interface{}, err error) {
	search := func() error {
		return database.WithCollection(collectionProvider.DbProvider, collectionProvider.CollectionName, func(collection *mgo.Collection) error {
			fn := collection.Find(query).Skip(skip).Limit(limit).All(&results)
			if limit < 0 {
				fn = collection.Find(query).Skip(skip).All(&results)
			}
			return fn
		})
	}
	err = search()
	return
}

func (collection Collection) Delete(collectionProvider database.CollectionProvider, id string) (err error) {
	search := func() error {
		return database.WithCollection(collectionProvider.DbProvider, collectionProvider.CollectionName, func(collection *mgo.Collection) error {
			fn := collection.RemoveId(id)
			return fn
		})
	}
	err = search()
	return
}

func (collection Collection) FindOne(collectionProvider database.CollectionProvider, query interface{}) (result interface{}, err error) {
	search := func() error {
		return database.WithCollection(collectionProvider.DbProvider, collectionProvider.CollectionName, func(collection *mgo.Collection) error {
			fn := collection.Find(query).One(&result)
			return fn
		})
	}
	err = search()
	return
}

func (collection Collection) UpdateOne(collectionProvider database.CollectionProvider, query interface{}, object interface{}) (err error) {
	update := func() error {
		return database.WithCollection(collectionProvider.DbProvider, collectionProvider.CollectionName, func(collection *mgo.Collection) error {
			fn := collection.Update(query, object)
			return fn
		})
	}
	err = update()
	return
}

func (collection Collection) UpdateById(collectionProvider database.CollectionProvider, id string, object interface{}) (err error) {
	update := func() error {
		return database.WithCollection(collectionProvider.DbProvider, collectionProvider.CollectionName, func(collection *mgo.Collection) error {
			fn := collection.Update(bson.M{"_id":id}, object)
			return fn
		})
	}
	err = update()
	return
}

func (collection Collection) CreateOne(collectionProvider database.CollectionProvider, object interface{}) (err error) {
	create := func() error {
		return database.WithCollection(collectionProvider.DbProvider, collectionProvider.CollectionName, func(collection *mgo.Collection) error {
			fn := collection.Insert(object)
			return fn
		})
	}
	err = create()
	return
}
