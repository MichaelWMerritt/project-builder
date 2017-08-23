package model

type BuildSupportScript struct {

	Id 				string								`bson:"_id" json:"_id"`
	DisplayName		string								`bson:"displayName" json:"displayName"`
	Description		string								`bson:"description" json:"description"`
	Url				string								`bson:"url" json:"url"`
	ExecutionPoint	BuildSupportScriptExecutionPoint	`bson:"executionPoint" json:"executionPoint"`
	VersionInfo		VersionInfo							`bson:"versionInfo" json:"versionInfo"`

}
