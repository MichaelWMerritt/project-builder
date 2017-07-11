package dao

import "github.com/michaelwmerritt/project-builder/model"

type Release interface {

	GetAllReleases() *[]model.Release

	GetRelease(releaseId string) *model.Release

	DeleteRelease(releaseId string)

}
