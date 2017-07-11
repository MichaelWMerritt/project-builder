package model

type UserInfo struct {

	FirstName	string	`bson:"firstName" json:"firstName"`
	LastName	string	`bson:"lastName" json:"lastName"`
	Email		string	`bson:"email" json:"email"`

}
