package database

import (
	//"fmt"
	//"log"

	"gopkg.in/mgo.v2"

	//"os"
	//"fmt"
	//"encoding/json"
	//"github.com/michaelwmerritt/project-builder/model"
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

	session.SetMode(mgo.Monotonic, true)
	session.SetSafe(&mgo.Safe{})

	return session.DB(databaseName)
}

//func LoadMongoDBConfiguration() model.MongoDBConfig {
//	var mongoDBConfig model.MongoDBConfig
//	mongoDBConfigFile, err := os.Open(mongodb_config_file)
//	defer mongoDBConfigFile.Close()
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//	jsonParser := json.NewDecoder(mongoDBConfigFile)
//	jsonParser.Decode(&mongoDBConfig)
//	return mongoDBConfig
//}
