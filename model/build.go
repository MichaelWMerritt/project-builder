package model

type Build struct {

	Id				string			`bson:"_id" json:"_id"`
	DisplayName		string			`bson:"displayName" json:"displayName"`
	Status			string			`bson:"status" json:"status"`

	Release			Release			`bson:"release" json:"release"`
	Modules			[]Module		`bson:"modules" json:"modules"`

	//additional data (i.e. UtilityId, ipaddress of servers, etc...)
		//this should probably be a map of key to value

}
