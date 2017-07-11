package dao

import "github.com/michaelwmerritt/project-builder/model"

type User interface {

	GetAllUsers() (*[]model.User, error)

	GetUser(userName string) (*model.User, error)

	DeleteUser(userName string) error

	UpdateUser(user model.User) error

}
