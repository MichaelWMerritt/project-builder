package database

import (
	"gopkg.in/mgo.v2"
	"log"
)
var (
	mgoSession *mgo.Session
)

func getSession () *mgo.Session {
	if mgoSession == nil {
		var err error
		mgoSession, err = mgo.Dial("localhost")
		if err != nil {
			log.Fatal("Unable to connect to database at: " + "localhost")
		}
		mgoSession.SetMode(mgo.Monotonic, true)
		mgoSession.SetSafe(&mgo.Safe{})
	}
	return mgoSession.Copy()
}

func WithCollection(dbProvider DBProvider, collectionName string, s func(*mgo.Collection) error) error {
	session := getSession()
	defer session.Close()
	collection := dbProvider.DB(session).C(collectionName)
	return s(collection)
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
