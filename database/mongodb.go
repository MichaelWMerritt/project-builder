package database

import (
	"gopkg.in/mgo.v2"
	"log"
)

const (
	HOST = "localhost"
)

var (
	mgoSession *mgo.Session
)

func getSession () *mgo.Session {
	if mgoSession == nil {
		var err error
		mgoSession, err = mgo.Dial(HOST)
		if err != nil {
			log.Fatal("Unable to connect to database at: " + HOST)
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

func WithGridFS(dbProvider DBProvider, gridFsPrefix string, s func(fs *mgo.GridFS) error) error {
	session := getSession()
	defer session.Close()
	gridFs := dbProvider.DB(session).GridFS(gridFsPrefix)
	return s(gridFs)
}

func WithGridFSFile(dbProvider DBProvider, gridFsPrefix string, s func(fs *mgo.GridFS) (*mgo.GridFile, error)) (*mgo.GridFile, error) {
	session := getSession()
	defer session.Close()
	gridFs := dbProvider.DB(session).GridFS(gridFsPrefix)
	return s(gridFs)
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
