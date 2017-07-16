package datastore

type Release struct {
	Base
}

func NewReleaseDatastore() Release {
	return Release{Base{}}
}