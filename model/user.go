package model

type User struct {

	Id        string   `bson:"_id" json:"_id"`
	UserName  string   `bson:"userName" json:"userName"`
	Password  string   `bson:"password" json:"password"`
	UserInfo  UserInfo `bson:"userInfo" json:"userInfo"`

}
