package database

import (
	//"fmt"
	//"log"

	"gopkg.in/mgo.v2"
	"os"
	"fmt"
	"encoding/json"
	"github.com/michaelwmerritt/project-builder/model"
)

const (
	mongodb_url = "localhost"
	mongodb_config_file = "/path/to/config/file"
)

func GetDB(databaseName string) *mgo.Database {
	session, err := mgo.Dial(mongodb_url)
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

func LoadMongoDBConfiguration() model.MongoDBConfig {
	var mongoDBConfig model.MongoDBConfig
	mongoDBConfigFile, err := os.Open(mongodb_config_file)
	defer mongoDBConfigFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(mongoDBConfigFile)
	jsonParser.Decode(&mongoDBConfig)
	return mongoDBConfig
}
