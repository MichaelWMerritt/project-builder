package dao

import "github.com/michaelwmerritt/project-builder/model"

type Module interface {

	GetAllModules() *[]model.Module

	GetModule(moduleId string) *model.Module

	DeleteModule(moduleId string)

}
