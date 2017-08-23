package model

type Build struct {

	Id				string			`bson:"_id" json:"_id"`

	BuildReference BuildReference	`bson:"buildReference" json:"buildReference"`

	Status			BuildStatus		`bson:"status" json:"status"`
	Message			string			`bson:"message" json:"message"`
	Link			string			`bson:"link" json:"link"`

	//additional data (i.e. UtilityId, ipaddress of servers, etc...)
		//this should probably be a map of key to value

}
