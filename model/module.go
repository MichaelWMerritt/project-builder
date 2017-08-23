package model

type Module struct {
	Id 					string 		`bson:"_id" json:"_id"`
	DisplayName			string		`bson:"displayName" json:"displayName"`
	Description 		string		`bson:"description" json:"description"`
	Dependencies		[]string	`bson:"dependencies" json:"dependencies"`
	BuildModule			bool		`bson:"buildModule" json:"buildModule"`
	BuildFile			string		`bson:"buildFile" json:"buildFile"`
	BuildInfrastructure bool		`bson:"buildInfrastructure" json:"buildInfrastructure"`
	BuilderDependencies bool		`bson:"builderDependencies" json:"builderDependencies"`
	VersionInfo 		VersionInfo `bson:"versionInfo" json:"versionInfo"`
	Group       		string      `bson:"group" json:"group"` //id of module to which this belongs (can be used for plugins, etc...)
}

func (module Module) CreateVCSUrl() (url string) {
	url = module.VersionInfo.Url + "/" + module.Id
	return
}

func (module Module) CreateModulePath(path string) (modulePath string) {
	modulePath = path + "/" + module.Id
	return
}

func (module Module) CreateBuildFilePath(path string) (buildFilePath string) {
	buildFilePath = path + "/" + module.Id + "/" + module.BuildFile
	return
}