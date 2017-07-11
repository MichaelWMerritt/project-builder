package model

type Release struct {
	VersionInfo VersionInfo `json:"versionInfo"`
	RepoType    RepoType    `json:"repoType"`
}
