package model

import (
	"gopkg.in/mgo.v2"
	"github.com/michaelwmerritt/project-builder/database"
)

type DBProvider uint

func (dbProvider DBProvider) DB() *mgo.Database {
	return database.GetDB(databases[dbProvider])
}

const (

	BUILD DBProvider = iota
	PROJECT DBProvider = iota

)

var databases = []string{"build", "project"}
