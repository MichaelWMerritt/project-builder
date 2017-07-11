package dao

import "github.com/michaelwmerritt/project-builder/model"

type Module interface {

	GetAllModules() ([]model.Module, error)

	GetModule(moduleId string) (model.Module, error)

	DeleteModule(moduleId string) error

}
