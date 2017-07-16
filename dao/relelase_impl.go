package dao

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/michaelwmerritt/project-builder/model"
	"github.com/michaelwmerritt/project-builder/database"
	"github.com/michaelwmerritt/project-builder/datastore"
)

type ReleaseDao struct {
	releaseDatastore datastore.Release
}

func NewReleaseDao() ReleaseDao {
	return ReleaseDao{releaseDatastore:datastore.NewReleaseDatastore()}
}

func (releaseDao ReleaseDao) GetAllReleases() (releases []model.Release, err error) {
	results, err := releaseDao.releaseDatastore.Find(getReleaseCollectionProvider(), bson.M{}, 0, 0)
	releases = make([]model.Release, len(results))
	for i, release := range results {
		var r model.Release
		bsonBytes, _ := bson.Marshal(release)
		bson.Unmarshal(bsonBytes, &r)
		releases[i] = r
	}
	return
}

func (releaseDao ReleaseDao) GetRelease(releaseId string) (release model.Release, err error) {
	//release := model.Release{}
	//r, err := getReleaseCollection(r).FindId(releaseId).One(&release)
	//return release, err
	result, err := releaseDao.releaseDatastore.FindOne(getReleaseCollectionProvider(), bson.M{"_id": releaseId})
	bsonBytes, _ := bson.Marshal(result)
	bson.Unmarshal(bsonBytes, &release)
	return
}

func (releaseDao ReleaseDao) DeleteRelease(releaseId string) (err error) {
	err = releaseDao.releaseDatastore.Delete(getReleaseCollectionProvider(), releaseId)
	return
}

func getReleaseCollectionProvider() (collectionProvider database.CollectionProvider) {
	collectionProvider = database.CollectionProvider{
		DbProvider:database.PROJECT,
		CollectionName:"release",
	}
	return
}

//func getReleaseCollection(r *http.Request) *mgo.Collection {
//	db := context.Get(r, server.CONTEXT_DB_KEY).(*mgo.Session)
//	return database.PROJECT.DB(db).C("release")
//}