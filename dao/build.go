package dao

import (
	"github.com/michaelwmerritt/project-builder/datastore"
	"github.com/michaelwmerritt/project-builder/model"
	"gopkg.in/mgo.v2/bson"
	"github.com/michaelwmerritt/project-builder/database"
	"os"
	"io"
)

type BuildDao struct {
	buildDatastore datastore.Build
}

func NewBuildDao() BuildDao {
	return BuildDao{buildDatastore:datastore.NewBuildDatastore()}
}

func (buildDao BuildDao) GetAllBuildReferences() (builds []model.Build, err error) {
	results, err := buildDao.buildDatastore.Collection.Find(getBuildCollectionProvider(".files"), bson.M{}, 0, 0)
	builds = make([]model.Build, len(results))
	for i, build := range results {
		var b model.Build
		bsonBytes, _ := bson.Marshal(build)
		bson.Unmarshal(bsonBytes, &b)
		builds[i] = b
	}
	return
}

func (buildDao BuildDao) GetBuildReference(buildId string) (build model.Build, err error) {
	result, err := buildDao.buildDatastore.Collection.FindOne(getBuildCollectionProvider(".files"), bson.M{"_id":buildId})
	bsonBytes, _ := bson.Marshal(result)
	bson.Unmarshal(bsonBytes, &build)
	return
}

func (buildDao BuildDao) DeleteBuild(buildId string) (err error) {
	err = buildDao.buildDatastore.GridFS.Delete(getBuildCollectionProvider(""), buildId)
	return
}

func (buildDao BuildDao) CreateBuild(build model.Build, resultZipPath string) (err error) {
	gridFile, err := buildDao.buildDatastore.CreateFile(getBuildCollectionProvider(""), resultZipPath); if err != nil {
		return
	}
	zipResults, err  := os.Open(resultZipPath); if err != nil {
		return
	}
	defer zipResults.Close()
	_, err = io.Copy(gridFile, zipResults); if err != nil {
		return
	}
	gridFile.SetMeta(build)
	gridFile.SetName(build.Id)

	err = gridFile.Close()
	return
}

func getBuildCollectionProvider(suffix string) (collectionProvider database.CollectionProvider) {
	collectionProvider = database.CollectionProvider{
		DbProvider:database.BUILD,
		CollectionName:"builds" + suffix,
	}
	return
}

func (buildDao BuildDao) CreateBuildReference(build model.Build) (err error) {
	err = buildDao.buildDatastore.Collection.CreateOne(getBuildCollectionProvider(".files"), build)
	return
}

func (buildDao BuildDao) UpdateBuildReference(build model.Build) (err error) {
	err = buildDao.buildDatastore.Collection.UpdateById(getBuildCollectionProvider(".files"), build.Id, build)
	return
}

func (buildDao BuildDao) GetBuildSupportScripts(version string) (buildSupportScriptList []model.BuildSupportScript, err error) {
	query := bson.M{
		"versionInfo.version" : version,
	}
	results, err := buildDao.buildDatastore.Collection.Find(database.CollectionProvider{
		DbProvider: database.BUILD,
		CollectionName: "build.support",
	}, query, 0, 0)
	
	buildSupportScriptList = make([]model.BuildSupportScript, len(results))
	for i, build := range results {
		var bss model.BuildSupportScript
		bsonBytes, _ := bson.Marshal(build)
		bson.Unmarshal(bsonBytes, &bss)
		buildSupportScriptList[i] = bss
	}
	return
}