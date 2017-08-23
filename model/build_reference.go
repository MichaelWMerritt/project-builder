package model

type BuildReference struct {

	DisplayName		string			`bson:"displayName" json:"displayName"`

	DateCreated		string			`bson:"dateCreated" json:"dateCreated"`
	DateCompleted	string			`bson:"dateCompleted" json:"dateCompleted"`

	Release			Release			`bson:"release" json:"release"`
	Modules			[]Module		`bson:"modules" json:"modules"`

	BuildType		BuildType		`bson:"buildType" json:"buildType"`

}
