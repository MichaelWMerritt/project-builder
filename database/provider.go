package database

import (
	"gopkg.in/mgo.v2"
)

type DBProvider uint

func (dbProvider DBProvider) DB(session *mgo.Session) *mgo.Database {
	return session.DB(databases[dbProvider])
}

const (

	BUILD DBProvider = iota
	PROJECT DBProvider = iota

)

var databases = []string{"build", "project"}
