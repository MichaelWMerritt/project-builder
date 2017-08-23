package model

type Release struct {

	Id					string			`bson:"_id" json:"_id"`
	DisplayName			string			`bson:"displayName" json:"displayName"`
	Description 		string			`bson:"description" json:"description"`
	VersionInfo 		VersionInfo 	`bson:"versionInfo" json:"versionInfo"`
	RepoType    		RepoType    	`bson:"repoType" json:"repoType"`

}
