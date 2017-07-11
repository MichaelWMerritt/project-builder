package model

type Module struct {
	Id 			string 		`bson:"_id" json:"_id"`
	DisplayName	string		`bson:"displayName" json:"displayName"`
	VersionInfo VersionInfo `bson:"versionInfo" json:"versionInfo"`
	Group       string      `json:"group"` //id of module to which this belongs
}
