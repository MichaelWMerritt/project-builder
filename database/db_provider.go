package database

import (
	"gopkg.in/mgo.v2"
)

type DBProvider uint

func (dbProvider DBProvider) DB() *mgo.Database {
	return GetDB(databases[dbProvider])
}

const (

	BUILD DBProvider = iota
	PROJECT DBProvider = iota

)

var databases = []string{"build", "project"}
