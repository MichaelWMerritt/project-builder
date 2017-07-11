package database

import (
	//"fmt"
	//"log"

	"gopkg.in/mgo.v2"
)

const (
	mognodb_url = "localhost"
)

func GetDB(databaseName string) *mgo.Database {
	session, err := mgo.Dial(mognodb_url)
	if err != nil {
		panic(err)
	}
	//defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	return session.DB(databaseName)
	//c := session.DB("test").C("people")
	//err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
	//	&Person{"Cla", "+55 53 8402 8510"})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//result := Person{}
	//err = c.Find(bson.M{"name": "Ale"}).One(&result)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println("Phone:", result.Phone)
}
