package dao

import "github.com/michaelwmerritt/project-builder/model"

type Build interface {

	GetAllBuilds() *[]model.Build

	GetBuild(buildId string) *model.Build

	DeleteBuild(buildId string)

	CreateBuild(build model.Build) *model.Build

}
