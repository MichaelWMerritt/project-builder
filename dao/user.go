package dao

import "github.com/michaelwmerritt/project-builder/model"

type User interface {

	GetAllUsers() *[]model.User

	GetUser(userName string) *model.User

	DeleteUser(userName string)

	UpdateUser(user model.User) *model.User

}
