package dao

import "github.com/michaelwmerritt/project-builder/model"

type Release interface {

	GetAllReleases() ([]model.Release, error)

	GetRelease(releaseId string) (model.Release, error)

	DeleteRelease(releaseId string) error

}
