package dao

import (
	"net/http"
	"gopkg.in/mgo.v2"
)

type BaseDao interface {
	GetCollectionName() string

	GetAll() (*[]interface{}, error)

	Get() (*interface{}, error)

	Delete() (*interface{}, error)

	Create() (error)

	getCollection(r *http.Request) *mgo.Collection
}
