package dao

import "github.com/michaelwmerritt/project-builder/model"

type Build interface {

	GetAllBuilds() (*[]model.Build, error)

	GetBuild(buildId string) (*model.Build, error)

	DeleteBuild(buildId string) error

	CreateBuild(build model.Build) (*model.Build, error)

}
