package builder

import (
	"github.com/michaelwmerritt/project-builder/dao"
	"github.com/michaelwmerritt/project-builder/model"
	"gopkg.in/mgo.v2/bson"
	"github.com/satori/go.uuid"
	"os"
	"github.com/pierrre/archivefile/zip"
	"io/ioutil"
	"path/filepath"
	"fmt"
	"regexp"
	"io"
	"log"
)

type BuildManager struct{}

var (
	projectPath = "/opt/project_builder/"
	releaseDao = dao.NewReleaseDao()
	moduleDao = dao.NewModuleDao()
	buildDao = dao.NewBuildDao()
)

func NewBuildManager() BuildManager {
	return BuildManager{}
}

func (buildManager BuildManager) ExecuteBuild(build model.Build) (err error) {
	err = buildDao.CreateBuildReference(build); if err != nil {
		log.Panic(err)
		return
	}
	saveBuild(build, model.IN_PROGRESS)

	release := model.Release{}
	modules := []model.Module{}

	release, err = releaseDao.GetRelease(build.BuildReference.Release.Id); if err != nil {
		handleBuildFailure(build, build.BuildReference.Release.Id + " not found")
		return
	}

	//Get Build Support Modules
	query := bson.M{"versionInfo.version":release.VersionInfo.Version}
	query["buildInfrastructure"] = true
	buildSupportModules, err := moduleDao.GetAllModules(query); if err != nil {
		handleBuildFailure(build,"Build Support Modules not found")
		return
	}
	modules = append(modules, buildSupportModules...)

	//Get Selected Modules
	for _, m := range build.BuildReference.Modules {
		module, err := moduleDao.GetModule(m.Id); if err != nil {
			handleBuildFailure(build, m.Id + " not found")
			return err
		}
		modules = append(modules, module)
	}

	//temp build folder
	tempBuildPath := projectPath + uuid.NewV4().String()
	if _, err := os.Stat(tempBuildPath); os.IsNotExist(err) {
		os.Mkdir(tempBuildPath, os.ModePerm)
	}

	log.Println("*******************************Starting VCS Manager Checkout*******************************")
	vcsManager := NewSvnManager()
	err = vcsManager.Checkout(release, modules, tempBuildPath); if err != nil {
		handleBuildFailure(build, "Unable to checkout modules")
		return
	}
	log.Println("*******************************Finished VCS Manager Checkout*******************************")

	log.Println("*******************************Starting Job Builder*******************************")
	jobBuilder := NewAntBuilder()
	err = jobBuilder.Build(release, modules, tempBuildPath); if err != nil {
		handleBuildFailure(build, "Unable to build project")
		return
	}
	log.Println("*******************************Finished Job Builder*******************************")

	log.Println("*******************************Starting Zipping Results*******************************")
	//folder to hold all build results that will be zipped for user
	zipResultsPath, err := ioutil.TempDir(tempBuildPath, "results")
	if err != nil {
		handleBuildFailure(build, "Unable to create folder 'results'")
		return
	}
	log.Println("*******************************Finished Zipping Results*******************************")

	if build.BuildReference.BuildType == model.DOCKER {
		//TODO do Docker build for each docker buildable in docker infrastructure (maybe execute scripts?)
		// do a docker save
		// build project based on wars or docker
	} else {
		filepath.Walk(tempBuildPath, func(path string, f os.FileInfo, _ error) error {
			if !f.IsDir() {
				r, err := regexp.MatchString(".war", f.Name())
				if err == nil && r {
					src_file, err := os.Open(path + "/" + f.Name()); if err != nil {
						return nil
					}
					defer src_file.Close()

					src_file_stat, err := src_file.Stat(); if err != nil {
						return nil
					}

					if !src_file_stat.Mode().IsRegular() {
						return nil
					}

					dst_file, err := os.Create(zipResultsPath + "/" + f.Name()); if err != nil {
						return nil
					}
					defer dst_file.Close()
					_, err = io.Copy(dst_file, src_file); if err != nil {
						log.Println("Unable to copy file: " + path + "/" + f.Name() + " because: " + err.Error())
					}
				}
			}
			return nil
		})
	}
	//zip and store results in gridfs
	zipAndPersistResults(build, zipResultsPath)
	saveBuild(build, model.COMPLETE)
	cleanUpBuild(build, tempBuildPath)
	return
}

func zipAndPersistResults(build model.Build, path string) (err error) {
	outFilePath := filepath.Join(path, build.Id + ".zip")
	progress := func(archivePath string) {
		fmt.Println(archivePath)
	}
	err = zip.ArchiveFile(path, outFilePath, progress); if err == nil {
		err = buildDao.CreateBuild(build, outFilePath)
	}
	return
}

func handleBuildFailure(build model.Build, message string) {
	build.Message = message
	saveBuild(build, model.FAILED)
}

func cleanUpBuild(build model.Build, path string) (err error) {
	if _, err := os.Stat(path); os.IsExist(err) {
		os.Remove(path)
	}; if err != nil {
		handleBuildFailure(build, "Failed to clean up build: " + build.Id)
	}
	return
}

func saveBuild(build model.Build, status model.BuildStatus) {
	build.Status = status
	buildDao.UpdateBuildReference(build)
}