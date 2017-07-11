package model

type User struct {
	Id        string `json:"_id"`
	UserName  string `json:"userName"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
